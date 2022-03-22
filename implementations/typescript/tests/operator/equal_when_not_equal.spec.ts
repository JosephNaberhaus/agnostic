import { expect } from 'chai';

describe('TypeScript', () => {
  describe('Operator', () => {
    it('EqualWhenNotEqual', () => {
      expect(1 == 42).to.be.false;
    });
  });
});
