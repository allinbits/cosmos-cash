#!/bin/bash

cosmos-cashd tx identifier create-identifier --from validator --chain-id cash -y

cosmos-cashd keys add auth2

cosmos-cashd tx identifier add-authentication did:cash:cosmos1g55vhuxc95q5mfaz7stl5wh72rkpumguemnt7r cosmospub1addwnpepqtqutllgy55y33078dw480zlrspnvtepnfyq7x3nzhpx8vgzju3gs0ungys --from validator --chain-id cash -y

cosmos-cashd tx identifier add-service did:cash:cosmos18qh73ky6dt8xzrnrswdp7rmjda3r8w8tv4f0r3 new-verifiable-cred-3 KYCCredential cosmos-cash:new-verifiable-cred-3 --from validator --chain-id cash -y
