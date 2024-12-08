const msgpack = require('@msgpack/msgpack');
const JSON5 = require('json5');
const { TIMESTAMP_EXT_TYPE, DURATION_EXT_TYPE } = require('./extensions');
const { Duration } = require('./duration');

// Custom extension unpacker for Date objects
const timestampExtUnpacker = {
    type: TIMESTAMP_EXT_TYPE,
    decode: (data) => {
        const seconds = new DataView(data.buffer).getUint64(0, false);
        const nanoseconds = new DataView(data.buffer).getUint32(4, false);
        return new Date(seconds * 1000 + nanoseconds / 1e6);
    },
};

// Custom extension unpacker for Duration objects
const durationExtUnpacker = {
    type: DURATION_EXT_TYPE,
    decode: (data) => {
        //const seconds = new DataView(data.buffer).getUint32(0, false);
        const nanoseconds = new DataView(data.buffer).getUint64(4, false);
        return new Duration(0, nanoseconds);
    },
};

const extensionCodec = new msgpack.ExtensionCodec();
extensionCodec.register(timestampExtUnpacker);
extensionCodec.register(durationExtUnpacker);


function deserialize(data) {
    if (data instanceof Uint8Array) {
        return msgpack.decode(data, { extensionCodec });
    } else if (typeof data === 'string') {
        return JSON5.parse(data, (key, value) => {
            if (value && value.$type === 'Date') {
                return new Date(value.$value);
            }
            return value;
        });
    } else {
        throw new Error('Invalid input type. Expected Uint8Array or string.');
    }
}

module.exports = { deserialize };
