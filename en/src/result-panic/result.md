# result and ?
`Result<T>` is an enum to describe possible errors. It has two variants: 

- `Ok(T)`: A value T was found
- `Err(e)`: An error was found with a value `e`

In short words, the expected outcome is `Ok`, while the unexpected outcome is `Err`.

1. 🌟🌟
```rust,editable

// FILL in the blanks and FIX the errors
use std::num::ParseIntError;

fn multiply(n1_str: &str, n2_str: &str) -> __ {
    let n1 = n1_str.parse::<i32>();
    let n2 = n2_str.parse::<i32>();
    Ok(n1.unwrap() * n2.unwrap())
}

fn main() {
    let result = multiply("10", "2");
    assert_eq!(result, __);

    let result = multiply("t", "2");
    assert_eq!(result.__, 8);

    println!("Success!");
}
```

### ? 
`?` is almost exactly equivalent to `unwrap`, but `?` returns instead of panic on `Err`.

2. 🌟🌟
```rust,editable

use std::num::ParseIntError;

// IMPLEMENT multiply with ?
// DON'T use unwrap here
fn multiply(n1_str: &str, n2_str: &str) -> __ {
}

fn main() {
    assert_eq!(multiply("3", "4").unwrap(), 12);
    println!("Success!");
}
```

3. 🌟🌟
```rust,editable

use std::fs::File;
use std::io::{self, Read};

fn read_file1() -> Result<String, io::Error> {
    let f = File::open("hello.txt");
    let mut f = match f {
        Ok(file) => file,
        Err(e) => return Err(e),
    };

    let mut s = String::new();
    match f.read_to_string(&mut s) {
        Ok(_) => Ok(s),
        Err(e) => Err(e),
    }
}

// FILL in the blanks with one code line
// DON'T change any code lines
fn read_file2() -> Result<String, io::Error> {
    let mut s = String::new();

    __;

    Ok(s)
}

fn main() {
    assert_eq!(read_file1().unwrap_err().to_string(), read_file2().unwrap_err().to_string());
    println!("Success!");
}
```

### map & and_then
[map](https://doc.rust-lang.org/stable/std/result/enum.Result.html#method.map) and [and_then](https://doc.rust-lang.org/stable/std/result/enum.Result.html#method.and_then) are two common combinators for `Result<T, E>` (also for `Option<T>`).

4. 🌟🌟 

```rust,editable
use std::num::ParseIntError;

// FILL in the blank in two ways: map, and then
fn add_two(n_str: &str) -> Result<i32, ParseIntError> {
   n_str.parse::<i32>().__
}

fn main() {
    assert_eq!(add_two("4").unwrap(), 6);

    println!("Success!");
}
```

5. 🌟🌟🌟
```rust,editable
use std::num::ParseIntError;

// With the return type rewritten, we use pattern matching without `unwrap()`.
// But it's so Verbose...
fn multiply(n1_str: &str, n2_str: &str) -> Result<i32, ParseIntError> {
    match n1_str.parse::<i32>() {
        Ok(n1)  => {
            match n2_str.parse::<i32>() {
                Ok(n2)  => {
                    Ok(n1 * n2)
                },
                Err(e) => Err(e),
            }
        },
        Err(e) => Err(e),
    }
}

// Rewriting `multiply` to make it succinct
// You should use BOTH of  `and_then` and `map` here.
fn multiply1(n1_str: &str, n2_str: &str) -> Result<i32, ParseIntError> {
    // IMPLEMENT...
}

fn print(result: Result<i32, ParseIntError>) {
    match result {
        Ok(n)  => println!("n is {}", n),
        Err(e) => println!("Error: {}", e),
    }
}

fn main() {
    // This still presents a reasonable answer.
    let twenty = multiply1("10", "2");
    print(twenty);

    // The following now provides a much more helpful error message.
    let tt = multiply("t", "2");
    print(tt);

    println!("Success!");
}
```

### Type alias
Using `std::result::Result<T, ParseIntError>` everywhere is verbose and tedious, we can use alias for this purpose.

At a module level, creating aliases can be particularly helpful. Errors found in  a specific module often has the same `Err` type, so a single alias can succinctly defined all associated `Results`. This is so useful even the `std` library supplies one: [`io::Result`](https://doc.rust-lang.org/std/io/type.Result.html).

6. 🌟
```rust,editable
use std::num::ParseIntError;

// FILL in the blank
type __;

// Use the above alias to refer to our specific `Result` type.
fn multiply(first_number_str: &str, second_number_str: &str) -> Res<i32> {
    first_number_str.parse::<i32>().and_then(|first_number| {
        second_number_str.parse::<i32>().map(|second_number| first_number * second_number)
    })
}

// Here, the alias again allows us to save some space.
fn print(result: Res<i32>) {
    match result {
        Ok(n)  => println!("n is {}", n),
        Err(e) => println!("Error: {}", e),
    }
}

fn main() {
    print(multiply("10", "2"));
    print(multiply("t", "2"));

    println!("Success!");
}
```

### Using Result in `fn main`
Typically `the` main function will look like this: 
```rust
fn main() {
    println!("Hello World!");
}
```

However `main` is also able to have a return type of `Result`. If an error occurs within the `main` function it will return an error code and print a debug representation of the error( Debug trait ).

The following example shows such a scenario:
```rust,editable

use std::num::ParseIntError;

fn main() -> Result<(), ParseIntError> {
    let number_str = "10";
    let number = match number_str.parse::<i32>() {
        Ok(number)  => number,
        Err(e) => return Err(e),
    };
    println!("{}", number);
    Ok(())
}
```
> You can find the solutions [here](https://github.com/sunface/rust-by-practice/blob/master/solutions/result-panic/result.md)(under the solutions path), but only use it when you need it :)
