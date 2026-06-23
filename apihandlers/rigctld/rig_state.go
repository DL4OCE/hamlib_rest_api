package rigctld

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// --- System, Status & Konfiguration ---

func HandleGetInfo(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "_")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"info": output[0]})
}

func HandleGetRigInfo(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_rig_info")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Failed to get rig info"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"response": output})
}

func HandleGetModes(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_modes")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Failed to get modes"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{
		"modes": output,
	})
}

func HandleGetDcd(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_dcd")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"dcd": output[0]})
}

// --- Twiddle & Cache ---

func HandleGetTwiddle(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_twiddle")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"twiddle": output[0]})
}

func HandleSetTwiddle(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("set_twiddle %s", body.NewValue))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

func HandleGetCache(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_cache")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"cache": output[0]})
}

func HandleSetCache(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	output, err := pollTrx(trxID, fmt.Sprintf("set_cache %s", body.NewValue))
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error setting cache"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"response": output})
}

// --- Power & State ---

func HandleGetPowerState(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_powerstat")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"power_state": output})
}

func HandleSetPowerState(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body ValuePayload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("set_powerstate %s", body.NewValue))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

func HandleSetStateDump(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "dump_state")
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to dump state"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"response": output})
}

// --- Raw Commands ---

func HandleSetRawCommand(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body struct {
		RawCommand string `json:"raw_command"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.RawCommand == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	output, err := pollTrx(trxID, fmt.Sprintf("w %s", body.RawCommand))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"response": output})
}

func HandleSetRawCommandRx(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body struct {
		RawCommand         string `json:"raw_command"`
		ExpectedBytesCount string `json:"number_of_expected_rx_bytes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.RawCommand == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("W %s %s", body.RawCommand, body.ExpectedBytesCount))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

// --- Power Conversion ---

func HandleGetMwPower(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body struct {
		PowerMW   string `json:"power_mW"`
		Frequency string `json:"frequency"`
		Mode      string `json:"mode"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	output, err := pollTrx(trxID, fmt.Sprintf("4 %s %s %s", body.PowerMW, body.Frequency, body.Mode))
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"power_factor": output[0]})
}

func HandleGetPowerMw(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body struct {
		PowerFactor string `json:"power_factor"`
		Frequency   string `json:"frequency"`
		Mode        string `json:"mode"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	output, err := pollTrx(trxID, fmt.Sprintf("2 %s %s %s", body.PowerFactor, body.Frequency, body.Mode))
	if err != nil || len(output) == 0 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"power_mw": output[0]})
}
