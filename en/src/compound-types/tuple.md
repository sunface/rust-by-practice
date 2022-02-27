# Tuple
🌟 Elements in a tuple can have different types. Tuple's type signature is `(T1, T2, ...)`, where `T1`, `T2` are the types of tuple's members.
```rust,editable

fn main() {
    let _t0: (u8,i16) = (0, -1);
    // Tuples can be tuple's members
    let _t1: (u8, (i16, u32)) = (0, (-1, 1));
    // fill the blanks to make the code work
    let t: (u8, __, i64, __, __) = (1u8, 2u16, 3i64, "hello", String::from(", world"));
}
```

🌟 Members can be extracted from the tuple using indexing.
```rust,editable

// make it works
fn main() {
   let t = ("i", "am", "sunface");
   assert_eq!(t.1, "sunface");
}
```

🌟 Long tuples  cannot be printed
```rust,editable

// fix the error
fn main() {
    let too_long_tuple = (1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13);
    println!("too long tuple: {:?}", too_long_tuple);
}
```

🌟 Destructuring tuple with patter.
```rust,editable

fn main() {
    let tup = (1, 6.4, "hello");

    // fill the blank to make the code work
    let __ = tup;

    assert_eq!(x, 1);
    assert_eq!(y, "hello");
    assert_eq!(z, 6.4);
}
```

🌟🌟 Destructure assignments.
```rust,editable
fn main() {
    let (x, y, z);

    __ = (1, 2, 3);
    
    assert_eq!(x, 3);
    assert_eq!(y, 1);
    assert_eq!(z, 2);
}
```

🌟🌟 Tuples can be used as function arguments and return values
```rust,editable

fn main() {
   // fill the blank to make it work
   let (x, y) = sum_multiply(__);

   assert_eq!(x, 5);
   assert_eq!(y, 6);
}

fn sum_multiply(nums: (i32, i32)) -> (i32, i32) {
    (nums.0 + nums.1, nums.0 * nums.1)
}
```
