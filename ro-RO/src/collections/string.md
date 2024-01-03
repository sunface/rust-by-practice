# String
`std::string::String` este un șir extensibil codificat în UTF-8. Este cel mai comun tip de șir folosit în dezvoltarea zilnică și are proprietatea conținutului șirului.

### Operații de bază
1. 🌟🌟
```rust,editable

// COMPLETAȚI spațiile goale și CORECTAȚI erorile
// 1. Nu utilizați to_string()
// 2. Nu adăugați/eliminați nicio linie de cod
fn main() {
    let mut s: String = "hello, ";
    s.push_str("world".to_string());
    s.push(__);

    move_ownership(s);

    assert_eq!(s, "hello, world!");

    println!("Success!");
}

fn move_ownership(s: String) {
    println!("ownership of \"{}\" is moved here!", s)
}
```

### String și &str
Un String este stocat sub formă de vector de octeți (Vec<u8>), dar este întotdeauna garantat să fie o secvență UTF-8 validă. String este alocat pe heap, poate crește dinamic și nu este terminat cu null.

&str este o felie (&[u8]) care indică întotdeauna către o secvență UTF-8 validă și poate fi utilizat pentru a vizualiza conținutul unui String, la fel cum &[T] este o vedere într-un Vec<T>.

2. 🌟🌟
```rust,editable
// COMPLETAȚI spațiile goale
fn main() {  
   let mut s = String::from("hello, world");

   let slice1: &str = __; // În două moduri
   assert_eq!(slice1, "hello, world");

   let slice2 = __;
   assert_eq!(slice2, "hello");

   let slice3: __ = __; 
   slice3.push('!');
   assert_eq!(slice3, "hello, world!");

   println!("Success!");
}
```

3. 🌟🌟
```rust,editable

// Întrebare: câte alocări de heap au loc aici?
// Răspunsul tău:
fn main() {  
    // Creează un tip String bazat pe &str
    // Tipul literalurilor de șir este &str
   let s: String = String::from("hello, world!");

   // Creează un `slice` care indică către String-ul s
   let slice: &str = &s;

   
    // Creează un tip String bazat pe `slice`-ul creat recent
   let s: String = slice.to_string();

   assert_eq!(s, "hello, world!");

   println!("Success!");
}
```

### UTF-8 și Indexare
Șirurile sunt întotdeauna formate din UTF-8 valid. Acest lucru are câteva implicații:

Prima dintre acestea este că, dacă aveți nevoie de un șir non-UTF-8, luați în considerare [OsString](https://doc.rust-lang.org/stable/std/ffi/struct.OsString.html). Este similar, dar fără constrângerea UTF-8.
A doua implicație este că nu puteți indexa într-un String.
Indexarea este menită să fie o operație cu timp constant, dar encodarea UTF-8 nu ne permite să facem acest lucru. În plus, nu este clar ce fel de informație ar trebui să returneze indexul: un octet, un punct de cod sau un cluster grafemic. Metodele bytes și chars returnează iteratori peste primele două, respectiv.

4. 🌟🌟🌟 Nu puteți folosi indexul pentru a accesa un caracter într-un șir, dar puteți utiliza felișea &s1[start..end].

```rust,editable

// COMPLETAȚI spațiile goale și CORECTAȚI erorile
fn main() {
    let s = String::from("hello, 世界");
    let slice1 = s[0]; // Sfaturi: h ocupă doar 1 byte în formatul UTF-8
    assert_eq!(slice1, "h");

    let slice2 = &s[3..5]; // Sfaturi: 中 ocupă 3 octeți în formatul UTF-8
    assert_eq!(slice2, "世");
    
    // Parcurgeți toate caracterele din s
    for (i, c) in s.__ {
        if i == 7 {
            assert_eq!(c, '世')
        }
    }

    println!("Success!");
}
```


#### UTF8_slice
Puteți utiliza [utf8_slice](https://docs.rs/utf8_slice/1.0.0/utf8_slice/fn.slice.html) pentru a felia un șir UTF8, acesta poate indexa caractere în loc de octeți.

**Exemplu**
```rust
use utf8_slice;
fn main() {
   let s = "The 🚀 goes to the 🌑!";

   let rocket = utf8_slice::slice(s, 4, 5);
   // Va fi egal cu "🚀"
}
```


5. 🌟🌟🌟
> Sfat: poate aveți nevoie de metoda from_utf8

```rust,editable

// COMPLETAȚI spațiile goale
fn main() {
    let mut s = String::new();
    __;

    // Niste octeți într-un vector.
    let v = vec![104, 101, 108, 108, 111];

    // Transformă un vector de octeți într-un șir de caractere (String)
    let s1 = __;
    
    
    assert_eq!(s, s1);

    println!("Success!");
}
```

### Reprezentare
Un șir de caractere (String) este compus din trei componente: un pointer către niște octeți, o lungime și o capacitate. 

Pointer-ul arată către un buffer intern pe care String îl folosește pentru a stoca datele sale. Lungimea reprezintă numărul de octeți stocați în prezent în buffer (întotdeauna stocați în heap), iar capacitatea este dimensiunea buffer-ului în octeți. Prin urmare, lungimea va fi întotdeauna mai mică sau egală cu capacitatea.

6. 🌟🌟 Dacă un String are suficientă capacitate, adăugarea de elemente la el nu va realoca memorie.
```rust,editable

// Modificați codul de mai jos pentru a afișa
// 25
// 25
// 25
// Aici nu este nevoie să alocăm mai multă memorie în interiorul buclei.
fn main() {
    let mut s = String::new();

    println!("{}", s.capacity());

    for _ in 0..2 {
        s.push_str("hello");
        println!("{}", s.capacity());
    }

    println!("Success!");
}
```

7. 🌟🌟🌟
```rust,editable

// UMPLEȚI spațiile goale
use std::mem;

fn main() {
    let story = String::from("Rust By Practice");

    // Împiedică eliminarea automată a datelor din String
    let mut story = mem::ManuallyDrop::new(story);

    let ptr = story.__();
    let len = story.__();
    let capacity = story.__();

    assert_eq!(16, len);

    // Putem reconstrui un String folosind ptr, len și capacity.
    // Totul este nesigur, deoarece suntem responsabili să ne asigurăm că
    // componentele sunt valide:
    let s = unsafe { String::from_raw_parts(ptr, len, capacity) };

    assert_eq!(*story, s);

    println!("Success!");
}
```


### Metode comune
Mai multe exerciții despre metodele String pot fi găsite [aici](../std/String.md).

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie.