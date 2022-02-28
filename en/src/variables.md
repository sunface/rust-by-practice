# Variables

### Binding and mutablity
🌟 A variable can be used only if it has been initialized.
```rust,editable

// fix the error below with least modifying
fn main() {
    let x: i32; // uninitialized but using, ERROR !
    let y: i32; // uninitialized but also unusing, only warning
    println!("{} is equal to 5", x); 
}
```

🌟 Use `mut` to mark a variable as mutable.
```rust,editable

// fill the blanks in code to make it compile
fn main() {
    let __ =  1;
    __ += 2; 
    
    println!("{} is equal to 3", x); 
}
```

### Scope
🌟 A scope is the range within the program for which the item is valid.
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

🌟🌟 
```rust,editable

// fix the error
fn main() {
    println!("{}, world", x); 
}

fn define_x() {
    let x = "hello";
}
```

### Shadowing
🌟🌟 You can declare a new variable with the same name as a previous variable, here we can say **the first one is shadowed by the second one.

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

🌟🌟 remove a line in code to make it compile
```rust,editable

fn main() {
    let mut x: i32 = 1;
    x = 7;
    // shadowing and re-binding
    let x = x; 
    x += 3;


    let y = 4;
    // shadowing
    let y = "I can also be bound to text!"; 
}
```

### Unused varibles
fix the warning below with :

- 🌟  one way
- 🌟🌟  two ways

> Note: there are two ways you can use, but none of them is removing the line `let x = 1` 

```rust,editable

fn main() {
    let x = 1; 
}

// warning: unused variable: `x`
```

### Destructing
🌟🌟 fix the error below with least modifying

> Tips: you can use Shadowing or Mutability

```rust,editable

fn main() {
    let (x, y) = (1, 2);
    x += 2;

    assert_eq!(x, 3);
    assert_eq!(y, 2);
}
```

### Destructuring assignments
🌟🌟 fix the code with two ways:

- Shadowing with adding `let`
- make types compatible


> Note: the feature `Destructuring assignments` need 1.59 or higher Rust version

```rust,editable

fn main() {
    let (mut x, mut y) = (1.0, 2.0);
    (x,y) = (3, 4);

    assert_eq!([x,y],[3, 4]);
} 
```