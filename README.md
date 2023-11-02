# medical-zkML-backend

This repository contains all code for the medical zkML backend.

## Getting started
Get the code
```shell
git clone https://github.com/storswiftlabs/medical-zkML-backend.git
```

Server startup
```shell
cd medical-zkML-backend
make build
```

Introducing the required libraries for Python
```shell
pip install -r requirements.txt
```

Configure database connections, IPFS information, and agreed addresses in the configuration file under the project root path (config.yaml).
database:
```yaml
database:
  mysql:
    driverName: MYSQL
    host: LOCALHOST
    port: '3306'
    user: ROOT
    password: EXAMPLE
    schema: 
    database: MEDICAL-ZKML
```
IPFS:
```yaml
ipfs:
  url: https://api.nft.storage/upload
  auth: IPFS_AUTH
```

Contract:       
The validator's contract is stored in the 'internal/plugin/abi' path of the project, deployed to the chain, and the contract address is filled in the corresponding 'address:' mapping
```yaml
contract:
  contract_function: verify
  decision_tree+Acute_Inflammations:
    address: CONTRACT_ADDRESS
  decision_tree+Breast_Cancer:
    address: CONTRACT_ADDRESS
  decision_tree+Chronic_Kidney_Disease:
    address: CONTRACT_ADDRESS
  decision_tree+Heart_Disease:
    address: CONTRACT_ADDRESS
  decision_tree+Heart_Failure_Clinical_Records:
    address: CONTRACT_ADDRESS
  decision_tree+Lymphography:
    address: CONTRACT_ADDRESS
  decision_tree+Parkinsons:
    address: CONTRACT_ADDRESS
  kMeans+Inflammations:
    address: CONTRACT_ADDRESS
  kMeans+Breast_Cancer:
    address: CONTRACT_ADDRESS
  kMeans+Chronic_Kidney_Disease:
    address: CONTRACT_ADDRESS
  kMeans+Heart_Disease:
    address: CONTRACT_ADDRESS
  kMeans+Heart_Failure_Clinical_Records:
    address: CONTRACT_ADDRESS
  kMeans+Lymphography:
    address: CONTRACT_ADDRESS
  kMeans+Parkinsons:
    address: CONTRACT_ADDRESS
  XGBoost+Inflammations:
    address: CONTRACT_ADDRESS
  XGBoost+Breast_Cancer:
    address: CONTRACT_ADDRESS
  XGBoost+Chronic_Kidney_Disease:
    address: CONTRACT_ADDRESS
  XGBoost+Heart_Disease:
    address: CONTRACT_ADDRESS
  XGBoost+Heart_Failure_Clinical_Records:
    address: CONTRACT_ADDRESS
  XGBoost+Lymphography:
    address: CONTRACT_ADDRESS
  XGBoost+Parkinsons:
    address: CONTRACT_ADDRESS
```

Start service
```shell
bash medical &
```