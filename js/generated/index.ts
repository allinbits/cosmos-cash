// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import AllinbitsCosmosCashAllinbitsCosmoscashDid from './allinbits/cosmos-cash/allinbits.cosmoscash.did'
import AllinbitsCosmosCashAllinbitsCosmoscashIssuer from './allinbits/cosmos-cash/allinbits.cosmoscash.issuer'
import AllinbitsCosmosCashAllinbitsCosmoscashVerifiablecredential from './allinbits/cosmos-cash/allinbits.cosmoscash.verifiablecredential'


export default { 
  AllinbitsCosmosCashAllinbitsCosmoscashDid: load(AllinbitsCosmosCashAllinbitsCosmoscashDid, 'allinbits.cosmoscash.did'),
  AllinbitsCosmosCashAllinbitsCosmoscashIssuer: load(AllinbitsCosmosCashAllinbitsCosmoscashIssuer, 'allinbits.cosmoscash.issuer'),
  AllinbitsCosmosCashAllinbitsCosmoscashVerifiablecredential: load(AllinbitsCosmosCashAllinbitsCosmoscashVerifiablecredential, 'allinbits.cosmoscash.verifiablecredential'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
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
