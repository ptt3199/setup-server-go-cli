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

### Build the Binary

```bash
# Build for both Linux and macOS
./build-cli.sh
```

This creates two binaries:
- `setup-server-cli-linux` - For Linux servers (x86_64)
- `setup-server-cli-mac` - For macOS (ARM64/Silicon)

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

### Project Structure
```
setup-server/
├── main.go           # Entry point
├── go.mod           # Go module definition
├── build-cli.sh     # Cross-compilation build script
└── cmd/
    ├── root.go      # Main CLI and interactive menu
    ├── accessories.go # System tools installation
    └── docker.go    # Docker installation
```

### Building from Source

```bash
# Install dependencies
go mod tidy

# Build for current platform
go build -o setup-server-cli .

# Cross-compile for Linux
GOOS=linux GOARCH=amd64 go build -o setup-server-cli-linux .
```

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

