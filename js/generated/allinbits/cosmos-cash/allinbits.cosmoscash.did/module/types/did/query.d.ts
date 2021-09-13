import { Reader, Writer } from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { DidDocument, DidMetadata } from "../did/did";
export declare const protobufPackage = "allinbits.cosmoscash.did";
/** QueryDidDocumentsRequest is request type for Query/DidDocuments RPC method. */
export interface QueryDidDocumentsRequest {
    /** status enables to query for validators matching a given status. */
    status: string;
    /** pagination defines an optional pagination for the request. */
    pagination: PageRequest | undefined;
}
/** QueryDidDocumentsResponse is response type for the Query/DidDocuments RPC method */
export interface QueryDidDocumentsResponse {
    /** validators contains all the queried validators. */
    didDocuments: DidDocument[];
    /** pagination defines the pagination in the response. */
    pagination: PageResponse | undefined;
}
/** QueryDidDocumentsRequest is request type for Query/DidDocuments RPC method. */
export interface QueryDidDocumentRequest {
    /** status enables to query for validators matching a given status. */
    id: string;
}
/** QueryDidDocumentsResponse is response type for the Query/DidDocuments RPC method */
export interface QueryDidDocumentResponse {
    /** validators contains all the queried validators. */
    didDocument: DidDocument | undefined;
    didMetadata: DidMetadata | undefined;
}
export declare const QueryDidDocumentsRequest: {
    encode(message: QueryDidDocumentsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryDidDocumentsRequest;
    fromJSON(object: any): QueryDidDocumentsRequest;
    toJSON(message: QueryDidDocumentsRequest): unknown;
    fromPartial(object: DeepPartial<QueryDidDocumentsRequest>): QueryDidDocumentsRequest;
};
export declare const QueryDidDocumentsResponse: {
    encode(message: QueryDidDocumentsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryDidDocumentsResponse;
    fromJSON(object: any): QueryDidDocumentsResponse;
    toJSON(message: QueryDidDocumentsResponse): unknown;
    fromPartial(object: DeepPartial<QueryDidDocumentsResponse>): QueryDidDocumentsResponse;
};
export declare const QueryDidDocumentRequest: {
    encode(message: QueryDidDocumentRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryDidDocumentRequest;
    fromJSON(object: any): QueryDidDocumentRequest;
    toJSON(message: QueryDidDocumentRequest): unknown;
    fromPartial(object: DeepPartial<QueryDidDocumentRequest>): QueryDidDocumentRequest;
};
export declare const QueryDidDocumentResponse: {
    encode(message: QueryDidDocumentResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryDidDocumentResponse;
    fromJSON(object: any): QueryDidDocumentResponse;
    toJSON(message: QueryDidDocumentResponse): unknown;
    fromPartial(object: DeepPartial<QueryDidDocumentResponse>): QueryDidDocumentResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** DidDocuments queries all did documents that match the given status. */
    DidDocuments(request: QueryDidDocumentsRequest): Promise<QueryDidDocumentsResponse>;
    /** DidDocument queries a did documents with an id. */
    DidDocument(request: QueryDidDocumentRequest): Promise<QueryDidDocumentResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    DidDocuments(request: QueryDidDocumentsRequest): Promise<QueryDidDocumentsResponse>;
    DidDocument(request: QueryDidDocumentRequest): Promise<QueryDidDocumentResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
