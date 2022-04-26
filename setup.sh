#!/bin/bash -e

cd src
echo "Compiling server"
cd server
go build server.go middleware.go handlers.go

echo "Making JWT key"
cd ../
touch .env
echo "KEY=$(openssl rand -hex 60)" > .env

cd setup-scripts
go run password.go

echo "In case you want to specify PORT number"
echo "  1: Open src/.env"
echo "  2: Write PORT=WXYZ"
echo "else you have to pass it as an environment varible like:"
echo "  PORT=WXYZ ./src/server/server"

echo "Ready to Go"
