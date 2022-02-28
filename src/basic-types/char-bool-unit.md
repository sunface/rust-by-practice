# Char, Bool and Unit

### Char
🌟
```rust, editable

// make it work
use std::mem::size_of_val;
fn main() {
    let c1 = 'a';
    assert_eq!(size_of_val(&c1),1); 

    let c2 = '中';
    assert_eq!(size_of_val(&c2),3); 
} 
```

🌟
```rust, editable

// make it work
fn main() {
    let c1 = "中";
    print_char(c1);
} 

fn print_char(c : char) {
    println!("{}", c);
}
```

### Bool
🌟
```rust, editable

// make the println! work
fn main() {
    let _f: bool = false;

    let t = true;
    if !t {
        println!("hello, world");
    }
} 
```

🌟
```rust, editable

// make it work
fn main() {
    let f = true;
    let t = true && false;
    assert_eq!(t, f);
}
```


### Unit type
🌟🌟
```rust,editable

// make it work, don't modify `implicitly_ret_unit` !
fn main() {
    let _v: () = ();

    let v = (2, 3);
    assert_eq!(v, implicitly_ret_unit())
}

fn implicitly_ret_unit() {
    println!("I will returen a ()")
}

// don't use this one
fn explicitly_ret_unit() -> () {
    println!("I will returen a ()")
}
```

🌟🌟 what's the size of the unit type?
```rust,editable

// modify `4` in assert to make it work
use std::mem::size_of_val;
fn main() {
    let unit: () = ();
    assert!(size_of_val(&unit) == 4);
}
```