#!/usr/bin/env bash

set -e

BASEDIR=$(dirname "$0")

VIDULUMD_BINARY=$(which vidulumd || (echo -e "\033[31mPlease add vidulumd to PATH\033[0m" 1>&2 && exit 1))
VIDULUMD_USER=$USER
VIDULUMD_BINARY_DIR=$(dirname $(which vidulumd))
VIDULUMD_USER_HOME=$(eval echo "~$USER")

sed "s#<VIDULUMD_BINARY>#$VIDULUMD_BINARY#g; s#<VIDULUMD_USER>#$VIDULUMD_USER#g; s#<VIDULUMD_BINARY_DIR>#$VIDULUMD_BINARY_DIR#g; s#<VIDULUMD_USER_HOME>#$VIDULUMD_USER_HOME#g" $BASEDIR/vidulumd.service.template > $BASEDIR/vidulumd.service

echo -e "\033[32mGenerated $BASEDIR/vidulumd.service\033[0m"

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  sudo cp $BASEDIR/vidulumd.service /etc/systemd/system/vidulumd.service
  sudo systemctl daemon-reload
  sudo systemctl enable vidulumd.service
  echo -e "\033[32mCreated /etc/systemd/system/vidulumd.service\033[0m"
else
  echo -e "\033[31mCan only create /etc/systemd/system/vidulumd.service for linux\033[0m" 1>&2
  exit 1
fi