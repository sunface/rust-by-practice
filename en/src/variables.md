# Variables

### Binding and mutablity
🌟 fix the error below with least change
```rust,editable

fn main() {
    let x: i32; // uninitialized but using, ERROR !
    let y: i32; // uninitialized but also unusing, only warning
    println!("{} is equal to 5", x); 
}
```

🌟🌟 fill the blanks in code to make it compile
```rust,editable

fn main() {
    // replace __ with a variable name
    let __ =  1;
    __ += 2; 
    
    println!("{} is equal to 3", x); 
}
```

### Scope
🌟 fix the error below with least change
```rust,editable

fn main() {
    let x: i32 = 10;
    {
        let y: i32 = 5;
        println!("The value of x is {} and value of y is {}", x, y);
    }
    println!("The value of x is {} and value of y is {}", x, y); 
}
```

🌟🌟 fix the error with the knowledge you grasped
```rust,editable

fn main() {
    println!("{}, world", x); 
}

fn define_x() {
    let x = "hello";
}
```

### Shadowing
🌟🌟 only change `assert_eq!` to make the `println!` work(print `42` in terminal)

```rust,editable

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
🌟🌟 fix the error below with least change

> Tips: you can use Shadowing or Mutability

```rust,editable

fn main() {
    let (x, y) = (1, 2);
    x += 2;

    assert_eq!(x, 3);
    assert_eq!(y, 2);
}
```