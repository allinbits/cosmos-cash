import { Reader, Writer } from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { VerifiableCredential } from "../verifiable-credential/verifiable-credential";
export declare const protobufPackage = "allinbits.cosmoscash.verifiablecredential";
/** QueryVerifiableCredentialsRequest is request type for Query/VerifiableCredentials RPC method. */
export interface QueryVerifiableCredentialsRequest {
    /** status enables to query for credentials matching a given status. */
    status: string;
    /** pagination defines an optional pagination for the request. */
    pagination: PageRequest | undefined;
}
/** QueryVerifiableCredentialsResponse is response type for the Query/Identifers RPC method */
export interface QueryVerifiableCredentialsResponse {
    /** validators contains all the queried validators. */
    vcs: VerifiableCredential[];
    /** pagination defines the pagination in the response. */
    pagination: PageResponse | undefined;
}
/** QueryVerifiableCredentialRequest is response type for the Query/VerifiableCredential RPC method */
export interface QueryVerifiableCredentialRequest {
    /** verifiable_credential_id defines the credential id to query for. */
    verifiableCredentialId: string;
}
/** QueryVerifiableCredentialResponse is response type for the Query/VerifiableCredential RPC method */
export interface QueryVerifiableCredentialResponse {
    /** verifiable_credential defines the the credential info. */
    verifiableCredential: VerifiableCredential | undefined;
}
/** QueryVerifiableCredentialResponse is response type for the Query/VerifiableCredential RPC method */
export interface QueryValidateVerifiableCredentialResponse {
    /** is_valid defines if the credential is signed by the correct public key. */
    isValid: boolean;
}
export declare const QueryVerifiableCredentialsRequest: {
    encode(message: QueryVerifiableCredentialsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryVerifiableCredentialsRequest;
    fromJSON(object: any): QueryVerifiableCredentialsRequest;
    toJSON(message: QueryVerifiableCredentialsRequest): unknown;
    fromPartial(object: DeepPartial<QueryVerifiableCredentialsRequest>): QueryVerifiableCredentialsRequest;
};
export declare const QueryVerifiableCredentialsResponse: {
    encode(message: QueryVerifiableCredentialsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryVerifiableCredentialsResponse;
    fromJSON(object: any): QueryVerifiableCredentialsResponse;
    toJSON(message: QueryVerifiableCredentialsResponse): unknown;
    fromPartial(object: DeepPartial<QueryVerifiableCredentialsResponse>): QueryVerifiableCredentialsResponse;
};
export declare const QueryVerifiableCredentialRequest: {
    encode(message: QueryVerifiableCredentialRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryVerifiableCredentialRequest;
    fromJSON(object: any): QueryVerifiableCredentialRequest;
    toJSON(message: QueryVerifiableCredentialRequest): unknown;
    fromPartial(object: DeepPartial<QueryVerifiableCredentialRequest>): QueryVerifiableCredentialRequest;
};
export declare const QueryVerifiableCredentialResponse: {
    encode(message: QueryVerifiableCredentialResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryVerifiableCredentialResponse;
    fromJSON(object: any): QueryVerifiableCredentialResponse;
    toJSON(message: QueryVerifiableCredentialResponse): unknown;
    fromPartial(object: DeepPartial<QueryVerifiableCredentialResponse>): QueryVerifiableCredentialResponse;
};
export declare const QueryValidateVerifiableCredentialResponse: {
    encode(message: QueryValidateVerifiableCredentialResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryValidateVerifiableCredentialResponse;
    fromJSON(object: any): QueryValidateVerifiableCredentialResponse;
    toJSON(message: QueryValidateVerifiableCredentialResponse): unknown;
    fromPartial(object: DeepPartial<QueryValidateVerifiableCredentialResponse>): QueryValidateVerifiableCredentialResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Identifers queries all validators that match the given status. */
    VerifiableCredentials(request: QueryVerifiableCredentialsRequest): Promise<QueryVerifiableCredentialsResponse>;
    /** VerifiableCredential queries validator info for given validator address. */
    VerifiableCredential(request: QueryVerifiableCredentialRequest): Promise<QueryVerifiableCredentialResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    VerifiableCredentials(request: QueryVerifiableCredentialsRequest): Promise<QueryVerifiableCredentialsResponse>;
    VerifiableCredential(request: QueryVerifiableCredentialRequest): Promise<QueryVerifiableCredentialResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
