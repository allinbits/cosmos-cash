/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "allinbits.cosmoscash.issuer";

/** MsgCreateIssuer defines an SDK message for creating an emoney token issuer. */
export interface MsgCreateIssuer {
  token: string;
  fee: number;
  issuerDid: string;
  licenseCredId: string;
  owner: string;
}

export interface MsgCreateIssuerResponse {}

/** MsgBurnToken defines a SDK message for burning issuer tokens. */
export interface MsgBurnToken {
  amount: string;
  owner: string;
}

export interface MsgBurnTokenResponse {}

/** MsgMintToken defines a SDK message for minting a token */
export interface MsgMintToken {
  amount: string;
  owner: string;
}

export interface MsgMintTokenResponse {}

const baseMsgCreateIssuer: object = {
  token: "",
  fee: 0,
  issuerDid: "",
  licenseCredId: "",
  owner: "",
};

export const MsgCreateIssuer = {
  encode(message: MsgCreateIssuer, writer: Writer = Writer.create()): Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.fee !== 0) {
      writer.uint32(16).int32(message.fee);
    }
    if (message.issuerDid !== "") {
      writer.uint32(26).string(message.issuerDid);
    }
    if (message.licenseCredId !== "") {
      writer.uint32(34).string(message.licenseCredId);
    }
    if (message.owner !== "") {
      writer.uint32(42).string(message.owner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateIssuer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateIssuer } as MsgCreateIssuer;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        case 2:
          message.fee = reader.int32();
          break;
        case 3:
          message.issuerDid = reader.string();
          break;
        case 4:
          message.licenseCredId = reader.string();
          break;
        case 5:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateIssuer {
    const message = { ...baseMsgCreateIssuer } as MsgCreateIssuer;
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = Number(object.fee);
    } else {
      message.fee = 0;
    }
    if (object.issuerDid !== undefined && object.issuerDid !== null) {
      message.issuerDid = String(object.issuerDid);
    } else {
      message.issuerDid = "";
    }
    if (object.licenseCredId !== undefined && object.licenseCredId !== null) {
      message.licenseCredId = String(object.licenseCredId);
    } else {
      message.licenseCredId = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: MsgCreateIssuer): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    message.fee !== undefined && (obj.fee = message.fee);
    message.issuerDid !== undefined && (obj.issuerDid = message.issuerDid);
    message.licenseCredId !== undefined &&
      (obj.licenseCredId = message.licenseCredId);
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateIssuer>): MsgCreateIssuer {
    const message = { ...baseMsgCreateIssuer } as MsgCreateIssuer;
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = 0;
    }
    if (object.issuerDid !== undefined && object.issuerDid !== null) {
      message.issuerDid = object.issuerDid;
    } else {
      message.issuerDid = "";
    }
    if (object.licenseCredId !== undefined && object.licenseCredId !== null) {
      message.licenseCredId = object.licenseCredId;
    } else {
      message.licenseCredId = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

const baseMsgCreateIssuerResponse: object = {};

export const MsgCreateIssuerResponse = {
  encode(_: MsgCreateIssuerResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateIssuerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateIssuerResponse,
    } as MsgCreateIssuerResponse;
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

  fromJSON(_: any): MsgCreateIssuerResponse {
    const message = {
      ...baseMsgCreateIssuerResponse,
    } as MsgCreateIssuerResponse;
    return message;
  },

  toJSON(_: MsgCreateIssuerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateIssuerResponse>
  ): MsgCreateIssuerResponse {
    const message = {
      ...baseMsgCreateIssuerResponse,
    } as MsgCreateIssuerResponse;
    return message;
  },
};

const baseMsgBurnToken: object = { amount: "", owner: "" };

export const MsgBurnToken = {
  encode(message: MsgBurnToken, writer: Writer = Writer.create()): Writer {
    if (message.amount !== "") {
      writer.uint32(10).string(message.amount);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBurnToken {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBurnToken } as MsgBurnToken;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.amount = reader.string();
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

  fromJSON(object: any): MsgBurnToken {
    const message = { ...baseMsgBurnToken } as MsgBurnToken;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: MsgBurnToken): unknown {
    const obj: any = {};
    message.amount !== undefined && (obj.amount = message.amount);
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgBurnToken>): MsgBurnToken {
    const message = { ...baseMsgBurnToken } as MsgBurnToken;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

const baseMsgBurnTokenResponse: object = {};

export const MsgBurnTokenResponse = {
  encode(_: MsgBurnTokenResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBurnTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBurnTokenResponse } as MsgBurnTokenResponse;
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

  fromJSON(_: any): MsgBurnTokenResponse {
    const message = { ...baseMsgBurnTokenResponse } as MsgBurnTokenResponse;
    return message;
  },

  toJSON(_: MsgBurnTokenResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgBurnTokenResponse>): MsgBurnTokenResponse {
    const message = { ...baseMsgBurnTokenResponse } as MsgBurnTokenResponse;
    return message;
  },
};

const baseMsgMintToken: object = { amount: "", owner: "" };

export const MsgMintToken = {
  encode(message: MsgMintToken, writer: Writer = Writer.create()): Writer {
    if (message.amount !== "") {
      writer.uint32(10).string(message.amount);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMintToken {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMintToken } as MsgMintToken;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.amount = reader.string();
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

  fromJSON(object: any): MsgMintToken {
    const message = { ...baseMsgMintToken } as MsgMintToken;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: MsgMintToken): unknown {
    const obj: any = {};
    message.amount !== undefined && (obj.amount = message.amount);
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgMintToken>): MsgMintToken {
    const message = { ...baseMsgMintToken } as MsgMintToken;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

const baseMsgMintTokenResponse: object = {};

export const MsgMintTokenResponse = {
  encode(_: MsgMintTokenResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMintTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMintTokenResponse } as MsgMintTokenResponse;
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

  fromJSON(_: any): MsgMintTokenResponse {
    const message = { ...baseMsgMintTokenResponse } as MsgMintTokenResponse;
    return message;
  },

  toJSON(_: MsgMintTokenResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgMintTokenResponse>): MsgMintTokenResponse {
    const message = { ...baseMsgMintTokenResponse } as MsgMintTokenResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateIssuer(request: MsgCreateIssuer): Promise<MsgCreateIssuerResponse>;
  BurnToken(request: MsgBurnToken): Promise<MsgBurnTokenResponse>;
  MintToken(request: MsgMintToken): Promise<MsgMintTokenResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateIssuer(request: MsgCreateIssuer): Promise<MsgCreateIssuerResponse> {
    const data = MsgCreateIssuer.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.issuer.Msg",
      "CreateIssuer",
      data
    );
    return promise.then((data) =>
      MsgCreateIssuerResponse.decode(new Reader(data))
    );
  }

  BurnToken(request: MsgBurnToken): Promise<MsgBurnTokenResponse> {
    const data = MsgBurnToken.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.issuer.Msg",
      "BurnToken",
      data
    );
    return promise.then((data) =>
      MsgBurnTokenResponse.decode(new Reader(data))
    );
  }

  MintToken(request: MsgMintToken): Promise<MsgMintTokenResponse> {
    const data = MsgMintToken.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.issuer.Msg",
      "MintToken",
      data
    );
    return promise.then((data) =>
      MsgMintTokenResponse.decode(new Reader(data))
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
