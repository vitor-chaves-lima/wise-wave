#!/bin/bash

echo "Base branch: $GITHUB_BASE_REF"
echo "Head branch: $GITHUB_HEAD_REF"

# Obter arquivos modificados na PR
echo git branch