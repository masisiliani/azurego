FROM golang:alpine

# Copy the local package files to the container’s workspace.
ADD . /go/src/azurego

# Install our dependencies  
RUN apk add git

# Install api binary globally within container
RUN go install azurego

# Set binary as entrypoint
ENTRYPOINT /go/bin/azurego

# Expose default port (5000)
EXPOSE 5000 