package rigctld

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func pollTrx(trxID int, command string) (string, error) {
	targetPort := 4532 + trxID
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", targetPort))
	if err != nil {
		return "", fmt.Errorf("Could not reach rigctld on port %d", targetPort)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "%s\n", strings.TrimSpace(command))

	resp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(resp), nil
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// GET /trx/{trx_id}/frequency -> get_trx_frequency
func HandleGetFrequency(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id")) // ID aus der URL holen

	rawFreq, err := pollTrx(trxID, "f")
	if err != nil {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}

	// build_response() Äquivalent als Map
	WriteJSON(w, http.StatusOK, map[string]string{"freq": rawFreq})
}

// POST /trx/{trx_id}/frequency -> set_trx_frequency
func HandleSetFrequency(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	// RequestBody dynamisch in eine Map dekodieren (json_decode)
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Ungültiger Body"})
		return
	}

	// Befehl an pollTrx übergeben (z.B. "F 145425000")
	_, err := pollTrx(trxID, "F "+body["newValue"])
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	// Leere Erfolgsantwort senden (array())
	WriteJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/mode -> get_trx_mode

// GET /trx/{trx_id}/mode -> get_trx_mode
func HandleGetMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "m")
	if err != nil || len(output) < 2 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	// FIX: Explicitly cast the outputs to string if they are treated as bytes/runes
	WriteJSON(w, http.StatusOK, map[string]string{
		"mode":     string(output[0]),
		"passband": string(output[1]),
	})
}

// POST /trx/{trx_id}/mode -> set_trx_mode
func HandleSetMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	// Decode incoming JSON dynamically (json_decode)
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["mode"] == "" || body["passband"] == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	// Construct and send command (e.g., "M LSB 2400")
	fullCmd := fmt.Sprintf("M %s %s", body["mode"], body["passband"])
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	// Return empty response array()
	WriteJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/split_frequency -> get_trx_split_frequency
func HandleGetSplitFrequency(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "i")
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	// Returns the split frequency (qrg) from the first line
	WriteJSON(w, http.StatusOK, map[string]string{
		"qrg": string(output[0]),
	})
}

// POST /trx/{trx_id}/split_frequency -> set_trx_split_frequency
func HandleSetSplitFrequency(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["newValue"] == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	// Command to set split frequency (e.g., "I 145425000")
	// fullCmd := fmt.Sprintf("I %s", body["newValue"])
	// _, err := pollTrx(trxID, fullCmd)

	fullCmd := fmt.Sprintf("F_VFO VFOB %s", body["newValue"])
	_, err := pollTrx(trxID, fullCmd)

	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/split_mode -> get_trx_split_mode
func HandleGetSplitMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "x")
	if err != nil || len(output) < 2 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	// Line 1: Split Mode, Line 2: Split Passband
	WriteJSON(w, http.StatusOK, map[string]string{
		"mode":     string(output[0]),
		"passband": string(output[1]),
	})
}

// POST /trx/{trx_id}/split_mode -> set_trx_split_mode
func HandleSetSplitMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["mode"] == "" || body["passband"] == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	// Command to set split mode and passband (e.g., "X LSB 2700")
	fullCmd := fmt.Sprintf("X %s %s", body["mode"], body["passband"])
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/split_frequency_mode -> get_trx_split_frequency_mode
func HandleGetSplitFrequencyMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "k")
	if err != nil || len(output) < 3 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	// Line 1: Frequency, Line 2: Mode, Line 3: Passband
	WriteJSON(w, http.StatusOK, map[string]string{
		"frequency": string(output[0]),
		"mode":      string(output[1]),
		"passband":  string(output[2]),
	})
}

// POST /trx/{trx_id}/split_frequency_mode -> set_trx_split_frequency_mode
func HandleSetSplitFrequencyMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["frequency"] == "" || body["mode"] == "" || body["passband"] == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	// Command to set all three split parameters (e.g., "K 145425000 LSB 2700")
	fullCmd := fmt.Sprintf("K %s %s %s", body["frequency"], body["mode"], body["passband"])
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/split_vfo -> get_trx_split_vfo
func HandleGetSplitVFO(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "s")
	if err != nil || len(output) < 2 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response from rigctld"})
		return
	}

	// Line 1: Split Mode (0 or 1), Line 2: TX VFO
	WriteJSON(w, http.StatusOK, map[string]string{
		"split_mode": string(output[0]),
		"tx_vfo":     string(output[1]),
	})
}

// POST /trx/{trx_id}/split_vfo -> set_trx_split_vfo
func HandleSetSplitVFO(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["split_mode"] == "" || body["tx_vfo"] == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	// Command: S <Split Mode> <TX VFO> (e.g., "S 1 VFOB")
	fullCmd := fmt.Sprintf("S %s %s", body["split_mode"], body["tx_vfo"])
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{})
}
