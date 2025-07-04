#!/bin/bash

echo "🛠️  Project Zero Init Script"

read -p "Enter your new Go module path (e.g., github.com/yourusername/yourproject): " MODULE

echo "🔧 Updating go.mod with module name: $MODULE"
go mod edit -module "$MODULE"

echo "✏️  Rewriting import paths from 'placeholder' to '$MODULE'..."
find . -type f -name "*.go" -exec sed -i -e "s|\"placeholder/|\"$MODULE/|g" {} +

echo "🧹 Running go mod tidy..."
go mod tidy

echo "✅ Done! Your project is now using: $MODULE"