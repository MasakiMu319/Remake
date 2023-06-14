#!/bin/zsh

count=0; while ((count<5)); do curl localhost:8080/ping;echo ""; ((count+=1)); sleep 1; done;