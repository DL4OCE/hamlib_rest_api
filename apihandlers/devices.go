package apihandlers

import (
	"encoding/json"
	"hamlib_rest_api/apihandlers/common"
	"hamlib_rest_api/apihandlers/rigctld"
	"hamlib_rest_api/apihandlers/rotctld"
	"net/http"
	"os"
)

type DevicesResponse struct {
	RigctldVersion string                      `json:"rigctld_version"`
	RotctldVersion string                      `json:"rotctld_version"`
	Rigs           []rigctld.TransceiverConfig `json:"rigs"`
	Rotators       []rotctld.RotatorConfig     `json:"rotators"`
}

// type DeviceList struct {
// 	Rigs     []rigctld.TransceiverConfig `json:"rigs"`
// 	Rotators []rotctld.RotatorConfig     `json:"rotators"`
// }

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

// 	for i := range rigs {
// 		if rigctld.IsRigctldInstanceRunning(rigs[i].ID) {
// 			rigs[i].ServiceStatus = "running"
// 		} else {
// 			rigs[i].ServiceStatus = "stopped"
// 		}
// 	}

// 	// Live-Status für Rotatoren ermitteln
// 	for i := range rots {
// 		if rotctld.IsRotctldInstanceRunning(rots[i].ID) {
// 			rots[i].ServiceStatus = "running"
// 		} else {
// 			rots[i].ServiceStatus = "stopped"
// 		}
// 	}

// 	if rigs == nil {
// 		rigs = []rigctld.TransceiverConfig{}
// 	}
// 	if rots == nil {
// 		rots = []rotctld.RotatorConfig{}
// 	}

// 	response := DeviceList{
// 		Rigs:     rigs,
// 		Rotators: rots,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(response)
// }

func HandleGetDevices(w http.ResponseWriter, r *http.Request) {
	// 1. Rigs laden & prüfen
	rigsFile, err1 := os.ReadFile("/etc/hamlib_rest_api/rigctld.json")
	var rigs []rigctld.TransceiverConfig
	if err1 == nil {
		_ = json.Unmarshal(rigsFile, &rigs)
	}
	for i := range rigs {
		if rigctld.IsRigctldInstanceRunning(rigs[i].ID) {
			rigs[i].ServiceStatus = "running"
		} else {
			rigs[i].ServiceStatus = "stopped"
		}
	}
	if rigs == nil {
		rigs = []rigctld.TransceiverConfig{}
	}

	// 2. Rotators laden & prüfen
	rotsFile, err2 := os.ReadFile("/etc/hamlib_rest_api/rotctld.json")
	var rotators []rotctld.RotatorConfig
	if err2 == nil {
		_ = json.Unmarshal(rotsFile, &rotators)
	}
	for i := range rotators {
		if rotctld.IsRotctldInstanceRunning(rotators[i].ID) {
			rotators[i].ServiceStatus = "running"
		} else {
			rotators[i].ServiceStatus = "stopped"
		}
	}
	if rotators == nil {
		rotators = []rotctld.RotatorConfig{}
	}

	// 3. Kombinierte Antwort mit beiden Versionen bauen
	response := DevicesResponse{
		RigctldVersion: rigctld.GetRigctldVersion(),
		RotctldVersion: rotctld.GetRotctldVersion(),
		Rigs:           rigs,
		Rotators:       rotators,
	}

	common.WriteJSON(w, http.StatusOK, response)
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
