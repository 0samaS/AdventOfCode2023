#!/bin/bash

# Check if the correct number of arguments is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <day_number>"
    exit 1
fi

input_file="sols/day_template.go"

# Check if the input file exists
if [ ! -f "$input_file" ]; then
    echo "Error: File '$input_file' not found."
    exit 1
fi

num=$1

mkdir -p inputs/day${num}/
touch inputs/day${num}/test.txt
touch inputs/day${num}/real.txt

output_file="sols/day${num}.go"
rm output_file

# Read the input file line by line, replace "X" with variable, and write to temporary file
while IFS= read -r line; do
    modified_line=$(echo "$line" | sed "s/X/$num/g")
    echo "$modified_line" >> "$output_file"
done < "$input_file"

echo "Modified file created: $output_file"
