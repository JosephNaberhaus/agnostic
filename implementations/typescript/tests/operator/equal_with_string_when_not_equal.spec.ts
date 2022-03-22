import { expect } from 'chai';

describe('TypeScript', () => {
  describe('Operator', () => {
    it('EqualWithStringWhenNotEqual', () => {
      expect('test' == 'hello').to.be.false;
    });
  });
});
