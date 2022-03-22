import { expect } from 'chai';

describe('TypeScript', () => {
  describe('Operator', () => {
    it('GreaterThanOrEqualToWhenLessThan', () => {
      expect(3 >= 5).to.be.false;
    });
  });
});
