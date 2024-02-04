# Numere

### Numere întregi

1. 🌟 

> Sfat: Dacă nu atribuim explicit un tip unei variabile, compilatorul îl va deduce pentru noi.

```rust,editable

// Eliminați ceva pentru a face să funcționeze
fn main() {
    let x: i32 = 5;
    let mut y: u32 = 5;

    y = x;
    
    let z = 10; // Ce tip are z ?

    println!("Success!");
}
```

2. 🌟
```rust,editable

// Umple spațiul liber
fn main() {
    let v: u16 = 38_u8 as __;

    println!("Success!");
}
```

3. 🌟🌟🌟  

> Sfat: Dacă nu atribuim explicit un tip unei variabile, compilatorul îl va deduce pentru noi.

```rust,editable

// Modificați `assert_eq!` pentru a face să funcționeze
fn main() {
    let x = 5;
    assert_eq!("u32".to_string(), type_of(&x));

    println!("Success!");
}

// Obțineți tipul variabilei date, returnați o reprezentare sub formă de șir a tipului, de exemplu "i8", "u8", "i32", "u32".
fn type_of<T>(_: &T) -> String {
    format!("{}", std::any::type_name::<T>())
}
```

4. 🌟🌟 
```rust,editable

// Completați spațiile goale pentru a face să funcționeze
fn main() {
    assert_eq!(i8::MAX, __); 
    assert_eq!(u8::MAX, __); 

    println!("Success!");
}
```

5. 🌟🌟 
```rust,editable

// Remediați erorile și panicele pentru a face să funcționeze.
fn main() {
   let v1 = 251_u8 + 8;
   let v2 = i8::checked_add(251, 8).unwrap();
   println!("{},{}",v1,v2);
}
```

6. 🌟🌟
```rust,editable

// Modificați `assert!` pentru a face să funcționeze
fn main() {
    let v = 1_024 + 0xff + 0o77 + 0b1111_1111;
    assert!(v == 1579);

    println!("Success!");
}
```


### Floating-Point
7. 🌟

```rust,editable

// Completați spațiul gol pentru a face să funcționeze

fn main() {
    let x = 1_000.000_1; // ?
    let y: f32 = 0.12; // f32
    let z = 0.01_f64; // f64

    assert_eq!(type_of(&x), "__".to_string());
    println!("Success!");
}

fn type_of<T>(_: &T) -> String {
    format!("{}", std::any::type_name::<T>())
}
```

8. 🌟🌟 Faceți să funcționeze în două moduri distincte.

```rust,editable

fn main() {
    assert!(0.1+0.2==0.3);

    println!("Success!");
}
```

### Range
9. 🌟🌟 Două obiective: 1. Modificați `assert!` pentru a face codul să funcționeze 2. Faceți `println!` să afișeze: 97 - 122

```rust,editable
fn main() {
    let mut sum = 0;
    for i in -3..2 {
        sum += i
    }

    assert!(sum == -3);

    for c in 'a'..='z' {
        println!("{}",c);
    }
}
```

10. 🌟🌟 
```rust,editable

// Umpleți spațiile goale
use std::ops::{Range, RangeInclusive};
fn main() {
    assert_eq!((1..__), Range{ start: 1, end: 5 });
    assert_eq!((1..__), RangeInclusive::new(1, 5));

    println!("Success!");
}
```

### Computations

11. 🌟 
```rust,editable

// Umpleți spațiile goale și rezolvați erorile
fn main() {
    // Adunarea numerelor întregi
    assert!(1u32 + 2 == __);

    // Scăderea numerelor întregi
    assert!(1i32 - 2 == __);
    assert!(1u8 - 2 == -1); 
    
    assert!(3 * 50 == __);

    assert!(9.6 / 3.2 == 3.0); // asigură-te că 9.6 împărțit la 3.2 este egal cu 3.0; // eroare! fă-l să funcționeze

    assert!(24 % 5 == __);
    // Logica booleană cu scurtcircuitare
    assert!(true && false == __);
    assert!(true || false == __);
    assert!(!true == __);

    // Operații pe biți
    println!("0011 AND 0101 is {:04b}", 0b0011u32 & 0b0101);
    println!("0011 OR 0101 is {:04b}", 0b0011u32 | 0b0101);
    println!("0011 XOR 0101 is {:04b}", 0b0011u32 ^ 0b0101);
    println!("1 << 5 is {}", 1u32 << 5);
    println!("0x80 >> 2 is 0x{:x}", 0x80u32 >> 2);
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice)(sub calea soluțiilor), dar utilizați-le doar atunci când aveți nevoie.
