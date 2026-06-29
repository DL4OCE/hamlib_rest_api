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
$REAL_USER ALL=(ALL) NOPASSWD: /usr/bin/systemctl start rotctld@*, /usr/bin/systemctl stop rotctld@*
EOF
chmod 0440 "$SUDOERS_FILE"
echo "Wrote sudoers file to $SUDOERS_FILE for user $REAL_USER"

# install hamlib 

ARCH=$(uname -m)
case "$ARCH" in
    x86_64)
        BIN_SUBDIR="amd64"
        ;;
    aarch64)
        BIN_SUBDIR="arm64"
        ;;
    *)
        BIN_SUBDIR="unknown"
        ;;
esac

echo "Detected system architecture: $ARCH"

LOCAL_BIN_DIR="./build/binaries/linux"
TARGET_DIR="/usr/bin" #/usr/local/bin

echo "Attempting to install Hamlib binaries for $BIN_SUBDIR..."
systemctl stop rigctld@*
systemctl stop rotctld@*

if [ "$BIN_SUBDIR" != "unknown" ] && [ -f "${LOCAL_BIN_DIR}/${BIN_SUBDIR}/rigctld" ] && [ -f "${LOCAL_BIN_DIR}/${BIN_SUBDIR}/rotctld" ]; then
    echo "[+] Found local binaries for $BIN_SUBDIR."
    echo "Installing pre-[x]-compiled Hamlib binaries to ${TARGET_DIR}..."
    
    # Ensure the target directory exists
    sudo mkdir -p "$TARGET_DIR"
    
    sudo cp "${LOCAL_BIN_DIR}/${BIN_SUBDIR}/rigctld" "${LOCAL_BIN_DIR}/${BIN_SUBDIR}/rotctld" "${TARGET_DIR}/"
    # sudo cp "${LOCAL_BIN_DIR}/${BIN_SUBDIR}/rigctld" "${TARGET_DIR}/rigctld"
    # sudo cp "${LOCAL_BIN_DIR}/${BIN_SUBDIR}/rotctld" "${TARGET_DIR}/rotctld"
    sudo chmod +x "${TARGET_DIR}/rigctld" "${TARGET_DIR}/rotctld"
    
    echo "✔ Custom binaries installed successfully!"
    "${TARGET_DIR}/rigctld" --version
    "${TARGET_DIR}/rotctld" --version
else
    echo "[-] No matching local binaries found (or architecture not supported)."
    echo "Falling back to distribution package..."
    echo "Searching for a suitable distribution package..."
    
    if [ -x "$(command -v apt-get)" ]; then
        echo "Debian/Ubuntu-based system detected."
        sudo apt-get update && sudo apt-get install -y libhamlib-utils jq curl tar
    elif [ -x "$(command -v pacman)" ]; then
        echo "Arch-based system detected."
        sudo pacman -Syu --noconfirm hamlib jq curl tar
    elif [ -x "$(command -v dnf)" ]; then
        echo "Fedora/RHEL-based system detected."
        sudo dnf install -y hamlib jq curl tar
    elif [ -x "$(command -v zypper)" ]; then
        echo "openSUSE-based system detected."
        sudo zypper install -y hamlib jq curl tar
    else
        echo "Error: No supported package manager for your distribution found. Please install Hamlib manually. Open a ticket"
        exit 65
    fi
    
    echo "✔ Distribution package installed successfully!"
    rigctld --version
    rotctld --version
fi

# install hamlib_rest_api binary from GitHub releases

REPO="DL4OCE/hamlib_rest_api"
BINARY_NAME="hamlib_rest_api"
INSTALL_DIR="/usr/local/bin"
ARCH=$(uname -m)
case "$ARCH" in
    x86_64)  GOARCH="amd64" ;;
    aarch64) GOARCH="arm64" ;;
    armv7l)  GOARCH="arm" ;; # Fallback for older 32-bit Pis if compiled
    *) echo "Error: unknown architecture $ARCH"; exit 1 ;;
esac
echo "Detected architecture: $ARCH, using GOARCH=$GOARCH"
API_URL="https://api.github.com/repos/$REPO/releases/latest"
DOWNLOAD_URL=$(curl -s "$API_URL" | jq -r ".assets[] | select(.name | contains(\"linux\") and contains(\"$GOARCH\") and endswith(\"tar.gz\")) | .browser_download_url")

if [ -z "$DOWNLOAD_URL" ] || [ "$DOWNLOAD_URL" = "null" ]; then
    echo "Error: Could not find a matching Linux-$GOARCH release asset on GitHub!"
    exit 1
fi
echo "Downloading latest binary release from $DOWNLOAD_URL ..."
curl -sL "$DOWNLOAD_URL" | tar -xz -C "$INSTALL_DIR" "$BINARY_NAME"
chmod 755 "$INSTALL_DIR/$BINARY_NAME"
echo "Installed $BINARY_NAME to $INSTALL_DIR/$BINARY_NAME"

# install hamlib_rest_api systemd service
API_SERVICE="/etc/systemd/system/hamlib_rest_api.service"
cp services/hamlib_rest_api.service "$API_SERVICE"
sed -i "s/{{USER}}/$REAL_USER/g" "$API_SERVICE"
systemctl daemon-reload
systemctl enable hamlib_rest_api.service
systemctl restart hamlib_rest_api.service
echo "Installed, enabled and started hamlib_rest_api systemd service as $REAL_USER"

# install multi-instance rigctld systemd service template
cp services/rigctld@.service services/rotctld@.service /etc/systemd/system/

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

# stop and disable all running rotctld services
echo "Stopping all running rotctld services..."
systemctl list-units "rotctld@*" --plain --no-legend
ACTIVE_SERVICES=$(systemctl list-units --all --plain --no-legend "rotctld@*" | awk '{print $1}')
for service in $ACTIVE_SERVICES; do
    if [ -n "$service" ]; then
        echo "Stopping and deactivating $service ..."
        systemctl stop "$service" || true
        systemctl disable "$service" || true
    fi
done

systemctl daemon-reload

bash update_hamlib_services.sh
