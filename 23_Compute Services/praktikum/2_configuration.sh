git clone https://github.com/timmtimm/go_timotius-wirawan
cd o_timotius-wirawan/23_Compute\ Services/praktikum

sudo systemctl docker
docker build -t learn-docker:1.0.0 .
docker run -itd --name myapp -p 8000:8000 learn-docker:1.0.0