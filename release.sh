#!/bin/bash

set -e

usage()
{
  echo "usage: release.sh [[-v version ] | [-h]]"
}


version=
while getopts v:h: option
do
  case "${option}"
  in
  v) version=${OPTARG};;
  h) usage;;
  esac
done

if [ "$version" == "" ]; then
  echo "Please specify version/tag to release in format X.X.X"
  exit 1
else
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
    echo "git push origin master"
    git push origin master
  fi

  # Add tag
  echo "git tag -a v${version} -m \"Release v$version\""
  git tag -a v${version} -m "Release v$version"
  # Push
  echo "git push origin v${version}"
  git push origin v${version}
  # Create release
  echo "goreleaser --rm-dist"
  goreleaser --rm-dist
fi
