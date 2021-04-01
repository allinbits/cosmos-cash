/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { DidDocument } from "../identifier/identifier";

export const protobufPackage = "allinbits.cosmoscash.identifier";

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

const baseQueryIdentifiersRequest: object = { status: "" };

export const QueryIdentifiersRequest = {
  encode(
    message: QueryIdentifiersRequest,
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

  decode(input: Reader | Uint8Array, length?: number): QueryIdentifiersRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryIdentifiersRequest,
    } as QueryIdentifiersRequest;
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

  fromJSON(object: any): QueryIdentifiersRequest {
    const message = {
      ...baseQueryIdentifiersRequest,
    } as QueryIdentifiersRequest;
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

  toJSON(message: QueryIdentifiersRequest): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryIdentifiersRequest>
  ): QueryIdentifiersRequest {
    const message = {
      ...baseQueryIdentifiersRequest,
    } as QueryIdentifiersRequest;
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

const baseQueryIdentifiersResponse: object = {};

export const QueryIdentifiersResponse = {
  encode(
    message: QueryIdentifiersResponse,
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
  ): QueryIdentifiersResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryIdentifiersResponse,
    } as QueryIdentifiersResponse;
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

  fromJSON(object: any): QueryIdentifiersResponse {
    const message = {
      ...baseQueryIdentifiersResponse,
    } as QueryIdentifiersResponse;
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

  toJSON(message: QueryIdentifiersResponse): unknown {
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
    object: DeepPartial<QueryIdentifiersResponse>
  ): QueryIdentifiersResponse {
    const message = {
      ...baseQueryIdentifiersResponse,
    } as QueryIdentifiersResponse;
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

const baseQueryIdentifierRequest: object = { id: "" };

export const QueryIdentifierRequest = {
  encode(
    message: QueryIdentifierRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryIdentifierRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryIdentifierRequest } as QueryIdentifierRequest;
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

  fromJSON(object: any): QueryIdentifierRequest {
    const message = { ...baseQueryIdentifierRequest } as QueryIdentifierRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    return message;
  },

  toJSON(message: QueryIdentifierRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryIdentifierRequest>
  ): QueryIdentifierRequest {
    const message = { ...baseQueryIdentifierRequest } as QueryIdentifierRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    return message;
  },
};

const baseQueryIdentifierResponse: object = {};

export const QueryIdentifierResponse = {
  encode(
    message: QueryIdentifierResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.didDocument !== undefined) {
      DidDocument.encode(
        message.didDocument,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryIdentifierResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryIdentifierResponse,
    } as QueryIdentifierResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.didDocument = DidDocument.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryIdentifierResponse {
    const message = {
      ...baseQueryIdentifierResponse,
    } as QueryIdentifierResponse;
    if (object.didDocument !== undefined && object.didDocument !== null) {
      message.didDocument = DidDocument.fromJSON(object.didDocument);
    } else {
      message.didDocument = undefined;
    }
    return message;
  },

  toJSON(message: QueryIdentifierResponse): unknown {
    const obj: any = {};
    message.didDocument !== undefined &&
      (obj.didDocument = message.didDocument
        ? DidDocument.toJSON(message.didDocument)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryIdentifierResponse>
  ): QueryIdentifierResponse {
    const message = {
      ...baseQueryIdentifierResponse,
    } as QueryIdentifierResponse;
    if (object.didDocument !== undefined && object.didDocument !== null) {
      message.didDocument = DidDocument.fromPartial(object.didDocument);
    } else {
      message.didDocument = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Identifers queries all validators that match the given status. */
  Identifiers(
    request: QueryIdentifiersRequest
  ): Promise<QueryIdentifiersResponse>;
  Identifier(request: QueryIdentifierRequest): Promise<QueryIdentifierResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Identifiers(
    request: QueryIdentifiersRequest
  ): Promise<QueryIdentifiersResponse> {
    const data = QueryIdentifiersRequest.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.identifier.Query",
      "Identifiers",
      data
    );
    return promise.then((data) =>
      QueryIdentifiersResponse.decode(new Reader(data))
    );
  }

  Identifier(
    request: QueryIdentifierRequest
  ): Promise<QueryIdentifierResponse> {
    const data = QueryIdentifierRequest.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.identifier.Query",
      "Identifier",
      data
    );
    return promise.then((data) =>
      QueryIdentifierResponse.decode(new Reader(data))
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
