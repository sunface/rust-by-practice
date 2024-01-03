# Char, Bool and Unit

### Char
1. 🌟
```rust,editable

// Faceți să funcționeze
use std::mem::size_of_val;
fn main() {
    let c1 = 'a';
    assert_eq!(size_of_val(&c1),1); 

    let c2 = '中';
    assert_eq!(size_of_val(&c2),3); 

    println!("Success!");
} 
```

2. 🌟
```rust,editable

// Faceți să funcționeze
fn main() {
    let c1 = "中";
    print_char(c1);
} 

fn print_char(c : char) {
    println!("{}", c);
}
```

### Bool
3. 🌟
```rust,editable

// Faceți println! să funcționeze
fn main() {
    let _f: bool = false;

    let t = true;
    if !t {
        println!("Success!");
    }
} 
```

4. 🌟
```rust,editable

// Faceți să funcționeze
fn main() {
    let f = true;
    let t = true && false;
    assert_eq!(t, f);

    println!("Success!");
}
```


### Unit type
5. 🌟🌟
```rust,editable

// Faceți să funcționeze fără a modifica `implicitly_ret_unit` !
fn main() {
    let _v: () = ();

    let v = (2, 3);
    assert_eq!(v, implicitly_ret_unit());

    println!("Success!");
}

fn implicitly_ret_unit() {
    println!("I will return a ()");
}

// Nu folosiți asta
fn explicitly_ret_unit() -> () {
    println!("I will return a ()");
}
```

6. 🌟🌟 What's the size of the unit type?
```rust,editable

// Modificați `4` in 'assert' pentru a face să funcționeze
use std::mem::size_of_val;
fn main() {
    let unit: () = ();
    assert!(size_of_val(&unit) == 4);

    println!("Success!");
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice)(sub calea soluțiilor), dar utilizați-le doar atunci când aveți nevoie.
