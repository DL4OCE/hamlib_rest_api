package rigctld

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// GET /trx/{trx_id}/function/list
func HandleGetFunctionList(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "u ?")
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	WriteJSON(w, http.StatusOK, map[string][]string{"capabilities": strings.Fields(string(output[0]))})
}

// GET /trx/{trx_id}/function/{param}
func HandleGetFunction(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	param := r.PathValue("param")
	output, err := pollTrx(trxID, fmt.Sprintf("u %s", param))
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	WriteJSON(w, http.StatusOK, map[string]string{"function": string(output[0])})
}

// POST /trx/{trx_id}/function/{param}
func HandleSetFunction(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	param := r.PathValue("param")
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("U %s %s", param, body["newValue"]))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/parameter/list
func HandleGetParameterList(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "p ?")
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	WriteJSON(w, http.StatusOK, map[string][]string{"capabilities": strings.Fields(string(output[0]))})
}

// GET /trx/{trx_id}/parameter/{param}
func HandleGetParameter(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	param := r.PathValue("param")
	output, err := pollTrx(trxID, fmt.Sprintf("p %s", param))
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	WriteJSON(w, http.StatusOK, map[string]string{"parameter": string(output[0])})
}

// POST /trx/{trx_id}/parameter/{param}
func HandleSetParameter(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	param := r.PathValue("param")
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("P %s %s", param, body["newValue"])) // FIX: Korrigiert auf 'P'
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/scan/list
func HandleGetScanList(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "g ?")
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	WriteJSON(w, http.StatusOK, map[string][]string{"capabilities": strings.Fields(string(output[0]))})
}

// GET /trx/{trx_id}/scan/{param}
func HandleGetScan(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	param := r.PathValue("param")
	output, err := pollTrx(trxID, fmt.Sprintf("g %s", param))
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	WriteJSON(w, http.StatusOK, map[string]string{"scan": string(output[0])})
}

// POST /trx/{trx_id}/scan/{param}
func HandleSetScan(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	param := r.PathValue("param")
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("G %s %s", param, body["newValue"])) // FIX: rigctld scan set ist 'G'
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/transceive/list
func HandleGetTransceiveList(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "A ?")
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	WriteJSON(w, http.StatusOK, map[string][]string{"capabilities": strings.Fields(string(output[0]))})
}

// GET /trx/{trx_id}/transceive
func HandleGetTransceive(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "a")
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	WriteJSON(w, http.StatusOK, map[string]string{"transceive": string(output[0])})
}

// POST /trx/{trx_id}/transceive
func HandleSetTransceive(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("A %s", body["newValue"])) // FIX: Korrigiert auf 'A'
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/capabilities
func HandleGetCapabilities(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "1")
	if err != nil {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
}

// GET /trx/{trx_id}/configuration
func HandleGetConfiguration(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "3")
	if err != nil {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
}
