#!/bin/sh
# docudex installer — downloads the release archive matching your OS/arch from
# GitHub, verifies its SHA-256 checksum, and installs the `docudex` binary.
#
#   curl -fsSL https://raw.githubusercontent.com/aureliushq/docudex/main/install.sh | sh
#
# Environment overrides:
#   DOCUDEX_VERSION      version/tag to install (e.g. v0.1.0); default: latest
#   DOCUDEX_INSTALL_DIR  directory to install into;            default: /usr/local/bin
#
# The naming contract below mirrors .goreleaser.yaml: release assets are the
# raw binary docudex_<version>_<os>_<arch> (version has NO leading "v", no
# extension on unix), published under the v<version> tag alongside a checksums.txt.

set -eu

REPO="aureliushq/docudex"
BINARY="docudex"

info() { printf '==> %s\n' "$1"; }
warn() { printf 'warning: %s\n' "$1" >&2; }
err() {
	printf 'error: %s\n' "$1" >&2
	exit 1
}

# download URL DEST — fetch URL to DEST using curl or wget, whichever exists.
download() {
	if command -v curl >/dev/null 2>&1; then
		curl -fsSL "$1" -o "$2"
	elif command -v wget >/dev/null 2>&1; then
		wget -qO "$2" "$1"
	else
		err "need curl or wget to download files"
	fi
}

# detect_os — map uname to a GoReleaser GOOS. curl|sh is a Unix idiom, so only
# darwin/linux are supported; Windows users take the archive from the releases page.
detect_os() {
	os=$(uname -s)
	case "$os" in
	Linux) echo linux ;;
	Darwin) echo darwin ;;
	*) err "unsupported OS '$os' — download a build from https://github.com/${REPO}/releases" ;;
	esac
}

# detect_arch — map uname -m to a GoReleaser GOARCH.
detect_arch() {
	arch=$(uname -m)
	case "$arch" in
	x86_64 | amd64) echo amd64 ;;
	arm64 | aarch64) echo arm64 ;;
	*) err "unsupported architecture '$arch' — download a build from https://github.com/${REPO}/releases" ;;
	esac
}

# latest_tag — resolve the newest release tag (e.g. v0.1.0) via the GitHub API.
latest_tag() {
	api="https://api.github.com/repos/${REPO}/releases/latest"
	tmp="$(mktemp)"
	download "$api" "$tmp" || err "could not query the latest release"
	tag=$(sed -n 's/.*"tag_name": *"\([^"]*\)".*/\1/p' "$tmp" | head -n1)
	rm -f "$tmp"
	[ -n "$tag" ] || err "could not determine the latest release tag"
	echo "$tag"
}

# verify_checksum ARCHIVE CHECKSUMS FILENAME — compare ARCHIVE's SHA-256 against
# the entry for FILENAME in the checksums file. Fatal on mismatch.
verify_checksum() {
	archive="$1"
	checksums="$2"
	filename="$3"

	expected=$(awk -v f="$filename" '$2 == f {print $1}' "$checksums")
	[ -n "$expected" ] || err "no checksum listed for ${filename}"

	if command -v sha256sum >/dev/null 2>&1; then
		actual=$(sha256sum "$archive" | awk '{print $1}')
	elif command -v shasum >/dev/null 2>&1; then
		actual=$(shasum -a 256 "$archive" | awk '{print $1}')
	else
		warn "no sha256sum/shasum found — skipping checksum verification"
		return
	fi

	[ "$expected" = "$actual" ] || err "checksum mismatch for ${filename} (expected ${expected}, got ${actual})"
	info "checksum verified"
}

# install_binary SRC DIR — move SRC to DIR/$BINARY, escalating with sudo only if
# DIR is not writable by the current user.
install_binary() {
	src="$1"
	dir="$2"
	dest="${dir}/${BINARY}"

	chmod +x "$src"
	if [ -w "$dir" ] || { [ ! -e "$dir" ] && mkdir -p "$dir" 2>/dev/null; }; then
		mv "$src" "$dest"
	elif command -v sudo >/dev/null 2>&1; then
		info "elevating with sudo to write to ${dir}"
		sudo mkdir -p "$dir"
		sudo mv "$src" "$dest"
	else
		err "cannot write to ${dir}; set DOCUDEX_INSTALL_DIR to a writable path"
	fi
	echo "$dest"
}

main() {
	os=$(detect_os)
	arch=$(detect_arch)

	tag="${DOCUDEX_VERSION:-}"
	[ -n "$tag" ] || tag=$(latest_tag)
	case "$tag" in
	v*) version="${tag#v}" ;;
	*)
		version="$tag"
		tag="v${tag}"
		;;
	esac

	install_dir="${DOCUDEX_INSTALL_DIR:-/usr/local/bin}"
	# Raw binary asset — no archive extension on darwin/linux (see .goreleaser.yaml).
	filename="${BINARY}_${version}_${os}_${arch}"
	base="https://github.com/${REPO}/releases/download/${tag}"

	info "installing ${BINARY} ${tag} (${os}/${arch})"

	tmpdir=$(mktemp -d)
	trap 'rm -rf "$tmpdir"' EXIT INT TERM

	info "downloading ${filename}"
	download "${base}/${filename}" "${tmpdir}/${filename}" ||
		err "download failed — is ${tag} a published release for ${os}/${arch}?"
	download "${base}/checksums.txt" "${tmpdir}/checksums.txt" ||
		err "could not download checksums.txt for ${tag}"

	verify_checksum "${tmpdir}/${filename}" "${tmpdir}/checksums.txt" "$filename"

	# The downloaded asset is the binary itself — no extraction needed.
	dest=$(install_binary "${tmpdir}/${filename}" "$install_dir")
	info "installed ${BINARY} to ${dest}"

	case ":${PATH}:" in
	*":${install_dir}:"*) ;;
	*) warn "${install_dir} is not on your PATH — add it to run '${BINARY}' directly" ;;
	esac
}

main "$@"
