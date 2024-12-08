const { serializeJSON5 } = require('../src/serialize');
const { Duration } = require('../src/duration');

describe('serializeJSON5', () => {
    test('serializes a simple object', () => {
        const input = {
            name: 'Alice',
            age: 30,
            city: 'New York',
        };

        const expected = `{
  name: 'Alice',
  age: 30,
  city: 'New York',
}`;

        const actual = serializeJSON5(input);

        console.log("actual = ", actual);
        
        expect(actual).toEqual(expected);
    });

    test('serializes NaN and Infinity', () => {
        const input = {
            nan: NaN,
            infinity: Infinity,
            negativeInfinity: -Infinity,
        };

        const expected = `{
  nan: NaN,
  infinity: Infinity,
  negativeInfinity: -Infinity,
}`;

        const actual = serializeJSON5(input);

        expect(actual).toEqual(expected);
    });

    test('serializes Date objects', () => {
        const input = {
            timestamp: new Date(1623430881123),
        };

        const expected = new Uint8Array([
            0x81, 0xd7, 0x01, 0x0c, 0x00, 0x00, 0x60, 0xc7, 0xc9, 0x91, 0x00, 0x00, 0x00, 0x00, 0x47, 0x53,
        ]);

        const actual = serialize(input);

        expect(actual).toEqual(expected);
    });

    test('serializes Duration objects', () => {
        const input = {
            duration: new Duration(123, 456000000),
        };

        const expected = new Uint8Array([
            0x81, 0xd7, 0x02, 0x08, 0x00, 0x00, 0x00, 0x7b, 0x00, 0x00, 0x00, 0x00, 0x1b, 0x48, 0x70, 0x00,
        ]);

        const actual = serialize(input);

        expect(actual).toEqual(expected);
    });
});
