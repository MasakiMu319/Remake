#!/usr/bin/env bash

read -p "Enter username: " username
read -s -p "Enter password: " pass
useradd "$username"
# echo "$pass" | passwd --stdin "$username"
# ubuntu 从 16 的版本后 passwd 就不再支持 stdin
echo "$username:$pass" | chpasswd