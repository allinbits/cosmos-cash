import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "allinbits.cosmoscash.verifiablecredentialservice";
export interface VerifiableCredential {
    context: string[];
    id: string;
    type: string[];
    issuer: string;
    issuanceDate: string;
    userCred: UserCredentialSubject | undefined;
    issuerCred: IssuerCredentialSubject | undefined;
    proof: Proof | undefined;
}
export interface UserCredentialSubject {
    id: string;
    hasKyc: boolean;
}
export interface IssuerCredentialSubject {
    id: string;
    isVerified: boolean;
}
export interface Proof {
    type: string;
    created: string;
    proofPurpose: string;
    verificationMethod: string;
    signature: string;
}
export declare const VerifiableCredential: {
    encode(message: VerifiableCredential, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): VerifiableCredential;
    fromJSON(object: any): VerifiableCredential;
    toJSON(message: VerifiableCredential): unknown;
    fromPartial(object: DeepPartial<VerifiableCredential>): VerifiableCredential;
};
export declare const UserCredentialSubject: {
    encode(message: UserCredentialSubject, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): UserCredentialSubject;
    fromJSON(object: any): UserCredentialSubject;
    toJSON(message: UserCredentialSubject): unknown;
    fromPartial(object: DeepPartial<UserCredentialSubject>): UserCredentialSubject;
};
export declare const IssuerCredentialSubject: {
    encode(message: IssuerCredentialSubject, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): IssuerCredentialSubject;
    fromJSON(object: any): IssuerCredentialSubject;
    toJSON(message: IssuerCredentialSubject): unknown;
    fromPartial(object: DeepPartial<IssuerCredentialSubject>): IssuerCredentialSubject;
};
export declare const Proof: {
    encode(message: Proof, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Proof;
    fromJSON(object: any): Proof;
    toJSON(message: Proof): unknown;
    fromPartial(object: DeepPartial<Proof>): Proof;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
