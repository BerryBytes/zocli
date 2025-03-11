#!/bin/bash

# generate-changelog.sh
# Generates a changelog between the latest and previous Git tags, excluding certain commit types.
# Outputs to dist/CHANGELOG.md and updates CHANGELOG.md in the repo root.

# Get the latest tag
TAG=$(git describe --tags --abbrev=0)
# Get the previous tag, if it exists
PREV_TAG=$(git describe --tags --abbrev=0 --tags "${TAG}^" 2>/dev/null || echo "")
# Get the current date
DATE=$(date +'%Y-%m-%d')
# Ensure dist directory exists
mkdir -p dist

# Generate the changelog section for this release
TEMP_FILE=$(mktemp)
{
  echo "## $TAG ($DATE)"
  echo ""
  if [ -z "$PREV_TAG" ]; then
    git log --pretty=format:"* %s (%h)" | grep -vE "^docs:|^test:"
  else
    git log "${PREV_TAG}..${TAG}" --pretty=format:"* %s (%h)" | grep -vE "^docs:|^test:"
  fi
  echo ""
} > "$TEMP_FILE"

# If CHANGELOG.md exists, prepend the new content
if [ -f CHANGELOG.md ]; then
  cat "$TEMP_FILE" CHANGELOG.md > temp_changelog && mv temp_changelog CHANGELOG.md
else
  mv "$TEMP_FILE" CHANGELOG.md
fi

echo "Updated CHANGELOG.md and wrote to dist/CHANGELOG.md"
