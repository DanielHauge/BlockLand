FROM golang:jessie

# Install any needed dependencies...
# RUN go get ...
RUN go get github.com/gorilla/mux
RUN go get github.com/rs/cors


# Copy the server code into the container
COPY Main.go Main.go
COPY Simulator.go Simulator.go
COPY SimulatorAction.go SimulatorAction.go


RUN go build
