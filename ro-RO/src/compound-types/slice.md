# Felie (Slice)
Feliile sunt similare cu tablourile, dar lungimea lor nu este cunoscută în timpul compilării, așa că nu poți utiliza feliile direct.

1. 🌟🌟 Aici, atât [i32] cât și str sunt tipuri de felii, dar utilizarea directă va provoca erori. Trebuie să folosiți referința feliei în schimb: &[i32], &str.
```rust,editable

// Remediază erorile, NU adăuga noi linii!
fn main() {
    let arr = [1, 2, 3];
    let s1: [i32] = arr[0..2];

    let s2: str = "hello, world" as str;

    println!("Success!");
}
```

O referință la o felie este un obiect format din două cuvinte, din motive de simplitate, de acum înainte vom folosi doar "felia" în loc de "referința la felie". Primul cuvânt este un pointer către date, iar al doilea cuvânt este lungimea feliei. Dimensiunea cuvântului este aceeași cu usize, determinată de arhitectura procesorului, de exemplu, 64 de biți pe un x86-64. Feliile pot fi utilizate pentru a împrumuta o secțiune a unui tablou și au semnătura de tip &[T].

2. 🌟🌟🌟
```rust,editable

fn main() {
    let arr: [char; 3] = ['中', '国', '人'];

    let slice = &arr[..2];
    
    // Modifică '8' pentru a face codul să funcționeze
    // SFAT: slice (referința la feliu) NU ESTE un tablou, dacă ar fi un tablou, atunci assert! ar fi trecut: Fiecare dintre cele două caractere '中' și '国' ocupă 4 octeți, 2 * 4 = 8
    assert!(std::mem::size_of_val(&slice) == 8);

    println!("Success!");
}
```

3. 🌟🌟
```rust,editable

fn main() {
    let arr: [i32; 5] = [1, 2, 3, 4, 5];
    // Completează spațiile goale pentru a face codul să funcționeze
    let slice: __ = __;
    assert_eq!(slice, &[2, 3, 4]);

    println!("Success!");
}
```

### Feliile de șiruri (String slices)
4. 🌟 
```rust,editable

fn main() {
    let s = String::from("hello");

    let slice1 = &s[0..2];
    // Completează spațiul gol pentru a face codul să funcționeze, NU FOLOSI 0..2 din nou
    let slice2 = &s[__];

    assert_eq!(slice1, slice2);

    println!("Success!");
}
```

5. 🌟
```rust,editable

fn main() {
    let s = "你好，世界";
    // Modifică această linie pentru a face codul să funcționeze
    let slice = &s[0..2];

    assert!(slice == "你");

    println!("Success!");
}
```

6. 🌟🌟 `&String` can be implicitly converted into `&str`.
```rust,editable

// Remediază erorile
fn main() {
    let mut s = String::from("hello world");

    // Aici, &s are tipul &String, dar first_word necesită un tip &str.
    // Funcționează deoarece &String poate fi convertit implicit la &str. Dacă vrei să afli mai multe, acest lucru se numește "coerciție Deref".
    let word = first_word(&s);

    s.clear(); // eroare!

    println!("the first word is: {}", word);
}
fn first_word(s: &str) -> &str {
    &s[..1]
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie.