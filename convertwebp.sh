#!/bin/bash

# Check if a directory argument is provided
if [ -z "$1" ]; then
  echo "Usage: $0 /path/to/your/folder"
  exit 1
fi

# Get the root directory from the argument
ROOT_DIR="$1"

# Check if the directory exists
if [ ! -d "$ROOT_DIR" ]; then
  echo "Error: Directory '$ROOT_DIR' does not exist."
  exit 1
fi

# Find all PNG files in the directory and subdirectories
find "$ROOT_DIR" -type f \( -iname "*.png" -o -iname "*.jpg" \) | while read -r file; do
  # Define the output file name by changing the extension to .webp
  case "$file" in
  *.png)
    output="${file%.png}.webp"
    ;;
  *.jpg)
    output="${file%.jpg}.webp"
    ;;
  *.jpeg)
    output="${file%.jpeg}.webp"
    ;;
  esac

  echo "converting $file to $output"

  # Encode the PNG to WebP using cwebp
  cwebp "$file" -o "$output"

  # Optionally, remove the original PNG file after conversion
  rm "$file"
done

echo "All PNG files have been converted to WebP."
