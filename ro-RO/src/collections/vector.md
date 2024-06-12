# Vector
Vectorii sunt tablouri redimensionabile. La fel ca și slice-urile, dimensiunea lor nu este cunoscută în timpul compilării, dar pot să crească sau să scadă în orice moment.


### Operații de bază
1. 🌟🌟🌟
```rust,editable

fn main() {
    let arr: [u8; 3] = [1, 2, 3];
    
    let v = Vec::from(arr);
    is_vec(&v);

    let v = vec![1, 2, 3];
    is_vec(&v);

    // vec!(..) și vec![..] sunt aceleași macrocomenzi, deci
    let v = vec!(1, 2, 3);
    is_vec(&v);

    // În cod-ul de mai jos, v este Vec<[u8; 3]>, nu Vec<u8>
    // FOLOSIȚI Vec::new și `for` pentru a rescrie codul de mai jos
    let v1 = vec!(arr);
    is_vec(&v1);
 
    assert_eq!(v, v1);

    println!("Success!");
}

fn is_vec(v: &Vec<u8>) {}
```



2. 🌟🌟 Un vector poate fi extins cu metoda `extend`
```rust,editable

// UMPLEȚI spațiile goale
fn main() {
    let mut v1 = Vec::from([1, 2, 4]);
    v1.pop();
    v1.push(3);
    
    let mut v2 = Vec::new();
    v2.__;

    assert_eq!(v1, v2);

    println!("Success!");
}
```

### Transformați X în Vec
3. 🌟🌟🌟
```rust,editable

// UMPLEȚI spațiile goale
fn main() {
    // Array -> Vec
    // impl From<[T; N]> pentru Vec
    let arr = [1, 2, 3];
    let v1 = __(arr);
    let v2: Vec<i32> = arr.__();
 
    assert_eq!(v1, v2);
 
    
    // String -> Vec
    // impl From<String> pentru Vec
    let s = "hello".to_string();
    let v1: Vec<u8> = s.__();

    let s = "hello".to_string();
    let v2 = s.into_bytes();
    assert_eq!(v1, v2);

    // impl<'_> From<&'_ str> pentru Vec
    let s = "hello";
    let v3 = Vec::__(s);
    assert_eq!(v2, v3);

    // Iteratoarele pot fi colectate în vectori
    let v4: Vec<i32> = [0; 10].into_iter().collect();
    assert_eq!(v4, vec![0; 10]);

    println!("Success!");
 }
```

### Indexarea
4. 🌟🌟🌟
```rust,editable

// REZOLVAȚI erorile și IMPLEMENTAȚI codul
fn main() {
    let mut v = Vec::from([1, 2, 3]);
    for i in 0..5 {
        println!("{:?}", v[i])
    }

    for i in 0..5 {
       // IMPLEMENTAȚI codul aici...
    }
    
    assert_eq!(v, vec![2, 3, 4, 5, 6]);

    println!("Success!");
}
```


### Felii (Slicing)
Un Vec poate fi mutabil. Pe de altă parte, slice-urile sunt obiecte doar pentru citire. Pentru a obține un slice, folosiți &.

În Rust, este mai obișnuit să se transmită slice-uri ca argumente în loc de vectori atunci când doriți doar să furnizați acces de citire. Același lucru este valabil pentru String și &str.

5. 🌟🌟
```rust,editable

// REZOLVAȚI erorile
fn main() {
    let mut v = vec![1, 2, 3];

    let slice1 = &v[..];
    // Accesul în afara limitelor va cauza o panică
    // Trebuie să utilizați v.len aici
    let slice2 = &v[0..4];
    
    assert_eq!(slice1, slice2);
    
    // Slice-urile sunt doar pentru citire
    // Notă: slice-urile și &Vec sunt diferite
    let vec_ref: &mut Vec<i32> = &mut v;
    (*vec_ref).push(4);
    let slice3 = &mut v[0..3];
    slice3.push(4);

    assert_eq!(slice3, &[1, 2, 3, 4]);

    println!("Success!");
}
```
### Capacitate
Capacitatea unui vector reprezintă spațiul alocat pentru orice elemente viitoare care vor fi adăugate la vector. Aceasta nu trebuie confundată cu lungimea unui vector, care specifică numărul real de elemente din vector. Dacă lungimea unui vector depășește capacitatea sa, capacitatea acestuia va fi crescută automat, dar elementele vor trebui să fie realocate.

De exemplu, un vector cu o capacitate de 10 și o lungime de 0 ar fi un vector gol cu spațiu pentru încă 10 elemente. Adăugarea a 10 sau mai puține elemente în vector nu va schimba capacitatea sau nu va provoca realocare. Cu toate acestea, dacă lungimea vectorului este crescută la 11, va trebui să fie realocat, ceea ce poate fi lent. Din acest motiv, se recomandă utilizarea Vec::with_capacity ori de câte ori este posibil pentru a specifica cât de mare se așteaptă să devină vectorul.

6. 🌟🌟
```rust,editable
// REZOLVAȚI erorile
fn main() {
    let mut vec = Vec::with_capacity(10);

    // Vectorul nu conține niciun element, chiar dacă are capacitate pentru mai multe.
    assert_eq!(vec.len(), __);
    assert_eq!(vec.capacity(), 10);

    // Toate acestea se fac fără a realoca memorie...
    for i in 0..10 {
        vec.push(i);
    }
    assert_eq!(vec.len(), __);
    assert_eq!(vec.capacity(), __);

    // ...dar acest lucru ar putea determina vectorul să realoce memorie.
    vec.push(11);
    assert_eq!(vec.len(), 11);
    assert!(vec.capacity() >= 11);

    // Completați cu o valoare potrivită pentru a face ca bucla for să se execute fără alocare suplimentară de memorie
    let mut vec = Vec::with_capacity(__);
    for i in 0..100 {
        vec.push(i);
    }

    assert_eq!(vec.len(), __);
    assert_eq!(vec.capacity(), __);
    
    println!("Success!");
}
```

### Stocarea de tipuri distincte într-un Vector
Elementele dintr-un vector trebuie să fie de același tip. De exemplu, codul de mai jos va genera o eroare:
```rust
fn main() {
   let v = vec![1, 2.0, 3];
}
```

Dar putem folosi enum-uri sau obiecte de tip trait pentru a stoca tipuri distincte.

7. 🌟🌟
```rust,editable
#[derive(Debug)]
enum IpAddr {
    V4(String),
    V6(String),
}
fn main() {
    // COMPLETAȚI spațiul liber
    let v : Vec<IpAddr>= __;
    
    // Pentru a compara două enum-uri, este necesar să derivăm traitul PartialEq.
    assert_eq!(v[0], IpAddr::V4("127.0.0.1".to_string()));
    assert_eq!(v[1], IpAddr::V6("::1".to_string()));

    println!("Success!");
}
```

8. 🌟🌟
```rust,editable
trait IpAddr {
    fn display(&self);
}

struct V4(String);
impl IpAddr for V4 {
    fn display(&self) {
        println!("ipv4: {:?}",self.0)
    }
}
struct V6(String);
impl IpAddr for V6 {
    fn display(&self) {
        println!("ipv6: {:?}",self.0)
    }
}

fn main() {
    // COMPLETAȚI spațiul liber
    let v: __= vec![
        Box::new(V4("127.0.0.1".to_string())),
        Box::new(V6("::1".to_string())),
    ];

    for ip in v {
        ip.display();
    }
}
```
