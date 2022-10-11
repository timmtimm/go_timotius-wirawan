# Praktikum Docker

Berikut isi Dockerfile

```
FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o dist

EXPOSE 1323

ENTRYPOINT [ "./dist" ]
```

Untuk membuild dan menjalankan image serta container yang ada bisa menjalankan `sh start.sh ${username_docker}` yang berisi

```sh
docker build -t $1/learn-docker:1.0.0 .
docker run -itd --name myapp -p 1323:1323 $1/learn-docker:1.0.0
```

Untuk menyelesaikan dan menghapus image serta container yang ada bisa menjalankan `sh start.sh ${username_docker}` yang berisi

```sh
docker stop myapp
docker rm -f myapp
docker rmi -f $1/learn-docker:1.0.0
```