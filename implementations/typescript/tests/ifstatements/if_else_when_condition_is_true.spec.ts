import { expect } from 'chai';

describe('TypeScript', () => {
  describe('IfStatements', () => {
    it('IfElseWhenConditionIsTrue', () => {
      let input = true;
      let ifOutput = false;
      let elseOutput = false;
      if (input) {
        ifOutput = true;
      } else {
        elseOutput = true;
      }
      expect(ifOutput).to.be.true;
      expect(elseOutput).to.be.false;
    });
  });
});
