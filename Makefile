# Filbot Makefile
# builds the extras stored in this repository.
#

.PHONY: filbot hello-world echo deps

buildHelloWorld: 
	cd hello_world/
	go install
	cd ../

buildEcho: 
	cd echo/
	go install 
	cd ../

filbot: 
	go install 

build: buildHelloWorld buildEcho filbot

all: build

deps: 
	go get -u -v github.com/go-chat-bot/bot/...

run: build
	source environment.sh
	filbot

