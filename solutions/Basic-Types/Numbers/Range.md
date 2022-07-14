# 9. 🌟🌟 Two goals: 1. Modify assert! to make it work 2. Make println! output: 97 - 122
```rs
fn main() {
    let mut sum = 0;
    for i in -3..2 {
        sum += i
    }

    assert!(sum == -3);

    for c in 'a'..='z' {
        println!("{}",c);
    }
}
```
# Solution:
```rs
fn main() {
    let mut sum = 0;
    for i in -3..=2 {
        sum += i;
    }

    assert!(sum == -3);

    for c in 'a'..='z' {
        println!("{}",c as u8); // one can use u8, i8, u16, i16, u32, i32, u64, i64 after c as for conversion of character to integer
    }
}
```

# 10. 🌟🌟
```rs

// Fill the blanks
use std::ops::{Range, RangeInclusive};
fn main() {
    assert_eq!((1..__), Range{ start: 1, end: 5 });
    assert_eq!((1..__), RangeInclusive::new(1, 5));

    println!("Success!");
}
```
# Solution:
```rs
use std::ops::{Range, RangeInclusive};
fn main() {
    assert_eq!((1..5), Range{ start: 1, end: 5 });
    assert_eq!((1..=5), RangeInclusive::new(1, 5));

    println!("Success!");
}
```