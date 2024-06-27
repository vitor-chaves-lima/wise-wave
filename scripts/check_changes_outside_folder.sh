#!/bin/bash

# Definir as referências padrão, se não estiverem definidas
GITHUB_BASE_REF=${GITHUB_BASE_REF:-$(git symbolic-ref --short HEAD)}
GITHUB_HEAD_REF=${GITHUB_HEAD_REF:-$(git merge-base HEAD HEAD^)}

# Variável para armazenar os diretórios modificados
changed_dirs=""

# Obtém os arquivos modificados na PR
files_changed=$(git diff --name-only "${GITHUB_BASE_REF}" "${GITHUB_HEAD_REF}")

# Loop pelos arquivos modificados para extrair diretórios únicos na raiz do repositório
for file in $files_changed; do
  dir=$(dirname "$file")
  # Extrai o diretório raiz do repositório
  root_dir=$(echo "$dir" | cut -d'/' -f1)
  # Adiciona o diretório raiz na lista, se ainda não estiver presente
  if [[ ! "$changed_dirs" =~ "$root_dir" ]]; then
    changed_dirs+=" $root_dir"
  fi
done

# Exibe os diretórios modificados na raiz do repositório
echo "Directories in the root with changes:"
echo "$changed_dirs