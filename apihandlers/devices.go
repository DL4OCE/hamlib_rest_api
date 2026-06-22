package apihandlers

import (
	"encoding/json"
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

	// 2. Parsen
	var rigs []RigInfo
	var rots []RotatorInfo

	json.Unmarshal(rigData, &rigs)
	json.Unmarshal(rotData, &rots)

	// 3. Zusammenführen
	response := DeviceList{
		Rigs:     rigs,
		Rotators: rots,
	}

	// 4. Senden
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
