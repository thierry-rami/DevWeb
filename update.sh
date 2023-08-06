#!/bin/bash
 mkdir /srv/opt
sudo echo "/opt /srv/opt none rw,bind 0 0" > /etc/fstab
sudo systemctl daemon-reload
sudo mount -a

sudo  rm /usr/lib/python3.11/EXTERNALLY-MANAGED
pip3 list -o | cut -f1 -d' ' | tr " " "\n" | awk '{if(NR>=3)print}' | cut -d' ' -f1 | xargs -n1 pip3 install -U


