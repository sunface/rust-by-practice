# Conversie prin `as`
În limbajul Rust, nu există conversie implicită de tip (coerciție) între tipurile de bază. Cu toate acestea, conversiile explicite de tip pot fi realizate folosind cuvântul cheie `as`.

1. 🌟
```rust,editable
// CORECTAȚI erorile și completați spațiile goale
// NU eliminați niciun cod
fn main() {
    let decimal = 97.123_f32;

    let integer: __ = decimal as u8;

    let c1: char = decimal as char;
    let c2 = integer as char;

    assert_eq!(integer, 'b' as u8);

    println!("Success!");
}
```

2. 🌟🌟 Implicit, depășirea valorilor maxime va provoca erori de compilare, dar putem adăuga o adnotare globală pentru a suprima aceste erori.
```rust,editable
fn main() {
    assert_eq!(u8::MAX, 255);
    // Maximul pentru `u8` este 255, conform afirmației anterioare.
    // Prin urmare, codul de mai jos va provoca o eroare de depășire: litera în afara intervalului pentru `u8`.
    // VA ROG să căutați indicii în erorile de compilare pentru a CORECTA aceasta.
    // NU modificați niciun cod în main.
    let v = 1000 as u8;

    println!("Success!");
}
```

3. 🌟🌟  La conversia oricărei valori la un tip nesemnat `T`, se adaugă sau se scade `T::MAX + 1` până când valoarea se potrivește în noul tip.
```rust,editable
fn main() {
    assert_eq!(1000 as u16, __);

    assert_eq!(1000 as u8, __);

    // Pentru numerele pozitive, acest lucru este echivalent cu modulo
    println!("1000 mod 256 is : {}", 1000 % 256);

    assert_eq!(-1_i8 as u8, __);
    
    // Începând cu Rust 1.45, cuvântul cheie `as` efectuează o *conversie cu săturare*
    // atunci când se face conversia de la float la int. Dacă valoarea cu virgulă depășește
    // limita superioară sau este mai mică decât limita inferioară, valoarea returnată
    // va fi egală cu limita depășită.
    assert_eq!(300.1_f32 as u8, __);
    assert_eq!(-100.1_f32 as u8, __);
    

    // Această comportare implică un cost mic la runtime și poate fi evitată
    // cu metode nesigure, totuși rezultatele ar putea depăși și returna **valori nesigure**.
    /* Use these methods wisely: */
    unsafe {
        // 300.0 este 44
        println!("300.0 is {}", 300.0_f32.to_int_unchecked::<u8>());
        // -100.0 ca u8 este 156
        println!("-100.0 as u8 is {}", (-100.0_f32).to_int_unchecked::<u8>());
        // nan ca u8 este 0
        println!("nan as u8 is {}", f32::NAN.to_int_unchecked::<u8>());
    }
}
```

4. 🌟🌟🌟 Pointerii raw pot fi convertiți la adrese de memorie (integer) și invers.
```rust,editable

// Completați spațiile goale
fn main() {
    let mut values: [i32; 2] = [1, 2];
    let p1: *mut i32 = values.as_mut_ptr();
    let first_address: usize = p1 __; 
    let second_address = first_address + 4; // 4 == std::mem::size_of::<i32>()
    let p2: *mut i32 = second_address __; // p2 indică către al doilea element în values
    unsafe {
        // Adăugați unu la al doilea element
        __
    }
    
    assert_eq!(values[1], 3);

    println!("Success!");
}
```


5. 🌟🌟🌟 
```rust,editable
fn main() {
    let arr :[u64; 13] = [0; 13];
    assert_eq!(std::mem::size_of_val(&arr), 8 * 13);
    let a: *const [u64] = &arr;
    let b = a as *const [u8];
    unsafe {
        assert_eq!(std::mem::size_of_val(&*b), __)
    }

    println!("Success!");
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)