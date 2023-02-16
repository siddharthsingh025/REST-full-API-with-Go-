n=$1
while [ $n -gt 0 ] 
do
curl http://localhost:8080/albums
echo $(( n-- ))
done
