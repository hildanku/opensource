interface User {
    name: string,
    class: string,
    studentId: number
}

let user: User = {
    name: "Hildan",
    class: "4IKRA",
    studentId: 11111
}

export function hiUser(user: User) {
    return `Hello ${user.name} from ${user.class} with studentId: ${user.studentId}`;
}

// export function sayHi(message: string): String {
//     return `Hello Taipsekrip,  ${message}`;
// }