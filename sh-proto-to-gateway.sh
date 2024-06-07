#!/bin/bash

# Set the project root to the current directory
PROJECT_ROOT="./"

# Define the source directory (shared_proto) and the destination directory (grpc-gateway/proto)
SOURCE_DIR="${PROJECT_ROOT}/shared_proto"
DESTINATION_DIR="${PROJECT_ROOT}/grpc-gateway/proto"

# Function to copy proto files, creating directories as needed
copy_protos() {
    local source_path=$1
    local dest_path=$2

    # Ensure the destination directory exists
    mkdir -p "$dest_path"

    # Loop through all files and directories in source path
    for entry in "$source_path"/*
    do
        # Get basename of the entry
        local base_entry=$(basename "$entry")

        # Check if entry is a directory
        if [ -d "$entry" ]; then
            # Recursively call function to handle subdirectories
            copy_protos "$entry" "$dest_path/$base_entry"
        else
            # If entry is a file, copy it
            cp "$entry" "$dest_path/"
        fi
    done
}

# Remove existing files and directories in the destination directory before copying
remove_existing_files() {
    echo "Removing existing files in $DESTINATION_DIR..."
    rm -rf "$DESTINATION_DIR"/*
}

# Check if the source directory exists before proceeding
if [ -d "$SOURCE_DIR" ]; then
    # Clear the grpc-gateway/proto directory first
    remove_existing_files
    # Start copying proto files
    copy_protos "$SOURCE_DIR" "$DESTINATION_DIR"
    echo "Proto files have been copied from shared_proto to grpc-gateway/proto directory."
else
    echo "Source directory does not exist. Please check the path: $SOURCE_DIR"
fi
