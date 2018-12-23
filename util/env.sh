#!/bin/bash

jsonize() {
    if [ "$#" -eq 0 ]; then
        echo "Gimme at least a db file"
    elif [ "$#" -eq 1 ]; then
        bsondump --pretty $1
    else
        bsondump --pretty --outFile $2 $1
    fi
}