#!/bin/bash

# Verifica se foi passado um diretório como argumento
if [ $# -eq 0 ]; then
  echo "Usage: $0 <directory>"
  exit 1
fi

directory="$1"

# Verificar se houve alterações fora do diretório especificado
if git diff --name-only "${GITHUB_BASE_REF}" "${GITHUB_HEAD_REF}" | grep -v "^${directory}/" | grep -q '^'; then
  echo "Changes outside '${directory}' context detected. Please always work with one context per PR"
  exit 1
else
  echo "No changes outside '${directory}'."
fi