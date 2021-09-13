/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import { Coin } from "../cosmos/base/v1beta1/coin";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "allinbits.cosmoscash.verifiablecredential";

/** VerifiableCredential represents a verifiable credential */
export interface VerifiableCredential {
  /** @context is spec for verifiable credential. */
  context: string[];
  /**
   * The value of the id property MUST be a single URI. It is RECOMMENDED
   * that the URI in the id be one which, if dereferenced, results in a
   * document containing machine-readable information about the id.
   */
  id: string;
  /**
   * The value of the type property MUST be, or map to (through interpretation
   * of the @context property), one or more URIs. If more than one URI is
   * provided, the URIs MUST be interpreted as an unordered set.
   */
  type: string[];
  /**
   * The value of the issuer property MUST be either a URI or an object
   * containing an id property. It is RECOMMENDED that the URI in the issuer
   * or its id be one which, if dereferenced, results in a document containing
   * machine-readable information about the issuer that can be used to verify
   * the information expressed in the credential.
   */
  issuer: string;
  /**
   * A credential MUST have an issuanceDate property. The value of the issuanceDate
   * property MUST be a string value of an [RFC3339] combined date and time string
   * representing the date and time the credential becomes valid, which could
   * be a date and time in the future. Note that this value represents the earliest
   * point in time at which the information associated with the credentialSubject
   * property becomes valid.
   */
  issuanceDate: Date | undefined;
  /**
   * The value of user_cred represents a privacy respecting verifiable
   * credential. This is used when adding sensitive information about
   * a credential subject. It allows the credential subject to create
   * and validate proofs about what is contained in a credential without
   * revealing the values contained in the credential otherwise known as
   * selective disclosure.
   */
  userCred: UserCredentialSubject | undefined;
  /**
   * The value of license_cred represents a license issued by a regulatory
   * body. The license can be used to define authorized actions by the
   * credential subject
   */
  licenseCred: LicenseCredentialSubject | undefined;
  /**
   * One or more cryptographic proofs that can be used to detect tampering
   * and verify the authorship of a credential or presentation. The specific
   * method used for an embedded proof MUST be included using the type property.
   */
  proof: Proof | undefined;
}

/**
 * UserCredentialSubject represents a privacy respecting
 * credential_subject of a verifiable credential. This
 * is used as an on chain verifiable credential.
 */
export interface UserCredentialSubject {
  id: string;
  root: string;
  isVerified: boolean;
}

/**
 * The LicenseCredential message makes reference to the classes of crypto assets
 * described in MiCA, but the license could easily be adopted as proof of
 * authority to issue various types of crypto or virtual assets. The LicenseCredential
 * is used a a credential_subject in a verifiable credential.
 */
export interface LicenseCredentialSubject {
  /** The value of id represents the ID of the credential_subject */
  id: string;
  /**
   * The license type is defined by the MICA regulation. This will
   * be used to identify different asset classes of tokens being issuedi
   * by the credential_subject.
   */
  licenseType: string;
  /** The country field represents the country the credential was issued in. */
  country: string;
  /** The authority field represents a licensing authority that issued the LicenseCredential */
  authority: string;
  /**
   * The circulation_limit represents the amount of a token
   * that can be minted by a credential_subject.
   */
  circulationLimit: Coin | undefined;
}

/**
 * The Proof message represents a cryptographic proof that the
 * credential has not been tampered with or changed without the issuersi
 * knowledge. This can be used to verify the verifiable credential.
 */
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
    if (message.issuanceDate !== undefined) {
      Timestamp.encode(
        toTimestamp(message.issuanceDate),
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.userCred !== undefined) {
      UserCredentialSubject.encode(
        message.userCred,
        writer.uint32(50).fork()
      ).ldelim();
    }
    if (message.licenseCred !== undefined) {
      LicenseCredentialSubject.encode(
        message.licenseCred,
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
          message.issuanceDate = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 6:
          message.userCred = UserCredentialSubject.decode(
            reader,
            reader.uint32()
          );
          break;
        case 7:
          message.licenseCred = LicenseCredentialSubject.decode(
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
      message.issuanceDate = fromJsonTimestamp(object.issuanceDate);
    } else {
      message.issuanceDate = undefined;
    }
    if (object.userCred !== undefined && object.userCred !== null) {
      message.userCred = UserCredentialSubject.fromJSON(object.userCred);
    } else {
      message.userCred = undefined;
    }
    if (object.licenseCred !== undefined && object.licenseCred !== null) {
      message.licenseCred = LicenseCredentialSubject.fromJSON(
        object.licenseCred
      );
    } else {
      message.licenseCred = undefined;
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
      (obj.issuanceDate =
        message.issuanceDate !== undefined
          ? message.issuanceDate.toISOString()
          : null);
    message.userCred !== undefined &&
      (obj.userCred = message.userCred
        ? UserCredentialSubject.toJSON(message.userCred)
        : undefined);
    message.licenseCred !== undefined &&
      (obj.licenseCred = message.licenseCred
        ? LicenseCredentialSubject.toJSON(message.licenseCred)
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
      message.issuanceDate = undefined;
    }
    if (object.userCred !== undefined && object.userCred !== null) {
      message.userCred = UserCredentialSubject.fromPartial(object.userCred);
    } else {
      message.userCred = undefined;
    }
    if (object.licenseCred !== undefined && object.licenseCred !== null) {
      message.licenseCred = LicenseCredentialSubject.fromPartial(
        object.licenseCred
      );
    } else {
      message.licenseCred = undefined;
    }
    if (object.proof !== undefined && object.proof !== null) {
      message.proof = Proof.fromPartial(object.proof);
    } else {
      message.proof = undefined;
    }
    return message;
  },
};

const baseUserCredentialSubject: object = {
  id: "",
  root: "",
  isVerified: false,
};

export const UserCredentialSubject = {
  encode(
    message: UserCredentialSubject,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.root !== "") {
      writer.uint32(18).string(message.root);
    }
    if (message.isVerified === true) {
      writer.uint32(24).bool(message.isVerified);
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
          message.root = reader.string();
          break;
        case 3:
          message.isVerified = reader.bool();
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
    if (object.root !== undefined && object.root !== null) {
      message.root = String(object.root);
    } else {
      message.root = "";
    }
    if (object.isVerified !== undefined && object.isVerified !== null) {
      message.isVerified = Boolean(object.isVerified);
    } else {
      message.isVerified = false;
    }
    return message;
  },

  toJSON(message: UserCredentialSubject): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.root !== undefined && (obj.root = message.root);
    message.isVerified !== undefined && (obj.isVerified = message.isVerified);
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
    if (object.root !== undefined && object.root !== null) {
      message.root = object.root;
    } else {
      message.root = "";
    }
    if (object.isVerified !== undefined && object.isVerified !== null) {
      message.isVerified = object.isVerified;
    } else {
      message.isVerified = false;
    }
    return message;
  },
};

const baseLicenseCredentialSubject: object = {
  id: "",
  licenseType: "",
  country: "",
  authority: "",
};

export const LicenseCredentialSubject = {
  encode(
    message: LicenseCredentialSubject,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.licenseType !== "") {
      writer.uint32(18).string(message.licenseType);
    }
    if (message.country !== "") {
      writer.uint32(26).string(message.country);
    }
    if (message.authority !== "") {
      writer.uint32(34).string(message.authority);
    }
    if (message.circulationLimit !== undefined) {
      Coin.encode(message.circulationLimit, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): LicenseCredentialSubject {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseLicenseCredentialSubject,
    } as LicenseCredentialSubject;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.licenseType = reader.string();
          break;
        case 3:
          message.country = reader.string();
          break;
        case 4:
          message.authority = reader.string();
          break;
        case 5:
          message.circulationLimit = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LicenseCredentialSubject {
    const message = {
      ...baseLicenseCredentialSubject,
    } as LicenseCredentialSubject;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.licenseType !== undefined && object.licenseType !== null) {
      message.licenseType = String(object.licenseType);
    } else {
      message.licenseType = "";
    }
    if (object.country !== undefined && object.country !== null) {
      message.country = String(object.country);
    } else {
      message.country = "";
    }
    if (object.authority !== undefined && object.authority !== null) {
      message.authority = String(object.authority);
    } else {
      message.authority = "";
    }
    if (
      object.circulationLimit !== undefined &&
      object.circulationLimit !== null
    ) {
      message.circulationLimit = Coin.fromJSON(object.circulationLimit);
    } else {
      message.circulationLimit = undefined;
    }
    return message;
  },

  toJSON(message: LicenseCredentialSubject): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.licenseType !== undefined &&
      (obj.licenseType = message.licenseType);
    message.country !== undefined && (obj.country = message.country);
    message.authority !== undefined && (obj.authority = message.authority);
    message.circulationLimit !== undefined &&
      (obj.circulationLimit = message.circulationLimit
        ? Coin.toJSON(message.circulationLimit)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<LicenseCredentialSubject>
  ): LicenseCredentialSubject {
    const message = {
      ...baseLicenseCredentialSubject,
    } as LicenseCredentialSubject;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.licenseType !== undefined && object.licenseType !== null) {
      message.licenseType = object.licenseType;
    } else {
      message.licenseType = "";
    }
    if (object.country !== undefined && object.country !== null) {
      message.country = object.country;
    } else {
      message.country = "";
    }
    if (object.authority !== undefined && object.authority !== null) {
      message.authority = object.authority;
    } else {
      message.authority = "";
    }
    if (
      object.circulationLimit !== undefined &&
      object.circulationLimit !== null
    ) {
      message.circulationLimit = Coin.fromPartial(object.circulationLimit);
    } else {
      message.circulationLimit = undefined;
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
