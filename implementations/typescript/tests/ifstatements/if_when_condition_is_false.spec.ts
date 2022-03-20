import { expect } from 'chai';

describe('TypeScript', () => {
  describe('IfStatements', () => {
    it('IfWhenConditionIsFalse', () => {
      let input = false;
      let output = false;
      if (input) {
      output = false;
      }
      expect(output).to.be.false;
    });
  });
});
