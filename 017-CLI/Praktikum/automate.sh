firstname="$1"
personal="$2"
profesional="$3"
var=$(date)
sudo mkdir -p about_me/personal/
temp="${BASH_SOURCE:-$0}"
real="$(realpath "${temp}")"
dir="$(dirname "${real}")"  
cd  $dir/about_me/personal/
sudo touch sosmed.txt
sudo chmod 777 sosmed.txt
echo "https://sosmed/$personal" > sosmed.txt 
cd ..
sudo mkdir -p profesional
cd $dir/about_me/profesional
sudo touch linkedin.txt
sudo chmod 777 linkedin.txt
echo "https://linkedin//$profesional" > linkedin.txt
cd $dir
CRUL="usr/bin/curl"
link="https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt"
flags="-f -s -S -k"
raw="$(curl $link $flags)"
sudo mkdir -p my_friends/
cd $dir/my_friends/
sudo touch list_of_my_friends.txt
sudo chmod 777 list_of_my_friends.txt
echo "$raw" > list_of_my_friends.txt
cd $dir
sudo mkdir -p my_system_info/
cd $dir/my_system_info/
sudo touch about_this_laptop.txt
sudo touch internet_connection.txt
sudo chmod 777 about_this_laptop.txt
sudo chmod 777 internet_connection.txt
sudo lshw > about_this_laptop.txt
echo "connection to google" > internet_connection.txt
ping -c 4 google.com >> internet_connection.txt
