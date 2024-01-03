# Debug și Display

Toate tipurile care doresc să fie afișate trebuie să implementeze trăsătura de formatare `std::fmt`: `std::fmt::Debug` sau `std::fmt::Display`.

Implementările automate sunt furnizate doar pentru tipuri precum cele din biblioteca `std`. Celelalte trebuie implementate manual.

## Debug
Implementarea `Debug` este foarte simplă: Toate tipurile pot deriva implementarea `std::fmt::Debug`. Aceasta nu este valabilă pentru `std::fmt::Display`, care trebuie implementată manual.

`{:?}` trebuie folosit pentru a afișa tipul care a implementat trăsătura `Debug`.

```rust
// Această structură nu poate fi tipărită nici cu `fmt::Display` nici cu
// `fmt::Debug`.
struct UnPrintable(i32);

// Pentru a face această structură imprimabilă cu `fmt::Debug`, putem deriva implementările automate furnizate de Rust
#[derive(Debug)]
struct DebugPrintable(i32);
```

1. 🌟
```rust,editable

/* Completați spațiile libere și remediați erorile */
struct Structure(i32);

fn main() {
    // Tipurile din std și Rust au implementată trăsătura fmt::Debug
    println!("__ months in a year.", 12);

    println!("Now __ will print!", Structure(3));
}
```

2. 🌟🌟 Deci, fmt::Debug cu siguranță face un tip imprimabil, dar sacrifică puțină eleganță. Poate putem obține ceva mai elegant înlocuind {:?} cu altceva (dar nu și {}!) 
```rust,editable
#[derive(Debug)]
struct Person {
    name: String,
    age: u8
}

fn main() {
    let person = Person { name:  "Sunface".to_string(), age: 18 };

    /* Faceți să afișeze: 
    Person {
        name: "Sunface",
        age: 18,
    }
    */
    println!("{:?}", person);
}
```

3. 🌟🌟 Putem implementa și manual trăsătura Debug pentru tipurile noastre
```rust,editable

#[derive(Debug)]
struct Structure(i32);

#[derive(Debug)]
struct Deep(Structure);


fn main() {    
    // Problema cu `derive` este că nu există control asupra modului
    // în care arată rezultatele. Ce se întâmplă dacă vreau să afișez doar un `7`?

    /* Faceți să afișeze: Acum 7 va fi tipărit! */
    println!("Now {:?} will print!", Deep(Structure(7)));
}
```

## Display
Da, Debug este simplu și ușor de folosit. Dar uneori vrem să personalizăm aspectul de ieșire al tipului nostru. Aici intervine Display.

Spre deosebire de Debug, nu există nici o modalitate de a deriva implementarea trăsăturii Display, trebuie să o implementăm manual.

Un alt lucru de remarcat: locul de substituire pentru Display este {}, nu {:?}.

4. 🌟🌟
```rust,editable

/* Faceți să funcționeze*/
use std::fmt;

struct Point2D {
    x: f64,
    y: f64,
}

impl fmt::Display for Point2D {
    /* Implementați.. */
}

impl fmt::Debug for Point2D {
    /* Implementați.. */
}

fn main() {
    let point = Point2D { x: 3.3, y: 7.2 };
    assert_eq!(format!("{}",point), "Display: 3.3 + 7.2i");
    assert_eq!(format!("{:?}",point), "Debug: Complex { real: 3.3, imag: 7.2 }");
    
    println!("Success!");
}
```


### Operatorul '?'

Implementarea 'fmt::Display' pentru o structură a cărei elemente trebuie gestionate separat este dificilă. Problema este că fiecare 'write!' generează un 'fmt::Result' care trebuie gestionat în același loc.

În mod fericit, Rust oferă operatorul '?' pentru a ne ajuta să eliminăm unele coduri inutile pentru tratarea rezultatului 'fmt::Result.'

5. 🌟🌟
```rust,editable

/* Faceți să funcționeze */
use std::fmt; 

struct List(Vec<i32>);

impl fmt::Display for List {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        // Extrageți valoarea folosind indexarea tuplului,
        // și creați o referință la `vec`.
        let vec = &self.0;

        write!(f, "[")?;

        // Iterați peste `v` în `vec` în timp ce enumerați iterația
        // numărul în `count`.
        for (count, v) in vec.iter().enumerate() {
            // Pentru fiecare element în afară de primul, adăugați o virgulă.
            // Utilizați operatorul ? pentru a reveni la erori.
            if count != 0 { write!(f, ", ")?; }
            write!(f, "{}", v)?;
        }

        // Închideți paranteza deschisă și returnați o valoare fmt::Result.
        write!(f, "]")
    }
}

fn main() {
    let v = List(vec![1, 2, 3]);
    assert_eq!(format!("{}",v), "[0: 1, 1: 2, 2: 3]");
    println!("Success!");
}
```


> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
