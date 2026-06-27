package rotctld

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// HandleGetInfo maps to GET /rotator/{rotator_id}/info
// Retrieves the hardware or firmware information string.
func HandleGetInfo(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}

	// Hamlib '_' command requests general backend info
	output, err := PollRotctlDaemon(rotID, "_")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"info": output,
	})
}

// HandleGetStatus maps to GET /rotator/{rotator_id}/status
// Retrieves the current mechanical status or operational state flag.
func HandleGetStatus(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}

	// Hamlib 's' command fetches status flags
	output, err := PollRotctlDaemon(rotID, "s")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"status": output,
	})
}

// HandleGetState maps to GET /rotator/{rotator_id}/state
// Requests a complete raw dump of the internal state values.
func HandleGetState(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}

	// Hamlib 'dump_state' prints out all config/state blocks
	output, err := PollRotctlDaemon(rotID, "dump_state")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"state": output,
	})
}

// HandleGetCapabilities maps to GET /rotator/{rotator_id}/capabilities
// Retrieves the rotator capabilities mask.
func HandleGetCapabilities(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}

	// Hamlib '1' lists out capability parameters
	output, err := PollRotctlDaemon(rotID, "1")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"capabilities": output,
	})
}

// HandleGetLevel maps to GET /rotator/{rotator_id}/level/{level}
func HandleGetLevel(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}
	levelParam := r.PathValue("level")

	// Hamlib 'v' syntax: v <level_parameter>
	output, err := PollRotctlDaemon(rotID, "v "+levelParam)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"level":    levelParam,
		"response": output,
	})
}

// HandleSetLevel maps to POST /rotator/{rotator_id}/level/{level}
func HandleSetLevel(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}
	levelParam := r.PathValue("level")

	var body struct {
		NewValue string `json:"newValue"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON payload"})
		return
	}

	// Hamlib 'V' syntax: V <level_parameter> <value>
	_, err = PollRotctlDaemon(rotID, "V "+levelParam+" "+body.NewValue)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

// HandleGetFunction maps to GET /rotator/{rotator_id}/function/{function}
func HandleGetFunction(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}
	funcParam := r.PathValue("function")

	output, err := PollRotctlDaemon(rotID, "u "+funcParam)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"function": funcParam,
		"response": output,
	})
}

// HandleSetFunction maps to POST /rotator/{rotator_id}/function/{function}
func HandleSetFunction(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}
	funcParam := r.PathValue("function")

	var body struct {
		NewValue string `json:"newValue"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON payload"})
		return
	}

	_, err = PollRotctlDaemon(rotID, "U "+funcParam+" "+body.NewValue)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

// HandleGetParameter maps to GET /rotator/{rotator_id}/parameter/{parameter}
func HandleGetParameter(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}
	paramName := r.PathValue("parameter")

	output, err := PollRotctlDaemon(rotID, "u "+paramName)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"parameter": paramName,
		"response":  output,
	})
}

// HandleSetParameter maps to POST /rotator/{rotator_id}/parameter/{parameter}
func HandleSetParameter(w http.ResponseWriter, r *http.Request) {
	rotID, err := strconv.Atoi(r.PathValue("rotator_id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid rotator ID format"})
		return
	}
	paramName := r.PathValue("parameter")

	var body struct {
		NewValue string `json:"newValue"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON payload"})
		return
	}

	_, err = PollRotctlDaemon(rotID, "U "+paramName+" "+body.NewValue)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"status": "success"})
}
