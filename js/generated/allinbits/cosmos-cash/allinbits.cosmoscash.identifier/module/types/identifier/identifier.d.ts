import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "allinbits.cosmoscash.identifier";
/** DidDocument represents a dencentralised identifer. */
export interface DidDocument {
    /** @context is spec for did document. */
    context: string;
    /** id represents the id for the did document. */
    id: string;
    /** authentication represents public key associated with the did document. */
    authentication: Authentication[];
    /** services represents each service associated with a did */
    services: Service[];
}
/** Authentication defines how to authenticate a did document. */
export interface Authentication {
    id: string;
    type: string;
    controller: string;
    publicKey: string;
}
/** Service defines how to find data associated with a identifer */
export interface Service {
    id: string;
    type: string;
    serviceEndpoint: string;
}
export declare const DidDocument: {
    encode(message: DidDocument, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): DidDocument;
    fromJSON(object: any): DidDocument;
    toJSON(message: DidDocument): unknown;
    fromPartial(object: DeepPartial<DidDocument>): DidDocument;
};
export declare const Authentication: {
    encode(message: Authentication, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Authentication;
    fromJSON(object: any): Authentication;
    toJSON(message: Authentication): unknown;
    fromPartial(object: DeepPartial<Authentication>): Authentication;
};
export declare const Service: {
    encode(message: Service, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Service;
    fromJSON(object: any): Service;
    toJSON(message: Service): unknown;
    fromPartial(object: DeepPartial<Service>): Service;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
