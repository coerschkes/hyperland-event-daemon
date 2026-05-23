#!/bin/bash

go generate ./...
go build .
systemctl --user stop hyprland-event-daemon.service
sudo mv ./hyprland-event-daemon /usr/local/bin
systemctl --user restart hyprland-event-daemon.service
