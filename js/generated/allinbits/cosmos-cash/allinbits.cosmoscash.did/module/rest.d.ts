/**
 * DidDocument represents a dencentralised identifer.
 */
export interface DidDidDocument {
    /** @context is spec for did document. */
    context?: string[];
    /** id represents the id for the did document. */
    id?: string;
    controller?: string[];
    verificationMethod?: DidVerificationMethod[];
    service?: DidService[];
    authentication?: string[];
    assertionMethod?: string[];
    keyAgreement?: string[];
    capabilityInvocation?: string[];
    capabilityDelegation?: string[];
}
export interface DidDidMetadata {
    versionId?: string;
    /** @format date-time */
    created?: string;
    /** @format date-time */
    updated?: string;
    deactivated?: boolean;
}
export declare type DidMsgAddServiceResponse = object;
export declare type DidMsgAddVerificationResponse = object;
export declare type DidMsgCreateDidDocumentResponse = object;
export declare type DidMsgDeleteServiceResponse = object;
export declare type DidMsgRevokeVerificationResponse = object;
export declare type DidMsgSetVerificationRelationshipsResponse = object;
export declare type DidMsgUpdateDidDocumentResponse = object;
export interface DidQueryDidDocumentResponse {
    /** validators contains all the queried validators. */
    didDocument?: DidDidDocument;
    didMetadata?: DidDidMetadata;
}
export interface DidQueryDidDocumentsResponse {
    /** validators contains all the queried validators. */
    didDocuments?: DidDidDocument[];
    /** pagination defines the pagination in the response. */
    pagination?: V1Beta1PageResponse;
}
export interface DidService {
    id?: string;
    type?: string;
    serviceEndpoint?: string;
}
export interface DidVerification {
    relationships?: string[];
    method?: DidVerificationMethod;
    context?: string[];
}
export interface DidVerificationMethod {
    id?: string;
    type?: string;
    controller?: string;
    blockchainAccountID?: string;
    publicKeyHex?: string;
    publicKeyMultibase?: string;
}
export interface ProtobufAny {
    "@type"?: string;
}
export interface RpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
    /**
     * key is a value returned in PageResponse.next_key to begin
     * querying the next page most efficiently. Only one of offset or key
     * should be set.
     * @format byte
     */
    key?: string;
    /**
     * offset is a numeric offset that can be used when key is unavailable.
     * It is less efficient than using key. Only one of offset or key should
     * be set.
     * @format uint64
     */
    offset?: string;
    /**
     * limit is the total number of results to be returned in the result page.
     * If left empty it will default to a value to be set by each app.
     * @format uint64
     */
    limit?: string;
    /**
     * count_total is set to true  to indicate that the result set should include
     * a count of the total number of items available for pagination in UIs.
     * count_total is only respected when offset is used. It is ignored when key
     * is set.
     */
    countTotal?: boolean;
    /** reverse is set to true if results are to be returned in the descending order. */
    reverse?: boolean;
}
/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
    /** @format byte */
    nextKey?: string;
    /** @format uint64 */
    total?: string;
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
export interface FullRequestParams extends Omit<RequestInit, "body"> {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean;
    /** request path */
    path: string;
    /** content type of request body */
    type?: ContentType;
    /** query params */
    query?: QueryParamsType;
    /** format of response (i.e. response.json() -> format: "json") */
    format?: keyof Omit<Body, "body" | "bodyUsed">;
    /** request body */
    body?: unknown;
    /** base url */
    baseUrl?: string;
    /** request cancellation token */
    cancelToken?: CancelToken;
}
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title did/did.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryDidDocuments
     * @summary DidDocuments queries all did documents that match the given status.
     * @request GET:/allinbits/did/dids
     */
    queryDidDocuments: (query?: {
        status?: string;
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<DidQueryDidDocumentsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryDidDocument
     * @summary DidDocument queries a did documents with an id.
     * @request GET:/allinbits/did/dids/{id}
     */
    queryDidDocument: (id: string, params?: RequestParams) => Promise<HttpResponse<DidQueryDidDocumentResponse, RpcStatus>>;
}
export {};
