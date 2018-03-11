FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/vitorgusta/red-venture/api

# Install our dependencies       
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/labstack/echo
RUN go get github.com/labstack/echo/middleware
RUN go get gopkg.in/mgo.v2

# Install api binary globally within container
RUN go install github.com/vitorgusta/red-venture/api

# Set binary as entrypoint
ENTRYPOINT /go/bin/api

# Expose default port (1323)
EXPOSE 1323 