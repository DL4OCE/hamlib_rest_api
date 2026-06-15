# --- TC: negative test (invalid TRXID) ---
GET http://localhost:8080/trx/99/frequency

## Das API sollte mit einem Fehler antworten
HTTP 503
[Asserts]
jsonpath "$.error" contains "nicht erreichbar"
