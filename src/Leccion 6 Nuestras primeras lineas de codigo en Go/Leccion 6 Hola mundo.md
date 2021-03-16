# Lección 6: Nuestras primeras lineas en Go

## Módulo 1: Hola mundo en Go

Es preciso notar, que existen dos maneras de correr un programa escrito en Go.

1. Con el comando ```build src/helloWorld.go``` se construye el archivo binario a partir del código fuente. Una vez que está creado el archivo binario, lo compilamos con el comando ```./helloWorld```. Se debe sustituir la parte de ```helloWorld``` por el nombre del archivo que se desea compilar.

2. El comando ```go run src/helloWorld.go``` construye el archivo binario, lo compila y después de la compilación se destruye el archivo binario.
