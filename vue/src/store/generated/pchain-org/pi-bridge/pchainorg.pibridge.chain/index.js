import { txClient, queryClient, MissingWalletError } from './module';
// @ts-ignore
import { SpVuexError } from '@starport/vuex';
import { Chain } from "./module/types/chain/chain";
export { Chain };
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
function mergeResults(value, next_values) {
    for (let prop of Object.keys(next_values)) {
        if (Array.isArray(next_values[prop])) {
            value[prop] = [...value[prop], ...next_values[prop]];
        }
        else {
            value[prop] = next_values[prop];
        }
    }
    return value;
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
        Chain: {},
        ChainAll: {},
        _Structure: {
            Chain: getStructure(Chain.fromPartial({})),
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
        getChain: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.Chain[JSON.stringify(params)] ?? {};
        },
        getChainAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ChainAll[JSON.stringify(params)] ?? {};
        },
        getTypeStructure: (state) => (type) => {
            return state._Structure[type].fields;
        }
    },
    actions: {
        init({ dispatch, rootGetters }) {
            console.log('Vuex module: pchainorg.pibridge.chain initialized!');
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
            state._Subscriptions.forEach(async (subscription) => {
                try {
                    await dispatch(subscription.action, subscription.payload);
                }
                catch (e) {
                    throw new SpVuexError('Subscriptions: ' + e.message);
                }
            });
        },
        async QueryChain({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryChain(key.index)).data;
                commit('QUERY', { query: 'Chain', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryChain', payload: { options: { all }, params: { ...key }, query } });
                return getters['getChain']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryChain', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryChainAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryChainAll(query)).data;
                while (all && value.pagination && value.pagination.nextKey != null) {
                    let next_values = (await queryClient.queryChainAll({ ...query, 'pagination.key': value.pagination.nextKey })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'ChainAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryChainAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getChainAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryChainAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async sendMsgCreateChain({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateChain(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgCreateChain:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateChain:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgUpdateChain({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgUpdateChain(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgUpdateChain:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgUpdateChain:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgDeleteChain({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgDeleteChain(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgDeleteChain:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgDeleteChain:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async MsgCreateChain({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateChain(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgCreateChain:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateChain:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgUpdateChain({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgUpdateChain(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgUpdateChain:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgUpdateChain:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgDeleteChain({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgDeleteChain(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgDeleteChain:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgDeleteChain:Create', 'Could not create message: ' + e.message);
                }
            }
        },
    }
};