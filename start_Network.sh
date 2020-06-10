cd ./fabric-network/first-network/
echo y | ./byfn.sh down

echo y | ./byfn.sh generate -a -s couchdb -l golang  

echo y | ./byfn.sh up -a  -s couchdb -l golang  

cd ../../website 

sudo rm -R wallet/

node enrollAdmin 1
node registerUser 1 
node app.js
