sudo systemctl start docker
sudo docker build -t learn-docker:1.0.0 .
sudo docker run -itd --name myapp -p 1323:1323 learn-docker:1.0.0