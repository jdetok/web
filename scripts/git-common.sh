#!/bin/bash

msg=""

while getopts "m:" opt; do
    case $opt in
        m) msg="$OPTARG" ;;
        *) read -p "msg: " msg 
    esac 
done


if [ -z "$msg" ]; then
    read -p "msg: " msg
fi 


git add . && \
git commit -m "$msg" && \
git push