package apihandlers

import (
	"encoding/json"
	"hamlib_rest_api/apihandlers/rigctld"
	"hamlib_rest_api/apihandlers/rotctld"
	"net/http"
	"os"
)

// type DeviceList struct {
// 	Rigs     []rigctld.TransceiverConfig `json:"rigs"`
// 	Rotators []rotctld.RotatorConfig     `json:"rotators"`
// }

func HandleListDevices(w http.ResponseWriter, r *http.Request) {
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

	// Live-Status für Rigs ermitteln
	for i := range rigs {
		if rigctld.IsRigctldInstanceRunning(rigs[i].ID) { // Nutzt deine Prüffunktion (ggf. Export-Paketpräfix anpassen)
			rigs[i].ServiceStatus = "RUNNING"
		} else {
			rigs[i].ServiceStatus = "STOPPED"
		}
	}

	// Live-Status für Rotatoren ermitteln
	for i := range rots {
		if rotctld.IsRotctldInstanceRunning(rots[i].ID) { // Nutzt deine Rotator-Prüffunktion
			rots[i].ServiceStatus = "RUNNING" // bzw. das Feld, das in RotatorConfig den Status hält
		} else {
			rots[i].ServiceStatus = "STOPPED"
		}
	}

	if rigs == nil {
		rigs = []rigctld.TransceiverConfig{}
	}
	if rots == nil {
		rots = []rotctld.RotatorConfig{}
	}

	response := DeviceList{
		Rigs:     rigs,
		Rotators: rots,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GET /devices
// func HandleListDevices(w http.ResponseWriter, r *http.Request) {
// 	rigData, err := os.ReadFile("/etc/hamlib_rest_api/rigctld.json")
// 	if err != nil {
// 		http.Error(w, "Could not read rig config", http.StatusInternalServerError)
// 		return
// 	}

// 	rotData, err := os.ReadFile("/etc/hamlib_rest_api/rotctld.json")
// 	if err != nil {
// 		http.Error(w, "Could not read rotator config", http.StatusInternalServerError)
// 		return
// 	}

// 	var rigs []rigctld.TransceiverConfig
// 	var rots []rotctld.RotatorConfig

// 	json.Unmarshal(rigData, &rigs)
// 	json.Unmarshal(rotData, &rots)

// 	response := DeviceList{
// 		Rigs:     rigs,
// 		Rotators: rots,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)

// }
