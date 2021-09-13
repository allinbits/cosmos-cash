/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { VerificationMethod, Service } from "../did/did";

export const protobufPackage = "allinbits.cosmoscash.did";

/**
 * Verification is a message that allows to assign a verification method
 * to one or more verification relationships
 */
export interface Verification {
  /**
   * verificationRelationships defines which relationships
   * are allowed to use the verification method
   */
  relationships: string[];
  /** public key associated with the did document. */
  method: VerificationMethod | undefined;
  /** additional contexts (json ld schemas) */
  context: string[];
}

/** MsgCreateDidDocument defines a SDK message for creating a new did. */
export interface MsgCreateDidDocument {
  /** the did */
  id: string;
  /** the controller did */
  controller: string;
  /** the list of verification methods and relationships */
  verifications: Verification[];
  /** the list of services */
  services: Service[];
  /** address of the account signing the message */
  signer: string;
}

export interface MsgCreateDidDocumentResponse {}

export interface MsgUpdateDidDocument {
  /** the did */
  id: string;
  /** update controllers */
  controller: string[];
  /** address of the account signing the message */
  signer: string;
}

export interface MsgUpdateDidDocumentResponse {}

export interface MsgAddVerification {
  /** the did */
  id: string;
  /** the verification to add */
  verification: Verification | undefined;
  /** address of the account signing the message */
  signer: string;
}

export interface MsgAddVerificationResponse {}

export interface MsgSetVerificationRelationships {
  /** the did */
  id: string;
  /** the verification method id */
  methodId: string;
  /** the list of relationships to set */
  relationships: string[];
  /** address of the account signing the message */
  signer: string;
}

export interface MsgSetVerificationRelationshipsResponse {}

export interface MsgRevokeVerification {
  /** the did */
  id: string;
  /** the verification method id */
  methodId: string;
  /** address of the account signing the message */
  signer: string;
}

export interface MsgRevokeVerificationResponse {}

export interface MsgAddService {
  /** the did */
  id: string;
  /** the service data to add */
  serviceData: Service | undefined;
  /** address of the account signing the message */
  signer: string;
}

export interface MsgAddServiceResponse {}

export interface MsgDeleteService {
  /** the did */
  id: string;
  /** the service id */
  serviceId: string;
  /** address of the account signing the message */
  signer: string;
}

export interface MsgDeleteServiceResponse {}

const baseVerification: object = { relationships: "", context: "" };

export const Verification = {
  encode(message: Verification, writer: Writer = Writer.create()): Writer {
    for (const v of message.relationships) {
      writer.uint32(10).string(v!);
    }
    if (message.method !== undefined) {
      VerificationMethod.encode(
        message.method,
        writer.uint32(18).fork()
      ).ldelim();
    }
    for (const v of message.context) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Verification {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVerification } as Verification;
    message.relationships = [];
    message.context = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.relationships.push(reader.string());
          break;
        case 2:
          message.method = VerificationMethod.decode(reader, reader.uint32());
          break;
        case 3:
          message.context.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Verification {
    const message = { ...baseVerification } as Verification;
    message.relationships = [];
    message.context = [];
    if (object.relationships !== undefined && object.relationships !== null) {
      for (const e of object.relationships) {
        message.relationships.push(String(e));
      }
    }
    if (object.method !== undefined && object.method !== null) {
      message.method = VerificationMethod.fromJSON(object.method);
    } else {
      message.method = undefined;
    }
    if (object.context !== undefined && object.context !== null) {
      for (const e of object.context) {
        message.context.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: Verification): unknown {
    const obj: any = {};
    if (message.relationships) {
      obj.relationships = message.relationships.map((e) => e);
    } else {
      obj.relationships = [];
    }
    message.method !== undefined &&
      (obj.method = message.method
        ? VerificationMethod.toJSON(message.method)
        : undefined);
    if (message.context) {
      obj.context = message.context.map((e) => e);
    } else {
      obj.context = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Verification>): Verification {
    const message = { ...baseVerification } as Verification;
    message.relationships = [];
    message.context = [];
    if (object.relationships !== undefined && object.relationships !== null) {
      for (const e of object.relationships) {
        message.relationships.push(e);
      }
    }
    if (object.method !== undefined && object.method !== null) {
      message.method = VerificationMethod.fromPartial(object.method);
    } else {
      message.method = undefined;
    }
    if (object.context !== undefined && object.context !== null) {
      for (const e of object.context) {
        message.context.push(e);
      }
    }
    return message;
  },
};

const baseMsgCreateDidDocument: object = { id: "", controller: "", signer: "" };

export const MsgCreateDidDocument = {
  encode(
    message: MsgCreateDidDocument,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.controller !== "") {
      writer.uint32(18).string(message.controller);
    }
    for (const v of message.verifications) {
      Verification.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.services) {
      Service.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.signer !== "") {
      writer.uint32(42).string(message.signer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateDidDocument {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateDidDocument } as MsgCreateDidDocument;
    message.verifications = [];
    message.services = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.controller = reader.string();
          break;
        case 3:
          message.verifications.push(
            Verification.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.services.push(Service.decode(reader, reader.uint32()));
          break;
        case 5:
          message.signer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateDidDocument {
    const message = { ...baseMsgCreateDidDocument } as MsgCreateDidDocument;
    message.verifications = [];
    message.services = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.controller !== undefined && object.controller !== null) {
      message.controller = String(object.controller);
    } else {
      message.controller = "";
    }
    if (object.verifications !== undefined && object.verifications !== null) {
      for (const e of object.verifications) {
        message.verifications.push(Verification.fromJSON(e));
      }
    }
    if (object.services !== undefined && object.services !== null) {
      for (const e of object.services) {
        message.services.push(Service.fromJSON(e));
      }
    }
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer);
    } else {
      message.signer = "";
    }
    return message;
  },

  toJSON(message: MsgCreateDidDocument): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.controller !== undefined && (obj.controller = message.controller);
    if (message.verifications) {
      obj.verifications = message.verifications.map((e) =>
        e ? Verification.toJSON(e) : undefined
      );
    } else {
      obj.verifications = [];
    }
    if (message.services) {
      obj.services = message.services.map((e) =>
        e ? Service.toJSON(e) : undefined
      );
    } else {
      obj.services = [];
    }
    message.signer !== undefined && (obj.signer = message.signer);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateDidDocument>): MsgCreateDidDocument {
    const message = { ...baseMsgCreateDidDocument } as MsgCreateDidDocument;
    message.verifications = [];
    message.services = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.controller !== undefined && object.controller !== null) {
      message.controller = object.controller;
    } else {
      message.controller = "";
    }
    if (object.verifications !== undefined && object.verifications !== null) {
      for (const e of object.verifications) {
        message.verifications.push(Verification.fromPartial(e));
      }
    }
    if (object.services !== undefined && object.services !== null) {
      for (const e of object.services) {
        message.services.push(Service.fromPartial(e));
      }
    }
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer;
    } else {
      message.signer = "";
    }
    return message;
  },
};

const baseMsgCreateDidDocumentResponse: object = {};

export const MsgCreateDidDocumentResponse = {
  encode(
    _: MsgCreateDidDocumentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateDidDocumentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateDidDocumentResponse,
    } as MsgCreateDidDocumentResponse;
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

  fromJSON(_: any): MsgCreateDidDocumentResponse {
    const message = {
      ...baseMsgCreateDidDocumentResponse,
    } as MsgCreateDidDocumentResponse;
    return message;
  },

  toJSON(_: MsgCreateDidDocumentResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateDidDocumentResponse>
  ): MsgCreateDidDocumentResponse {
    const message = {
      ...baseMsgCreateDidDocumentResponse,
    } as MsgCreateDidDocumentResponse;
    return message;
  },
};

const baseMsgUpdateDidDocument: object = { id: "", controller: "", signer: "" };

export const MsgUpdateDidDocument = {
  encode(
    message: MsgUpdateDidDocument,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    for (const v of message.controller) {
      writer.uint32(18).string(v!);
    }
    if (message.signer !== "") {
      writer.uint32(42).string(message.signer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateDidDocument {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateDidDocument } as MsgUpdateDidDocument;
    message.controller = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.controller.push(reader.string());
          break;
        case 5:
          message.signer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateDidDocument {
    const message = { ...baseMsgUpdateDidDocument } as MsgUpdateDidDocument;
    message.controller = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.controller !== undefined && object.controller !== null) {
      for (const e of object.controller) {
        message.controller.push(String(e));
      }
    }
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer);
    } else {
      message.signer = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateDidDocument): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    if (message.controller) {
      obj.controller = message.controller.map((e) => e);
    } else {
      obj.controller = [];
    }
    message.signer !== undefined && (obj.signer = message.signer);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateDidDocument>): MsgUpdateDidDocument {
    const message = { ...baseMsgUpdateDidDocument } as MsgUpdateDidDocument;
    message.controller = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.controller !== undefined && object.controller !== null) {
      for (const e of object.controller) {
        message.controller.push(e);
      }
    }
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer;
    } else {
      message.signer = "";
    }
    return message;
  },
};

const baseMsgUpdateDidDocumentResponse: object = {};

export const MsgUpdateDidDocumentResponse = {
  encode(
    _: MsgUpdateDidDocumentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateDidDocumentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateDidDocumentResponse,
    } as MsgUpdateDidDocumentResponse;
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

  fromJSON(_: any): MsgUpdateDidDocumentResponse {
    const message = {
      ...baseMsgUpdateDidDocumentResponse,
    } as MsgUpdateDidDocumentResponse;
    return message;
  },

  toJSON(_: MsgUpdateDidDocumentResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateDidDocumentResponse>
  ): MsgUpdateDidDocumentResponse {
    const message = {
      ...baseMsgUpdateDidDocumentResponse,
    } as MsgUpdateDidDocumentResponse;
    return message;
  },
};

const baseMsgAddVerification: object = { id: "", signer: "" };

export const MsgAddVerification = {
  encode(
    message: MsgAddVerification,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.verification !== undefined) {
      Verification.encode(
        message.verification,
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.signer !== "") {
      writer.uint32(26).string(message.signer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddVerification {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddVerification } as MsgAddVerification;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.verification = Verification.decode(reader, reader.uint32());
          break;
        case 3:
          message.signer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddVerification {
    const message = { ...baseMsgAddVerification } as MsgAddVerification;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.verification !== undefined && object.verification !== null) {
      message.verification = Verification.fromJSON(object.verification);
    } else {
      message.verification = undefined;
    }
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer);
    } else {
      message.signer = "";
    }
    return message;
  },

  toJSON(message: MsgAddVerification): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.verification !== undefined &&
      (obj.verification = message.verification
        ? Verification.toJSON(message.verification)
        : undefined);
    message.signer !== undefined && (obj.signer = message.signer);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddVerification>): MsgAddVerification {
    const message = { ...baseMsgAddVerification } as MsgAddVerification;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.verification !== undefined && object.verification !== null) {
      message.verification = Verification.fromPartial(object.verification);
    } else {
      message.verification = undefined;
    }
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer;
    } else {
      message.signer = "";
    }
    return message;
  },
};

const baseMsgAddVerificationResponse: object = {};

export const MsgAddVerificationResponse = {
  encode(
    _: MsgAddVerificationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddVerificationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddVerificationResponse,
    } as MsgAddVerificationResponse;
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

  fromJSON(_: any): MsgAddVerificationResponse {
    const message = {
      ...baseMsgAddVerificationResponse,
    } as MsgAddVerificationResponse;
    return message;
  },

  toJSON(_: MsgAddVerificationResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgAddVerificationResponse>
  ): MsgAddVerificationResponse {
    const message = {
      ...baseMsgAddVerificationResponse,
    } as MsgAddVerificationResponse;
    return message;
  },
};

const baseMsgSetVerificationRelationships: object = {
  id: "",
  methodId: "",
  relationships: "",
  signer: "",
};

export const MsgSetVerificationRelationships = {
  encode(
    message: MsgSetVerificationRelationships,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.methodId !== "") {
      writer.uint32(18).string(message.methodId);
    }
    for (const v of message.relationships) {
      writer.uint32(26).string(v!);
    }
    if (message.signer !== "") {
      writer.uint32(34).string(message.signer);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSetVerificationRelationships {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSetVerificationRelationships,
    } as MsgSetVerificationRelationships;
    message.relationships = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.methodId = reader.string();
          break;
        case 3:
          message.relationships.push(reader.string());
          break;
        case 4:
          message.signer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetVerificationRelationships {
    const message = {
      ...baseMsgSetVerificationRelationships,
    } as MsgSetVerificationRelationships;
    message.relationships = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.methodId !== undefined && object.methodId !== null) {
      message.methodId = String(object.methodId);
    } else {
      message.methodId = "";
    }
    if (object.relationships !== undefined && object.relationships !== null) {
      for (const e of object.relationships) {
        message.relationships.push(String(e));
      }
    }
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer);
    } else {
      message.signer = "";
    }
    return message;
  },

  toJSON(message: MsgSetVerificationRelationships): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.methodId !== undefined && (obj.methodId = message.methodId);
    if (message.relationships) {
      obj.relationships = message.relationships.map((e) => e);
    } else {
      obj.relationships = [];
    }
    message.signer !== undefined && (obj.signer = message.signer);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSetVerificationRelationships>
  ): MsgSetVerificationRelationships {
    const message = {
      ...baseMsgSetVerificationRelationships,
    } as MsgSetVerificationRelationships;
    message.relationships = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.methodId !== undefined && object.methodId !== null) {
      message.methodId = object.methodId;
    } else {
      message.methodId = "";
    }
    if (object.relationships !== undefined && object.relationships !== null) {
      for (const e of object.relationships) {
        message.relationships.push(e);
      }
    }
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer;
    } else {
      message.signer = "";
    }
    return message;
  },
};

const baseMsgSetVerificationRelationshipsResponse: object = {};

export const MsgSetVerificationRelationshipsResponse = {
  encode(
    _: MsgSetVerificationRelationshipsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSetVerificationRelationshipsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSetVerificationRelationshipsResponse,
    } as MsgSetVerificationRelationshipsResponse;
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

  fromJSON(_: any): MsgSetVerificationRelationshipsResponse {
    const message = {
      ...baseMsgSetVerificationRelationshipsResponse,
    } as MsgSetVerificationRelationshipsResponse;
    return message;
  },

  toJSON(_: MsgSetVerificationRelationshipsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSetVerificationRelationshipsResponse>
  ): MsgSetVerificationRelationshipsResponse {
    const message = {
      ...baseMsgSetVerificationRelationshipsResponse,
    } as MsgSetVerificationRelationshipsResponse;
    return message;
  },
};

const baseMsgRevokeVerification: object = { id: "", methodId: "", signer: "" };

export const MsgRevokeVerification = {
  encode(
    message: MsgRevokeVerification,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.methodId !== "") {
      writer.uint32(18).string(message.methodId);
    }
    if (message.signer !== "") {
      writer.uint32(26).string(message.signer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRevokeVerification {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRevokeVerification } as MsgRevokeVerification;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.methodId = reader.string();
          break;
        case 3:
          message.signer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRevokeVerification {
    const message = { ...baseMsgRevokeVerification } as MsgRevokeVerification;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.methodId !== undefined && object.methodId !== null) {
      message.methodId = String(object.methodId);
    } else {
      message.methodId = "";
    }
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer);
    } else {
      message.signer = "";
    }
    return message;
  },

  toJSON(message: MsgRevokeVerification): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.methodId !== undefined && (obj.methodId = message.methodId);
    message.signer !== undefined && (obj.signer = message.signer);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRevokeVerification>
  ): MsgRevokeVerification {
    const message = { ...baseMsgRevokeVerification } as MsgRevokeVerification;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.methodId !== undefined && object.methodId !== null) {
      message.methodId = object.methodId;
    } else {
      message.methodId = "";
    }
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer;
    } else {
      message.signer = "";
    }
    return message;
  },
};

const baseMsgRevokeVerificationResponse: object = {};

export const MsgRevokeVerificationResponse = {
  encode(
    _: MsgRevokeVerificationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRevokeVerificationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRevokeVerificationResponse,
    } as MsgRevokeVerificationResponse;
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

  fromJSON(_: any): MsgRevokeVerificationResponse {
    const message = {
      ...baseMsgRevokeVerificationResponse,
    } as MsgRevokeVerificationResponse;
    return message;
  },

  toJSON(_: MsgRevokeVerificationResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRevokeVerificationResponse>
  ): MsgRevokeVerificationResponse {
    const message = {
      ...baseMsgRevokeVerificationResponse,
    } as MsgRevokeVerificationResponse;
    return message;
  },
};

const baseMsgAddService: object = { id: "", signer: "" };

export const MsgAddService = {
  encode(message: MsgAddService, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.serviceData !== undefined) {
      Service.encode(message.serviceData, writer.uint32(18).fork()).ldelim();
    }
    if (message.signer !== "") {
      writer.uint32(26).string(message.signer);
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
          message.signer = reader.string();
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
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer);
    } else {
      message.signer = "";
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
    message.signer !== undefined && (obj.signer = message.signer);
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
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer;
    } else {
      message.signer = "";
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

const baseMsgDeleteService: object = { id: "", serviceId: "", signer: "" };

export const MsgDeleteService = {
  encode(message: MsgDeleteService, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.serviceId !== "") {
      writer.uint32(18).string(message.serviceId);
    }
    if (message.signer !== "") {
      writer.uint32(26).string(message.signer);
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
          message.signer = reader.string();
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
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer);
    } else {
      message.signer = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteService): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.serviceId !== undefined && (obj.serviceId = message.serviceId);
    message.signer !== undefined && (obj.signer = message.signer);
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
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer;
    } else {
      message.signer = "";
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
  CreateDidDocument(
    request: MsgCreateDidDocument
  ): Promise<MsgCreateDidDocumentResponse>;
  /** UpdateDidDocument defines a method for updating an identity. */
  UpdateDidDocument(
    request: MsgUpdateDidDocument
  ): Promise<MsgUpdateDidDocumentResponse>;
  /** AddVerificationMethod adds a new verification method */
  AddVerification(
    request: MsgAddVerification
  ): Promise<MsgAddVerificationResponse>;
  /** RevokeVerification remove the verification method and all associated verification Relations */
  RevokeVerification(
    request: MsgRevokeVerification
  ): Promise<MsgRevokeVerificationResponse>;
  /** SetVerificationRelationships overwrite current verification relationships */
  SetVerificationRelationships(
    request: MsgSetVerificationRelationships
  ): Promise<MsgSetVerificationRelationshipsResponse>;
  /** AddService add a new service */
  AddService(request: MsgAddService): Promise<MsgAddServiceResponse>;
  /** DeleteService delete an existing service */
  DeleteService(request: MsgDeleteService): Promise<MsgDeleteServiceResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateDidDocument(
    request: MsgCreateDidDocument
  ): Promise<MsgCreateDidDocumentResponse> {
    const data = MsgCreateDidDocument.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.did.Msg",
      "CreateDidDocument",
      data
    );
    return promise.then((data) =>
      MsgCreateDidDocumentResponse.decode(new Reader(data))
    );
  }

  UpdateDidDocument(
    request: MsgUpdateDidDocument
  ): Promise<MsgUpdateDidDocumentResponse> {
    const data = MsgUpdateDidDocument.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.did.Msg",
      "UpdateDidDocument",
      data
    );
    return promise.then((data) =>
      MsgUpdateDidDocumentResponse.decode(new Reader(data))
    );
  }

  AddVerification(
    request: MsgAddVerification
  ): Promise<MsgAddVerificationResponse> {
    const data = MsgAddVerification.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.did.Msg",
      "AddVerification",
      data
    );
    return promise.then((data) =>
      MsgAddVerificationResponse.decode(new Reader(data))
    );
  }

  RevokeVerification(
    request: MsgRevokeVerification
  ): Promise<MsgRevokeVerificationResponse> {
    const data = MsgRevokeVerification.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.did.Msg",
      "RevokeVerification",
      data
    );
    return promise.then((data) =>
      MsgRevokeVerificationResponse.decode(new Reader(data))
    );
  }

  SetVerificationRelationships(
    request: MsgSetVerificationRelationships
  ): Promise<MsgSetVerificationRelationshipsResponse> {
    const data = MsgSetVerificationRelationships.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.did.Msg",
      "SetVerificationRelationships",
      data
    );
    return promise.then((data) =>
      MsgSetVerificationRelationshipsResponse.decode(new Reader(data))
    );
  }

  AddService(request: MsgAddService): Promise<MsgAddServiceResponse> {
    const data = MsgAddService.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.did.Msg",
      "AddService",
      data
    );
    return promise.then((data) =>
      MsgAddServiceResponse.decode(new Reader(data))
    );
  }

  DeleteService(request: MsgDeleteService): Promise<MsgDeleteServiceResponse> {
    const data = MsgDeleteService.encode(request).finish();
    const promise = this.rpc.request(
      "allinbits.cosmoscash.did.Msg",
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
