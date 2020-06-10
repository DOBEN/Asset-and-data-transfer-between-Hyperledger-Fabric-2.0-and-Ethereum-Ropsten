pragma solidity >=0.4.22 <0.7.0;


contract Storage {

    mapping(string=>string) assets;

 
    function store(string memory _key,string memory _value) public {
        assets[_key]=_value;
    }


    function retreive(string memory _key) public view returns (string memory){
        return assets[_key];
    }
}


//This smart contract was uploaded to the ropsten test network at "https://ropsten.etherscan.io/address/0xe3Bb1A8129d5Be6684c75ea3Fe884303e97638f5"
