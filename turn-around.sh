#!/bin/bash

go generate ./...
go build .
systemctl --user restart hyprland-event-daemon.service
