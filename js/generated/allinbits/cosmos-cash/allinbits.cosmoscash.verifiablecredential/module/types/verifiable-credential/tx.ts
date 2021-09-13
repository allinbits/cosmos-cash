/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { VerifiableCredential } from "../verifiable-credential/verifiable-credential";

export const protobufPackage = "allinbits.cosmoscash.verifiablecredential";

/** MsgCreateVerifiableCredential defines a SDK message for creating a new identifer. */
export interface MsgCreateVerifiableCredential {
  verifiableCredential: VerifiableCredential | undefined;
  /** owner represents the user creating the message */
  owner: string;
}

export interface MsgCreateVerifiableCredentialResponse {}

/** MsgDeleteVerifiableCredential defines a SDK message for updating a credential. */
export interface MsgDeleteVerifiableCredential {
  verifiableCredentialId: string;
  issuerDid: string;
  /** owner represents the user creating the message */
  owner: string;
}

export interface MsgDeleteVerifiableCredentialResponse {}

const baseMsgCreateVerifiableCredential: object = { owner: "" };

export const MsgCreateVerifiableCredential = {
  encode(
    message: MsgCreateVerifiableCredential,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.verifiableCredential !== undefined) {
      VerifiableCredential.encode(
        message.verifiableCredential,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateVerifiableCredential {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateVerifiableCredential,
    } as MsgCreateVerifiableCredential;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.verifiableCredential = VerifiableCredential.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateVerifiableCredential {
    const message = {
      ...baseMsgCreateVerifiableCredential,
    } as MsgCreateVerifiableCredential;
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
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: MsgCreateVerifiableCredential): unknown {
    const obj: any = {};
    message.verifiableCredential !== undefined &&
      (obj.verifiableCredential = message.verifiableCredential
        ? VerifiableCredential.toJSON(message.verifiableCredential)
        : undefined);
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateVerifiableCredential>
  ): MsgCreateVerifiableCredential {
    const message = {
      ...baseMsgCreateVerifiableCredential,
    } as MsgCreateVerifiableCredential;
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
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

const baseMsgCreateVerifiableCredentialResponse: object = {};

export const MsgCreateVerifiableCredentialResponse = {
  encode(
    _: MsgCreateVerifiableCredentialResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateVerifiableCredentialResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateVerifiableCredentialResponse,
    } as MsgCreateVerifiableCredentialResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateVerifiableCredentialResponse {
    const message = {
      ...baseMsgCreateVerifiableCredentialResponse,
    } as MsgCreateVerifiableCredentialResponse;
    return message;
  },

  toJSON(_: MsgCreateVerifiableCredentialResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateVerifiableCredentialResponse>
  ): MsgCreateVerifiableCredentialResponse {
    const message = {
      ...baseMsgCreateVerifiableCredentialResponse,
    } as MsgCreateVerifiableCredentialResponse;
    return message;
  },
};

const baseMsgDeleteVerifiableCredential: object = {
  verifiableCredentialId: "",
  issuerDid: "",
  owner: "",
};

export const MsgDeleteVerifiableCredential = {
  encode(
    message: MsgDeleteVerifiableCredential,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.verifiableCredentialId !== "") {
      writer.uint32(10).string(message.verifiableCredentialId);
    }
    if (message.issuerDid !== "") {
      writer.uint32(18).string(message.issuerDid);
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteVerifiableCredential {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteVerifiableCredential,
    } as MsgDeleteVerifiableCredential;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.verifiableCredentialId = reader.string();
          break;
        case 2:
          message.issuerDid = reader.string();
          break;
        case 3:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteVerifiableCredential {
    const message = {
      ...baseMsgDeleteVerifiableCredential,
    } as MsgDeleteVerifiableCredential;
    if (
      object.verifiableCredentialId !== undefined &&
      object.verifiableCredentialId !== null
    ) {
      message.verifiableCredentialId = String(object.verifiableCredentialId);
    } else {
      message.verifiableCredentialId = "";
    }
    if (object.issuerDid !== undefined && object.issuerDid !== null) {
      message.issuerDid = String(object.issuerDid);
    } else {
      message.issuerDid = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteVerifiableCredential): unknown {
    const obj: any = {};
    message.verifiableCredentialId !== undefined &&
      (obj.verifiableCredentialId = message.verifiableCredentialId);
    message.issuerDid !== undefined && (obj.issuerDid = message.issuerDid);
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteVerifiableCredential>
  ): MsgDeleteVerifiableCredential {
    const message = {
      ...baseMsgDeleteVerifiableCredential,
    } as MsgDeleteVerifiableCredential;
    if (
      object.verifiableCredentialId !== undefined &&
      object.verifiableCredentialId !== null
    ) {
      message.verifiableCredentialId = object.verifiableCredentialId;
    } else {
      message.verifiableCredentialId = "";
    }
    if (object.issuerDid !== undefined && object.issuerDid !== null) {
      message.issuerDid = object.issuerDid;
    } else {
      message.issuerDid = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

const baseMsgDeleteVerifiableCredentialResponse: object = {};

export const MsgDeleteVerifiableCredentialResponse = {
  encode(
    _: MsgDeleteVerifiableCredentialResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteVerifiableCredentialResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteVerifiableCredentialResponse,
    } as MsgDeleteVerifiableCredentialResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteVerifiableCredentialResponse {
    const message = {
      ...baseMsgDeleteVerifiableCredentialResponse,
    } as MsgDeleteVerifiableCredentialResponse;
    return message;
  },

  toJSON(_: MsgDeleteVerifiableCredentialResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteVerifiableCredentialResponse>
  ): MsgDeleteVerifiableCredentialResponse {
    const message = {
      ...baseMsgDeleteVerifiableCredentialResponse,
    } as MsgDeleteVerifiableCredentialResponse;
    return message;
  },
};

/** Msg defines the identity Msg service. */
export interface Msg {
  CreateVerifiableCredential(
    request: MsgCreateVerifiableCredential
  ): Promise<MsgCreateVerifiableCredentialResponse>;
  DeleteVerifiableCredential(
    request: MsgDeleteVerifiableCredential
  ): Promise<MsgDeleteVerifiableCredentialResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateVerifiableCredential(
    request: MsgCreateVerifiableCredential
  ): Promise<MsgCreateVerifiableCredentialResponse> {
    const data = MsgCreateVerifiableCredential.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.verifiablecredential.Msg",
      "CreateVerifiableCredential",
      data
    );
    return promise.then((data) =>
      MsgCreateVerifiableCredentialResponse.decode(new Reader(data))
    );
  }

  DeleteVerifiableCredential(
    request: MsgDeleteVerifiableCredential
  ): Promise<MsgDeleteVerifiableCredentialResponse> {
    const data = MsgDeleteVerifiableCredential.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.verifiablecredential.Msg",
      "DeleteVerifiableCredential",
      data
    );
    return promise.then((data) =>
      MsgDeleteVerifiableCredentialResponse.decode(new Reader(data))
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
