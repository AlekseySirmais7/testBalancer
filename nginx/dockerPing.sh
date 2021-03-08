#!bun/bash


while true
do
curtime=$(date +"%H:%M:%S:") 
echo $curtime "docker nginx ping is OK"
sleep 5000
done
