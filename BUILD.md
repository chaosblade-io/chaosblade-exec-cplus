# Build Guide

This project supports multi-platform builds without Docker environment and implements a complete version management flow.

## Version Management Flow

```
master/main → Git Tag → CI gets version → Build time injection → Compiled product contains version
```

## Build Requirements

- Go 1.20+
- Make
- Git

## Get Help

View all available build commands:
```bash
make help
```

## Version Information

### Automatic Version Detection
The project automatically gets version information from Git Tags:
```bash
# Display version information
make version

# View help (includes version information)
make help
```

### Version Information Fields
- **Version Number**: Automatically obtained from Git Tag (e.g., 1.4.0)
- **Git Commit**: Current commit hash value
- **Build Time**: Build timestamp
- **Build Type**: release (tagged version) or dev (development version)

## Build Targets

### Default Build (Current Platform)
```bash
make build
```

### Multi-Platform Build

#### Linux
```bash
# Linux AMD64
make linux_amd64

# Linux ARM64
make linux_arm64
```

#### macOS
```bash
# macOS AMD64
make darwin_amd64

# macOS ARM64 (Apple Silicon)
make darwin_arm64
```

#### Windows
```bash
# Windows AMD64
make windows_amd64
```

## Version Control

### Automatic Version Detection (Recommended)
```bash
# Automatically get version from Git Tag
make build
```

### Manual Version Specification
```bash
# Build with specified version
BLADE_VERSION=1.8.0 make build
```

## Build Output

After build completion, output files are located in the `target/chaosblade-$(BLADE_VERSION)/` directory:

- `lib/cplus/chaosblade-exec-cplus` - Executable file (contains version information)
- `yaml/chaosblade-cplus-spec-$(BLADE_VERSION).yaml` - Configuration file with version information
- `script/` - Script file directory

## Version Information Display

### View During Build
```bash
make version
```

### View at Runtime
```bash
# View executable file version
./target/chaosblade-1.4.0/lib/cplus/chaosblade-exec-cplus --version
```

## Other Commands

### Clean
Clean all build products:
```bash
make clean
```

### Test
Run tests:
```bash
make test
```

### Complete Build
Build and test:
```bash
make all
```

### Help
Display all available commands:
```bash
make help
```

## CI/CD Integration

The project includes GitHub Actions CI configuration supporting:

- **Automatic Testing**: Runs when pushing to master/main branch
- **Automatic Building**: Supports multi-platform build matrix
- **Automatic Release**: Automatically creates Release when pushing version tags

### Trigger Conditions
- Push to branch: Run tests and build
- Push tags (v*): Run tests, build, and release

## Version Release Process

```bash
# 1. Ensure code is committed to master branch
git add .
git commit -m "feat: prepare for v1.5.0 release"
git push origin master

# 2. Create version tag
git tag v1.5.0
git push origin v1.5.0

# 3. CI automatically triggers build and release
```

## Related Documentation

- [Detailed Version Management Guide](VERSION_MANAGEMENT.md)
- [GitHub Actions CI Configuration](.github/workflows/ci.yml)
- [Version Script](version/version.sh)
