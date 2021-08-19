#!/bin/bash

echo "Starting money mule sceneario"

# sets up an account with a verifiable credential
# takes 1 arg the name of the account to set up
set_up_account() {
	echo 'y' | cosmos-cashd keys add aml$2$1 --keyring-backend test

	### send native tokens so user can create did
	cosmos-cashd tx bank send \
		$(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show aml$2$1 --keyring-backend test -a) 100000stake --from validator --chain-id cash -y

	### create did
        cosmos-cashd tx did create-did --from aml$2$1 --keyring-backend test --chain-id cash -y

	### issue credential
        cosmos-cashd tx verifiablecredential create-verifiable-credential \
		did:cash:$(cosmos-cashd keys show aml$2$1 --keyring-backend test -a) $2-cred-$1 secret$1 name address dob nationalId phoneNumber --from validator --chain-id cash -y

	### attach credential to did document
        cosmos-cashd tx did add-service \
		did:cash:$(cosmos-cashd keys show aml$2$1 --keyring-backend test -a) $2-cred-$1 KYCCredential cash:$2-cred-$1 \
		--from aml$2$1 --keyring-backend test --chain-id cash -y


	echo "Querying data for aml$2$1"
	cosmos-cashd query did did did:cash:$(cosmos-cashd keys show aml$2$1 --keyring-backend test -a) --output json | jq
	cosmos-cashd query verifiablecredential verifiable-credential $2-cred-$1 --output json | jq

	sleep 1
}


# Destination accounts – set up 10 accounts to receive the coins
for i in {0..10}
do
	set_up_account $i dest &
	wait
done

# “mule accounts” – we need 10,000 of these setting up.
for i in {0..9}
do
	for j in {0..9} 
	do
		set_up_account ${i}${j} mule &
	done
	wait
done

# Originating accounts; we begin with 15 accounts all with a balance of €10K.
for i in {0..9}
do
	set_up_account $i origin

	### send token to fraud accounts
	cosmos-cashd tx bank send \
		$(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show amlorigin$i --keyring-backend test -a) 10000seuro --from validator --chain-id cash -y

	for j in {0..9} 
	do
		### send token from fraud accounts to mules
		cosmos-cashd tx bank send \
			$(cosmos-cashd keys show amlorigin$i --keyring-backend test -a) $(cosmos-cashd keys show amlmule$i$j --keyring-backend test -a) 100seuro --from amlorigin$i --chain-id cash -y
	
		echo "Querying data for amlmule$i$j"
		cosmos-cashd query bank balances $(cosmos-cashd keys show amlmule$i$j --keyring-backend test -a) --output json | jq
	done
done

# Transactions: We create 10,000 transactions of €150 (or if you want a more realistic version
# make the average transaction of €150 but each one varies between €135 and €165). These transactions are sent to the mule accounts. 

# The mule accounts deduct €20 and send the transactions to the destination accounts as suggested.
for i in {0..9}
do
	R=$(($RANDOM%10))
	echo $R
	for j in {0..9} 
	do
		### send token from mule accounts to dest
		cosmos-cashd tx bank send \
			 $(cosmos-cashd keys show amlmule${i}${j} --keyring-backend test -a) $(cosmos-cashd keys show amldest$R --keyring-backend test -a) 80seuro \
			 --from $(cosmos-cashd keys show amlmule${i}${j} --keyring-backend test -a) --chain-id cash -y &
	done
	wait

	echo "Querying data for amldest$R"
	cosmos-cashd query bank balances $(cosmos-cashd keys show amldest$R --keyring-backend test -a) --output json | jq
done

