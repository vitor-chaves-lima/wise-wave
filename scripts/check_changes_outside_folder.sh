#!/bin/bash

# Check if a directory argument is passed
if [ $# -eq 0 ]; then
  echo "Usage: $0 <directory>"
  exit 1
fi

directory="$1"

# Variable to track if changes outside specified directory are detected
changes_detected=false
changed_dirs=""

# Get the commit range for comparison
GITHUB_BASE_REF=${GITHUB_BASE_REF:-$(git merge-base HEAD HEAD^)}

# Loop through modified files to check directories and list changed directories
echo "Modified directories:"
git diff --name-only "$GITHUB_BASE_REF" HEAD | while read -r file; do
  file_dir=$(dirname "${file}")
  if [[ ! "$changed_dirs" =~ "$file_dir" ]]; then
    changed_dirs+=" $file_dir"
    echo "$file_dir"
  fi
done
echo ""

# Check if changes outside specified directory are detected
for dir in $changed_dirs; do
  if [[ "$dir" != "$directory" ]]; then
    changes_detected=true
  fi
done

# If changes outside specified directory are detected, display error message
if [ "$changes_detected" = true ]; then
  echo "Changes detected outside '${directory}' context in directories:"
  echo "$changed_dirs"
  echo ""
  echo "Please create a separate PR for changes in these directories."
  exit 1
else
  echo "No changes outside '${directory}'."
fi