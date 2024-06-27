#!/bin/bash

echo "Base branch: $GITHUB_BASE_REF"
echo "Head branch: $GITHUB_HEAD_REF"

echo "Switching to $GITHUB_HEAD_REF branch:"
git checkout $GITHUB_HEAD_REF

# Obter arquivos modificados na PR
files_changed="$(git diff --name-only ${GITHUB_BASE_REF} ${GITHUB_HEAD_REF})"

# Exibir os arquivos modificados
echo "Changed files in the PR:"
echo "$files_changed"