#!/bin/bash

# Heights to download traces, native logs, eth logs, and tipsets
heights=(14107 1467665 1162295)

# Function to download files
download_files() {
    local height=$1
    ./tracedl get --type traces --compress gz --height "$height" --outPath ../../data/heights
    ./tracedl get --type nativelog --compress gz --height "$height" --outPath ../../data/heights
    ./tracedl get --type ethlog --compress gz --height "$height" --outPath ../../data/heights
    ./tracedl get --type tipset --compress gz --height "$height" --outPath ../../data/heights
}

# Download files for each height
for height in "${heights[@]}"; do
    download_files "$height"
done
