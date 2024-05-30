typescript
import { hiUser } from "../userObject";
import { User } from "../userObject";

// Membuat objek user yang sama dengan yang ada di file userObject
const user: User = {
    name: "Hildan",
    class: "4IKRA",
    studentId: 11111
};

describe('sayHi', function(): void {
    it('should return "Hello Hildan from 4IKRA with studentId: 11111"', function () {
        expect(hiUser(user)).toBe("Hello Hildan from 4IKRA with studentId: 11111");
    });
});