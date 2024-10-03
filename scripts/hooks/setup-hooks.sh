#!/bin/bash
# scripts/hooks/setup_hooks.sh

echo "Setting up Git hooks..."
# Check if the script is being run from the root of the repository
if [ ! -d ".git" ]; then
    echo "This script must be run from the root of the Git repository."
    exit 1
fi
# Check if the .git/hooks directory exists
if [ ! -d ".git/hooks" ]; then
    echo "This directory does not appear to be a Git repository."
    exit 1
fi

# Copy the specific hooks to the correct directory
cp ./scripts/hooks/pre-commit .git/hooks/
cp ./scripts/hooks/pre-push .git/hooks/

# Make the hooks executable
chmod +x .git/hooks/pre-commit
chmod +x .git/hooks/pre-push

echo "Git hooks successfully configured."
