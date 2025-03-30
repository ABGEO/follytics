import { DEFAULTS } from '.';

describe('constants', () => {
  it('should equal', () => {
    expect(DEFAULTS.width).toEqual(960);
    expect(DEFAULTS.height).toEqual(500);
  });
});
