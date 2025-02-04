#!/usr/bin/env bash

docker build . -t gtfs_tracker/env

docker run -v .:/wa/gtfs_tracker -it gtfs_tracker/env /bin/bash
