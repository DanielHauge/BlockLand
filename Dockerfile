FROM golang:jessie

# Install any needed dependencies...
# RUN go get ...
RUN go get github.com/gorilla/mux
RUN go get github.com/rs/cors
RUN go get -u gopkg.in/russross/blackfriday.v2
RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/boltdb/bolt


# Copy the server code into the container
COPY BlockChainHQ.go BlockChainHQ.go
COPY BlocksConstructionOffice.go BlocksConstructionOffice.go
COPY Council.go Council.go
COPY CouncilParticipantsCheck.go CouncilParticipantsCheck.go
COPY CouncilPropositionProtocol.go CouncilPropositionProtocol.go
COPY Handlers.go Handlers.go
COPY Logger.go Logger.go
Copy Main.go Main.go
COPY Messages.go Messages.go
COPY MinersGuild.go MinersGuild.go
COPY PeerCommunicationsProtocol.go PeerCommunicationsProtocol.go
COPY PeersControlCenter.go PeersControlCenter.go
COPY Router.go Router.go
COPY Routes.go Routes.go


EXPOSE 8080
EXPOSE 8081
EXPOSE 8082
EXPOSE 8083
EXPOSE 8084

RUN go build
