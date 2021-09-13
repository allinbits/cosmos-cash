/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import { Coin } from "../cosmos/base/v1beta1/coin";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "allinbits.cosmoscash.verifiablecredential";
const baseVerifiableCredential = {
    context: "",
    id: "",
    type: "",
    issuer: "",
};
export const VerifiableCredential = {
    encode(message, writer = Writer.create()) {
        for (const v of message.context) {
            writer.uint32(10).string(v);
        }
        if (message.id !== "") {
            writer.uint32(18).string(message.id);
        }
        for (const v of message.type) {
            writer.uint32(26).string(v);
        }
        if (message.issuer !== "") {
            writer.uint32(34).string(message.issuer);
        }
        if (message.issuanceDate !== undefined) {
            Timestamp.encode(toTimestamp(message.issuanceDate), writer.uint32(42).fork()).ldelim();
        }
        if (message.userCred !== undefined) {
            UserCredentialSubject.encode(message.userCred, writer.uint32(50).fork()).ldelim();
        }
        if (message.licenseCred !== undefined) {
            LicenseCredentialSubject.encode(message.licenseCred, writer.uint32(58).fork()).ldelim();
        }
        if (message.proof !== undefined) {
            Proof.encode(message.proof, writer.uint32(66).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseVerifiableCredential };
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
                    message.issuanceDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
                    break;
                case 6:
                    message.userCred = UserCredentialSubject.decode(reader, reader.uint32());
                    break;
                case 7:
                    message.licenseCred = LicenseCredentialSubject.decode(reader, reader.uint32());
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
    fromJSON(object) {
        const message = { ...baseVerifiableCredential };
        message.context = [];
        message.type = [];
        if (object.context !== undefined && object.context !== null) {
            for (const e of object.context) {
                message.context.push(String(e));
            }
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.type !== undefined && object.type !== null) {
            for (const e of object.type) {
                message.type.push(String(e));
            }
        }
        if (object.issuer !== undefined && object.issuer !== null) {
            message.issuer = String(object.issuer);
        }
        else {
            message.issuer = "";
        }
        if (object.issuanceDate !== undefined && object.issuanceDate !== null) {
            message.issuanceDate = fromJsonTimestamp(object.issuanceDate);
        }
        else {
            message.issuanceDate = undefined;
        }
        if (object.userCred !== undefined && object.userCred !== null) {
            message.userCred = UserCredentialSubject.fromJSON(object.userCred);
        }
        else {
            message.userCred = undefined;
        }
        if (object.licenseCred !== undefined && object.licenseCred !== null) {
            message.licenseCred = LicenseCredentialSubject.fromJSON(object.licenseCred);
        }
        else {
            message.licenseCred = undefined;
        }
        if (object.proof !== undefined && object.proof !== null) {
            message.proof = Proof.fromJSON(object.proof);
        }
        else {
            message.proof = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.context) {
            obj.context = message.context.map((e) => e);
        }
        else {
            obj.context = [];
        }
        message.id !== undefined && (obj.id = message.id);
        if (message.type) {
            obj.type = message.type.map((e) => e);
        }
        else {
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
    fromPartial(object) {
        const message = { ...baseVerifiableCredential };
        message.context = [];
        message.type = [];
        if (object.context !== undefined && object.context !== null) {
            for (const e of object.context) {
                message.context.push(e);
            }
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.type !== undefined && object.type !== null) {
            for (const e of object.type) {
                message.type.push(e);
            }
        }
        if (object.issuer !== undefined && object.issuer !== null) {
            message.issuer = object.issuer;
        }
        else {
            message.issuer = "";
        }
        if (object.issuanceDate !== undefined && object.issuanceDate !== null) {
            message.issuanceDate = object.issuanceDate;
        }
        else {
            message.issuanceDate = undefined;
        }
        if (object.userCred !== undefined && object.userCred !== null) {
            message.userCred = UserCredentialSubject.fromPartial(object.userCred);
        }
        else {
            message.userCred = undefined;
        }
        if (object.licenseCred !== undefined && object.licenseCred !== null) {
            message.licenseCred = LicenseCredentialSubject.fromPartial(object.licenseCred);
        }
        else {
            message.licenseCred = undefined;
        }
        if (object.proof !== undefined && object.proof !== null) {
            message.proof = Proof.fromPartial(object.proof);
        }
        else {
            message.proof = undefined;
        }
        return message;
    },
};
const baseUserCredentialSubject = {
    id: "",
    root: "",
    isVerified: false,
};
export const UserCredentialSubject = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseUserCredentialSubject };
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
    fromJSON(object) {
        const message = { ...baseUserCredentialSubject };
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.root !== undefined && object.root !== null) {
            message.root = String(object.root);
        }
        else {
            message.root = "";
        }
        if (object.isVerified !== undefined && object.isVerified !== null) {
            message.isVerified = Boolean(object.isVerified);
        }
        else {
            message.isVerified = false;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.root !== undefined && (obj.root = message.root);
        message.isVerified !== undefined && (obj.isVerified = message.isVerified);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseUserCredentialSubject };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.root !== undefined && object.root !== null) {
            message.root = object.root;
        }
        else {
            message.root = "";
        }
        if (object.isVerified !== undefined && object.isVerified !== null) {
            message.isVerified = object.isVerified;
        }
        else {
            message.isVerified = false;
        }
        return message;
    },
};
const baseLicenseCredentialSubject = {
    id: "",
    licenseType: "",
    country: "",
    authority: "",
};
export const LicenseCredentialSubject = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseLicenseCredentialSubject,
        };
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
    fromJSON(object) {
        const message = {
            ...baseLicenseCredentialSubject,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.licenseType !== undefined && object.licenseType !== null) {
            message.licenseType = String(object.licenseType);
        }
        else {
            message.licenseType = "";
        }
        if (object.country !== undefined && object.country !== null) {
            message.country = String(object.country);
        }
        else {
            message.country = "";
        }
        if (object.authority !== undefined && object.authority !== null) {
            message.authority = String(object.authority);
        }
        else {
            message.authority = "";
        }
        if (object.circulationLimit !== undefined &&
            object.circulationLimit !== null) {
            message.circulationLimit = Coin.fromJSON(object.circulationLimit);
        }
        else {
            message.circulationLimit = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
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
    fromPartial(object) {
        const message = {
            ...baseLicenseCredentialSubject,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.licenseType !== undefined && object.licenseType !== null) {
            message.licenseType = object.licenseType;
        }
        else {
            message.licenseType = "";
        }
        if (object.country !== undefined && object.country !== null) {
            message.country = object.country;
        }
        else {
            message.country = "";
        }
        if (object.authority !== undefined && object.authority !== null) {
            message.authority = object.authority;
        }
        else {
            message.authority = "";
        }
        if (object.circulationLimit !== undefined &&
            object.circulationLimit !== null) {
            message.circulationLimit = Coin.fromPartial(object.circulationLimit);
        }
        else {
            message.circulationLimit = undefined;
        }
        return message;
    },
};
const baseProof = {
    type: "",
    created: "",
    proofPurpose: "",
    verificationMethod: "",
    signature: "",
};
export const Proof = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseProof };
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
    fromJSON(object) {
        const message = { ...baseProof };
        if (object.type !== undefined && object.type !== null) {
            message.type = String(object.type);
        }
        else {
            message.type = "";
        }
        if (object.created !== undefined && object.created !== null) {
            message.created = String(object.created);
        }
        else {
            message.created = "";
        }
        if (object.proofPurpose !== undefined && object.proofPurpose !== null) {
            message.proofPurpose = String(object.proofPurpose);
        }
        else {
            message.proofPurpose = "";
        }
        if (object.verificationMethod !== undefined &&
            object.verificationMethod !== null) {
            message.verificationMethod = String(object.verificationMethod);
        }
        else {
            message.verificationMethod = "";
        }
        if (object.signature !== undefined && object.signature !== null) {
            message.signature = String(object.signature);
        }
        else {
            message.signature = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.type !== undefined && (obj.type = message.type);
        message.created !== undefined && (obj.created = message.created);
        message.proofPurpose !== undefined &&
            (obj.proofPurpose = message.proofPurpose);
        message.verificationMethod !== undefined &&
            (obj.verificationMethod = message.verificationMethod);
        message.signature !== undefined && (obj.signature = message.signature);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseProof };
        if (object.type !== undefined && object.type !== null) {
            message.type = object.type;
        }
        else {
            message.type = "";
        }
        if (object.created !== undefined && object.created !== null) {
            message.created = object.created;
        }
        else {
            message.created = "";
        }
        if (object.proofPurpose !== undefined && object.proofPurpose !== null) {
            message.proofPurpose = object.proofPurpose;
        }
        else {
            message.proofPurpose = "";
        }
        if (object.verificationMethod !== undefined &&
            object.verificationMethod !== null) {
            message.verificationMethod = object.verificationMethod;
        }
        else {
            message.verificationMethod = "";
        }
        if (object.signature !== undefined && object.signature !== null) {
            message.signature = object.signature;
        }
        else {
            message.signature = "";
        }
        return message;
    },
};
function toTimestamp(date) {
    const seconds = date.getTime() / 1000;
    const nanos = (date.getTime() % 1000) * 1000000;
    return { seconds, nanos };
}
function fromTimestamp(t) {
    let millis = t.seconds * 1000;
    millis += t.nanos / 1000000;
    return new Date(millis);
}
function fromJsonTimestamp(o) {
    if (o instanceof Date) {
        return o;
    }
    else if (typeof o === "string") {
        return new Date(o);
    }
    else {
        return fromTimestamp(Timestamp.fromJSON(o));
    }
}
