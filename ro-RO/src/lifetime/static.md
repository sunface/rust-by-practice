# &'static și T: 'static
`'static` este un nume rezervat pentru durata de viață și l-ați putut întâlni de mai multe ori:
```rust
// O referință cu durata de viață 'static:
let s: &'static str = "hello world";

// 'static ca parte a unei constrângeri de trăsătură (trait bound):
fn generic<T>(x: T) where T: 'static {}
```

Chiar dacă toate sunt 'static, ele sunt subtil diferite.

## &'static
Ca durată de viață pentru o referință, &'static indică faptul că datele la care se referă referința trăiesc pe toată durata de execuție a programului. Cu toate acestea, poate fi încă coercită (coerced/constrânsă) la o durată de viață mai scurtă.



1. 🌟🌟 Există mai multe moduri de a crea o variabilă cu durată de viață `'static`, două dintre ele sunt stocate în memoria de tip citire a binarului.

```rust,editable

/* Completați spațiile libere în două moduri */
fn main() {
    __;
    need_static(v);

    println!("Success!")
}

fn need_static(r : &'static str) {
    assert_eq!(r, "hello");
}
```

2. 🌟🌟🌟🌟 O altă modalitate de a obține durata de viață `'static` este utilizarea `Box::leak`.
```rust,editable
#[derive(Debug)]
struct Config {
    a: String,
    b: String,
}
static mut config: Option<&mut Config> = None;

/* Faceți-l să funcționeze fără a schimba semnăturile funcțiilor din `init` */
fn init() -> Option<&'static mut Config> {
    Some(&mut Config {
        a: "A".to_string(),
        b: "B".to_string(),
    })
}


fn main() {
    unsafe {
        config = init();

        println!("{:?}",config)
    }
}
```

3. 🌟 `&'static` indică doar faptul că datele pot trăi pentru totdeauna, nu și referința. Aceasta din urmă va fi restricționată de domeniul său.
```rust,editable
fn main() {
    {
        // Creați un șir literal și tipăriți-l:
        let static_string = "I'm in read-only memory";
        println!("static_string: {}", static_string);

        // Când `static_string` iese din domeniu, referința
        // nu mai poate fi utilizată, dar datele rămân în binar.
    }

    println!("static_string reference remains alive: {}", static_string);
}
```

4. `&'static` poate fi coercitată la o durată de viață mai scurtă.

**Exemplu**
```rust,editable
// Creați o constantă cu durata de viață `'static`.
static NUM: i32 = 18;

// Returnează o referință la `NUM` în care durata de viață `'static`
// este coercitată la cea a argumentului de intrare.
fn coerce_static<'a>(_: &'a i32) -> &'a i32 {
    &NUM
}

fn main() {
    {
        // Faceți un număr întreg pentru a-l folosi în `coerce_static`:
        let lifetime_num = 9;

        // Coercitați `NUM` la durata de viață a lui `lifetime_num`:
        let coerced_static = coerce_static(&lifetime_num);

        println!("coerced_static: {}", coerced_static);
    }

    println!("NUM: {} stays accessible!", NUM);
}
```



##  T: 'static
Ca o constrângere a trăsăturii, înseamnă că tipul nu conține nicio referință non-'static. De exemplu, receptorul poate reține tipul pentru cât timp doresc și acesta nu va deveni niciodată invalid până când îl abandonează.

Este important să înțelegeți că acest lucru înseamnă că orice date deținute întotdeauna trec printr-o durată de viață 'static, dar o referință la aceste date deținute de obicei nu o face.


5. 🌟🌟
```rust,editable
/* Faceți-l să funcționeze */
use std::fmt::Debug;

fn print_it<T: Debug + 'static>( input: T) {
    println!( "'static value passed in is: {:?}", input );
}

fn print_it1( input: impl Debug + 'static ) {
    println!( "'static value passed in is: {:?}", input );
}


fn print_it2<T: Debug + 'static>( input: &T) {
    println!( "'static value passed in is: {:?}", input );
}

fn main() {
    // i este deținut și nu conține referințe, așa că este 'static:
    let i = 5;
    print_it(i);

    // ups, &i are doar durata de viață definită de domeniul
    // main(), deci nu este 'static:
    print_it(&i);

    print_it1(&i);

    // dar acesta FUNCȚIONEAZĂ!
    print_it2(&i);
}
```


6. 🌟🌟🌟
```rust,editable
use std::fmt::Display;

fn main() {
  let mut string = "First".to_owned();

  string.push_str(string.to_uppercase().as_str());
  print_a(&string);
  print_b(&string);
  print_c(&string); // Eroare la compilare
  print_d(&string); // Eroare la compilare
  print_e(&string);
  print_f(&string);
  print_g(&string); // Eroare la compilare
}

fn print_a<T: Display + 'static>(t: &T) {
  println!("{}", t);
}

fn print_b<T>(t: &T)
where
  T: Display + 'static,
{
  println!("{}", t);
}

fn print_c(t: &'static dyn Display) {
  println!("{}", t)
}

fn print_d(t: &'static impl Display) {
  println!("{}", t)
}

fn print_e(t: &(dyn Display + 'static)) {
  println!("{}", t)
}

fn print_f(t: &(impl Display + 'static)) {
  println!("{}", t)
}

fn print_g(t: &'static String) {
  println!("{}", t);
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
