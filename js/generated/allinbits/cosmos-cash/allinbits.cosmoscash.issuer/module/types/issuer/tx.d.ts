import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "allinbits.cosmoscash.issuer";
/** MsgCreateIssuer defines a SDK message for creating a new identifer. */
export interface MsgCreateIssuer {
    token: string;
    fee: number;
    owner: string;
}
export interface MsgCreateIssuerResponse {
}
export declare const MsgCreateIssuer: {
    encode(message: MsgCreateIssuer, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateIssuer;
    fromJSON(object: any): MsgCreateIssuer;
    toJSON(message: MsgCreateIssuer): unknown;
    fromPartial(object: DeepPartial<MsgCreateIssuer>): MsgCreateIssuer;
};
export declare const MsgCreateIssuerResponse: {
    encode(_: MsgCreateIssuerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateIssuerResponse;
    fromJSON(_: any): MsgCreateIssuerResponse;
    toJSON(_: MsgCreateIssuerResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateIssuerResponse>): MsgCreateIssuerResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateIssuer(request: MsgCreateIssuer): Promise<MsgCreateIssuerResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateIssuer(request: MsgCreateIssuer): Promise<MsgCreateIssuerResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
