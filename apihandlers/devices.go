package apihandlers

import (
	"encoding/json"
	"hamlib_rest_api/apihandlers/rigctld"
	"hamlib_rest_api/apihandlers/rotctld"
	"net/http"
	"os"
)

// GET /devices
func HandleListDevices(w http.ResponseWriter, r *http.Request) {
	// 1. Dateien lesen (Annahme: Pfade liegen in deinen Config-Files)
	rigData, err := os.ReadFile("/etc/hamlib_rest_api/rigctld.json")
	if err != nil {
		http.Error(w, "Could not read rig config", http.StatusInternalServerError)
		return
	}

	rotData, err := os.ReadFile("/etc/hamlib_rest_api/rotctld.json")
	if err != nil {
		http.Error(w, "Could not read rotator config", http.StatusInternalServerError)
		return
	}

	var rigs []rigctld.TransceiverConfig
	var rots []rotctld.RotatorConfig

	json.Unmarshal(rigData, &rigs)
	json.Unmarshal(rotData, &rots)

	response := DeviceList{
		Rigs:     rigs,
		Rotators: rots,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
