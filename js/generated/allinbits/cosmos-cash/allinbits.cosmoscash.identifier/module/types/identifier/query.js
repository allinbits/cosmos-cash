/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { PageRequest, PageResponse, } from "../cosmos/base/query/v1beta1/pagination";
import { DidDocument } from "../identifier/identifier";
export const protobufPackage = "allinbits.cosmoscash.identifier";
const baseQueryIdentifiersRequest = { status: "" };
export const QueryIdentifiersRequest = {
    encode(message, writer = Writer.create()) {
        if (message.status !== "") {
            writer.uint32(10).string(message.status);
        }
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryIdentifiersRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.status = reader.string();
                    break;
                case 2:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
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
            ...baseQueryIdentifiersRequest,
        };
        if (object.status !== undefined && object.status !== null) {
            message.status = String(object.status);
        }
        else {
            message.status = "";
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.status !== undefined && (obj.status = message.status);
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryIdentifiersRequest,
        };
        if (object.status !== undefined && object.status !== null) {
            message.status = object.status;
        }
        else {
            message.status = "";
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryIdentifiersResponse = {};
export const QueryIdentifiersResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.didDocuments) {
            DidDocument.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryIdentifiersResponse,
        };
        message.didDocuments = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.didDocuments.push(DidDocument.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
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
            ...baseQueryIdentifiersResponse,
        };
        message.didDocuments = [];
        if (object.didDocuments !== undefined && object.didDocuments !== null) {
            for (const e of object.didDocuments) {
                message.didDocuments.push(DidDocument.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.didDocuments) {
            obj.didDocuments = message.didDocuments.map((e) => e ? DidDocument.toJSON(e) : undefined);
        }
        else {
            obj.didDocuments = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryIdentifiersResponse,
        };
        message.didDocuments = [];
        if (object.didDocuments !== undefined && object.didDocuments !== null) {
            for (const e of object.didDocuments) {
                message.didDocuments.push(DidDocument.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryIdentifierRequest = { id: "" };
export const QueryIdentifierRequest = {
    encode(message, writer = Writer.create()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryIdentifierRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryIdentifierRequest };
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryIdentifierRequest };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        return message;
    },
};
const baseQueryIdentifierResponse = {};
export const QueryIdentifierResponse = {
    encode(message, writer = Writer.create()) {
        if (message.didDocument !== undefined) {
            DidDocument.encode(message.didDocument, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryIdentifierResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.didDocument = DidDocument.decode(reader, reader.uint32());
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
            ...baseQueryIdentifierResponse,
        };
        if (object.didDocument !== undefined && object.didDocument !== null) {
            message.didDocument = DidDocument.fromJSON(object.didDocument);
        }
        else {
            message.didDocument = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.didDocument !== undefined &&
            (obj.didDocument = message.didDocument
                ? DidDocument.toJSON(message.didDocument)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryIdentifierResponse,
        };
        if (object.didDocument !== undefined && object.didDocument !== null) {
            message.didDocument = DidDocument.fromPartial(object.didDocument);
        }
        else {
            message.didDocument = undefined;
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Identifiers(request) {
        const data = QueryIdentifiersRequest.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.identifier.Query", "Identifiers", data);
        return promise.then((data) => QueryIdentifiersResponse.decode(new Reader(data)));
    }
    Identifier(request) {
        const data = QueryIdentifierRequest.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.identifier.Query", "Identifier", data);
        return promise.then((data) => QueryIdentifierResponse.decode(new Reader(data)));
    }
}
