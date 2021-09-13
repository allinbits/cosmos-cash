import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAddService } from "./types/did/tx";
import { MsgUpdateDidDocument } from "./types/did/tx";
import { MsgAddVerification } from "./types/did/tx";
import { MsgCreateDidDocument } from "./types/did/tx";
import { MsgSetVerificationRelationships } from "./types/did/tx";
import { MsgDeleteService } from "./types/did/tx";
import { MsgRevokeVerification } from "./types/did/tx";
export declare const MissingWalletError: Error;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => Promise<import("@cosmjs/stargate").BroadcastTxResponse>;
    msgAddService: (data: MsgAddService) => EncodeObject;
    msgUpdateDidDocument: (data: MsgUpdateDidDocument) => EncodeObject;
    msgAddVerification: (data: MsgAddVerification) => EncodeObject;
    msgCreateDidDocument: (data: MsgCreateDidDocument) => EncodeObject;
    msgSetVerificationRelationships: (data: MsgSetVerificationRelationships) => EncodeObject;
    msgDeleteService: (data: MsgDeleteService) => EncodeObject;
    msgRevokeVerification: (data: MsgRevokeVerification) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
