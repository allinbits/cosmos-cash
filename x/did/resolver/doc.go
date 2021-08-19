/*

Package resolver implements a universal resolver compatible REST API

https://github.com/decentralized-identity/universal-resolver


The configuration for this resolver are the following:


	{
		"pattern": "^(did:cosmos:.+)$",
		"url": "http://uni-resolver-driver-did-uport:8081/",
		"testIdentifiers": [
			"did:cosmos:cash:alice",
			"did:cosmos:key:cosmos1000000000000000000000000000000000000000000"
		]
	}


*/
package resolver
