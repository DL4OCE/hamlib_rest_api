#!/bin/bash
set -e

if [ "$EUID" -ne 0 ]; then
  echo "Please run as root (sudo ./install.sh)"
  exit 1
fi

apt update && sudo apt install -y libhamlib-utils jq

cp rigctld@.service /etc/systemd/system/
mkdir -p /etc/hamlib_rest_api

echo "Stopping rigctld services..."

systemctl list-units "rigctld@*" --plain --no-legend

ACTIVE_SERVICES=$(systemctl list-units --all --plain --no-legend "rigctld@*" | awk '{print $1}')
for service in $ACTIVE_SERVICES; do
    if [ -n "$service" ]; then
        echo "Stopping and deactivating $service ..."
        systemctl stop "$service" || true
        systemctl disable "$service" || true
    fi
done

systemctl daemon-reload

echo "Please modify rigctld.config according to your needs and run update_rigctld_services.sh"
