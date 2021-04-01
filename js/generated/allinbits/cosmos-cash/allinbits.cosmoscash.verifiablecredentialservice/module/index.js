// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateVerifiableCredential } from "./types/verifiable-credential-service/tx";
import { MsgCreateIssuerVerifiableCredential } from "./types/verifiable-credential-service/tx";
const types = [
    ["/allinbits.cosmoscash.verifiablecredentialservice.MsgCreateVerifiableCredential", MsgCreateVerifiableCredential],
    ["/allinbits.cosmoscash.verifiablecredentialservice.MsgCreateIssuerVerifiableCredential", MsgCreateIssuerVerifiableCredential],
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
        msgCreateVerifiableCredential: (data) => ({ typeUrl: "/allinbits.cosmoscash.verifiablecredentialservice.MsgCreateVerifiableCredential", value: data }),
        msgCreateIssuerVerifiableCredential: (data) => ({ typeUrl: "/allinbits.cosmoscash.verifiablecredentialservice.MsgCreateIssuerVerifiableCredential", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
