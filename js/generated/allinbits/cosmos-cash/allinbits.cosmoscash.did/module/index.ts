// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAddService } from "./types/did/tx";
import { MsgUpdateDidDocument } from "./types/did/tx";
import { MsgAddVerification } from "./types/did/tx";
import { MsgCreateDidDocument } from "./types/did/tx";
import { MsgSetVerificationRelationships } from "./types/did/tx";
import { MsgDeleteService } from "./types/did/tx";
import { MsgRevokeVerification } from "./types/did/tx";


const types = [
  ["/allinbits.cosmoscash.did.MsgAddService", MsgAddService],
  ["/allinbits.cosmoscash.did.MsgUpdateDidDocument", MsgUpdateDidDocument],
  ["/allinbits.cosmoscash.did.MsgAddVerification", MsgAddVerification],
  ["/allinbits.cosmoscash.did.MsgCreateDidDocument", MsgCreateDidDocument],
  ["/allinbits.cosmoscash.did.MsgSetVerificationRelationships", MsgSetVerificationRelationships],
  ["/allinbits.cosmoscash.did.MsgDeleteService", MsgDeleteService],
  ["/allinbits.cosmoscash.did.MsgRevokeVerification", MsgRevokeVerification],
  
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
    msgAddService: (data: MsgAddService): EncodeObject => ({ typeUrl: "/allinbits.cosmoscash.did.MsgAddService", value: data }),
    msgUpdateDidDocument: (data: MsgUpdateDidDocument): EncodeObject => ({ typeUrl: "/allinbits.cosmoscash.did.MsgUpdateDidDocument", value: data }),
    msgAddVerification: (data: MsgAddVerification): EncodeObject => ({ typeUrl: "/allinbits.cosmoscash.did.MsgAddVerification", value: data }),
    msgCreateDidDocument: (data: MsgCreateDidDocument): EncodeObject => ({ typeUrl: "/allinbits.cosmoscash.did.MsgCreateDidDocument", value: data }),
    msgSetVerificationRelationships: (data: MsgSetVerificationRelationships): EncodeObject => ({ typeUrl: "/allinbits.cosmoscash.did.MsgSetVerificationRelationships", value: data }),
    msgDeleteService: (data: MsgDeleteService): EncodeObject => ({ typeUrl: "/allinbits.cosmoscash.did.MsgDeleteService", value: data }),
    msgRevokeVerification: (data: MsgRevokeVerification): EncodeObject => ({ typeUrl: "/allinbits.cosmoscash.did.MsgRevokeVerification", value: data }),
    
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
