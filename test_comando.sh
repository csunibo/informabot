#!/bin/bash

# Comando per rimuovere il file 'informabot' se esiste
rm -f informabot

# Export del token come variabile di ambiente
export TOKEN="6924907539:AAFNh_3VeTxCQrE6euNEjgEvV30zUZUaeww"

# Build del programma Go
go build

# Esecuzione di './informabot'
./informabot
