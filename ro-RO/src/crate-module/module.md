# Module
Modulele ne permit să organizăm codul într-o cutie în grupuri pentru lizibilitate și ușurință în reutilizare. Modulele controlează, de asemenea, confidențialitatea elementelor, adică dacă un element poate fi văzut de codul extern (public) sau este doar o implementare internă și nu este disponibil pentru codul extern (privat).


Am creat un pachet numit hello-package în capitolul anterior și arată așa:
```shell
.
├── Cargo.toml
├── src
│   ├── lib.rs
│   └── main.rs
```

Acum este momentul să creăm câteva module în cutia de bibliotecă și să le utilizăm în cutia binară, să începem.

1. 🌟🌟 Implementează modulul front_of_house bazat pe arborele de module de mai jos:
```shell
library crate root
 └── front_of_house
     ├── hosting
     │   ├── add_to_waitlist
     │   └── seat_at_table
     └── serving
         ├── take_order
         ├── serve_order
         ├── take_payment
         └── complain
```

```rust,editable
// COMPLETEAZĂ spațiile libere
// în __.rs

mod front_of_house {
    // IMPLEMENTEAZĂ acest modul...
}
```


2. 🌟🌟 Să apelăm add_to_waitlist dintr-o funcție eat_at_restaurant care se află în rădăcina cutiei de bibliotecă.

```rust,editable
// În lib.rs

// COMPLETEAZĂ spațiile libere și CORECTEAZĂ erorile
// Trebuie să faci ceva public cu pub pentru a asigura accesibilitatea pentru codul extern fn eat_at_restaurant()
mod front_of_house {
    /* ...snip... */
}

pub fn eat_at_restaurant() {
    // Apelează 'add_to_waitlist' cu **cale absolută**:
    __.add_to_waitlist();

    // Apelează cu **cale relativă**
     __.add_to_waitlist();
}
```

3. 🌟🌟 Poți folosi 'super' pentru a importa elemente în cadrul modulului părinte
```rust,editable
// În lib.rs

mod back_of_house {
    fn fix_incorrect_order() {
        cook_order();
        // COMPLETEAZĂ spațiul liber în trei moduri
        // 1. folosind cuvântul cheie 'super'
        // 2. folosind calea absolută
        __.serve_order();
    }

    fn cook_order() {}
}
```


### Separarea modulelor în fișiere diferite
```rust,editable
// În lib.rs
pub mod front_of_house {
    pub mod hosting {
        pub fn add_to_waitlist() {}

        pub fn seat_at_table() -> String {
            String::from("sit down please")
        }
    }

    pub mod serving {
        pub fn take_order() {}

        pub fn serve_order() {}

        pub fn take_payment() {}

        // Poate nu vrei ca oaspeții să audă plângerile tale despre ei
        // Așa că fă-le private
        fn complain() {} 
    }
}

pub fn eat_at_restaurant() -> String {
    front_of_house::hosting::add_to_waitlist();
    
    back_of_house::cook_order();

    String::from("yummy yummy!")
}

pub mod back_of_house {
    pub fn fix_incorrect_order() {
        cook_order();
        crate::front_of_house::serving::serve_order();
    }

    pub fn cook_order() {}
}
```

4. 🌟🌟🌟🌟 Te rog să separi modulele și codurile de mai sus în fișiere aflate în arborele de directoare de mai jos:
```shell
.
├── Cargo.toml
├── src
│   ├── back_of_house.rs
│   ├── front_of_house
│   │   ├── hosting.rs
│   │   ├── mod.rs
│   │   └── serving.rs
│   ├── lib.rs
│   └── main.rs
```

```rust,editable
// În src/lib.rs

// IMPLEMENTEAZĂ...
```

```rust,editable
// În src/back_of_house.rs

// IMPLEMENTEAZĂ...
```


```rust,editable
// În src/front_of_house/mod.rs

// IMPLEMENTEAZĂ...
```

```rust,editable
// În src/front_of_house/hosting.rs

// IMPLEMENTEAZĂ...
```

```rust,editable
// În src/front_of_house/serving.rs

// IMPLEMENTEAZĂ...
```

### Accesarea codului în cutia de bibliotecă din cutia binară
**Te rog să te asiguri că ai finalizat a patra exercițiu înainte de a face progrese suplimentare.**

Ar trebui să ai structurile de mai jos și codurile corespunzătoare în ele când ajungi aici:
```shell
.
├── Cargo.toml
├── src
│   ├── back_of_house.rs
│   ├── front_of_house
│   │   ├── hosting.rs
│   │   ├── mod.rs
│   │   └── serving.rs
│   ├── lib.rs
│   └── main.rs
```

5. 🌟🌟🌟 Acum vom apela câteva funcții de bibliotecă din cutia binară.

```rust,editable
// În src/main.rs

// COMPLETEAZĂ spațiul liber și CORECTEAZĂ erorile
fn main() {
    assert_eq!(__, "sit down please");
    assert_eq!(__,"yummy yummy!");
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
