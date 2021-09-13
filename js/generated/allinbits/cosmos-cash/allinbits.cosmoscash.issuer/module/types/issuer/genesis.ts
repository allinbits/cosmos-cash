/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "allinbits.cosmoscash.issuer";

/** GenesisState defines the issuer module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  regulatorsParams: RegulatorsParams | undefined;
}

/** RegulatorsParams defines the addresses of the regulators */
export interface RegulatorsParams {
  /**
   * the addresses of the regualtors for the chain. The addresses will be used to
   * generate DID documents at genesis.
   */
  addresses: string[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.regulatorsParams !== undefined) {
      RegulatorsParams.encode(
        message.regulatorsParams,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.regulatorsParams = RegulatorsParams.decode(
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

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    if (
      object.regulatorsParams !== undefined &&
      object.regulatorsParams !== null
    ) {
      message.regulatorsParams = RegulatorsParams.fromJSON(
        object.regulatorsParams
      );
    } else {
      message.regulatorsParams = undefined;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.regulatorsParams !== undefined &&
      (obj.regulatorsParams = message.regulatorsParams
        ? RegulatorsParams.toJSON(message.regulatorsParams)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    if (
      object.regulatorsParams !== undefined &&
      object.regulatorsParams !== null
    ) {
      message.regulatorsParams = RegulatorsParams.fromPartial(
        object.regulatorsParams
      );
    } else {
      message.regulatorsParams = undefined;
    }
    return message;
  },
};

const baseRegulatorsParams: object = { addresses: "" };

export const RegulatorsParams = {
  encode(message: RegulatorsParams, writer: Writer = Writer.create()): Writer {
    for (const v of message.addresses) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): RegulatorsParams {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRegulatorsParams } as RegulatorsParams;
    message.addresses = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.addresses.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegulatorsParams {
    const message = { ...baseRegulatorsParams } as RegulatorsParams;
    message.addresses = [];
    if (object.addresses !== undefined && object.addresses !== null) {
      for (const e of object.addresses) {
        message.addresses.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: RegulatorsParams): unknown {
    const obj: any = {};
    if (message.addresses) {
      obj.addresses = message.addresses.map((e) => e);
    } else {
      obj.addresses = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<RegulatorsParams>): RegulatorsParams {
    const message = { ...baseRegulatorsParams } as RegulatorsParams;
    message.addresses = [];
    if (object.addresses !== undefined && object.addresses !== null) {
      for (const e of object.addresses) {
        message.addresses.push(e);
      }
    }
    return message;
  },
};

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
