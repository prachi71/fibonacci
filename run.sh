#!/bin/bash

psport=`docker ps | grep 5432`

#TODO check for GIN port

if [ -z "${psport}" ]; then
    make pg.start
fi

make runLocal