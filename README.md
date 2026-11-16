hash-identify

A lightweight, cross-platform hash identification CLI written in Go.

https://img.shields.io/badge/Go-1.21%2B-00ADD8?logo=go&logoColor=white
https://img.shields.io/badge/License-Apache%202.0-blue.svg
https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows%20%7C%20BSD%20%7C%20WASM-informational
https://img.shields.io/github/v/release/nanoctf/ctfbox-app-hashidfy?color=success

Identify possible hash algorithms using hash length, character patterns, and format prefixes. Built as a compact standalone executable with zero third-party dependencies.

---

✨ Features

 
🔎 Broad Detection Recognizes numerous common hash algorithms and variants.
📊 Ranked Results Separates possible matches from less likely matches.
💬 Interactive & CLI Modes Enter hashes interactively or pass them as arguments.
📦 Standalone Binary No Python, runtime installation, or dependency management.
🌍 Cross-Platform Builds for multiple operating systems and architectures.
🪶 Lightweight Optimized builds with stripped debugging information.

---

🚀 Usage

Direct input

```bash
./hash-identify 5d41402abc4b2a76b9719d911017c592
```

Interactive mode

```bash
./hash-identify
```

Enter a hash when prompted:

```text
--------------------------------------------------
HASH: 5d41402abc4b2a76b9719d911017c592

Possible Hashs:
  [+] MD5
  [+] NTLM
```

Press <kbd>Ctrl</kbd>+<kbd>C</kbd> to exit.

---

📦 Installation

Prebuilt binaries

Download the archive for your operating system and architecture from the
GitHub Releases page.

Extract the archive and execute the binary.

Build from source

Requirements: Go version specified in go.mod.

```bash
git clone https://github.com/nanoctf/ctfbox-app-hashidfy.git
cd hash-identify

CGO_ENABLED=0 go build \
  -trimpath \
  -ldflags="-s -w -buildid=" \
  -o hash-identify .
```

---

🖥️ Supported platforms

Release builds target multiple operating systems and architectures, including:

Operating System Architectures
Linux AMD64, ARM64, ARM, 386, RISC-V, PPC64, S390X, LoongArch
Windows AMD64, ARM64, 386
macOS AMD64, ARM64
FreeBSD AMD64, ARM64, 386
OpenBSD AMD64, ARM64, 386, PPC64
NetBSD AMD64, ARM64, ARM, 386
DragonFly BSD AMD64
Solaris AMD64
illumos AMD64
WebAssembly WASM

ℹ️ Some targets may not be available with every Go version or build environment.
WebAssembly binaries require a compatible runtime.

---

🏷️ Releases

Releases are automated using GitHub Actions.

To publish a new version:

```bash
git tag v1.0.0
git push origin v1.0.0
```

The release workflow will:

1. Build all supported targets
2. Package successful builds
3. Publish them as GitHub Release assets

---

🔍 How it works

Hash identification is based on structural characteristics:

· Length of the input
· Character set
· Numeric and alphabetic composition
· Known format prefixes and separators

⚠️ Important: This is heuristic identification, not cryptographic verification. Different algorithms can produce hashes with identical formats, so multiple matches are possible.

---

📄 License

Licensed under the Apache License 2.0.

See LICENSE for the full license text.

---

<div align="center">

Built with Go 🐹

</div>