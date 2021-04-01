/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { VerifiableCredential } from "../verifiable-credential-service/verifiable-credential";
export const protobufPackage = "allinbits.cosmoscash.verifiablecredentialservice";
const baseMsgCreateVerifiableCredential = { owner: "" };
export const MsgCreateVerifiableCredential = {
    encode(message, writer = Writer.create()) {
        if (message.verifiableCredential !== undefined) {
            VerifiableCredential.encode(message.verifiableCredential, writer.uint32(10).fork()).ldelim();
        }
        if (message.owner !== "") {
            writer.uint32(18).string(message.owner);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateVerifiableCredential,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.verifiableCredential = VerifiableCredential.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.owner = reader.string();
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
            ...baseMsgCreateVerifiableCredential,
        };
        if (object.verifiableCredential !== undefined &&
            object.verifiableCredential !== null) {
            message.verifiableCredential = VerifiableCredential.fromJSON(object.verifiableCredential);
        }
        else {
            message.verifiableCredential = undefined;
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = String(object.owner);
        }
        else {
            message.owner = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.verifiableCredential !== undefined &&
            (obj.verifiableCredential = message.verifiableCredential
                ? VerifiableCredential.toJSON(message.verifiableCredential)
                : undefined);
        message.owner !== undefined && (obj.owner = message.owner);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgCreateVerifiableCredential,
        };
        if (object.verifiableCredential !== undefined &&
            object.verifiableCredential !== null) {
            message.verifiableCredential = VerifiableCredential.fromPartial(object.verifiableCredential);
        }
        else {
            message.verifiableCredential = undefined;
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = "";
        }
        return message;
    },
};
const baseMsgCreateVerifiableCredentialResponse = {};
export const MsgCreateVerifiableCredentialResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateVerifiableCredentialResponse,
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
            ...baseMsgCreateVerifiableCredentialResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgCreateVerifiableCredentialResponse,
        };
        return message;
    },
};
const baseMsgCreateIssuerVerifiableCredential = { owner: "" };
export const MsgCreateIssuerVerifiableCredential = {
    encode(message, writer = Writer.create()) {
        if (message.verifiableCredential !== undefined) {
            VerifiableCredential.encode(message.verifiableCredential, writer.uint32(10).fork()).ldelim();
        }
        if (message.owner !== "") {
            writer.uint32(18).string(message.owner);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateIssuerVerifiableCredential,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.verifiableCredential = VerifiableCredential.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.owner = reader.string();
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
            ...baseMsgCreateIssuerVerifiableCredential,
        };
        if (object.verifiableCredential !== undefined &&
            object.verifiableCredential !== null) {
            message.verifiableCredential = VerifiableCredential.fromJSON(object.verifiableCredential);
        }
        else {
            message.verifiableCredential = undefined;
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = String(object.owner);
        }
        else {
            message.owner = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.verifiableCredential !== undefined &&
            (obj.verifiableCredential = message.verifiableCredential
                ? VerifiableCredential.toJSON(message.verifiableCredential)
                : undefined);
        message.owner !== undefined && (obj.owner = message.owner);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgCreateIssuerVerifiableCredential,
        };
        if (object.verifiableCredential !== undefined &&
            object.verifiableCredential !== null) {
            message.verifiableCredential = VerifiableCredential.fromPartial(object.verifiableCredential);
        }
        else {
            message.verifiableCredential = undefined;
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = "";
        }
        return message;
    },
};
const baseMsgCreateIssuerVerifiableCredentialResponse = {};
export const MsgCreateIssuerVerifiableCredentialResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateIssuerVerifiableCredentialResponse,
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
            ...baseMsgCreateIssuerVerifiableCredentialResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgCreateIssuerVerifiableCredentialResponse,
        };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateVerifiableCredential(request) {
        const data = MsgCreateVerifiableCredential.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.verifiablecredentialservice.Msg", "CreateVerifiableCredential", data);
        return promise.then((data) => MsgCreateVerifiableCredentialResponse.decode(new Reader(data)));
    }
    CreateIssuerVerifiableCredential(request) {
        const data = MsgCreateIssuerVerifiableCredential.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.verifiablecredentialservice.Msg", "CreateIssuerVerifiableCredential", data);
        return promise.then((data) => MsgCreateIssuerVerifiableCredentialResponse.decode(new Reader(data)));
    }
}
