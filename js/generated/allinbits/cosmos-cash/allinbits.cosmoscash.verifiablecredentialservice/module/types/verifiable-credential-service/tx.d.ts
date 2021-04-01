import { Reader, Writer } from "protobufjs/minimal";
import { VerifiableCredential } from "../verifiable-credential-service/verifiable-credential";
export declare const protobufPackage = "allinbits.cosmoscash.verifiablecredentialservice";
/** MsgCreateVerifiableCredential defines a SDK message for creating a new identifer. */
export interface MsgCreateVerifiableCredential {
    verifiableCredential: VerifiableCredential | undefined;
    /** owner represents the user creating the message */
    owner: string;
}
export interface MsgCreateVerifiableCredentialResponse {
}
export interface MsgCreateIssuerVerifiableCredential {
    verifiableCredential: VerifiableCredential | undefined;
    /** owner represents the user creating the message */
    owner: string;
}
export interface MsgCreateIssuerVerifiableCredentialResponse {
}
export declare const MsgCreateVerifiableCredential: {
    encode(message: MsgCreateVerifiableCredential, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateVerifiableCredential;
    fromJSON(object: any): MsgCreateVerifiableCredential;
    toJSON(message: MsgCreateVerifiableCredential): unknown;
    fromPartial(object: DeepPartial<MsgCreateVerifiableCredential>): MsgCreateVerifiableCredential;
};
export declare const MsgCreateVerifiableCredentialResponse: {
    encode(_: MsgCreateVerifiableCredentialResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateVerifiableCredentialResponse;
    fromJSON(_: any): MsgCreateVerifiableCredentialResponse;
    toJSON(_: MsgCreateVerifiableCredentialResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateVerifiableCredentialResponse>): MsgCreateVerifiableCredentialResponse;
};
export declare const MsgCreateIssuerVerifiableCredential: {
    encode(message: MsgCreateIssuerVerifiableCredential, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateIssuerVerifiableCredential;
    fromJSON(object: any): MsgCreateIssuerVerifiableCredential;
    toJSON(message: MsgCreateIssuerVerifiableCredential): unknown;
    fromPartial(object: DeepPartial<MsgCreateIssuerVerifiableCredential>): MsgCreateIssuerVerifiableCredential;
};
export declare const MsgCreateIssuerVerifiableCredentialResponse: {
    encode(_: MsgCreateIssuerVerifiableCredentialResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateIssuerVerifiableCredentialResponse;
    fromJSON(_: any): MsgCreateIssuerVerifiableCredentialResponse;
    toJSON(_: MsgCreateIssuerVerifiableCredentialResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateIssuerVerifiableCredentialResponse>): MsgCreateIssuerVerifiableCredentialResponse;
};
/** Msg defines the identity Msg service. */
export interface Msg {
    CreateVerifiableCredential(request: MsgCreateVerifiableCredential): Promise<MsgCreateVerifiableCredentialResponse>;
    CreateIssuerVerifiableCredential(request: MsgCreateIssuerVerifiableCredential): Promise<MsgCreateIssuerVerifiableCredentialResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateVerifiableCredential(request: MsgCreateVerifiableCredential): Promise<MsgCreateVerifiableCredentialResponse>;
    CreateIssuerVerifiableCredential(request: MsgCreateIssuerVerifiableCredential): Promise<MsgCreateIssuerVerifiableCredentialResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
