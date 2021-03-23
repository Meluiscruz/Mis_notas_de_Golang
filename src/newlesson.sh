#!/bin/bash

echo "Ingrese el nombre de la lección:"
read titulo
mkdir ./"$titulo" && cd ./"$titulo" && touch "$titulo.md"
echo "Ingrese el nombre del script (sin la extensión): "
read script
touch "$script.go"
echo "Todo listo, jefe!!"
sleep 2
clear
exit 0