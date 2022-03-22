import { expect } from 'chai';

describe('TypeScript', () => {
  describe('Operator', () => {
    it('EqualWithStringWhenEqual', () => {
      expect('test' == 'test').to.be.true;
    });
  });
});
