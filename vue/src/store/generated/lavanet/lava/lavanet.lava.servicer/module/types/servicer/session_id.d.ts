import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "lavanet.lava.servicer";
export interface SessionID {
    num: number;
}
export declare const SessionID: {
    encode(message: SessionID, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SessionID;
    fromJSON(object: any): SessionID;
    toJSON(message: SessionID): unknown;
    fromPartial(object: DeepPartial<SessionID>): SessionID;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};