#!/bin/bash

# Build the Go application
echo "Building utl application..."
go build -o utl main.go

# Check if the build was successful
if [ $? -ne 0 ]; then
    echo "Build failed. Exiting."
    exit 1
fi

# Move the executable to /usr/local/bin
echo "Installing utl command..."
sudo mv utl /usr/local/bin/

# Make it executable
sudo chmod +x /usr/local/bin/utl