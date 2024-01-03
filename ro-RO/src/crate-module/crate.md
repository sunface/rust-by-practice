# Pachet și ladă (crate)
Un pachet este un proiect pe care il creăm cu "cargo" (in general). Pachetul conține un fișier 'Cargo.toml'.

1. 🌟 Creați un pachet cu ierarhia de mai jos
```shell
.
├── Cargo.toml
└── src
    └── main.rs

1 directory, 2 files
```

```toml
# in Cargo.toml
[package]
name = "hello-package"
version = "0.1.0"
edition = "2021"
```

> În acest pachet vom lucra pe parcursul întregului capitol!

2. 🌟 Creați un pachet cu ierarhia de mai jos
```shell
.
├── Cargo.toml
└── src
    └── lib.rs

1 directory, 2 files
```

```toml
# in Cargo.toml
[package]
name = "hello-package1"
version = "0.1.0"
edition = "2021"
```

> Acest pachet poate fi șters datorită existenței primului.

3. 🌟 
```rust,editable
/* Întrebare: */

// Î: Care este diferența dintre cele două pachete?
// R: __
```


## Crate
O ladă este un fișier binar sau o bibliotecă. Rădăcina lăzii este un fișier sursă de unde începe compilatorul Rust și formează modulul rădăcină al lăzii.

În pachetul hello-package, există o ladă binară cu același nume ca și pachetul: hello-package, iar src/main.rs este rădăcina lăzii acestei lăzi binare.

Similar cu 'hello-package', 'hello-package1' are, de asemenea, o ladă în el, însă acest pachet nu conține o ladă binară, ci o ladă de bibliotecă, iar src/lib.rs este rădăcina lăzii.

4. 🌟
```rust,editable
/* Întrebare: */

// Î: Care este numele lăzii de bibliotecă în pachetul hello-package1?
// R: __
```


5. 🌟🌟 Adaugă o cutie de bibliotecă pentru 'hello-package' și descrie structura sa de fișiere mai jos:
```shell,editable
# COMPLETEAZĂ spațiile libere
.
├── Cargo.lock
├── Cargo.toml
├── src
│   ├── __
│   └── __
```

După acest pas, în pachetul hello-package ar trebui să existe două cutii: o cutie binară și o cutie de bibliotecă, ambele având același nume ca și pachetul.

6. 🌟🌟🌟 Un pachet poate conține cel mult o cutie de bibliotecă, dar poate conține oricâte cutii binare dorești plasând fișiere în directorul src/bin: fiecare fișier va fi o cutie binară separată cu același nume ca și fișierul.

```shell,editable
# Creează un pachet care conține
# 1. trei cutii binare: hello-package, main1 și main2
# 2. o cutie de bibliotecă
# descrie structura directorului mai jos
.
├── Cargo.toml
├── Cargo.lock
├── src
│   ├── __
│   ├── __
│   └── __
│       └── __
│       └── __
├── tests # director pentru fișierele de teste integrate
│   └── some_integration_tests.rs
├── benches # director pentru fișierele de referință de performanță
│   └── simple_bench.rs
└── examples # director pentru fișierele de exemplu
    └── simple_example.rs
```

Așa cum poți observa, structura pachetului de mai sus este foarte standard și este folosită pe scară largă în multe proiecte Rust.


> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)