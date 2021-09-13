import { Reader, Writer } from "protobufjs/minimal";
import { VerifiableCredential } from "../verifiable-credential/verifiable-credential";
export declare const protobufPackage = "allinbits.cosmoscash.verifiablecredential";
/** MsgCreateVerifiableCredential defines a SDK message for creating a new identifer. */
export interface MsgCreateVerifiableCredential {
    verifiableCredential: VerifiableCredential | undefined;
    /** owner represents the user creating the message */
    owner: string;
}
export interface MsgCreateVerifiableCredentialResponse {
}
/** MsgDeleteVerifiableCredential defines a SDK message for updating a credential. */
export interface MsgDeleteVerifiableCredential {
    verifiableCredentialId: string;
    issuerDid: string;
    /** owner represents the user creating the message */
    owner: string;
}
export interface MsgDeleteVerifiableCredentialResponse {
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
export declare const MsgDeleteVerifiableCredential: {
    encode(message: MsgDeleteVerifiableCredential, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteVerifiableCredential;
    fromJSON(object: any): MsgDeleteVerifiableCredential;
    toJSON(message: MsgDeleteVerifiableCredential): unknown;
    fromPartial(object: DeepPartial<MsgDeleteVerifiableCredential>): MsgDeleteVerifiableCredential;
};
export declare const MsgDeleteVerifiableCredentialResponse: {
    encode(_: MsgDeleteVerifiableCredentialResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteVerifiableCredentialResponse;
    fromJSON(_: any): MsgDeleteVerifiableCredentialResponse;
    toJSON(_: MsgDeleteVerifiableCredentialResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteVerifiableCredentialResponse>): MsgDeleteVerifiableCredentialResponse;
};
/** Msg defines the identity Msg service. */
export interface Msg {
    CreateVerifiableCredential(request: MsgCreateVerifiableCredential): Promise<MsgCreateVerifiableCredentialResponse>;
    DeleteVerifiableCredential(request: MsgDeleteVerifiableCredential): Promise<MsgDeleteVerifiableCredentialResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateVerifiableCredential(request: MsgCreateVerifiableCredential): Promise<MsgCreateVerifiableCredentialResponse>;
    DeleteVerifiableCredential(request: MsgDeleteVerifiableCredential): Promise<MsgDeleteVerifiableCredentialResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
