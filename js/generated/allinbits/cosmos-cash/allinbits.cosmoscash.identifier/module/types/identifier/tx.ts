/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Authentication, Service } from "../identifier/identifier";

export const protobufPackage = "allinbits.cosmoscash.identifier";

/** MsgCreateIdentifier defines a SDK message for creating a new identifer. */
export interface MsgCreateIdentifier {
  context: string;
  id: string;
  /** authentication represents public key associated with the did document. */
  authentication: Authentication[];
  /** services represents each service associated with a did */
  services: Service[];
  /** owner represents the user creating the message */
  owner: string;
}

export interface MsgCreateIdentifierResponse {}

export interface MsgAddAuthentication {
  id: string;
  /** authentication represents public key associated with the did document. */
  authentication: Authentication | undefined;
  /** owner is the address of the user creating the message */
  owner: string;
}

export interface MsgAddAuthenticationResponse {}

export interface MsgAddService {
  id: string;
  /** authentication represents public key associated with the did document. */
  serviceData: Service | undefined;
  /** owner is the address of the user creating the message */
  owner: string;
}

export interface MsgAddServiceResponse {}

export interface MsgDeleteAuthentication {
  id: string;
  key: string;
  /** owner is the address of the user creating the message */
  owner: string;
}

export interface MsgDeleteAuthenticationResponse {}

export interface MsgDeleteService {
  id: string;
  serviceId: string;
  /** owner is the address of the user creating the message */
  owner: string;
}

export interface MsgDeleteServiceResponse {}

const baseMsgCreateIdentifier: object = { context: "", id: "", owner: "" };

export const MsgCreateIdentifier = {
  encode(
    message: MsgCreateIdentifier,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.context !== "") {
      writer.uint32(10).string(message.context);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    for (const v of message.authentication) {
      Authentication.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.services) {
      Service.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.owner !== "") {
      writer.uint32(42).string(message.owner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateIdentifier {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateIdentifier } as MsgCreateIdentifier;
    message.authentication = [];
    message.services = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.context = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        case 3:
          message.authentication.push(
            Authentication.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.services.push(Service.decode(reader, reader.uint32()));
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

  fromJSON(object: any): MsgCreateIdentifier {
    const message = { ...baseMsgCreateIdentifier } as MsgCreateIdentifier;
    message.authentication = [];
    message.services = [];
    if (object.context !== undefined && object.context !== null) {
      message.context = String(object.context);
    } else {
      message.context = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.authentication !== undefined && object.authentication !== null) {
      for (const e of object.authentication) {
        message.authentication.push(Authentication.fromJSON(e));
      }
    }
    if (object.services !== undefined && object.services !== null) {
      for (const e of object.services) {
        message.services.push(Service.fromJSON(e));
      }
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: MsgCreateIdentifier): unknown {
    const obj: any = {};
    message.context !== undefined && (obj.context = message.context);
    message.id !== undefined && (obj.id = message.id);
    if (message.authentication) {
      obj.authentication = message.authentication.map((e) =>
        e ? Authentication.toJSON(e) : undefined
      );
    } else {
      obj.authentication = [];
    }
    if (message.services) {
      obj.services = message.services.map((e) =>
        e ? Service.toJSON(e) : undefined
      );
    } else {
      obj.services = [];
    }
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateIdentifier>): MsgCreateIdentifier {
    const message = { ...baseMsgCreateIdentifier } as MsgCreateIdentifier;
    message.authentication = [];
    message.services = [];
    if (object.context !== undefined && object.context !== null) {
      message.context = object.context;
    } else {
      message.context = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.authentication !== undefined && object.authentication !== null) {
      for (const e of object.authentication) {
        message.authentication.push(Authentication.fromPartial(e));
      }
    }
    if (object.services !== undefined && object.services !== null) {
      for (const e of object.services) {
        message.services.push(Service.fromPartial(e));
      }
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

const baseMsgCreateIdentifierResponse: object = {};

export const MsgCreateIdentifierResponse = {
  encode(
    _: MsgCreateIdentifierResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateIdentifierResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateIdentifierResponse,
    } as MsgCreateIdentifierResponse;
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

  fromJSON(_: any): MsgCreateIdentifierResponse {
    const message = {
      ...baseMsgCreateIdentifierResponse,
    } as MsgCreateIdentifierResponse;
    return message;
  },

  toJSON(_: MsgCreateIdentifierResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateIdentifierResponse>
  ): MsgCreateIdentifierResponse {
    const message = {
      ...baseMsgCreateIdentifierResponse,
    } as MsgCreateIdentifierResponse;
    return message;
  },
};

const baseMsgAddAuthentication: object = { id: "", owner: "" };

export const MsgAddAuthentication = {
  encode(
    message: MsgAddAuthentication,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.authentication !== undefined) {
      Authentication.encode(
        message.authentication,
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddAuthentication {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddAuthentication } as MsgAddAuthentication;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.authentication = Authentication.decode(
            reader,
            reader.uint32()
          );
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

  fromJSON(object: any): MsgAddAuthentication {
    const message = { ...baseMsgAddAuthentication } as MsgAddAuthentication;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.authentication !== undefined && object.authentication !== null) {
      message.authentication = Authentication.fromJSON(object.authentication);
    } else {
      message.authentication = undefined;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: MsgAddAuthentication): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.authentication !== undefined &&
      (obj.authentication = message.authentication
        ? Authentication.toJSON(message.authentication)
        : undefined);
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddAuthentication>): MsgAddAuthentication {
    const message = { ...baseMsgAddAuthentication } as MsgAddAuthentication;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.authentication !== undefined && object.authentication !== null) {
      message.authentication = Authentication.fromPartial(
        object.authentication
      );
    } else {
      message.authentication = undefined;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

const baseMsgAddAuthenticationResponse: object = {};

export const MsgAddAuthenticationResponse = {
  encode(
    _: MsgAddAuthenticationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddAuthenticationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddAuthenticationResponse,
    } as MsgAddAuthenticationResponse;
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

  fromJSON(_: any): MsgAddAuthenticationResponse {
    const message = {
      ...baseMsgAddAuthenticationResponse,
    } as MsgAddAuthenticationResponse;
    return message;
  },

  toJSON(_: MsgAddAuthenticationResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgAddAuthenticationResponse>
  ): MsgAddAuthenticationResponse {
    const message = {
      ...baseMsgAddAuthenticationResponse,
    } as MsgAddAuthenticationResponse;
    return message;
  },
};

const baseMsgAddService: object = { id: "", owner: "" };

export const MsgAddService = {
  encode(message: MsgAddService, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.serviceData !== undefined) {
      Service.encode(message.serviceData, writer.uint32(18).fork()).ldelim();
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddService {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddService } as MsgAddService;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.serviceData = Service.decode(reader, reader.uint32());
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

  fromJSON(object: any): MsgAddService {
    const message = { ...baseMsgAddService } as MsgAddService;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.serviceData !== undefined && object.serviceData !== null) {
      message.serviceData = Service.fromJSON(object.serviceData);
    } else {
      message.serviceData = undefined;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: MsgAddService): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.serviceData !== undefined &&
      (obj.serviceData = message.serviceData
        ? Service.toJSON(message.serviceData)
        : undefined);
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddService>): MsgAddService {
    const message = { ...baseMsgAddService } as MsgAddService;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.serviceData !== undefined && object.serviceData !== null) {
      message.serviceData = Service.fromPartial(object.serviceData);
    } else {
      message.serviceData = undefined;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

const baseMsgAddServiceResponse: object = {};

export const MsgAddServiceResponse = {
  encode(_: MsgAddServiceResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddServiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddServiceResponse } as MsgAddServiceResponse;
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

  fromJSON(_: any): MsgAddServiceResponse {
    const message = { ...baseMsgAddServiceResponse } as MsgAddServiceResponse;
    return message;
  },

  toJSON(_: MsgAddServiceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgAddServiceResponse>): MsgAddServiceResponse {
    const message = { ...baseMsgAddServiceResponse } as MsgAddServiceResponse;
    return message;
  },
};

const baseMsgDeleteAuthentication: object = { id: "", key: "", owner: "" };

export const MsgDeleteAuthentication = {
  encode(
    message: MsgDeleteAuthentication,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteAuthentication {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteAuthentication,
    } as MsgDeleteAuthentication;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.key = reader.string();
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

  fromJSON(object: any): MsgDeleteAuthentication {
    const message = {
      ...baseMsgDeleteAuthentication,
    } as MsgDeleteAuthentication;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteAuthentication): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.key !== undefined && (obj.key = message.key);
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteAuthentication>
  ): MsgDeleteAuthentication {
    const message = {
      ...baseMsgDeleteAuthentication,
    } as MsgDeleteAuthentication;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

const baseMsgDeleteAuthenticationResponse: object = {};

export const MsgDeleteAuthenticationResponse = {
  encode(
    _: MsgDeleteAuthenticationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteAuthenticationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteAuthenticationResponse,
    } as MsgDeleteAuthenticationResponse;
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

  fromJSON(_: any): MsgDeleteAuthenticationResponse {
    const message = {
      ...baseMsgDeleteAuthenticationResponse,
    } as MsgDeleteAuthenticationResponse;
    return message;
  },

  toJSON(_: MsgDeleteAuthenticationResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteAuthenticationResponse>
  ): MsgDeleteAuthenticationResponse {
    const message = {
      ...baseMsgDeleteAuthenticationResponse,
    } as MsgDeleteAuthenticationResponse;
    return message;
  },
};

const baseMsgDeleteService: object = { id: "", serviceId: "", owner: "" };

export const MsgDeleteService = {
  encode(message: MsgDeleteService, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.serviceId !== "") {
      writer.uint32(18).string(message.serviceId);
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteService {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteService } as MsgDeleteService;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.serviceId = reader.string();
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

  fromJSON(object: any): MsgDeleteService {
    const message = { ...baseMsgDeleteService } as MsgDeleteService;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.serviceId !== undefined && object.serviceId !== null) {
      message.serviceId = String(object.serviceId);
    } else {
      message.serviceId = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteService): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.serviceId !== undefined && (obj.serviceId = message.serviceId);
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteService>): MsgDeleteService {
    const message = { ...baseMsgDeleteService } as MsgDeleteService;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.serviceId !== undefined && object.serviceId !== null) {
      message.serviceId = object.serviceId;
    } else {
      message.serviceId = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

const baseMsgDeleteServiceResponse: object = {};

export const MsgDeleteServiceResponse = {
  encode(
    _: MsgDeleteServiceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteServiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteServiceResponse,
    } as MsgDeleteServiceResponse;
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

  fromJSON(_: any): MsgDeleteServiceResponse {
    const message = {
      ...baseMsgDeleteServiceResponse,
    } as MsgDeleteServiceResponse;
    return message;
  },

  toJSON(_: MsgDeleteServiceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteServiceResponse>
  ): MsgDeleteServiceResponse {
    const message = {
      ...baseMsgDeleteServiceResponse,
    } as MsgDeleteServiceResponse;
    return message;
  },
};

/** Msg defines the identity Msg service. */
export interface Msg {
  /** CreateDidDocument defines a method for creating a new identity. */
  CreateIdentifier(
    request: MsgCreateIdentifier
  ): Promise<MsgCreateIdentifierResponse>;
  AddAuthentication(
    request: MsgAddAuthentication
  ): Promise<MsgAddAuthenticationResponse>;
  AddService(request: MsgAddService): Promise<MsgAddServiceResponse>;
  DeleteAuthentication(
    request: MsgDeleteAuthentication
  ): Promise<MsgDeleteAuthenticationResponse>;
  DeleteService(request: MsgDeleteService): Promise<MsgDeleteServiceResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateIdentifier(
    request: MsgCreateIdentifier
  ): Promise<MsgCreateIdentifierResponse> {
    const data = MsgCreateIdentifier.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.identifier.Msg",
      "CreateIdentifier",
      data
    );
    return promise.then((data) =>
      MsgCreateIdentifierResponse.decode(new Reader(data))
    );
  }

  AddAuthentication(
    request: MsgAddAuthentication
  ): Promise<MsgAddAuthenticationResponse> {
    const data = MsgAddAuthentication.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.identifier.Msg",
      "AddAuthentication",
      data
    );
    return promise.then((data) =>
      MsgAddAuthenticationResponse.decode(new Reader(data))
    );
  }

  AddService(request: MsgAddService): Promise<MsgAddServiceResponse> {
    const data = MsgAddService.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.identifier.Msg",
      "AddService",
      data
    );
    return promise.then((data) =>
      MsgAddServiceResponse.decode(new Reader(data))
    );
  }

  DeleteAuthentication(
    request: MsgDeleteAuthentication
  ): Promise<MsgDeleteAuthenticationResponse> {
    const data = MsgDeleteAuthentication.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.identifier.Msg",
      "DeleteAuthentication",
      data
    );
    return promise.then((data) =>
      MsgDeleteAuthenticationResponse.decode(new Reader(data))
    );
  }

  DeleteService(request: MsgDeleteService): Promise<MsgDeleteServiceResponse> {
    const data = MsgDeleteService.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.identifier.Msg",
      "DeleteService",
      data
    );
    return promise.then((data) =>
      MsgDeleteServiceResponse.decode(new Reader(data))
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
