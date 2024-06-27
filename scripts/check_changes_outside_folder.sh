#!/bin/bash

# Definir as referências padrão, se não estiverem definidas
GITHUB_BASE_REF=${GITHUB_BASE_REF:-$(git symbolic-ref --short HEAD)}
GITHUB_HEAD_REF=${GITHUB_HEAD_REF:-$(git merge-base HEAD HEAD^)}

echo "Base branch: $GITHUB_BASE_REF"
echo "Head branch: $GITHUB_HEAD_REF"