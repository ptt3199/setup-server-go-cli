# Setup Server Go CLI

An interactive command-line tool to setup your server from scratch with essential tools and services. Perfect for quickly configuring new Ubuntu/Debian servers.

## ✨ Features

### 🔧 Interactive CLI
- **Main Menu**: Choose between different installation options
- **Selective Installation**: Pick specific tools to install
- **Progress Feedback**: Real-time installation progress with emojis
- **Return to Menu**: Continue installing after each operation completes

### 📦 Available Tools

#### System Accessories
- **tmux** - Terminal multiplexer for session management
- **git** - Version control system
- **curl** - Command line tool for data transfer
- **htop** - Interactive process viewer
- **ufw** - Uncomplicated Firewall
- **zsh** - Advanced shell

#### Services
- **SSH Server** - Secure remote access (openssh-server)
- **Docker** - Complete Docker installation with:
  - Docker Engine (CE)
  - Docker CLI
  - containerd.io
  - Docker Buildx Plugin
  - Docker Compose Plugin
  - User added to docker group

## 🚀 Quick Start

### Download Pre-built Binary

**Linux (x86_64):**
```bash
wget https://github.com/ptt3199/setup-server-go-cli/releases/latest/download/setup-server-cli_Linux_x86_64.tar.gz
tar -xzf setup-server-cli_Linux_x86_64.tar.gz
cd setup-server-cli_Linux_x86_64
chmod +x setup-server-cli
./setup-server-cli
```

**macOS (Apple Silicon):**
```bash
wget https://github.com/ptt3199/setup-server-go-cli/releases/latest/download/setup-server-cli_Darwin_arm64.tar.gz
tar -xzf setup-server-cli_Darwin_arm64.tar.gz
cd setup-server-cli_Darwin_arm64
chmod +x setup-server-cli
./setup-server-cli
```

### Build from Source

**Development builds:**
```bash
# Build for both Linux and macOS
./build-cli.sh
```

This creates two binaries:
- `setup-server-cli-linux` - For Linux servers (x86_64)
- `setup-server-cli-mac` - For macOS (ARM64/Silicon)

**Release builds with GoReleaser:**
```bash
# Install GoReleaser
go install github.com/goreleaser/goreleaser@latest

# Create snapshot release (no git tag needed)
./build-cli.sh release snapshot

# Create tagged release (requires git tag)
git tag v1.0.0
./build-cli.sh release
```

### Run Interactively

```bash
# On Linux server
./setup-server-cli-linux

# On macOS (development)
./setup-server-cli-mac
```

### Direct Installation (Non-interactive)

```bash
# Install all accessories directly
./setup-server-cli-linux install
```

## 🎯 Usage Examples

### Interactive Mode
```
🚀 Server Setup CLI
Choose an option:
1. Install accessories
2. Install Docker
3. Exit
Enter your choice (1-3): 1

Choose accessories to install:
1. tmux
2. git
3. curl
4. htop
5. ufw
6. zsh
Enter the number of the accessory to install (e.g. 1, 2, 3, etc.) or 'all' to install all accessories: 1 3 5

Installing tmux...
✅ Done: sudo apt-get install -y tmux
Installing curl...
✅ Done: sudo apt-get install -y curl
Installing ufw...
✅ Done: sudo apt-get install -y ufw
Accessories installed successfully

✅ Press Enter to return to main menu...
```

### Multiple Selection Options
- `all` - Install all accessories
- `1 2 3` - Install multiple specific tools
- `1` - Install single tool

## 🛠 Development

### Prerequisites
- Go 1.24.3 or higher
- Linux/Ubuntu target system for deployment
- [GoReleaser](https://goreleaser.com/) (for releases)

### Project Structure
```
setup-server/
├── main.go             # Entry point
├── go.mod              # Go module definition
├── .goreleaser.yaml    # GoReleaser configuration
├── build-cli.sh        # Build script (dev + release)
├── .env                # Environment variables (GITHUB_TOKEN)
└── cmd/
    ├── root.go         # Main CLI and interactive menu
    ├── accessories.go  # System tools installation
    └── docker.go       # Docker installation
```

### Building from Source

**Development builds:**
```bash
# Install dependencies
go mod tidy

# Build for current platform
go build -o setup-server-cli .

# Cross-compile for Linux
GOOS=linux GOARCH=amd64 go build -o setup-server-cli-linux .

# Use build script (recommended)
./build-cli.sh
```

**Release builds:**
```bash
# Setup GitHub token in .env file
echo "GITHUB_TOKEN=your_github_token" >> .env

# Install GoReleaser
go install github.com/goreleaser/goreleaser@latest

# Test release build (no publishing)
./build-cli.sh release snapshot

# Create actual release (requires git tag)
git tag v1.0.0
git push origin v1.0.0
./build-cli.sh release
```

## 📦 Release Management

### GoReleaser Features
- **Multi-platform builds**: Linux (x86_64, ARM64), macOS (x86_64, ARM64)
- **GitHub Releases**: Automatic release creation with changelog
- **Archives**: Properly named tar.gz files with checksums
- **Package formats**: DEB and RPM packages for Linux
- **Homebrew**: Automatic formula creation (future)

### Release Process
1. **Development**: Use `./build-cli.sh` for local testing
2. **Snapshot**: Use `./build-cli.sh release snapshot` for testing releases
3. **Tagged Release**: Create git tag and use `./build-cli.sh release`
4. **Distribution**: Users download from GitHub Releases

### Available Downloads
Each release provides:
- `setup-server-cli_Linux_x86_64.tar.gz` - Linux Intel/AMD
- `setup-server-cli_Linux_arm64.tar.gz` - Linux ARM (Raspberry Pi, etc.)
- `setup-server-cli_Darwin_x86_64.tar.gz` - macOS Intel
- `setup-server-cli_Darwin_arm64.tar.gz` - macOS Apple Silicon
- `setup-server-cli_1.0.0_linux_amd64.deb` - Debian package
- `setup-server-cli_1.0.0_linux_amd64.rpm` - RedHat package
- `checksums.txt` - SHA256 checksums for verification

## 🎯 Use Cases

- **New Server Setup**: Quickly install essential tools on fresh Ubuntu/Debian servers
- **Development Environment**: Set up consistent tooling across development machines
- **Docker Installation**: Properly install and configure Docker with all plugins
- **System Administration**: Interactive tool for selective software installation

## 📋 Requirements

- Ubuntu/Debian-based Linux distribution
- sudo privileges for package installation
- Internet connection for downloading packages

## 🔮 Future Enhancements

- [ ] Install cloudflared tunnel
- [ ] Zsh configuration and Oh My Zsh setup
- [ ] Nginx installation and basic configuration
- [ ] SSL certificate setup with Let's Encrypt
- [ ] Firewall rule configuration
- [ ] User management utilities

