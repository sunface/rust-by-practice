# Closure
Închideri (Closures) pot captura mediile înconjurate. De exemplu, putem captura variabila 'x':
```rust
fn main() {
    let x = 1;
    let closure = |val| val + x;
    assert_eq!(closure(2), 3);
}
```

Din sintaxă, putem observa că închiderile sunt foarte convenabile pentru utilizarea pe loc. Spre deosebire de funcții, atât tipurile de intrare, cât și cele de ieșire ale unei închideri pot fi inferate de compilator.

```rust
fn main() {
    // Incrementare prin închideri și funcții.
    fn function(i: i32) -> i32 { i + 1 }

    // Închiderile sunt anonime, aici le legăm la referințe
    // 
    // Aceste funcții fără nume sunt atribuite unor variabile cu denumiri corespunzătoare.
    let closure_annotated = |i: i32| -> i32 { i + 1 };
    let closure_inferred  = |i     |          i + 1  ;

    let i = 1;
    // Apelăm funcția și închiderile.
    println!("function: {}", function(i));
    println!("closure_annotated: {}", closure_annotated(i));
    println!("closure_inferred: {}", closure_inferred(i));

    // O închidere care nu primește argumente și returnează un `i32`.
    // Tipul de returnare este inferat.
    let one = || 1;
    println!("closure returning one: {}", one());

}
```

## Capturarea
Închiderile pot captura variabile prin împrumut sau mutare. Dar ele preferă să captureze prin împrumut și să treacă la mutare doar când este necesar:

- Prin referință: '&T'
- Prin referință mutabilă: '&mut T'
- Prin valoare: 'T'


1. 🌟

```rust,editable
/* Faceți să funcționeze cu cel mai mic număr de modificări posibile */
fn main() {
    let color = String::from("green");

    let print = move || println!("`color`: {}", color);

    print();
    print();

    // `color` poate fi împrumutat imutabil din nou, deoarece închiderea deține doar
    // o referință imutabilă la `color`.
    let _reborrow = &color;

    println!("{}",color);
}
```


2. 🌟🌟

```rust,editable
/* Faceți să funcționeze 
- Nu utilizați `_reborrow` și `_count_reborrowed`
- Nu modificați `assert_eq`
*/
fn main() {
    let mut count = 0;

    let mut inc = || {
        count += 1;
        println!("`count`: {}", count);
    };

    inc();


    let _reborrow = &count; 

    inc();

    // Închiderea nu mai are nevoie să împrumute `&mut count`. Prin urmare, este
    // posibil să se facă un împrumut din nou fără eroare
    let _count_reborrowed = &mut count; 

    assert_eq!(count, 0);
}
```


3. 🌟🌟

```rust,editable
/* Faceți să funcționeze în două moduri, niciunul dintre ele nu 
este să eliminați `take(movable)` din cod */
fn main() {
     let movable = Box::new(3);

     let consume = || {
         println!("`movable`: {:?}", movable);
         take(movable);
     };

     consume();
     consume();
}

fn take<T>(_v: T) {}
```

În comparație, următorul cod nu are eroare:
```rust
fn main() {
     let movable = Box::new(3);

     let consume = move || {
         println!("`movable`: {:?}", movable);
     };

     consume();
     consume();
}
```

## Tipul inferat
Următoarele patru închideri nu au diferențe în tipurile de intrare și ieșire.

```rust
fn  add_one_v1   (x: u32) -> u32 { x + 1 }
let add_one_v2 = |x: u32| -> u32 { x + 1 };
let add_one_v3 = |x|             { x + 1 };
let add_one_v4 = |x|               x + 1  ;
```


4. 🌟

```rust,editable
fn main() {
    let example_closure = |x| x;

    let s = example_closure(String::from("hello"));

    /* Faceți să funcționeze, schimbând doar următoarea linie */
    let n = example_closure(5);
}
```

## Fn, FnMut, FnOnce
Atunci când se preia o închidere ca parametru de intrare, tipul complet al închiderii trebuie să fie adnotat folosind una dintre următoarele trăsături:

- Fn: închiderea folosește valoarea capturată prin referință (&T)
- FnMut: închiderea folosește valoarea capturată prin referință mutabilă (&mut T)
- FnOnce: închiderea folosește valoarea capturată prin valoare (T)


5. 🌟🌟

```rust,editable
/* Faceți să funcționeze schimbând bound-ul trăsăturii, în două feluri*/
fn fn_once<F>(func: F)
where
    F: FnOnce(usize) -> bool,
{
    println!("{}", func(3));
    println!("{}", func(4));
}

fn main() {
    let x = vec![1, 2, 3];
    fn_once(|z|{z == x.len()})
}
```

6. 🌟🌟
```rust,editable
fn main() {
    let mut s = String::new();

    let update_string = |str| s.push_str(str);

    exec(update_string);

    println!("{:?}",s);
}

/* Completați spațiile libere */
fn exec<'a, F: __>(mut f: F)  {
    f("hello")
}
```
 
#### Pe care trăsătură preferă compilatorul să o folosească?
- Fn: închiderea folosește valoarea capturată prin referință (&T)
- FnMut: închiderea folosește valoarea capturată prin referință mutabilă (&mut T)
- FnOnce: închiderea folosește valoarea capturată prin valoare (T)

În mod individual pentru fiecare variabilă, compilatorul va captura variabilele în cel mai puțin restrictiv mod posibil.

De exemplu, să luăm în considerare un parametru notat ca FnOnce. Acest lucru specifică că închiderea poate captura prin '&T', '&mut T' sau 'T', dar compilatorul va alege în funcție de modul în care variabilele capturate sunt utilizate în închidere.
Trăsătura de utilizare este determinată de ceea ce face închiderea cu valorile capturate.

Acest lucru se întâmplă pentru că dacă o mutare este posibilă, atunci orice tip de împrumut ar trebui să fie, de asemenea, posibil. Observați că inversul nu este adevărat. Dacă parametrul este notat ca Fn, atunci capturarea variabilelor prin '&mut T' sau 'T' nu este permisă.


7. 🌟🌟

```rust,editable
/* Completați spațiile libere */

// O funcție care primește o închidere ca argument și o apelează.
// <F> indică faptul că F este un "Parametru de tip generic"
fn apply<F>(f: F) where
    // Închiderea nu primește niciun argument și nu returnează nimic.
    F: __ {

    f();
}

// O funcție care primește o închidere și returnează un `i32`.
fn apply_to_3<F>(f: F) -> i32 where
    // Închiderea primește un `i32` și returnează un `i32`.
    F: Fn(i32) -> i32 {

    f(3)
}

fn main() {
    use std::mem;

    let greeting = "hello";
    // Un tip care nu se copiază.
    // `to_owned` creează date deținute din cele împrumutate
    let mut farewell = "goodbye".to_owned();

    // Capturăm 2 variabile: `greeting` prin referință și
    // `farewell` prin valoare.
    let diary = || {
        // `greeting` este prin referință: necesită `Fn`.
        println!("I said {}.", greeting);

        // Mutarea forțează `farewell` să fie capturat
        // prin referință mutabilă. Acum necesită `FnMut`.
        farewell.push_str("!!!");
        println!("Then I screamed {}.", farewell);
        println!("Now I can sleep. zzzzz");

        // Apelăm manual drop pentru a forța `farewell` să fie
        // capturat prin valoare. Acum necesită `FnOnce`.
        mem::drop(farewell);
    };

    // Apelăm funcția care aplică închiderea.
    apply(diary);

    // `double` satisface condiția de trăsătură a lui `apply_to_3`
    let double = |x| 2 * x;

    println!("3 doubled: {}", apply_to_3(double));
}
```

Închiderile care mută pot implementa în continuare 'Fn' sau 'FnMut', chiar dacă capturează variabile prin mutare. Acest lucru se datorează faptului că trăsăturile implementate de un tip de închidere sunt determinate de ceea ce face închiderea cu valorile capturate, nu cum le capturează. Cuvântul cheie 'move' specifică doar ultimul aspect.

```rust
fn main() {
    let s = String::new();

    let update_string = move || println!("{}",s);

    exec(update_string);
}

fn exec<F: FnOnce()>(f: F)  {
    f()
}
```

Următorul cod nu are nicio eroare:
```rust
fn main() {
    let s = String::new();

    let update_string = move || println!("{}",s);

    exec(update_string);
}

fn exec<F: Fn()>(f: F)  {
    f()
}
```


8. 🌟🌟

```rust,editable
/* Completați spațiile libere */
fn main() {
    let mut s = String::new();

    let update_string = |str| -> String {s.push_str(str); s };

    exec(update_string);
}

fn exec<'a, F: __>(mut f: F) {
    f("hello");
}
```


## Funcții de intrare
Deoarece închiderile pot fi utilizate ca argumente, poate vă întrebați dacă putem utiliza și funcții ca argumente? Și cu siguranță că putem.


9. 🌟🌟

```rust,editable

/* Implementați `call_me` pentru a funcționa */
fn call_me {
    f();
}

fn function() {
    println!("I'm a function!");
}

fn main() {
    let closure = || println!("I'm a closure!");

    call_me(closure);
    call_me(function);
}
```

## Închiderea ca tipuri de returnare
Returnarea unei închideri este mult mai dificilă decât ai fi crezut.


10. 🌟🌟

```rust,editable
/* Completați spațiul liber folosind două abordări,
și rezolvați eroarea */
fn create_fn() -> __ {
    let num = 5;

    // Cum capturează închiderea următoarea variabilă de mediu `num`
    // &T, &mut T, T ?
    |x| x + num
}


fn main() {
    let fn_plain = create_fn();
    fn_plain(1);
}
```


11. 🌟🌟

```rust,editable
/* Completați spațiul liber și rezolvați eroarea*/
fn factory(x:i32) -> __ {

    let num = 5;

    if x > 1{
        move |x| x + num
    } else {
        move |x| x + num
    }
}
```


## Închideri în structuri

**Exemplu**
```rust
struct Cacher<T,E>
where
    T: Fn(E) -> E,
    E: Copy
{
    query: T,
    value: Option<E>,
}

impl<T,E> Cacher<T,E>
where
    T: Fn(E) -> E,
    E: Copy
{
    fn new(query: T) -> Cacher<T,E> {
        Cacher {
            query,
            value: None,
        }
    }

    fn value(&mut self, arg: E) -> E {
        match self.value {
            Some(v) => v,
            None => {
                let v = (self.query)(arg);
                self.value = Some(v);
                v
            }
        }
    }
}
fn main() {
  
}

#[test]
fn call_with_different_values() {
    let mut c = Cacher::new(|a| a);

    let v1 = c.value(1);
    let v2 = c.value(2);

    assert_eq!(v2, 1);
}
```
> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
