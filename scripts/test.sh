#!/usr/bin/env bash

###
# Runs the service tests.
###

function read_env_file {
  set -a
  [ -f .env.test ] && source .env.test
  set +a
}

##
# Main function
##
function main {
  # Read test env file
  read_env_file

  # Run tests.
  go test -v ./...
}

# And so, it begins...
main
