package rotctld

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// // RotatorConfig holds the setup data for a single rotator daemon
// type RotatorConfig struct {
// 	ID       int    `json:"id"`
// 	Model    int    `json:"model"`
// 	Device   string `json:"device"`
// 	Baudrate int    `json:"baudrate"`
// 	Port     int    `json:"port"`
// }

// var (
// 	activeRotators []RotatorConfig
// 	rotConfigMu    sync.RWMutex
// )

type RotatorConfig struct {
	ID            int    `json:"id"`
	Model         int    `json:"model"`
	Device        string `json:"device"`
	Baudrate      int    `json:"baudrate"`
	Port          int    `json:"port"`
	ServiceStatus string `json:"status"` // Wird beim Laden des Zustands gesetzt
}

// LoadRotatorConfig reads the rotator configurations from a JSON file
func LoadRotatorConfig(filePath string) error {
	// rotConfigMu.Lock()
	// defer rotConfigMu.Unlock()

	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read rotator config file: %w", err)
	}

	var configs []RotatorConfig
	if err := json.Unmarshal(file, &configs); err != nil {
		return fmt.Errorf("failed to parse rotator json data: %w", err)
	}

	// activeRotators = configs
	return nil
}

// GetRotators returns a thread-safe copy of all configured rotators
// func GetRotators() []RotatorConfig {
// read list of rotators from file

// rotConfigMu.RLock()
// defer rotConfigMu.RUnlock()

// dst := make([]RotatorConfig, len(activeRotators))
// copy(dst, activeRotators)
// return dst
// }

func GetRotators() ([]RotatorConfig, error) {
	jsonPath := "/etc/hamlib_rest_api/rotctld.json"

	// 1. Datei lesen
	file, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}

	// 2. Parsen
	var rots []RotatorConfig
	if err := json.Unmarshal(file, &rots); err != nil {
		return nil, err
	}

	return rots, nil
}

// GetRotctldStatus queries the live systemd controller state safely using the exit code
func GetRotctldStatus(id int) string {
	cmd := exec.Command("systemctl", "is-active", fmt.Sprintf("rotctld@%d.service", id))
	output, _ := cmd.Output()

	status := strings.TrimSpace(string(output))
	if status == "active" {
		return "RUNNING"
	}
	// "inactive", "activating", "deactivating" etc. maps to STOPPED
	return "STOPPED"
}

func PollRotctlDaemon(rotID int, command string) ([]string, error) {
	rots, err := GetRotators()
	if err != nil {
		return nil, fmt.Errorf("failed to get rotators: %w", err)
	}

	var rot *RotatorConfig
	for _, r := range rots {
		if r.ID == rotID {
			rot = &r
			break
		}
	}
	if rot == nil {
		return nil, fmt.Errorf("rotator with id=%d not defined", rotID)
	}

	address := net.JoinHostPort("127.0.0.1", strconv.Itoa(rot.Port))
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("connection to rotctld failed on %s: %w", address, err)
	}
	defer conn.Close()

	_, err = fmt.Fprintf(conn, "%s\n", strings.TrimSpace(command))
	if err != nil {
		return nil, fmt.Errorf("failed to write command: %w", err)
	}

	var lines []string
	scanner := bufio.NewScanner(conn)

	for {
		conn.SetReadDeadline(time.Now().Add(50 * time.Millisecond))

		if !scanner.Scan() {
			break
		}

		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		if strings.HasPrefix(text, "RPRT") {
			break
		}

		lines = append(lines, text)
	}

	conn.SetReadDeadline(time.Time{})

	if err := scanner.Err(); err != nil && !os.IsTimeout(err) {
		return nil, fmt.Errorf("error reading response from daemon: %w", err)
	}

	return lines, nil
}

// WriteJSON is a reusable HTTP utility for sending standardized JSON payloads
func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// HandleRawCommand targets the direct backend query endpoint: /rotator/{rotator_id}/command
func HandleRawCommand(w http.ResponseWriter, r *http.Request) {
	rotIDStr := r.PathValue("rotator_id")
	rotID, err := strconv.Atoi(rotIDStr)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}

	var body struct {
		Command string `json:"command"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body payload"})
		return
	}

	if strings.TrimSpace(body.Command) == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Command parameter cannot be empty"})
		return
	}

	output, err := PollRotctlDaemon(rotID, body.Command)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"rotator_id": rotID,
		"command":    body.Command,
		"response":   output,
	})
}

// HandleListRotators maps to GET /rotators
// Reads the configuration live from disk, matching the rigctld behavior
func HandleListRotators(w http.ResponseWriter, r *http.Request) {
	jsonPath := "/etc/hamlib_rest_api/rotctld.json"

	// 1. Read file directly on request
	file, err := os.ReadFile(jsonPath)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Failed to read rotator configuration: %v", err),
		})
		return
	}

	var rots []RotatorConfig
	if err := json.Unmarshal(file, &rots); err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Failed to parse rotator configuration JSON: %v", err),
		})
		return
	}

	for i := range rots {
		if IsRotctldInstanceRunning(rots[i].ID) {
			rots[i].ServiceStatus = "running"
		} else {
			rots[i].ServiceStatus = "stopped"
		}
	}

	if rots == nil {
		rots = []RotatorConfig{}
	}

	WriteJSON(w, http.StatusOK, rots)
}

// HandleStartService maps to POST /rotator/{rotator_id}/start
func HandleStartService(w http.ResponseWriter, r *http.Request) {
	rotIDStr := r.PathValue("rotator_id")
	serviceName := fmt.Sprintf("rotctld@%s.service", rotIDStr)

	cmd := exec.Command("sudo", "systemctl", "start", serviceName)
	if err := cmd.Run(); err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"status": "success", "message": serviceName + " started"})
}

// HandleStopService maps to POST /rotator/{rotator_id}/stop
func HandleStopService(w http.ResponseWriter, r *http.Request) {
	rotIDStr := r.PathValue("rotator_id")
	serviceName := fmt.Sprintf("rotctld@%s.service", rotIDStr)

	cmd := exec.Command("sudo", "systemctl", "stop", serviceName)
	if err := cmd.Run(); err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"status": "success", "message": serviceName + " stopped"})
}

func IsRotctldInstanceRunning(rotID int) bool {
	serviceName := fmt.Sprintf("rotctld@%d.service", rotID)
	cmd := exec.Command("systemctl", "is-active", serviceName)
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(output)) == "active"
}
