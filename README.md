hash-identify

A lightweight, cross-platform hash identification CLI written in Go.

"hash-identify" inspects a hash's length and character pattern to suggest possible hash algorithms. It is designed as a small standalone executable, with no external runtime or third-party dependencies.

Features

- Detects many common hash formats and variants.
- Displays possible and least-likely matches.
- Supports interactive input and command-line arguments.
- Builds as a standalone executable.
- Cross-compiles for multiple operating systems and architectures.
- No Python, runtime installation, or dependency setup required.

«Note: Identification is based on hash format heuristics such as length, character set, and prefixes. It does not cryptographically verify which algorithm produced a hash. Multiple algorithms may match the same input.»

Usage

Pass a hash as an argument:

./hash-identify 5d41402abc4b2a76b9719d911017c592

Or start interactive mode:

./hash-identify

Enter hashes at the "HASH:" prompt. Use "Ctrl+C" to exit.

Build from source

Requirements:

- Go version specified in "go.mod"

Clone and build:

git clone https://github.com/OWNER/hash-identify.git
cd hash-identify

go build -trimpath -ldflags="-s -w -buildid=" -o hash-identify .

For a Linux/macOS build with CGO disabled:

CGO_ENABLED=0 go build \
  -trimpath \
  -ldflags="-s -w -buildid=" \
  -o hash-identify .

Downloads

Prebuilt binaries are published on the project's "GitHub Releases" (https://github.com/OWNER/hash-identify/releases) page.

Download the archive matching your operating system and architecture, extract it, and run the executable.

Supported build targets

The automated release workflow discovers supported targets from "go tool dist list". It builds ordinary CLI targets with CGO disabled, excluding Android and iOS, which require platform-specific packaging.

The workflow can also produce WebAssembly builds where supported by Go.

Releases

To publish a release, push a version tag:

git tag v1.0.0
git push origin v1.0.0

GitHub Actions builds the available targets, packages the binaries, and attaches the archives to a GitHub Release.

License

See "LICENSE" (LICENSE) for license information. WIP