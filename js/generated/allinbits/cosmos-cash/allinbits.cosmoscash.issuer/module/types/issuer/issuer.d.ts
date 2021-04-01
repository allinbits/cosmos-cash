import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "allinbits.cosmoscash.issuer";
/** Issuer represents an e-money token issuer */
export interface Issuer {
    token: string;
    fee: number;
    address: string;
}
export declare const Issuer: {
    encode(message: Issuer, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Issuer;
    fromJSON(object: any): Issuer;
    toJSON(message: Issuer): unknown;
    fromPartial(object: DeepPartial<Issuer>): Issuer;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
