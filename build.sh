#!/bin/bash

DEV_DIR=./dev
BUILD_DIR=./build
MAKEIGNORE=buildignore

API_URL_DEV="http://localhost:3000"
API_URL_PROD="https://test.com"

# Reset (clean build directory)
reset() {
    echo -e "\033[35mCleaning $BUILD_DIR\033[0m"
    rm -rf "$BUILD_DIR"/*
    echo -e "\033[32mCleaning completed.\033[0m"
}

# Replace localhost URLs with production API URL in all build files
replace_url() {
    echo -e "\033[35mReplacing localhost with actual API URL in all files\033[0m"
    find "$BUILD_DIR" -type f -exec sed -i "s|$API_URL_DEV|$API_URL_PROD|g" {} +
    echo -e "\033[32mAPI URL replacement completed.\033[0m"
}

# Minify JS files
minify() {
    echo -e "\033[35mMinifying mobile JS files in $BUILD_DIR\033[0m"

    # Mobile JS minify
    find "./dev/frontend/mobile/src/js" -type f -name "*.js" -exec bash -c '
        SRC_DIR="$1"
        DEST_DIR="$2"
        for file in "${@:3}"; do
            minified_path="$DEST_DIR/${file#$SRC_DIR/}"
            mkdir -p "$(dirname "$minified_path")"
            terser "$file" --compress --mangle -o "$minified_path"
        done
    ' _ "./dev/frontend/mobile/src/js" "./build/frontend/mobile/src/js" {} +

    echo -e "\033[32mMobile JS minify completed.\033[0m"

    echo -e "\033[35mMinifying desktop JS files in $BUILD_DIR\033[0m"

    # Desktop JS minify
    find "./dev/frontend/desktop/src/js" -type f -name "*.js" -exec bash -c '
        SRC_DIR="$1"
        DEST_DIR="$2"
        for file in "${@:3}"; do
            minified_path="$DEST_DIR/${file#$SRC_DIR/}"
            mkdir -p "$(dirname "$minified_path")"
            terser "$file" --compress --mangle -o "$minified_path"
        done
    ' _ "./dev/frontend/desktop/src/js" "./build/frontend/desktop/src/js" {} +

    echo -e "\033[32mDesktop JS minify completed.\033[0m"

    replace_url
}

# Build (reset + copy + minify)
build() {
    reset
    echo -e "\033[35mCopying files from $DEV_DIR to $BUILD_DIR\033[0m"

    mkdir -p "$BUILD_DIR"

    tar --exclude-from="$MAKEIGNORE" -cf - -C "$DEV_DIR" . | tar -xf - -C "$BUILD_DIR"

    echo -e "\033[32mCopy completed.\033[0m"
    minify
}


# Default action
all() {
    build
}

# Run the script based on first argument or default to all
case "$1" in
    reset) reset ;;
    build) build ;;
    minify) minify ;;
    replace_url) replace_url ;;
    all|"") all ;;
    *) echo "Usage: $0 {reset|build|minify|replace_url|all}" ;;
esac
