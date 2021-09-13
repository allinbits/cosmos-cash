/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { VerificationMethod, Service } from "../did/did";
export const protobufPackage = "allinbits.cosmoscash.did";
const baseVerification = { relationships: "", context: "" };
export const Verification = {
    encode(message, writer = Writer.create()) {
        for (const v of message.relationships) {
            writer.uint32(10).string(v);
        }
        if (message.method !== undefined) {
            VerificationMethod.encode(message.method, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.context) {
            writer.uint32(26).string(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseVerification };
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
    fromJSON(object) {
        const message = { ...baseVerification };
        message.relationships = [];
        message.context = [];
        if (object.relationships !== undefined && object.relationships !== null) {
            for (const e of object.relationships) {
                message.relationships.push(String(e));
            }
        }
        if (object.method !== undefined && object.method !== null) {
            message.method = VerificationMethod.fromJSON(object.method);
        }
        else {
            message.method = undefined;
        }
        if (object.context !== undefined && object.context !== null) {
            for (const e of object.context) {
                message.context.push(String(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.relationships) {
            obj.relationships = message.relationships.map((e) => e);
        }
        else {
            obj.relationships = [];
        }
        message.method !== undefined &&
            (obj.method = message.method
                ? VerificationMethod.toJSON(message.method)
                : undefined);
        if (message.context) {
            obj.context = message.context.map((e) => e);
        }
        else {
            obj.context = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseVerification };
        message.relationships = [];
        message.context = [];
        if (object.relationships !== undefined && object.relationships !== null) {
            for (const e of object.relationships) {
                message.relationships.push(e);
            }
        }
        if (object.method !== undefined && object.method !== null) {
            message.method = VerificationMethod.fromPartial(object.method);
        }
        else {
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
const baseMsgCreateDidDocument = { id: "", controller: "", signer: "" };
export const MsgCreateDidDocument = {
    encode(message, writer = Writer.create()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.controller !== "") {
            writer.uint32(18).string(message.controller);
        }
        for (const v of message.verifications) {
            Verification.encode(v, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.services) {
            Service.encode(v, writer.uint32(34).fork()).ldelim();
        }
        if (message.signer !== "") {
            writer.uint32(42).string(message.signer);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateDidDocument };
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
                    message.verifications.push(Verification.decode(reader, reader.uint32()));
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
    fromJSON(object) {
        const message = { ...baseMsgCreateDidDocument };
        message.verifications = [];
        message.services = [];
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.controller !== undefined && object.controller !== null) {
            message.controller = String(object.controller);
        }
        else {
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
        }
        else {
            message.signer = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.controller !== undefined && (obj.controller = message.controller);
        if (message.verifications) {
            obj.verifications = message.verifications.map((e) => e ? Verification.toJSON(e) : undefined);
        }
        else {
            obj.verifications = [];
        }
        if (message.services) {
            obj.services = message.services.map((e) => e ? Service.toJSON(e) : undefined);
        }
        else {
            obj.services = [];
        }
        message.signer !== undefined && (obj.signer = message.signer);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateDidDocument };
        message.verifications = [];
        message.services = [];
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.controller !== undefined && object.controller !== null) {
            message.controller = object.controller;
        }
        else {
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
        }
        else {
            message.signer = "";
        }
        return message;
    },
};
const baseMsgCreateDidDocumentResponse = {};
export const MsgCreateDidDocumentResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateDidDocumentResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgCreateDidDocumentResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgCreateDidDocumentResponse,
        };
        return message;
    },
};
const baseMsgUpdateDidDocument = { id: "", controller: "", signer: "" };
export const MsgUpdateDidDocument = {
    encode(message, writer = Writer.create()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        for (const v of message.controller) {
            writer.uint32(18).string(v);
        }
        if (message.signer !== "") {
            writer.uint32(42).string(message.signer);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateDidDocument };
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
    fromJSON(object) {
        const message = { ...baseMsgUpdateDidDocument };
        message.controller = [];
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.controller !== undefined && object.controller !== null) {
            for (const e of object.controller) {
                message.controller.push(String(e));
            }
        }
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        if (message.controller) {
            obj.controller = message.controller.map((e) => e);
        }
        else {
            obj.controller = [];
        }
        message.signer !== undefined && (obj.signer = message.signer);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateDidDocument };
        message.controller = [];
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.controller !== undefined && object.controller !== null) {
            for (const e of object.controller) {
                message.controller.push(e);
            }
        }
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = "";
        }
        return message;
    },
};
const baseMsgUpdateDidDocumentResponse = {};
export const MsgUpdateDidDocumentResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgUpdateDidDocumentResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgUpdateDidDocumentResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgUpdateDidDocumentResponse,
        };
        return message;
    },
};
const baseMsgAddVerification = { id: "", signer: "" };
export const MsgAddVerification = {
    encode(message, writer = Writer.create()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.verification !== undefined) {
            Verification.encode(message.verification, writer.uint32(18).fork()).ldelim();
        }
        if (message.signer !== "") {
            writer.uint32(26).string(message.signer);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddVerification };
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
    fromJSON(object) {
        const message = { ...baseMsgAddVerification };
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.verification !== undefined && object.verification !== null) {
            message.verification = Verification.fromJSON(object.verification);
        }
        else {
            message.verification = undefined;
        }
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.verification !== undefined &&
            (obj.verification = message.verification
                ? Verification.toJSON(message.verification)
                : undefined);
        message.signer !== undefined && (obj.signer = message.signer);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgAddVerification };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.verification !== undefined && object.verification !== null) {
            message.verification = Verification.fromPartial(object.verification);
        }
        else {
            message.verification = undefined;
        }
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = "";
        }
        return message;
    },
};
const baseMsgAddVerificationResponse = {};
export const MsgAddVerificationResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgAddVerificationResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgAddVerificationResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgAddVerificationResponse,
        };
        return message;
    },
};
const baseMsgSetVerificationRelationships = {
    id: "",
    methodId: "",
    relationships: "",
    signer: "",
};
export const MsgSetVerificationRelationships = {
    encode(message, writer = Writer.create()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.methodId !== "") {
            writer.uint32(18).string(message.methodId);
        }
        for (const v of message.relationships) {
            writer.uint32(26).string(v);
        }
        if (message.signer !== "") {
            writer.uint32(34).string(message.signer);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgSetVerificationRelationships,
        };
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
    fromJSON(object) {
        const message = {
            ...baseMsgSetVerificationRelationships,
        };
        message.relationships = [];
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.methodId !== undefined && object.methodId !== null) {
            message.methodId = String(object.methodId);
        }
        else {
            message.methodId = "";
        }
        if (object.relationships !== undefined && object.relationships !== null) {
            for (const e of object.relationships) {
                message.relationships.push(String(e));
            }
        }
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.methodId !== undefined && (obj.methodId = message.methodId);
        if (message.relationships) {
            obj.relationships = message.relationships.map((e) => e);
        }
        else {
            obj.relationships = [];
        }
        message.signer !== undefined && (obj.signer = message.signer);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgSetVerificationRelationships,
        };
        message.relationships = [];
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.methodId !== undefined && object.methodId !== null) {
            message.methodId = object.methodId;
        }
        else {
            message.methodId = "";
        }
        if (object.relationships !== undefined && object.relationships !== null) {
            for (const e of object.relationships) {
                message.relationships.push(e);
            }
        }
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = "";
        }
        return message;
    },
};
const baseMsgSetVerificationRelationshipsResponse = {};
export const MsgSetVerificationRelationshipsResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgSetVerificationRelationshipsResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgSetVerificationRelationshipsResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgSetVerificationRelationshipsResponse,
        };
        return message;
    },
};
const baseMsgRevokeVerification = { id: "", methodId: "", signer: "" };
export const MsgRevokeVerification = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRevokeVerification };
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
    fromJSON(object) {
        const message = { ...baseMsgRevokeVerification };
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.methodId !== undefined && object.methodId !== null) {
            message.methodId = String(object.methodId);
        }
        else {
            message.methodId = "";
        }
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.methodId !== undefined && (obj.methodId = message.methodId);
        message.signer !== undefined && (obj.signer = message.signer);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgRevokeVerification };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.methodId !== undefined && object.methodId !== null) {
            message.methodId = object.methodId;
        }
        else {
            message.methodId = "";
        }
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = "";
        }
        return message;
    },
};
const baseMsgRevokeVerificationResponse = {};
export const MsgRevokeVerificationResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgRevokeVerificationResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgRevokeVerificationResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgRevokeVerificationResponse,
        };
        return message;
    },
};
const baseMsgAddService = { id: "", signer: "" };
export const MsgAddService = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddService };
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
    fromJSON(object) {
        const message = { ...baseMsgAddService };
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.serviceData !== undefined && object.serviceData !== null) {
            message.serviceData = Service.fromJSON(object.serviceData);
        }
        else {
            message.serviceData = undefined;
        }
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.serviceData !== undefined &&
            (obj.serviceData = message.serviceData
                ? Service.toJSON(message.serviceData)
                : undefined);
        message.signer !== undefined && (obj.signer = message.signer);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgAddService };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.serviceData !== undefined && object.serviceData !== null) {
            message.serviceData = Service.fromPartial(object.serviceData);
        }
        else {
            message.serviceData = undefined;
        }
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = "";
        }
        return message;
    },
};
const baseMsgAddServiceResponse = {};
export const MsgAddServiceResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddServiceResponse };
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
    fromJSON(_) {
        const message = { ...baseMsgAddServiceResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgAddServiceResponse };
        return message;
    },
};
const baseMsgDeleteService = { id: "", serviceId: "", signer: "" };
export const MsgDeleteService = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeleteService };
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
    fromJSON(object) {
        const message = { ...baseMsgDeleteService };
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.serviceId !== undefined && object.serviceId !== null) {
            message.serviceId = String(object.serviceId);
        }
        else {
            message.serviceId = "";
        }
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.serviceId !== undefined && (obj.serviceId = message.serviceId);
        message.signer !== undefined && (obj.signer = message.signer);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgDeleteService };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.serviceId !== undefined && object.serviceId !== null) {
            message.serviceId = object.serviceId;
        }
        else {
            message.serviceId = "";
        }
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = "";
        }
        return message;
    },
};
const baseMsgDeleteServiceResponse = {};
export const MsgDeleteServiceResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgDeleteServiceResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgDeleteServiceResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgDeleteServiceResponse,
        };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateDidDocument(request) {
        const data = MsgCreateDidDocument.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.did.Msg", "CreateDidDocument", data);
        return promise.then((data) => MsgCreateDidDocumentResponse.decode(new Reader(data)));
    }
    UpdateDidDocument(request) {
        const data = MsgUpdateDidDocument.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.did.Msg", "UpdateDidDocument", data);
        return promise.then((data) => MsgUpdateDidDocumentResponse.decode(new Reader(data)));
    }
    AddVerification(request) {
        const data = MsgAddVerification.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.did.Msg", "AddVerification", data);
        return promise.then((data) => MsgAddVerificationResponse.decode(new Reader(data)));
    }
    RevokeVerification(request) {
        const data = MsgRevokeVerification.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.did.Msg", "RevokeVerification", data);
        return promise.then((data) => MsgRevokeVerificationResponse.decode(new Reader(data)));
    }
    SetVerificationRelationships(request) {
        const data = MsgSetVerificationRelationships.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.did.Msg", "SetVerificationRelationships", data);
        return promise.then((data) => MsgSetVerificationRelationshipsResponse.decode(new Reader(data)));
    }
    AddService(request) {
        const data = MsgAddService.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.did.Msg", "AddService", data);
        return promise.then((data) => MsgAddServiceResponse.decode(new Reader(data)));
    }
    DeleteService(request) {
        const data = MsgDeleteService.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.did.Msg", "DeleteService", data);
        return promise.then((data) => MsgDeleteServiceResponse.decode(new Reader(data)));
    }
}
