package rigctld

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// --- Repeater, Tone, SQL, VFO ---

func HandleGetRepeaterShift(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "r")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"shift_direction": output[0]})
}

func HandleSetRepeaterShift(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("R %s", body.NewValue))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

func HandleGetRepeaterOffset(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "o")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"repeater_offset": output[0]})
}

func HandleSetRepeaterOffset(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("O %s", body.NewValue))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

func HandleGetCtcssTone(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "c")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"ctcss_tone": output[0]})
}

func HandleSetCtcssTone(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("C %s", body.NewValue))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

func HandleGetDcsTone(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "d")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"dcs_tone": output[0]})
}

func HandleSetDcsTone(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("D %s", body.NewValue))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

func HandleGetCtcssSql(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_ctcss_sql")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"ctcss_sql": output[0]})
}

func HandleSetCtcssSql(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	output, err := pollTrx(trxID, fmt.Sprintf("set_ctcss_sql %s", body.NewValue))
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error setting SQL"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"response": output})
}

func HandleGetDcsSql(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_dcs_sql")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"dcs_sql": output[0]})
}

func HandleSetDcsSql(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	output, err := pollTrx(trxID, fmt.Sprintf("set_dcs_sql %s", body.NewValue))
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error setting SQL"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"response": output})
}

func HandleGetVFO(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "v")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"vfo_name": output[0]})
}

func HandleSetVFO(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("V %s", body.NewValue))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}
