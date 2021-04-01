import { txClient, queryClient } from './module';
// @ts-ignore
import { SpVuexError } from '@starport/vuex';
import { VerifiableCredential } from "./module/types/verifiable-credential-service/verifiable-credential";
import { UserCredentialSubject } from "./module/types/verifiable-credential-service/verifiable-credential";
import { IssuerCredentialSubject } from "./module/types/verifiable-credential-service/verifiable-credential";
import { Proof } from "./module/types/verifiable-credential-service/verifiable-credential";
async function initTxClient(vuexGetters) {
    return await txClient(vuexGetters['common/wallet/signer'], {
        addr: vuexGetters['common/env/apiTendermint']
    });
}
async function initQueryClient(vuexGetters) {
    return await queryClient({
        addr: vuexGetters['common/env/apiCosmos']
    });
}
function getStructure(template) {
    let structure = { fields: [] };
    for (const [key, value] of Object.entries(template)) {
        let field = {};
        field.name = key;
        field.type = typeof value;
        structure.fields.push(field);
    }
    return structure;
}
const getDefaultState = () => {
    return {
        VerifiableCredentials: {},
        VerifiableCredential: {},
        ValidateVerifiableCredential: {},
        _Structure: {
            VerifiableCredential: getStructure(VerifiableCredential.fromPartial({})),
            UserCredentialSubject: getStructure(UserCredentialSubject.fromPartial({})),
            IssuerCredentialSubject: getStructure(IssuerCredentialSubject.fromPartial({})),
            Proof: getStructure(Proof.fromPartial({})),
        },
        _Subscriptions: new Set(),
    };
};
// initial state
const state = getDefaultState();
export default {
    namespaced: true,
    state,
    mutations: {
        RESET_STATE(state) {
            Object.assign(state, getDefaultState());
        },
        QUERY(state, { query, key, value }) {
            state[query][JSON.stringify(key)] = value;
        },
        SUBSCRIBE(state, subscription) {
            state._Subscriptions.add(subscription);
        },
        UNSUBSCRIBE(state, subscription) {
            state._Subscriptions.delete(subscription);
        }
    },
    getters: {
        getVerifiableCredentials: (state) => (params = {}) => {
            if (!params.query) {
                params.query = null;
            }
            return state.VerifiableCredentials[JSON.stringify(params)] ?? {};
        },
        getVerifiableCredential: (state) => (params = {}) => {
            if (!params.query) {
                params.query = null;
            }
            return state.VerifiableCredential[JSON.stringify(params)] ?? {};
        },
        getValidateVerifiableCredential: (state) => (params = {}) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ValidateVerifiableCredential[JSON.stringify(params)] ?? {};
        },
        getTypeStructure: (state) => (type) => {
            return state._Structure[type].fields;
        }
    },
    actions: {
        init({ dispatch, rootGetters }) {
            console.log('init');
            if (rootGetters['common/env/client']) {
                rootGetters['common/env/client'].on('newblock', () => {
                    dispatch('StoreUpdate');
                });
            }
        },
        resetState({ commit }) {
            commit('RESET_STATE');
        },
        unsubscribe({ commit }, subscription) {
            commit('UNSUBSCRIBE', subscription);
        },
        async StoreUpdate({ state, dispatch }) {
            state._Subscriptions.forEach((subscription) => {
                dispatch(subscription.action, subscription.payload);
            });
        },
        async QueryVerifiableCredentials({ commit, rootGetters, getters }, { options: { subscribe = false, all = false }, params: { ...key }, query = null }) {
            try {
                let value = query ? (await (await initQueryClient(rootGetters)).queryVerifiableCredentials(query)).data : (await (await initQueryClient(rootGetters)).queryVerifiableCredentials()).data;
                while (all && value.pagination && value.pagination.nextKey != null) {
                    let next_values = (await (await initQueryClient(rootGetters)).queryVerifiableCredentials({ ...query, 'pagination.key': value.pagination.nextKey })).data;
                    for (let prop of Object.keys(next_values)) {
                        if (Array.isArray(next_values[prop])) {
                            value[prop] = [...value[prop], ...next_values[prop]];
                        }
                        else {
                            value[prop] = next_values[prop];
                        }
                    }
                }
                commit('QUERY', { query: 'VerifiableCredentials', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryVerifiableCredentials', payload: { options: { all }, params: { ...key }, query } });
                return getters['getVerifiableCredentials']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                console.error(new SpVuexError('QueryClient:QueryVerifiableCredentials', 'API Node Unavailable. Could not perform query.'));
                return {};
            }
        },
        async QueryVerifiableCredential({ commit, rootGetters, getters }, { options: { subscribe = false, all = false }, params: { ...key }, query = null }) {
            try {
                let value = query ? (await (await initQueryClient(rootGetters)).queryVerifiableCredential(key.verifiable_credential_id, query)).data : (await (await initQueryClient(rootGetters)).queryVerifiableCredential(key.verifiable_credential_id)).data;
                commit('QUERY', { query: 'VerifiableCredential', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryVerifiableCredential', payload: { options: { all }, params: { ...key }, query } });
                return getters['getVerifiableCredential']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                console.error(new SpVuexError('QueryClient:QueryVerifiableCredential', 'API Node Unavailable. Could not perform query.'));
                return {};
            }
        },
        async QueryValidateVerifiableCredential({ commit, rootGetters, getters }, { options: { subscribe = false, all = false }, params: { ...key }, query = null }) {
            try {
                let value = query ? (await (await initQueryClient(rootGetters)).queryValidateVerifiableCredential(key.verifiable_credential_id, key.issuer_pubkey, query)).data : (await (await initQueryClient(rootGetters)).queryValidateVerifiableCredential(key.verifiable_credential_id, key.issuer_pubkey)).data;
                commit('QUERY', { query: 'ValidateVerifiableCredential', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryValidateVerifiableCredential', payload: { options: { all }, params: { ...key }, query } });
                return getters['getValidateVerifiableCredential']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                console.error(new SpVuexError('QueryClient:QueryValidateVerifiableCredential', 'API Node Unavailable. Could not perform query.'));
                return {};
            }
        },
        async sendMsgCreateVerifiableCredential({ rootGetters }, { value, fee, memo }) {
            try {
                const msg = await (await initTxClient(rootGetters)).msgCreateVerifiableCredential(value);
                const result = await (await initTxClient(rootGetters)).signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e.toString() == 'wallet is required') {
                    throw new SpVuexError('TxClient:MsgCreateVerifiableCredential:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateVerifiableCredential:Send', 'Could not broadcast Tx.');
                }
            }
        },
        async sendMsgCreateIssuerVerifiableCredential({ rootGetters }, { value, fee, memo }) {
            try {
                const msg = await (await initTxClient(rootGetters)).msgCreateIssuerVerifiableCredential(value);
                const result = await (await initTxClient(rootGetters)).signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e.toString() == 'wallet is required') {
                    throw new SpVuexError('TxClient:MsgCreateIssuerVerifiableCredential:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateIssuerVerifiableCredential:Send', 'Could not broadcast Tx.');
                }
            }
        },
        async MsgCreateVerifiableCredential({ rootGetters }, { value }) {
            try {
                const msg = await (await initTxClient(rootGetters)).msgCreateVerifiableCredential(value);
                return msg;
            }
            catch (e) {
                if (e.toString() == 'wallet is required') {
                    throw new SpVuexError('TxClient:MsgCreateVerifiableCredential:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateVerifiableCredential:Create', 'Could not create message.');
                }
            }
        },
        async MsgCreateIssuerVerifiableCredential({ rootGetters }, { value }) {
            try {
                const msg = await (await initTxClient(rootGetters)).msgCreateIssuerVerifiableCredential(value);
                return msg;
            }
            catch (e) {
                if (e.toString() == 'wallet is required') {
                    throw new SpVuexError('TxClient:MsgCreateIssuerVerifiableCredential:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateIssuerVerifiableCredential:Create', 'Could not create message.');
                }
            }
        },
    }
};
