#!/bin/bash
set -e

if [ "$EUID" -ne 0 ]; then
  echo "Please run as root (sudo ./install.sh)"
  exit 1
fi

ARCH=$(uname -m)
case "$ARCH" in
    x86_64)
        ARCH="amd64" ;;
    aarch64)
        ARCH="arm64" ;;
    armv7l)
        ARCH="arm" ;;    
    *)
        echo "Error: unknown architecture $ARCH"
        exit 1
        ;;
esac
echo "Detected system architecture: $ARCH"

echo "Detecting Linux distribution..."
if [ -f /etc/os-release ]; then
    . /etc/os-release
    DISTRO=$ID
    DISTRO_LIKE=$ID_LIKE
else
    echo "Error: /etc/os-release not found. Cannot determine the distribution."
    exit 1
fi

case "$DISTRO" in
    debian|ubuntu|raspbian|pop|linuxmint|zorin)
        echo "Installing build tools for Debian/Ubuntu-based systems..."
        #sudo apt-get update
        sudo apt-get install -y build-essential autoconf automake libtool pkg-config git libusb-1.0-0-dev
        ;;
        
    fedora|rhel|centos|rocky|almalinux)
        echo "Installing build tools for Fedora/RHEL-based systems..."
        sudo dnf install -y @development-tools autoconf automake libtool pkgconf git libusb1-devel
        ;;
        
    arch|manjaro|endeavouros)
        echo "Installing build tools for Arch Linux..."
        sudo pacman -Sy --needed base-devel autoconf automake libtool pkgconf git libusb
        ;;
        
    opensuse*|sles)
        echo "Installing build tools for openSUSE..."
        sudo zypper install -t pattern devel_basis
        sudo zypper install -y autoconf automake libtool pkg-config git libusb-1_0-devel
        ;;
        
    *)
        if [[ "$DISTRO_LIKE" =~ "debian" ]]; then
            echo "Fallback: using Debian-compatible routine..."
            sudo apt-get update
            sudo apt-get install -y build-essential autoconf automake libtool pkg-config git libusb-1.0-0-dev
        elif [[ "$DISTRO_LIKE" =~ "rhel" || "$DISTRO_LIKE" =~ "fedora" ]]; then
            echo "Fallback: using RedHat-compatible routine..."
            sudo dnf install -y @development-tools autoconf automake libtool pkgconf git libusb1-devel
        else
            echo "Error: Unknown distribution '$DISTRO'. Please install the build tools (gcc, make, autotools, libusb-dev) manually."
            exit 1
        fi
        ;;
esac

echo "Patching sudoers file to allow current user to start/stop rigctld services..."
REAL_USER=${SUDO_USER:-$USER}
SUDOERS_FILE="/etc/sudoers.d/hamlib_rest_api"
rm -f "$SUDOERS_FILE"
cat << EOF > "$SUDOERS_FILE"
# Auto-generated sudoers file for hamlib_rest_api - DO NOT EDIT
$REAL_USER ALL=(ALL) NOPASSWD: /usr/binsystemctl start rigctld@*, /usr/bin/systemctl stop rigctld@*
$REAL_USER ALL=(ALL) NOPASSWD: /usr/bin/systemctl start rotctld@*, /usr/bin/systemctl stop rotctld@*
EOF
chmod 0440 "$SUDOERS_FILE"
echo "Wrote sudoers file to $SUDOERS_FILE for user $REAL_USER"

LOCAL_BIN_DIR="./build/binaries/linux"

systemctl stop "rigctld@*" 2>/dev/null || true
systemctl stop "rotctld@*" 2>/dev/null || true

HAMLIB_VERSION="4.7.2"
CURRENT_VERSION=$(rigctl --version 2>/dev/null | head -n 1 | awk '{print $3}')

if [ "$(printf '%s\n' "$HAMLIB_VERSION" "$CURRENT_VERSION" | sort -V | head -n1)" != "$HAMLIB_VERSION" ]; then
    echo "Need to install Hamlib version ${HAMLIB_VERSION} for ${ARCH} from source..."
    wget https://github.com/Hamlib/Hamlib/releases/download/${HAMLIB_VERSION}/hamlib-${HAMLIB_VERSION}.tar.gz
    tar -xzf hamlib-${HAMLIB_VERSION}.tar.gz
    cd hamlib-${HAMLIB_VERSION}
    ./configure
    make -j$(nproc)
    sudo make install
    sudo ldconfig
    rigctl --version
    rotctl --version
    rm -rf hamlib-${HAMLIB_VERSION} hamlib-${HAMLIB_VERSION}.tar.gz
    hash -r
else
    echo "Hamlib version ${HAMLIB_VERSION} is already installed."
fi

REPO="DL4OCE/hamlib_rest_api"
BINARY_NAME="hamlib_rest_api"
INSTALL_DIR="/usr/local/bin"
echo "Installing hamlib_rest_api for architecture: $ARCH"
API_URL="https://api.github.com/repos/$REPO/releases/latest"
DOWNLOAD_URL=$(curl -s "$API_URL" | jq -r ".assets[] | select(.name | contains(\"linux\") and contains(\"$ARCH\") and endswith(\"tar.gz\")) | .browser_download_url")

if [ -z "$DOWNLOAD_URL" ] || [ "$DOWNLOAD_URL" = "null" ]; then
    echo "Error: Could not find a matching Linux-$ARCH release asset on GitHub!"
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
