// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgBurnToken } from "./types/issuer/tx";
import { MsgMintToken } from "./types/issuer/tx";
import { MsgCreateIssuer } from "./types/issuer/tx";
const types = [
    ["/allinbits.cosmoscash.issuer.MsgBurnToken", MsgBurnToken],
    ["/allinbits.cosmoscash.issuer.MsgMintToken", MsgMintToken],
    ["/allinbits.cosmoscash.issuer.MsgCreateIssuer", MsgCreateIssuer],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgBurnToken: (data) => ({ typeUrl: "/allinbits.cosmoscash.issuer.MsgBurnToken", value: data }),
        msgMintToken: (data) => ({ typeUrl: "/allinbits.cosmoscash.issuer.MsgMintToken", value: data }),
        msgCreateIssuer: (data) => ({ typeUrl: "/allinbits.cosmoscash.issuer.MsgCreateIssuer", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
