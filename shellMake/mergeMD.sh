#!/bin/bash

# Define the directory containing the files
dir="./my_folder"

# Define the extension of the files you want to merge
ext=".md"

# Create an empty file to store the merged content
merged_file="merged_files.md"
touch $merged_file

# Loop through all files in the directory
for file in "$dir"/*$ext; do
    # Append the content of each file to the merged file
    cat "$file" >> $merged_file
    echo "" >> $merged_file
done