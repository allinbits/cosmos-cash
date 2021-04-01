import { Reader, Writer } from "protobufjs/minimal";
import { Authentication, Service } from "../identifier/identifier";
export declare const protobufPackage = "allinbits.cosmoscash.identifier";
/** MsgCreateIdentifier defines a SDK message for creating a new identifer. */
export interface MsgCreateIdentifier {
    context: string;
    id: string;
    /** authentication represents public key associated with the did document. */
    authentication: Authentication[];
    /** services represents each service associated with a did */
    services: Service[];
    /** owner represents the user creating the message */
    owner: string;
}
export interface MsgCreateIdentifierResponse {
}
export interface MsgAddAuthentication {
    id: string;
    /** authentication represents public key associated with the did document. */
    authentication: Authentication | undefined;
    /** owner is the address of the user creating the message */
    owner: string;
}
export interface MsgAddAuthenticationResponse {
}
export interface MsgAddService {
    id: string;
    /** authentication represents public key associated with the did document. */
    serviceData: Service | undefined;
    /** owner is the address of the user creating the message */
    owner: string;
}
export interface MsgAddServiceResponse {
}
export interface MsgDeleteAuthentication {
    id: string;
    key: string;
    /** owner is the address of the user creating the message */
    owner: string;
}
export interface MsgDeleteAuthenticationResponse {
}
export interface MsgDeleteService {
    id: string;
    serviceId: string;
    /** owner is the address of the user creating the message */
    owner: string;
}
export interface MsgDeleteServiceResponse {
}
export declare const MsgCreateIdentifier: {
    encode(message: MsgCreateIdentifier, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateIdentifier;
    fromJSON(object: any): MsgCreateIdentifier;
    toJSON(message: MsgCreateIdentifier): unknown;
    fromPartial(object: DeepPartial<MsgCreateIdentifier>): MsgCreateIdentifier;
};
export declare const MsgCreateIdentifierResponse: {
    encode(_: MsgCreateIdentifierResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateIdentifierResponse;
    fromJSON(_: any): MsgCreateIdentifierResponse;
    toJSON(_: MsgCreateIdentifierResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateIdentifierResponse>): MsgCreateIdentifierResponse;
};
export declare const MsgAddAuthentication: {
    encode(message: MsgAddAuthentication, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddAuthentication;
    fromJSON(object: any): MsgAddAuthentication;
    toJSON(message: MsgAddAuthentication): unknown;
    fromPartial(object: DeepPartial<MsgAddAuthentication>): MsgAddAuthentication;
};
export declare const MsgAddAuthenticationResponse: {
    encode(_: MsgAddAuthenticationResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddAuthenticationResponse;
    fromJSON(_: any): MsgAddAuthenticationResponse;
    toJSON(_: MsgAddAuthenticationResponse): unknown;
    fromPartial(_: DeepPartial<MsgAddAuthenticationResponse>): MsgAddAuthenticationResponse;
};
export declare const MsgAddService: {
    encode(message: MsgAddService, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddService;
    fromJSON(object: any): MsgAddService;
    toJSON(message: MsgAddService): unknown;
    fromPartial(object: DeepPartial<MsgAddService>): MsgAddService;
};
export declare const MsgAddServiceResponse: {
    encode(_: MsgAddServiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddServiceResponse;
    fromJSON(_: any): MsgAddServiceResponse;
    toJSON(_: MsgAddServiceResponse): unknown;
    fromPartial(_: DeepPartial<MsgAddServiceResponse>): MsgAddServiceResponse;
};
export declare const MsgDeleteAuthentication: {
    encode(message: MsgDeleteAuthentication, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteAuthentication;
    fromJSON(object: any): MsgDeleteAuthentication;
    toJSON(message: MsgDeleteAuthentication): unknown;
    fromPartial(object: DeepPartial<MsgDeleteAuthentication>): MsgDeleteAuthentication;
};
export declare const MsgDeleteAuthenticationResponse: {
    encode(_: MsgDeleteAuthenticationResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteAuthenticationResponse;
    fromJSON(_: any): MsgDeleteAuthenticationResponse;
    toJSON(_: MsgDeleteAuthenticationResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteAuthenticationResponse>): MsgDeleteAuthenticationResponse;
};
export declare const MsgDeleteService: {
    encode(message: MsgDeleteService, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteService;
    fromJSON(object: any): MsgDeleteService;
    toJSON(message: MsgDeleteService): unknown;
    fromPartial(object: DeepPartial<MsgDeleteService>): MsgDeleteService;
};
export declare const MsgDeleteServiceResponse: {
    encode(_: MsgDeleteServiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteServiceResponse;
    fromJSON(_: any): MsgDeleteServiceResponse;
    toJSON(_: MsgDeleteServiceResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteServiceResponse>): MsgDeleteServiceResponse;
};
/** Msg defines the identity Msg service. */
export interface Msg {
    /** CreateDidDocument defines a method for creating a new identity. */
    CreateIdentifier(request: MsgCreateIdentifier): Promise<MsgCreateIdentifierResponse>;
    AddAuthentication(request: MsgAddAuthentication): Promise<MsgAddAuthenticationResponse>;
    AddService(request: MsgAddService): Promise<MsgAddServiceResponse>;
    DeleteAuthentication(request: MsgDeleteAuthentication): Promise<MsgDeleteAuthenticationResponse>;
    DeleteService(request: MsgDeleteService): Promise<MsgDeleteServiceResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateIdentifier(request: MsgCreateIdentifier): Promise<MsgCreateIdentifierResponse>;
    AddAuthentication(request: MsgAddAuthentication): Promise<MsgAddAuthenticationResponse>;
    AddService(request: MsgAddService): Promise<MsgAddServiceResponse>;
    DeleteAuthentication(request: MsgDeleteAuthentication): Promise<MsgDeleteAuthenticationResponse>;
    DeleteService(request: MsgDeleteService): Promise<MsgDeleteServiceResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
