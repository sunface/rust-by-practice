# Advance Traits

## Associated types
The use of "Associated types" improves the overall readability of code by moving inner types locally into a trait as output types. For example :
```rust
pub trait CacheableItem: Clone + Default + fmt::Debug + Decodable + Encodable {
  type Address: AsRef<[u8]> + Clone + fmt::Debug + Eq + Hash;
  fn is_null(&self) -> bool;
}
```

Using of `Address` is much more clearer and convenient than `AsRef<[u8]> + Clone + fmt::Debug + Eq + Hash`.

1. 🌟🌟🌟
```rust,editable

struct Container(i32, i32);

// USING associated types to re-implement trait Contains.
// trait Contains {
//    type A;
//    type B;

trait Contains<A, B> {
    fn contains(&self, _: &A, _: &B) -> bool;
    fn first(&self) -> i32;
    fn last(&self) -> i32;
}

impl Contains<i32, i32> for Container {
    fn contains(&self, number_1: &i32, number_2: &i32) -> bool {
        (&self.0 == number_1) && (&self.1 == number_2)
    }
    // Grab the first number.
    fn first(&self) -> i32 { self.0 }

    // Grab the last number.
    fn last(&self) -> i32 { self.1 }
}

fn difference<A, B, C: Contains<A, B>>(container: &C) -> i32 {
    container.last() - container.first()
}

fn main() {
    let number_1 = 3;
    let number_2 = 10;

    let container = Container(number_1, number_2);

    println!("Does container contain {} and {}: {}",
        &number_1, &number_2,
        container.contains(&number_1, &number_2));
    println!("First number: {}", container.first());
    println!("Last number: {}", container.last());
    
    println!("The difference is: {}", difference(&container));
}
```

## Default Generic Type Parameters
When we use generic type parameters, we can specify a default concrete type for the generic type. This eliminates the need for implementors of the trait to specify a concrete type if the default type works.

2. 🌟🌟
```rust,editable

use std::ops::Sub;

#[derive(Debug, PartialEq)]
struct Point<T> {
    x: T,
    y: T,
}

// FILL in the blank in three ways: two of them use the default generic  parameters, the other one not.
// Notice that the implementation uses the associated type `Output`.
impl __ {
    type Output = Self;

    fn sub(self, other: Self) -> Self::Output {
        Point {
            x: self.x - other.x,
            y: self.y - other.y,
        }
    }
}

fn main() {
    assert_eq!(Point { x: 2, y: 3 } - Point { x: 1, y: 0 },
        Point { x: 1, y: 3 });

    println!("Success!");
}
```

## Fully Qualified Syntax
Nothing in Rust prevents a trait from having a method with the same name as another trait’s method, nor does Rust prevent you from implementing both traits on one type. It’s also possible to implement a method directly on the type with the same name as methods from traits.

When calling methods with the same name, we have to use  Fully Qualified Syntax.

#### Example
```rust,editable
trait UsernameWidget {
    // Get the selected username out of this widget
    fn get(&self) -> String;
}

trait AgeWidget {
    // Get the selected age out of this widget
    fn get(&self) -> u8;
}

// A form with both a UsernameWidget and an AgeWidget.
struct Form {
    username: String,
    age: u8,
}

impl UsernameWidget for Form {
    fn get(&self) -> String {
        self.username.clone()
    }
}

impl AgeWidget for Form {
    fn get(&self) -> u8 {
        self.age
    }
}

fn main() {
    let form = Form{
        username: "rustacean".to_owned(),
        age: 28,
    };

    // If you uncomment this line, you'll get an error saying 
    // "multiple `get` found". Because, after all, there are multiple methods
    // named `get`.
    // println!("{}", form.get());
    
    let username = UsernameWidget::get(&form);
    assert_eq!("rustacean".to_owned(), username);
    let age = AgeWidget::get(&form); // You can also use `<Form as AgeWidget>::get`
    assert_eq!(28, age);

    println!("Success!");
}
```

#### Exercise
3. 🌟🌟
```rust,editable
trait Pilot {
    fn fly(&self) -> String;
}

trait Wizard {
    fn fly(&self) -> String;
}

struct Human;

impl Pilot for Human {
    fn fly(&self) -> String {
        String::from("This is your captain speaking.")
    }
}

impl Wizard for Human {
    fn fly(&self) -> String {
        String::from("Up!")
    }
}

impl Human {
    fn fly(&self) -> String {
        String::from("*waving arms furiously*")
    }
}

fn main() {
    let person = Human;

    assert_eq!(__, "This is your captain speaking.");
    assert_eq!(__, "Up!");

    assert_eq!(__, "*waving arms furiously*");

    println!("Success!");
}
```

## Supertraits
Sometimes, you might need one trait to use another trait’s functionality( like the "inheritance" in other languages ). In this case, you need to rely on the dependent trait also being implemented. The trait you rely on is a `supertrait` of the trait you’re implementing.

4. 🌟🌟🌟
```rust,editable

trait Person {
    fn name(&self) -> String;
}

// Person is a supertrait of Student.
// Implementing Student requires you to also impl Person.
trait Student: Person {
    fn university(&self) -> String;
}

trait Programmer {
    fn fav_language(&self) -> String;
}

// CompSciStudent (computer science student) is a subtrait of both Programmer 
// and Student. Implementing CompSciStudent requires you to impl both supertraits.
trait CompSciStudent: Programmer + Student {
    fn git_username(&self) -> String;
}

fn comp_sci_student_greeting(student: &dyn CompSciStudent) -> String {
    format!(
        "My name is {} and I attend {}. My favorite language is {}. My Git username is {}",
        student.name(),
        student.university(),
        student.fav_language(),
        student.git_username()
    )
}

struct CSStudent {
    name: String,
    university: String,
    fav_language: String,
    git_username: String
}

// IMPLEMENT the necessary traits for CSStudent to make the code work
impl ...

fn main() {
    let student = CSStudent {
        name: "Sunfei".to_string(),
        university: "XXX".to_string(),
        fav_language: "Rust".to_string(),
        git_username: "sunface".to_string()
    };

    // FILL in the blank
    println!("{}", comp_sci_student_greeting(__));
}
```

## Orphan Rules
We can’t implement external traits on external types. For example, we can’t implement the `Display` trait on `Vec<T>` within our own crate, because `Display` and `Vec<T>` are defined in the standard library and aren’t local to our crate. 

This restriction is often called the orphan rule, so named because the parent type is not present. This rule ensures that other people’s code can’t break your code and vice versa. 

It’s possible to get around this restriction using the newtype pattern, which involves creating a new type in a tuple struct.

5. 🌟🌟
```rust,editable
use std::fmt;

// DEFINE a newtype `Pretty` here


impl fmt::Display for Pretty {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "\"{}\"", self.0.clone() + ", world")
    }
}

fn main() {
    let w = Pretty("hello".to_string());
    println!("w = {}", w);
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it :)