#!/bin/bash
set -e

if [ "$EUID" -ne 0 ]; then
  echo "Please run as root (sudo ./update_rigctld_services.sh)"
  exit 1
fi

JSON_FILE="rigctld.json"

if [ ! -f "$JSON_FILE" ]; then
    echo "Error: $JSON_FILE does not exist!"
    exit 1
fi

mkdir -p /etc/hamlib_rest_api
cp $JSON_FILE /etc/hamlib_rest_api/
chmod 644 /etc/hamlib_rest_api/$JSON_FILE

ACTIVE_SERVICES=$(systemctl list-units --all --plain --no-legend "rigctld@*" | awk '{print $1}')
for service in $ACTIVE_SERVICES; do
    echo "Stopping and deactivating servicee $service ..."
    systemctl stop "$service"
    systemctl disable "$service"
done

systemctl daemon-reload

TRX_IDS=$(jq -r '.[] | .id' "$JSON_FILE")
if [ -z "$TRX_IDS" ]; then
    echo "Warning: No IDs found in $JSON_FILE."
    exit 0
fi

for id in $TRX_IDS; do
    echo "  -> Aktiviere und starte Dienst für TRX ID: $id (rigctld@$id.service)"
    systemctl enable rigctld@"$id".service
    systemctl start rigctld@"$id".service
done

echo "=== rigctld services overview ==="
systemctl list-units "rigctld@*" --plain --no-legend

#cp hamlib_rest_api.service /etc/systemd/system
#cp rigctld@.service /etc/systemd/system

#systemctl enable --now rigctld@*.service
# hamlib_rest_api.service

