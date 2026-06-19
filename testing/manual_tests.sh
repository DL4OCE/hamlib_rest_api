#!/bin/bash

# ==============================================================================
# CONFIGURATION
# ==============================================================================
BASE_URL="http://localhost:8080/trx/1"
HEADER="Content-Type: application/json"

echo "=== STARTING HAMLIB REST API MANUAL CURL TESTS ==="
echo "Target URL: $BASE_URL"
echo "--------------------------------------------------"

# ==============================================================================
# 1. CORE FUNCTIONS (FREQUENCY & MODE)
# ==============================================================================
echo -e "\n[1. Frequency & Mode]"

echo "GET /frequency:"
curl -s -X GET "$BASE_URL/frequency"
echo -e "\nPOST /frequency:"
curl -s -X POST "$BASE_URL/frequency" -H "$HEADER" -d '{"newValue":"145500000"}'

echo -e "\n\nGET /mode:"
curl -s -X GET "$BASE_URL/mode"
echo -e "\nPOST /mode:"
curl -s -X POST "$BASE_URL/mode" -H "$HEADER" -d '{"mode":"USB","passband":"2400"}'


# ==============================================================================
# 2. SPLIT REQS
# ==============================================================================
echo -e "\n\n[2. Split Settings]"

echo "GET /split_frequency:"
curl -s -X GET "$BASE_URL/split_frequency"
echo -e "\nPOST /split_frequency:"
curl -s -X POST "$BASE_URL/split_frequency" -H "$HEADER" -d '{"newValue":"145525000"}'

echo -e "\n\nGET /split_mode:"
curl -s -X GET "$BASE_URL/split_mode"
echo -e "\nPOST /split_mode:"
curl -s -X POST "$BASE_URL/split_mode" -H "$HEADER" -d '{"mode":"LSB","passband":"2700"}'

echo -e "\n\nGET /split_frequency_mode:"
curl -s -X GET "$BASE_URL/split_frequency_mode"
echo -e "\nPOST /split_frequency_mode:"
curl -s -X POST "$BASE_URL/split_frequency_mode" -H "$HEADER" -d '{"frequency":"145450000","mode":"CW","passband":"500"}'

echo -e "\n\nGET /split_vfo:"
curl -s -X GET "$BASE_URL/split_vfo"
echo -e "\nPOST /split_vfo:"
curl -s -X POST "$BASE_URL/split_vfo" -H "$HEADER" -d '{"split_mode":"1","tx_vfo":"VFOB"}'


# ==============================================================================
# 3. TUNING & LEVELS
# ==============================================================================
echo -e "\n\n[3. Tuning & Levels]"

echo "GET /tuningstep:"
curl -s -X GET "$BASE_URL/tuningstep"
echo -e "\nPOST /tuningstep:"
curl -s -X POST "$BASE_URL/tuningstep" -H "$HEADER" -d '{"newValue":"500"}'

echo -e "\n\nGET /level/list:"
curl -s -X GET "$BASE_URL/level/list"
echo -e "\nGET /level/RFPOWER:"
curl -s -X GET "$BASE_URL/level/RFPOWER"
echo -e "\nPOST /level/RFPOWER:"
curl -s -X POST "$BASE_URL/level/RFPOWER" -H "$HEADER" -d '{"newValue":"0.500000"}'


# ==============================================================================
# 4. RIG CAPABILITIES (FUNCTIONS, PARAMS, SCANS, TRANSCEIVE)
# ==============================================================================
echo -e "\n\n[4. Functions, Parameters & Scans]"

echo "GET /function/list:"
curl -s -X GET "$BASE_URL/function/list"
echo -e "\nGET /function/RIT:"
curl -s -X GET "$BASE_URL/function/RIT"
echo -e "\nPOST /function/RIT:"
curl -s -X POST "$BASE_URL/function/RIT" -H "$HEADER" -d '{"newValue":"1"}'

echo -e "\n\nGET /parameter/list:"
curl -s -X GET "$BASE_URL/parameter/list"
echo -e "\nGET /parameter/APO:"
curl -s -X GET "$BASE_URL/parameter/APO"
echo -e "\nPOST /parameter/APO:"
curl -s -X POST "$BASE_URL/parameter/APO" -H "$HEADER" -d '{"newValue":"60"}'

echo -e "\n\nGET /scan/list:"
curl -s -X GET "$BASE_URL/scan/list"
#echo -e "\nGET /scan/MEM:"
#curl -s -X GET "$BASE_URL/scan/MEM"
echo -e "\nPOST /scan/MEM:"
curl -s -X POST "$BASE_URL/scan/MEM" -H "$HEADER" -d '{"newValue":"1"}'

echo -e "\n\nGET /transceive/list:"
curl -s -X GET "$BASE_URL/transceive/list"
echo -e "\nGET /transceive:"
curl -s -X GET "$BASE_URL/transceive"
echo -e "\nPOST /transceive:"
curl -s -X POST "$BASE_URL/transceive" -H "$HEADER" -d '{"newValue":"1"}'


# ==============================================================================
# 5. REPEATER, TONES & VFO
# ==============================================================================
echo -e "\n\n[5. Repeater & Tones]"

echo "GET /repeater/shift:"
curl -s -X GET "$BASE_URL/repeater/shift"
echo -e "\nPOST /repeater/shift:"
curl -s -X POST "$BASE_URL/repeater/shift" -H "$HEADER" -d '{"newValue":"+"}'

echo -e "\n\nGET /repeater/offset:"
curl -s -X GET "$BASE_URL/repeater/offset"
echo -e "\nPOST /repeater/offset:"
curl -s -X POST "$BASE_URL/repeater/offset" -H "$HEADER" -d '{"newValue":"600000"}'

echo -e "\n\nGET /tone/ctcss:"
curl -s -X GET "$BASE_URL/tone/ctcss"
echo -e "\nPOST /tone/ctcss:"
curl -s -X POST "$BASE_URL/tone/ctcss" -H "$HEADER" -d '{"newValue":"88.5"}'

echo -e "\n\nGET /tone/dcs:"
curl -s -X GET "$BASE_URL/tone/dcs"
echo -e "\nPOST /tone/dcs:"
curl -s -X POST "$BASE_URL/tone/dcs" -H "$HEADER" -d '{"newValue":"023"}'

echo -e "\n\nGET /vfo:"
curl -s -X GET "$BASE_URL/vfo"
echo -e "\nPOST /vfo:"
curl -s -X POST "$BASE_URL/vfo" -H "$HEADER" -d '{"newValue":"VFOA"}'


# ==============================================================================
# 6. HARDWARE STATES (PTT, MEMORY, CHANNEL, ANTENNA)
# ==============================================================================
echo -e "\n\n[6. Hardware States]"

echo "GET /ptt:"
curl -s -X GET "$BASE_URL/ptt"
echo -e "\nPOST /ptt:"
curl -s -X POST "$BASE_URL/ptt" -H "$HEADER" -d '{"newValue":"0"}'

echo -e "\n\nGET /memory:"
curl -s -X GET "$BASE_URL/memory"
echo -e "\nPOST /memory:"
curl -s -X POST "$BASE_URL/memory" -H "$HEADER" -d '{"newValue":"1"}'

#echo -e "\n\nGET /channel:"
#curl -s -X GET "$BASE_URL/channel"
#echo -e "\nPOST /channel:"
#curl -s -X POST "$BASE_URL/channel" -H "$HEADER" -d '{"newValue":"12"}'

echo -e "\n\nGET /info:"
curl -s -X GET "$BASE_URL/info"

echo -e "\n\nGET /rit:"
curl -s -X GET "$BASE_URL/rit"
echo -e "\nPOST /rit:"
curl -s -X POST "$BASE_URL/rit" -H "$HEADER" -d '{"newValue":"0"}'

echo -e "\n\nGET /xit:"
curl -s -X GET "$BASE_URL/xit"
echo -e "\nPOST /xit:"
curl -s -X POST "$BASE_URL/xit" -H "$HEADER" -d '{"newValue":"0"}'

echo -e "\n\nGET /antenna:"
curl -s -X GET "$BASE_URL/antenna"
#echo -e "\nPOST /antenna:"
#curl -s -X POST "$BASE_URL/antenna" -H "$HEADER" -d '{"newValue":"A"}'


# ==============================================================================
# 7. MORSE, RIG INFO & CACHE
# ==============================================================================
echo -e "\n\n[7. Morse, Rig State & Conversions]"

echo "POST /morse:"
curl -s -X POST "$BASE_URL/morse" -H "$HEADER" -d '{"text":"CQ CQ DE DL4OCE"}'
echo -e "\nPOST /morse/stop:"
curl -s -X POST "$BASE_URL/morse/stop"

#echo -e "\n\nGET /morse/wait:"
#curl -s -X GET "$BASE_URL/morse/wait"

#echo -e "\n\nGET /dcd:"
#curl -s -X GET "$BASE_URL/dcd"

#echo -e "\n\nGET /twiddle:"
#curl -s -X GET "$BASE_URL/twiddle"
echo -e "\nPOST /twiddle:"
curl -s -X POST "$BASE_URL/twiddle" -H "$HEADER" -d '{"newValue":"1"}'

#echo -e "\n\nGET /cache:"
#curl -s -X GET "$BASE_URL/cache"
echo -e "\nPOST /cache:"
curl -s -X POST "$BASE_URL/cache" -H "$HEADER" -d '{"newValue":"10"}'

echo -e "\n\nGET /capabilities:"
curl -s -X GET "$BASE_URL/capabilities"

echo -e "\n\nGET /configuration:"
curl -s -X GET "$BASE_URL/configuration"

echo -e "\n\nPOST /state/dump:"
curl -s -X POST "$BASE_URL/state/dump"

#echo -e "\n\nGET /rig_info:"
#curl -s -X GET "$BASE_URL/rig_info"

#echo -e "\n\nGET /modes:"
#curl -s -X GET "$BASE_URL/modes"

#echo -e "\n\nGET /power_state:"
#curl -s -X GET "$BASE_URL/power_state"
echo -e "\nPOST /power_state:"
curl -s -X POST "$BASE_URL/power_state" -H "$HEADER" -d '{"newValue":"1"}'


# ==============================================================================
# 8. SQL EXTENSIONS & DTMF
# ==============================================================================
echo -e "\n\n[8. SQL Extensions & DTMF]"

#echo "GET /sql/ctcss:"
#curl -s -X GET "$BASE_URL/sql/ctcss"
echo -e "\nPOST /sql/ctcss:"
curl -s -X POST "$BASE_URL/sql/ctcss" -H "$HEADER" -d '{"newValue":"1"}'

#echo -e "\n\nGET /sql/dcs:"
#curl -s -X GET "$BASE_URL/sql/dcs"
echo -e "\nPOST /sql/dcs:"
curl -s -X POST "$BASE_URL/sql/dcs" -H "$HEADER" -d '{"newValue":"1"}'

echo -e "\n\nGET /dtmf:"
curl -s -X GET "$BASE_URL/dtmf"
echo -e "\nPOST /dtmf:"
curl -s -X POST "$BASE_URL/dtmf" -H "$HEADER" -d '{"newValue":"123"}'

echo -e "\n\nPOST /voice_mem:"
curl -s -X POST "$BASE_URL/voice_mem" -H "$HEADER" -d '{"newValue":"1"}'


# ==============================================================================
# 9. SPECIAL POWER CONVERSIONS (BODY PAYLOAD CONVERTED TO POST)
# ==============================================================================
echo -e "\n\n[9. Power Conversion Endpoints]"

echo "POST /power/to_factor:"
curl -s -X POST "$BASE_URL/power/to_factor" -H "$HEADER" -d '{"power_mW":"50000", "frequency":"145500000", "mode":"USB"}'

echo -e "\n\nPOST /power/to_mw:"
curl -s -X POST "$BASE_URL/power/to_mw" -H "$HEADER" -d '{"power_factor":"0.50", "frequency":"145500000", "mode":"USB"}'


# ==============================================================================
# 10. RAW COMMANDS (USE WITH CAUTION)
# ==============================================================================
#echo -e "\n\n[10. Raw Commands]"

#echo "POST /raw (Command: f):"
#curl -s -X POST "$BASE_URL/raw" -H "$HEADER" -d '{"raw_command":"f"}'

#echo -e "\n\nPOST /raw_rx (Command: F 145500000):"
#curl -s -X POST "$BASE_URL/raw_rx" -H "$HEADER" -d '{"raw_command":"F 145500000", "number_of_expected_rx_bytes":"0"}'

echo -e "\n\n=== TESTS COMPLETE ==="