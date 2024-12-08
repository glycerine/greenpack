const { serializeJSON5 } = require('../src/serialize');

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

});
