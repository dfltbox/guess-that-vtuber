#!/bin/bash

# Loop through all subdirectories
for dir in */; do
  # Check if it's a directory
  if [ -d "$dir" ]; then
    # Change into the subdirectory
    cd "$dir" || continue

    # Run the ImageMagick commands in the current directory
    magick main.jpg -charcoal 1 -blur 0x5 medium.jpg
    magick main.jpg -attenuate 6 +noise Impulse -blur 0x5 -posterize 5 easy.jpg
    magick main.jpg -attenuate 4.5 +noise Impulse -blur 0x1 -threshold 50% extreme.jpg

    # Return to the original directory
    cd - || exit
  fi
done
