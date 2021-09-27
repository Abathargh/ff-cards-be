#!/usr/bin/env bash

default_conf='{
  "port": 6666
}'

build_conf() {
  echo "Insert the port number: "
  read -r portNumber
  conf='{
    "port": '$portNumber'
  }'

  $conf > conf.json
}

if [ "$#" -gt 1 ]; then
    echo "wrong number of args"
    usage
    exit
fi

if [ -f conf.json ]; then
    echo "a conf file already exists"
    exit
fi

case $# in
    0) build_conf ù;;
    1)
      if [ "$1" = "default" ]; then
          $default_conf > conf.json
      else
        echo "expected 'default' arg, got $1"
      fi
esac
