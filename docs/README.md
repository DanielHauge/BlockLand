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
[![Demo Test run](https://media.giphy.com/media/xUOxf6LTn2NMizxUUE/giphy.gif)](https://www.youtube.com/watch?v=PJya4nOu0hg&feature=youtu.be)
