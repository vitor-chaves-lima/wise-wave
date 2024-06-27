#!/bin/bash

# Definir as referências padrão, se não estiverem definidas
GITHUB_BASE_REF=${GITHUB_BASE_REF:-$(git symbolic-ref --short HEAD)}
GITHUB_HEAD_REF=${GITHUB_HEAD_REF:-$(git merge-base HEAD HEAD^)}

# Obter arquivos modificados na PR
files_changed=$(git diff --name-only "${GITHUB_BASE_REF}" "${GITHUB_HEAD_REF}")

# Exibir os arquivos modificados
echo "Changed files in the PR:"
echo "$files_changed"