# Tuplu
1. 🌟 Elementele dintr-un tuplu pot avea diferite tipuri. Semnătura de tip tuplu este `(T1, T2, ...)`, unde `T1`, `T2` sunt tipurile de membri ai tuplului.
```rust,editable

fn main() {
    let _t0: (u8,i16) = (0, -1);
    // Tuples can be tuple's members
    let _t1: (u8, (i16, u32)) = (0, (-1, 1));
    // Fill the blanks to make the code work
    let t: (u8, __, i64, __, __) = (1u8, 2u16, 3i64, "hello", String::from(", world"));

    println!("Success!");
}
```

2. 🌟 Membrii pot fi extrași din tuplu folosind indexare.
```rust,editable

// Faceți să funcționeze
fn main() {
    let t = ("i", "am", "sunface");
    assert_eq!(t.1, "sunface");

    println!("Success!");
}
```

3. 🌟 Tuplurile lungi nu pot fi imprimate
```rust,editable

// Remediați eroarea
fn main() {
    let too_long_tuple = (1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13);
    println!("too long tuple: {:?}", too_long_tuple);
}
```

4. 🌟 Destructurarea tuplului cu model.
```rust,editable

fn main() {
    let tup = (1, 6.4, "hello");

    // Completați spațiul liber pentru ca codul să funcționeze
    let __ = tup;

    assert_eq!(x, 1);
    assert_eq!(y, "hello");
    assert_eq!(z, 6.4);

    println!("Success!");
}
```

5. 🌟🌟 Destructurați sarcinile.
```rust,editable
fn main() {
    let (x, y, z);

    // Completați spațiul liber
    __ = (1, 2, 3);
    
    assert_eq!(x, 3);
    assert_eq!(y, 1);
    assert_eq!(z, 2);

    println!("Success!");
}
```

6. 🌟🌟 Tuplurile pot fi folosite ca argumente de funcție și valori returnate
```rust,editable

fn main() {
    // Completați spațiul liber, aveți nevoie de câteva calcule aici.
    let (x, y) = sum_multiply(__);

    assert_eq!(x, 5);
    assert_eq!(y, 6);

    println!("Success!");
}

fn sum_multiply(nums: (i32, i32)) -> (i32, i32) {
    (nums.0 + nums.1, nums.0 * nums.1)
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie.
