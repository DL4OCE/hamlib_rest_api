package rigctld

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// --- Tuning & Level ---

func HandleGetTuningStep(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "n")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"tuningstep": output[0]})
}

func HandleSetTuningStep(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("N %s", body.NewValue))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

func HandleGetLevelList(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "l ?")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"levels": strings.Fields(output[0])})
}

func HandleGetLevel(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	levelParam := r.PathValue("level_param")
	if levelParam == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Missing parameter"})
		return
	}
	output, err := pollTrx(trxID, fmt.Sprintf("l %s", levelParam))
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"level": output[0]})
}

func HandleSetLevel(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	levelParam := r.PathValue("level_param")
	if levelParam == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Missing parameter"})
		return
	}
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("L %s %s", levelParam, body.NewValue))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}
