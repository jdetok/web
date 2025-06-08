#!/bin/bash

# run git add ., git commit, & git push with the following command: 
# gi -m <commit message goes here>

# add the following line to ~/.vshrc (and source it after saving)
# alias gi=~/scripts/git-common.sh

# accept -m as arg & assign to msg for commit
# if a different arg is passed, prompt for the msg 
while getopts "m:" opt; do
    case $opt in
        m) msg="$OPTARG" ;;
        *) read -p "invalid opt, enter commit msg: " msg 
    esac 
done

# if no message is passed prompt for one
if [ -z "$msg" ]; then
    read -p "commit msg can't be empty: " msg
fi 

# run git commands
git add . && git commit -m "$msg" && git push