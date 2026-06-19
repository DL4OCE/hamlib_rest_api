#!/bin/bash
set -e

if [ "$EUID" -ne 0 ]; then
  echo "Please run as root (sudo ./install.sh)"
  exit 1
fi

REAL_USER=${SUDO_USER:-$USER}
SUDOERS_FILE="/etc/sudoers.d/hamlib_rest_api"
rm -f "$SUDOERS_FILE"
cat << EOF > "$SUDOERS_FILE"
# Auto-generated sudoers file for hamlib_rest_api - DO NOT EDIT
$REAL_USER ALL=(ALL) NOPASSWD: /usr/bin/systemctl start rigctld@*, /usr/bin/systemctl stop rigctld@*
EOF
chmod 0440 "$SUDOERS_FILE"
echo "Wrote sudoers file to $SUDOERS_FILE for user $REAL_USER"
systemctl daemon-reload

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
