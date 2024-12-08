const msgpack = require('@msgpack/msgpack');
const JSON5 = require('json5');

function deserialize(data) {
    if (data instanceof Uint8Array) {
        return msgpack.decode(data);
    } else if (typeof data === 'string') {
        return JSON5.parse(data);
    } else {
        throw new Error('Invalid input type. Expected Uint8Array or string.');
    }
}

module.exports = { deserialize };
