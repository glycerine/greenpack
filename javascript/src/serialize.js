const JSON5 = require('json5');
const msgpack = require('@msgpack/msgpack');
const { TIMESTAMP_EXT_TYPE, DURATION_EXT_TYPE } = require('./extensions');

// Custom extension packer for Date objects
const timestampExtPacker = {
    type: TIMESTAMP_EXT_TYPE,
    encode: (timestamp) => {
        const seconds = Math.floor(timestamp / 1000);
        const nanoseconds = (timestamp % 1000) * 1e6;
        const data = new Uint8Array(12);
        new DataView(data.buffer).setUint32(0, seconds, false);
        new DataView(data.buffer).setUint32(4, nanoseconds, false);
        return data;
    },
    decode: (data) => {
        const seconds = new DataView(data.buffer).getUint32(0, false);
        const nanoseconds = new DataView(data.buffer).getUint32(4, false);
        return seconds * 1000 + nanoseconds / 1e6;
    },
};

// Custom extension packer for Duration objects
const durationExtPacker = {
    type: DURATION_EXT_TYPE,
    encode: (duration) => {
        const data = new Uint8Array(8);
        new DataView(data.buffer).setUint32(0, duration.seconds, false);
        new DataView(data.buffer).setUint32(4, duration.nanoseconds, false);
        return data;
    },
    decode: (data) => {
        const seconds = new DataView(data.buffer).getUint32(0, false);
        const nanoseconds = new DataView(data.buffer).getUint32(4, false);
        return new Duration(seconds, nanoseconds);
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
