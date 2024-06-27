#!/bin/bash

# Obter arquivos modificados na PR usando GitHub environment variables
files_changed=$(git diff --name-only "${{ github.event.pull_request.base.sha }}" "${{ github.sha }}")

# Exibir os arquivos modificados
echo "Changed files in the PR:"
echo "$files_changed"