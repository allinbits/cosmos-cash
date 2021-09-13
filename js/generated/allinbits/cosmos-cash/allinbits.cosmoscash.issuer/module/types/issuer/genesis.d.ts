import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "allinbits.cosmoscash.issuer";
/** GenesisState defines the issuer module's genesis state. */
export interface GenesisState {
    /** this line is used by starport scaffolding # genesis/proto/state */
    regulatorsParams: RegulatorsParams | undefined;
}
/** RegulatorsParams defines the addresses of the regulators */
export interface RegulatorsParams {
    /**
     * the addresses of the regualtors for the chain. The addresses will be used to
     * generate DID documents at genesis.
     */
    addresses: string[];
}
export declare const GenesisState: {
    encode(message: GenesisState, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisState;
    fromJSON(object: any): GenesisState;
    toJSON(message: GenesisState): unknown;
    fromPartial(object: DeepPartial<GenesisState>): GenesisState;
};
export declare const RegulatorsParams: {
    encode(message: RegulatorsParams, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): RegulatorsParams;
    fromJSON(object: any): RegulatorsParams;
    toJSON(message: RegulatorsParams): unknown;
    fromPartial(object: DeepPartial<RegulatorsParams>): RegulatorsParams;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
