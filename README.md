# BlockLand
Blockchain project

## Programming:
- Simulator : ?
- P2P Node
- User Interface? (React) from mobile?

## Research Areas:
- IP Discovery

## setup:
- CI/CD

## Design Decisions
##### Connection Policies
- Everyone has to be connected to everyone
>- Health Check every x Min?
>- Disconnect & Reconnect at points?
- Everyone does not need to be connected to everyone
>- Node roles?
>- Message distribution
##### Blockchain Policies
- Blockchain syncing
>- Every x min
>- Ask everyone how they blockchain looks (Keep majority)
- Who has the rights to add
>- Single person can add a block
>- Multiple people needs to acknowlegde the block
- Who should mine
>- Everyone?
>- Person trying to add.
- Queing system for adding blocks?
##### Deployment strategy
- Dockerswarm?
>- How many machines?
>- Replica, or global?
- Docker Compose?
>- How many machines
>- UI+node stack?
##### Simulator
- Language?
>- C#
>- Java
>- Golang?
##### Topic
- What should the system actully do?
>- Transactions?
>- Chat system? (Simulator could use clever bot API haha xD, imagine 4 bots talking together)[Cleverbot API](https://www.cleverbot.com/api/)
>- Other thing?
