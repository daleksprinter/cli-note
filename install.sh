#!/bin/bash
go build -o /usr/local/bin/note
cp ./config-example.yaml $HOME/.note_conf.yaml
