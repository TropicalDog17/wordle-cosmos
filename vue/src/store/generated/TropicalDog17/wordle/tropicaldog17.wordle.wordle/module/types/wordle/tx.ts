/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "tropicaldog17.wordle.wordle";

export interface MsgCreateGame {
  creator: string;
  player: string;
  secret: string;
}

export interface MsgCreateGameResponse {
  gameIndex: string;
}

export interface MsgDoGuess {
  creator: string;
  gameIndex: string;
  letter: string;
}

export interface MsgDoGuessResponse {
  guessState: string;
  win: string;
}

const baseMsgCreateGame: object = { creator: "", player: "", secret: "" };

export const MsgCreateGame = {
  encode(message: MsgCreateGame, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.player !== "") {
      writer.uint32(18).string(message.player);
    }
    if (message.secret !== "") {
      writer.uint32(26).string(message.secret);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateGame } as MsgCreateGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.player = reader.string();
          break;
        case 3:
          message.secret = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateGame {
    const message = { ...baseMsgCreateGame } as MsgCreateGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = String(object.player);
    } else {
      message.player = "";
    }
    if (object.secret !== undefined && object.secret !== null) {
      message.secret = String(object.secret);
    } else {
      message.secret = "";
    }
    return message;
  },

  toJSON(message: MsgCreateGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.player !== undefined && (obj.player = message.player);
    message.secret !== undefined && (obj.secret = message.secret);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateGame>): MsgCreateGame {
    const message = { ...baseMsgCreateGame } as MsgCreateGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = object.player;
    } else {
      message.player = "";
    }
    if (object.secret !== undefined && object.secret !== null) {
      message.secret = object.secret;
    } else {
      message.secret = "";
    }
    return message;
  },
};

const baseMsgCreateGameResponse: object = { gameIndex: "" };

export const MsgCreateGameResponse = {
  encode(
    message: MsgCreateGameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.gameIndex !== "") {
      writer.uint32(10).string(message.gameIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateGameResponse } as MsgCreateGameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateGameResponse {
    const message = { ...baseMsgCreateGameResponse } as MsgCreateGameResponse;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = String(object.gameIndex);
    } else {
      message.gameIndex = "";
    }
    return message;
  },

  toJSON(message: MsgCreateGameResponse): unknown {
    const obj: any = {};
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateGameResponse>
  ): MsgCreateGameResponse {
    const message = { ...baseMsgCreateGameResponse } as MsgCreateGameResponse;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = object.gameIndex;
    } else {
      message.gameIndex = "";
    }
    return message;
  },
};

const baseMsgDoGuess: object = { creator: "", gameIndex: "", letter: "" };

export const MsgDoGuess = {
  encode(message: MsgDoGuess, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.gameIndex !== "") {
      writer.uint32(18).string(message.gameIndex);
    }
    if (message.letter !== "") {
      writer.uint32(26).string(message.letter);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDoGuess {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDoGuess } as MsgDoGuess;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.gameIndex = reader.string();
          break;
        case 3:
          message.letter = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDoGuess {
    const message = { ...baseMsgDoGuess } as MsgDoGuess;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = String(object.gameIndex);
    } else {
      message.gameIndex = "";
    }
    if (object.letter !== undefined && object.letter !== null) {
      message.letter = String(object.letter);
    } else {
      message.letter = "";
    }
    return message;
  },

  toJSON(message: MsgDoGuess): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    message.letter !== undefined && (obj.letter = message.letter);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDoGuess>): MsgDoGuess {
    const message = { ...baseMsgDoGuess } as MsgDoGuess;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = object.gameIndex;
    } else {
      message.gameIndex = "";
    }
    if (object.letter !== undefined && object.letter !== null) {
      message.letter = object.letter;
    } else {
      message.letter = "";
    }
    return message;
  },
};

const baseMsgDoGuessResponse: object = { guessState: "", win: "" };

export const MsgDoGuessResponse = {
  encode(
    message: MsgDoGuessResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.guessState !== "") {
      writer.uint32(10).string(message.guessState);
    }
    if (message.win !== "") {
      writer.uint32(18).string(message.win);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDoGuessResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDoGuessResponse } as MsgDoGuessResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.guessState = reader.string();
          break;
        case 2:
          message.win = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDoGuessResponse {
    const message = { ...baseMsgDoGuessResponse } as MsgDoGuessResponse;
    if (object.guessState !== undefined && object.guessState !== null) {
      message.guessState = String(object.guessState);
    } else {
      message.guessState = "";
    }
    if (object.win !== undefined && object.win !== null) {
      message.win = String(object.win);
    } else {
      message.win = "";
    }
    return message;
  },

  toJSON(message: MsgDoGuessResponse): unknown {
    const obj: any = {};
    message.guessState !== undefined && (obj.guessState = message.guessState);
    message.win !== undefined && (obj.win = message.win);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDoGuessResponse>): MsgDoGuessResponse {
    const message = { ...baseMsgDoGuessResponse } as MsgDoGuessResponse;
    if (object.guessState !== undefined && object.guessState !== null) {
      message.guessState = object.guessState;
    } else {
      message.guessState = "";
    }
    if (object.win !== undefined && object.win !== null) {
      message.win = object.win;
    } else {
      message.win = "";
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DoGuess(request: MsgDoGuess): Promise<MsgDoGuessResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse> {
    const data = MsgCreateGame.encode(request).finish();
    const promise = this.rpc.request(
      "tropicaldog17.wordle.wordle.Msg",
      "CreateGame",
      data
    );
    return promise.then((data) =>
      MsgCreateGameResponse.decode(new Reader(data))
    );
  }

  DoGuess(request: MsgDoGuess): Promise<MsgDoGuessResponse> {
    const data = MsgDoGuess.encode(request).finish();
    const promise = this.rpc.request(
      "tropicaldog17.wordle.wordle.Msg",
      "DoGuess",
      data
    );
    return promise.then((data) => MsgDoGuessResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
