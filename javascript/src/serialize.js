const JSON5 = require('json5');
const msgpack = require('@msgpack/msgpack');
const { TIMESTAMP_EXT_TYPE, DURATION_EXT_TYPE } = require('./extensions');

// Custom extension packer for Date objects
const timestampExtPacker = {
    type: TIMESTAMP_EXT_TYPE,
    encode: (timestamp) => {
        const seconds = BigInt(Math.floor(timestamp / 1000));
        const nanoseconds = (timestamp % 1000) * 1e6;
        const data = new Uint8Array(12);
        new DataView(data.buffer).setBigUint64(0, seconds, false); // false for big-endian, msgpack requires.
        new DataView(data.buffer).setUint32(8, nanoseconds, false);
        return data;
    },
    decode: (data) => {
        const seconds = new DataView(data.buffer).getBigUint64(0, false);
        const nanoseconds = new DataView(data.buffer).getUint32(8, false);
        const fullPrecision = seconds * BigInt(1e9) + BigInt(nanoseconds);
        
        // lossy: the least significant digit of the nanoseconds is lost.
        // Does not matter if javascript is giving us msec precision anyway.
        return Number(fullPrecision)/1e6; 
    },
};

// Custom extension packer for Duration objects
const durationExtPacker = {
    type: DURATION_EXT_TYPE,
    encode: (duration) => {
        const data = new Uint8Array(8);
        const nanoseconds = BigInt(duration.seconds) * BigInt(1e9) + BigInt(duration.nanoseconds);
        new DataView(data.buffer).setBigUint64(0, nanoseconds, false);
        return data;
    },
    decode: (data) => {
        const nanoseconds = new DataView(data.buffer).getBigUint64(0, false);
        const seconds = Number(nanoseconds / BigInt(1e9));
        const ns      = Number(nanoseconds % BigInt(1e9));
        return new Duration(seconds, ns);
    },
};

const extensionCodec = new msgpack.ExtensionCodec();
extensionCodec.register(timestampExtPacker);
extensionCodec.register(durationExtPacker);

function serialize(obj) {
    return msgpack.encode(obj, {
        extensionCodec,
        context: {
            encodeDate: (date) => date.getTime(),
        },
    });
}

function serializeJSON5(obj) {
    return JSON5.stringify(obj, (key, value) => {
        if (value instanceof Date) {
            return {
                $type: 'Date',
                $value: value.getTime(),
            };
        }
        return value;
    }, 2);
}

module.exports = {
    serialize,
    serializeJSON5,
};
