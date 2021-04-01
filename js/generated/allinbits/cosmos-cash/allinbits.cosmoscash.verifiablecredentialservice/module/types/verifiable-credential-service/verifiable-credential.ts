/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage =
  "allinbits.cosmoscash.verifiablecredentialservice";

export interface VerifiableCredential {
  context: string[];
  id: string;
  type: string[];
  issuer: string;
  issuanceDate: string;
  userCred: UserCredentialSubject | undefined;
  issuerCred: IssuerCredentialSubject | undefined;
  proof: Proof | undefined;
}

export interface UserCredentialSubject {
  id: string;
  hasKyc: boolean;
}

export interface IssuerCredentialSubject {
  id: string;
  isVerified: boolean;
}

export interface Proof {
  type: string;
  created: string;
  proofPurpose: string;
  verificationMethod: string;
  signature: string;
}

const baseVerifiableCredential: object = {
  context: "",
  id: "",
  type: "",
  issuer: "",
  issuanceDate: "",
};

export const VerifiableCredential = {
  encode(
    message: VerifiableCredential,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.context) {
      writer.uint32(10).string(v!);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    for (const v of message.type) {
      writer.uint32(26).string(v!);
    }
    if (message.issuer !== "") {
      writer.uint32(34).string(message.issuer);
    }
    if (message.issuanceDate !== "") {
      writer.uint32(42).string(message.issuanceDate);
    }
    if (message.userCred !== undefined) {
      UserCredentialSubject.encode(
        message.userCred,
        writer.uint32(50).fork()
      ).ldelim();
    }
    if (message.issuerCred !== undefined) {
      IssuerCredentialSubject.encode(
        message.issuerCred,
        writer.uint32(58).fork()
      ).ldelim();
    }
    if (message.proof !== undefined) {
      Proof.encode(message.proof, writer.uint32(66).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VerifiableCredential {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVerifiableCredential } as VerifiableCredential;
    message.context = [];
    message.type = [];
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
          message.type.push(reader.string());
          break;
        case 4:
          message.issuer = reader.string();
          break;
        case 5:
          message.issuanceDate = reader.string();
          break;
        case 6:
          message.userCred = UserCredentialSubject.decode(
            reader,
            reader.uint32()
          );
          break;
        case 7:
          message.issuerCred = IssuerCredentialSubject.decode(
            reader,
            reader.uint32()
          );
          break;
        case 8:
          message.proof = Proof.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VerifiableCredential {
    const message = { ...baseVerifiableCredential } as VerifiableCredential;
    message.context = [];
    message.type = [];
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
    if (object.type !== undefined && object.type !== null) {
      for (const e of object.type) {
        message.type.push(String(e));
      }
    }
    if (object.issuer !== undefined && object.issuer !== null) {
      message.issuer = String(object.issuer);
    } else {
      message.issuer = "";
    }
    if (object.issuanceDate !== undefined && object.issuanceDate !== null) {
      message.issuanceDate = String(object.issuanceDate);
    } else {
      message.issuanceDate = "";
    }
    if (object.userCred !== undefined && object.userCred !== null) {
      message.userCred = UserCredentialSubject.fromJSON(object.userCred);
    } else {
      message.userCred = undefined;
    }
    if (object.issuerCred !== undefined && object.issuerCred !== null) {
      message.issuerCred = IssuerCredentialSubject.fromJSON(object.issuerCred);
    } else {
      message.issuerCred = undefined;
    }
    if (object.proof !== undefined && object.proof !== null) {
      message.proof = Proof.fromJSON(object.proof);
    } else {
      message.proof = undefined;
    }
    return message;
  },

  toJSON(message: VerifiableCredential): unknown {
    const obj: any = {};
    if (message.context) {
      obj.context = message.context.map((e) => e);
    } else {
      obj.context = [];
    }
    message.id !== undefined && (obj.id = message.id);
    if (message.type) {
      obj.type = message.type.map((e) => e);
    } else {
      obj.type = [];
    }
    message.issuer !== undefined && (obj.issuer = message.issuer);
    message.issuanceDate !== undefined &&
      (obj.issuanceDate = message.issuanceDate);
    message.userCred !== undefined &&
      (obj.userCred = message.userCred
        ? UserCredentialSubject.toJSON(message.userCred)
        : undefined);
    message.issuerCred !== undefined &&
      (obj.issuerCred = message.issuerCred
        ? IssuerCredentialSubject.toJSON(message.issuerCred)
        : undefined);
    message.proof !== undefined &&
      (obj.proof = message.proof ? Proof.toJSON(message.proof) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<VerifiableCredential>): VerifiableCredential {
    const message = { ...baseVerifiableCredential } as VerifiableCredential;
    message.context = [];
    message.type = [];
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
    if (object.type !== undefined && object.type !== null) {
      for (const e of object.type) {
        message.type.push(e);
      }
    }
    if (object.issuer !== undefined && object.issuer !== null) {
      message.issuer = object.issuer;
    } else {
      message.issuer = "";
    }
    if (object.issuanceDate !== undefined && object.issuanceDate !== null) {
      message.issuanceDate = object.issuanceDate;
    } else {
      message.issuanceDate = "";
    }
    if (object.userCred !== undefined && object.userCred !== null) {
      message.userCred = UserCredentialSubject.fromPartial(object.userCred);
    } else {
      message.userCred = undefined;
    }
    if (object.issuerCred !== undefined && object.issuerCred !== null) {
      message.issuerCred = IssuerCredentialSubject.fromPartial(
        object.issuerCred
      );
    } else {
      message.issuerCred = undefined;
    }
    if (object.proof !== undefined && object.proof !== null) {
      message.proof = Proof.fromPartial(object.proof);
    } else {
      message.proof = undefined;
    }
    return message;
  },
};

const baseUserCredentialSubject: object = { id: "", hasKyc: false };

export const UserCredentialSubject = {
  encode(
    message: UserCredentialSubject,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.hasKyc === true) {
      writer.uint32(16).bool(message.hasKyc);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): UserCredentialSubject {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUserCredentialSubject } as UserCredentialSubject;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.hasKyc = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserCredentialSubject {
    const message = { ...baseUserCredentialSubject } as UserCredentialSubject;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.hasKyc !== undefined && object.hasKyc !== null) {
      message.hasKyc = Boolean(object.hasKyc);
    } else {
      message.hasKyc = false;
    }
    return message;
  },

  toJSON(message: UserCredentialSubject): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.hasKyc !== undefined && (obj.hasKyc = message.hasKyc);
    return obj;
  },

  fromPartial(
    object: DeepPartial<UserCredentialSubject>
  ): UserCredentialSubject {
    const message = { ...baseUserCredentialSubject } as UserCredentialSubject;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.hasKyc !== undefined && object.hasKyc !== null) {
      message.hasKyc = object.hasKyc;
    } else {
      message.hasKyc = false;
    }
    return message;
  },
};

const baseIssuerCredentialSubject: object = { id: "", isVerified: false };

export const IssuerCredentialSubject = {
  encode(
    message: IssuerCredentialSubject,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.isVerified === true) {
      writer.uint32(16).bool(message.isVerified);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): IssuerCredentialSubject {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseIssuerCredentialSubject,
    } as IssuerCredentialSubject;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.isVerified = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IssuerCredentialSubject {
    const message = {
      ...baseIssuerCredentialSubject,
    } as IssuerCredentialSubject;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.isVerified !== undefined && object.isVerified !== null) {
      message.isVerified = Boolean(object.isVerified);
    } else {
      message.isVerified = false;
    }
    return message;
  },

  toJSON(message: IssuerCredentialSubject): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.isVerified !== undefined && (obj.isVerified = message.isVerified);
    return obj;
  },

  fromPartial(
    object: DeepPartial<IssuerCredentialSubject>
  ): IssuerCredentialSubject {
    const message = {
      ...baseIssuerCredentialSubject,
    } as IssuerCredentialSubject;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.isVerified !== undefined && object.isVerified !== null) {
      message.isVerified = object.isVerified;
    } else {
      message.isVerified = false;
    }
    return message;
  },
};

const baseProof: object = {
  type: "",
  created: "",
  proofPurpose: "",
  verificationMethod: "",
  signature: "",
};

export const Proof = {
  encode(message: Proof, writer: Writer = Writer.create()): Writer {
    if (message.type !== "") {
      writer.uint32(10).string(message.type);
    }
    if (message.created !== "") {
      writer.uint32(18).string(message.created);
    }
    if (message.proofPurpose !== "") {
      writer.uint32(26).string(message.proofPurpose);
    }
    if (message.verificationMethod !== "") {
      writer.uint32(34).string(message.verificationMethod);
    }
    if (message.signature !== "") {
      writer.uint32(42).string(message.signature);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Proof {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseProof } as Proof;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.string();
          break;
        case 2:
          message.created = reader.string();
          break;
        case 3:
          message.proofPurpose = reader.string();
          break;
        case 4:
          message.verificationMethod = reader.string();
          break;
        case 5:
          message.signature = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Proof {
    const message = { ...baseProof } as Proof;
    if (object.type !== undefined && object.type !== null) {
      message.type = String(object.type);
    } else {
      message.type = "";
    }
    if (object.created !== undefined && object.created !== null) {
      message.created = String(object.created);
    } else {
      message.created = "";
    }
    if (object.proofPurpose !== undefined && object.proofPurpose !== null) {
      message.proofPurpose = String(object.proofPurpose);
    } else {
      message.proofPurpose = "";
    }
    if (
      object.verificationMethod !== undefined &&
      object.verificationMethod !== null
    ) {
      message.verificationMethod = String(object.verificationMethod);
    } else {
      message.verificationMethod = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    return message;
  },

  toJSON(message: Proof): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = message.type);
    message.created !== undefined && (obj.created = message.created);
    message.proofPurpose !== undefined &&
      (obj.proofPurpose = message.proofPurpose);
    message.verificationMethod !== undefined &&
      (obj.verificationMethod = message.verificationMethod);
    message.signature !== undefined && (obj.signature = message.signature);
    return obj;
  },

  fromPartial(object: DeepPartial<Proof>): Proof {
    const message = { ...baseProof } as Proof;
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = "";
    }
    if (object.created !== undefined && object.created !== null) {
      message.created = object.created;
    } else {
      message.created = "";
    }
    if (object.proofPurpose !== undefined && object.proofPurpose !== null) {
      message.proofPurpose = object.proofPurpose;
    } else {
      message.proofPurpose = "";
    }
    if (
      object.verificationMethod !== undefined &&
      object.verificationMethod !== null
    ) {
      message.verificationMethod = object.verificationMethod;
    } else {
      message.verificationMethod = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
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
