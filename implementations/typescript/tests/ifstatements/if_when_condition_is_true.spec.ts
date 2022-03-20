import { expect } from 'chai';

describe('TypeScript', () => {
  describe('IfStatements', () => {
    it('IfWhenConditionIsTrue', () => {
      let input = true;
      let output = false;
      if (input) {
      output = true;
      }
      expect(output).to.be.true;
    });
  });
});
