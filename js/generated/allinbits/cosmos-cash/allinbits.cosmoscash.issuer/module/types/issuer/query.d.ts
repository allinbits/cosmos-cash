import { Reader, Writer } from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Issuer } from "../issuer/issuer";
export declare const protobufPackage = "allinbits.cosmoscash.issuer";
export interface QueryIssuersRequest {
    /** status enables to query for validators matching a given status. */
    status: string;
    /** pagination defines an optional pagination for the request. */
    pagination: PageRequest | undefined;
}
/** QueryIdentifersResponse is response type for the Query/Identifers RPC method */
export interface QueryIssuersResponse {
    /** validators contains all the queried validators. */
    issuers: Issuer[];
    /** pagination defines the pagination in the response. */
    pagination: PageResponse | undefined;
}
export declare const QueryIssuersRequest: {
    encode(message: QueryIssuersRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryIssuersRequest;
    fromJSON(object: any): QueryIssuersRequest;
    toJSON(message: QueryIssuersRequest): unknown;
    fromPartial(object: DeepPartial<QueryIssuersRequest>): QueryIssuersRequest;
};
export declare const QueryIssuersResponse: {
    encode(message: QueryIssuersResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryIssuersResponse;
    fromJSON(object: any): QueryIssuersResponse;
    toJSON(message: QueryIssuersResponse): unknown;
    fromPartial(object: DeepPartial<QueryIssuersResponse>): QueryIssuersResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    Issuers(request: QueryIssuersRequest): Promise<QueryIssuersResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Issuers(request: QueryIssuersRequest): Promise<QueryIssuersResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
