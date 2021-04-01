/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { PageRequest, PageResponse, } from "../cosmos/base/query/v1beta1/pagination";
import { VerifiableCredential } from "../verifiable-credential-service/verifiable-credential";
export const protobufPackage = "allinbits.cosmoscash.verifiablecredentialservice";
const baseQueryVerifiableCredentialsRequest = { status: "" };
export const QueryVerifiableCredentialsRequest = {
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
            ...baseQueryVerifiableCredentialsRequest,
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
            ...baseQueryVerifiableCredentialsRequest,
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
            ...baseQueryVerifiableCredentialsRequest,
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
const baseQueryVerifiableCredentialsResponse = {};
export const QueryVerifiableCredentialsResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.vcs) {
            VerifiableCredential.encode(v, writer.uint32(10).fork()).ldelim();
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
            ...baseQueryVerifiableCredentialsResponse,
        };
        message.vcs = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vcs.push(VerifiableCredential.decode(reader, reader.uint32()));
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
            ...baseQueryVerifiableCredentialsResponse,
        };
        message.vcs = [];
        if (object.vcs !== undefined && object.vcs !== null) {
            for (const e of object.vcs) {
                message.vcs.push(VerifiableCredential.fromJSON(e));
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
        if (message.vcs) {
            obj.vcs = message.vcs.map((e) => e ? VerifiableCredential.toJSON(e) : undefined);
        }
        else {
            obj.vcs = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryVerifiableCredentialsResponse,
        };
        message.vcs = [];
        if (object.vcs !== undefined && object.vcs !== null) {
            for (const e of object.vcs) {
                message.vcs.push(VerifiableCredential.fromPartial(e));
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
const baseQueryVerifiableCredentialRequest = {
    verifiableCredentialId: "",
};
export const QueryVerifiableCredentialRequest = {
    encode(message, writer = Writer.create()) {
        if (message.verifiableCredentialId !== "") {
            writer.uint32(10).string(message.verifiableCredentialId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryVerifiableCredentialRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.verifiableCredentialId = reader.string();
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
            ...baseQueryVerifiableCredentialRequest,
        };
        if (object.verifiableCredentialId !== undefined &&
            object.verifiableCredentialId !== null) {
            message.verifiableCredentialId = String(object.verifiableCredentialId);
        }
        else {
            message.verifiableCredentialId = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.verifiableCredentialId !== undefined &&
            (obj.verifiableCredentialId = message.verifiableCredentialId);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryVerifiableCredentialRequest,
        };
        if (object.verifiableCredentialId !== undefined &&
            object.verifiableCredentialId !== null) {
            message.verifiableCredentialId = object.verifiableCredentialId;
        }
        else {
            message.verifiableCredentialId = "";
        }
        return message;
    },
};
const baseQueryVerifiableCredentialResponse = {};
export const QueryVerifiableCredentialResponse = {
    encode(message, writer = Writer.create()) {
        if (message.verifiableCredential !== undefined) {
            VerifiableCredential.encode(message.verifiableCredential, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryVerifiableCredentialResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.verifiableCredential = VerifiableCredential.decode(reader, reader.uint32());
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
            ...baseQueryVerifiableCredentialResponse,
        };
        if (object.verifiableCredential !== undefined &&
            object.verifiableCredential !== null) {
            message.verifiableCredential = VerifiableCredential.fromJSON(object.verifiableCredential);
        }
        else {
            message.verifiableCredential = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.verifiableCredential !== undefined &&
            (obj.verifiableCredential = message.verifiableCredential
                ? VerifiableCredential.toJSON(message.verifiableCredential)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryVerifiableCredentialResponse,
        };
        if (object.verifiableCredential !== undefined &&
            object.verifiableCredential !== null) {
            message.verifiableCredential = VerifiableCredential.fromPartial(object.verifiableCredential);
        }
        else {
            message.verifiableCredential = undefined;
        }
        return message;
    },
};
const baseQueryValidateVerifiableCredentialRequest = {
    verifiableCredentialId: "",
    issuerPubkey: "",
};
export const QueryValidateVerifiableCredentialRequest = {
    encode(message, writer = Writer.create()) {
        if (message.verifiableCredentialId !== "") {
            writer.uint32(10).string(message.verifiableCredentialId);
        }
        if (message.issuerPubkey !== "") {
            writer.uint32(18).string(message.issuerPubkey);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryValidateVerifiableCredentialRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.verifiableCredentialId = reader.string();
                    break;
                case 2:
                    message.issuerPubkey = reader.string();
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
            ...baseQueryValidateVerifiableCredentialRequest,
        };
        if (object.verifiableCredentialId !== undefined &&
            object.verifiableCredentialId !== null) {
            message.verifiableCredentialId = String(object.verifiableCredentialId);
        }
        else {
            message.verifiableCredentialId = "";
        }
        if (object.issuerPubkey !== undefined && object.issuerPubkey !== null) {
            message.issuerPubkey = String(object.issuerPubkey);
        }
        else {
            message.issuerPubkey = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.verifiableCredentialId !== undefined &&
            (obj.verifiableCredentialId = message.verifiableCredentialId);
        message.issuerPubkey !== undefined &&
            (obj.issuerPubkey = message.issuerPubkey);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryValidateVerifiableCredentialRequest,
        };
        if (object.verifiableCredentialId !== undefined &&
            object.verifiableCredentialId !== null) {
            message.verifiableCredentialId = object.verifiableCredentialId;
        }
        else {
            message.verifiableCredentialId = "";
        }
        if (object.issuerPubkey !== undefined && object.issuerPubkey !== null) {
            message.issuerPubkey = object.issuerPubkey;
        }
        else {
            message.issuerPubkey = "";
        }
        return message;
    },
};
const baseQueryValidateVerifiableCredentialResponse = {
    isValid: false,
};
export const QueryValidateVerifiableCredentialResponse = {
    encode(message, writer = Writer.create()) {
        if (message.isValid === true) {
            writer.uint32(8).bool(message.isValid);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryValidateVerifiableCredentialResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.isValid = reader.bool();
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
            ...baseQueryValidateVerifiableCredentialResponse,
        };
        if (object.isValid !== undefined && object.isValid !== null) {
            message.isValid = Boolean(object.isValid);
        }
        else {
            message.isValid = false;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.isValid !== undefined && (obj.isValid = message.isValid);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryValidateVerifiableCredentialResponse,
        };
        if (object.isValid !== undefined && object.isValid !== null) {
            message.isValid = object.isValid;
        }
        else {
            message.isValid = false;
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    VerifiableCredentials(request) {
        const data = QueryVerifiableCredentialsRequest.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.verifiablecredentialservice.Query", "VerifiableCredentials", data);
        return promise.then((data) => QueryVerifiableCredentialsResponse.decode(new Reader(data)));
    }
    VerifiableCredential(request) {
        const data = QueryVerifiableCredentialRequest.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.verifiablecredentialservice.Query", "VerifiableCredential", data);
        return promise.then((data) => QueryVerifiableCredentialResponse.decode(new Reader(data)));
    }
    ValidateVerifiableCredential(request) {
        const data = QueryValidateVerifiableCredentialRequest.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.verifiablecredentialservice.Query", "ValidateVerifiableCredential", data);
        return promise.then((data) => QueryValidateVerifiableCredentialResponse.decode(new Reader(data)));
    }
}
