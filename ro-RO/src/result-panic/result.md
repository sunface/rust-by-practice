# Result și ?
`Result<T>` este o enumerare folosită pentru a descrie posibilele erori. Are două variante:

- Ok(T): A fost găsită o valoare T
- Err(e): A fost găsită o eroare cu valoarea `e`
Pe scurt, rezultatul așteptat este `Ok`, în timp ce rezultatul neașteptat este `Err`.

1. 🌟🌟
```rust,editable

// UMPLE blank-urile și REZOLVĂ erorile
use std::num::ParseIntError;

fn multiply(n1_str: &str, n2_str: &str) -> __ {
    let n1 = n1_str.parse::<i32>();
    let n2 = n2_str.parse::<i32>();
    Ok(n1.unwrap() * n2.unwrap())
}

fn main() {
    let result = multiply("10", "2");
    assert_eq!(result, __);

    let result = multiply("t", "2");
    assert_eq!(result.__, 8);

    println!("Success!");
}
```

### ? 
`?` este aproape echivalent cu `unwrap`, dar `?` întoarce rezultatul în loc să provoace panică la `Err`.

2. 🌟🌟
```rust,editable

use std::num::ParseIntError;

// IMPLEMENTEAZĂ multiply cu ?
// NU folosi unwrap aici
fn multiply(n1_str: &str, n2_str: &str) -> __ {
}

fn main() {
    assert_eq!(multiply("3", "4").unwrap(), 12);
    println!("Success!");
}
```

3. 🌟🌟
```rust,editable

use std::fs::File;
use std::io::{self, Read};

fn read_file1() -> Result<String, io::Error> {
    let f = File::open("hello.txt");
    let mut f = match f {
        Ok(file) => file,
        Err(e) => return Err(e),
    };

    let mut s = String::new();
    match f.read_to_string(&mut s) {
        Ok(_) => Ok(s),
        Err(e) => Err(e),
    }
}

// UMPLE blank-urile cu o linie de cod
// NU schimba nicio linie de cod
fn read_file2() -> Result<String, io::Error> {
    let mut s = String::new();

    __;

    Ok(s)
}

fn main() {
    assert_eq!(read_file1().unwrap_err().to_string(), read_file2().unwrap_err().to_string());
    println!("Success!");
}
```

### map & and_then
[map](https://doc.rust-lang.org/stable/std/result/enum.Result.html#method.map) și [and_then](https://doc.rust-lang.org/stable/std/result/enum.Result.html#method.and_then) sunt doi combinatori comuni pentru `Result<T, E>` (și pentru `Option<T>`).

4. 🌟🌟 

```rust,editable
use std::num::ParseIntError;

// UMPLE blank-ul în două moduri: map și and_then
fn add_two(n_str: &str) -> Result<i32, ParseIntError> {
   n_str.parse::<i32>().__
}

fn main() {
    assert_eq!(add_two("4").unwrap(), 6);

    println!("Success!");
}
```

5. 🌟🌟🌟
```rust,editable
use std::num::ParseIntError;

// Cu tipul returnat rescris, folosim potrivirea modelelor fără `unwrap()`.
// Dar este atât de detaliat...
fn multiply(n1_str: &str, n2_str: &str) -> Result<i32, ParseIntError> {
    match n1_str.parse::<i32>() {
        Ok(n1)  => {
            match n2_str.parse::<i32>() {
                Ok(n2)  => {
                    Ok(n1 * n2)
                },
                Err(e) => Err(e),
            }
        },
        Err(e) => Err(e),
    }
}

// Rescrierea `multiply` pentru a o face concisă
// Ar trebui să folosești ATÂT `and_then` cât și `map` aici.
fn multiply1(n1_str: &str, n2_str: &str) -> Result<i32, ParseIntError> {
    // IMPLEMENTEAZĂ...
}

fn print(result: Result<i32, ParseIntError>) {
    match result {
        Ok(n)  => println!("n is {}", n),
        Err(e) => println!("Error: {}", e),
    }
}

fn main() {
    // Aceasta oferă încă un răspuns rezonabil.
    let twenty = multiply1("10", "2");
    print(twenty);

    // Următorul oferă acum un mesaj de eroare mult mai util.
    let tt = multiply("t", "2");
    print(tt);

    println!("Success!");
}
```

### Alias de tipuri
Folosirea `std::result::Result<T, ParseIntError>` peste tot este stufos și plictisitor, putem folosi un alias în acest scop.

La nivel de modul, crearea de aliasuri poate fi deosebit de utilă. Erorile găsite într-un modul specific au adesea același tip `Err`, astfel că un singur alias poate defini concis toate `Result`-urile asociate. Aceasta este atât de utilă încât biblioteca standard furnizează una: [`io::Result`](https://doc.rust-lang.org/std/io/type.Result.html).

6. 🌟
```rust,editable
use std::num::ParseIntError;

// UMPLE blank-ul
type __;

// Folosește aliasul de mai sus pentru a te referi la tipul nostru specific `Result`.
fn multiply(first_number_str: &str, second_number_str: &str) -> Res<i32> {
    first_number_str.parse::<i32>().and_then(|first_number| {
        second_number_str.parse::<i32>().map(|second_number| first_number * second_number)
    })
}

// Aici, aliasul permite din nou să economisim spațiu.
fn print(result: Res<i32>) {
    match result {
        Ok(n)  => println!("n is {}", n),
        Err(e) => println!("Error: {}", e),
    }
}

fn main() {
    print(multiply("10", "2"));
    print(multiply("t", "2"));

    println!("Success!");
}
```

### Utilizarea Result în `fn main`
De obicei, funcția `main` va arăta așa:
```rust
fn main() {
    println!("Hello World!");
}
```

Cu toate acestea, main poate avea și un tip de returnare `Result`. Dacă apare o eroare în cadrul funcției `main`, aceasta va returna un cod de eroare și va imprima o reprezentare de depanare a erorii (traitul Debug).

Următorul exemplu arată o astfel de situație:
```rust,editable

use std::num::ParseIntError;

fn main() -> Result<(), ParseIntError> {
    let number_str = "10";
    let number = match number_str.parse::<i32>() {
        Ok(number)  => number,
        Err(e) => return Err(e),
    };
    println!("{}", number);
    Ok(())
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
