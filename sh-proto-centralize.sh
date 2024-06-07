#!/bin/bash

PROJECT_ROOT="./"

rm -rf "${PROJECT_ROOT}/shared_proto"
mkdir -p "${PROJECT_ROOT}/shared_proto"

# Function to copy proto files, prioritizing directories with the same name as their parent
copy_protos() {
    local source_path=$1
    local dest_path=$2
    local parent_name=$(basename "$(dirname "$source_path")")

    # Loop through all files and directories in source path
    for entry in "$source_path"/*
    do
        # Get basename of the entry
        local base_entry=$(basename "$entry")
        # Check if entry is a directory
        if [ -d "$entry" ]; then
            # Check if the directory's name matches the parent directory's name
            if [ "$base_entry" == "$parent_name" ]; then
                # Prioritize by copying immediately
                mkdir -p "$dest_path/$base_entry"
                copy_protos "$entry" "$dest_path/$base_entry"
            fi
        fi
    done

    # Loop again for non-priority directories and files
    for entry in "$source_path"/*
    do
        local base_entry=$(basename "$entry")
        if [ -d "$entry" ]; then
            if [ ! -d "$dest_path/$base_entry" ] && [ "$base_entry" != "$parent_name" ]; then
                mkdir -p "$dest_path/$base_entry"
                copy_protos "$entry" "$dest_path/$base_entry"
            fi
        else
            # If entry is a file, copy it
            cp "$entry" "$dest_path/"
        fi
    done
}

# Scan the services directory and handle each proto directory found
for service in "${PROJECT_ROOT}/services"/*
do
    if [ -d "$service/proto" ]; then
        copy_protos "$service/proto" "${PROJECT_ROOT}/shared_proto"
    fi
done

echo "Proto files from all services have been copied to shared_proto directory."
