#!/bin/bash

# Version management script for chaosblade-exec-cplus
# This script extracts version information from git tags and provides it to the build system

set -e

# Get the latest git tag
get_latest_tag() {
    git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0"
}

# Get the current git commit hash
get_commit_hash() {
    git rev-parse --short HEAD 2>/dev/null || echo "unknown"
}

# Get the build time
get_build_time() {
    date -u '+%Y-%m-%dT%H:%M:%SZ'
}

# Get version from git tag (remove 'v' prefix)
get_version() {
    local tag=$(get_latest_tag)
    echo "${tag#v}"
}

# Get full version string
get_full_version() {
    local version=$(get_version)
    local commit=$(get_commit_hash)
    local build_time=$(get_build_time)
    echo "${version}-${commit}-${build_time}"
}

# Check if we're on a tagged commit
is_tagged_commit() {
    local tag=$(get_latest_tag)
    local current_commit=$(git rev-parse HEAD)
    local tag_commit=$(git rev-parse "$tag" 2>/dev/null || echo "")
    
    if [ "$current_commit" = "$tag_commit" ]; then
        echo "true"
    else
        echo "false"
    fi
}

# Get build type
get_build_type() {
    if [ "$(is_tagged_commit)" = "true" ]; then
        echo "release"
    else
        echo "dev"
    fi
}

# Main function
main() {
    case "${1:-version}" in
        "version")
            get_version
            ;;
        "full-version")
            get_full_version
            ;;
        "commit")
            get_commit_hash
            ;;
        "build-time")
            get_build_time
            ;;
        "tag")
            get_latest_tag
            ;;
        "build-type")
            get_build_type
            ;;
        "is-tagged")
            is_tagged_commit
            ;;
        *)
            echo "Usage: $0 {version|full-version|commit|build-time|tag|build-type|is-tagged}"
            exit 1
            ;;
    esac
}

main "$@"
