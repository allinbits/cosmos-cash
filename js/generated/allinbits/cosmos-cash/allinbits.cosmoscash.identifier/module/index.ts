// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
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
  if (!wallet) throw new Error("wallet is required");

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee=defaultFee, memo=null }: SignAndBroadcastOptions) => memo?client.signAndBroadcast(address, msgs, fee,memo):client.signAndBroadcast(address, msgs, fee),
    msgDeleteAuthentication: (data: MsgDeleteAuthentication): EncodeObject => ({ typeUrl: "/allinbits.cosmoscash.identifier.MsgDeleteAuthentication", value: data }),
    msgAddService: (data: MsgAddService): EncodeObject => ({ typeUrl: "/allinbits.cosmoscash.identifier.MsgAddService", value: data }),
    msgCreateIdentifier: (data: MsgCreateIdentifier): EncodeObject => ({ typeUrl: "/allinbits.cosmoscash.identifier.MsgCreateIdentifier", value: data }),
    msgAddAuthentication: (data: MsgAddAuthentication): EncodeObject => ({ typeUrl: "/allinbits.cosmoscash.identifier.MsgAddAuthentication", value: data }),
    msgDeleteService: (data: MsgDeleteService): EncodeObject => ({ typeUrl: "/allinbits.cosmoscash.identifier.MsgDeleteService", value: data }),
    
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
