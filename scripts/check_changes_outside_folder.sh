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

# Loop through modified files to check directories
git diff --name-only "${GITHUB_BASE_REF}" "${GITHUB_HEAD_REF}" | while read -r file; do
  file_dir=$(dirname "${file}")
  if [[ "$file_dir" != "$directory" && ! "$changed_dirs" =~ "$file_dir" ]]; then
    changes_detected=true
    changed_dirs+=" $file_dir"
  fi
done

# If changes outside specified directory are detected, display error message
if [ "$changes_detected" = true ]; then
  echo "Changes detected outside '${directory}' context in directories:${changed_dirs}"
  echo "Please create a separate PR for changes in these directories."
  exit 1
else
  echo "No changes outside '${directory}'."
fi