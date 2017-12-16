echo "Initializing BlockLand setup!"
echo "Installing Docker! & Docker-compose"
wget -O - https://bit.ly/docker-install | bash
echo "--------------------------------------"
echo "--------------------------------------"
echo "----------Getting BlockLand-----------"
echo "--------------------------------------"
echo "--------------------------------------"
git clone https://github.com/DanielHauge/BlockLand.git
echo "--------------------------------------"
echo "--------------------------------------"
echo "----------Building Images-------------"
echo "--------------------------------------"
echo "--------------------------------------"
sudo docker build -t p2pbl:fast $(pwd)/BlockLand/PeerToPeer/.
sudo docker build -t p2pbl:slow $(pwd)/BlockLand/PeerToPeerSlow/.
sudo docker build -t blsim:latest $(pwd)/BlockLand/Simulator/.
echo "--------------------------------------"
echo "--------------------------------------"
echo "--------Starting up Blockland---------"
echo "--------------------------------------"
echo "--------------------------------------"
sudo docker-compose -f $(pwd)/BlockLand/docker-compose.yml up -d 
echo "--------------------------------------"
echo "--------------------------------------"
echo "--------BlockLand started ------------"
echo "--------------------------------------"
echo "--------------------------------------"
echo "----Grace period: sleeping 30 sec.----"
echo "Please resize the terminal to an adequate big size"
sleep 30s
clear
cat <<"EOF"
 .----------------.  .----------------.  .----------------.  .----------------.  .----------------. 
| .--------------. || .--------------. || .--------------. || .--------------. || .--------------. |
| |   ______     | || |   _____      | || |     ____     | || |     ______   | || |  ___  ____   | |
| |  |_   _ \    | || |  |_   _|     | || |   .'    `.   | || |   .' ___  |  | || | |_  ||_  _|  | |
| |    | |_) |   | || |    | |       | || |  /  .--.  \  | || |  / .'   \_|  | || |   | |_/ /    | |
| |    |  __'.   | || |    | |   _   | || |  | |    | |  | || |  | |         | || |   |  __'.    | |
| |   _| |__) |  | || |   _| |__/ |  | || |  \  `--'  /  | || |  \ `.___.'\  | || |  _| |  \ \_  | |
| |  |_______/   | || |  |________|  | || |   `.____.'   | || |   `._____.'  | || | |____||____| | |
| |              | || |              | || |              | || |              | || |              | |
| '--------------' || '--------------' || '--------------' || '--------------' || '--------------' |
 '----------------'  '----------------'  '----------------'  '----------------'  '----------------' 
 
 .----------------.  .----------------.  .-----------------. .----------------. 
| .--------------. || .--------------. || .--------------. || .--------------. |
| |   _____      | || |      __      | || | ____  _____  | || |  ________    | |
| |  |_   _|     | || |     /  \     | || ||_   \|_   _| | || | |_   ___ `.  | |
| |    | |       | || |    / /\ \    | || |  |   \ | |   | || |   | |   `. \ | |
| |    | |   _   | || |   / ____ \   | || |  | |\ \| |   | || |   | |    | | | |
| |   _| |__/ |  | || | _/ /    \ \_ | || | _| |_\   |_  | || |  _| |___.' / | |
| |  |________|  | || ||____|  |____|| || ||_____|\____| | || | |________.'  | |
| |              | || |              | || |              | || |              | |
| '--------------' || '--------------' || '--------------' || '--------------' |
 '----------------'  '----------------'  '----------------'  '----------------'
EOF


echo -e "By Group L / 8: Emmely (cph-el69) & Kristian (cph-kf96) & Daniel (cph-dh136)"   
echo -e "Node API's are available on 8080, 8081, 8082, 8083, 8084, 8085, 8086"
echo -e "Grafana Monitoring is running on port 3000."
echo -e "However, to import the dashboard do the following steps"
echo -e " - Go to grafana on port 3000 and create a datasource as: Prometheus datasource with source as http://Prom:9090 and click save. "
echo -e " - Copy the content of this file: $(pwd)/BlockLand/dashboard.json "
echo -e " - Import the dashboard by pasting it into Import dashboard with json text"
echo -e " From here you can monitor the network, by going to port 3000 "