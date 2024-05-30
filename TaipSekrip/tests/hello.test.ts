import { sayHi } from "../Hello";

describe('sayHi', function(): void {
  it('Harus return berupa, Hello Taipskrip, aku sedang belajar kamu', function () {
    expect(sayHi('aku sedang belajar kamu')).toBe('Hello Taipsekrip,  aku sedang belajar kamu');
  });
});