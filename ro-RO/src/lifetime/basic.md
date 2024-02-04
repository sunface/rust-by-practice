## Durată de viață
Compilatorul utilizează durata de viață pentru a asigura că toate împrumuturile sunt valide. În mod obișnuit, durata de viață a unei variabile începe când este creată și se încheie când este distrusă.

## Domeniul de viață
1. 🌟
```rust,editable
/* Adăugați adnotările pentru duratele de viață ale `i` și `borrow2` */

// Duratele de viață sunt adnotate mai jos cu linii care indică crearea
// și distrugerea fiecărei variabile.
// `i` are cea mai lungă durată de viață pentru că domeniul său încadrează
// complet atât `borrow1`, cât și `borrow2`. Durata lui `borrow1` comparativ
// cu cea a lui `borrow2` este irelevantă, deoarece acestea sunt distincte.
fn main() {
    let i = 3;                                             
    {                                                    
        let borrow1 = &i; // // Durata de viață a lui `borrow1` începe. ──┐
        //                                                                │
        println!("borrow1: {}", borrow1); //                              │
    } // `borrow1` se încheie ────────────────────────────────────────────┘
    {                                                    
        let borrow2 = &i; 
                                                        
        println!("borrow2: {}", borrow2);               
    }                                                   
}   
```

2. 🌟🌟

**Exemplu**
```rust
{
    let x = 5;            // ----------+-- 'b
                          //           |
    let r = &x;           // --+-- 'a  |
                          //   |       |
    println!("r: {}", r); //   |       |
                          // --+       |
}                         // ----------+
```


```rust,editable
/* Adăugați adnotările pentru `r` și `x` ca mai sus și explicați de ce acest cod nu compilează din perspectiva duratei de viață. */

fn main() {  
    {
        let r;                // ---------+-- 'a
                              //          |
        {                     //          |
            let x = 5;        // -+-- 'b  |
            r = &x;           //  |       |
        }                     // -+       |
                              //          |
        println!("r: {}", r); //          |
    }                         // ---------+
}
```

## Adnotarea Duratei de Viață
**Verificatorul de împrumuturi utilizează adnotările explicite ale duratei de viață** pentru a determina cât timp trebuie să fie valabilă o referință.

Dar pentru noi, utilizatorii, în majoritatea cazurilor, nu este nevoie să adnotăm durata de viață, deoarece există mai multe reguli de omisiune, iar înainte de a învăța aceste reguli, trebuie să știm cum să adnotăm manual durata de viață.

#### Funcții
Ignorând regulile de omisiune, duratele de viață în semnăturile funcțiilor au câteva constrângeri:

- Orice referință trebuie să aibă o durată de viață adnotată.
- Oricare referință care este returnată trebuie să aibă aceeași durată de viață ca una dintre intrările sau să fie statică.

**Exemplu**
```rust,editable
// O referință de intrare cu durata de viață `'a`, care trebuie să trăiască
// cel puțin pe durata funcției.
fn print_one<'a>(x: &'a i32) {
    println!("`print_one`: x is {}", x);
}

// Referințe mutable sunt posibile și cu durate de viață.
fn add_one<'a>(x: &'a mut i32) {
    *x += 1;
}

// Mai multe elemente cu durate de viață diferite. În acest caz,
// ar fi în regulă ca ambele să aibă aceeași durată de viață `'a`, dar
// în cazuri mai complexe, pot fi necesare durate de viață diferite.
fn print_multi<'a, 'b>(x: &'a i32, y: &'b i32) {
    println!("`print_multi`: x is {}, y is {}", x, y);
}

// Returnarea referințelor care au fost trecute ca intrări este acceptabilă.
// Cu toate acestea, trebuie returnată durata de viață corectă.
fn pass_x<'a, 'b>(x: &'a i32, _: &'b i32) -> &'a i32 { x }

fn main() {
    let x = 7;
    let y = 9;
    
    print_one(&x);
    print_multi(&x, &y);
    
    let z = pass_x(&x, &y);
    print_one(z);

    let mut t = 3;
    add_one(&mut t);
    print_one(&t);
}
```

3. 🌟
```rust,editable
/* Faceți-l să funcționeze adăugând adnotarea adecvată pentru durata de viață */
fn longest(x: &str, y: &str) -> &str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn main() {}
```
4. 🌟🌟🌟
```rust,editable
// `'a` trebuie să trăiască mai mult decât funcția.
// Aici, `&String::from("foo")` ar crea un `String`, urmat de un
// referință. Apoi datele sunt abandonate la ieșirea din domeniu, lăsând
// o referință către date invalide pentru a fi returnată.

/* Reparați eroarea în trei moduri */
fn invalid_output<'a>() -> &'a String { 
    &String::from("foo") 
}

fn main() {
}
```

5. 🌟🌟
```rust,editable
// `print_refs` primește două referințe la `i32` care au durate de viață diferite
// `'a` și `'b`. Aceste două durate de viață trebuie să fie ambele cel puțin
// la fel de lungi ca funcția `print_refs`.
fn print_refs<'a, 'b>(x: &'a i32, y: &'b i32) {
    println!("x is {} and y is {}", x, y);
}

/* Faceți-l să funcționeze */
// O funcție care nu primește argumente, dar are un parametru de durată de viață `'a`.
fn failed_borrow<'a>() {
    let _x = 12;

    // EROARE: `_x` nu trăiește destul de mult
    let y: &'a i32 = &_x;
    // Încercarea de a utiliza durata de viață `'a` ca adnotare de tip explicită 
    // în interiorul funcției va eșua pentru că durata de viață a `&_x` este mai scurtă
    // decât `'a`. O durată de viață scurtă nu poate fi transformată într-una mai lungă.
}

fn main() {
    let (four, nine) = (4, 9);
    
    // Împrumuturile (`&`) ale ambelor variabile sunt trecute în funcție.
    print_refs(&four, &nine);
    // Orice intrare împrumutată trebuie să supraviețuiască împrumutătorului. 
    // Cu alte cuvinte, durata de viață a lui `four` și `nine` trebuie să 
    // fie mai lungă decât cea a funcției `print_refs`.
    
    failed_borrow();
    // `failed_borrow` nu conține referințe pentru a forța `'a` să fie
    // mai lung decât durata de viață a funcției, dar `'a` este mai lung.
    // Deoarece durata de viață nu este niciodată restricționată, aceasta
    // devine implicită `'static`.
}
```

#### Structs
6. 🌟
```rust,editable
/* Faceți-l să funcționeze adăugând adnotarea adecvată pentru durata de viață */

// Un tip `Borrowed` care conține o referință la un
// `i32`. Referința la `i32` trebuie să trăiască mai mult decât `Borrowed`.
#[derive(Debug)]
struct Borrowed(&i32);

// Similar, ambele referințe aici trebuie să trăiască mai mult decât această structură.
#[derive(Debug)]
struct NamedBorrowed {
    x: &i32,
    y: &i32,
}

// Un enum care este fie un `i32`, fie o referință la unul.
#[derive(Debug)]
enum Either {
    Num(i32),
    Ref(&i32),
}

fn main() {
    let x = 18;
    let y = 15;

    let single = Borrowed(&x);
    let double = NamedBorrowed { x: &x, y: &y };
    let reference = Either::Ref(&x);
    let number    = Either::Num(y);

    println!("x is borrowed in {:?}", single);
    println!("x and y are borrowed in {:?}", double);
    println!("x is borrowed in {:?}", reference);
    println!("y is *not* borrowed in {:?}", number);
}
```


7. 🌟🌟
```rust,editable
/* Faceți-l să funcționeze */

#[derive(Debug)]
struct NoCopyType {}

#[derive(Debug)]
struct Example<'a, 'b> {
    a: &'a u32,
    b: &'b NoCopyType
}

fn main()
{ 
  /* 'a legat de stiva fn-main */
  let var_a = 35;
  let example: Example;
  
  {
    /* Durata de viață 'b legată de un nou cadru/scope */ 
    let var_b = NoCopyType {};
    
    /* fixme */
    example = Example { a: &var_a, b: &var_b };
  }
  
  println!("(Success!) {:?}", example);
}
```


8. 🌟🌟
```rust,editable

#[derive(Debug)]
struct NoCopyType {}

#[derive(Debug)]
#[allow(dead_code)]
struct Example<'a, 'b> {
    a: &'a u32,
    b: &'b NoCopyType
}

/* Fixați semnătura funcției */
fn fix_me(foo: &Example) -> &NoCopyType
{ foo.b }

fn main()
{
    let no_copy = NoCopyType {};
    let example = Example { a: &1, b: &no_copy };
    fix_me(&example);
    println!("Success!")
}
```

## Metodă
Metodele sunt adnotate similar cu funcțiile.

**Exemplu**
```rust,editable
struct Owner(i32);

impl Owner {
    // Adnotare pentru duratele de viață la fel ca într-o funcție independentă.
    fn add_one<'a>(&'a mut self) { self.0 += 1; }
    fn print<'a>(&'a self) {
        println!("`print`: {}", self.0);
    }
}

fn main() {
    let mut owner = Owner(18);

    owner.add_one();
    owner.print();
}
```

9. 🌟🌟
```rust,editable
/* Faceți-l să funcționeze adăugând adnotările adecvate pentru durata de viață */
struct ImportantExcerpt {
    part: &str,
}

impl ImportantExcerpt {
    fn level(&'a self) -> i32 {
        3
    }
}

fn main() {}
```

## Omisiune (Elision)
Unele modele de durate de viață sunt atât de comune încât verificatorul de împrumuturi va permite să le omiteți pentru a salva tastarea și a îmbunătăți citirea.

Acest lucru este cunoscut sub numele de **omisiune**. Omisiunea există în Rust doar pentru că aceste modele sunt comune.

Pentru o înțelegere mai cuprinzătoare a omisiunii, vă rugăm să consultați [omisiunea duratei de viață](https://doc.rust-lang.org/book/ch10-03-lifetime-syntax.html#lifetime-elision) în cartea oficială.

10. 🌟🌟
```rust,editable
/* Eliminați toate duratele de viață care pot fi omise */

fn input<'a>(x: &'a i32) {
    println!("`annotated_input`: {}", x);
}

fn pass<'a>(x: &'a i32) -> &'a i32 { x }

fn longest<'a, 'b>(x: &'a str, y: &'b str) -> &'a str {
    x
}

struct Owner(i32);

impl Owner {
    // Adnotați duratele de viață la fel ca într-o funcție independentă.
    fn add_one<'a>(&'a mut self) { self.0 += 1; }
    fn print<'a>(&'a self) {
        println!("`print`: {}", self.0);
    }
}

struct Person<'a> {
    age: u8,
    name: &'a str,
}

enum Either<'a> {
    Num(i32),
    Ref(&'a i32),
}

fn main() {}
```
> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
