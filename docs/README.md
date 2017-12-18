# Blockland

## Setup
#### Pre-requisite
Before you are able to setup a demo of blockland, you will need to: 
- Setup a linux machine and have access to root user (Be root user when running one-liner). with atleast 512 mb of ram.
- Machine has to have git installed, if not use this command: ```sudo apt-get install git-all``` (Most debian distros have git installed for the box)

#### One-line setup
To install and setup blockland type in this command in a bash terminal **(Read Pre-requisite)**
```
wget -O - https://raw.githubusercontent.com/DanielHauge/BlockLand/master/SetupBlockLand.sh | bash
```

#### Runs that are tested and works:
- We have tested with a digital ocean droplet of 512mb size, which works.
- We have tested on a ubuntu 16.04 vagrant image.

```ruby
Vagrant.configure("2") do |config|

  config.vm.box = "bento/ubuntu-16.04"
  config.vm.network "forwarded_port", guest: 8888, host: 8888, host_ip: "127.0.0.1"
  config.vm.network "private_network", ip: "192.168.33.11"
  
  config.vm.define "Linux", primary: true do |dockermachine|
	dockermachine.vm.hostname = "Linux"
	dockermachine.vm.provision "shell", inline: <<-SHELL
		sudo echo "Ready at 192.168.33.11"
	SHELL
end
end
```
- We have also tested with a debian vagrant image hat works:

```ruby
Vagrant.configure("2") do |config|

  config.vm.box = "debian/jessie64"
  config.vm.network "forwarded_port", guest: 8888, host: 8888, host_ip: "127.0.0.1"
  config.vm.network "private_network", ip: "192.168.33.13"
  config.vm.synced_folder ".", "/home/vagrant/sync", disabled: true
  
  config.vm.define "Linux", primary: true do |dockermachine|
	dockermachine.vm.hostname = "Linux"
	dockermachine.vm.provision "shell", inline: <<-SHELL
		
		sudo echo "Ready at 192.168.33.13"

		
		
	SHELL
end

end
```

## Documentation

#### Test run video
Click me!


[![Demo Test run](https://media.giphy.com/media/xUOxf6LTn2NMizxUUE/giphy.gif)](https://www.youtube.com/watch?v=PJya4nOu0hg&feature=youtu.be)


#### The BlockChain Program - a brief description

We have written the program in Go. We use the following technologies to set up the program, Go lang, Docker, Docker Compose and a Linux machine.
In this documentation we will only link to [PeerToPeer](https://github.com/DanielHauge/BlockLand/tree/master/PeerToPeer) software however the [PeerToPeerSlow](https://github.com/DanielHauge/BlockLand/tree/master/PeerToPeerSlow) works in the same way.

#### Architecture

Our BlockChain program runs on a network of distributed nodes. 
It uses a peer to peer architecture consisting of peers that are equal. A Peer is also referred as a node in the program and documentation.
Each peer runs on a docker container. The Peers communicate via TCP as a network protocol and JSON is used as the data exchange format.
Each peer runs the BlockChain software and has a API endpoint. The [API endpoint](https://github.com/DanielHauge/BlockLand/wiki/API-Endpoint-calls) is used to trigger actions and BlockChain transactions.
 
[![https://raw.githubusercontent.com/DanielHauge/BlockLand/master/Documentation/UML/Main.png](https://raw.githubusercontent.com/DanielHauge/BlockLand/master/Documentation/UML/Main.png)](https://raw.githubusercontent.com/DanielHauge/BlockLand/master/Documentation/UML/Main.png)


#### BlockChain - Implemented actions

- Connect to Peer to Peer network. 
 The peers (nodes) are instantiated with a [docker-compose.yml](https://github.com/DanielHauge/BlockLand/blob/master/docker-compose.yml).
It uses the Docker image build with the [dockerfile](https://github.com/DanielHauge/BlockLand/blob/master/PeerToPeer/Dockerfile). 
The arguments for the docker image can be found in the docker-compose file under ```entrypoint: ./go 127.0.0.1 Daniel :8080```.   
 
- Disconnect to Peer to Peer network. - This functionality is implemented but not used.
 
#### BlockChain Transactions

[![https://raw.githubusercontent.com/DanielHauge/BlockLand/master/Documentation/UML/Usecasediagram1.png](https://raw.githubusercontent.com/DanielHauge/BlockLand/master/Documentation/UML/Usecasediagram1.png)](https://raw.githubusercontent.com/DanielHauge/BlockLand/master/Documentation/UML/Usecasediagram1.png)

#### Mining of Blocks 

We are using the Hash algorithm of Char56.
For the proof of work we are using nonce that give six zeros in the beginning of the hash.


#### Consensus algorithm of your choice, etc.

[![https://raw.githubusercontent.com/DanielHauge/BlockLand/master/Documentation/NEW_BlocklandCouncilMeetingFlow.png](https://raw.githubusercontent.com/DanielHauge/BlockLand/master/Documentation/NEW_BlocklandCouncilMeetingFlow.png)


Flaw: If something occurs at the exact same time, the business logic becomes very heavy, they will need to agree upon who came first and stuff.

Link to [Wiki-page](https://github.com/DanielHauge/BlockLand/wiki).

