package apihandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GET /trx/{trx_id}/ptt
func handleGetPTT(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "t")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"ptt": string(output[0])})
}

// POST /trx/{trx_id}/ptt
func handleSetPTT(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("T %s", body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/memory
func handleGetMemory(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "e")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"memory": string(output[0])})
}

// POST /trx/{trx_id}/memory
func handleSetMemory(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("E %s", body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/channel
func handleGetChannel(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "h")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"channel": string(output[0])})
}

// POST /trx/{trx_id}/channel
func handleSetChannel(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("H %s", body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/antenna
func handleGetAntenna(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "y 0")
	if err != nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	writeJSON(w, http.StatusOK, map[string][]string{"antennas": []string{string(output[0])}})
}

// POST /trx/{trx_id}/antenna
func handleSetAntenna(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("Y %s", body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/rit
func handleGetRit(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "j")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"rit": string(output[0])})
}

// POST /trx/{trx_id}/rit
func handleSetRit(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("J %s", body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/xit
func handleGetXit(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "z")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"xit": string(output[0])})
}

// POST /trx/{trx_id}/xit
func handleSetXit(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("Z %s", body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}

// POST /trx/{trx_id}/morse
func handleSetMorse(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["text"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	output, err := pollTrx(trxID, fmt.Sprintf("b %s", body["text"]))
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"response": string(output[0])})
}

// POST /trx/{trx_id}/morse/stop
func handleSetMorseStop(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "stop_morse")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"response": string(output[0])})
}

// GET /trx/{trx_id}/morse/wait
func handleGetMorseWait(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "wait_morse")
	if err != nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string][]string{"text": []string{string(output[0])}})
}

// GET /trx/{trx_id}/dtmf
func handleGetDtmf(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "recv_dtmf")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"dtmf": string(output[0])})
}

// POST /trx/{trx_id}/dtmf
func handleSetDtmf(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("send_dtmf %s", body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}

// POST /trx/{trx_id}/voice_mem
func handleSetVoiceMem(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("send_voice_mem %s", body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}
