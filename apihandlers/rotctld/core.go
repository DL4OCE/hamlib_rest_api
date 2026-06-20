package rotctld

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

// RotatorConfig holds the setup data for a single rotator daemon
type RotatorConfig struct {
	ID       int    `json:"id"`
	Model    int    `json:"model"`
	Device   string `json:"device"`
	Baudrate int    `json:"baudrate"`
	Port     int    `json:"port"`
}

var (
	activeRotators []RotatorConfig
	rotConfigMu    sync.RWMutex
)

// LoadRotatorConfig reads the rotator configurations from a JSON file
func LoadRotatorConfig(filePath string) error {
	rotConfigMu.Lock()
	defer rotConfigMu.Unlock()

	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read rotator config file: %w", err)
	}

	var configs []RotatorConfig
	if err := json.Unmarshal(file, &configs); err != nil {
		return fmt.Errorf("failed to parse rotator json data: %w", err)
	}

	activeRotators = configs
	return nil
}

// GetRotators returns a thread-safe copy of all configured rotators
func GetRotators() []RotatorConfig {
	rotConfigMu.RLock()
	defer rotConfigMu.RUnlock()

	dst := make([]RotatorConfig, len(activeRotators))
	copy(dst, activeRotators)
	return dst
}

// PollRotctlDaemon connects to the backend rotctld TCP port and sends a raw command
func PollRotctlDaemon(rotID int, command string) ([]string, error) {
	var rot *RotatorConfig
	for _, r := range GetRotators() {
		if r.ID == rotID {
			rot = &r
			break
		}
	}
	if rot == nil {
		return nil, fmt.Errorf("rotator with id=%d not defined", rotID)
	}

	// Establish TCP network connection to the respective rotctld daemon
	address := net.JoinHostPort("127.0.0.1", strconv.Itoa(rot.Port))
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("connection to rotctld failed on %s: %w", address, err)
	}
	defer conn.Close()

	// Ensure the command ends with a newline character as required by hamlib
	if !strings.HasSuffix(command, "\n") {
		command += "\n"
	}

	_, err = conn.Write([]byte(command))
	if err != nil {
		return nil, fmt.Errorf("failed to write command to daemon: %w", err)
	}

	var lines []string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		// Break early for non-dump commands since single commands return immediate results
		if !strings.HasPrefix(command, "dump_state") && !strings.HasPrefix(command, "_") {
			break
		}
	}

	if err := scanner.Err(); err != nil {
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
