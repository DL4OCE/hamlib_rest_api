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

// Struktur für die JSON-Antwort
type FrequencyResponse struct {
	TRX       int    `json:"trx"`
	Port      int    `json:"port"`
	Frequency string `json:"frequency"`
	Status    string `json:"status"`
}

// Struktur für Fehler-Antworten
type ErrorResponse struct {
	Error  string `json:"error"`
	Status string `json:"status"`
}

func getFrequencyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. TRX-ID aus der URL extrahieren (Platzhalter {trx_id})
	trxParam := r.PathValue("trx_id")
	trxID, err := strconv.Atoi(trxParam)
	if err != nil || trxID < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Ungültige TRX-ID im Pfad", Status: "error"})
		return
	}

	// 2. Zielport für diesen spezifischen rigctld berechnen
	targetPort := 4532 + trxID
	targetAddress := fmt.Sprintf("localhost:%d", targetPort)

	// 3. TCP-Verbindung zum rigctld aufbauen (mit Timeout)
	conn, err := net.Dial("tcp", targetAddress)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:  fmt.Sprintf("Verbindung zu rigctld auf Port %d fehlgeschlagen. Läuft der Dienst?", targetPort),
			Status: "error",
		})
		return
	}
	defer conn.Close()

	// 4. Hamlib-Befehl 'f' (Frequenz abfragen) senden
	_, err = fmt.Fprint(conn, "f\n")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Fehler beim Senden an rigctld", Status: "error"})
		return
	}

	// 5. Antwort von rigctld lesen
	reader := bufio.NewReader(conn)
	rawResponse, err := reader.ReadString('\n')
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Fehler beim Lesen von rigctld", Status: "error"})
		return
	}

	// Antwort säubern (Whitespace/Newlines entfernen)
	frequency := strings.TrimSpace(rawResponse)

	// 6. JSON-Ergebnis generieren und ausgeben
	response := FrequencyResponse{
		TRX:       trxID,
		Port:      targetPort,
		Frequency: frequency,
		Status:    "ok",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()

	// Definition der Route mit dem Platzhalter {trx_id}
	// Wichtig: Seit Go 1.22 kann die HTTP-Methode direkt vorangestellt werden
	mux.HandleFunc("GET /api/v1/trx/{trx_id}/frequency", getFrequencyHandler)

	fmt.Println("=== Hamlib REST-API Wrapper ===")
	fmt.Println("Server startet auf http://localhost:8080 ...")

	// Server auf Port 8080 starten
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Server-Fehler: %v\n", err)
	}
}
