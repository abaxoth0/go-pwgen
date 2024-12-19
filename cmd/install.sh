#!/bin/bash

HELP="Available installation options:
        \n\t '--local' - install in /usr/local/bin
        \n\t '--user' - install in /usr/bin
        \n\t '--global' - install in /bin"

if [ ! "$#" == 1 ]; then
    echo "Invalid amount of parameters. Only one parameter required - installation option"
    echo -e $HELP
    exit 1
fi

opts=("--local" "--user" "--global")

if [[ ! " ${opts[*]} " =~ " $1 " ]]; then
    echo "Invalid installation option: ${1}"
    echo -e $HELP
    exit 1
fi

go build main.go

if [[ $? -ne 0 ]]; then
    echo "Failed to build pwgen"
    exit 1
fi

chmod +x main

if [ $1 == "--local" ]; then
    sudo mv main /usr/local/bin/pwgen
fi

if [ $1 == "--user" ]; then
    sudo mv main /usr/bin/pwgen
fi

if [ $1 == "--global" ]; then
    sudo mv main /bin/pwgen
fi

if [[ -e "main" ]]; then
    echo "clean up"
    rm main
fi

if [[ ! $? -ne 0 ]]; then
    echo "pwgen installed"
fi

