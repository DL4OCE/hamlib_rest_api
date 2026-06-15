package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
)

// Globaler Helfer, um JSON-Antworten einheitlich zu senden (entspricht deinem build_response)
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// 1. ZENTRALE ROUTINE: Ersetzt dein "poll_trx"
// Sie nimmt die TRX-ID und das rohe Hamlib-Kommando, spricht mit dem Socket und gibt die Antwort zurück.
func pollTrx(trxID int, command string) (string, error) {
	targetPort := 4532 + trxID
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", targetPort))
	if err != nil {
		return "", fmt.Errorf("rigctld auf Port %d nicht erreichbar", targetPort)
	}
	defer conn.Close()

	// Befehl senden (mit Newline)
	fmt.Fprintf(conn, "%s\n", strings.TrimSpace(command))

	// Erste Zeile der Antwort lesen
	resp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(resp), nil
}

// === API ENDPUNKTE (Deine Wrapper-Routinen) ===

// GET /trx/{trx_id}/frequency -> get_trx_frequency
func handleGetFrequency(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id")) // ID aus der URL holen

	rawFreq, err := pollTrx(trxID, "f")
	if err != nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}

	// build_response() Äquivalent als Map
	writeJSON(w, http.StatusOK, map[string]string{"freq": rawFreq})
}

// POST /trx/{trx_id}/frequency -> set_trx_frequency
func handleSetFrequency(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	// RequestBody dynamisch in eine Map dekodieren (json_decode)
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Ungültiger Body"})
		return
	}

	// Befehl an pollTrx übergeben (z.B. "F 145425000")
	_, err := pollTrx(trxID, "F "+body["newValue"])
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	// Leere Erfolgsantwort senden (array())
	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/mode -> get_trx_mode

// GET /trx/{trx_id}/mode -> get_trx_mode
func handleGetMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "m")
	if err != nil || len(output) < 2 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	// FIX: Explicitly cast the outputs to string if they are treated as bytes/runes
	writeJSON(w, http.StatusOK, map[string]string{
		"mode":     string(output[0]),
		"passband": string(output[1]),
	})
}

// POST /trx/{trx_id}/mode -> set_trx_mode
func handleSetMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	// Decode incoming JSON dynamically (json_decode)
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["mode"] == "" || body["passband"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	// Construct and send command (e.g., "M LSB 2400")
	fullCmd := fmt.Sprintf("M %s %s", body["mode"], body["passband"])
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	// Return empty response array()
	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/split_frequency -> get_trx_split_frequency
func handleGetSplitFrequency(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "i")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	// Returns the split frequency (qrg) from the first line
	writeJSON(w, http.StatusOK, map[string]string{
		"qrg": string(output[0]),
	})
}

// POST /trx/{trx_id}/split_frequency -> set_trx_split_frequency
func handleSetSplitFrequency(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	// Command to set split frequency (e.g., "I 145425000")
	fullCmd := fmt.Sprintf("I %s", body["newValue"])
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/split_mode -> get_trx_split_mode
func handleGetSplitMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "x")
	if err != nil || len(output) < 2 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	// Line 1: Split Mode, Line 2: Split Passband
	writeJSON(w, http.StatusOK, map[string]string{
		"mode":     string(output[0]),
		"passband": string(output[1]),
	})
}

// POST /trx/{trx_id}/split_mode -> set_trx_split_mode
func handleSetSplitMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["mode"] == "" || body["passband"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	// Command to set split mode and passband (e.g., "X LSB 2700")
	fullCmd := fmt.Sprintf("X %s %s", body["mode"], body["passband"])
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/split_frequency_mode -> get_trx_split_frequency_mode
func handleGetSplitFrequencyMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "k")
	if err != nil || len(output) < 3 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	// Line 1: Frequency, Line 2: Mode, Line 3: Passband
	writeJSON(w, http.StatusOK, map[string]string{
		"frequency": string(output[0]),
		"mode":      string(output[1]),
		"passband":  string(output[2]),
	})
}

// POST /trx/{trx_id}/split_frequency_mode -> set_trx_split_frequency_mode
func handleSetSplitFrequencyMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["frequency"] == "" || body["mode"] == "" || body["passband"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	// Command to set all three split parameters (e.g., "K 145425000 LSB 2700")
	fullCmd := fmt.Sprintf("K %s %s %s", body["frequency"], body["mode"], body["passband"])
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/split_vfo -> get_trx_split_vfo
func handleGetSplitVFO(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "s")
	if err != nil || len(output) < 2 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	// Line 1: Split Mode (0 or 1), Line 2: TX VFO
	writeJSON(w, http.StatusOK, map[string]string{
		"split_mode": string(output[0]),
		"tx_vfo":     string(output[1]),
	})
}

// POST /trx/{trx_id}/split_vfo -> set_trx_split_vfo
func handleSetSplitVFO(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["split_mode"] == "" || body["tx_vfo"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	// Command: S <Split Mode> <TX VFO> (e.g., "S 1 VFOB")
	fullCmd := fmt.Sprintf("S %s %s", body["split_mode"], body["tx_vfo"])
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/tuningstep -> get_trx_tuningstep
func handleGetTuningStep(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "n")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"tuningstep": string(output[0]),
	})
}

// POST /trx/{trx_id}/tuningstep -> set_trx_tuningstep
func handleSetTuningStep(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	fullCmd := fmt.Sprintf("N %s", body["newValue"])
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/level/list -> get_trx_level_list
func handleGetLevelList(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "l ?")
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	// Explode by space equivalent: strings.Fields automatically splits by whitespace and removes empty tokens
	capabilities := strings.Fields(string(output[0]))

	writeJSON(w, http.StatusOK, map[string][]string{
		"capabilities": capabilities,
	})
}

// GET /trx/{trx_id}/level/{level_param} -> get_trx_level
func handleGetLevel(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	levelParam := r.PathValue("level_param")

	if levelParam == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Missing level parameter"})
		return
	}

	fullCmd := fmt.Sprintf("l %s", levelParam)
	output, err := pollTrx(trxID, fullCmd)
	if err != nil || len(output) < 1 {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"level": string(output[0]),
	})
}

// POST /trx/{trx_id}/level/{level_param} -> set_trx_level
func handleSetLevel(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	levelParam := r.PathValue("level_param")

	if levelParam == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Missing level parameter"})
		return
	}

	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	// Command: L <Level Name> <Value> (e.g., "L AF STRENGTH")
	fullCmd := fmt.Sprintf("L %s %s", levelParam, body["newValue"])
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{})
}

// === MAIN INTERFACES / ROUTER ===

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /trx/{trx_id}/frequency", handleGetFrequency)
	mux.HandleFunc("POST /trx/{trx_id}/frequency", handleSetFrequency)
	mux.HandleFunc("GET /trx/{trx_id}/mode", handleGetMode)
	mux.HandleFunc("POST /trx/{trx_id}/mode", handleSetMode)
	mux.HandleFunc("GET /trx/{trx_id}/split_frequency_mode", handleGetSplitFrequencyMode)
	mux.HandleFunc("POST /trx/{trx_id}/split_frequency_mode", handleSetSplitFrequencyMode)
	mux.HandleFunc("GET /trx/{trx_id}/split_frequency", handleGetSplitFrequency)
	mux.HandleFunc("POST /trx/{trx_id}/split_frequency", handleSetSplitFrequency)
	mux.HandleFunc("GET /trx/{trx_id}/split_mode", handleGetSplitMode)
	mux.HandleFunc("POST /trx/{trx_id}/split_mode", handleSetSplitMode)
	mux.HandleFunc("GET /trx/{trx_id}/level/{level_param}", handleGetLevel)
	mux.HandleFunc("POST /trx/{trx_id}/level/{level_param}", handleSetLevel)
	mux.HandleFunc("GET /trx/{trx_id}/level/list", handleGetLevelList)
	mux.HandleFunc("GET /trx/{trx_id}/tuningstep", handleGetTuningStep)
	mux.HandleFunc("POST /trx/{trx_id}/tuningstep", handleSetTuningStep)
	mux.HandleFunc("GET /trx/{trx_id}/split_vfo", handleGetSplitVFO)
	mux.HandleFunc("POST /trx/{trx_id}/split_vfo", handleSetSplitVFO)

	fmt.Println("Hamlib Go-API läuft auf http://localhost:8080 ...")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Schwerwiegender Server-Fehler: %v\n", err)
	}
}
