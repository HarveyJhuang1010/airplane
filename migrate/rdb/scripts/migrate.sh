#!/usr/bin/env bash

# 記得去 brew install flyway
# 記得去 brew install openjdk

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <path_to_flyway_config>"
    exit 1
fi

execEnv=$1

if [ "$execEnv" == "prod" ]; then
    read -p "You are about to run the migration in the 'prod' environment. Please confirm by typing 'prod': " confirmation
    if [ "$confirmation" != "prod" ]; then
        echo "Confirmation failed. Aborting migration."
        exit 1
    fi
fi

GITROOT=$(git rev-parse --show-toplevel)
flywayPath="${GITROOT}/migrate/rdb"
configPath="scripts/flyway-${execEnv}.conf"

cd ${flywayPath}
echo $pwd
flyway -configFiles=${configPath} migrate
#rm report.html
#rm report.json
