#!/bin/bash

# Written by Kenneth (Alex) Jenkins for use in ArgonV2

# Define guarded apps and locations
guardedApps=("JetBrains") # Add more apps separated by space
guardedLocations=("~/Library/Caches" "~/Library/Logs" "/Library/Caches" "/Library/Logs")
userPrefLogs="~/Library/Preferences"
systemLogs="/var/log"
iosPhotoCache="~/Pictures/iPhoto Library/iPod Photo Cache"

# Function to clean files from a location excluding guarded apps
clean_location() {
    local location="$1"
    echo "Cleaning files from $location"
    find "$location" -mindepth 1 -maxdepth 1 \( ! -path "*JetBrains*" \) -exec rm -rf {} +
    echo "Done cleaning files from $location"
}

# Clean guarded locations
for guardedLocation in "${guardedLocations[@]}"; do
    clean_location "$guardedLocation"
done

# Clean user preference logs
echo "Cleaning user preference logs"
# Uncomment the line below if you want to clean user preference logs
#rm -rf "$userPrefLogs"/*
echo "Done cleaning user preference logs"

# Clean system logs
echo "Cleaning system logs from $systemLogs"
sudo rm -rf "$systemLogs"/*
echo "Done cleaning system logs"

# Clean iOS photo caches
echo "Cleaning iOS photo caches"
rm -rf "$iosPhotoCache"/*
echo "Done cleaning iOS photo caches"

# Clean application caches
echo "Cleaning application caches"
for appDir in ~/Library/Containers/*; do
    if [ -d "$appDir" ]; then
        echo "Cleaning $appDir/Data/Library/Caches/"
        rm -rf "$appDir/Data/Library/Caches"/*
        echo "Done cleaning $appDir/Data/Library/Caches"
    fi
done
echo "Done cleaning application caches"
echo "All done!\nYou should restart your Mac as soon as possible!"