#!/bin/bash

test_get_api() { 
    echo -e "GET $1"; curl -s -X GET "$1"; 
    echo -e "\n-----------------------------------"; 
}
test_post_api() { 
    echo -e "POST $1 -d $2"; curl -s -X POST "$1" -H "Content-Type: application/json" -d "$2"; 
    echo -e "\n-----------------------------------"; 
}

API_URL="http://localhost:8080/api/v1"
ROT_URL="$API_URL/rotators/1"

# === Service & List ===
test_get_api "$API_URL/devices/rotators"
test_post_api "$ROT_URL/service/stop" '{}'
test_post_api "$ROT_URL/service/start" '{}'

sleep 2

# === Raw & Control === 
# test_post_api "$ROT_URL/command" '{"command":"p"}'
test_post_api "$ROT_URL/move" '{"direction":"1", "speed":"5"}'
test_post_api "$ROT_URL/stop" '{}'
test_post_api "$ROT_URL/park" '{}'
# test_post_api "$ROT_URL/reset" '{"reset_type":"1"}'

# === Position ===
test_get_api "$ROT_URL/position"
test_post_api "$ROT_URL/position" '{"azimuth":"180", "elevation":"45"}'

# === Metadata ===
# test_get_api "$ROT_URL/info"
test_get_api "$ROT_URL/status"
# test_get_api "$ROT_URL/state"
test_get_api "$ROT_URL/capabilities"

# === DYNAMIC PARAMETER TESTS (LOOPED) ===
LEVELS=("ACCURACY" "SPEED")
for lvl in "${LEVELS[@]}"; do
    test_get_api "$ROT_URL/level/$lvl"
    test_post_api "$ROT_URL/level/$lvl" '{"newValue":"1"}'
done

FUNCTIONS=("ON" "LOCK")
for func in "${FUNCTIONS[@]}"; do
    test_get_api "$ROT_URL/function/$func"
    test_post_api "$ROT_URL/function/$func" '{"newValue":"1"}'
done

PARAMS=("TIMEOUT" "OFFSET")
for p in "${PARAMS[@]}"; do
    test_get_api "$ROT_URL/parameter/$p"
    test_post_api "$ROT_URL/parameter/$p" '{"newValue":"10"}'
done

