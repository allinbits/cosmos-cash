/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { VerifiableCredential } from "../verifiable-credential/verifiable-credential";
export const protobufPackage = "allinbits.cosmoscash.verifiablecredential";
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
const baseMsgDeleteVerifiableCredential = {
    verifiableCredentialId: "",
    issuerDid: "",
    owner: "",
};
export const MsgDeleteVerifiableCredential = {
    encode(message, writer = Writer.create()) {
        if (message.verifiableCredentialId !== "") {
            writer.uint32(10).string(message.verifiableCredentialId);
        }
        if (message.issuerDid !== "") {
            writer.uint32(18).string(message.issuerDid);
        }
        if (message.owner !== "") {
            writer.uint32(26).string(message.owner);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgDeleteVerifiableCredential,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.verifiableCredentialId = reader.string();
                    break;
                case 2:
                    message.issuerDid = reader.string();
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
    fromJSON(object) {
        const message = {
            ...baseMsgDeleteVerifiableCredential,
        };
        if (object.verifiableCredentialId !== undefined &&
            object.verifiableCredentialId !== null) {
            message.verifiableCredentialId = String(object.verifiableCredentialId);
        }
        else {
            message.verifiableCredentialId = "";
        }
        if (object.issuerDid !== undefined && object.issuerDid !== null) {
            message.issuerDid = String(object.issuerDid);
        }
        else {
            message.issuerDid = "";
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
        message.verifiableCredentialId !== undefined &&
            (obj.verifiableCredentialId = message.verifiableCredentialId);
        message.issuerDid !== undefined && (obj.issuerDid = message.issuerDid);
        message.owner !== undefined && (obj.owner = message.owner);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgDeleteVerifiableCredential,
        };
        if (object.verifiableCredentialId !== undefined &&
            object.verifiableCredentialId !== null) {
            message.verifiableCredentialId = object.verifiableCredentialId;
        }
        else {
            message.verifiableCredentialId = "";
        }
        if (object.issuerDid !== undefined && object.issuerDid !== null) {
            message.issuerDid = object.issuerDid;
        }
        else {
            message.issuerDid = "";
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
const baseMsgDeleteVerifiableCredentialResponse = {};
export const MsgDeleteVerifiableCredentialResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgDeleteVerifiableCredentialResponse,
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
            ...baseMsgDeleteVerifiableCredentialResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgDeleteVerifiableCredentialResponse,
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
        const promise = this.rpc.request("allinbits.cosmoscash.verifiablecredential.Msg", "CreateVerifiableCredential", data);
        return promise.then((data) => MsgCreateVerifiableCredentialResponse.decode(new Reader(data)));
    }
    DeleteVerifiableCredential(request) {
        const data = MsgDeleteVerifiableCredential.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.verifiablecredential.Msg", "DeleteVerifiableCredential", data);
        return promise.then((data) => MsgDeleteVerifiableCredentialResponse.decode(new Reader(data)));
    }
}
