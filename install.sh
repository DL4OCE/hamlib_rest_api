#!/bin/bash
set -e

if [ "$EUID" -ne 0 ]; then
  echo "Please run as root (sudo ./install.sh)"
  exit 1
fi

# patch sudoers file to allow current user to start/stop rigctld services
REAL_USER=${SUDO_USER:-$USER}
SUDOERS_FILE="/etc/sudoers.d/hamlib_rest_api"
rm -f "$SUDOERS_FILE"
cat << EOF > "$SUDOERS_FILE"
# Auto-generated sudoers file for hamlib_rest_api - DO NOT EDIT
$REAL_USER ALL=(ALL) NOPASSWD: /usr/bin/systemctl start rigctld@*, /usr/bin/systemctl stop rigctld@*
EOF
chmod 0440 "$SUDOERS_FILE"
echo "Wrote sudoers file to $SUDOERS_FILE for user $REAL_USER"

# install dependencies
apt update && sudo apt install -y libhamlib-utils jq curl tar

# get latest binary release from github
REPO="DL4OCE/hamlib_rest_api"
BINARY_NAME="hamlib_rest_api"
INSTALL_DIR="/usr/local/bin"
ARCH=$(uname -m)
case "$ARCH" in
    x86_64)  GOARCH="amd64" ;;
    aarch64) GOARCH="arm64" ;;
    armv7l)  GOARCH="arm" ;; # Fallback for older 32-bit Pis if compiled
    *) echo "Fehler: Unbekannte Architektur $ARCH"; exit 1 ;;
esac
echo "Detected architecture: $ARCH, using GOARCH=$GOARCH"
API_URL="https://api.github.com/repos/$REPO/releases/latest"
DOWNLOAD_URL=$(curl -s "$API_URL" | jq -r ".assets[] | select(.name | contains(\"linux\") and contains(\"$GOARCH\") and endswith(\"tar.gz\")) | .browser_download_url")

if [ -z "$DOWNLOAD_URL" ] || [ "$DOWNLOAD_URL" = "null" ]; then
    echo "Fehler: Konnte kein passendes Linux-$GOARCH Release-Asset auf GitHub finden!"
    exit 1
fi
echo "Downloading latest binary release from $DOWNLOAD_URL ..."
curl -sL "$DOWNLOAD_URL" | tar -xz -C "$INSTALL_DIR" "$BINARY_NAME"
chmod 755 "$INSTALL_DIR/$BINARY_NAME"
echo "Installed $BINARY_NAME to $INSTALL_DIR/$BINARY_NAME"

# install hamlib_rest_api systemd service
API_SERVICE="/etc/systemd/system/hamlib_rest_api.service"
cp hamlib_rest_api.service "$API_SERVICE"
sed -i "s/{{USER}}/$REAL_USER/g" "$API_SERVICE"
systemctl daemon-reload
systemctl enable hamlib_rest_api.service
systemctl restart hamlib_rest_api.service
echo "Installed, enabled and started hamlib_rest_api systemd service as $REAL_USER"

# install multi-instance rigctld systemd service template
cp rigctld@.service /etc/systemd/system/

# stop and disable all running rigctld services
echo "Stopping all running rigctld services..."
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

bash update_rigctld_services.sh
