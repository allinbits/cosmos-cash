import { Reader, Writer } from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { DidDocument } from "../identifier/identifier";
export declare const protobufPackage = "allinbits.cosmoscash.identifier";
/** QueryIdentifersRequest is request type for Query/Identifers RPC method. */
export interface QueryIdentifiersRequest {
    /** status enables to query for validators matching a given status. */
    status: string;
    /** pagination defines an optional pagination for the request. */
    pagination: PageRequest | undefined;
}
/** QueryIdentifersResponse is response type for the Query/Identifers RPC method */
export interface QueryIdentifiersResponse {
    /** validators contains all the queried validators. */
    didDocuments: DidDocument[];
    /** pagination defines the pagination in the response. */
    pagination: PageResponse | undefined;
}
/** QueryIdentifersRequest is request type for Query/Identifers RPC method. */
export interface QueryIdentifierRequest {
    /** status enables to query for validators matching a given status. */
    id: string;
}
/** QueryIdentifersResponse is response type for the Query/Identifers RPC method */
export interface QueryIdentifierResponse {
    /** validators contains all the queried validators. */
    didDocument: DidDocument | undefined;
}
export declare const QueryIdentifiersRequest: {
    encode(message: QueryIdentifiersRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryIdentifiersRequest;
    fromJSON(object: any): QueryIdentifiersRequest;
    toJSON(message: QueryIdentifiersRequest): unknown;
    fromPartial(object: DeepPartial<QueryIdentifiersRequest>): QueryIdentifiersRequest;
};
export declare const QueryIdentifiersResponse: {
    encode(message: QueryIdentifiersResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryIdentifiersResponse;
    fromJSON(object: any): QueryIdentifiersResponse;
    toJSON(message: QueryIdentifiersResponse): unknown;
    fromPartial(object: DeepPartial<QueryIdentifiersResponse>): QueryIdentifiersResponse;
};
export declare const QueryIdentifierRequest: {
    encode(message: QueryIdentifierRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryIdentifierRequest;
    fromJSON(object: any): QueryIdentifierRequest;
    toJSON(message: QueryIdentifierRequest): unknown;
    fromPartial(object: DeepPartial<QueryIdentifierRequest>): QueryIdentifierRequest;
};
export declare const QueryIdentifierResponse: {
    encode(message: QueryIdentifierResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryIdentifierResponse;
    fromJSON(object: any): QueryIdentifierResponse;
    toJSON(message: QueryIdentifierResponse): unknown;
    fromPartial(object: DeepPartial<QueryIdentifierResponse>): QueryIdentifierResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Identifers queries all validators that match the given status. */
    Identifiers(request: QueryIdentifiersRequest): Promise<QueryIdentifiersResponse>;
    Identifier(request: QueryIdentifierRequest): Promise<QueryIdentifierResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Identifiers(request: QueryIdentifiersRequest): Promise<QueryIdentifiersResponse>;
    Identifier(request: QueryIdentifierRequest): Promise<QueryIdentifierResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
