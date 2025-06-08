#!/bin/bash

read -p "msg: " msg
git add . && \
git commit -m "$msg" && \
git push