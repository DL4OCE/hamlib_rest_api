

#!/bin/bash
TRX_ID=$1
CONFIG_FILE="/etc/hamlib_rest_api/rigctld.json"

# Werte aus der JSON-Datei für die spezifische ID extrahieren
MODEL=$(jq -r ".\"$TRX_ID\".model" $CONFIG_FILE)
DEVICE=$(jq -r ".\"$TRX_ID\".device" $CONFIG_FILE)
BAUD=$(jq -r ".\"$TRX_ID\".baud" $CONFIG_FILE)
CIV=$(jq -r ".\"$TRX_ID\".civaddr" "$CONFIG_FILE")
PORT=$((4532 + TRX_ID))

echo "Starte rigctld für TRX $TRX_ID auf Port $PORT mit Modell $MODEL an $DEVICE $CIV_STR..."
CMD="exec rigctld -m $MODEL -r $DEVICE -s $BAUD -t $PORT"

if [ ! -z "$CIV" ] && [ "$CIV" != "null" ] && [ "$CIV" != "" ]; then
    CMD="$CMD -C civaddr=$CIV"
    CIV_STR="(CIV address: $CIV)"
fi

$CMD