<p align="right"><img align="center" src="https://raw.githubusercontent.com/xvoland/xvoland/main/images/qr_airdrop.jpg" alt="DOTOCA Ltd." height="50" width="50" /></a>
</p>

# CLI airdrop 📤

CLI utility for Apple AirDrop — send files via AirDrop directly from the terminal.

## What is this? 🤔

**airdrop** is a command-line tool that lets you send files to nearby Apple devices (iPhone, iPad, Mac) using Apple's AirDrop feature — but ***directly from the terminal!***

Instead of dragging files to the AirDrop icon, you can just type:

```bash
airdrop image.jpg
```

```bash
airdrop myfile.pdf
```

## Features ✨

- Send files via AirDrop from command line
- Support for multiple files at once
- Pipe support — send files via stdin
- Automatic file type detection
- Works on Apple Silicon and Intel Macs

## Installation 📥

### Method 1: Homebrew (Easiest!) 🍺

If you have Homebrew installed (most Mac users do):

```bash
brew tap xvoland/tap
brew install airdrop
```

### Method 2: Manual Download

1. Go to [GitHub Releases](https://github.com/xvoland/airdrop/releases)
2. Download the latest `airdrop_darwin_arm64.tar.gz` (for Apple Silicon) or `airdrop_darwin_x86_64.tar.gz` (for Intel)
3. Extract the file
4. Run it!

Or use terminal:

```bash
# For Apple Silicon (M1/M2/M3)
curl -L -o airdrop.tar.gz https://github.com/xvoland/airdrop/releases/download/v0.3.6/airdrop_darwin_arm64.tar.gz
tar -xzf airdrop.tar.gz
./airdrop yourfile.pdf
```

## How to Use 📖

### Basic Example: Send One File

Let's say you have a file called `photo.jpg` and want to send it to your iPhone:

```bash
airdrop photo.jpg
```

That's it! Airdrop will open and let you choose where to send the file.

### Send Multiple Files

You can send several files at once:

```bash
airdrop photo.jpg document.pdf screenshot.png
```

### Send via Pipe (Advanced)

If you want to send content that comes from another program:

```bash
cat myfile.pdf | airdrop
```

This sends the content of `myfile.pdf` through AirDrop.

## Command Options ⚙️

### Help

See all options:
```bash
airdrop --help
```

### Version

Check which version you have:
```bash
airdrop --version
```

### Verbose Mode

Want to see more details about what's happening?
```bash
airdrop -v file.pdf
```

## Troubleshooting 🔧

### "Permission denied" or "Blocked by macOS"

macOS has a security feature called Gatekeeper that blocks apps from unknown developers. If this happens:

**Solution 1: Run this command in terminal**
```bash
xattr -cr $(which airdrop)
```

**Solution 2: Allow manually**
1. Open **System Settings** → **Privacy & Security**
2. Look for the message about blocked airdrop
3. Click **Allow anyway**
4. Try running airdrop again

### "AirDrop failed"

Make sure:
- ✅ AirDrop is turned ON on your Mac (Control Center → AirDrop)
- ✅ AirDrop is turned ON on your receiving device (iPhone/iPad)
- ✅ Devices are close to each other
- ✅ You accepted the incoming request on the receiving device

## Requirements 📋

- macOS 10.15 (Catalina) or later
- AirDrop enabled on your Mac
- Receiving device nearby (iPhone, iPad, or another Mac)

## Building from Source (For Developers) 👨‍💻

Want to build it yourself? Here's how:

```bash
# 1. Clone the project
git clone https://github.com/xvoland/airdrop.git
cd airdrop

# 2. Build it
make

# 3. Run it!
./airdrop file.pdf
```

### Building for Different Macs

```bash
# For Apple Silicon (M1/M2/M3)
make all ARCH=arm64

# For Intel Mac
make all ARCH=x86_64
```

## License 📄

MIT License — see [LICENSE](LICENSE)

## Donation

<p>I’ll keep improving the app no matter what because I love seeing people use it to reach their goals. But if you’d like to support my work, even a $1 donation makes a big difference—it helps cover hosting costs and the time I put into coding. Every little bit helps, and I’d truly appreciate it.</p>
<p>If you enjoy the my work, consider donating today. Thank you so much! 🙌</p>

<p align="center">
  <a href="https://paypal.me/xvoland" target="blank"><img align="center" src="https://raw.githubusercontent.com/xvoland/xvoland/main/images/paypal.png" alt="PayPal" width="250" /></a>
</p>

### Crypto

**BTC (ERC20):** 0x17496b75d241d377334717f8cbc16cc1a5b80396<br />
**USDT (TRC20):** TAAsGXjNoQRJ7ewxSBL2W3DUCoG7h8LCT6


## HOME 👤

🌐 https://dotoca.net/airdrop  

## Other
### ☎️  Connect with me:

<p align="left">
  <a href="https://youtube.com/xvoland" target="blank"><img align="center" src="https://raw.githubusercontent.com/xvoland/xvoland/main/images/youtube.svg" alt="Youtube channel" height="30" width="40" /></a>
  <a href="https://instagram.com/xvoland" target="blank"><img align="center" src="https://raw.githubusercontent.com/xvoland/xvoland/main/images/instagram.svg" alt="xVoLAnD" height="30" width="40" /></a>
  <a href="https://www.linkedin.com/in/vitalij-terescsuk-02b4689/" target="blank"><img align="center" src="https://raw.githubusercontent.com/xvoland/xvoland/main/images/linked-in-alt.svg" alt="xVoLAnD" height="30" width="40" /></a>
  <a href="https://dotoca.net" target="blank"><img align="center" src="https://raw.githubusercontent.com/xvoland/xvoland/main/images/logo-dotoca.svg" alt="DOTOCA Ltd." height="50" width="80" /></a>
</p>

<br />
<br />

📺 Latest YouTube Videos
<!-- BLOG-POST-LIST:START -->
- [🔴 WWDC App&#39;s Wall Screensaver | Apple Store Apps Wall  #live #screensaver4k #relax](https://www.youtube.com/watch?v=tZ3UaYibMso)
- [OptiGrid – Optical Illusion Pattern Generator](https://www.youtube.com/shorts/uJ_X90zbwUQ)
- [Gemini 3.1 Nano Banana 2 - AI Photoshop Plugin jsxNanaBananaPro v0.6.4](https://www.youtube.com/watch?v=GtspFSN7VlI)
- [How To Change The Thumbnail for Shorts using the YouTube app for iOS #youtube #techtips #tutorial](https://www.youtube.com/shorts/_1UBbKtWUT4)
- [🍌 Photoshop Plugin for Gemini 3 Pro Nano Banana #photoshop #photoshoptutorial](https://www.youtube.com/shorts/lIJcYdDSH5M)
<!-- BLOG-POST-LIST:END -->

---

=======
Made with ❤️ for macOS users!
