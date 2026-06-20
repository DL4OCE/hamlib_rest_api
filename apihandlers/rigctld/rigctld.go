package rigctld

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type TransceiverConfig struct {
	ID int `json:"id"`
}

type TransceiverExtendedConfig struct {
	ID        int    `json:"id"`
	Model     string `json:"model"`
	Device    string `json:"device"`
	Baud      string `json:"baud"`
	Port      string `json:"port"`
	IsRunning bool   `json:"is_running"`
}

func isTrxIDDefined(trxID int) bool {
	jsonPath := "/etc/hamlib_rest_api/rigctld.json"
	file, err := os.ReadFile(jsonPath)
	if err != nil {
		return false
	}
	var trxs []TransceiverConfig
	if err := json.Unmarshal(file, &trxs); err != nil {
		return false
	}
	for _, trx := range trxs {
		if trx.ID == trxID {
			return true
		}
	}
	return false
}

func isRigctldInstanceRunning(trxID int) bool {
	serviceName := fmt.Sprintf("rigctld@%d.service", trxID)

	cmd := exec.Command("systemctl", "is-active", serviceName)
	output, _ := cmd.Output()

	status := strings.TrimSpace(string(output))
	return status == "active"
}

// POST /trx/{trx_id}/start
func HandleStartRigctld(w http.ResponseWriter, r *http.Request) {
	trxIDStr := r.PathValue("trx_id")
	trxID, err := strconv.Atoi(trxIDStr)
	if err != nil || trxID <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid TRX ID"})
		return
	}

	if !isTrxIDDefined(trxID) {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": fmt.Sprintf("TRX ID %d not defined in rigctld.json", trxID),
		})
		return
	}

	if isRigctldInstanceRunning(trxID) {
		writeJSON(w, http.StatusConflict, map[string]string{
			"error": fmt.Sprintf("rigctld for TRX ID %d is already running", trxID),
		})
		return
	}

	serviceName := fmt.Sprintf("rigctld@%d.service", trxID)
	cmd := exec.Command("sudo", "systemctl", "start", serviceName)

	if err := cmd.Run(); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Error starting %s: %v", serviceName, err),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Service %s has been started, successfully", serviceName),
	})
}

// POST /trx/{trx_id}/stop
func HandleStopRigctld(w http.ResponseWriter, r *http.Request) {
	trxIDStr := r.PathValue("trx_id")
	trxID, err := strconv.Atoi(trxIDStr)
	if err != nil || trxID <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid TRX ID"})
		return
	}

	if !isTrxIDDefined(trxID) {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": fmt.Sprintf("TRX ID %d not defined in rigctld.json", trxID),
		})
		return
	}

	if !isRigctldInstanceRunning(trxID) {
		writeJSON(w, http.StatusConflict, map[string]string{
			"error": fmt.Sprintf("rigctld for TRX ID %d not running", trxID),
		})
		return
	}

	serviceName := fmt.Sprintf("rigctld@%d.service", trxID)
	cmd := exec.Command("sudo", "systemctl", "stop", serviceName)

	if err := cmd.Run(); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Error stopping %s: %v", serviceName, err),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Service %s has been stopped, successfully", serviceName),
	})
}

// GET /trxs -> list_rigs
func HandleListRigs(w http.ResponseWriter, r *http.Request) {
	jsonPath := "/etc/hamlib_rest_api/rigctld.json"

	file, err := os.ReadFile(jsonPath)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Failed to read transceiver configuration: %v", err),
		})
		return
	}

	var trxs []TransceiverExtendedConfig
	if err := json.Unmarshal(file, &trxs); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Failed to parse configuration JSON: %v", err),
		})
		return
	}

	for i := range trxs {
		trxs[i].IsRunning = isRigctldInstanceRunning(trxs[i].ID)
	}

	// Return the enriched list
	writeJSON(w, http.StatusOK, trxs)
}
