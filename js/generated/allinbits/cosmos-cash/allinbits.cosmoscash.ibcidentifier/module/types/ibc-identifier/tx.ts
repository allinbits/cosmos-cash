/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Height } from "../ibc/core/client/v1/client";

export const protobufPackage = "allinbits.cosmoscash.ibcidentifier";

/** MsgTransferIdentifierIBC defines a SDK message for creating a new identifer. */
export interface MsgTransferIdentifierIBC {
  id: string;
  /** the port by which the packet will be sent */
  sourcePort: string;
  /** the channel by which the packet will be sent */
  sourceChannel: string;
  /**
   * Timeout height relative to the current block height.
   * The timeout is disabled when set to 0.
   */
  timeoutHeight: Height | undefined;
  /**
   * Timeout timestamp (in nanoseconds) relative to the current block timestamp.
   * The timeout is disabled when set to 0.
   */
  timeoutTimestamp: number;
  owner: string;
}

export interface MsgTransferIdentifierIBCResponse {}

const baseMsgTransferIdentifierIBC: object = {
  id: "",
  sourcePort: "",
  sourceChannel: "",
  timeoutTimestamp: 0,
  owner: "",
};

export const MsgTransferIdentifierIBC = {
  encode(
    message: MsgTransferIdentifierIBC,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.sourcePort !== "") {
      writer.uint32(18).string(message.sourcePort);
    }
    if (message.sourceChannel !== "") {
      writer.uint32(26).string(message.sourceChannel);
    }
    if (message.timeoutHeight !== undefined) {
      Height.encode(message.timeoutHeight, writer.uint32(34).fork()).ldelim();
    }
    if (message.timeoutTimestamp !== 0) {
      writer.uint32(40).uint64(message.timeoutTimestamp);
    }
    if (message.owner !== "") {
      writer.uint32(50).string(message.owner);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgTransferIdentifierIBC {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgTransferIdentifierIBC,
    } as MsgTransferIdentifierIBC;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.sourcePort = reader.string();
          break;
        case 3:
          message.sourceChannel = reader.string();
          break;
        case 4:
          message.timeoutHeight = Height.decode(reader, reader.uint32());
          break;
        case 5:
          message.timeoutTimestamp = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgTransferIdentifierIBC {
    const message = {
      ...baseMsgTransferIdentifierIBC,
    } as MsgTransferIdentifierIBC;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.sourcePort !== undefined && object.sourcePort !== null) {
      message.sourcePort = String(object.sourcePort);
    } else {
      message.sourcePort = "";
    }
    if (object.sourceChannel !== undefined && object.sourceChannel !== null) {
      message.sourceChannel = String(object.sourceChannel);
    } else {
      message.sourceChannel = "";
    }
    if (object.timeoutHeight !== undefined && object.timeoutHeight !== null) {
      message.timeoutHeight = Height.fromJSON(object.timeoutHeight);
    } else {
      message.timeoutHeight = undefined;
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = Number(object.timeoutTimestamp);
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: MsgTransferIdentifierIBC): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.sourcePort !== undefined && (obj.sourcePort = message.sourcePort);
    message.sourceChannel !== undefined &&
      (obj.sourceChannel = message.sourceChannel);
    message.timeoutHeight !== undefined &&
      (obj.timeoutHeight = message.timeoutHeight
        ? Height.toJSON(message.timeoutHeight)
        : undefined);
    message.timeoutTimestamp !== undefined &&
      (obj.timeoutTimestamp = message.timeoutTimestamp);
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgTransferIdentifierIBC>
  ): MsgTransferIdentifierIBC {
    const message = {
      ...baseMsgTransferIdentifierIBC,
    } as MsgTransferIdentifierIBC;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.sourcePort !== undefined && object.sourcePort !== null) {
      message.sourcePort = object.sourcePort;
    } else {
      message.sourcePort = "";
    }
    if (object.sourceChannel !== undefined && object.sourceChannel !== null) {
      message.sourceChannel = object.sourceChannel;
    } else {
      message.sourceChannel = "";
    }
    if (object.timeoutHeight !== undefined && object.timeoutHeight !== null) {
      message.timeoutHeight = Height.fromPartial(object.timeoutHeight);
    } else {
      message.timeoutHeight = undefined;
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = object.timeoutTimestamp;
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

const baseMsgTransferIdentifierIBCResponse: object = {};

export const MsgTransferIdentifierIBCResponse = {
  encode(
    _: MsgTransferIdentifierIBCResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgTransferIdentifierIBCResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgTransferIdentifierIBCResponse,
    } as MsgTransferIdentifierIBCResponse;
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

  fromJSON(_: any): MsgTransferIdentifierIBCResponse {
    const message = {
      ...baseMsgTransferIdentifierIBCResponse,
    } as MsgTransferIdentifierIBCResponse;
    return message;
  },

  toJSON(_: MsgTransferIdentifierIBCResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgTransferIdentifierIBCResponse>
  ): MsgTransferIdentifierIBCResponse {
    const message = {
      ...baseMsgTransferIdentifierIBCResponse,
    } as MsgTransferIdentifierIBCResponse;
    return message;
  },
};

/** Msg defines the identity Msg service. */
export interface Msg {
  /** TransferIdentifierIBC defines a method for transfering an identifier to another chain. */
  TransferIdentifierIBC(
    request: MsgTransferIdentifierIBC
  ): Promise<MsgTransferIdentifierIBCResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  TransferIdentifierIBC(
    request: MsgTransferIdentifierIBC
  ): Promise<MsgTransferIdentifierIBCResponse> {
    const data = MsgTransferIdentifierIBC.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.ibcidentifier.Msg",
      "TransferIdentifierIBC",
      data
    );
    return promise.then((data) =>
      MsgTransferIdentifierIBCResponse.decode(new Reader(data))
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

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
