#!/bin/bash

set -e

usage()
{
  echo "usage: release.sh [[-v version ] | [-h]]"
}


version=
comment=
while getopts v:c:h: option
do
  case "${option}"
  in
  v) version=${OPTARG};;
  c) comment=${OPTARG};;
  h) usage;;
  esac
done

# while [ "$1" != "" ]; do
#     case $1 in
#         -v | --version )        version=$1
#                                 ;;
#         -c | --comment )        version=$1
#                                 ;;
#         -h | --help )           usage
#                                 exit
#                                 ;;
#         * )                     usage
#                                 exit 1
#     esac
#     shift
# done

if [ "$version" == "" ]; then
  echo "Please specify version/tag to release in format X.X.X"
  exit 1
else
  if [ "$comment" == "" ]; then
    comment="Release v$version"
  fi

  cli_version=`cat cli.json | jq ".commands[0].version" | tr -d '"'`
  echo "Current CLI version: ${cli_version}"
  echo "Proposed version: ${version}"
  if [ "$version" != "$cli_version" ]; then
    # Set version in cli.json
    jq ".commands[0].version = \"$version\"" cli.json > tmp_cli.json
    mv tmp_cli.json cli.json

    # Commit cli.json changes
    echo "git add cli.json"
    git add cli.json
    echo "git commit -m 'Updated version in cli.json'"
    git commit -m 'Updated version in cli.json'
  fi
    echo "dep ensure -update"
    dep ensure -update
    dep status
    echo "git add -A && git commit -a -m '$comment'"
    git add -A && git commit -a -m "$comment"
    echo "git tag -a v$version -m '$comment'"
    git tag -a v$version -m "$comment"
    echo "git push origin master --tags"
    git push origin master --tags
fi