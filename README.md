# medical-zkML-backend

This repository contains all code for the medical zkML backend.

## Quick Start
Configure database connections, IPFS information, and contract addresses in the configuration file.
database:
```yaml
database:
  mysql:
    driverName:
    host:
    port:
    user:
    password:
    schema:
    database:
```
ipfs:
```yaml
ipfs:
  url:
  auth:
```

contract:       
The validator's contract is stored in the 'internal/plugin/abi' path of the project, deployed to the chain, and the contract address is filled in the corresponding 'address:' mapping
```yaml
contract:
  contract_function: verify
  decision_tree+Acute_Inflammations:
    address:
  decision_tree+Breast_Cancer:
    address:
  decision_tree+Chronic_Kidney_Disease:
    address:
  decision_tree+Heart_Disease:
    address:
  decision_tree+Heart_Failure_Clinical_Records:
    address:
  decision_tree+Lymphography:
    address:
  decision_tree+Parkinsons:
    address:
  kMeans+Inflammations:
    address:
  kMeans+Breast_Cancer:
    address:
  kMeans+Chronic_Kidney_Disease:
    address:
  kMeans+Heart_Disease:
    address:
  kMeans+Heart_Failure_Clinical_Records:
    address:
  kMeans+Lymphography:
    address:
  kMeans+Parkinsons:
    address:
  XGBoost+Inflammations:
    address:
  XGBoost+Breast_Cancer:
    address:
  XGBoost+Chronic_Kidney_Disease:
    address:
  XGBoost+Heart_Disease:
    address:
  XGBoost+Heart_Failure_Clinical_Records:
    address:
  XGBoost+Lymphography:
    address:
  XGBoost+Parkinsons:
    address:
```

Introducing the required libraries for Python
```shell
pip install -r requirements.txt
```

Start Service
```shell
make build
bash medical &
```