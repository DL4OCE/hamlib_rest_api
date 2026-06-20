package rigctld

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GET /trx/{trx_id}/repeater/shift
func HandleGetRepeaterShift(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "r")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"shift_direction": string(output[0])})
}

// POST /trx/{trx_id}/repeater/shift
func HandleSetRepeaterShift(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("R %s", body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/repeater/offset
func HandleGetRepeaterOffset(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "o")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	writeJSON(w, http.StatusOK, map[string]string{"repeater_offset": string(output[0])})
}

// POST /trx/{trx_id}/repeater/offset
func HandleSetRepeaterOffset(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("O %s", body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/tone/ctcss
func HandleGetCtcssTone(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "c")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"ctcss_tone": string(output[0])})
}

// POST /trx/{trx_id}/tone/ctcss
func HandleSetCtcssTone(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("C %s", body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/tone/dcs
func HandleGetDcsTone(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "d")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"dcs_tone": string(output[0])})
}

// POST /trx/{trx_id}/tone/dcs
func HandleSetDcsTone(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("D %s", body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/sql/ctcss
func HandleGetCtcssSql(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_ctcss_sql")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"ctcss_sql": string(output[0])})
}

// POST /trx/{trx_id}/sql/ctcss
func HandleSetCtcssSql(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	output, err := pollTrx(trxID, fmt.Sprintf("set_ctcss_sql %s", body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
}

// GET /trx/{trx_id}/sql/dcs
func HandleGetDcsSql(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_dcs_sql")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"dcs_sql": string(output[0])})
}

// POST /trx/{trx_id}/sql/dcs
func HandleSetDcsSql(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	output, err := pollTrx(trxID, fmt.Sprintf("set_dcs_sql %s", body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
}

// GET /trx/{trx_id}/vfo
func HandleGetVFO(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "v")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"vfo_name": string(output)})
}

// POST /trx/{trx_id}/vfo
func HandleSetVFO(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("V %s", body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}
