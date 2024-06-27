#!/bin/bash

echo "Base branch: $GITHUB_BASE_REF"
echo "Head branch: $GITHUB_HEAD_REF"

# Obter arquivos modificados na PR
git diff --name-only ${GITHUB_BASE_REF} ${GITHUB_HEAD_REF}