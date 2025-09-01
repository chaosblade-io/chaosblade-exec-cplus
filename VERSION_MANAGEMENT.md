# Version Management Guide

This document describes the version management process for the ChaosBlade C++ Executor project, from Git Tag to CI build to final product with complete version information.

## Version Management Flow

```
master/main → Git Tag → CI gets version → Build time injection → Compiled product contains version
```

### 1. Git Tag Management

#### Create Version Tags
```bash
# Create new version tag
git tag v1.5.0

# Push tag to remote repository
git push origin v1.5.0
```

#### Version Naming Convention
- Use semantic versioning: `vMajor.Minor.Patch`
- Examples: `v1.5.0`, `v2.0.0`, `v1.5.1`

### 2. Automatic Version Detection

The project uses the `version/version.sh` script to automatically detect version information:

```bash
# Get version number (remove 'v' prefix)
./version/version.sh version        # Output: 1.5.0

# Get full version information
./version/version.sh full-version   # Output: 1.5.0-a05596d-2025-09-01T03:45:45Z

# Get Git commit hash
./version/version.sh commit         # Output: a05596d

# Get build time
./version/version.sh build-time     # Output: 2025-09-01T03:45:45Z

# Get build type
./version/version.sh build-type     # Output: release or dev

# Check if it's a tagged version
./version/version.sh is-tagged      # Output: true or false
```

### 3. Build Time Version Injection

#### Automatic Version Detection
```bash
# Automatically get version from Git Tag
make build

# Build with specified version
BLADE_VERSION=1.8.0 make build
```

#### Version Information Fields
The following version information is automatically injected during build:
- `Version`: Version number (e.g., 1.5.0)
- `GitCommit`: Git commit hash
- `BuildTime`: Build timestamp
- `BuildType`: Build type (release/dev)

### 4. Version Information Display

#### View Version During Build
```bash
# Display version information
make version

# Display help information (includes version)
make help
```

#### View Version at Runtime
```bash
# View executable file version
./target/chaosblade-1.5.0/lib/cplus/chaosblade-exec-cplus --version
```

### 5. CI/CD Integration

#### GitHub Actions Trigger Conditions
- Push to `master` or `main` branch: Run tests and build
- Push tags (`v*`): Run tests, build, and release

#### Build Matrix
Supports multi-platform builds:
- Linux: AMD64, ARM64
- macOS: AMD64, ARM64  
- Windows: AMD64

#### Automatic Release
When pushing version tags, CI automatically:
1. Runs tests
2. Builds all platform versions
3. Creates GitHub Release
4. Uploads build artifacts

### 6. Version Information in Build Products

#### Executable Files
- View embedded version information via `strings` command
- Supports `--version` parameter to display version

#### Configuration Files
- YAML filename includes version number: `chaosblade-cplus-spec-1.5.0.yaml`
- File content includes version information

#### Directory Structure
```
target/
└── chaosblade-1.5.0/
    ├── lib/
    │   └── cplus/
    │       ├── chaosblade-exec-cplus
    │       └── script/
    └── yaml/
        └── chaosblade-cplus-spec-1.5.0.yaml
```

## Best Practices

### 1. Version Release Process
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

### 2. Development Version Management
- Development branches use `dev` build type
- Release versions use `release` build type
- Version numbers are automatically obtained from Git Tags

### 3. Version Rollback
```bash
# If version rollback is needed
git tag -d v1.5.0
git push origin :refs/tags/v1.5.0

# Recreate correct tag
git tag v1.5.0
git push origin v1.5.0
```

## Troubleshooting

### 1. Version Detection Failure
- Ensure Git repository has tags
- Check `version/version.sh` script permissions
- Verify Git command availability

### 2. Build Failure
- Check if version variables are correctly injected
- Verify version references in Makefile
- Confirm CI environment configuration

### 3. Incomplete Version Information
- Check `-ldflags` parameter settings
- Verify version variable declarations in Go code
- Confirm build script execution

## Related Files

- `version/version.sh`: Version detection script
- `Makefile`: Build configuration and version injection
- `.github/workflows/ci.yml`: CI/CD configuration
- `main.go`: Version information display
- `build/spec.go`: Configuration file generation
