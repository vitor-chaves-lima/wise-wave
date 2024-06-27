#!/bin/bash

# Definir a referência de base da PR
GITHUB_BASE_REF=${GITHUB_BASE_REF:-$(git symbolic-ref --short HEAD)}

# Obtém o commit base da PR
GITHUB_BASE_SHA=$(git merge-base "${GITHUB_BASE_REF}" HEAD)

# Obter arquivos modificados na PR
files_changed=$(git diff --name-only "${GITHUB_BASE_SHA}" HEAD)

# Exibir os arquivos modificados
echo "Changed files in the PR:"
echo "$files_changed"