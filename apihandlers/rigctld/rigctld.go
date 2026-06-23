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

type ValuePayload struct {
	NewValue string `json:"newValue"`
}

type TransceiverConfig struct {
	ID            int    `json:"id"`
	Model         int    `json:"model"`
	Device        string `json:"device"`
	Baudrate      int    `json:"baudrate"`
	Port          int    `json:"port"`
	ServiceStatus string `json:"status"`
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
	return strings.TrimSpace(string(output)) == "active"
}

// --- Service Lifecycle Management ---

func HandleStartRigctld(w http.ResponseWriter, r *http.Request) {
	trxID, err := strconv.Atoi(r.PathValue("trx_id"))
	if err != nil || trxID <= 0 {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid TRX ID"})
		return
	}

	if !isTrxIDDefined(trxID) {
		WriteJSON(w, http.StatusNotFound, map[string]string{"error": "TRX ID not defined"})
		return
	}

	if isRigctldInstanceRunning(trxID) {
		WriteJSON(w, http.StatusConflict, map[string]string{"error": "Service already running"})
		return
	}

	serviceName := fmt.Sprintf("rigctld@%d.service", trxID)
	if err := exec.Command("sudo", "systemctl", "start", serviceName).Run(); err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to start service"})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"status": "success", "message": serviceName + " started"})
}

func HandleStopRigctld(w http.ResponseWriter, r *http.Request) {
	trxID, err := strconv.Atoi(r.PathValue("trx_id"))
	if err != nil || trxID <= 0 {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid TRX ID"})
		return
	}

	if !isRigctldInstanceRunning(trxID) {
		WriteJSON(w, http.StatusConflict, map[string]string{"error": "Service not running"})
		return
	}

	serviceName := fmt.Sprintf("rigctld@%d.service", trxID)
	if err := exec.Command("sudo", "systemctl", "stop", serviceName).Run(); err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to stop service"})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"status": "success", "message": serviceName + " stopped"})
}

func HandleListRigs(w http.ResponseWriter, r *http.Request) {
	jsonPath := "/etc/hamlib_rest_api/rigctld.json"
	file, err := os.ReadFile(jsonPath)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to read config"})
		return
	}

	var trxs []TransceiverConfig
	if err := json.Unmarshal(file, &trxs); err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to parse config"})
		return
	}

	for i := range trxs {
		if isRigctldInstanceRunning(trxs[i].ID) {
			trxs[i].ServiceStatus = "running"
		} else {
			trxs[i].ServiceStatus = "stopped"
		}
	}

	if trxs == nil {
		trxs = []TransceiverConfig{}
	}

	WriteJSON(w, http.StatusOK, trxs)
}
