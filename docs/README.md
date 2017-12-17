# Blockland

## Setup
#### Pre-requisite
Before you are able to setup a demo of blockland, you will need to: 
- Setup a linux machine and have access to root user (Be root user). with atleast 512 mb of ram.
>- If you are using things like vagrant or other tools to manage VM's that have super user permisions but are not inherently root users from the get go. Follow these steps
>>- ```passwd root```
>>- Type in new password for root
>>- ```su - root```
>>- Type in the newly made password to change user to root
- Machine has to have git installed, if not use this command: ```sudo apt-get install git-all```
>- Most debian distributions have git installed allready

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
something
```

## Documentation
