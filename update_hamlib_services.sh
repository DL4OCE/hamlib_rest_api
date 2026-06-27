#!/bin/bash
set -e

if [ "$EUID" -ne 0 ]; then
  echo "Please run as root (sudo ./update_hamlib_services.sh)"
  exit 1
fi

mkdir -p /etc/hamlib_rest_api

JSON_FILE_RIGCTLD="config/rigctld.json"
JSON_FILE_ROTCTLD="config/rotctld.json"

if [ ! -f "$JSON_FILE_RIGCTLD" ]; then
    echo "Error: $JSON_FILE_RIGCTLD does not exist!"
    exit 1
fi

if [ ! -f "$JSON_FILE_ROTCTLD" ]; then
    echo "Error: $JSON_FILE_ROTCTLD does not exist!"
    exit 1
fi

cp $JSON_FILE_RIGCTLD $JSON_FILE_ROTCTLD /etc/hamlib_rest_api/
chmod 644 /etc/hamlib_rest_api/*

ACTIVE_SERVICES=$(systemctl list-units --all --plain --no-legend "rigctld@*" | awk '{print $1}')
for service in $ACTIVE_SERVICES; do
    echo "Stopping and deactivating service $service ..."
    systemctl stop "$service"
    systemctl disable "$service"
done

ACTIVE_SERVICES=$(systemctl list-units --all --plain --no-legend "rotctld@*" | awk '{print $1}')
for service in $ACTIVE_SERVICES; do
    echo "Stopping and deactivating service $service ..."
    systemctl stop "$service"
    systemctl disable "$service"
done

systemctl daemon-reload

TRX_IDS=$(jq -r '.[] | .id' "$JSON_FILE_RIGCTLD")
if [ -z "$TRX_IDS" ]; then
    echo "Warning: No TRX IDs found in $JSON_FILE_RIGCTLD."
    exit 0
fi

ROT_IDS=$(jq -r '.[] | .id' "$JSON_FILE_ROTCTLD")
if [ -z "$ROT_IDS" ]; then
    echo "Warning: No rotator IDs found in $JSON_FILE_ROTCTLD."
    exit 0
fi

for id in $TRX_IDS; do
    echo "Activating and starting service for TRX ID: $id (rigctld@$id.service)"
    systemctl enable rigctld@"$id".service
    systemctl start rigctld@"$id".service
done

for id in $ROT_IDS; do
    echo "Activating and starting service for rotator ID: $id (rotctld@$id.service)"
    systemctl enable rotctld@"$id".service
    systemctl start rotctld@"$id".service
done

echo -e "\n=== rigctld services overview ==="
systemctl list-units "rigctld@*" --plain --no-legend
echo -e "\n=== rotctld services overview ==="
systemctl list-units "rotctld@*" --plain --no-legend
