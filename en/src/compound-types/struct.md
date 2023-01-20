# Struct

### The types of structs
1. 🌟 We must specify concrete values for each of the fields in struct.
```rust,editable

// Fix the error
struct Person {
    name: String,
    age: u8,
    hobby: String
}
fn main() {
    let age = 30;
    let p = Person {
        name: String::from("sunface"),
        age,
    };

    println!("Success!");
} 
```


2. 🌟 Unit struct don't have any fields. It can be useful when you need to implement a trait on some type but don’t have any data that you want to store in the type itself. 
```rust,editable

struct Unit;
trait SomeTrait {
    // ...Some behaviors defined here.
}

// We don't care about what fields  are  in the Unit, but we care about its behaviors.
// So we use a struct with no fields and implement some behaviors for it
impl SomeTrait for Unit {  }
fn main() {
    let u = Unit;
    do_something_with_unit(u);

    println!("Success!");
} 

// Fill the blank to make the code work
fn do_something_with_unit(u: __) {   }
```

3. 🌟🌟🌟 Tuple struct looks similar to tuples, it has added meaning the struct name provides but has no named fields. It's useful when you want to give the whole tuple a name, but don't care about the fields's names.

```rust,editable

// Fix the error and fill the blanks
struct Color(i32, i32, i32);
struct Point(i32, i32, i32);
fn main() {
    let v = Point(__, __, __);
    check_color(v);

    println!("Success!");
}   

fn check_color(p: Color) {
    let (x, _, _) = p;
    assert_eq!(x, 0);
    assert_eq!(p.1, 127);
    assert_eq!(__, 255);
 }
```

### Operating on structs

4. 🌟 You can make a whole struct mutable when instantiating it, but Rust doesn't allow us to mark only certain fields as mutable.

```rust,editable

// Fill the blank and fix the error without adding/removing new line
struct Person {
    name: String,
    age: u8,
}
fn main() {
    let age = 18;
    let p = Person {
        name: String::from("sunface"),
        age,
    };

    // How can you believe sunface is only 18? 
    p.age = 30;

    // Fill the blank
    __ = String::from("sunfei");

    println!("Success!");
}
```

5. 🌟 Using *field init shorthand syntax* to reduce repetitions.
```rust,editable

// Fill the blank
struct Person {
    name: String,
    age: u8,
}
fn main() {
    println!("Success!");
} 

fn build_person(name: String, age: u8) -> Person {
    Person {
        age,
        __
    }
}
```

6. 🌟 You can create instance from other instance with *struct update syntax*
```rust,editable

// Fill the blank to make the code work
struct User {
    active: bool,
    username: String,
    email: String,
    sign_in_count: u64,
}
fn main() {
    let u1 = User {
        email: String::from("someone@example.com"),
        username: String::from("sunface"),
        active: true,
        sign_in_count: 1,
    };

    let u2 = set_email(u1);

    println!("Success!");
} 

fn set_email(u: User) -> User {
    User {
        email: String::from("contact@im.dev"),
        __
    }
}
```

### Print the structs
7. 🌟🌟 We can use `#[derive(Debug)]` to [make a struct printable](https://doc.rust-lang.org/book/ch05-02-example-structs.html?highlight=%23%5Bderive(Debug)%5D#adding-useful-functionality-with-derived-traits).

```rust,editable

// Fill the blanks to make the code work
#[__]
struct Rectangle {
    width: u32,
    height: u32,
}

fn main() {
    let scale = 2;
    let rect1 = Rectangle {
        width: dbg!(30 * scale), // Print debug info to stderr and assign the value of  `30 * scale` to `width`
        height: 50,
    };

    dbg!(&rect1); // Print debug info to stderr

    println!(__, rect1); // Print debug info to stdout
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

8. 🌟🌟
```rust,editable

// Fix errors to make it work
#[derive(Debug)]
struct File {
    name: String,
    data: String,
}
fn main() {
    let f = File {
        name: String::from("readme.md"),
        data: "Rust By Practice".to_string()
    };

    let _name = f.name;

    // ONLY modify this line
    println!("{}, {}, {:?}",f.name, f.data, f);
} 
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice/blob/master/solutions/compound-types/struct.md)(under the solutions path), but only use it when you need it