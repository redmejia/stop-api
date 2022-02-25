#!/bin/bash

clear

echo "Stop Store"

# .env enviroment variables file
DOT_ENV=.env

if [ -e "$DOT_ENV" ]; then
	echo "$DOT_ENV found setting and running server"
	source .env

	# Server app
	go run cmd/*.*go
else
	# crate .env file and edit with the enviroment variables
	echo "$DOT_ENV was not found"
	echo "Creating"
	touch .env
fi

