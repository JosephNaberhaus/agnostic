import { expect } from 'chai';

describe('TypeScript', () => {
  describe('IfStatements', () => {
    it('IfElseWhenConditionIsFalse', () => {
      let input = false;
      let ifOutput = false;
      let elseOutput = false;
      if (input) {
        ifOutput = true;
      } else {
        elseOutput = true;
      }
      expect(ifOutput).to.be.false;
      expect(elseOutput).to.be.true;
    });
  });
});
