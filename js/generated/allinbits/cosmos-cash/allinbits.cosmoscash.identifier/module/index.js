// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteAuthentication } from "./types/identifier/tx";
import { MsgAddService } from "./types/identifier/tx";
import { MsgCreateIdentifier } from "./types/identifier/tx";
import { MsgAddAuthentication } from "./types/identifier/tx";
import { MsgDeleteService } from "./types/identifier/tx";
const types = [
    ["/allinbits.cosmoscash.identifier.MsgDeleteAuthentication", MsgDeleteAuthentication],
    ["/allinbits.cosmoscash.identifier.MsgAddService", MsgAddService],
    ["/allinbits.cosmoscash.identifier.MsgCreateIdentifier", MsgCreateIdentifier],
    ["/allinbits.cosmoscash.identifier.MsgAddAuthentication", MsgAddAuthentication],
    ["/allinbits.cosmoscash.identifier.MsgDeleteService", MsgDeleteService],
];
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw new Error("wallet is required");
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee = defaultFee, memo = null }) => memo ? client.signAndBroadcast(address, msgs, fee, memo) : client.signAndBroadcast(address, msgs, fee),
        msgDeleteAuthentication: (data) => ({ typeUrl: "/allinbits.cosmoscash.identifier.MsgDeleteAuthentication", value: data }),
        msgAddService: (data) => ({ typeUrl: "/allinbits.cosmoscash.identifier.MsgAddService", value: data }),
        msgCreateIdentifier: (data) => ({ typeUrl: "/allinbits.cosmoscash.identifier.MsgCreateIdentifier", value: data }),
        msgAddAuthentication: (data) => ({ typeUrl: "/allinbits.cosmoscash.identifier.MsgAddAuthentication", value: data }),
        msgDeleteService: (data) => ({ typeUrl: "/allinbits.cosmoscash.identifier.MsgDeleteService", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
