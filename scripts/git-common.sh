#!/bin/bash

# run git add ., git commit, & git push with the following command: 
# gi -m <commit message goes here>

# add the following line to ~/.vshrc (and source it after saving)
# alias gi=~/scripts/git-common.sh

# msg=""

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
