import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "allinbits.cosmoscash.issuer";
/** MsgCreateIssuer defines an SDK message for creating an emoney token issuer. */
export interface MsgCreateIssuer {
    token: string;
    fee: number;
    issuerDid: string;
    licenseCredId: string;
    owner: string;
}
export interface MsgCreateIssuerResponse {
}
/** MsgBurnToken defines a SDK message for burning issuer tokens. */
export interface MsgBurnToken {
    amount: string;
    owner: string;
}
export interface MsgBurnTokenResponse {
}
/** MsgMintToken defines a SDK message for minting a token */
export interface MsgMintToken {
    amount: string;
    owner: string;
}
export interface MsgMintTokenResponse {
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
export declare const MsgBurnToken: {
    encode(message: MsgBurnToken, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgBurnToken;
    fromJSON(object: any): MsgBurnToken;
    toJSON(message: MsgBurnToken): unknown;
    fromPartial(object: DeepPartial<MsgBurnToken>): MsgBurnToken;
};
export declare const MsgBurnTokenResponse: {
    encode(_: MsgBurnTokenResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgBurnTokenResponse;
    fromJSON(_: any): MsgBurnTokenResponse;
    toJSON(_: MsgBurnTokenResponse): unknown;
    fromPartial(_: DeepPartial<MsgBurnTokenResponse>): MsgBurnTokenResponse;
};
export declare const MsgMintToken: {
    encode(message: MsgMintToken, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgMintToken;
    fromJSON(object: any): MsgMintToken;
    toJSON(message: MsgMintToken): unknown;
    fromPartial(object: DeepPartial<MsgMintToken>): MsgMintToken;
};
export declare const MsgMintTokenResponse: {
    encode(_: MsgMintTokenResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgMintTokenResponse;
    fromJSON(_: any): MsgMintTokenResponse;
    toJSON(_: MsgMintTokenResponse): unknown;
    fromPartial(_: DeepPartial<MsgMintTokenResponse>): MsgMintTokenResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateIssuer(request: MsgCreateIssuer): Promise<MsgCreateIssuerResponse>;
    BurnToken(request: MsgBurnToken): Promise<MsgBurnTokenResponse>;
    MintToken(request: MsgMintToken): Promise<MsgMintTokenResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateIssuer(request: MsgCreateIssuer): Promise<MsgCreateIssuerResponse>;
    BurnToken(request: MsgBurnToken): Promise<MsgBurnTokenResponse>;
    MintToken(request: MsgMintToken): Promise<MsgMintTokenResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
