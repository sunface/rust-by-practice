# Ownership

🌟🌟 
```rust,editable

fn main() {
    // modify this line only! use as many approaches as you can
    let x = String::from("hello, world");
    let y = x;
    println!("{},{}",x,y);
}
```

🌟🌟
```rust,editable
// Don't modify code in main!
fn main() {
    let s1 = String::from("hello, world");
    let s2 = take_ownership(s1);

    println!("{}", s2);
}

// Only modify the code below!
fn take_ownership(s: String) {
    println!("{}", s);
}
```


🌟🌟
```rust,editable

fn main() {
    let s = give_ownership();
    println!("{}", s);
}

// Only modify the code below!
fn give_ownership() -> String {
    let s = String::from("hello, world");
    // convert String to Vec
    let _s = s.into_bytes();
    s
}
```

🌟🌟
```rust,editable
// use clone to fix it
fn main() {
    let s = String::from("hello, world");

    print_str(s);

    println!("{}", s);
}

fn print_str(s: String)  {
    println!("{}",s)
}
```

🌟🌟 
```rust, editable
// don't use clone ,use copy instead
fn main() {
    let x = (1, 2, (), "hello");
    let y = x.clone();
    println!("{:?}, {:?}", x, y);
}
```

#### Mutability
Mutability can be changed when ownership is transferred.

🌟
```rust,editable

fn main() {
    let s = String::from("hello, ");
    
    // modify this line only !
    let s1 = s;

    s1.push_str("world")
}
```

🌟🌟🌟
```rust,editable

fn main() {
    let x = Box::new(5);
    
    let ...      // implement this line, dont change other lines!
    
    *y = 4;
    
    assert_eq!(*x, 5);
}
```

### Partial move
Within the destructuring of a single variable, both by-move and by-reference pattern bindings can be used at the same time. Doing this will result in a partial move of the variable, which means that parts of the variable will be moved while other parts stay. In such a case, the parent variable cannot be used afterwards as a whole, however the parts that are only referenced (and not moved) can still be used.

#### Example
```rust,editable

fn main() {
    #[derive(Debug)]
    struct Person {
        name: String,
        age: Box<u8>,
    }

    let person = Person {
        name: String::from("Alice"),
        age: Box::new(20),
    };

    // `name` is moved out of person, but `age` is referenced
    let Person { name, ref age } = person;

    println!("The person's age is {}", age);

    println!("The person's name is {}", name);

    // Error! borrow of partially moved value: `person` partial move occurs
    //println!("The person struct is {:?}", person);

    // `person` cannot be used but `person.age` can be used as it is not moved
    println!("The person's age from person struct is {}", person.age);
}
```

#### Exercises

🌟
```rust,editable

fn main() {
   let t = (String::from("hello"), String::from("world"));

   let _s = t.0;

   // modify this line only, don't use `_s`
   println!("{:?}", t);
}
```

🌟🌟
```rust,editable

fn main() {
   let t = (String::from("hello"), String::from("world"));

   // fill the blanks
   let (__, __) = t;

   println!("{:?}, {:?}, {:?}", s1, s2, t);
}
```
