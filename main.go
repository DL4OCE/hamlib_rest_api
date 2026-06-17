package main

import (
	"fmt"
	"net/http"
)

// Globaler Helfer, um JSON-Antworten einheitlich zu senden (entspricht deinem build_response)
// func writeJSON(w http.ResponseWriter, status int, data any) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	json.NewEncoder(w).Encode(data)
// }

// 1. ZENTRALE ROUTINE: Ersetzt dein "poll_trx"
// Sie nimmt die TRX-ID und das rohe Hamlib-Kommando, spricht mit dem Socket und gibt die Antwort zurück.
// func pollTrx(trxID int, command string) (string, error) {
// 	targetPort := 4532 + trxID
// 	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", targetPort))
// 	if err != nil {
// 		return "", fmt.Errorf("rigctld auf Port %d nicht erreichbar", targetPort)
// 	}
// 	defer conn.Close()

// 	// Befehl senden (mit Newline)
// 	fmt.Fprintf(conn, "%s\n", strings.TrimSpace(command))

// 	// Erste Zeile der Antwort lesen
// 	resp, err := bufio.NewReader(conn).ReadString('\n')
// 	if err != nil {
// 		return "", err
// 	}
// 	return strings.TrimSpace(resp), nil
// }

// === API ENDPUNKTE (Deine Wrapper-Routinen) ===

// // GET /trx/{trx_id}/frequency -> get_trx_frequency
// func handleGetFrequency(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id")) // ID aus der URL holen

// 	rawFreq, err := pollTrx(trxID, "f")
// 	if err != nil {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
// 		return
// 	}

// 	// build_response() Äquivalent als Map
// 	writeJSON(w, http.StatusOK, map[string]string{"freq": rawFreq})
// }

// // POST /trx/{trx_id}/frequency -> set_trx_frequency
// func handleSetFrequency(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

// 	// RequestBody dynamisch in eine Map dekodieren (json_decode)
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Ungültiger Body"})
// 		return
// 	}

// 	// Befehl an pollTrx übergeben (z.B. "F 145425000")
// 	_, err := pollTrx(trxID, "F "+body["newValue"])
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}

// 	// Leere Erfolgsantwort senden (array())
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/mode -> get_trx_mode

// // GET /trx/{trx_id}/mode -> get_trx_mode
// func handleGetMode(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

// 	output, err := pollTrx(trxID, "m")
// 	if err != nil || len(output) < 2 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}

// 	// FIX: Explicitly cast the outputs to string if they are treated as bytes/runes
// 	writeJSON(w, http.StatusOK, map[string]string{
// 		"mode":     string(output[0]),
// 		"passband": string(output[1]),
// 	})
// }

// // POST /trx/{trx_id}/mode -> set_trx_mode
// func handleSetMode(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

// 	// Decode incoming JSON dynamically (json_decode)
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["mode"] == "" || body["passband"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
// 		return
// 	}

// 	// Construct and send command (e.g., "M LSB 2400")
// 	fullCmd := fmt.Sprintf("M %s %s", body["mode"], body["passband"])
// 	_, err := pollTrx(trxID, fullCmd)
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}

// 	// Return empty response array()
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/split_frequency -> get_trx_split_frequency
// func handleGetSplitFrequency(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

// 	output, err := pollTrx(trxID, "i")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}

// 	// Returns the split frequency (qrg) from the first line
// 	writeJSON(w, http.StatusOK, map[string]string{
// 		"qrg": string(output[0]),
// 	})
// }

// // POST /trx/{trx_id}/split_frequency -> set_trx_split_frequency
// func handleSetSplitFrequency(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
// 		return
// 	}

// 	// Command to set split frequency (e.g., "I 145425000")
// 	fullCmd := fmt.Sprintf("I %s", body["newValue"])
// 	_, err := pollTrx(trxID, fullCmd)
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}

// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/split_mode -> get_trx_split_mode
// func handleGetSplitMode(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

// 	output, err := pollTrx(trxID, "x")
// 	if err != nil || len(output) < 2 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}

// 	// Line 1: Split Mode, Line 2: Split Passband
// 	writeJSON(w, http.StatusOK, map[string]string{
// 		"mode":     string(output[0]),
// 		"passband": string(output[1]),
// 	})
// }

// // POST /trx/{trx_id}/split_mode -> set_trx_split_mode
// func handleSetSplitMode(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["mode"] == "" || body["passband"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
// 		return
// 	}

// 	// Command to set split mode and passband (e.g., "X LSB 2700")
// 	fullCmd := fmt.Sprintf("X %s %s", body["mode"], body["passband"])
// 	_, err := pollTrx(trxID, fullCmd)
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}

// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/split_frequency_mode -> get_trx_split_frequency_mode
// func handleGetSplitFrequencyMode(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

// 	output, err := pollTrx(trxID, "k")
// 	if err != nil || len(output) < 3 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}

// 	// Line 1: Frequency, Line 2: Mode, Line 3: Passband
// 	writeJSON(w, http.StatusOK, map[string]string{
// 		"frequency": string(output[0]),
// 		"mode":      string(output[1]),
// 		"passband":  string(output[2]),
// 	})
// }

// // POST /trx/{trx_id}/split_frequency_mode -> set_trx_split_frequency_mode
// func handleSetSplitFrequencyMode(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["frequency"] == "" || body["mode"] == "" || body["passband"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
// 		return
// 	}

// 	// Command to set all three split parameters (e.g., "K 145425000 LSB 2700")
// 	fullCmd := fmt.Sprintf("K %s %s %s", body["frequency"], body["mode"], body["passband"])
// 	_, err := pollTrx(trxID, fullCmd)
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}

// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/split_vfo -> get_trx_split_vfo
// func handleGetSplitVFO(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

// 	output, err := pollTrx(trxID, "s")
// 	if err != nil || len(output) < 2 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}

// 	// Line 1: Split Mode (0 or 1), Line 2: TX VFO
// 	writeJSON(w, http.StatusOK, map[string]string{
// 		"split_mode": string(output[0]),
// 		"tx_vfo":     string(output[1]),
// 	})
// }

// // POST /trx/{trx_id}/split_vfo -> set_trx_split_vfo
// func handleSetSplitVFO(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["split_mode"] == "" || body["tx_vfo"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
// 		return
// 	}

// 	// Command: S <Split Mode> <TX VFO> (e.g., "S 1 VFOB")
// 	fullCmd := fmt.Sprintf("S %s %s", body["split_mode"], body["tx_vfo"])
// 	_, err := pollTrx(trxID, fullCmd)
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}

// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/tuningstep -> get_trx_tuningstep
// func handleGetTuningStep(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

// 	output, err := pollTrx(trxID, "n")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}

// 	writeJSON(w, http.StatusOK, map[string]string{
// 		"tuningstep": string(output[0]),
// 	})
// }

// // POST /trx/{trx_id}/tuningstep -> set_trx_tuningstep
// func handleSetTuningStep(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
// 		return
// 	}

// 	fullCmd := fmt.Sprintf("N %s", body["newValue"])
// 	_, err := pollTrx(trxID, fullCmd)
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}

// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/level/list -> get_trx_level_list
// func handleGetLevelList(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

// 	output, err := pollTrx(trxID, "l ?")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}

// 	// Explode by space equivalent: strings.Fields automatically splits by whitespace and removes empty tokens
// 	capabilities := strings.Fields(string(output[0]))

// 	writeJSON(w, http.StatusOK, map[string][]string{
// 		"capabilities": capabilities,
// 	})
// }

// // GET /trx/{trx_id}/level/{level_param} -> get_trx_level
// func handleGetLevel(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	levelParam := r.PathValue("level_param")

// 	if levelParam == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Missing level parameter"})
// 		return
// 	}

// 	fullCmd := fmt.Sprintf("l %s", levelParam)
// 	output, err := pollTrx(trxID, fullCmd)
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}

// 	writeJSON(w, http.StatusOK, map[string]string{
// 		"level": string(output[0]),
// 	})
// }

// // POST /trx/{trx_id}/level/{level_param} -> set_trx_level
// func handleSetLevel(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	levelParam := r.PathValue("level_param")

// 	if levelParam == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Missing level parameter"})
// 		return
// 	}

// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
// 		return
// 	}

// 	// Command: L <Level Name> <Value> (e.g., "L AF STRENGTH")
// 	fullCmd := fmt.Sprintf("L %s %s", levelParam, body["newValue"])
// 	_, err := pollTrx(trxID, fullCmd)
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}

// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// ==============================================================================
// FUNCTIONS, PARAMETERS, SCANS & TRANSCEIVE
// ==============================================================================

// // GET /trx/{trx_id}/function/list
// func handleGetFunctionList(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "u ?")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}
// 	// FIX: Explicitly cast output[0] to string
// 	writeJSON(w, http.StatusOK, map[string][]string{"capabilities": strings.Fields(string(output[0]))})
// }

// // GET /trx/{trx_id}/function/{param}
// func handleGetFunction(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	param := r.PathValue("param")
// 	output, err := pollTrx(trxID, fmt.Sprintf("u %s", param))
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}
// 	// FIX: Explicitly cast output[0] to string
// 	writeJSON(w, http.StatusOK, map[string]string{"function": string(output[0])})
// }

// // POST /trx/{trx_id}/function/{param}
// func handleSetFunction(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	param := r.PathValue("param")
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("U %s %s", param, body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/parameter/list
// func handleGetParameterList(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "p ?")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}
// 	// FIX: Explicitly cast output[0] to string
// 	writeJSON(w, http.StatusOK, map[string][]string{"capabilities": strings.Fields(string(output[0]))})
// }

// // GET /trx/{trx_id}/parameter/{param}
// func handleGetParameter(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	param := r.PathValue("param")
// 	output, err := pollTrx(trxID, fmt.Sprintf("p %s", param))
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}
// 	// FIX: Explicitly cast output[0] to string
// 	writeJSON(w, http.StatusOK, map[string]string{"parameter": string(output[0])})
// }

// // POST /trx/{trx_id}/parameter/{param}
// func handleSetParameter(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	param := r.PathValue("param")
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("P %s %s", param, body["newValue"])) // FIX: Korrigiert auf 'P'
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/scan/list
// func handleGetScanList(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "g ?")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}
// 	// FIX: Explicitly cast output[0] to string
// 	writeJSON(w, http.StatusOK, map[string][]string{"capabilities": strings.Fields(string(output[0]))})
// }

// // GET /trx/{trx_id}/scan/{param}
// func handleGetScan(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	param := r.PathValue("param")
// 	output, err := pollTrx(trxID, fmt.Sprintf("g %s", param))
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}
// 	// FIX: Explicitly cast output[0] to string
// 	writeJSON(w, http.StatusOK, map[string]string{"scan": string(output[0])})
// }

// // POST /trx/{trx_id}/scan/{param}
// func handleSetScan(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	param := r.PathValue("param")
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("G %s %s", param, body["newValue"])) // FIX: rigctld scan set ist 'G'
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/transceive/list
// func handleGetTransceiveList(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "A ?")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}
// 	// FIX: Explicitly cast output[0] to string
// 	writeJSON(w, http.StatusOK, map[string][]string{"capabilities": strings.Fields(string(output[0]))})
// }

// // GET /trx/{trx_id}/transceive
// func handleGetTransceive(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "a")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
// 		return
// 	}
// 	// FIX: Explicitly cast output[0] to string
// 	writeJSON(w, http.StatusOK, map[string]string{"transceive": string(output[0])})
// }

// // POST /trx/{trx_id}/transceive
// func handleSetTransceive(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("A %s", body["newValue"])) // FIX: Korrigiert auf 'A'
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// ==============================================================================
// REPEATER SHIFT, OFFSET, TONES & VFO
// ==============================================================================

// GET /trx/{trx_id}/repeater/shift
// func handleGetRepeaterShift(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "r")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	// FIX: Explicitly cast output[0] to string
// 	writeJSON(w, http.StatusOK, map[string]string{"shift_direction": string(output[0])})
// }

// // POST /trx/{trx_id}/repeater/shift
// func handleSetRepeaterShift(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("R %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/repeater/offset
// func handleGetRepeaterOffset(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "o")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	// FIX: Explicitly cast output[0] to string
// 	writeJSON(w, http.StatusOK, map[string]string{"repeater_offset": string(output[0])})
// }

// // POST /trx/{trx_id}/repeater/offset
// func handleSetRepeaterOffset(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("O %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/tone/ctcss
// func handleGetCtcssTone(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "c")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"ctcss_tone": string(output[0])})
// }

// // POST /trx/{trx_id}/tone/ctcss
// func handleSetCtcssTone(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("C %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/tone/dcs
// func handleGetDcsTone(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "d")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"dcs_tone": string(output[0])})
// }

// // POST /trx/{trx_id}/tone/dcs
// func handleSetDcsTone(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("D %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/vfo
// func handleGetVFO(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "v")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"vfo_name": string(output)})
// }

// // POST /trx/{trx_id}/vfo
// func handleSetVFO(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("V %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// ==============================================================================
// PTT, MEMORY, CHANNELS, RIT/XIT & ANTENNA
// ==============================================================================

// // GET /trx/{trx_id}/ptt
// func handleGetPTT(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "t")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"ptt": string(output[0])})
// }

// // POST /trx/{trx_id}/ptt
// func handleSetPTT(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("T %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/memory
// func handleGetMemory(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "e")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"memory": string(output[0])})
// }

// // POST /trx/{trx_id}/memory
// func handleSetMemory(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("E %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/channel
// func handleGetChannel(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "h")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"channel": string(output[0])})
// }

// // POST /trx/{trx_id}/channel
// func handleSetChannel(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("H %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/info
// func handleGetInfo(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "_")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"info": string(output[0])})
// }

// // GET /trx/{trx_id}/rit
// func handleGetRit(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "j")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"rit": string(output[0])})
// }

// // POST /trx/{trx_id}/rit
// func handleSetRit(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("J %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/xit
// func handleGetXit(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "z")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"xit": string(output[0])})
// }

// // POST /trx/{trx_id}/xit
// func handleSetXit(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("Z %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/antenna
// func handleGetAntenna(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "y 0")
// 	if err != nil {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string][]string{"antennas": []string{string(output[0])}})
// }

// // POST /trx/{trx_id}/antenna
// func handleSetAntenna(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("Y %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// ==============================================================================
// RAW RAW_COMMANDS, MORSE & POWER CONVERSIONS
// ==============================================================================

// // POST /trx/{trx_id}/raw
// func handleSetRawCommand(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["raw_command"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	output, err := pollTrx(trxID, fmt.Sprintf("w %s", body["raw_command"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output)}})
// }

// // POST /trx/{trx_id}/raw_rx
// func handleSetRawCommandRx(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["raw_command"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	// Combines command and optional expected bytes length
// 	fullCmd := fmt.Sprintf("W %s %s", body["raw_command"], body["number_of_expected_rx_bytes"])
// 	_, err := pollTrx(trxID, fullCmd)
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // POST /trx/{trx_id}/power/to_factor (PHP: get_trx_mw_power)
// func handleGetMwPower(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	fullCmd := fmt.Sprintf("4 %s %s %s", body["power_mW"], body["frequency"], body["mode"])
// 	output, err := pollTrx(trxID, fullCmd)
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"power_factor": string(output[0])})
// }

// // POST /trx/{trx_id}/power/to_mw (PHP: get_trx_power_mw)
// func handleGetPowerMw(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	fullCmd := fmt.Sprintf("2 %s %s %s", body["power_factor"], body["frequency"], body["mode"])
// 	output, err := pollTrx(trxID, fullCmd)
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"power_mw": string(output[0])})
// }

// // GET /trx/{trx_id}/capabilities
// func handleGetCapabilities(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "1")
// 	if err != nil {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
// }

// // GET /trx/{trx_id}/configuration
// func handleGetConfiguration(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "3")
// 	if err != nil {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
// }

// // POST /trx/{trx_id}/morse
// func handleSetMorse(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["text"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	output, err := pollTrx(trxID, fmt.Sprintf("b %s", body["text"]))
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"response": string(output[0])})
// }

// // POST /trx/{trx_id}/morse/stop
// func handleSetMorseStop(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "stop_morse")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"response": string(output[0])})
// }

// // GET /trx/{trx_id}/sql/ctcss
// func handleGetCtcssSql(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "get_ctcss_sql")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"ctcss_sql": string(output[0])})
// }

// // POST /trx/{trx_id}/sql/ctcss
// func handleSetCtcssSql(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	output, err := pollTrx(trxID, fmt.Sprintf("set_ctcss_sql %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
// }

// // GET /trx/{trx_id}/sql/dcs
// func handleGetDcsSql(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "get_dcs_sql")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"dcs_sql": string(output[0])})
// }

// // POST /trx/{trx_id}/sql/dcs
// func handleSetDcsSql(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	output, err := pollTrx(trxID, fmt.Sprintf("set_dcs_sql %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
// }

// // GET /trx/{trx_id}/dtmf
// func handleGetDtmf(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "recv_dtmf")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"dtmf": string(output[0])})
// }

// // POST /trx/{trx_id}/dtmf
// func handleSetDtmf(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("send_dtmf %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/morse/wait
// func handleGetMorseWait(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "wait_morse")
// 	if err != nil {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string][]string{"text": []string{string(output[0])}})
// }

// // GET /trx/{trx_id}/dcd
// func handleGetDcd(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "get_dcd")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"dcd": string(output[0])})
// }

// // GET /trx/{trx_id}/twiddle
// func handleGetTwiddle(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "get_twiddle")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"twiddle": string(output[0])})
// }

// // POST /trx/{trx_id}/twiddle
// func handleSetTwiddle(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("set_twiddle %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // GET /trx/{trx_id}/cache
// func handleGetCache(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "get_cache")
// 	if err != nil || len(output) < 1 {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]string{"cache": string(output[0])})
// }

// // POST /trx/{trx_id}/cache
// func handleSetCache(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	output, err := pollTrx(trxID, fmt.Sprintf("set_cache %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
// }

// // POST /trx/{trx_id}/state/dump
// func handleSetStateDump(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "dump_state")
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
// }

// // GET /trx/{trx_id}/rig_info
// func handleGetRigInfo(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "get_rig_info")
// 	if err != nil {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
// }

// // GET /trx/{trx_id}/modes
// func handleGetModes(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "get_modes")
// 	if err != nil {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string][]string{"response": []string{string(output[0])}})
// }

// // GET /trx/{trx_id}/power_state
// func handleGetPowerState(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	output, err := pollTrx(trxID, "get_powerstat")
// 	if err != nil {
// 		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string][]string{"power_state": []string{string(output[0])}})
// }

// // POST /trx/{trx_id}/power_state
// func handleSetPowerState(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("set_powerstate %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// // POST /trx/{trx_id}/voice_mem
// func handleSetVoiceMem(w http.ResponseWriter, r *http.Request) {
// 	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
// 	var body map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
// 		return
// 	}
// 	_, err := pollTrx(trxID, fmt.Sprintf("send_voice_mem %s", body["newValue"]))
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	writeJSON(w, http.StatusOK, map[string]any{})
// }

// === MAIN INTERFACES / ROUTER ===

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("GET /trx/{trx_id}/frequency", handleGetFrequency)
	// mux.HandleFunc("POST /trx/{trx_id}/frequency", handleSetFrequency)
	// mux.HandleFunc("GET /trx/{trx_id}/mode", handleGetMode)
	// mux.HandleFunc("POST /trx/{trx_id}/mode", handleSetMode)
	// mux.HandleFunc("GET /trx/{trx_id}/split_frequency_mode", handleGetSplitFrequencyMode)
	// mux.HandleFunc("POST /trx/{trx_id}/split_frequency_mode", handleSetSplitFrequencyMode)
	// mux.HandleFunc("GET /trx/{trx_id}/split_frequency", handleGetSplitFrequency)
	// mux.HandleFunc("POST /trx/{trx_id}/split_frequency", handleSetSplitFrequency)
	// mux.HandleFunc("GET /trx/{trx_id}/split_mode", handleGetSplitMode)
	// mux.HandleFunc("POST /trx/{trx_id}/split_mode", handleSetSplitMode)
	// mux.HandleFunc("GET /trx/{trx_id}/level/{level_param}", handleGetLevel)
	// mux.HandleFunc("POST /trx/{trx_id}/level/{level_param}", handleSetLevel)
	// mux.HandleFunc("GET /trx/{trx_id}/level/list", handleGetLevelList)
	// mux.HandleFunc("GET /trx/{trx_id}/tuningstep", handleGetTuningStep)
	// mux.HandleFunc("POST /trx/{trx_id}/tuningstep", handleSetTuningStep)
	// mux.HandleFunc("GET /trx/{trx_id}/split_vfo", handleGetSplitVFO)
	// mux.HandleFunc("POST /trx/{trx_id}/split_vfo", handleSetSplitVFO)

	// // Functions, Parameters, Scans & Transceive
	// mux.HandleFunc("GET /trx/{trx_id}/function/list", handleGetFunctionList)
	// mux.HandleFunc("GET /trx/{trx_id}/function/{param}", handleGetFunction)
	// mux.HandleFunc("POST /trx/{trx_id}/function/{param}", handleSetFunction)

	// mux.HandleFunc("GET /trx/{trx_id}/parameter/list", handleGetParameterList)
	// mux.HandleFunc("GET /trx/{trx_id}/parameter/{param}", handleGetParameter)
	// mux.HandleFunc("POST /trx/{trx_id}/parameter/{param}", handleSetParameter)

	// mux.HandleFunc("GET /trx/{trx_id}/scan/list", handleGetScanList)
	// mux.HandleFunc("GET /trx/{trx_id}/scan/{param}", handleGetScan)
	// mux.HandleFunc("POST /trx/{trx_id}/scan/{param}", handleSetScan)

	// mux.HandleFunc("GET /trx/{trx_id}/transceive/list", handleGetTransceiveList)
	// mux.HandleFunc("GET /trx/{trx_id}/transceive", handleGetTransceive)
	// mux.HandleFunc("POST /trx/{trx_id}/transceive", handleSetTransceive)

	// // Repeater Shift, Offset, Tones & VFO
	// mux.HandleFunc("GET /trx/{trx_id}/repeater/shift", handleGetRepeaterShift)
	// mux.HandleFunc("POST /trx/{trx_id}/repeater/shift", handleSetRepeaterShift)
	// mux.HandleFunc("GET /trx/{trx_id}/repeater/offset", handleGetRepeaterOffset)
	// mux.HandleFunc("POST /trx/{trx_id}/repeater/offset", handleSetRepeaterOffset)
	// mux.HandleFunc("GET /trx/{trx_id}/tone/ctcss", handleGetCtcssTone)
	// mux.HandleFunc("POST /trx/{trx_id}/tone/ctcss", handleSetCtcssTone)
	// mux.HandleFunc("GET /trx/{trx_id}/tone/dcs", handleGetDcsTone)
	// mux.HandleFunc("POST /trx/{trx_id}/tone/dcs", handleSetDcsTone)
	// mux.HandleFunc("GET /trx/{trx_id}/vfo", handleGetVFO)
	// mux.HandleFunc("POST /trx/{trx_id}/vfo", handleSetVFO)

	// // PTT, Memory, Channels, RIT/XIT & Antenna
	// mux.HandleFunc("GET /trx/{trx_id}/ptt", handleGetPTT)
	// mux.HandleFunc("POST /trx/{trx_id}/ptt", handleSetPTT)
	// mux.HandleFunc("GET /trx/{trx_id}/memory", handleGetMemory)
	// mux.HandleFunc("POST /trx/{trx_id}/memory", handleSetMemory)
	// mux.HandleFunc("GET /trx/{trx_id}/channel", handleGetChannel)
	// mux.HandleFunc("POST /trx/{trx_id}/channel", handleSetChannel)
	// mux.HandleFunc("GET /trx/{trx_id}/info", handleGetInfo)
	// mux.HandleFunc("GET /trx/{trx_id}/rit", handleGetRit)
	// mux.HandleFunc("POST /trx/{trx_id}/rit", handleSetRit)
	// mux.HandleFunc("GET /trx/{trx_id}/xit", handleGetXit)
	// mux.HandleFunc("POST /trx/{trx_id}/xit", handleSetXit)
	// mux.HandleFunc("GET /trx/{trx_id}/antenna", handleGetAntenna)
	// mux.HandleFunc("POST /trx/{trx_id}/antenna", handleSetAntenna)

	// // Raw Commands, Morse & Power Conversions
	// mux.HandleFunc("POST /trx/{trx_id}/raw", handleSetRawCommand)
	// mux.HandleFunc("POST /trx/{trx_id}/raw_rx", handleSetRawCommandRx)
	// mux.HandleFunc("POST /trx/{trx_id}/power/to_factor", handleGetMwPower)
	// mux.HandleFunc("POST /trx/{trx_id}/power/to_mw", handleGetPowerMw)
	// mux.HandleFunc("GET /trx/{trx_id}/capabilities", handleGetCapabilities)
	// mux.HandleFunc("GET /trx/{trx_id}/configuration", handleGetConfiguration)
	// mux.HandleFunc("POST /trx/{trx_id}/morse", handleSetMorse)
	// mux.HandleFunc("POST /trx/{trx_id}/morse/stop", handleSetMorseStop)

	// // SQL Extensions, Rig State & Misc
	// mux.HandleFunc("GET /trx/{trx_id}/sql/ctcss", handleGetCtcssSql)
	// mux.HandleFunc("POST /trx/{trx_id}/sql/ctcss", handleSetCtcssSql)
	// mux.HandleFunc("GET /trx/{trx_id}/sql/dcs", handleGetDcsSql)
	// mux.HandleFunc("POST /trx/{trx_id}/sql/dcs", handleSetDcsSql)
	// mux.HandleFunc("GET /trx/{trx_id}/dtmf", handleGetDtmf)
	// mux.HandleFunc("POST /trx/{trx_id}/dtmf", handleSetDtmf)
	// mux.HandleFunc("GET /trx/{trx_id}/morse/wait", handleGetMorseWait)
	// mux.HandleFunc("GET /trx/{trx_id}/dcd", handleGetDcd)
	// mux.HandleFunc("GET /trx/{trx_id}/twiddle", handleGetTwiddle)
	// mux.HandleFunc("POST /trx/{trx_id}/twiddle", handleSetTwiddle)
	// mux.HandleFunc("GET /trx/{trx_id}/cache", handleGetCache)
	// mux.HandleFunc("POST /trx/{trx_id}/cache", handleSetCache)
	// mux.HandleFunc("POST /trx/{trx_id}/state/dump", handleSetStateDump)
	// mux.HandleFunc("GET /trx/{trx_id}/rig_info", handleGetRigInfo)
	// mux.HandleFunc("GET /trx/{trx_id}/modes", handleGetModes)
	// mux.HandleFunc("GET /trx/{trx_id}/power_state", handleGetPowerState)
	// mux.HandleFunc("POST /trx/{trx_id}/power_state", handleSetPowerState)
	// mux.HandleFunc("POST /trx/{trx_id}/voice_mem", handleSetVoiceMem)

	fmt.Println("Hamlib Go-API läuft auf http://localhost:8080 ...")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Schwerwiegender Server-Fehler: %v\n", err)
	}
}
