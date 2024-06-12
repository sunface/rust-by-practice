# Functions
1. 🌟🌟🌟
```rust,editable

fn main() {
    // Nu modificați următoarele două linii!
    let (x, y) = (1, 2);
    let s = sum(x, y);

    assert_eq!(s, 3);

    println!("Success!");
}

fn sum(x, y: i32) {
    x + y;
}
```


2. 🌟
```rust,editable
fn main() {
   print();
}

// Inlocuiți i32 cu un alt tip de date
fn print() -> i32 {
   println!("Success!");
}
```


3. 🌟🌟🌟

```rust,editable
// Rezolvați in două moduri
// NU lăsați `println!` să ruleze
fn main() {
    never_return();

    println!("Failed!");
}

fn never_return() -> ! {
    // Implementați această funcție, nu modificați semnăturile funcției
    
}
```

### Diverging functions 
Funcțiile divergente nu se întorc niciodată la apelant, astfel că pot fi folosite în locuri unde se așteaptă o valoare de orice tip.

4. 🌟🌟
```rust,editable

fn main() {
    println!("Success!");
}

fn get_option(tp: u8) -> Option<i32> {
    match tp {
        1 => {
            // TODO
        }
        _ => {
            // TODO
        }
    };
    
    // În loc să returnăm un "None", folosim în schimb o funcție divergentă.
    never_return_fn()
}

// IMPLMENTAȚI această funcție în TREI moduri
fn never_return_fn() -> ! {
    
}
```

5. 🌟🌟
```rust,editable

fn main() {
    // COMPLETEAZĂ spațiul liber
    let b = __;

    let _v = match b {
        true => 1,
        // Funcțiile divergente pot fi, de asemenea, utilizate în expresii match pentru a înlocui o valoare de orice tip.
        false => {
            println!("Success!");
            panic!("we have no value for `false`, but we can panic");
        }
    };

    println!("Exercise Failed if printing out this line!");
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice)(sub calea soluțiilor), dar utilizați-le doar atunci când aveți nevoie.
