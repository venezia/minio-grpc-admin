#!/usr/bin/env sh

if [[ "$#" -ne 2 ]]; then
  echo "Usage: `basename $0`<path to chart> <release version>"
  exit 1
fi

parse.py --file ${1}/Chart.yaml --key-val version=${2}
parse.py --file ${1}/values.yaml --key-val image.tag=${2}
parse.py --file ${1}/Chart.yaml --key-val appVersion=$(git rev-parse --short HEAD)