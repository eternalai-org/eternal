#!/bin/bash

target_env=$1

ENVS=$(conda env list | awk '{print $1}' | grep -w $target_env)

conda init --all --dry-run --verbose

if [[ "$ENVS" == "" ]] 
then
    echo "Creating new environment"
    conda create -y -n $target_env python=3.9
    conda run -n $target_env python -V
    exit
else 
    echo $ENVS
    conda run -n $target_env python -V
fi;