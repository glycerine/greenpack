const msgpack = require('@msgpack/msgpack');
const JSON5 = require('json5');
const { TIMESTAMP_EXT_TYPE, DURATION_EXT_TYPE } = require('./extensions');
const { Duration } = require('./duration');

// Custom extension unpacker for Date objects
const timestampExtUnpacker = msgpack.createExtensionCodec({
    type: TIMESTAMP_EXT_TYPE,
    encode: () => { throw new Error('Encoding not supported for timestamp ext type'); },
    decode: (data) => {
        const seconds = new DataView(data.buffer).getUint32(0, false);
        const nanoseconds = new DataView(data.buffer).getUint32(4, false);
        return new Date(seconds * 1000 + nanoseconds / 1e6);
    },
});

// Custom extension unpacker for Duration objects
const durationExtUnpacker = msgpack.createExtensionCodec({
    type: DURATION_EXT_TYPE,
    encode: () => { throw new Error('Encoding not supported for duration ext type'); },
    decode: (data) => {
        const seconds = new DataView(data.buffer).getUint32(0, false);
        const nanoseconds = new DataView(data.buffer).getUint32(4, false);
        return new Duration(seconds, nanoseconds);
    },
});

function deserialize(data) {
    if (data instanceof Uint8Array) {
        return msgpack.decode(data, { extensionCodec: msgpack.ExtensionCodec.concat([timestampExtUnpacker, durationExtUnpacker]) });
    } else if (typeof data === 'string') {
        return JSON5.parse(data);
    } else {
        throw new Error('Invalid input type. Expected Uint8Array or string.');
    }
}

module.exports = { deserialize };
