/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { VerifiableCredential } from "../verifiable-credential-service/verifiable-credential";

export const protobufPackage =
  "allinbits.cosmoscash.verifiablecredentialservice";

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

/** QueryValidateVerifiableCredentialRequest is response type for the Query/VerifiableCredential RPC method */
export interface QueryValidateVerifiableCredentialRequest {
  /** verifiable_credential_id defines the credential id to query for. */
  verifiableCredentialId: string;
  /** issuer_pubkey is used to validate to verifiable_credential */
  issuerPubkey: string;
}

/** QueryVerifiableCredentialResponse is response type for the Query/VerifiableCredential RPC method */
export interface QueryValidateVerifiableCredentialResponse {
  /** is_valid defines if the credential is signed by the correct public key. */
  isValid: boolean;
}

const baseQueryVerifiableCredentialsRequest: object = { status: "" };

export const QueryVerifiableCredentialsRequest = {
  encode(
    message: QueryVerifiableCredentialsRequest,
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
  ): QueryVerifiableCredentialsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryVerifiableCredentialsRequest,
    } as QueryVerifiableCredentialsRequest;
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

  fromJSON(object: any): QueryVerifiableCredentialsRequest {
    const message = {
      ...baseQueryVerifiableCredentialsRequest,
    } as QueryVerifiableCredentialsRequest;
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

  toJSON(message: QueryVerifiableCredentialsRequest): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVerifiableCredentialsRequest>
  ): QueryVerifiableCredentialsRequest {
    const message = {
      ...baseQueryVerifiableCredentialsRequest,
    } as QueryVerifiableCredentialsRequest;
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

const baseQueryVerifiableCredentialsResponse: object = {};

export const QueryVerifiableCredentialsResponse = {
  encode(
    message: QueryVerifiableCredentialsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.vcs) {
      VerifiableCredential.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryVerifiableCredentialsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryVerifiableCredentialsResponse,
    } as QueryVerifiableCredentialsResponse;
    message.vcs = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vcs.push(
            VerifiableCredential.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryVerifiableCredentialsResponse {
    const message = {
      ...baseQueryVerifiableCredentialsResponse,
    } as QueryVerifiableCredentialsResponse;
    message.vcs = [];
    if (object.vcs !== undefined && object.vcs !== null) {
      for (const e of object.vcs) {
        message.vcs.push(VerifiableCredential.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryVerifiableCredentialsResponse): unknown {
    const obj: any = {};
    if (message.vcs) {
      obj.vcs = message.vcs.map((e) =>
        e ? VerifiableCredential.toJSON(e) : undefined
      );
    } else {
      obj.vcs = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVerifiableCredentialsResponse>
  ): QueryVerifiableCredentialsResponse {
    const message = {
      ...baseQueryVerifiableCredentialsResponse,
    } as QueryVerifiableCredentialsResponse;
    message.vcs = [];
    if (object.vcs !== undefined && object.vcs !== null) {
      for (const e of object.vcs) {
        message.vcs.push(VerifiableCredential.fromPartial(e));
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

const baseQueryVerifiableCredentialRequest: object = {
  verifiableCredentialId: "",
};

export const QueryVerifiableCredentialRequest = {
  encode(
    message: QueryVerifiableCredentialRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.verifiableCredentialId !== "") {
      writer.uint32(10).string(message.verifiableCredentialId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryVerifiableCredentialRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryVerifiableCredentialRequest,
    } as QueryVerifiableCredentialRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.verifiableCredentialId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryVerifiableCredentialRequest {
    const message = {
      ...baseQueryVerifiableCredentialRequest,
    } as QueryVerifiableCredentialRequest;
    if (
      object.verifiableCredentialId !== undefined &&
      object.verifiableCredentialId !== null
    ) {
      message.verifiableCredentialId = String(object.verifiableCredentialId);
    } else {
      message.verifiableCredentialId = "";
    }
    return message;
  },

  toJSON(message: QueryVerifiableCredentialRequest): unknown {
    const obj: any = {};
    message.verifiableCredentialId !== undefined &&
      (obj.verifiableCredentialId = message.verifiableCredentialId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVerifiableCredentialRequest>
  ): QueryVerifiableCredentialRequest {
    const message = {
      ...baseQueryVerifiableCredentialRequest,
    } as QueryVerifiableCredentialRequest;
    if (
      object.verifiableCredentialId !== undefined &&
      object.verifiableCredentialId !== null
    ) {
      message.verifiableCredentialId = object.verifiableCredentialId;
    } else {
      message.verifiableCredentialId = "";
    }
    return message;
  },
};

const baseQueryVerifiableCredentialResponse: object = {};

export const QueryVerifiableCredentialResponse = {
  encode(
    message: QueryVerifiableCredentialResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.verifiableCredential !== undefined) {
      VerifiableCredential.encode(
        message.verifiableCredential,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryVerifiableCredentialResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryVerifiableCredentialResponse,
    } as QueryVerifiableCredentialResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.verifiableCredential = VerifiableCredential.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryVerifiableCredentialResponse {
    const message = {
      ...baseQueryVerifiableCredentialResponse,
    } as QueryVerifiableCredentialResponse;
    if (
      object.verifiableCredential !== undefined &&
      object.verifiableCredential !== null
    ) {
      message.verifiableCredential = VerifiableCredential.fromJSON(
        object.verifiableCredential
      );
    } else {
      message.verifiableCredential = undefined;
    }
    return message;
  },

  toJSON(message: QueryVerifiableCredentialResponse): unknown {
    const obj: any = {};
    message.verifiableCredential !== undefined &&
      (obj.verifiableCredential = message.verifiableCredential
        ? VerifiableCredential.toJSON(message.verifiableCredential)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVerifiableCredentialResponse>
  ): QueryVerifiableCredentialResponse {
    const message = {
      ...baseQueryVerifiableCredentialResponse,
    } as QueryVerifiableCredentialResponse;
    if (
      object.verifiableCredential !== undefined &&
      object.verifiableCredential !== null
    ) {
      message.verifiableCredential = VerifiableCredential.fromPartial(
        object.verifiableCredential
      );
    } else {
      message.verifiableCredential = undefined;
    }
    return message;
  },
};

const baseQueryValidateVerifiableCredentialRequest: object = {
  verifiableCredentialId: "",
  issuerPubkey: "",
};

export const QueryValidateVerifiableCredentialRequest = {
  encode(
    message: QueryValidateVerifiableCredentialRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.verifiableCredentialId !== "") {
      writer.uint32(10).string(message.verifiableCredentialId);
    }
    if (message.issuerPubkey !== "") {
      writer.uint32(18).string(message.issuerPubkey);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryValidateVerifiableCredentialRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryValidateVerifiableCredentialRequest,
    } as QueryValidateVerifiableCredentialRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.verifiableCredentialId = reader.string();
          break;
        case 2:
          message.issuerPubkey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryValidateVerifiableCredentialRequest {
    const message = {
      ...baseQueryValidateVerifiableCredentialRequest,
    } as QueryValidateVerifiableCredentialRequest;
    if (
      object.verifiableCredentialId !== undefined &&
      object.verifiableCredentialId !== null
    ) {
      message.verifiableCredentialId = String(object.verifiableCredentialId);
    } else {
      message.verifiableCredentialId = "";
    }
    if (object.issuerPubkey !== undefined && object.issuerPubkey !== null) {
      message.issuerPubkey = String(object.issuerPubkey);
    } else {
      message.issuerPubkey = "";
    }
    return message;
  },

  toJSON(message: QueryValidateVerifiableCredentialRequest): unknown {
    const obj: any = {};
    message.verifiableCredentialId !== undefined &&
      (obj.verifiableCredentialId = message.verifiableCredentialId);
    message.issuerPubkey !== undefined &&
      (obj.issuerPubkey = message.issuerPubkey);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryValidateVerifiableCredentialRequest>
  ): QueryValidateVerifiableCredentialRequest {
    const message = {
      ...baseQueryValidateVerifiableCredentialRequest,
    } as QueryValidateVerifiableCredentialRequest;
    if (
      object.verifiableCredentialId !== undefined &&
      object.verifiableCredentialId !== null
    ) {
      message.verifiableCredentialId = object.verifiableCredentialId;
    } else {
      message.verifiableCredentialId = "";
    }
    if (object.issuerPubkey !== undefined && object.issuerPubkey !== null) {
      message.issuerPubkey = object.issuerPubkey;
    } else {
      message.issuerPubkey = "";
    }
    return message;
  },
};

const baseQueryValidateVerifiableCredentialResponse: object = {
  isValid: false,
};

export const QueryValidateVerifiableCredentialResponse = {
  encode(
    message: QueryValidateVerifiableCredentialResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.isValid === true) {
      writer.uint32(8).bool(message.isValid);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryValidateVerifiableCredentialResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryValidateVerifiableCredentialResponse,
    } as QueryValidateVerifiableCredentialResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.isValid = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryValidateVerifiableCredentialResponse {
    const message = {
      ...baseQueryValidateVerifiableCredentialResponse,
    } as QueryValidateVerifiableCredentialResponse;
    if (object.isValid !== undefined && object.isValid !== null) {
      message.isValid = Boolean(object.isValid);
    } else {
      message.isValid = false;
    }
    return message;
  },

  toJSON(message: QueryValidateVerifiableCredentialResponse): unknown {
    const obj: any = {};
    message.isValid !== undefined && (obj.isValid = message.isValid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryValidateVerifiableCredentialResponse>
  ): QueryValidateVerifiableCredentialResponse {
    const message = {
      ...baseQueryValidateVerifiableCredentialResponse,
    } as QueryValidateVerifiableCredentialResponse;
    if (object.isValid !== undefined && object.isValid !== null) {
      message.isValid = object.isValid;
    } else {
      message.isValid = false;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Identifers queries all validators that match the given status. */
  VerifiableCredentials(
    request: QueryVerifiableCredentialsRequest
  ): Promise<QueryVerifiableCredentialsResponse>;
  /** VerifiableCredential queries validator info for given validator address. */
  VerifiableCredential(
    request: QueryVerifiableCredentialRequest
  ): Promise<QueryVerifiableCredentialResponse>;
  /** ValidateVerifiableCredential queries validator info for given validator address. */
  ValidateVerifiableCredential(
    request: QueryValidateVerifiableCredentialRequest
  ): Promise<QueryValidateVerifiableCredentialResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  VerifiableCredentials(
    request: QueryVerifiableCredentialsRequest
  ): Promise<QueryVerifiableCredentialsResponse> {
    const data = QueryVerifiableCredentialsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.verifiablecredentialservice.Query",
      "VerifiableCredentials",
      data
    );
    return promise.then((data) =>
      QueryVerifiableCredentialsResponse.decode(new Reader(data))
    );
  }

  VerifiableCredential(
    request: QueryVerifiableCredentialRequest
  ): Promise<QueryVerifiableCredentialResponse> {
    const data = QueryVerifiableCredentialRequest.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.verifiablecredentialservice.Query",
      "VerifiableCredential",
      data
    );
    return promise.then((data) =>
      QueryVerifiableCredentialResponse.decode(new Reader(data))
    );
  }

  ValidateVerifiableCredential(
    request: QueryValidateVerifiableCredentialRequest
  ): Promise<QueryValidateVerifiableCredentialResponse> {
    const data = QueryValidateVerifiableCredentialRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.verifiablecredentialservice.Query",
      "ValidateVerifiableCredential",
      data
    );
    return promise.then((data) =>
      QueryValidateVerifiableCredentialResponse.decode(new Reader(data))
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
