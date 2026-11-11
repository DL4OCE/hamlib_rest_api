package rigctld

import (
	"encoding/json"
	"fmt"
	"hamlib_rest_api/apihandlers/common"
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

type RigListResponse struct {
	RigctldVersion string              `json:"rigctld_version"`
	Rigs           []TransceiverConfig `json:"rigs"`
}

func GetRigctldVersion() string {
	cmd := exec.Command("rigctld", "--version")
	out, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	lines := strings.Split(string(out), "\n")
	if len(lines) > 0 {
		fields := strings.Fields(lines[0])
		if len(fields) > 0 {
			return fields[2]
		}
	}
	return "unknown"
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

func IsRigctldInstanceRunning(trxID int) bool {
	serviceName := fmt.Sprintf("rigctld@%d.service", trxID)
	cmd := exec.Command("systemctl", "is-active", serviceName)
	output, _ := cmd.Output()
	return strings.TrimSpace(string(output)) == "active"
}

// --- Service Lifecycle Management ---

func HandleStartRigctld(w http.ResponseWriter, r *http.Request) {
	trxID, err := strconv.Atoi(r.PathValue("trx_id"))
	if err != nil || trxID <= 0 {
		common.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid TRX ID"})
		return
	}

	if !isTrxIDDefined(trxID) {
		common.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "TRX ID not defined"})
		return
	}

	if IsRigctldInstanceRunning(trxID) {
		common.WriteJSON(w, http.StatusConflict, map[string]string{"error": "Service already running"})
		return
	}

	serviceName := fmt.Sprintf("rigctld@%d.service", trxID)
	if err := exec.Command("sudo", "systemctl", "start", serviceName).Run(); err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to start service"})
		return
	}

	common.WriteJSON(w, http.StatusOK, map[string]string{"status": "success", "message": serviceName + " started"})
}

func HandleStopRigctld(w http.ResponseWriter, r *http.Request) {
	trxID, err := strconv.Atoi(r.PathValue("trx_id"))
	if err != nil || trxID <= 0 {
		common.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid TRX ID"})
		return
	}

	if !IsRigctldInstanceRunning(trxID) {
		common.WriteJSON(w, http.StatusConflict, map[string]string{"error": "Service not running"})
		return
	}

	serviceName := fmt.Sprintf("rigctld@%d.service", trxID)
	if err := exec.Command("sudo", "systemctl", "stop", serviceName).Run(); err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to stop service"})
		return
	}

	common.WriteJSON(w, http.StatusOK, map[string]string{"status": "success", "message": serviceName + " stopped"})
}

// func HandleListRigs(w http.ResponseWriter, r *http.Request) {
// jsonPath := "/etc/hamlib_rest_api/rigctld.json"
// file, err := os.ReadFile(jsonPath)
// if err != nil {
// 	common.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to read config"})
// 	return
// }

// var trxs []TransceiverConfig
// if err := json.Unmarshal(file, &trxs); err != nil {
// 	common.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to parse config"})
// 	return
// }

// 	for i := range trxs {
// 		if IsRigctldInstanceRunning(trxs[i].ID) {
// 			trxs[i].ServiceStatus = "running"
// 		} else {
// 			trxs[i].ServiceStatus = "stopped"
// 		}
// 	}

// 	if trxs == nil {
// 		trxs = []TransceiverConfig{}
// 	}

// 	WriteJSON(w, http.StatusOK, trxs)
// }

func HandleListRigs(w http.ResponseWriter, r *http.Request) {
	jsonPath := "/etc/hamlib_rest_api/rigctld.json"
	file, err := os.ReadFile(jsonPath)
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to read config"})
		return
	}

	var trxs []TransceiverConfig
	if err := json.Unmarshal(file, &trxs); err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to parse config"})
		return
	}

	for i := range trxs {
		if IsRigctldInstanceRunning(trxs[i].ID) {
			trxs[i].ServiceStatus = "running"
		} else {
			trxs[i].ServiceStatus = "stopped"
		}
	}

	if trxs == nil {
		trxs = []TransceiverConfig{}
	}

	// Antwort als Wrapper-Objekt inklusive Version
	response := RigListResponse{
		RigctldVersion: GetRigctldVersion(),
		Rigs:           trxs,
	}

	common.WriteJSON(w, http.StatusOK, response)
}
