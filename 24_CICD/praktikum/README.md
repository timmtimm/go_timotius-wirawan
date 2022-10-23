# Praktikum CI/CD

Pada praktikum ini ada pada repository https://github.com/timmtimm/practice-ci-cd

Langkah:
1. Mengambil template dari github action tentang Docker
2. Menambah konfigurasi pada settings/Actions/Runners lalu tambahkan self-hosted runner dan jalankan serangkaian command sesuai dengan OS yang digunakan di server
3. Mengganti config sebagai berikut
    ```yml
    name: Docker Image CI

    on:
    push:
        branches: [ "main" ]
    pull_request:
        branches: [ "main" ]

    jobs:

    build:

        runs-on: self-hosted

        steps:
        - uses: actions/checkout@v3
        - name: Stopping old docker container
        run: sudo docker stop /myapp || true
        - name: Removing old docker container
        run: sudo docker rm /myapp || true
        - name: Build docker images
        run: sudo docker build -t learn-docker:1.0.0 .
        - name: Run docker
        run: sudo docker run -itd -p 1323:1323 --name myapp learn-docker:1.0.0
    ```
4. Sebagai tambahan, untuk best practice sebenarnya bisa ditambah testing dan linting tetapi karena pada source code belum ada maka belum bisa ditambahkan
5. untuk mengakses hasil berikut bisa dijalanakan pada [URL berikut](http://13.212.21.39:1323)