import { Reader, Writer } from "protobufjs/minimal";
import { Height } from "../ibc/core/client/v1/client";
export declare const protobufPackage = "allinbits.cosmoscash.ibcidentifier";
/** MsgTransferIdentifierIBC defines a SDK message for creating a new identifer. */
export interface MsgTransferIdentifierIBC {
    id: string;
    /** the port by which the packet will be sent */
    sourcePort: string;
    /** the channel by which the packet will be sent */
    sourceChannel: string;
    /**
     * Timeout height relative to the current block height.
     * The timeout is disabled when set to 0.
     */
    timeoutHeight: Height | undefined;
    /**
     * Timeout timestamp (in nanoseconds) relative to the current block timestamp.
     * The timeout is disabled when set to 0.
     */
    timeoutTimestamp: number;
    owner: string;
}
export interface MsgTransferIdentifierIBCResponse {
}
export declare const MsgTransferIdentifierIBC: {
    encode(message: MsgTransferIdentifierIBC, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgTransferIdentifierIBC;
    fromJSON(object: any): MsgTransferIdentifierIBC;
    toJSON(message: MsgTransferIdentifierIBC): unknown;
    fromPartial(object: DeepPartial<MsgTransferIdentifierIBC>): MsgTransferIdentifierIBC;
};
export declare const MsgTransferIdentifierIBCResponse: {
    encode(_: MsgTransferIdentifierIBCResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgTransferIdentifierIBCResponse;
    fromJSON(_: any): MsgTransferIdentifierIBCResponse;
    toJSON(_: MsgTransferIdentifierIBCResponse): unknown;
    fromPartial(_: DeepPartial<MsgTransferIdentifierIBCResponse>): MsgTransferIdentifierIBCResponse;
};
/** Msg defines the identity Msg service. */
export interface Msg {
    /** TransferIdentifierIBC defines a method for transfering an identifier to another chain. */
    TransferIdentifierIBC(request: MsgTransferIdentifierIBC): Promise<MsgTransferIdentifierIBCResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    TransferIdentifierIBC(request: MsgTransferIdentifierIBC): Promise<MsgTransferIdentifierIBCResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
