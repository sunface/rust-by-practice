# Flow control

### if/else
🌟 
```rust,editable

// fill in the blanks
fn main() {
    let n = 5;

    if n < 0 {
        println!("{} is negative", n);
    } __ n > 0 {
        println!("{} is positive", n);
    } __ {
        println!("{} is zero", n);
    }
} 
```

🌟🌟 `if/else` expression can be used in assignments.
```rust,editable

// fix the errors
fn main() {
    let n = 5;

    let big_n =
        if n < 10 && n > -10 {
            println!(", and is a small number, increase ten-fold");

            10 * n
        } else {
            println!(", and is a big number, halve the number");

            n / 2.0 ;
        }

    println!("{} -> {}", n, big_n);
} 
```

### for
🌟 The `for in` construct can be used to iterate through an Iterator, e.g a range `a..b`.

```rust,editable

fn main() {
    for n in 1..=100 { // modify this line to make the code work
        if n == 100 {
            panic!("NEVER LET THIS RUN")
        }
    }
} 
```


🌟🌟 
```rust,editable

// fix the errors without adding or removing lines
fn main() {
    let names = [String::from("liming"),String::from("hanmeimei")];
    for name in names {
        // do something with name...
    }

    println!("{:?}", names);

    let numbers = [1, 2, 3];
    // the elements in numbers are Copy，so there is no move here
    for n in numbers {
        // do something with name...
    }
    
    println!("{:?}", numbers);
} 
```

🌟
```rust,editable
fn main() {
    let a = [4, 3, 2, 1];

    // iterate the indexing and value in 'a'
    for (i,v) in a.__ {
        println!("The {}th element is {}",i+1,v);
    }
}
```

### while
The `while` keyword can be used to run a loop when a condition is true.

```rust,editable

// fill in the blanks to make the last println! work !
fn main() {
    // A counter variable
    let mut n = 1;

    // Loop while the condition is true
    while n __ 10 {
        if n % 15 == 0 {
            println!("fizzbuzz");
        } else if n % 3 == 0 {
            println!("fizz");
        } else if n % 5 == 0 {
            println!("buzz");
        } else {
            println!("{}", n);
        }


        __;
    }

    println!("n reached {}, soloop is over",n);
}
```

### continue and break
🌟 use `break` to break the loop.
```rust,editable

// fill in the blank
fn main() {
    let mut n = 0;
    for i in 0..=100 {
       if n == 66 {
           __
       }
       n += 1;
    }

    assert_eq!(n, 66);
}
```

🌟🌟 `continue` will skip over the remaining code in current iteration and go to the next iteration.
```rust,editable

// fill in the blanks
fn main() {
    let mut n = 0;
    for i in 0..=100 {
       if n != 66 {
           n+=1;
           __;
       }
       
       __
    }

    assert_eq!(n, 66);
}
```

### loop 

🌟🌟 loop is usually used together with `break` or `continue`.

```rust,editable

// fill in the blanks
fn main() {
    let mut count = 0u32;

    println!("Let's count until infinity!");

    // Infinite loop
    loop {
        count += 1;

        if count == 3 {
            println!("three");

            // Skip the rest of this iteration
            __;
        }

        println!("{}", count);

        if count == 5 {
            println!("OK, that's enough");

            __;
        }
    }

    assert_eq!(count, 5);
}
```

🌟🌟 loop is an expression, so we can use it with `break` to return a value
```rust,editable

// fill in the blank
fn main() {
    let mut counter = 0;

    let result = loop {
        counter += 1;

        if counter == 10 {
            __;
        }
    };

    assert_eq!(result, 20);
}
```

🌟🌟🌟 It's possible to break or continue outer loops when dealing with nested loops. In these cases, the loops must be annotated with some 'label, and the label must be passed to the break/continue statement.

```rust,editable

// fill in the blank
fn main() {
    let mut count = 0;
    'outer: loop {
        'inner1: loop {
            if count >= 20 {
                // This would break only the inner1 loop
                break 'inner1; // `break` is also ok 
            }
            count += 2;
        }

        count += 5;

        'inner2: loop {
            if count >= 30 {
                // This breaks the outer loop
                break 'outer;
            }

            // This will continue the outer loop
            continue 'outer;
        }
    }

    assert!(count == __)
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it