package rotctld

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// HandleGetPosition maps to GET /rotator/{rotator_id}/position
// Retrieves the current Azimuth and Elevation from the rotator.
func HandleGetPosition(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}

	// Hamlib 'p' command returns Azimuth and Elevation on separate lines
	output, err := PollRotctlDaemon(rotID, "p")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	azimuth := "0.0"
	elevation := "0.0"
	if len(output) > 0 {
		azimuth = output[0]
	}
	if len(output) > 1 {
		elevation = output[1]
	}

	WriteJSON(w, http.StatusOK, map[string]string{
		"azimuth":   azimuth,
		"elevation": elevation,
	})
}

// HandleSetPosition maps to POST /rotator/{rotator_id}/position
// Sets the target Azimuth and Elevation for the rotator.
func HandleSetPosition(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}

	var body struct {
		Azimuth   string `json:"azimuth"`
		Elevation string `json:"elevation"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON request payload"})
		return
	}

	// Hamlib 'P' command syntax: P <azimuth> <elevation>
	command := "P " + body.Azimuth + " " + body.Elevation
	_, err = PollRotctlDaemon(rotID, command)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"status": "success", "message": "Target position updated"})
}

// HandlePark maps to POST /rotator/{rotator_id}/park
// Commands the rotator to go to its predefined park position.
func HandlePark(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}

	// Hamlib 'K' commands the rotator to park
	_, err = PollRotctlDaemon(rotID, "K")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"status": "success", "message": "Rotator moving to park position"})
}

// HandleStop maps to POST /rotator/{rotator_id}/stop
// Immediately aborts any ongoing movement of the rotator.
func HandleStop(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}

	// Hamlib 'S' immediately stops the rotator
	_, err = PollRotctlDaemon(rotID, "S")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"status": "success", "message": "Rotator stopped immediately"})
}

// HandleMove maps to POST /rotator/{rotator_id}/move
// Commands continuous movement in a given direction at a specific speed.
func HandleMove(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}

	var body struct {
		Direction string `json:"direction"`
		Speed     string `json:"speed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON request payload"})
		return
	}

	// Hamlib 'M' command syntax: M <direction> <speed>
	command := "M " + body.Direction + " " + body.Speed
	_, err = PollRotctlDaemon(rotID, command)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"status": "success", "message": "Rotator continuous movement initialized"})
}

// HandleReset maps to POST /rotator/{rotator_id}/reset
// Triggers a hardware or controller reset.
func HandleReset(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}

	// Hamlib 'R' resets the rotator controller
	_, err = PollRotctlDaemon(rotID, "R")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"status": "success", "message": "Rotator controller reset command issued"})
}
