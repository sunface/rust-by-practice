# Folosire și publicare
1. 🌟 Putem aduce două tipuri cu același nume în același domeniu cu 'use', dar avem nevoie de cuvântul cheie 'as'.

```rust,editable
use std::fmt::Result;
use std::io::Result;

fn main() {}
```

2. 🌟🌟 Dacă folosim mai multe elemente definite în aceeași creație sau modul, atunci listarea fiecărui element pe linie proprie va ocupa prea mult spațiu vertical.

```rust,editable

// COMPLETEAZĂ spațiul liber în două moduri
// NU adăuga linii noi de cod
use std::collections::__;

fn main() {
    let _c1:HashMap<&str, i32> = HashMap::new();
    let mut c2 = BTreeMap::new();
    c2.insert(1, "a");
    let _c3: HashSet<i32> = HashSet::new();
}
```

### Re-exportarea numelor cu 'pub use'
3. 🌟🌟🌟 În pachetul nostru creat recent, hello-package, adaugă ceva pentru a face codul de mai jos să funcționeze.
```rust,editable
fn main() {
    assert_eq!(hello_package::hosting::seat_at_table(), "sit down please");
     assert_eq!(hello_package::eat_at_restaurant(),"yummy yummy!");
}
```


### Pub(in Crate) 
Uneori dorim ca un element să fie public doar pentru o anumită creație. Pentru asta putem folosi sintaxa pub(in Crate).

#### Exemplu
```rust,editable
pub mod a {
    pub const I: i32 = 3;

    fn semisecret(x: i32) -> i32 {
        use self::b::c::J;
        x + J
    }

    pub fn bar(z: i32) -> i32 {
        semisecret(I) * z
    }
    pub fn foo(y: i32) -> i32 {
        semisecret(I) + y
    }

    mod b {
        pub(in crate::a) mod c {
            pub(in crate::a) const J: i32 = 4;
        }
    }
}
```

### Cod Complet
Codul complet al hello-package se află [aici](https://github.com/sunface/rust-by-practice/tree/master/practices/hello-package).


> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)