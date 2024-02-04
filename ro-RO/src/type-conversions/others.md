# Altele

### Conversia oricărui tip în șir de caractere (String)

Pentru a converti orice tip în `String`, puteți folosi pur și simplu trait-ul `ToString` pentru acel tip. În loc să faceți asta direct, ar trebui să implementați trait-ul `fmt::Display`, care va oferi automat și `ToString` și vă permite să afișați tipul cu `println!`.

1. 🌟🌟
```rust,editable
use std::fmt;

struct Point {
    x: i32,
    y: i32,
}

impl fmt::Display for Point {
    // IMPLEMENTAȚI metoda fmt
}

fn main() {
    let origin = Point { x: 0, y: 0 };
    // COMPLETAȚI spațiile goale
    assert_eq!(origin.__, "The point is (0, 0)");
    assert_eq!(format!(__), "The point is (0, 0)");

    println!("Success!");
}
```

### Parsarea unui șir de caractere (String)
2. 🌟🌟🌟 Putem folosi metoda `parse` pentru a converti un număr `i32` dintr-un șir de caractere (`String`), acest lucru se datorează faptului că `FromStr` este implementat pentru tipul `i32` în biblioteca standard: `impl FromStr for i32`
```rust,editable
// Pentru a utiliza metoda `from_str`, trebuie să introduceți acest trait în domeniul de vizibilitate curent.
use std::str::FromStr;
fn main() {
    let parsed: i32 = "5".__.unwrap();
    let turbo_parsed = "10".__.unwrap();
    let from_str = __.unwrap();
    let sum = parsed + turbo_parsed + from_str;
    assert_eq!(sum, 35);

    println!("Success!");
}
``` 


3. 🌟🌟 Putem, de asemenea, să implementăm trait-ul `FromStr` pentru tipurile noastre personalizate
```rust,editable
use std::str::FromStr;
use std::num::ParseIntError;

#[derive(Debug, PartialEq)]
struct Point {
    x: i32,
    y: i32
}

impl FromStr for Point {
    type Err = ParseIntError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let coords: Vec<&str> = s.trim_matches(|p| p == '(' || p == ')' )
                                 .split(',')
                                 .map(|x| x.trim())
                                 .collect();

        let x_fromstr = coords[0].parse::<i32>()?;
        let y_fromstr = coords[1].parse::<i32>()?;

        Ok(Point { x: x_fromstr, y: y_fromstr })
    }
}
fn main() {
    // COMPLETAȚI spațiile goale în două moduri
    // NU schimbați codul în niciun alt loc
    let p = __;
    assert_eq!(p.unwrap(), Point{ x: 3, y: 4} );

    println!("Success!");
}
```

### Deref
You can find all the examples and exercises of the `Deref` trait [here](https://practice.rs/smart-pointers/deref.html).

### Transmute
**`std::mem::transmute`** este o **funcție nesigură** care poate fi folosită pentru a reinterpreta biții unei valori de un tip ca fiind de alt tip. Ambele tipuri original și rezultat trebuie să aibă aceeași dimensiune și niciunul dintre ele nu poate fi nevalid.

**`transmute`** este echivalent semantic cu o mutare pe biți a unei valori din sursă în destinație. Copiază biții din valoarea sursă în valoarea destinație, apoi uită originalul, părând a fi echivalent cu memcpy din C sub capotă.

Așadar, **`transmute` este incredibil de nesigură!** Cel care o apelează trebuie să se asigure singur de toate aspectele de siguranță!

### Exemple
`transmute` poate fi folosită pentru a transforma un pointer într-un pointer la o funcție, acest lucru nu este portabil pe mașini în care pointer-ul la funcție și pointer-ul la date au dimensiuni diferite.

```rust,editable
fn foo() -> i32 {
    0
}

fn main() {
    let pointer = foo as *const ();
    let function = unsafe {
        std::mem::transmute::<*const (), fn() -> i32>(pointer)
    };
    assert_eq!(function(), 0);
}
```

2. Extinderea unei perioade de valabilitate (lifetime) sau scurtarea unei perioade de valabilitate a unei invariante este o utilizare avansată a `transmute`.
```rust,editable
struct R<'a>(&'a i32);
unsafe fn extend_lifetime<'b>(r: R<'b>) -> R<'static> {
    std::mem::transmute::<R<'b>, R<'static>>(r)
}

unsafe fn shorten_invariant_lifetime<'b, 'c>(r: &'b mut R<'static>)
                                             -> &'b mut R<'c> {
    std::mem::transmute::<&'b mut R<'static>, &'b mut R<'c>>(r)
}
```

3. În loc să folosiți `transmute`, puteți utiliza unele alternative în schimb.
```rust,editable
fn main() {
    /*Transformarea de biți bruti (&[u8]) în u32, f64, etc.: */
    let raw_bytes = [0x78, 0x56, 0x34, 0x12];

    let num = unsafe { std::mem::transmute::<[u8; 4], u32>(raw_bytes) };

    // Folosiți `u32::from_ne_bytes` în schimb
    let num = u32::from_ne_bytes(raw_bytes);
    // Sau folosiți `u32::from_le_bytes` sau `u32::from_be_bytes` pentru a specifica endianitatea
    let num = u32::from_le_bytes(raw_bytes);
    assert_eq!(num, 0x12345678);
    let num = u32::from_be_bytes(raw_bytes);
    assert_eq!(num, 0x78563412);

    /*Transformarea unui pointer într-un usize: */
    let ptr = &0;
    let ptr_num_transmute = unsafe { std::mem::transmute::<&i32, usize>(ptr) };

    // Folosiți o conversie `as` în schimb
    let ptr_num_cast = ptr as *const i32 as usize;

     /*Transformarea unui &mut T într-un &mut U: */
    let ptr = &mut 0;
    let val_transmuted = unsafe { std::mem::transmute::<&mut i32, &mut u32>(ptr) };

    // Acum, combinați `as` și reîmprumutați - observați concatenarea `as`
    // `as` nu este tranzitiv
    let val_casts = unsafe { &mut *(ptr as *mut i32 as *mut u32) };

     /*Transformarea unui &str într-un &[u8]: */
    // Aceasta nu este o modalitate bună de a face acest lucru.
    let slice = unsafe { std::mem::transmute::<&str, &[u8]>("Rust") };
    assert_eq!(slice, &[82, 117, 115, 116]);

    // Ați putea folosi `str::as_bytes`
    let slice = "Rust".as_bytes();
    assert_eq!(slice, &[82, 117, 115, 116]);

    // Sau, pur și simplu, utilizați un șir de octeți, dacă
    // aveți control asupra literalului de șir
    assert_eq!(b"Rust", &[82, 117, 115, 116]);
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
