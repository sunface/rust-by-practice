# Tablou (Array)
Tipul de tablou este [T; Lungime], așa cum poți observa, lungimea tabloului este parte a semnăturii lor de tip. Prin urmare, lungimea lor trebuie să fie cunoscută în timpul compilării.

De exemplu, nu poți inițializa un tablou în felul următor:
```rust
fn init_arr(n: i32) {
    let arr = [1; n];
}
```

Acest lucru va provoca o eroare, deoarece compilatorul nu are nicio idee despre dimensiunea exactă a tabloului în timpul compilării.

1. 🌟 
```rust,editable

fn main() {
    // Completează spațiul gol cu tipul de tablou potrivit
    let arr: __ = [1, 2, 3, 4, 5];

    // Modifică codul de mai jos pentru a-l face funcțional
    assert!(arr.len() == 4);

    println!("Success!");
}
```

2. 🌟🌟
```rust,editable

fn main() {
    // Putem să ignorăm părți ale tipului de tablou sau chiar întregul tip, să lăsăm compilatorul să-l deducă pentru noi
    let arr0 = [1, 2, 3];
    let arr: [_; 3] = ['a', 'b', 'c'];
    
    // Completează spațiul gol
    // Tablourile sunt alocate pe stivă, std::mem::size_of_val returnează numărul de octeți pe care îi ocupă un tablou
    // Un caracter în Rust ocupă 4 octeți: caracter Unicode
    assert!(std::mem::size_of_val(&arr) == __);

    println!("Success!");
}
```

3. 🌟 Toate elementele dintr-un tablou pot fi inițializate la aceeași valoare în același timp.

```rust,editable

fn main() {
    // Completează spațiul gol
    let list: [i32; 100] = __ ;

    assert!(list[0] == 1);
    assert!(list.len() == 100);

    println!("Success!");
}
```

4. 🌟 Toate elementele dintr-un tablou trebuie să fie de același tip.
```rust,editable

fn main() {
    // Remediază eroarea
    let _arr = [1, 2, '3'];

    println!("Success!");
}
```

5. 🌟 Indexarea începe de la 0.
```rust,editable

fn main() {
    let arr = ['a', 'b', 'c'];
    
    let ele = arr[1]; // Modifică doar această linie pentru a face ca codul să funcționeze!

    assert!(ele == 'a');

    println!("Success!");
}
```

6. 🌟 Out of bounds indexing causes `panic`.
```rust,editable

// Remediază eroarea
fn main() {
    let names = [String::from("Sunfei"), "Sunface".to_string()];
    
    // Get returnează un Option<T>, este sigur să-l folosești
    let name0 = names.get(0).unwrap();

    // Dar indexarea nu este sigură
    let _name1 = &names[2];

    println!("Success!");
}

```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie.