package rigctld

import (
	"encoding/json"
	"fmt"
	"hamlib_rest_api/apihandlers/common"
	"net/http"
	"strconv"
)

// --- PTT, Memory, Channel, Antenna, RIT, XIT ---

func HandleGetPTT(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "t")
	if err != nil || len(output) == 0 {
		common.WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]string{"ptt": output[0]})
}

func HandleSetPTT(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		common.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("T %s", body.NewValue))
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]any{})
}

func HandleGetMemoryChannelNumber(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "e")
	if err != nil || len(output) == 0 {
		common.WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]string{"memory": output[0]})
}

func HandleSetMemoryChannelNumber(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		common.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("E %s", body.NewValue))
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]any{})
}

func HandleGetChannel(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "h")
	if err != nil || len(output) == 0 {
		common.WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]string{"channel": output[0]})
}

func HandleSetChannel(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		common.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("H %s", body.NewValue))
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]any{})
}

func HandleGetAntennaNumber(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "y 0")
	if err != nil || len(output) == 0 {
		common.WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string][]string{"antennas": output})
}

func HandleSetAntennaNumber(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		common.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("Y %s", body.NewValue))
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]any{})
}

func HandleGetRit(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "j")
	if err != nil || len(output) == 0 {
		common.WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]string{"rit": output[0]})
}

func HandleSetRit(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		common.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("J %s", body.NewValue))
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]any{})
}

func HandleGetXit(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "z")
	if err != nil || len(output) == 0 {
		common.WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]string{"xit": output[0]})
}

func HandleSetXit(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		common.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("Z %s", body.NewValue))
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]any{})
}

// --- Morse, DTMF, Voice ---

func HandleSetMorse(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body struct {
		Text string `json:"text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Text == "" {
		common.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	output, err := pollTrx(trxID, fmt.Sprintf("b %s", body.Text))
	if err != nil || len(output) == 0 {
		common.WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]string{"response": output[0]})
}

func HandleSetMorseStop(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "stop_morse")
	if err != nil || len(output) == 0 {
		common.WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]string{"response": output[0]})
}

func HandleGetMorseWait(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "wait_morse")
	if err != nil || len(output) == 0 {
		common.WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string][]string{"text": output})
}

func HandleGetDtmf(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "recv_dtmf")
	if err != nil || len(output) == 0 {
		common.WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]string{"dtmf": output[0]})
}

func HandleSetDtmf(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		common.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("send_dtmf %s", body.NewValue))
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]any{})
}

func HandleSetVoiceMem(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		common.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("send_voice_mem %s", body.NewValue))
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	common.WriteJSON(w, http.StatusOK, map[string]any{})
}
