#!/usr/bin/env bash
mkdir ~/.config/sprint/

go build -o sprint cmd/sprint/main.go &&
    rm ~/.local/bin/sprint &&
    mv sprint ~/.local/bin &&
    cp sprint-config/* ~/.config/sprint/
