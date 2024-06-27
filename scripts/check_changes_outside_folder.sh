#!/bin/bash

echo "Base branch: $GITHUB_BASE_REF"
echo "Head branch: $GITHUB_HEAD_REF - Current commit: $GITHUB_SHA"

# Obter arquivos modificados na PR
files_changed=$(git diff --name-only "${GITHUB_BASE_REF}")

# Exibir os arquivos modificados
echo "Changed files in the PR:"
echo "$files_changed"