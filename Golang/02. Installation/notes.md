# Golang Installation Guide

## Prerequisites
Before installing Go, ensure your system meets the following requirements:
- Operating system: Windows, macOS, or Linux
- Sufficient disk space (approximately 1GB)
- Internet connection for downloading

## Installation Steps

### 1. Download Go
- Visit the official Go downloads page: https://golang.org/dl/
- Choose the appropriate version for your operating system
- Download the installer package

### 2. Installation Process

#### For Windows:
1. Run the MSI installer
2. Follow the installation wizard
3. Default installation path: `C:\Go`

#### For macOS:
1. Using Homebrew (recommended):
```bash
brew install go
```

2. Manual Installation:
- Open the downloaded package
- Follow the installation prompts
- Default installation path: `/usr/local/go`

#### For Linux:
```bash
# Extract the archive to /usr/local
sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

### 3. Configure Environment Variables
Add the following to your profile (~/.profile, ~/.bashrc, ~/.zshrc):
```bash
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

### 4. Verify Installation
Open a terminal and run:
```bash
go version
```

## Post-Installation Setup

### 1. Create Workspace
Create your Go workspace directory:
```bash
mkdir -p $HOME/go/{bin,src,pkg}
```

### 2. Test Installation
Create and run a simple Go program:
```go
package main

func main() {
    println("Hello, Go!")
}
```

## Common Issues and Troubleshooting
- Ensure PATH is correctly set
- Verify GOPATH configuration
- Check file permissions
- Restart terminal after installation

## Additional Resources
- Official Documentation: https://golang.org/doc/
- Go Tour: https://tour.golang.org/
- Go Playground: https://play.golang.org/