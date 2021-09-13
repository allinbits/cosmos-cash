/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "allinbits.cosmoscash.did";

/** DidDocument represents a dencentralised identifer. */
export interface DidDocument {
  /** @context is spec for did document. */
  context: string[];
  /** id represents the id for the did document. */
  id: string;
  /**
   * A DID controller is an entity that is authorized to make changes to a DID document.
   * cfr. https://www.w3.org/TR/did-core/#did-controller
   */
  controller: string[];
  /**
   * A DID document can express verification methods,
   * such as cryptographic public keys, which can be used
   * to authenticate or authorize interactions with the DID subject or associated parties.
   * https://www.w3.org/TR/did-core/#verification-methods
   */
  verificationMethod: VerificationMethod[];
  /**
   * Services are used in DID documents to express ways of communicating
   * with the DID subject or associated entities.
   * https://www.w3.org/TR/did-core/#services
   */
  service: Service[];
  /**
   * Authentication represents public key associated with the did document.
   * cfr. https://www.w3.org/TR/did-core/#authentication
   */
  authentication: string[];
  /**
   * Used to specify how the DID subject is expected to express claims,
   * such as for the purposes of issuing a Verifiable Credential.
   * cfr. https://www.w3.org/TR/did-core/#assertion
   */
  assertionMethod: string[];
  /**
   * used to specify how an entity can generate encryption material
   * in order to transmit confidential information intended for the DID subject.
   * https://www.w3.org/TR/did-core/#key-agreement
   */
  keyAgreement: string[];
  /**
   * Used to specify a verification method that might be used by the DID subject
   * to invoke a cryptographic capability, such as the authorization
   * to update the DID Document.
   * https://www.w3.org/TR/did-core/#capability-invocation
   */
  capabilityInvocation: string[];
  /**
   * Used to specify a mechanism that might be used by the DID subject
   * to delegate a cryptographic capability to another party.
   * https://www.w3.org/TR/did-core/#capability-delegation
   */
  capabilityDelegation: string[];
}

/**
 * A DID document can express verification methods,
 * such as cryptographic public keys, which can be used
 * to authenticate or authorize interactions
 * with the DID subject or associated parties.
 * https://www.w3.org/TR/did-core/#verification-methods
 */
export interface VerificationMethod {
  id: string;
  type: string;
  controller: string;
  blockchainAccountID: string | undefined;
  publicKeyHex: string | undefined;
  publicKeyMultibase: string | undefined;
}

/** Service defines how to find data associated with a identifer */
export interface Service {
  id: string;
  type: string;
  serviceEndpoint: string;
}

/**
 * DidMetadata defines metadata associated to a did document such as
 * the status of the DID document
 */
export interface DidMetadata {
  versionId: string;
  created: Date | undefined;
  updated: Date | undefined;
  deactivated: boolean;
}

const baseDidDocument: object = {
  context: "",
  id: "",
  controller: "",
  authentication: "",
  assertionMethod: "",
  keyAgreement: "",
  capabilityInvocation: "",
  capabilityDelegation: "",
};

export const DidDocument = {
  encode(message: DidDocument, writer: Writer = Writer.create()): Writer {
    for (const v of message.context) {
      writer.uint32(10).string(v!);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    for (const v of message.controller) {
      writer.uint32(26).string(v!);
    }
    for (const v of message.verificationMethod) {
      VerificationMethod.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.service) {
      Service.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.authentication) {
      writer.uint32(50).string(v!);
    }
    for (const v of message.assertionMethod) {
      writer.uint32(58).string(v!);
    }
    for (const v of message.keyAgreement) {
      writer.uint32(66).string(v!);
    }
    for (const v of message.capabilityInvocation) {
      writer.uint32(74).string(v!);
    }
    for (const v of message.capabilityDelegation) {
      writer.uint32(82).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DidDocument {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDidDocument } as DidDocument;
    message.context = [];
    message.controller = [];
    message.verificationMethod = [];
    message.service = [];
    message.authentication = [];
    message.assertionMethod = [];
    message.keyAgreement = [];
    message.capabilityInvocation = [];
    message.capabilityDelegation = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.context.push(reader.string());
          break;
        case 2:
          message.id = reader.string();
          break;
        case 3:
          message.controller.push(reader.string());
          break;
        case 4:
          message.verificationMethod.push(
            VerificationMethod.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.service.push(Service.decode(reader, reader.uint32()));
          break;
        case 6:
          message.authentication.push(reader.string());
          break;
        case 7:
          message.assertionMethod.push(reader.string());
          break;
        case 8:
          message.keyAgreement.push(reader.string());
          break;
        case 9:
          message.capabilityInvocation.push(reader.string());
          break;
        case 10:
          message.capabilityDelegation.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DidDocument {
    const message = { ...baseDidDocument } as DidDocument;
    message.context = [];
    message.controller = [];
    message.verificationMethod = [];
    message.service = [];
    message.authentication = [];
    message.assertionMethod = [];
    message.keyAgreement = [];
    message.capabilityInvocation = [];
    message.capabilityDelegation = [];
    if (object.context !== undefined && object.context !== null) {
      for (const e of object.context) {
        message.context.push(String(e));
      }
    }
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
    if (
      object.verificationMethod !== undefined &&
      object.verificationMethod !== null
    ) {
      for (const e of object.verificationMethod) {
        message.verificationMethod.push(VerificationMethod.fromJSON(e));
      }
    }
    if (object.service !== undefined && object.service !== null) {
      for (const e of object.service) {
        message.service.push(Service.fromJSON(e));
      }
    }
    if (object.authentication !== undefined && object.authentication !== null) {
      for (const e of object.authentication) {
        message.authentication.push(String(e));
      }
    }
    if (
      object.assertionMethod !== undefined &&
      object.assertionMethod !== null
    ) {
      for (const e of object.assertionMethod) {
        message.assertionMethod.push(String(e));
      }
    }
    if (object.keyAgreement !== undefined && object.keyAgreement !== null) {
      for (const e of object.keyAgreement) {
        message.keyAgreement.push(String(e));
      }
    }
    if (
      object.capabilityInvocation !== undefined &&
      object.capabilityInvocation !== null
    ) {
      for (const e of object.capabilityInvocation) {
        message.capabilityInvocation.push(String(e));
      }
    }
    if (
      object.capabilityDelegation !== undefined &&
      object.capabilityDelegation !== null
    ) {
      for (const e of object.capabilityDelegation) {
        message.capabilityDelegation.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: DidDocument): unknown {
    const obj: any = {};
    if (message.context) {
      obj.context = message.context.map((e) => e);
    } else {
      obj.context = [];
    }
    message.id !== undefined && (obj.id = message.id);
    if (message.controller) {
      obj.controller = message.controller.map((e) => e);
    } else {
      obj.controller = [];
    }
    if (message.verificationMethod) {
      obj.verificationMethod = message.verificationMethod.map((e) =>
        e ? VerificationMethod.toJSON(e) : undefined
      );
    } else {
      obj.verificationMethod = [];
    }
    if (message.service) {
      obj.service = message.service.map((e) =>
        e ? Service.toJSON(e) : undefined
      );
    } else {
      obj.service = [];
    }
    if (message.authentication) {
      obj.authentication = message.authentication.map((e) => e);
    } else {
      obj.authentication = [];
    }
    if (message.assertionMethod) {
      obj.assertionMethod = message.assertionMethod.map((e) => e);
    } else {
      obj.assertionMethod = [];
    }
    if (message.keyAgreement) {
      obj.keyAgreement = message.keyAgreement.map((e) => e);
    } else {
      obj.keyAgreement = [];
    }
    if (message.capabilityInvocation) {
      obj.capabilityInvocation = message.capabilityInvocation.map((e) => e);
    } else {
      obj.capabilityInvocation = [];
    }
    if (message.capabilityDelegation) {
      obj.capabilityDelegation = message.capabilityDelegation.map((e) => e);
    } else {
      obj.capabilityDelegation = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<DidDocument>): DidDocument {
    const message = { ...baseDidDocument } as DidDocument;
    message.context = [];
    message.controller = [];
    message.verificationMethod = [];
    message.service = [];
    message.authentication = [];
    message.assertionMethod = [];
    message.keyAgreement = [];
    message.capabilityInvocation = [];
    message.capabilityDelegation = [];
    if (object.context !== undefined && object.context !== null) {
      for (const e of object.context) {
        message.context.push(e);
      }
    }
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
    if (
      object.verificationMethod !== undefined &&
      object.verificationMethod !== null
    ) {
      for (const e of object.verificationMethod) {
        message.verificationMethod.push(VerificationMethod.fromPartial(e));
      }
    }
    if (object.service !== undefined && object.service !== null) {
      for (const e of object.service) {
        message.service.push(Service.fromPartial(e));
      }
    }
    if (object.authentication !== undefined && object.authentication !== null) {
      for (const e of object.authentication) {
        message.authentication.push(e);
      }
    }
    if (
      object.assertionMethod !== undefined &&
      object.assertionMethod !== null
    ) {
      for (const e of object.assertionMethod) {
        message.assertionMethod.push(e);
      }
    }
    if (object.keyAgreement !== undefined && object.keyAgreement !== null) {
      for (const e of object.keyAgreement) {
        message.keyAgreement.push(e);
      }
    }
    if (
      object.capabilityInvocation !== undefined &&
      object.capabilityInvocation !== null
    ) {
      for (const e of object.capabilityInvocation) {
        message.capabilityInvocation.push(e);
      }
    }
    if (
      object.capabilityDelegation !== undefined &&
      object.capabilityDelegation !== null
    ) {
      for (const e of object.capabilityDelegation) {
        message.capabilityDelegation.push(e);
      }
    }
    return message;
  },
};

const baseVerificationMethod: object = { id: "", type: "", controller: "" };

export const VerificationMethod = {
  encode(
    message: VerificationMethod,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.type !== "") {
      writer.uint32(18).string(message.type);
    }
    if (message.controller !== "") {
      writer.uint32(26).string(message.controller);
    }
    if (message.blockchainAccountID !== undefined) {
      writer.uint32(34).string(message.blockchainAccountID);
    }
    if (message.publicKeyHex !== undefined) {
      writer.uint32(42).string(message.publicKeyHex);
    }
    if (message.publicKeyMultibase !== undefined) {
      writer.uint32(50).string(message.publicKeyMultibase);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VerificationMethod {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVerificationMethod } as VerificationMethod;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.type = reader.string();
          break;
        case 3:
          message.controller = reader.string();
          break;
        case 4:
          message.blockchainAccountID = reader.string();
          break;
        case 5:
          message.publicKeyHex = reader.string();
          break;
        case 6:
          message.publicKeyMultibase = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VerificationMethod {
    const message = { ...baseVerificationMethod } as VerificationMethod;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = String(object.type);
    } else {
      message.type = "";
    }
    if (object.controller !== undefined && object.controller !== null) {
      message.controller = String(object.controller);
    } else {
      message.controller = "";
    }
    if (
      object.blockchainAccountID !== undefined &&
      object.blockchainAccountID !== null
    ) {
      message.blockchainAccountID = String(object.blockchainAccountID);
    } else {
      message.blockchainAccountID = undefined;
    }
    if (object.publicKeyHex !== undefined && object.publicKeyHex !== null) {
      message.publicKeyHex = String(object.publicKeyHex);
    } else {
      message.publicKeyHex = undefined;
    }
    if (
      object.publicKeyMultibase !== undefined &&
      object.publicKeyMultibase !== null
    ) {
      message.publicKeyMultibase = String(object.publicKeyMultibase);
    } else {
      message.publicKeyMultibase = undefined;
    }
    return message;
  },

  toJSON(message: VerificationMethod): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.type !== undefined && (obj.type = message.type);
    message.controller !== undefined && (obj.controller = message.controller);
    message.blockchainAccountID !== undefined &&
      (obj.blockchainAccountID = message.blockchainAccountID);
    message.publicKeyHex !== undefined &&
      (obj.publicKeyHex = message.publicKeyHex);
    message.publicKeyMultibase !== undefined &&
      (obj.publicKeyMultibase = message.publicKeyMultibase);
    return obj;
  },

  fromPartial(object: DeepPartial<VerificationMethod>): VerificationMethod {
    const message = { ...baseVerificationMethod } as VerificationMethod;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = "";
    }
    if (object.controller !== undefined && object.controller !== null) {
      message.controller = object.controller;
    } else {
      message.controller = "";
    }
    if (
      object.blockchainAccountID !== undefined &&
      object.blockchainAccountID !== null
    ) {
      message.blockchainAccountID = object.blockchainAccountID;
    } else {
      message.blockchainAccountID = undefined;
    }
    if (object.publicKeyHex !== undefined && object.publicKeyHex !== null) {
      message.publicKeyHex = object.publicKeyHex;
    } else {
      message.publicKeyHex = undefined;
    }
    if (
      object.publicKeyMultibase !== undefined &&
      object.publicKeyMultibase !== null
    ) {
      message.publicKeyMultibase = object.publicKeyMultibase;
    } else {
      message.publicKeyMultibase = undefined;
    }
    return message;
  },
};

const baseService: object = { id: "", type: "", serviceEndpoint: "" };

export const Service = {
  encode(message: Service, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.type !== "") {
      writer.uint32(18).string(message.type);
    }
    if (message.serviceEndpoint !== "") {
      writer.uint32(26).string(message.serviceEndpoint);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Service {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseService } as Service;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.type = reader.string();
          break;
        case 3:
          message.serviceEndpoint = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Service {
    const message = { ...baseService } as Service;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = String(object.type);
    } else {
      message.type = "";
    }
    if (
      object.serviceEndpoint !== undefined &&
      object.serviceEndpoint !== null
    ) {
      message.serviceEndpoint = String(object.serviceEndpoint);
    } else {
      message.serviceEndpoint = "";
    }
    return message;
  },

  toJSON(message: Service): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.type !== undefined && (obj.type = message.type);
    message.serviceEndpoint !== undefined &&
      (obj.serviceEndpoint = message.serviceEndpoint);
    return obj;
  },

  fromPartial(object: DeepPartial<Service>): Service {
    const message = { ...baseService } as Service;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = "";
    }
    if (
      object.serviceEndpoint !== undefined &&
      object.serviceEndpoint !== null
    ) {
      message.serviceEndpoint = object.serviceEndpoint;
    } else {
      message.serviceEndpoint = "";
    }
    return message;
  },
};

const baseDidMetadata: object = { versionId: "", deactivated: false };

export const DidMetadata = {
  encode(message: DidMetadata, writer: Writer = Writer.create()): Writer {
    if (message.versionId !== "") {
      writer.uint32(10).string(message.versionId);
    }
    if (message.created !== undefined) {
      Timestamp.encode(
        toTimestamp(message.created),
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.updated !== undefined) {
      Timestamp.encode(
        toTimestamp(message.updated),
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.deactivated === true) {
      writer.uint32(32).bool(message.deactivated);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DidMetadata {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDidMetadata } as DidMetadata;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.versionId = reader.string();
          break;
        case 2:
          message.created = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.updated = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.deactivated = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DidMetadata {
    const message = { ...baseDidMetadata } as DidMetadata;
    if (object.versionId !== undefined && object.versionId !== null) {
      message.versionId = String(object.versionId);
    } else {
      message.versionId = "";
    }
    if (object.created !== undefined && object.created !== null) {
      message.created = fromJsonTimestamp(object.created);
    } else {
      message.created = undefined;
    }
    if (object.updated !== undefined && object.updated !== null) {
      message.updated = fromJsonTimestamp(object.updated);
    } else {
      message.updated = undefined;
    }
    if (object.deactivated !== undefined && object.deactivated !== null) {
      message.deactivated = Boolean(object.deactivated);
    } else {
      message.deactivated = false;
    }
    return message;
  },

  toJSON(message: DidMetadata): unknown {
    const obj: any = {};
    message.versionId !== undefined && (obj.versionId = message.versionId);
    message.created !== undefined &&
      (obj.created =
        message.created !== undefined ? message.created.toISOString() : null);
    message.updated !== undefined &&
      (obj.updated =
        message.updated !== undefined ? message.updated.toISOString() : null);
    message.deactivated !== undefined &&
      (obj.deactivated = message.deactivated);
    return obj;
  },

  fromPartial(object: DeepPartial<DidMetadata>): DidMetadata {
    const message = { ...baseDidMetadata } as DidMetadata;
    if (object.versionId !== undefined && object.versionId !== null) {
      message.versionId = object.versionId;
    } else {
      message.versionId = "";
    }
    if (object.created !== undefined && object.created !== null) {
      message.created = object.created;
    } else {
      message.created = undefined;
    }
    if (object.updated !== undefined && object.updated !== null) {
      message.updated = object.updated;
    } else {
      message.updated = undefined;
    }
    if (object.deactivated !== undefined && object.deactivated !== null) {
      message.deactivated = object.deactivated;
    } else {
      message.deactivated = false;
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}
