# panic!
Cel mai simplu mecanism de gestionare a erorilor este să folosim panic. Aceasta doar afișează un mesaj de eroare și începe să deruleze stiva, în cele din urmă ieșind din thread-ul curent:

- dacă panică apare în thread-ul main, atunci programul va fi închis.
- dacă apare într-un thread generat, atunci acel thread va fi terminat, dar programul nu va fi întrerupt.


1. 🌟🌟
```rust,editable

// UMPLI blank-urile
fn drink(beverage: &str) {
    if beverage == "lemonade" {
        println!("Success!");
        // IMPLEMENTEAZĂ codul de mai jos
        __
     }

    println!("Exercise Failed if printing out this line!");
}

fn main() {
    drink(__);

    println!("Exercise Failed if printing out this line!");
}
```

## Cazuri comune de panică
2. 🌟🌟
```rust,editable
// Faceți codul să funcționeze rezolvând toate panicele
fn main() {
    assert_eq!("abc".as_bytes(), [96, 97, 98]);

    let v = vec![1, 2, 3];
    let ele = v[3];
    // "unwrap" poate genera panică când "get" returnează "None"
    let ele = v.get(3).unwrap();

    // Uneori, compilatorul nu poate găsi erorile de overflow pentru tine în timpul compilării, astfel că va apărea panică
    let v = production_rate_per_hour(2);

    // din același motiv ca mai sus, trebuie să îl învelim într-o funcție pentru a provoca panică
    divide(15, 0);

    println!("Success!")
}

fn divide(x:u8, y:u8) {
    println!("{}", x / y)
}

fn production_rate_per_hour(speed: u8) -> f64 {
    let cph: u8 = 221;
    match speed {
        1..=4 => (speed * cph) as f64,
        5..=8 => (speed * cph) as f64 * 0.9,
        9..=10 => (speed * cph) as f64 * 0.77,
        _ => 0 as f64,
    }
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    (production_rate_per_hour(speed) / 60 as f64) as u32
}
```

### Stivă de apel detaliată
În mod implicit, derularea stivei va arăta ceva de genul:
```shell
thread 'main' panicked at 'index out of bounds: the len is 3 but the index is 99', src/main.rs:4:5
note: run with `RUST_BACKTRACE=1` environment variable to display a backtrace
```

Chiar dacă este indicat motivul panicii și linia de cod unde a avut loc panica, uneori vrem să obținem mai multe informații despre stiva de apel.

3. 🌟
```shell
## UMPLI blank-ul pentru a afișa întreaga stivă de apel
## Sugestie: poți găsi indicii în informațiile implicite despre panică
$ __ cargo run
thread 'main' panicked at 'assertion failed: `(left == right)`
  left: `[97, 98, 99]`,
 right: `[96, 97, 98]`', src/main.rs:3:5
stack backtrace:
   0: rust_begin_unwind
             at /rustc/9d1b2106e23b1abd32fce1f17267604a5102f57a/library/std/src/panicking.rs:498:5
   1: core::panicking::panic_fmt
             at /rustc/9d1b2106e23b1abd32fce1f17267604a5102f57a/library/core/src/panicking.rs:116:14
   2: core::panicking::assert_failed_inner
   3: core::panicking::assert_failed
             at /rustc/9d1b2106e23b1abd32fce1f17267604a5102f57a/library/core/src/panicking.rs:154:5
   4: study_cargo::main
             at ./src/main.rs:3:5
   5: core::ops::function::FnOnce::call_once
             at /rustc/9d1b2106e23b1abd32fce1f17267604a5102f57a/library/core/src/ops/function.rs:227:5
note: Some details are omitted, run with `RUST_BACKTRACE=full` for a verbose backtrace.
```

### `unwinding` și `abort`
În mod implicit, atunci când apare o panic, programul începe să deruleze, ceea ce înseamnă că Rust se plimbă înapoi pe stivă și curăță datele din fiecare funcție întâlnită.

Dar acest proces de parcurgere înapoi și de curățare este o muncă foarte mare. Alternativa este să întrerupem imediat programul fără a face curățenie.

Dacă în proiectul tău ai nevoie să faci binarul rezultat cât mai mic posibil, poți trece de la derulare la întrerupere adăugând conținutul de mai jos în `Cargo.toml`:
```toml
[profile.release]
panic = 'abort'
```


> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
