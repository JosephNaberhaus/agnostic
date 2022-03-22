import { expect } from 'chai';

describe('TypeScript', () => {
  describe('Operator', () => {
    it('LessThanWhenGreaterThan', () => {
      expect(6 < 5).to.be.false;
    });
  });
});
