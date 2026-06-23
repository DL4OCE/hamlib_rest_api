#!/bin/bash

test_get_api() {
    echo -e "GET $1 $2"
    curl -s -X GET "$1"
    echo -e "------------------------------------------------------------------------------------------------------------------------------------------------------"
}

test_post_api() {
    echo -e "POST $1 $2"
    if [ -n "$3" ]; then
        curl -s -X POST "$1" -H "Content-Type: application/json" -d "$3"
    else
        curl -s -X POST "$1"
    fi
    echo -e "------------------------------------------------------------------------------------------------------------------------------------------------------"
}

# ==============================================================================
# CONFIGURATION
# ==============================================================================
API_URL="http://localhost:8080"
API_BASE="/api/v1"
RIG_BASE_URL="$API_URL$API_BASE/rigs/1"
ROTATOR_BASE_URL="$API_URL$API_BASE/rotators/1"
HEADER="Content-Type: application/json"

echo "=== STARTING HAMLIB REST API MANUAL CURL TESTS ==="
echo "Target API URL: $API_URL"
echo "Target TRX URL: $BASE_URL"
echo "--------------------------------------------------"

# ==============================================================================
# === NEW === SYSTEM & SERVICE MANAGEMENT (GLOBAL ENDPOINTS)
# ==============================================================================
echo -e "[0. System & Service Management]"
test_get_api $API_URL$API_BASE/devices/rigs "(list rigs (transceivers))"
test_get_api $API_URL$API_BASE/devices/rotators "(list rotators)"
test_post_api $RIG_BASE_URL/service/stop "(Stop TRX 1)"
test_post_api $RIG_BASE_URL/service/start "(Start TRX 1)"
test_post_api $ROTATOR_BASE_URL/service/stop "(Stop Rotator 1)"
test_post_api $ROTATOR_BASE_URL/service/start "(Start Rotator 1)"
sleep 2

# ==============================================================================
# 1. CORE FUNCTIONS (FREQUENCY & MODE)
# ==============================================================================

echo -e "\n[1. Frequency & Mode]"
test_get_api $RIG_BASE_URL/frequency "(Get current frequency)"
test_post_api $RIG_BASE_URL/frequency "(Set frequency to 145.500 MHz)" '{"newValue":"145500000"}'
test_get_api $RIG_BASE_URL/mode "(Get current mode)"
test_post_api $RIG_BASE_URL/mode "(Set mode to USB with 2.4kHz passband)" '{"mode":"USB","passband":"2400"}'


# ==============================================================================
# 2. SPLIT REQS
# ==============================================================================
echo -e "\n[2. Split Settings]"
test_get_api $RIG_BASE_URL/split_frequency "(Get current split frequency)"
test_post_api $RIG_BASE_URL/split_frequency "(Set split frequency to 145.525 MHz)" '{"newValue":"145525000"}'

test_get_api $RIG_BASE_URL/split_mode "(Get current split mode)"
test_post_api $RIG_BASE_URL/split_mode "(Set split mode to LSB with 2.7kHz passband)" '{"mode":"LSB","passband":"2700"}'

test_get_api $RIG_BASE_URL/split_frequency_mode "(Get current split mode)"
test_post_api $RIG_BASE_URL/split_frequency_mode "(Set split mode to CW with 500 Hz passband)" '{"mode":"CW","passband":"500"}'

test_get_api $RIG_BASE_URL/split_vfo "(Get current VFO mode)"
test_post_api $RIG_BASE_URL/split_vfo "(Set VFO mode to split enabled and active VFO = VFO B)" '{"split_mode":"1","tx_vfo":"VFOB"}'
exit 0
# ==============================================================================
# 3. TUNING & LEVELS
# ==============================================================================
echo -e "\n\n[3. Tuning & Levels]"


echo "GET /tuningstep:"
curl -s -X GET "$RIG_BASE_URL/tuningstep"
echo -e "\nPOST /tuningstep:"
curl -s -X POST "$RIG_BASE_URL/tuningstep" -H "$HEADER" -d '{"newValue":"500"}'

echo -e "\n\nGET /level/list:"
curl -s -X GET "$RIG_BASE_URL/level/list"
echo -e "\nGET /level/RFPOWER:"
curl -s -X GET "$RIG_BASE_URL/level/RFPOWER"
echo -e "\nPOST /level/RFPOWER:"
curl -s -X POST "$RIG_BASE_URL/level/RFPOWER" -H "$HEADER" -d '{"newValue":"0.500000"}'


# ==============================================================================
# 4. RIG CAPABILITIES (FUNCTIONS, PARAMS, SCANS, TRANSCEIVE)
# ==============================================================================
echo -e "\n\n[4. Functions, Parameters & Scans]"

echo "GET /function/list:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/function/list"
echo -e "\nGET /function/RIT:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/function/RIT"
echo -e "\nPOST /function/RIT:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/function/RIT" -H "$HEADER" -d '{"newValue":"1"}'

echo -e "\n\nGET /parameter/list:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/parameter/list"
echo -e "\nGET /parameter/APO:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/parameter/APO"
echo -e "\nPOST /parameter/APO:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/parameter/APO" -H "$HEADER" -d '{"newValue":"60"}'

echo -e "\n\nGET /scan/list:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/scan/list"
#echo -e "\nGET /scan/MEM:"
#curl -s -X GET "$BASE_URL/rigs/1/scan/MEM"
echo -e "\nPOST /scan/MEM:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/scan/MEM" -H "$HEADER" -d '{"newValue":"1"}'

echo -e "\n\nGET /transceive/list:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/transceive/list"
echo -e "\nGET /transceive:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/transceive"
echo -e "\nPOST /transceive:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/transceive" -H "$HEADER" -d '{"newValue":"1"}'


# ==============================================================================
# 5. REPEATER, TONES & VFO
# ==============================================================================
echo -e "\n\n[5. Repeater & Tones]"

echo "GET /repeater/shift:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/repeater/shift"
echo -e "\nPOST /repeater/shift:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/repeater/shift" -H "$HEADER" -d '{"newValue":"+"}'

echo -e "\n\nGET /repeater/offset:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/repeater/offset"
echo -e "\nPOST /repeater/offset:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/repeater/offset" -H "$HEADER" -d '{"newValue":"600000"}'

echo -e "\n\nGET /tone/ctcss:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/tone/ctcss"
echo -e "\nPOST /tone/ctcss:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/tone/ctcss" -H "$HEADER" -d '{"newValue":"88.5"}'

echo -e "\n\nexport GET /tone/dcs:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/tone/dcs"
echo -e "\nPOST /tone/dcs:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/tone/dcs" -H "$HEADER" -d '{"newValue":"023"}'

echo -e "\n\nGET /vfo:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/vfo"
echo -e "\nPOST /vfo:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/vfo" -H "$HEADER" -d '{"newValue":"VFOA"}'


# ==============================================================================
# 6. HARDWARE STATES (PTT, MEMORY, CHANNEL, ANTENNA)
# ==============================================================================
echo -e "\n\n[6. Hardware States]"

echo "GET /ptt:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/ptt"
echo -e "\nPOST /ptt:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/ptt" -H "$HEADER" -d '{"newValue":"0"}'

echo -e "\n\nGET /memory:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/memory"
echo -e "\nPOST /memory:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/memory" -H "$HEADER" -d '{"newValue":"1"}'

#echo -e "\n\nGET /channel:"
#curl -s -X GET "$RIG_BASE_URL/rigs/1/channel"
#echo -e "\nPOST /channel:"
#curl -s -X POST "$RIG_BASE_URL/rigs/1/channel" -H "$HEADER" -d '{"newValue":"12"}'

echo -e "\n\nGET /info:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/info"

echo -e "\n\nGET /rit:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/rit"
echo -e "\nPOST /rit:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/rit" -H "$HEADER" -d '{"newValue":"0"}'

echo -e "\n\nGET /xit:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/xit"
echo -e "\nPOST /xit:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/xit" -H "$HEADER" -d '{"newValue":"0"}'

echo -e "\n\nGET /antenna:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/antenna"
#echo -e "\nPOST /antenna:"
#curl -s -X POST "$RIG_BASE_URL/rigs/1/antenna" -H "$HEADER" -d '{"newValue":"A"}'


# ==============================================================================
# 7. MORSE, RIG INFO & CACHE
# ==============================================================================
echo -e "\n\n[7. Morse, Rig State & Conversions]"

echo "POST /morse:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/morse" -H "$HEADER" -d '{"text":"CQ CQ DE DL4OCE"}'
echo -e "\nPOST /morse/stop:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/morse/stop"

#echo -e "\n\nGET /morse/wait:"
#curl -s -X GET "$RIG_BASE_URL/rigs/1/morse/wait"

#echo -e "\n\nGET /dcd:"
#curl -s -X GET "$RIG_BASE_URL/rigs/1/dcd"

#echo -e "\n\nGET /twiddle:"
#curl -s -X GET "$RIG_BASE_URL/rigs/1/twiddle"
echo -e "\nPOST /twiddle:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/twiddle" -H "$HEADER" -d '{"newValue":"1"}'

#echo -e "\n\nGET /cache:"
#curl -s -X GET "$RIG_BASE_URL/rigs/1/cache"
echo -e "\nPOST /cache:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/cache" -H "$HEADER" -d '{"newValue":"10"}'

echo -e "\n\nGET /capabilities:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/capabilities"

echo -e "\n\nGET /configuration:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/configuration"

echo -e "\n\nPOST /state/dump:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/state/dump" -H "$HEADER" -d '{}'
#echo -e "\n\nGET /rigs/1/rig_info:"
#curl -s -X GET "$RIG_BASE_URL/rigs/1/rig_info"

#echo -e "\n\nGET /modes:"
#curl -s -X GET "$RIG_BASE_URL/rigs/1/modes"

#echo -e "\n\nGET /power_state:"
#curl -s -X GET "$RIG_BASE_URL/rigs/1/power_state"
echo -e "\nPOST /power_state:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/power_state" -H "$HEADER" -d '{"newValue":"1"}'


# ==============================================================================
# 8. SQL EXTENSIONS & DTMF
# ==============================================================================
echo -e "\n\n[8. SQL Extensions & DTMF]"

#echo "GET /sql/ctcss:"
#curl -s -X GET "$RIG_BASE_URL/rigs/1/sql/ctcss"
echo -e "\nPOST /sql/ctcss:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/sql/ctcss" -H "$HEADER" -d '{"newValue":"1"}'

#echo -e "\n\nGET /sql/dcs:"
#curl -s -X GET "$RIG_BASE_URL/rigs/1/sql/dcs"
echo -e "\nPOST /sql/dcs:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/sql/dcs" -H "$HEADER" -d '{"newValue":"1"}'

echo -e "\n\nGET /dtmf:"
curl -s -X GET "$RIG_BASE_URL/rigs/1/dtmf"
echo -e "\nPOST /dtmf:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/dtmf" -H "$HEADER" -d '{"newValue":"123"}'

echo -e "\n\nPOST /voice_mem:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/voice_mem" -H "$HEADER" -d '{"newValue":"1"}'


# ==============================================================================
# 9. SPECIAL POWER CONVERSIONS (BODY PAYLOAD CONVERTED TO POST)
# ==============================================================================
echo -e "\n\n[9. Power Conversion Endpoints]"

echo "POST /power/to_factor:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/power/to_factor" -H "$HEADER" -d '{"power_mW":"50000", "frequency":"145500000", "mode":"USB"}'

echo -e "\n\nPOST /power/to_mw:"
curl -s -X POST "$RIG_BASE_URL/rigs/1/power/to_mw" -H "$HEADER" -d '{"power_factor":"0.50", "frequency":"145500000", "mode":"USB"}'


# ==============================================================================
# 10. RAW COMMANDS (USE WITH CAUTION)
# ==============================================================================
#echo -e "\n\n[10. Raw Commands]"

#echo "POST /raw (Command: f):"
#curl -s -X POST "$RIG_BASE_URL/rigs/1/raw" -H "$HEADER" -d '{"raw_command":"f"}'

#echo -e "\n\nPOST /raw_rx (Command: F 145500000):"
#curl -s -X POST "$RIG_BASE_URL/rigs/1/raw_rx" -H "$HEADER" -d '{"raw_command":"F 145500000", "number_of_expected_rx_bytes":"0"}'

echo -e "\n\n=== TESTS COMPLETE ==="