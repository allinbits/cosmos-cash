// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateVerifiableCredential } from "./types/verifiable-credential/tx";
import { MsgDeleteVerifiableCredential } from "./types/verifiable-credential/tx";


const types = [
  ["/allinbits.cosmoscash.verifiablecredential.MsgCreateVerifiableCredential", MsgCreateVerifiableCredential],
  ["/allinbits.cosmoscash.verifiablecredential.MsgDeleteVerifiableCredential", MsgDeleteVerifiableCredential],
  
];
export const MissingWalletError = new Error("wallet is required");

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreateVerifiableCredential: (data: MsgCreateVerifiableCredential): EncodeObject => ({ typeUrl: "/allinbits.cosmoscash.verifiablecredential.MsgCreateVerifiableCredential", value: data }),
    msgDeleteVerifiableCredential: (data: MsgDeleteVerifiableCredential): EncodeObject => ({ typeUrl: "/allinbits.cosmoscash.verifiablecredential.MsgDeleteVerifiableCredential", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
