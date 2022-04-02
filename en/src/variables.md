# Variables

### Binding and mutability
1. 🌟 A variable can be used only if it has been initialized.
```rust,editable

// fix the error below with least modifying
fn main() {
    let x: i32; // uninitialized but using, ERROR !
    let y: i32; // uninitialized but also unusing, only warning

    assert_eq!(x, 5);
    println!("Success!")
}
```

2. 🌟 Use `mut` to mark a variable as mutable.
```rust,editable

// fill the blanks in code to make it compile
fn main() {
    let __ =  1;
    __ += 2; 
    
    assert_eq!(x, 3);
    println!("Success!")
}
```

### Scope
A scope is the range within the program for which the item is valid.

3. 🌟 
```rust,editable

// fix the error below with least modifying
fn main() {
    let x: i32 = 10;
    {
        let y: i32 = 5;
        println!("The value of x is {} and value of y is {}", x, y);
    }
    println!("The value of x is {} and value of y is {}", x, y); 
}
```

4. 🌟🌟 
```rust,editable

// fix the error with using of define_x
fn main() {
    println!("{}, world", x); 
}

fn define_x() {
    let x = "hello";
}
```

### Shadowing
You can declare a new variable with the same name as a previous variable, here we can say **the first one is shadowed by the second one.

5. 🌟🌟 
```rust,editable

// only modify `assert_eq!` to make the `println!` work(print `42` in terminal)
fn main() {
    let x: i32 = 5;
    {
        let x = 12;
        assert_eq!(x, 5);
    }

    assert_eq!(x, 12);

    let x =  42;
    println!("{}", x); // Prints "42".
}
```

6. 🌟🌟 
```rust,editable

// remove a line in code to make it compile
fn main() {
    let mut x: i32 = 1;
    x = 7;
    // shadowing and re-binding
    let x = x; 
    x += 3;


    let y = 4;
    // shadowing
    let y = "I can also be bound to text!"; 

    println!("Success!")
}
```

### Unused variables
1. fix the warning below with :

- 🌟  only one solution
- 🌟🌟  two distinct solutions

> Note: none of the solutions is to remove the line `let x = 1` 

```rust,editable

fn main() {
    let x = 1; 
}

// warning: unused variable: `x`
```

### Destructuring
8. 🌟🌟 We can use a pattern with `let` to destructure a tuple to separate variables.

> Tips: you can use Shadowing or Mutability

```rust,editable

// fix the error below with least modifying
fn main() {
    let (x, y) = (1, 2);
    x += 2;

    assert_eq!(x, 3);
    assert_eq!(y, 2);

    println!("Success!")
}
```

### Destructuring assignments
Introduced in Rust 1.59: You can now use tuple, slice, and struct patterns as the left-hand side of an assignment.

9. 🌟🌟

> Note: the feature `Destructuring assignments` need 1.59 or higher Rust version

```rust,editable

fn main() {
    let (x, y);
    (x,..) = (3, 4);
    [.., y] = [1, 2];
    // fill the blank to make the code work
    assert_eq!([x,y], __);

    println!("Success!")
} 
```


> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it
