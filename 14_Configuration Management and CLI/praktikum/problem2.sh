#!/bin/bash

folder_name="$1 at $(date)"

mkdir -p "$folder_name/my_system_info"

mkdir -p "$folder_name/about_me/personal"
echo "https://www.facebook.com/$2" > "$folder_name/about_me/personal/facebook.txt"

mkdir -p "$folder_name/about_me/professional"
echo "https://www.linkedin.com/in/$3" > "$folder_name/about_me/professional/linkedin.txt"


mkdir -p "$folder_name/my_friends"
touch "$folder_name/my_friends/list_of_my_friends.txt"

echo "My username: $(hostname)" > "$folder_name/my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "$folder_name/my_system_info/about_this_laptop.txt"

echo "$(ping google.com -c 3)" >> "$folder_name/my_system_info/internet_connection.txt"
