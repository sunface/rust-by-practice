# Instrucțiuni și Expresii

### Exemple
```rust,editable
fn main() {
    let x = 5u32;

    let y = {
        let x_squared = x * x;
        let x_cube = x_squared * x;

        // Această expresie va fi asignată variabilei y
        x_cube + x_squared + x
    };

    let z = {
        
        // Punctul și virgulă suprimă această expresie și () este asignat variabilei z
        2 * x;
    };

    println!("x is {:?}", x);
    println!("y is {:?}", y);
    println!("z is {:?}", z);
}
```

### Exercises
1. 🌟🌟
```rust,editable
// Faceți să funcționeze în două moduri
fn main() {
   let v = {
       let mut x = 1;
       x += 2
   };

   assert_eq!(v, 3);

   println!("Success!");
}
```

2. 🌟
```rust,editable

fn main() {
   let v = (let x = 3);

   assert!(v == 3);

   println!("Success!");
}
```

3. 🌟
```rust,editable

fn main() {
    let s = sum(1 , 2);
    assert_eq!(s, 3);

    println!("Success!");
}

fn sum(x: i32, y: i32) -> i32 {
    x + y;
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice)(sub calea soluțiilor), dar utilizați-le doar atunci când aveți nevoie.