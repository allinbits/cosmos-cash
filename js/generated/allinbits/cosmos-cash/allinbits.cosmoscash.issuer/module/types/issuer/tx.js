/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
export const protobufPackage = "allinbits.cosmoscash.issuer";
const baseMsgCreateIssuer = {
    token: "",
    fee: 0,
    issuerDid: "",
    licenseCredId: "",
    owner: "",
};
export const MsgCreateIssuer = {
    encode(message, writer = Writer.create()) {
        if (message.token !== "") {
            writer.uint32(10).string(message.token);
        }
        if (message.fee !== 0) {
            writer.uint32(16).int32(message.fee);
        }
        if (message.issuerDid !== "") {
            writer.uint32(26).string(message.issuerDid);
        }
        if (message.licenseCredId !== "") {
            writer.uint32(34).string(message.licenseCredId);
        }
        if (message.owner !== "") {
            writer.uint32(42).string(message.owner);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateIssuer };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.token = reader.string();
                    break;
                case 2:
                    message.fee = reader.int32();
                    break;
                case 3:
                    message.issuerDid = reader.string();
                    break;
                case 4:
                    message.licenseCredId = reader.string();
                    break;
                case 5:
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
        const message = { ...baseMsgCreateIssuer };
        if (object.token !== undefined && object.token !== null) {
            message.token = String(object.token);
        }
        else {
            message.token = "";
        }
        if (object.fee !== undefined && object.fee !== null) {
            message.fee = Number(object.fee);
        }
        else {
            message.fee = 0;
        }
        if (object.issuerDid !== undefined && object.issuerDid !== null) {
            message.issuerDid = String(object.issuerDid);
        }
        else {
            message.issuerDid = "";
        }
        if (object.licenseCredId !== undefined && object.licenseCredId !== null) {
            message.licenseCredId = String(object.licenseCredId);
        }
        else {
            message.licenseCredId = "";
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
        message.token !== undefined && (obj.token = message.token);
        message.fee !== undefined && (obj.fee = message.fee);
        message.issuerDid !== undefined && (obj.issuerDid = message.issuerDid);
        message.licenseCredId !== undefined &&
            (obj.licenseCredId = message.licenseCredId);
        message.owner !== undefined && (obj.owner = message.owner);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateIssuer };
        if (object.token !== undefined && object.token !== null) {
            message.token = object.token;
        }
        else {
            message.token = "";
        }
        if (object.fee !== undefined && object.fee !== null) {
            message.fee = object.fee;
        }
        else {
            message.fee = 0;
        }
        if (object.issuerDid !== undefined && object.issuerDid !== null) {
            message.issuerDid = object.issuerDid;
        }
        else {
            message.issuerDid = "";
        }
        if (object.licenseCredId !== undefined && object.licenseCredId !== null) {
            message.licenseCredId = object.licenseCredId;
        }
        else {
            message.licenseCredId = "";
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
const baseMsgCreateIssuerResponse = {};
export const MsgCreateIssuerResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateIssuerResponse,
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
            ...baseMsgCreateIssuerResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgCreateIssuerResponse,
        };
        return message;
    },
};
const baseMsgBurnToken = { amount: "", owner: "" };
export const MsgBurnToken = {
    encode(message, writer = Writer.create()) {
        if (message.amount !== "") {
            writer.uint32(10).string(message.amount);
        }
        if (message.owner !== "") {
            writer.uint32(18).string(message.owner);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgBurnToken };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.amount = reader.string();
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
        const message = { ...baseMsgBurnToken };
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = String(object.amount);
        }
        else {
            message.amount = "";
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
        message.amount !== undefined && (obj.amount = message.amount);
        message.owner !== undefined && (obj.owner = message.owner);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgBurnToken };
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = object.amount;
        }
        else {
            message.amount = "";
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
const baseMsgBurnTokenResponse = {};
export const MsgBurnTokenResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgBurnTokenResponse };
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
        const message = { ...baseMsgBurnTokenResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgBurnTokenResponse };
        return message;
    },
};
const baseMsgMintToken = { amount: "", owner: "" };
export const MsgMintToken = {
    encode(message, writer = Writer.create()) {
        if (message.amount !== "") {
            writer.uint32(10).string(message.amount);
        }
        if (message.owner !== "") {
            writer.uint32(18).string(message.owner);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgMintToken };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.amount = reader.string();
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
        const message = { ...baseMsgMintToken };
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = String(object.amount);
        }
        else {
            message.amount = "";
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
        message.amount !== undefined && (obj.amount = message.amount);
        message.owner !== undefined && (obj.owner = message.owner);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgMintToken };
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = object.amount;
        }
        else {
            message.amount = "";
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
const baseMsgMintTokenResponse = {};
export const MsgMintTokenResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgMintTokenResponse };
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
        const message = { ...baseMsgMintTokenResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgMintTokenResponse };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateIssuer(request) {
        const data = MsgCreateIssuer.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.issuer.Msg", "CreateIssuer", data);
        return promise.then((data) => MsgCreateIssuerResponse.decode(new Reader(data)));
    }
    BurnToken(request) {
        const data = MsgBurnToken.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.issuer.Msg", "BurnToken", data);
        return promise.then((data) => MsgBurnTokenResponse.decode(new Reader(data)));
    }
    MintToken(request) {
        const data = MsgMintToken.encode(request).finish();
        const promise = this.rpc.request("allinbits.cosmoscash.issuer.Msg", "MintToken", data);
        return promise.then((data) => MsgMintTokenResponse.decode(new Reader(data)));
    }
}
