docker build -t $1/learn-docker:1.0.0 .
docker run -itd --name myapp -p 1323:1323 $1/learn-docker:1.0.0