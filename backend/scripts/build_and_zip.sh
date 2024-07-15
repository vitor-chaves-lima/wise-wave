#!/bin/bash

# Verifica se o caminho da pasta foi passado como argumento
if [ -z "$1" ]; then
  echo "Erro: Nenhum caminho de pasta fornecido."
  echo "Uso: $0 caminho/para/pasta"
  exit 1
fi

# Define o caminho da pasta
FOLDER_PATH="$1"

# Verifica se a pasta existe
if [ ! -d "$FOLDER_PATH" ];then
  echo "Erro: A pasta '$FOLDER_PATH' não foi encontrada."
  exit 1
fi

# Define o caminho para o main.go
MAIN_GO_PATH="$FOLDER_PATH/main.go"

# Verifica se o arquivo main.go existe
if [ ! -f "$MAIN_GO_PATH" ]; then
  echo "Erro: O arquivo '$MAIN_GO_PATH' não foi encontrado em '$FOLDER_PATH'."
  exit 1
fi

# Define o nome da pasta para o executável e o arquivo zip
FOLDER_NAME=$(basename "$FOLDER_PATH")
EXECUTABLE="bootstrap"
BUILD_DIR="$FOLDER_PATH/build"
ZIP_FILE="$BUILD_DIR/${FOLDER_NAME}.zip"

# Cria o diretório build se não existir
mkdir -p $BUILD_DIR

# Compila o código Go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $EXECUTABLE $MAIN_GO_PATH
if [ $? -ne 0 ]; then
  echo "Erro: A compilação do Go falhou."
  exit 1
fi

# Verifica se o executável foi criado
if [ ! -f "$EXECUTABLE" ]; then
  echo "Erro: O executável '$EXECUTABLE' não foi criado."
  exit 1
fi

# Cria o arquivo zip no diretório build
zip $ZIP_FILE $EXECUTABLE
if [ $? -ne 0 ]; then
  echo "Erro: Falha ao criar o arquivo zip."
  exit 1
fi

# Remove o executável
rm $EXECUTABLE
if [ $? -ne 0 ]; then
  echo "Aviso: Falha ao remover o executável '$EXECUTABLE'."
fi

echo "Compilação e empacotamento concluídos com sucesso. Arquivo gerado: $ZIP_FILE"