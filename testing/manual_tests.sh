#!/bin/bash

# ==============================================================================
# CONFIGURATION
# ==============================================================================
# === CHANGED === (Basis-URL für globale Endpunkte hinzugefügt)
API_URL="http://localhost:8080"
BASE_URL="$API_URL/api/v1/rigs/1"
HEADER="Content-Type: application/json"

echo "=== STARTING HAMLIB REST API MANUAL CURL TESTS ==="
echo "Target API URL: $API_URL"
echo "Target TRX URL: $BASE_URL"
echo "--------------------------------------------------"

# ==============================================================================
# === NEW === SYSTEM & SERVICE MANAGEMENT (GLOBAL ENDPOINTS)
# ==============================================================================
echo -e "\n[0. System & Service Management]"

echo "GET /trxs (List Transceivers):"
curl  -X GET "$API_URL/devices/rigs"

echo -e "\n\nGET /rotators (List Rotators):"
curl -s -X GET "$API_URL/rotators"

# === CHANGED === (Hier wird der angepasste /service/-Pfad für die Rigs getestet)
echo -e "\n\nPOST /rigs/1/service/stop (Stop TRX 1):"
curl -s -X POST "$API_URL/rigs/1/service/stop"
echo -e "\nPOST /rigs/1/service/start (Start TRX 1):"
curl -s -X POST "$API_URL/rigs/1/service/start"

echo -e "\n\nPOST /rotators/1/service/stop (Stop Rotator 1):"
curl -s -X POST "$API_URL/rotators/1/service/stop"
echo -e "\nPOST /rotators/1/service/start (Start Rotator 1):"
curl -s -X POST "$API_URL/rotators/1/service/start"


# ==============================================================================
# 1. CORE FUNCTIONS (FREQUENCY & MODE)
# ==============================================================================
echo -e "\n\n[1. Frequency & Mode]"

echo "GET /frequency:"
curl -s -X GET "$BASE_URL/rigs/1/frequency"
echo -e "\nPOST /frequency:"
curl -s -X POST "$BASE_URL/rigs/1/frequency" -H "$HEADER" -d '{"newValue":"145500000"}'

echo -e "\n\nGET /mode:"
curl -s -X GET "$BASE_URL/rigs/1/mode"
echo -e "\nPOST /mode:"
curl -s -X POST "$BASE_URL/rigs/1/mode" -H "$HEADER" -d '{"mode":"USB","passband":"2400"}'


# ==============================================================================
# 2. SPLIT REQS
# ==============================================================================
echo -e "\n\n[2. Split Settings]"

echo "GET /split_frequency:"
curl -s -X GET "$BASE_URL/rigs/1/split_frequency"
echo -e "\nPOST /split_frequency:"
curl -s -X POST "$BASE_URL/rigs/1/split_frequency" -H "$HEADER" -d '{"newValue":"145525000"}'

echo -e "\n\nGET /split_mode:"
curl -s -X GET "$BASE_URL/rigs/1/split_mode"
echo -e "\nPOST /split_mode:"
curl -s -X POST "$BASE_URL/rigs/1/split_mode" -H "$HEADER" -d '{"mode":"LSB","passband":"2700"}'

echo -e "\n\nGET /split_frequency_mode:"
curl -s -X GET "$BASE_URL/rigs/1/split_frequency_mode"
echo -e "\nPOST /split_frequency_mode:"
curl -s -X POST "$BASE_URL/rigs/1/split_frequency_mode" -H "$HEADER" -d '{"frequency":"145450000","mode":"CW","passband":"500"}'

echo -e "\n\nGET /split_vfo:"
curl -s -X GET "$BASE_URL/rigs/1/split_vfo"
echo -e "\nPOST /split_vfo:"
curl -s -X POST "$BASE_URL/rigs/1/split_vfo" -H "$HEADER" -d '{"split_mode":"1","tx_vfo":"VFOB"}'


# ==============================================================================
# 3. TUNING & LEVELS
# ==============================================================================
echo -e "\n\n[3. Tuning & Levels]"

echo "GET /tuningstep:"
curl -s -X GET "$BASE_URL/rigs/1/tuningstep"
echo -e "\nPOST /tuningstep:"
curl -s -X POST "$BASE_URL/rigs/1/tuningstep" -H "$HEADER" -d '{"newValue":"500"}'

echo -e "\n\nGET /level/list:"
curl -s -X GET "$BASE_URL/rigs/1/level/list"
echo -e "\nGET /level/RFPOWER:"
curl -s -X GET "$BASE_URL/rigs/1/level/RFPOWER"
echo -e "\nPOST /level/RFPOWER:"
curl -s -X POST "$BASE_URL/rigs/1/level/RFPOWER" -H "$HEADER" -d '{"newValue":"0.500000"}'


# ==============================================================================
# 4. RIG CAPABILITIES (FUNCTIONS, PARAMS, SCANS, TRANSCEIVE)
# ==============================================================================
echo -e "\n\n[4. Functions, Parameters & Scans]"

echo "GET /function/list:"
curl -s -X GET "$BASE_URL/rigs/1/function/list"
echo -e "\nGET /function/RIT:"
curl -s -X GET "$BASE_URL/rigs/1/function/RIT"
echo -e "\nPOST /function/RIT:"
curl -s -X POST "$BASE_URL/rigs/1/function/RIT" -H "$HEADER" -d '{"newValue":"1"}'

echo -e "\n\nGET /parameter/list:"
curl -s -X GET "$BASE_URL/rigs/1/parameter/list"
echo -e "\nGET /parameter/APO:"
curl -s -X GET "$BASE_URL/rigs/1/parameter/APO"
echo -e "\nPOST /parameter/APO:"
curl -s -X POST "$BASE_URL/rigs/1/parameter/APO" -H "$HEADER" -d '{"newValue":"60"}'

echo -e "\n\nGET /scan/list:"
curl -s -X GET "$BASE_URL/rigs/1/scan/list"
#echo -e "\nGET /scan/MEM:"
#curl -s -X GET "$BASE_URL/rigs/1/scan/MEM"
echo -e "\nPOST /scan/MEM:"
curl -s -X POST "$BASE_URL/rigs/1/scan/MEM" -H "$HEADER" -d '{"newValue":"1"}'

echo -e "\n\nGET /transceive/list:"
curl -s -X GET "$BASE_URL/rigs/1/transceive/list"
echo -e "\nGET /transceive:"
curl -s -X GET "$BASE_URL/rigs/1/transceive"
echo -e "\nPOST /transceive:"
curl -s -X POST "$BASE_URL/rigs/1/transceive" -H "$HEADER" -d '{"newValue":"1"}'


# ==============================================================================
# 5. REPEATER, TONES & VFO
# ==============================================================================
echo -e "\n\n[5. Repeater & Tones]"

echo "GET /repeater/shift:"
curl -s -X GET "$BASE_URL/rigs/1/repeater/shift"
echo -e "\nPOST /repeater/shift:"
curl -s -X POST "$BASE_URL/rigs/1/repeater/shift" -H "$HEADER" -d '{"newValue":"+"}'

echo -e "\n\nGET /repeater/offset:"
curl -s -X GET "$BASE_URL/rigs/1/repeater/offset"
echo -e "\nPOST /repeater/offset:"
curl -s -X POST "$BASE_URL/rigs/1/repeater/offset" -H "$HEADER" -d '{"newValue":"600000"}'

echo -e "\n\nGET /tone/ctcss:"
curl -s -X GET "$BASE_URL/rigs/1/tone/ctcss"
echo -e "\nPOST /tone/ctcss:"
curl -s -X POST "$BASE_URL/rigs/1/tone/ctcss" -H "$HEADER" -d '{"newValue":"88.5"}'

echo -e "\n\nexport GET /tone/dcs:"
curl -s -X GET "$BASE_URL/rigs/1/tone/dcs"
echo -e "\nPOST /tone/dcs:"
curl -s -X POST "$BASE_URL/rigs/1/tone/dcs" -H "$HEADER" -d '{"newValue":"023"}'

echo -e "\n\nGET /vfo:"
curl -s -X GET "$BASE_URL/rigs/1/vfo"
echo -e "\nPOST /vfo:"
curl -s -X POST "$BASE_URL/rigs/1/vfo" -H "$HEADER" -d '{"newValue":"VFOA"}'


# ==============================================================================
# 6. HARDWARE STATES (PTT, MEMORY, CHANNEL, ANTENNA)
# ==============================================================================
echo -e "\n\n[6. Hardware States]"

echo "GET /ptt:"
curl -s -X GET "$BASE_URL/rigs/1/ptt"
echo -e "\nPOST /ptt:"
curl -s -X POST "$BASE_URL/rigs/1/ptt" -H "$HEADER" -d '{"newValue":"0"}'

echo -e "\n\nGET /memory:"
curl -s -X GET "$BASE_URL/rigs/1/memory"
echo -e "\nPOST /memory:"
curl -s -X POST "$BASE_URL/rigs/1/memory" -H "$HEADER" -d '{"newValue":"1"}'

#echo -e "\n\nGET /channel:"
#curl -s -X GET "$BASE_URL/rigs/1/channel"
#echo -e "\nPOST /channel:"
#curl -s -X POST "$BASE_URL/rigs/1/channel" -H "$HEADER" -d '{"newValue":"12"}'

echo -e "\n\nGET /info:"
curl -s -X GET "$BASE_URL/rigs/1/info"

echo -e "\n\nGET /rit:"
curl -s -X GET "$BASE_URL/rigs/1/rit"
echo -e "\nPOST /rit:"
curl -s -X POST "$BASE_URL/rigs/1/rit" -H "$HEADER" -d '{"newValue":"0"}'

echo -e "\n\nGET /xit:"
curl -s -X GET "$BASE_URL/rigs/1/xit"
echo -e "\nPOST /xit:"
curl -s -X POST "$BASE_URL/rigs/1/xit" -H "$HEADER" -d '{"newValue":"0"}'

echo -e "\n\nGET /antenna:"
curl -s -X GET "$BASE_URL/rigs/1/antenna"
#echo -e "\nPOST /antenna:"
#curl -s -X POST "$BASE_URL/rigs/1/antenna" -H "$HEADER" -d '{"newValue":"A"}'


# ==============================================================================
# 7. MORSE, RIG INFO & CACHE
# ==============================================================================
echo -e "\n\n[7. Morse, Rig State & Conversions]"

echo "POST /morse:"
curl -s -X POST "$BASE_URL/rigs/1/morse" -H "$HEADER" -d '{"text":"CQ CQ DE DL4OCE"}'
echo -e "\nPOST /morse/stop:"
curl -s -X POST "$BASE_URL/rigs/1/morse/stop"

#echo -e "\n\nGET /morse/wait:"
#curl -s -X GET "$BASE_URL/rigs/1/morse/wait"

#echo -e "\n\nGET /dcd:"
#curl -s -X GET "$BASE_URL/rigs/1/dcd"

#echo -e "\n\nGET /twiddle:"
#curl -s -X GET "$BASE_URL/rigs/1/twiddle"
echo -e "\nPOST /twiddle:"
curl -s -X POST "$BASE_URL/rigs/1/twiddle" -H "$HEADER" -d '{"newValue":"1"}'

#echo -e "\n\nGET /cache:"
#curl -s -X GET "$BASE_URL/rigs/1/cache"
echo -e "\nPOST /cache:"
curl -s -X POST "$BASE_URL/rigs/1/cache" -H "$HEADER" -d '{"newValue":"10"}'

echo -e "\n\nGET /capabilities:"
curl -s -X GET "$BASE_URL/rigs/1/capabilities"

echo -e "\n\nGET /configuration:"
curl -s -X GET "$BASE_URL/rigs/1/configuration"

echo -e "\n\nPOST /state/dump:"
curl -s -X POST "$BASE_URL/rigs/1/state/dump" -H "$HEADER" -d '{}'
#echo -e "\n\nGET /rigs/1/rig_info:"
#curl -s -X GET "$BASE_URL/rigs/1/rig_info"

#echo -e "\n\nGET /modes:"
#curl -s -X GET "$BASE_URL/rigs/1/modes"

#echo -e "\n\nGET /power_state:"
#curl -s -X GET "$BASE_URL/rigs/1/power_state"
echo -e "\nPOST /power_state:"
curl -s -X POST "$BASE_URL/rigs/1/power_state" -H "$HEADER" -d '{"newValue":"1"}'


# ==============================================================================
# 8. SQL EXTENSIONS & DTMF
# ==============================================================================
echo -e "\n\n[8. SQL Extensions & DTMF]"

#echo "GET /sql/ctcss:"
#curl -s -X GET "$BASE_URL/rigs/1/sql/ctcss"
echo -e "\nPOST /sql/ctcss:"
curl -s -X POST "$BASE_URL/rigs/1/sql/ctcss" -H "$HEADER" -d '{"newValue":"1"}'

#echo -e "\n\nGET /sql/dcs:"
#curl -s -X GET "$BASE_URL/rigs/1/sql/dcs"
echo -e "\nPOST /sql/dcs:"
curl -s -X POST "$BASE_URL/rigs/1/sql/dcs" -H "$HEADER" -d '{"newValue":"1"}'

echo -e "\n\nGET /dtmf:"
curl -s -X GET "$BASE_URL/rigs/1/dtmf"
echo -e "\nPOST /dtmf:"
curl -s -X POST "$BASE_URL/rigs/1/dtmf" -H "$HEADER" -d '{"newValue":"123"}'

echo -e "\n\nPOST /voice_mem:"
curl -s -X POST "$BASE_URL/rigs/1/voice_mem" -H "$HEADER" -d '{"newValue":"1"}'


# ==============================================================================
# 9. SPECIAL POWER CONVERSIONS (BODY PAYLOAD CONVERTED TO POST)
# ==============================================================================
echo -e "\n\n[9. Power Conversion Endpoints]"

echo "POST /power/to_factor:"
curl -s -X POST "$BASE_URL/rigs/1/power/to_factor" -H "$HEADER" -d '{"power_mW":"50000", "frequency":"145500000", "mode":"USB"}'

echo -e "\n\nPOST /power/to_mw:"
curl -s -X POST "$BASE_URL/rigs/1/power/to_mw" -H "$HEADER" -d '{"power_factor":"0.50", "frequency":"145500000", "mode":"USB"}'


# ==============================================================================
# 10. RAW COMMANDS (USE WITH CAUTION)
# ==============================================================================
#echo -e "\n\n[10. Raw Commands]"

#echo "POST /raw (Command: f):"
#curl -s -X POST "$BASE_URL/rigs/1/raw" -H "$HEADER" -d '{"raw_command":"f"}'

#echo -e "\n\nPOST /raw_rx (Command: F 145500000):"
#curl -s -X POST "$BASE_URL/rigs/1/raw_rx" -H "$HEADER" -d '{"raw_command":"F 145500000", "number_of_expected_rx_bytes":"0"}'

echo -e "\n\n=== TESTS COMPLETE ==="