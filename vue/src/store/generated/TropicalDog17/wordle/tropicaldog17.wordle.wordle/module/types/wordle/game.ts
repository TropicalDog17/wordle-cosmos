/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "tropicaldog17.wordle.wordle";

export interface Game {
  index: string;
  secret: string;
  player: string;
  moveCount: number;
  isWin: boolean;
}

const baseGame: object = {
  index: "",
  secret: "",
  player: "",
  moveCount: 0,
  isWin: false,
};

export const Game = {
  encode(message: Game, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.secret !== "") {
      writer.uint32(18).string(message.secret);
    }
    if (message.player !== "") {
      writer.uint32(26).string(message.player);
    }
    if (message.moveCount !== 0) {
      writer.uint32(32).uint64(message.moveCount);
    }
    if (message.isWin === true) {
      writer.uint32(40).bool(message.isWin);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Game {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGame } as Game;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.secret = reader.string();
          break;
        case 3:
          message.player = reader.string();
          break;
        case 4:
          message.moveCount = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.isWin = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Game {
    const message = { ...baseGame } as Game;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.secret !== undefined && object.secret !== null) {
      message.secret = String(object.secret);
    } else {
      message.secret = "";
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = String(object.player);
    } else {
      message.player = "";
    }
    if (object.moveCount !== undefined && object.moveCount !== null) {
      message.moveCount = Number(object.moveCount);
    } else {
      message.moveCount = 0;
    }
    if (object.isWin !== undefined && object.isWin !== null) {
      message.isWin = Boolean(object.isWin);
    } else {
      message.isWin = false;
    }
    return message;
  },

  toJSON(message: Game): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.secret !== undefined && (obj.secret = message.secret);
    message.player !== undefined && (obj.player = message.player);
    message.moveCount !== undefined && (obj.moveCount = message.moveCount);
    message.isWin !== undefined && (obj.isWin = message.isWin);
    return obj;
  },

  fromPartial(object: DeepPartial<Game>): Game {
    const message = { ...baseGame } as Game;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.secret !== undefined && object.secret !== null) {
      message.secret = object.secret;
    } else {
      message.secret = "";
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = object.player;
    } else {
      message.player = "";
    }
    if (object.moveCount !== undefined && object.moveCount !== null) {
      message.moveCount = object.moveCount;
    } else {
      message.moveCount = 0;
    }
    if (object.isWin !== undefined && object.isWin !== null) {
      message.isWin = object.isWin;
    } else {
      message.isWin = false;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
