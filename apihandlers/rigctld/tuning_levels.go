package rigctld

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// GET /trx/{trx_id}/tuningstep -> get_trx_tuningstep
func HandleGetTuningStep(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "n")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"tuningstep": string(output[0]),
	})
}

// POST /trx/{trx_id}/tuningstep -> set_trx_tuningstep
func HandleSetTuningStep(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	fullCmd := fmt.Sprintf("N %s", body["newValue"])
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/level/list -> get_trx_level_list
func HandleGetLevelList(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "l ?")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	// Explode by space equivalent: strings.Fields automatically splits by whitespace and removes empty tokens
	capabilities := strings.Fields(string(output[0]))

	writeJSON(w, http.StatusOK, map[string][]string{
		"capabilities": capabilities,
	})
}

// GET /trx/{trx_id}/level/{level_param} -> get_trx_level
func HandleGetLevel(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	levelParam := r.PathValue("level_param")

	if levelParam == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Missing level parameter"})
		return
	}

	fullCmd := fmt.Sprintf("l %s", levelParam)
	output, err := pollTrx(trxID, fullCmd)
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"level": string(output[0]),
	})
}

// POST /trx/{trx_id}/level/{level_param} -> set_trx_level
func HandleSetLevel(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	levelParam := r.PathValue("level_param")

	if levelParam == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Missing level parameter"})
		return
	}

	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	// Command: L <Level Name> <Value> (e.g., "L AF STRENGTH")
	fullCmd := fmt.Sprintf("L %s %s", levelParam, body["newValue"])
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{})
}
