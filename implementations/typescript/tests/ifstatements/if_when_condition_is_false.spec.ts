import { expect } from 'chai';

describe('TypeScript', () => {
  describe('IfStatements', () => {
    it('IfWhenConditionIsFalse', () => {
      let input = false;
      let output = false;
      if (input) {
        output = true;
      }
      expect(output).to.be.false;
    });
  });
});
