# airdrop

CLI utility for Apple AirDrop — send files via AirDrop directly from the terminal.

## Features

- Send files via AirDrop from command line
- Support for multiple files
- Pipe support — send files via stdin
- Automatic file type detection
- Cross-architecture support (Apple Silicon & Intel)

## Installation

### Homebrew (Recommended)

```bash
brew tap xvoland/tap
brew install airdrop
```

### Manual

Download the latest release from [GitHub Releases](https://github.com/xvoland/airdrop/releases):

```bash
curl -L -o airdrop.tar.gz https://github.com/xvoland/airdrop/releases/download/v0.3.6/airdrop_darwin_arm64.tar.gz
tar -xzf airdrop.tar.gz
./airdrop file.pdf
```

## Usage

### Send single file

```bash
airdrop file.pdf
```

### Send multiple files

```bash
airdrop file1.jpg file2.png document.pdf
```

### Send via pipe

```bash
cat file.pdf | airdrop
```

### Options

```bash
airdrop --help
```

- `-v` — verbose logging
- `--version` — show version information

## Requirements

- macOS (macOS 10.15 or later)
- Apple AirDrop enabled

## Gatekeeper / Security Notice

If you get a security error when running airdrop after installation, you need to bypass Gatekeeper:

### Option 1: Remove quarantine attribute (recommended)

```bash
xattr -cr $(which airdrop)
```

### Option 2: Allow in System Settings

1. Go to **System Settings → Privacy & Security**
2. Scroll down to **Security**
3. Click **Allow anyway** next to the blocked airdrop

### Option 3: Sign the binary yourself (for developers)

```bash
codesign --force --deep -s - ./airdrop
```

---

## Development

### Build from source

```bash
# Clone the repository
git clone https://github.com/xvoland/airdrop.git
cd airdrop

# Build
make

# Run
./airdrop file.pdf
```

### Build for different architectures

```bash
# Apple Silicon (M1/M2/M3)
make all ARCH=arm64

# Intel
make all ARCH=x86_64
```

## License

MIT License — see [LICENSE](LICENSE)

## Author

Vitalii Tereshchuk  
https://dotoca.net/airdrop
