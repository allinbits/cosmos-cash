import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteAuthentication } from "./types/identifier/tx";
import { MsgAddService } from "./types/identifier/tx";
import { MsgCreateIdentifier } from "./types/identifier/tx";
import { MsgAddAuthentication } from "./types/identifier/tx";
import { MsgDeleteService } from "./types/identifier/tx";
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions) => Promise<import("@cosmjs/stargate").BroadcastTxResponse>;
    msgDeleteAuthentication: (data: MsgDeleteAuthentication) => EncodeObject;
    msgAddService: (data: MsgAddService) => EncodeObject;
    msgCreateIdentifier: (data: MsgCreateIdentifier) => EncodeObject;
    msgAddAuthentication: (data: MsgAddAuthentication) => EncodeObject;
    msgDeleteService: (data: MsgDeleteService) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
