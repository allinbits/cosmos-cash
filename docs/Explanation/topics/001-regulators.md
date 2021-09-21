# Credentials verification

Verification of data on the Cosmos Cash network happens on two levels: the first level is the classic transactions signature, that is used on consensus level to maintain the correct and tamper proof state of the blockchain; the second level is provided by the verifiable credentials proof scheme, where a verifiable credentials is signed by the issuer of the credentials.

Both types of verifications happen in a Cosmos Cash powered blockchain, but the two verifications serve a different goal, one to make sure that the transaction committed on chain are correct in terms of consensus, the other that the verifiable credentials contained in a transaction are signed by the issuer of the credentials. 


## Regulators

Regulators are the root of trust in the cosmos-cash chain: regulators chain addresses are know before chain launch and their addresses are stored in the genesis file. 

Regulators have the authority to issue registration and license verifiable credentials to E-Money issuers accounts, and through those licenses the issuers can create/mint/burn/redeem tokens on a cosoms-cash powered chain.

## Activation 

A regulator must perform an activation to be able to start issuing verifiable credentials, the activation consists of issuing a **regulator credential** to a DID of its choice. The transaction to issue the regulator credential MUST be signed by an account listed in the regulator addresses in the genesis file. 

In DID indicated as the issuer in the regulator credentials MUST be:

- resolvable: it must be possible to retrieve the DID document that describes the vc issuer DID
- with a verification method listed in the `authentication` or `assertionMethod` relationships
- with the verification method listing a public key or an address with a know public key; in both cases the signature verification scheme used is secp256k1

The following state diagram illustrates the process of activation of a regulator did

![](001-regulators-activate.svg)

