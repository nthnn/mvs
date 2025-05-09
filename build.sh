#!/usr/bin/env bash
set -euo pipefail

GO_VERSION=$(go version | cut -d' ' -f3-)
VERSION_NUM="0.0.1"
VERSION="${VERSION:-v${VERSION_NUM}-alpha}"
OUT_DIR="dist"
PKG_DIR="pkg"
DEB_NAME="mvs_${VERSION}_amd64.deb"

LDFLAGS=(
    "-w" "-s"
    "-X main.ActualVersion=${VERSION}"
    "-extldflags=-static"
)

printf "\e[32m→\e[0m Go version: %s\n" "${GO_VERSION}"
printf "\e[32m→\e[0m Building version: %s\n" "${VERSION}"

rm -rf "${OUT_DIR}" "${PKG_DIR}"
mkdir -p "${OUT_DIR}"

go build -ldflags="${LDFLAGS[*]}" -o "${OUT_DIR}/mvs" .
printf "\e[32m→\e[0m Built mvs binary.\n"

mkdir -p "${PKG_DIR}/usr/local/bin"
cp "${OUT_DIR}/mvs" "${PKG_DIR}/usr/local/bin/"

mkdir -p "${PKG_DIR}/DEBIAN"
cat > "${PKG_DIR}/DEBIAN/control" <<EOF
Package: mvs
Version: ${VERSION_NUM}
Section: vcs
Priority: optional
Architecture: amd64
Maintainer: Nathanne Isip <nathanneisip@gmail.com>
Description: Minimal Versioning System - A lightweight versioning system
 MVS provides init, add, remove, commit, amend, log, branch, checkout,
 status, and tree commands, with msgpack metadata, Ed25519 signatures,
 gzip compression, and YAML global configuration.
EOF

cat > "${PKG_DIR}/DEBIAN/postinst" <<'EOF'
#!/usr/bin/env bash
set -e

USER_NAME="${SUDO_USER:-$(logname)}"
USER_HOME=$(getent passwd "$USER_NAME" | cut -d: -f6)
MVS_CONF_DIR="${USER_HOME}/.local/mvs"
MVS_CONF_FILE="${MVS_CONF_DIR}/globals.yaml"

mkdir -p "${MVS_CONF_DIR}"
chown "$USER_NAME":"$USER_NAME" "${MVS_CONF_DIR}"

if [ ! -f "${MVS_CONF_FILE}" ]; then
    RAND_PUB=$(tr -dc 'A-Za-z0-9' < /dev/urandom | head -c 12)
    RAND_PRIV=$(tr -dc 'A-Za-z0-9' < /dev/urandom | head -c 12)

    cat > "${MVS_CONF_FILE}" <<YAML
name: ${USER_NAME}
email: none@none.com
public_key: ${RAND_PUB}
private_key: ${RAND_PRIV}
branch: main
YAML

    chown "$USER_NAME":"$USER_NAME" "${MVS_CONF_FILE}"
    echo "→ A new MVS global config was created at ${MVS_CONF_FILE}." >&2
    echo "   Please edit it now to set your name, email, and keys." >&2
fi

exit 0
EOF
chmod 0755 "${PKG_DIR}/DEBIAN/postinst"

cat > "${PKG_DIR}/DEBIAN/prerm" <<'EOF'
#!/usr/bin/env bash
set -e

case "$1" in
    remove|purge)
        rm -f /usr/local/bin/mvs
        ;;
esac

exit 0
EOF
chmod 0755 "${PKG_DIR}/DEBIAN/prerm"

cat > "${PKG_DIR}/DEBIAN/postrm" <<'EOF'
#!/usr/bin/env bash
set -e

if [ "$1" = "purge" ]; then
    USER_NAME="${SUDO_USER:-$(logname)}"
    USER_HOME=$(getent passwd "$USER_NAME" | cut -d: -f6)
    rm -rf "${USER_HOME}/.local/mvs"
fi

exit 0
EOF
chmod 0755 "${PKG_DIR}/DEBIAN/postrm"

dpkg-deb --build "${PKG_DIR}" "${OUT_DIR}/${DEB_NAME}" > /dev/null
printf "\e[32m→\e[0m Debian package created: %s\n" "${DEB_NAME}"

rm -rf "${OUT_DIR}/mvs" "${PKG_DIR}"

printf "\e[32m→\e[0m Build and packaging complete!\n"
