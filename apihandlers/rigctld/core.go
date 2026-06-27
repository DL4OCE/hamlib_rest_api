package rigctld

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func pollTrx(trxID int, command string) ([]string, error) {
	targetPort := 4499 + trxID
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", targetPort))
	if err != nil {
		return nil, fmt.Errorf("Could not reach rigctld on port %d", targetPort)
	}
	defer conn.Close()

	_, err = fmt.Fprintf(conn, "%s\n", strings.TrimSpace(command))
	if err != nil {
		return nil, fmt.Errorf("failed to send command: %w", err)
	}

	var lines []string
	scanner := bufio.NewScanner(conn)

	for {
		conn.SetReadDeadline(time.Now().Add(50 * time.Millisecond))

		if !scanner.Scan() {
			break
		}

		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		if strings.HasPrefix(text, "RPRT") {
			break
		}

		lines = append(lines, text)
	}

	conn.SetReadDeadline(time.Time{})

	if err := scanner.Err(); err != nil && !os.IsTimeout(err) {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return lines, nil
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// GET /trx/{trx_id}/frequency -> get_trx_frequency
func HandleGetFrequency(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "f")
	if err != nil || len(output) < 1 {
		errMsg := "Failed to get frequency from rigctld"
		if err != nil {
			errMsg = err.Error()
		}
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": errMsg})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{"freq": output[0]})
}

// POST /trx/{trx_id}/frequency -> set_trx_frequency
func HandleSetFrequency(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))
	// var body map[string]string
	var body struct {
		NewValue string `json:"newValue"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON body"})
		return
	}
	_, err := pollTrx(trxID, "F "+body.NewValue)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/mode -> get_trx_mode
func HandleGetMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "m")

	// //
	// if err != nil {
	// 	fmt.Printf("DEBUG: pollTrx error: %v\n", err)
	// } else {
	// 	fmt.Printf("DEBUG: Received %d lines: %v\n", len(output), output)
	// }
	// //

	if err != nil || len(output) < 2 {
		errMsg := "Failed to get mode"
		if err != nil {
			errMsg = err.Error()
		}
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": errMsg})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{
		"mode":     output[0],
		"passband": output[1],
	})
}

// POST /trx/{trx_id}/mode -> set_trx_mode
func HandleSetMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	var body struct {
		Mode     string `json:"mode"`
		Passband string `json:"passband"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Mode == "" || body.Passband == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	fullCmd := fmt.Sprintf("M %s %s", body.Mode, body.Passband)
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{})
}

// GET /trx/{trx_id}/split_frequency -> get_trx_split_frequency
func HandleGetSplitFrequency(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	output, err := pollTrx(trxID, "i")
	if err != nil || len(output) < 1 {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"qrg": output[0]})
}

// POST /trx/{trx_id}/split_frequency -> set_trx_split_frequency
func HandleSetSplitFrequency(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	var body struct {
		NewValue string `json:"newValue"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.NewValue == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	fullCmd := fmt.Sprintf("F_VFO VFOB %s", body.NewValue)
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
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{
		"mode":     output[0],
		"passband": output[1],
	})
}

// POST /trx/{trx_id}/split_mode -> set_trx_split_mode
func HandleSetSplitMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	var body struct {
		Mode     string `json:"mode"`
		Passband string `json:"passband"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Mode == "" || body.Passband == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	fullCmd := fmt.Sprintf("X %s %s", body.Mode, body.Passband)
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
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{
		"frequency": output[0],
		"mode":      output[1],
		"passband":  output[2],
	})
}

// POST /trx/{trx_id}/split_frequency_mode -> set_trx_split_frequency_mode
func HandleSetSplitFrequencyMode(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	var body struct {
		Frequency string `json:"frequency"`
		Mode      string `json:"mode"`
		Passband  string `json:"passband"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Frequency == "" || body.Mode == "" || body.Passband == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	fullCmd := fmt.Sprintf("K %s %s %s", body.Frequency, body.Mode, body.Passband)
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
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "Invalid response"})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{
		"split_mode": output[0],
		"tx_vfo":     output[1],
	})
}

// POST /trx/{trx_id}/split_vfo -> set_trx_split_vfo
func HandleSetSplitVFO(w http.ResponseWriter, r *http.Request) {
	trxID, _ := strconv.Atoi(r.PathValue("trx_id"))

	var body struct {
		SplitMode string `json:"split_mode"`
		TxVFO     string `json:"tx_vfo"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.SplitMode == "" || body.TxVFO == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid or incomplete body JSON"})
		return
	}

	fullCmd := fmt.Sprintf("S %s %s", body.SplitMode, body.TxVFO)
	_, err := pollTrx(trxID, fullCmd)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{})
}
