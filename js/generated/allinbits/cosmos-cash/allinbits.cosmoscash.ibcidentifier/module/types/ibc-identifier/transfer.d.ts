import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "allinbits.cosmoscash.ibcidentifier";
/** IdentifierPacketData defines a struct for the packet payload */
export interface IdentifierPacketData {
    id: number;
}
export declare const IdentifierPacketData: {
    encode(message: IdentifierPacketData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): IdentifierPacketData;
    fromJSON(object: any): IdentifierPacketData;
    toJSON(message: IdentifierPacketData): unknown;
    fromPartial(object: DeepPartial<IdentifierPacketData>): IdentifierPacketData;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
