/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { DidDocument, DidMetadata } from "../did/did";

export const protobufPackage = "allinbits.cosmoscash.did";

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

const baseQueryDidDocumentsRequest: object = { status: "" };

export const QueryDidDocumentsRequest = {
  encode(
    message: QueryDidDocumentsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.status !== "") {
      writer.uint32(10).string(message.status);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDidDocumentsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDidDocumentsRequest,
    } as QueryDidDocumentsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.string();
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDidDocumentsRequest {
    const message = {
      ...baseQueryDidDocumentsRequest,
    } as QueryDidDocumentsRequest;
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryDidDocumentsRequest): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDidDocumentsRequest>
  ): QueryDidDocumentsRequest {
    const message = {
      ...baseQueryDidDocumentsRequest,
    } as QueryDidDocumentsRequest;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryDidDocumentsResponse: object = {};

export const QueryDidDocumentsResponse = {
  encode(
    message: QueryDidDocumentsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.didDocuments) {
      DidDocument.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDidDocumentsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDidDocumentsResponse,
    } as QueryDidDocumentsResponse;
    message.didDocuments = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.didDocuments.push(
            DidDocument.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDidDocumentsResponse {
    const message = {
      ...baseQueryDidDocumentsResponse,
    } as QueryDidDocumentsResponse;
    message.didDocuments = [];
    if (object.didDocuments !== undefined && object.didDocuments !== null) {
      for (const e of object.didDocuments) {
        message.didDocuments.push(DidDocument.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryDidDocumentsResponse): unknown {
    const obj: any = {};
    if (message.didDocuments) {
      obj.didDocuments = message.didDocuments.map((e) =>
        e ? DidDocument.toJSON(e) : undefined
      );
    } else {
      obj.didDocuments = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDidDocumentsResponse>
  ): QueryDidDocumentsResponse {
    const message = {
      ...baseQueryDidDocumentsResponse,
    } as QueryDidDocumentsResponse;
    message.didDocuments = [];
    if (object.didDocuments !== undefined && object.didDocuments !== null) {
      for (const e of object.didDocuments) {
        message.didDocuments.push(DidDocument.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryDidDocumentRequest: object = { id: "" };

export const QueryDidDocumentRequest = {
  encode(
    message: QueryDidDocumentRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryDidDocumentRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDidDocumentRequest,
    } as QueryDidDocumentRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDidDocumentRequest {
    const message = {
      ...baseQueryDidDocumentRequest,
    } as QueryDidDocumentRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    return message;
  },

  toJSON(message: QueryDidDocumentRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDidDocumentRequest>
  ): QueryDidDocumentRequest {
    const message = {
      ...baseQueryDidDocumentRequest,
    } as QueryDidDocumentRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    return message;
  },
};

const baseQueryDidDocumentResponse: object = {};

export const QueryDidDocumentResponse = {
  encode(
    message: QueryDidDocumentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.didDocument !== undefined) {
      DidDocument.encode(
        message.didDocument,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.didMetadata !== undefined) {
      DidMetadata.encode(
        message.didMetadata,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDidDocumentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDidDocumentResponse,
    } as QueryDidDocumentResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.didDocument = DidDocument.decode(reader, reader.uint32());
          break;
        case 2:
          message.didMetadata = DidMetadata.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDidDocumentResponse {
    const message = {
      ...baseQueryDidDocumentResponse,
    } as QueryDidDocumentResponse;
    if (object.didDocument !== undefined && object.didDocument !== null) {
      message.didDocument = DidDocument.fromJSON(object.didDocument);
    } else {
      message.didDocument = undefined;
    }
    if (object.didMetadata !== undefined && object.didMetadata !== null) {
      message.didMetadata = DidMetadata.fromJSON(object.didMetadata);
    } else {
      message.didMetadata = undefined;
    }
    return message;
  },

  toJSON(message: QueryDidDocumentResponse): unknown {
    const obj: any = {};
    message.didDocument !== undefined &&
      (obj.didDocument = message.didDocument
        ? DidDocument.toJSON(message.didDocument)
        : undefined);
    message.didMetadata !== undefined &&
      (obj.didMetadata = message.didMetadata
        ? DidMetadata.toJSON(message.didMetadata)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDidDocumentResponse>
  ): QueryDidDocumentResponse {
    const message = {
      ...baseQueryDidDocumentResponse,
    } as QueryDidDocumentResponse;
    if (object.didDocument !== undefined && object.didDocument !== null) {
      message.didDocument = DidDocument.fromPartial(object.didDocument);
    } else {
      message.didDocument = undefined;
    }
    if (object.didMetadata !== undefined && object.didMetadata !== null) {
      message.didMetadata = DidMetadata.fromPartial(object.didMetadata);
    } else {
      message.didMetadata = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** DidDocuments queries all did documents that match the given status. */
  DidDocuments(
    request: QueryDidDocumentsRequest
  ): Promise<QueryDidDocumentsResponse>;
  /** DidDocument queries a did documents with an id. */
  DidDocument(
    request: QueryDidDocumentRequest
  ): Promise<QueryDidDocumentResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  DidDocuments(
    request: QueryDidDocumentsRequest
  ): Promise<QueryDidDocumentsResponse> {
    const data = QueryDidDocumentsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.did.Query",
      "DidDocuments",
      data
    );
    return promise.then((data) =>
      QueryDidDocumentsResponse.decode(new Reader(data))
    );
  }

  DidDocument(
    request: QueryDidDocumentRequest
  ): Promise<QueryDidDocumentResponse> {
    const data = QueryDidDocumentRequest.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.did.Query",
      "DidDocument",
      data
    );
    return promise.then((data) =>
      QueryDidDocumentResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
