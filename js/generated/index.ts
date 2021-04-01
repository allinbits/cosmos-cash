// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import AllinbitsCosmosCashAllinbitsCosmoscashVerifiablecredentialservice from './allinbits/cosmos-cash/allinbits.cosmoscash.verifiablecredentialservice'
import AllinbitsCosmosCashAllinbitsCosmoscashIdentifier from './allinbits/cosmos-cash/allinbits.cosmoscash.identifier'
import AllinbitsCosmosCashAllinbitsCosmoscashIssuer from './allinbits/cosmos-cash/allinbits.cosmoscash.issuer'
import AllinbitsCosmosCashAllinbitsCosmoscashIbcidentifier from './allinbits/cosmos-cash/allinbits.cosmoscash.ibcidentifier'


export default { 
  AllinbitsCosmosCashAllinbitsCosmoscashVerifiablecredentialservice: load(AllinbitsCosmosCashAllinbitsCosmoscashVerifiablecredentialservice, 'allinbits.cosmoscash.verifiablecredentialservice'),
  AllinbitsCosmosCashAllinbitsCosmoscashIdentifier: load(AllinbitsCosmosCashAllinbitsCosmoscashIdentifier, 'allinbits.cosmoscash.identifier'),
  AllinbitsCosmosCashAllinbitsCosmoscashIssuer: load(AllinbitsCosmosCashAllinbitsCosmoscashIssuer, 'allinbits.cosmoscash.issuer'),
  AllinbitsCosmosCashAllinbitsCosmoscashIbcidentifier: load(AllinbitsCosmosCashAllinbitsCosmoscashIbcidentifier, 'allinbits.cosmoscash.ibcidentifier'),
  
}


function load(mod, fullns) {
    return function init(store) {
        const fullnsLevels = fullns.split('/')
        for (let i = 1; i < fullnsLevels.length; i++) {
            let ns = fullnsLevels.slice(0, i)
            if (!store.hasModule(ns)) {
                store.registerModule(ns, { namespaced: true })
            }
        }
        if (store.hasModule(fullnsLevels)) {
            throw new Error('Duplicate module name detected: '+ fullnsLevels.pop())
        }else{
            store.registerModule(fullnsLevels, mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}
