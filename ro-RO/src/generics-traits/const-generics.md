# Generice Constante
Genericele constante reprezintă argumente generice care variază asupra valorilor constante, în loc de tipuri sau durate de viață. Acest lucru permite, de exemplu, tipurilor să fie parametrizate de numere întregi. De fapt, există un exemplu de tipuri generice constante încă de la începutul dezvoltării limbajului Rust: tipurile de tablou [T; N], pentru un anumit tip T și N: usize. Cu toate acestea, nu a existat anterior o modalitate de a abstrage asupra tablourilor de dimensiune arbitrară: dacă doreați să implementați o trăsătură pentru tablouri de orice dimensiune, trebuia să faceți acest lucru manual pentru fiecare valoare posibilă. Timp îndelungat, chiar și metodele standard ale bibliotecii pentru tablouri au fost limitate la tablouri de lungime cel mult 32 din cauza acestei probleme.

## Exemple
1. Iată un exemplu de tip și implementare care folosește generice constante: un tip care înconjoară o pereche de tablouri de aceeași dimensiune.
```rust,editable
struct ArrayPair<T, const N: usize> {
    left: [T; N],
    right: [T; N],
}

impl<T: Debug, const N: usize> Debug for ArrayPair<T, N> {
    // ...
}
```


2. În prezent, parametrii const pot fi instanțiați numai de argumente const ale formelor următoare:

- Un parametru const separat.
- O literă (adică un întreg, bool sau caracter).
- O expresie constantă concretă (închisă între {}), care nu implică niciun parametru generic.
  
```rust,editable
fn foo<const N: usize>() {}

fn bar<T, const M: usize>() {
    foo::<M>(); // Corect: `M` este un parametru const
    foo::<2021>(); // Corect: `2021` este o literă
    foo::<{20 * 100 + 20 * 10 + 1}>(); // Corect: expresia const conține zero parametri generici
    
    foo::<{ M + 1 }>(); // Eroare: expresia const conține parametrul generic `M`
    foo::<{ std::mem::size_of::<T>() }>(); // Eroare: expresia const conține parametrul generic `T`
    
    
    let _: [u8; M]; // Corect: `M` este un parametru const
    let _: [u8; std::mem::size_of::<T>()]; // Eroare: expresia const conține parametrul generic `T`
}

fn main() {}
```

3. Genericele constante ne pot permite, de asemenea, să evităm unele verificări la timpul de execuție.
```rust
/// O regiune de memorie conținând cel puțin `N` elemente de tip `T`.
pub struct MinSlice<T, const N: usize> {
    /// Regiunea mărginită de memorie. Exact `N` elemente de tip `T`.
    pub head: [T; N],
    /// Zero sau mai multe elemente `T` rămase după cele `N` din regiunea mărginită.
    pub tail: [T],
}

fn main() {
    let slice: &[u8] = b"Hello, world";
    let reference: Option<&u8> = slice.get(6);
    // Știm că această valoare este `Some(b' ')`,
    // dar compilatorul nu poate ști asta.
    assert!(reference.is_some());

    let slice: &[u8] = b"Hello, world";
    // Verificarea lungimii este efectuată când construim un MinSlice,
    // și se știe la timpul de compilare că are lungimea 12.
    // Dacă `unwrap()` reușește, nu mai sunt necesare verificări
    // pe durata vieții `MinSlice`.
    let minslice = MinSlice::<u8, 12>::from_slice(slice).unwrap();
    let value: u8 = minslice.head[6];
    assert_eq!(value, b' ')
}
```


## Exerciții
1. 🌟🌟 '<T, const N: usize>' face parte din tipul structurii, ceea ce înseamnă că 'Array<i32, 3>' și 'Array<i32, 4>' sunt tipuri diferite.
   
```rust,editable
struct Array<T, const N: usize> {
    data : [T; N]
}

fn main() {
    let arrays = [
        Array{
            data: [1, 2, 3],
        },
        Array {
            data: [1.0, 2.0, 3.0],
        },
        Array {
            data: [1, 2]
        }
    ];

    println!("Success!");
}
```

2. 🌟🌟 
```rust,editable

// Completați spațiile libere pentru a face codul funcțional.
fn print_array<__>(__) {
    println!("{:?}", arr);
}
fn main() {
    let arr = [1, 2, 3];
    print_array(arr);

    let arr = ["hello", "world"];
    print_array(arr);
}
```

3. 🌟🌟🌟 Uneori vrem să limităm dimensiunea unei variabile, de exemplu, când o folosim în medii integrate, apoi 'const expressions' vor fi soluția potrivită.
   
```rust,editable
#![allow(incomplete_features)]
#![feature(generic_const_exprs)]

fn check_size<T>(val: T)
where
    Assert<{ core::mem::size_of::<T>() < 768 }>: IsTrue,
{
    //...
}

// Rezolvați erorile din funcția principală.
fn main() {
    check_size([0u8; 767]); 
    check_size([0i32; 191]);
    check_size(["hello你好"; __]); // Dimensiunea &str?
    check_size([(); __].map(|_| "hello你好".to_string()));  // Dimensiunea String?
    check_size(['中'; __]); // Dimensiunea char ?

    println!("Success!");
}



pub enum Assert<const CHECK: bool> {}

pub trait IsTrue {}

impl IsTrue for Assert<true> {}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)