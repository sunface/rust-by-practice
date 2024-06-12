# Referințe și Împrumutare

### Referințe
1. 🌟
```rust,editable

fn main() {
   let x = 5;
   // Completați spațiile libere
   let p = __;

   println!("the memory address of x is {:p}", p); // Una dintre posibilele ieșiri: 0x16fa3ac84
}
```

2. 🌟
```rust,editable

fn main() {
    let x = 5;
    let y = &x;

    // Modificați doar această linie
    assert_eq!(5, y);

    println!("Success!");
}
```

3. 🌟
```rust,editable

// Corectați eroarea
fn main() {
    let mut s = String::from("hello, ");

    borrow_object(s);

    println!("Success!");
}

fn borrow_object(s: &String) {}
```

4. 🌟
```rust,editable

// Corectați eroarea
fn main() {
    let mut s = String::from("hello, ");

    push_str(s);

    println!("Success!");
}

fn push_str(s: &mut String) {
    s.push_str("world")
}
```

5. 🌟🌟
```rust,editable

fn main() {
    let mut s = String::from("hello, ");

    // Completați spațiile libere pentru a face ca codul să funcționeze
    let p = __;
    
    p.push_str("world");

    println!("Success!");
}
```

#### Ref
`ref` poate fi folosit pentru a obține referințe la o valoare, similar cu `&`.

6. 🌟🌟🌟
```rust,editable

fn main() {
    let c = '中';

    let r1 = &c;
    // Completați spațiile libere, fără a schimba alte coduri
    let __ r2 = c;

    assert_eq!(*r1, *r2);
    
    // Verificați egalitatea celor două șiruri de adrese
    assert_eq!(get_addr(r1),get_addr(r2));

    println!("Success!");
}

// Obțineți șirul de adresă de memorie
fn get_addr(r: &char) -> String {
    format!("{:p}", r)
}
```

### Reguli de Împrumut
7. 🌟
```rust,editable

// Eliminați ceva pentru a face ca codul să funcționeze
// Nu ștergeți o linie întreagă!
fn main() {
    let mut s = String::from("hello");

    let r1 = &mut s;
    let r2 = &mut s;

    println!("{}, {}", r1, r2);

    println!("Success!");
}
```

#### Mutabilitate
8. 🌟 Eroare: Împrumutați un obiect imutabil ca mutabil
```rust,editable

fn main() {
    // Corectați eroarea modificând această linie
    let  s = String::from("hello, ");

    borrow_object(&mut s);

    println!("Success!");
}

fn borrow_object(s: &mut String) {}
```

9. 🌟🌟 Ok: Împrumutați un obiect mutabil ca imutabil
```rust,editable

// Acest cod nu are erori!
fn main() {
    let mut s = String::from("hello, ");

    borrow_object(&s);
    
    s.push_str("world");

    println!("Success!");
}

fn borrow_object(s: &String) {}
```

### NLL (Non-Lexical Lifetimes)
10. 🌟🌟
```rust,editable

// Comentați o linie pentru a face ca codul să funcționeze
fn main() {
    let mut s = String::from("hello, ");

    let r1 = &mut s;
    r1.push_str("world");
    let r2 = &mut s;
    r2.push_str("!");
    
    println!("{}",r1);
}
```

11. 🌟🌟
```rust,editable

fn main() {
    let mut s = String::from("hello, ");

    let r1 = &mut s;
    let r2 = &mut s;

    // Adăugați o linie mai jos pentru a obține o eroare de compilator: nu se poate împrumuta `s` ca mutabil de mai multe ori simultan
    // Nu puteți utiliza r1 și r2 în același timp
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)