import { expect } from 'chai';

describe('TypeScript', () => {
  describe('Operator', () => {
    it('LessThanOrEqualToWhenGreaterThan', () => {
      expect(5 <= 3).to.be.false;
    });
  });
});
