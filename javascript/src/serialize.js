const JSON5 = require('json5');
const msgpack = require('@msgpack/msgpack');

function serialize(obj) {
    return msgpack.encode(obj);
}

function serializeJSON5(obj) {
    return JSON5.stringify(obj, null, 2);
}

module.exports = {
    serialize,
    serializeJSON5,
};
