# go-loader

_go-loader es una aplicación que permite la carga de archivos masivos independiente de la lógica de negocio que necesite usarlo. Es una solución que puede ser integrada a una arquitectura orientada a eventos. El objetivo es convertir cada linea del archivo en un stream object en forma de mensaje, para que otros sistemas que necesiten la información pueda procesarla acorde a sus necesidades. Aunque inicialmente funciona con Redis, puede evolucionar facilmente para usar otros motores de mensajería de stream como Kafka_


## Pre-requisitos 📋


```
1. Go instalado
2. IDE de tu preferencia para el lenguaje Go
3. Docker, Podman, u otro gestor de containers instalado, o un servidor de Redis.
4. Si tomó la opción de Docker o Podman, puede ejecutar el siguiente comando para levantar un contenedor de Docker:

docker pull redis
docker run -p 6379:6379 redis --protected-mode no

```
## Variables de Entorno

```
  1. LOADER_CONFIG_PATH=config.json  >> Ruta del Archivo de Configuración que se explica en el siguiente item, por defecto hay un archivo en la raiz del proyecto.
  2. LOADER_UPLOAD_FOLDER=/home/user/upload >> Ruta del folder para almacenar los archivos subidos 
  3. PORT=8082 >> Puerto en el que corre la aplicación
  4. REDIS_HOST=localhost >> Server host de Redis
  5. REDIS_PORT=6379 >> Server port de Redis

NOTA: Esta versión solo se conecta a Redis sin autenticación.

LOADER_CONFIG_PATH=config.json LOADER_UPLOAD_FOLDER=/home/user/upload REDIS_HOST=localhost REDIS_PORT=6379 go run main.go
  
```
## Archivo de configuración: config.json

_El archivo de configuración permite configurar el formato del archivo, los tamaños de lectura y escritura del archivo, y la configuración, nombres y orden de los campos._

Por defecto tiene las siguientes opciones:

```
{
  "loaderName": "LoaderName", //Un nombre para nuestro cargador
  "streamName": "saveItems", //Un nombre para el stream donde se publicarán los eventos
  "delimiter": ",", //Delimitador de campos (aplica para archivos en formato CSV)
  "fetchSize": 100, //Tamaño de registros de lectura (Permite controlar la cantidad de registros que cargamos a memoria)
  "chunkSize": 20, //Tamaño de registros de escritura (Permite controlar la cantidad de registros que queremos escribir)
  "skipFirstLine": true, //Para archivos con headers, permite omitir la primera linea
  "formatFile": "CSV", //Formato que soportará la carga del archivo (soporta CSV y JSON)
  "fields": [ //Permite configurar los nombres y las posiciones de los campos
    {
      "fieldName": "site", //Nombre con el que se almacenará en el mensaje
      "index": 1 //Posición de donde tomará el valor en el archivo, aplica solo para CSV
    },
    {
      "fieldName": "id",
      "index": 2
    }
  ]
}
```
## Arquitectura

![alt text](https://github.com/enavarrom/go-loader/blob/main/Arquitectura.drawio.png?raw=tr)

La arquitectura se basa en la arquitectura de Spring Batch, pero lo más basico de él. Las clases mas importantes son:

Loader: Ejecuta el proceso teniendo encuenta las interfaces ItemReader, ItemProcessor, ItemWriter.
ItemReader: Es una interfaz que permite obtener información registro a registro de una fuente de datos.
ItemProcessor: Se encarga de traducir el registro de modo que pueda ser interpretado por el proceso de escritura.
ItemWriter: Se encarga de poner la información leida en un destino de datos.

Para nuestro caso se ha implementado Csv y Json Item Readers, que permite sacar la información de estos dos origenes. Y se escribe la información en Redis a traves de la implementación de Redis Stream Item Reader.

## Ejecución de la aplicación

Se puede ejecutar la aplicación haciendo el build y luego run del archivo generado. O solo descargando el proyecto y correr el comando:

```
LOADER_CONFIG_PATH=config.json LOADER_UPLOAD_FOLDER=/home/user/upload REDIS_HOST=localhost REDIS_PORT=6379 go run main.go
```
Recuerde modificar los valores de las variables por los de su entorno de trabajo.


Una vez arriba puede probar con alguna herramienta como postman a la siguiente URL: POST localhost:8080/upload y debe pasar el archivo en el cuerpo del mensaje.

curl --location 'localhost:8080/upload' --form 'file=@"/home/technical_challenge_data.csv"'




