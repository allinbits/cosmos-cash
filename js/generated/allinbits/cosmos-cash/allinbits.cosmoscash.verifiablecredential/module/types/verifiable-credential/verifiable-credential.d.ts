import { Coin } from "../cosmos/base/v1beta1/coin";
import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "allinbits.cosmoscash.verifiablecredential";
/** VerifiableCredential represents a verifiable credential */
export interface VerifiableCredential {
    /** @context is spec for verifiable credential. */
    context: string[];
    /**
     * The value of the id property MUST be a single URI. It is RECOMMENDED
     * that the URI in the id be one which, if dereferenced, results in a
     * document containing machine-readable information about the id.
     */
    id: string;
    /**
     * The value of the type property MUST be, or map to (through interpretation
     * of the @context property), one or more URIs. If more than one URI is
     * provided, the URIs MUST be interpreted as an unordered set.
     */
    type: string[];
    /**
     * The value of the issuer property MUST be either a URI or an object
     * containing an id property. It is RECOMMENDED that the URI in the issuer
     * or its id be one which, if dereferenced, results in a document containing
     * machine-readable information about the issuer that can be used to verify
     * the information expressed in the credential.
     */
    issuer: string;
    /**
     * A credential MUST have an issuanceDate property. The value of the issuanceDate
     * property MUST be a string value of an [RFC3339] combined date and time string
     * representing the date and time the credential becomes valid, which could
     * be a date and time in the future. Note that this value represents the earliest
     * point in time at which the information associated with the credentialSubject
     * property becomes valid.
     */
    issuanceDate: Date | undefined;
    /**
     * The value of user_cred represents a privacy respecting verifiable
     * credential. This is used when adding sensitive information about
     * a credential subject. It allows the credential subject to create
     * and validate proofs about what is contained in a credential without
     * revealing the values contained in the credential otherwise known as
     * selective disclosure.
     */
    userCred: UserCredentialSubject | undefined;
    /**
     * The value of license_cred represents a license issued by a regulatory
     * body. The license can be used to define authorized actions by the
     * credential subject
     */
    licenseCred: LicenseCredentialSubject | undefined;
    /**
     * One or more cryptographic proofs that can be used to detect tampering
     * and verify the authorship of a credential or presentation. The specific
     * method used for an embedded proof MUST be included using the type property.
     */
    proof: Proof | undefined;
}
/**
 * UserCredentialSubject represents a privacy respecting
 * credential_subject of a verifiable credential. This
 * is used as an on chain verifiable credential.
 */
export interface UserCredentialSubject {
    id: string;
    root: string;
    isVerified: boolean;
}
/**
 * The LicenseCredential message makes reference to the classes of crypto assets
 * described in MiCA, but the license could easily be adopted as proof of
 * authority to issue various types of crypto or virtual assets. The LicenseCredential
 * is used a a credential_subject in a verifiable credential.
 */
export interface LicenseCredentialSubject {
    /** The value of id represents the ID of the credential_subject */
    id: string;
    /**
     * The license type is defined by the MICA regulation. This will
     * be used to identify different asset classes of tokens being issuedi
     * by the credential_subject.
     */
    licenseType: string;
    /** The country field represents the country the credential was issued in. */
    country: string;
    /** The authority field represents a licensing authority that issued the LicenseCredential */
    authority: string;
    /**
     * The circulation_limit represents the amount of a token
     * that can be minted by a credential_subject.
     */
    circulationLimit: Coin | undefined;
}
/**
 * The Proof message represents a cryptographic proof that the
 * credential has not been tampered with or changed without the issuersi
 * knowledge. This can be used to verify the verifiable credential.
 */
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
export declare const LicenseCredentialSubject: {
    encode(message: LicenseCredentialSubject, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): LicenseCredentialSubject;
    fromJSON(object: any): LicenseCredentialSubject;
    toJSON(message: LicenseCredentialSubject): unknown;
    fromPartial(object: DeepPartial<LicenseCredentialSubject>): LicenseCredentialSubject;
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
