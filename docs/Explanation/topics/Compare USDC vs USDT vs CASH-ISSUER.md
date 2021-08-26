# USDC vs USDT vs Cosmos Cash Issuer

- [USDC vs USDT vs Cosmos Cash Issuer](#usdc-vs-usdt-vs-cosmos-cash-issuer)
  - [Summary](#summary)
  - [Methodology](#methodology)
  - [TL;DR;](#tldr)
  - [USDT ERC20 Smart Contracts](#usdt-erc20-smart-contracts)
    - [Source](#source)
    - [Analysis](#analysis)
    - [Roles](#roles)
    - [Functions](#functions)
  - [USDC ERC-20 smart contracts](#usdc-erc-20-smart-contracts)
    - [Source](#source-1)
    - [Analysis](#analysis-1)
    - [Roles](#roles-1)
    - [Functions](#functions-1)
  - [Cosmos Cash Issuer Module](#cosmos-cash-issuer-module)
    - [Source](#source-2)
    - [Functions](#functions-2)
  - [Gap Analysis](#gap-analysis)


---
## Summary

USDC and USDT are both Ethereum, ERC-20 smart contract-based stablecoins pegged one-to-one to the US Dollar (USD). However, each USD-backed stablecoin has different implementations. This article will

- Perform a high-level functional gap analysis of each smart contract
- Compare to the current Cosmos Cash proof-of-concept implementation
- Conclude best practices that will inform the Issuer ADR for Cosmos Cash.

> Why USDT and  USDC? These are currently the leading stablecoin tokens with total market capitalisation of over 90Bn USD (source: coinmarketcap.com)

## Methodology

Each function is labelled based on the following criteria:

- **BESPOKE:** custom functionality
- **[ERC20](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/ERC20.sol):** part of the ERC20 standard
- **[PAUSABLE](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/security/Pausable.sol):** part of the solidity pausable contract standard
- **[OWNABLE](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/access/Ownable.sol):** part of the solidity ownable contract standard 
- **[PROXY](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/proxy/Proxy.sol):** these functions are part of Solidity’s delegate proxy upgradability pattern
- **[MINTABLE](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol):** part of the solidity mintable contract standard
- **[BURNABLE](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/extensions/ERC20Burnable.sol):** part of the solidity burnable contract standard
- **[RBAC](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/access/AccessControl.sol):** part of the solidity roles contract standard 

## TL;DR;

| Label     | USDT                                 | USDC                                                                                                                   | Cosmos Cash Issuer              |
|:--------- |:------------------------------------ |:---------------------------------------------------------------------------------------------------------------------- |:------------------------------- |
| Mintable  | issue(amount)                        | mint(\_to, \_amount)                                                                                                   | mintToken(amount, owner)        |
| Burnable  | redeem(amount)                       | burn(\_amount)                                                                                                         | burnToken(amount, owner)        |
| Pausable  | pause                                | pause                                                                                                                  | TO DO                           |
| Pausable  | unpause                              | unpause                                                                                                                | TO DO                           |
| blacklist | addBlacklist                         | blacklist(\_account)                                                                                                   | TO DO                           |
| blacklist | removeBlacklist(\_clearedUser)       | unBlacklist(\_account)                                                                                                 | TO DO                           |
| blacklist | destroyBlackFunds(\_blackListedUser) | n/a                                                                                                                    | TO DO                           |
| blacklist | n/a                                  | rescueERC20(tokenContract, to, amount)                                                                                 | TO DO                           |
| ERC20     | transfer(\_from, \_value)            | transfer(\_from, \_value)                                                                                              |                                 |
| ERC20     | transferFrom(\_from, \_to, \_value)  | transferFrom(\_from, \_to, \_value)                                                                                    |                                 |
| ERC20     | approve(\_spender, \_value)          | approve(\_spender, \_value)                                                                                            |                                 |
| Ownable   | transferOwnership(newOwner)          | transferOwnership(newOwner)                                                                                            |                                 |
| Bespoke   | setParams(newBasisPoints, newMaxFee) | n/a                                                                                                                    |                                 |
| Bespoke   | n/a                                  | transferWithAuthorization(from, to, value, validAfter, validBefore, nonce, v, r, s)                                    |                                 |
| Bespoke   | n/a                                  | configureMinter(minter, minterAllowedAmount)                                                                           |                                 |
| Bespoke   | n/a                                  | receiveWithAuthorization(from, to, value, validAfter, validBefore, nonce, v, r, s)                                     |                                 |
| Bespoke   | n/a                                  | cancelAuthorization(authorizer, nonce, v, r, s)                                                                        |                                 |
| RBAC      |                                      | increaseAllowance(spender, increment)                                                                                  |                                 |
| RBAC      |                                      | decreaseAllowance(spender, decrement)                                                                                  |                                 |
| RBAC      | n/a                                  | removeMinter(minter)                                                                                                   |                                 |
| RBAC      | n/a                                  | updateBlacklister(\_newBlacklister)                                                                                    |                                 |
| RBAC      | n/a                                  | updateMasterMinter(\_newMasterMinter)                                                                                  |                                 |
| RBAC      | n/a                                  | updatePauser(\_newPauser)                                                                                              |                                 |
| RBAC      | n/a                                  | updateRescuer(newRescuer)                                                                                              |                                 |
| Proxy   | n/a                                  | initialize(tokenName, tokenSymbol, tokenCurrency, tokenDecimals, newMasterMinter, newPauser, newBlacklister, newOwner) | createIssuer(token, fee, owner) |
| Proxy   | n/a                                  | initializeV2(string)                                                                                                   |                                 |
| Proxy   | n/a                                  | initializeV2\_1(newName)                                                                                               |                                 |
| Proxy   | deprecate(\_upgradedAddress)         | n/a                                                                                                                    |                                 |


---
## USDT ERC20 Smart Contracts

### Source

This analysis used this definition of the [USDT ERC-20 smart contract](https://etherscan.io/token/0xdac17f958d2ee523a2206206994597c13d831ec7#writeContract). This is found at  0xdac17f958d2ee523a2206206994597c13d831ec7

### Analysis

The USDT smart contracts have **five** main areas of note:

- Tokens freezing PAUASABLE
- Tokens minting, burning MINTABLE, BURNABLE
- Role-Based Access Control
- User denylisting
- Upgrading function (albeit limited)

### Roles
This contract has admin privileges for **TWO** roles, each comprising of one address:

- **OWNER** - owner owns the contract and can call admin functions
- **BLACKLISTER** - can add and remove users from a blacklist

### Functions

`issue(amount)` - **BESPOKE/MINTABLE:** mints a certain amount of tokens.

`redeem(amount)` - **BESPOKE/BURNABLE:** redeems or burns an amount tokens.

`pause` - **PAUSABLE:** pauses the token transfers. This function can only be called by the contract owner, the issuer of USDT tokens.

`unpause` - **PAUSABLE:** unpauses token transfers only callable by the token.

`addBlacklist(_evilUser)` - **BESPOKE:** this function freezes a user's assets by adding an address to a denylist.

`removeBlacklist(_clearedUser)` - **BESPOKE:** removes a user from a denylist serves, effectively unfreezing a user account.

`destroyBlackFunds(_blackListedUser)` - **BESPOKE:** burns tokens of a blacklisted address.

`transfer(_from, _value)` - **ERC20:** transfer tokens from one user to another.

`transferFrom(_from, _to, _value)` - **ERC20:** Transfers tokens from one address to another, usually called in conjunction with the `approve` function.

`approve(_spender, _value)` - **ERC20:** this function approves an address or contracts to use funds on behalf of another user.

`transferOwnership(newOwner)` - **OWNABLE:** transfers the owner of the contracts to another address.

`setParams(newBasisPoints, newMaxFee)` - This is **BESPOKE:** functionality that sets specific params on the contract, probable something to do with earning on each transaction; currently, both are zero.

`deprecate(_upgradedAddress)` - **BESPOKE:** this function is used as part of an upgradability pattern and sets a previous contract as deprecated. The use of this pattern is specific to the upgrading of Ethereum contracts. The reasons for this e are explained in this [Coinbase design blog post](https://blog.coinbase.com/usdc-v2-upgrading-a-multi-billion-dollar-erc-20-token-b57cd9437096)

---
## USDC ERC-20 smart contracts
### Source

This analysis used this definition of the [USDC ERC-20 smart contract](https://etherscan.io/token/0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48#writeProxyContract). This is found at  0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48

### Analysis

The USDC contract is more complex than USDT, but like USDT, the functions are based upon **five** function types:

- Tokens freezing.
- Tokens minting/burning.
- RBAC.
- User denylisting.
- Upgrades, using a **Delegate Proxy** pattern.

### Roles

Unlike USDT, USDC has more granular permissions. There are **FIVE** roles, each comprising one address:

- **OWNER** - owner owns the contract and can call admin functions.
- **BLACKLISTER** - can add and remove users from a blacklist.
- **MASTERMINTER** - can add and remove minters and update minter allowances.
- **RESCUER** - can rescue funds from a user, currently a NULL address.
- **PAUSER** - can call the pause and unpause functions on the contract.

### Functions

`mint(_to, _amount)` - **MINTABLE:** mint tokens to a given user.

`burn(_amount)` - **BURNABLE:** burns a certain number of tokens.

`pause` - **PAUSABLE:** pauses the token transfers in the contract, only callable by the owner.

`unpause` - **PAUSABLE:** unpauses token transfers only callable.

`blacklist(_account)` - **BESPOKE:** this function freezes a user's assets by adding the address to a denylist.

`unBlacklist(_account)` - **BESPOKE:** removes a user from a denylist serves as freezing an accounts.

`rescueERC20(tokenContract, to, amount)` - **BESPOKE:** function to rescue funds.

`transfer(_from, _value)` - **ERC20:** transfer tokens from one user to another.

`transferFrom(_from, _to, _value)` - **ERC20:** Transfers tokens from one address to another, usually called in conjunction with the `approve` function.

`approve(_spender, _value)` - **ERC20:** this function approves an address or contract to use funds on behalf of another user.

`transferOwnership(newOwner)` - **OWNABLE:** transfers the owner of the contracts to another address.

`transferWithAuthorization(from, to, value, validAfter, validBefore, nonce, v, r, s)` - **BESPOKE**: This allows a user to send tokens with a signature from another user. 

`cancelAuthorization(authoriser, nonce, v, r, s)` - **BESPOKE:** This stops a user from sending tokens with a signature from a user. 
 
`configureMinter(minter, minterAllowedAmount)` - **BESPOKE:** allows a minter to mint a certain amount of tokens.

`increaseAllowance(spender, increment)` - **BESPOKE:** This increases the amount a user can spend per transaction for another user.

`decreaseAllowance(spender, decrement)` - **BESPOKE:** This decreases the amount a user can spend per transaction for another user.

`recieveWithAuthorization(from, to, value, validAfter, validBefore, nonce, v, r, s)` - **BESPOKE**: This allows a user to send tokens with a signature from a user.

`removeMinter(minter)` - **RBAC:** removes an address from the minter role callable by **MASTERMINTER**.

`updateBlacklister(_newBlacklister)` - **RBAC:** update the admin role **BLACKLIST**.

`updateMasterMinter(_newMasterMinter)` - **RBAC:** update the admin role **MASTERMINTER**.

`updatePauser(_newPauser)` - **RBAC:** update the admin role **PAUSER**.

`updateRescuer(newRescuer)` - **RBAC:** update the admin role **RESCUER**.

`initialize(tokenName, tokenSymbol, tokenCurrency, tokenDecimals, newMasterMinter, newPauser, newBlacklister, newOwner)` - **PROXY:** This functionality is used as part of the delegate proxy pattern. It initialises the smart contract and allows delegate proxy functionality. The use of this pattern is specific to the upgrading of Ethereum contracts. The reasons for this e are explained in this [Coinbase design blog post](https://blog.coinbase.com/usdc-v2-upgrading-a-multi-billion-dollar-erc-20-token-b57cd9437096).

`initializeV2(string)` - **PROXY:** This functionality is used as part of the delegate proxy pattern; it initialises the contract and allows delegate proxy functionality.

`initializeV2_1(newName)` - **PROXY:** This functionality is used as part of the delegate proxy pattern; it initialises the contract and allows delegate proxy functionality.

---
## Cosmos Cash Issuer Module

### Source

This analysis is based on the [Cosmos Cash proof of concept issuer implementation](https://github.com/allinbits/cosmos-cash/blob/main/x/issuer/keeper/handler.go#L21-L29).

### Functions

`createIssuer(token, fee, owner)` - This creates an issuer and a token. This initialises the token, and the issuer is the token owner. This is where we could define the token’s PAUSER, BLACKLISTER, MASTERMINTER or other ROLES.

`mintToken(amount, owner)` - mints new tokens to the owner’s address.

`burnToken(amount, owner)` - removes tokens from the owner’s address.

---
## Gap Analysis

The gap between the issuer module and the USDC and USDT smart contracts is not significant. For the Issuer module to be compatible with EVM based smart contracts, the following features must be implemented.

1.  **BESPOKE:** `Blockedlist` 
    - The blocklist functionality is adding and removing a user from a blocked list
        - addToBlockedList
        - removeFromBlocklist
        - burnBlockedlistTokens

2. **PAUSABLE:** `Freeze/Pause` token 
    - Freeze token functionality is a kill switch that stops all trading with an issuer token
        - freezeToken/pauseToken
        - unfreezeToken/unpauseToken

3.  **RBAC:** A `Role-Based Access Control` credential system is required to interact with the contract or module. There are different implementation options available:
    1.  Defining admin roles in the genesis block.
    2.  This could be implemented using Decentralized IDentity and **public** verifiable credentials such that:
        - The verifiable credential can be issued by a regulator actor or DID to an issuer DID.
        - The verifiable credential can allow minting and burning of tokens plus other functions
        - The Issuer DID can have multiple DID controllers who can represent different functions in the issuers, such as Operations and Compliance.

4.  **ERC20:** Some `ERC20` token functions are not in the SDK bank module
        - transferFrom(to, from, amount)
        - approve(address)
