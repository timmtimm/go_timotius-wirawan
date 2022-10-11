bash
docker build -t learn-docker:1.0.0 .
docker run -itd --name myapp -p 1322:1321 learn-docker:1.0.0