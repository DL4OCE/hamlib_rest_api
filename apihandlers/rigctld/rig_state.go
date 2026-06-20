package rigctld

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GET /trx/{trx_id}/info
func HandleGetInfo(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "_")
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"info": string(output[0])})
}

// GET /trx/{trx_id}/rig_info
func HandleGetRigInfo(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_rig_info")
	if err != nil {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
}

// GET /trx/{trx_id}/modes
func HandleGetModes(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_modes")
	if err != nil {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
}

// GET /trx/{trx_id}/dcd
func HandleGetDcd(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_dcd")
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"dcd": string(output[0])})
}

// GET /trx/{trx_id}/twiddle
func HandleGetTwiddle(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_twiddle")
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"twiddle": string(output[0])})
}

// POST /trx/{trx_id}/twiddle
func HandleSetTwiddle(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("set_twiddle %s", body["newValue"]))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/cache
func HandleGetCache(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_cache")
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"cache": string(output[0])})
}

// POST /trx/{trx_id}/cache
func HandleSetCache(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	output, err := pollTrx(trxID, fmt.Sprintf("set_cache %s", body["newValue"]))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
}

// GET /trx/{trx_id}/power_state
func HandleGetPowerState(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "get_powerstat")
	if err != nil {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"power_state": []string{string(output[0])}})
}

// POST /trx/{trx_id}/power_state
func HandleSetPowerState(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("set_powerstate %s", body["newValue"]))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

// POST /trx/{trx_id}/state/dump
func HandleSetStateDump(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "dump_state")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
}

// POST /trx/{trx_id}/raw
func HandleSetRawCommand(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["raw_command"] == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	output, err := pollTrx(trxID, fmt.Sprintf("w %s", body["raw_command"]))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output)}})
}

// POST /trx/{trx_id}/raw_rx
func HandleSetRawCommandRx(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["raw_command"] == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	// Combines command and optional expected bytes length
	fullCmd := fmt.Sprintf("W %s %s", body["raw_command"], body["number_of_expected_rx_bytes"])
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

// POST /trx/{trx_id}/power/to_factor (PHP: get_trx_mw_power)
func HandleGetMwPower(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	fullCmd := fmt.Sprintf("4 %s %s %s", body["power_mW"], body["frequency"], body["mode"])
	output, err := pollTrx(trxID, fullCmd)
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"power_factor": string(output[0])})
}

// POST /trx/{trx_id}/power/to_mw (PHP: get_trx_power_mw)
func HandleGetPowerMw(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	fullCmd := fmt.Sprintf("2 %s %s %s", body["power_factor"], body["frequency"], body["mode"])
	output, err := pollTrx(trxID, fullCmd)
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"power_mw": string(output[0])})
}
