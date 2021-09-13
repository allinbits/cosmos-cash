/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "allinbits.cosmoscash.issuer";
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.regulatorsParams !== undefined) {
            RegulatorsParams.encode(message.regulatorsParams, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.regulatorsParams = RegulatorsParams.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        if (object.regulatorsParams !== undefined &&
            object.regulatorsParams !== null) {
            message.regulatorsParams = RegulatorsParams.fromJSON(object.regulatorsParams);
        }
        else {
            message.regulatorsParams = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.regulatorsParams !== undefined &&
            (obj.regulatorsParams = message.regulatorsParams
                ? RegulatorsParams.toJSON(message.regulatorsParams)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        if (object.regulatorsParams !== undefined &&
            object.regulatorsParams !== null) {
            message.regulatorsParams = RegulatorsParams.fromPartial(object.regulatorsParams);
        }
        else {
            message.regulatorsParams = undefined;
        }
        return message;
    },
};
const baseRegulatorsParams = { addresses: "" };
export const RegulatorsParams = {
    encode(message, writer = Writer.create()) {
        for (const v of message.addresses) {
            writer.uint32(18).string(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseRegulatorsParams };
        message.addresses = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 2:
                    message.addresses.push(reader.string());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseRegulatorsParams };
        message.addresses = [];
        if (object.addresses !== undefined && object.addresses !== null) {
            for (const e of object.addresses) {
                message.addresses.push(String(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.addresses) {
            obj.addresses = message.addresses.map((e) => e);
        }
        else {
            obj.addresses = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseRegulatorsParams };
        message.addresses = [];
        if (object.addresses !== undefined && object.addresses !== null) {
            for (const e of object.addresses) {
                message.addresses.push(e);
            }
        }
        return message;
    },
};
