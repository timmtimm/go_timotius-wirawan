# Praktikum Compute Services

## Prequesities

Already make EC2 and RDS instances

## Initation

Download depedencies needed with running 1_initiate.sh

```
sudo yum update -y
sudo yum install -y golang
sudo yum install -y docker
```

## Configuration

To clone private repository, github authentication only allowed you to input the password column with github token, so the password column not for password but for github token. Github token can generated on https://github.com/settings/tokens. Choose the repository you want to clone and choose the persmission as you need.

```
git clone https://github.com/timmtimm/go_timotius-wirawan
cd go_timotius-wirawan/23_Compute\ Services/praktikum
```

Token (expired on 19 November 2022):
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

In AWS EC 2:
Edit inbound rules
- Add custom TCP Port 1323 0.0.0.0/0
- Add custom TCP 1323 ::/0

Step:
1. Clone repo and change it to the directory
2. Run docker
3. Build images and container

In AWS RDS:
- Add Inbound Rules and Outbound Rules: All Traffic with source Anywhere IPv4 and Ipv6

And then run 2_configuration.sh
```
sudo systemctl start docker
sudo docker build -t learn-docker:1.0.0 .
sudo docker run -itd --name myapp -p 1323:1323 learn-docker:1.0.0
```