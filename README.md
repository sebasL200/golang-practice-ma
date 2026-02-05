Introducción al Proyecto

Este proyecto consiste en una API REST básica para un CRUD (Crear, Leer, Actualizar, Eliminar) de libros. Está desarrollado en Go, sin depender de ningún framework externo, y utiliza una base de datos SQLite de forma local para su persistencia.
Guía Rápida para Configurar un Proyecto Go
Este tutorial explica los pasos esenciales para iniciar y ejecutar un proyecto básico en Go.
1. Instalación de Go:
Dirígete a la página oficial de Go (https://go.dev/doc/install) e instala la versión más reciente del instalador compatible con tu sistema operativo.
Verificación: Confirma la instalación abriendo una terminal y ejecutando el comando go version.
2. Configuración del Proyecto:
Crea la carpeta de tu proyecto.
Dentro de la carpeta, crea tu archivo principal con la extensión .go (ej. main.go).
3. Inicialización de Módulos (Gestión de Dependencias):
Ejecuta el comando go mod init <nombre del proyecto> para inicializar el módulo del proyecto.
Opcionalmente, usa go mod tidy para limpiar y gestionar las dependencias importadas, especialmente si trabajas con múltiples archivos.
4. Recomendación de Editor:
Si utilizas Visual Studio Code, se recomienda instalar la extensión oficial de Go llamada "Go".
5. Uso de Librerías (Ejemplo con SQLite):
Para integrar paquetes externos, como la librería de SQLite, utiliza el comando go get modernc.org/sqlite.
6. Ejecución del Proyecto:
Para ejecutar tu código Go (por ejemplo, el archivo main.go), utiliza el comando:
go run main.go
