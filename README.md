# BlockLand
Blockchain project

## Programming:
- Simulator : ?
- P2P Node
- User Interface? (React) from mobile if we have time.

## Research Areas:
- IP Discovery

## setup:
- CI/CD

## Design Decisions
##### Connection Policies
- Everyone has to be connected to everyone
>- Health Check every 5 min
>- Disconnect & Reconnect is possible - when reconnecting needs to get updated
>- Node role - all nodes are equal
>- Message distribution - TCP wire tap
##### Blockchain Policies
- Blockchain syncing
>- Everyone should be synced real-time 
>- All connected parties should be involved when a new block is added - everyone mines.
- Proof of Work
>- Three ones in the start
- Who has the rights to add - everyone has to agree
>- Single person can add a block
>- Multiple people needs to acknowlegde the block
- Who should mine
>- Everyone
- Queuing system for adding blocks?
##### Deployment strategy
- Individual servers. It would be nice to have Dockerswarm.
>- We will set up between 5 and 10 nodes in the peer to peer network.
>- The noes will act as an API where we could connect a GUI.
>- Replica, or global?
- Docker Compose?
>- How many machines
>- UI+node stack?
##### Simulator
- Language:
>- Golang
##### Topic
- What should the system actually do?
>- Queueing system
Other candidates
>- Transactions?
>- Chat system? (Simulator could use clever bot API haha xD, imagine 4 bots talking together)[Cleverbot API](https://www.cleverbot.com/api/)
>- Other thing?
##### Task
- Simulator
- Peer to peer program
- RAD document
- Sequence diagram
- CI/CD
- One line set up.
- Extra: Mobile app