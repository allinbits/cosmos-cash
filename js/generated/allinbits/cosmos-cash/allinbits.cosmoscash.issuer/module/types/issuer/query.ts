/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Issuer } from "../issuer/issuer";

export const protobufPackage = "allinbits.cosmoscash.issuer";

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

const baseQueryIssuersRequest: object = { status: "" };

export const QueryIssuersRequest = {
  encode(
    message: QueryIssuersRequest,
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

  decode(input: Reader | Uint8Array, length?: number): QueryIssuersRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryIssuersRequest } as QueryIssuersRequest;
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

  fromJSON(object: any): QueryIssuersRequest {
    const message = { ...baseQueryIssuersRequest } as QueryIssuersRequest;
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

  toJSON(message: QueryIssuersRequest): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryIssuersRequest>): QueryIssuersRequest {
    const message = { ...baseQueryIssuersRequest } as QueryIssuersRequest;
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

const baseQueryIssuersResponse: object = {};

export const QueryIssuersResponse = {
  encode(
    message: QueryIssuersResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.issuers) {
      Issuer.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryIssuersResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryIssuersResponse } as QueryIssuersResponse;
    message.issuers = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.issuers.push(Issuer.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryIssuersResponse {
    const message = { ...baseQueryIssuersResponse } as QueryIssuersResponse;
    message.issuers = [];
    if (object.issuers !== undefined && object.issuers !== null) {
      for (const e of object.issuers) {
        message.issuers.push(Issuer.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryIssuersResponse): unknown {
    const obj: any = {};
    if (message.issuers) {
      obj.issuers = message.issuers.map((e) =>
        e ? Issuer.toJSON(e) : undefined
      );
    } else {
      obj.issuers = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryIssuersResponse>): QueryIssuersResponse {
    const message = { ...baseQueryIssuersResponse } as QueryIssuersResponse;
    message.issuers = [];
    if (object.issuers !== undefined && object.issuers !== null) {
      for (const e of object.issuers) {
        message.issuers.push(Issuer.fromPartial(e));
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

/** Query defines the gRPC querier service. */
export interface Query {
  Issuers(request: QueryIssuersRequest): Promise<QueryIssuersResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Issuers(request: QueryIssuersRequest): Promise<QueryIssuersResponse> {
    const data = QueryIssuersRequest.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.issuer.Query",
      "Issuers",
      data
    );
    return promise.then((data) =>
      QueryIssuersResponse.decode(new Reader(data))
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
