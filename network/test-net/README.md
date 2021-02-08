### Network defines a BSN version of cosmos cash

### Overview
Terraform script for deploying a cluster to Digital Ocean and setting up a testnet with relayers configureed

Add your digital ocean token to `cluster/vars.tf`

### Prerequisites
Install [teraform](https://www.terraform.io/)

--- 

### k8s cluster
Deploy k8s cluster

```shell
cd cluster
terraform init
terraform plan
terraform approve
```
This will deploy a k8s cluster and download the `kube_config`

#### Destroy k8s cluster
```shell
terraform destory --auto-approve
```

### Deploy the BSN framwwork 
Please set up the k8s cluster before running these commands
```shell
cd test-net
kubectl create namespace ghost
terraform init
terraform plan
terraform approve
```

#### Destroy deployments

```shell
terraform destory --auto-approve
```

--- 

### k8s get pods

```shell
kubectl get pods --all-namespaces
```

### Get service accounts

```shell
kubectl get serviceaccounts --all-namespaces
```

#### Get namespaces

```shell
kubectl get namespaces
```

#### Set context 

```sh
kubectl config set-context --current --namespace=ghost
```

#### Ssh into machine 

```sh
kubectl exec --stdin --tty pod-name /bin/bash
```

#### Port forwarding 

```sh
kubectl port-forward pod 26657:26657
```

