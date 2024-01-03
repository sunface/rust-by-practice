# From/Into
Trait-ul `From` permite unui tip să definească modul în care se poate crea singur dintr-un alt tip, oferind astfel o modalitate foarte simplă de a converti între mai multe tipuri.

Trait-urile `From` și `Into` sunt intrinsec legate, și acest lucru face parte din implementarea lor. Acest lucru înseamnă că dacă scriem ceva de genul: `impl From<T>` pentru U, atunci putem utiliza let u: `U = U::from(T)` sau `let u:U = T.into()`.

Trait-ul Into este pur și simplu reciprocul trait-ului `From`. Adică, dacă ați implementat trait-ul `From` pentru tipul vostru, atunci trait-ul `Into` va fi implementat automat pentru același tip.

Utilizarea trait-ului `Into` va necesita în mod tipic adnotările de tip, deoarece compilatorul nu poate determina acest lucru de cele mai multe ori.

De exemplu, putem converti ușor `&str` în `String`:
```rust
fn main() {
    let my_str = "hello";

    // cele trei conversii de mai jos depind toate de faptul că String implementează From<&str>:
    let string1 = String::from(my_str);
    let string2 = my_str.to_string();
    // Este necesară adnotarea explicită de tip aici
    let string3: String = my_str.into();
}
```

Deoarece biblioteca standard a implementat deja acest lucru pentru noi: `impl From<&'_ str>` pentru String.

Unele implementări ale trait-ului `From` pot fi găsite [aici](https://doc.rust-lang.org/stable/std/convert/trait.From.html#implementors).

1. 🌟🌟🌟
```rust,editable
fn main() {
    // impl From<bool> pentru i32
    let i1: i32 = false.into();
    let i2: i32 = i32::from(false);
    assert_eq!(i1, i2);
    assert_eq!(i1, 0);

    // CORECTAȚI eroarea în două moduri
    /* 1. utilizați un tip similar care `impl From<char>`, poate ar trebui să verificați documentația menționată mai sus pentru a găsi răspunsul */
    // 2. un cuvânt cheie din ultimul capitol
    let i3: i32 = 'a'.into();

    // CORECTAȚI eroarea în două moduri
    let s: String = 'a' as String;

    println!("Success!");
}
```

### Implementare `From` pentru tipuri personalizate
2. 🌟🌟
```rust,editable
// `From` este acum inclus în `std::prelude`, deci nu este nevoie să-l introducem în domeniul de vizibilitate actual
// use std::convert::From;

#[derive(Debug)]
struct Number {
    value: i32,
}

impl From<i32> for Number {
    // IMPLEMENTAȚI metoda `from`
}

// COMPLETAȚI spațiile goale
fn main() {
    let num = __(30);
    assert_eq!(num.value, 30);

    let num: Number = __;
    assert_eq!(num.value, 30);

    println!("Success!");
}
```

3. 🌟🌟🌟 Când efectuăm gestionarea erorilor, este adesea util să implementăm trait-ul From pentru propria noastră eroare. Atunci putem utiliza ? pentru a converti automat tipul de eroare subiacent la propria noastră eroare.
```rust,editable
use std::fs;
use std::io;
use std::num;

enum CliError {
    IoError(io::Error),
    ParseError(num::ParseIntError),
}

impl From<io::Error> for CliError {
    // IMPLEMENTAȚI metoda `from`
}

impl From<num::ParseIntError> for CliError {
    // IMPLEMENTAȚI metoda `from`
}

fn open_and_parse_file(file_name: &str) -> Result<i32, CliError> {
    // ? convertește automat io::Error la CliError
    let contents = fs::read_to_string(&file_name)?;
    // num::ParseIntError -> CliError
    let num: i32 = contents.trim().parse()?;
    Ok(num)
}

fn main() {
    println!("Success!");
}
```


### TryFrom/TryInto
Similar cu `From` și `Into`, `TryFrom` și `TryInto` sunt trait-uri generice pentru conversii între tipuri.

Spre deosebire de `From/Into`, `TryFrom` și `TryInto` sunt folosite pentru conversiile care pot eșua și returnează un Result în loc de o valoare simplă.

4. 🌟🌟
```rust,editable
// TryFrom și TryInto sunt incluse în `std::prelude`, deci nu este nevoie să le introducem în domeniul de vizibilitate actual
// use std::convert::TryInto;

fn main() {
    let n: i16 = 256;

    // Trait-ul Into are o metodă `into`,
    // prin urmare TryInto are o metodă ?
    let n: u8 = match n.__() {
        Ok(n) => n,
        Err(e) => {
            println!("there is an error when converting: {:?}, but we catch it", e.to_string());
            0
        }
    };

    assert_eq!(n, __);

    println!("Success!");
}
```

5. 🌟🌟🌟
```rust,editable
#[derive(Debug, PartialEq)]
struct EvenNum(i32);

impl TryFrom<i32> for EvenNum {
    type Error = ();

    // IMPLEMENTAȚI `try_from`
    fn try_from(value: i32) -> Result<Self, Self::Error> {
        if value % 2 == 0 {
            Ok(EvenNum(value))
        } else {
            Err(())
        }
    }
}

fn main() {
    assert_eq!(EvenNum::try_from(8), Ok(EvenNum(8)));
    assert_eq!(EvenNum::try_from(5), Err(()));

    // COMPLETAȚI spațiile goale
    let result: Result<EvenNum, ()> = 8i32.try_into();
    assert_eq!(result, __);
    let result: Result<EvenNum, ()> = 5i32.try_into();
    assert_eq!(result, __);

    println!("Success!");
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
