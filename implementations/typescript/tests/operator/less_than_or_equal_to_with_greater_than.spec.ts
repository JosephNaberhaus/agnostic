import { expect } from 'chai';

describe('TypeScript', () => {
  describe('Operator', () => {
    it('LessThanOrEqualToWithGreaterThan', () => {
      expect(5 <= 3).to.be.false;
    });
  });
});
