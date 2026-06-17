package apihandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// GET /trx/{trx_id}/function/list
func handleGetFunctionList(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "u ?")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	writeJSON(w, http.StatusOK, map[string][]string{"capabilities": strings.Fields(string(output[0]))})
}

// GET /trx/{trx_id}/function/{param}
func handleGetFunction(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	param := r.PathValue("param")
	output, err := pollTrx(trxID, fmt.Sprintf("u %s", param))
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	writeJSON(w, http.StatusOK, map[string]string{"function": string(output[0])})
}

// POST /trx/{trx_id}/function/{param}
func handleSetFunction(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	param := r.PathValue("param")
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("U %s %s", param, body["newValue"]))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/parameter/list
func handleGetParameterList(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "p ?")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	writeJSON(w, http.StatusOK, map[string][]string{"capabilities": strings.Fields(string(output[0]))})
}

// GET /trx/{trx_id}/parameter/{param}
func handleGetParameter(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	param := r.PathValue("param")
	output, err := pollTrx(trxID, fmt.Sprintf("p %s", param))
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	writeJSON(w, http.StatusOK, map[string]string{"parameter": string(output[0])})
}

// POST /trx/{trx_id}/parameter/{param}
func handleSetParameter(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	param := r.PathValue("param")
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("P %s %s", param, body["newValue"])) // FIX: Korrigiert auf 'P'
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/scan/list
func handleGetScanList(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "g ?")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	writeJSON(w, http.StatusOK, map[string][]string{"capabilities": strings.Fields(string(output[0]))})
}

// GET /trx/{trx_id}/scan/{param}
func handleGetScan(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	param := r.PathValue("param")
	output, err := pollTrx(trxID, fmt.Sprintf("g %s", param))
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	writeJSON(w, http.StatusOK, map[string]string{"scan": string(output[0])})
}

// POST /trx/{trx_id}/scan/{param}
func handleSetScan(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	param := r.PathValue("param")
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("G %s %s", param, body["newValue"])) // FIX: rigctld scan set ist 'G'
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/transceive/list
func handleGetTransceiveList(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "A ?")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	writeJSON(w, http.StatusOK, map[string][]string{"capabilities": strings.Fields(string(output[0]))})
}

// GET /trx/{trx_id}/transceive
func handleGetTransceive(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "a")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}
	// FIX: Explicitly cast output[0] to string
	writeJSON(w, http.StatusOK, map[string]string{"transceive": string(output[0])})
}

// POST /trx/{trx_id}/transceive
func handleSetTransceive(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, fmt.Sprintf("A %s", body["newValue"])) // FIX: Korrigiert auf 'A'
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/capabilities
func handleGetCapabilities(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "1")
	if err != nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
}

// GET /trx/{trx_id}/configuration
func handleGetConfiguration(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	output, err := pollTrx(trxID, "3")
	if err != nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
}
