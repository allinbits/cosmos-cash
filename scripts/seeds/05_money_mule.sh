#!/bin/bash

echo "Starting money mule sceneario"

NUM_OF_DEST_ACCOUNTS=2
NUM_OF_MULE_ACCOUNTS=1
NUM_OF_ORIGIN_ACCOUNTS=2

# sets up an account with a verifiable credential
# takes 2 args the name of the account to set up, and the number of the account
set_up_account() {
	echo 'y' | cosmos-cashd keys add aml$2$1 --keyring-backend test

	### send native tokens so user can create did
	cosmos-cashd tx bank send \
		$(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show aml$2$1 --keyring-backend test -a) 100000stake --from validator --chain-id cash -y

	### create identifier
        cosmos-cashd tx identifier create-identifier --from aml$2$1 --keyring-backend test --chain-id cash -y

	### issue credential
        cosmos-cashd tx verifiablecredentialservice create-verifiable-credential \
		did:cash:$(cosmos-cashd keys show aml$2$1 --keyring-backend test -a) $2-cred-$1 --from validator --chain-id cash -y

	### attach credential to did document
        cosmos-cashd tx identifier add-service \
		did:cash:$(cosmos-cashd keys show aml$2$1 --keyring-backend test -a) $2-cred-$1 KYCCredential cash:$2-cred-$1 \
		--from aml$2$1 --keyring-backend test --chain-id cash -y


	echo "Querying data for aml$2$1"
	cosmos-cashd query identifier identifier did:cash:$(cosmos-cashd keys show aml$2$1 --keyring-backend test -a) --output json | jq
	cosmos-cashd query verifiablecredentialservice verifiable-credential $2-cred-$1 --output json | jq

	sleep 1
}


# Destination accounts – set up 10 accounts to receive the coins
for i in `seq 0 $NUM_OF_DEST_ACCOUNTS`
do
	set_up_account $i dest &
	wait
done

# “mule accounts” – we need 10,000 of these setting up.
for i in `seq 0 $NUM_OF_MULE_ACCOUNTS`
do
	for j in {0..9} 
	do
		set_up_account ${i}${j} mule &
	done
	wait
done

# Originating accounts; we begin with 15 accounts all with a balance of €10K.
for i in `seq 0 $NUM_OF_ORIGIN_ACCOUNTS`
do
	set_up_account $i origin

	### send token to fraud accounts
	cosmos-cashd tx bank send \
		$(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show amlorigin$i --keyring-backend test -a) 10000seuro --from validator --chain-id cash -y

	for j in {0..9} 
	do
		### send token from fraud accounts to mules
		cosmos-cashd tx bank send \
			$(cosmos-cashd keys show amlorigin$i --keyring-backend test -a) $(cosmos-cashd keys show amlmule$i$j --keyring-backend test -a) 100seuro \
			--from amlorigin$i --chain-id cash -y
	
		echo "Querying data for amlmule$i$j"
		cosmos-cashd query bank balances $(cosmos-cashd keys show amlmule$i$j --keyring-backend test -a) --output json | jq
	done
done

# Transactions: We create 10,000 transactions of €150 (or if you want a more realistic version
# make the average transaction of €150 but each one varies between €135 and €165). These transactions are sent to the mule accounts. 

# The mule accounts deduct €20 and send the transactions to the destination accounts as suggested.
for i in `seq 0 $NUM_OF_MULE_ACCOUNTS`
do
	R=$(($RANDOM%$NUM_OF_DEST_ACCOUNTS))
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

