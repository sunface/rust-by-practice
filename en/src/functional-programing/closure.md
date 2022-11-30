# Closure
Closures can capture the enclosing environments. For example we can capture the `x` variable :
```rust
fn main() {
    let x = 1;
    let closure = |val| val + x;
    assert_eq!(closure(2), 3);
}
```

From the syntax, we can see that closures are very convenient for on the fly usage. Unlike functions, both the input and return types of a closure can be inferred by the compiler.

```rust
fn main() {
    // Increment via closures and functions.
    fn function(i: i32) -> i32 { i + 1 }

    // Closures are anonymous, here we are binding them to references
    // 
    // These nameless functions are assigned to appropriately named variables.
    let closure_annotated = |i: i32| -> i32 { i + 1 };
    let closure_inferred  = |i     |          i + 1  ;

    let i = 1;
    // Call the function and closures.
    println!("function: {}", function(i));
    println!("closure_annotated: {}", closure_annotated(i));
    println!("closure_inferred: {}", closure_inferred(i));

    // A closure taking no arguments which returns an `i32`.
    // The return type is inferred.
    let one = || 1;
    println!("closure returning one: {}", one());

}
```

## Capturing
Closures can capture variables by borrowing or moving. But they prefer to capture by borrowing and only go lower when required:
- By reference: `&T`
- By mutable reference: `&mut T`
- By value: `T`

1.🌟
```rust,editable
/* Make it work with least amount of changes*/
fn main() {
    let color = String::from("green");

    let print = move || println!("`color`: {}", color);

    print();
    print();

    // `color` can be borrowed immutably again, because the closure only holds
    // an immutable reference to `color`. 
    let _reborrow = &color;

    println!("{}",color);
}
```

2.🌟🌟
```rust,editable
/* Make it work 
- Dont use `_reborrow` and `_count_reborrowed`
- Dont modify `assert_eq`
*/
fn main() {
    let mut count = 0;

    let mut inc = || {
        count += 1;
        println!("`count`: {}", count);
    };

    inc();


    let _reborrow = &count; 

    inc();

    // The closure no longer needs to borrow `&mut count`. Therefore, it is
    // possible to reborrow without an error
    let _count_reborrowed = &mut count; 

    assert_eq!(count, 0);
}
```

3.🌟🌟
```rust,editable
/* Make it work in two ways, none of them is to remove `take(movable)` away from the code
*/
fn main() {
     let movable = Box::new(3);

     let consume = || {
         println!("`movable`: {:?}", movable);
         take(movable);
     };

     consume();
     consume();
}

fn take<T>(_v: T) {}
```

For comparison, the following code has no error: 
```rust
fn main() {
     let movable = Box::new(3);

     let consume = move || {
         println!("`movable`: {:?}", movable);
     };

     consume();
     consume();
}
```

## Type inferred
The following four closures has no difference in input and return  types.

```rust
fn  add_one_v1   (x: u32) -> u32 { x + 1 }
let add_one_v2 = |x: u32| -> u32 { x + 1 };
let add_one_v3 = |x|             { x + 1 };
let add_one_v4 = |x|               x + 1  ;
```

4.🌟
```rust,editable
fn main() {
    let example_closure = |x| x;

    let s = example_closure(String::from("hello"));

    /* Make it work, only change the following line */
    let n = example_closure(5);
}
```

## Fn, FnMut, FnOnce
When taking a closure as an input parameter, the closure's complete type must be annotated using one of the following traits:

- Fn: the closure uses the captured value by reference (&T)
- FnMut: the closure uses the captured value by mutable reference (&mut T)
- FnOnce: the closure uses the captured value by value (T)

5.🌟🌟
```rust,editable
/* Make it work by changing the trait bound, in two ways*/
fn fn_once<F>(func: F)
where
    F: FnOnce(usize) -> bool,
{
    println!("{}", func(3));
    println!("{}", func(4));
}

fn main() {
    let x = vec![1, 2, 3];
    fn_once(|z|{z == x.len()})
}
```

6. 🌟🌟
```rust,editable
fn main() {
    let mut s = String::new();

    let update_string = |str| s.push_str(str);

    exec(update_string);

    println!("{:?}",s);
}

/* Fill in the blank */
fn exec<'a, F: __>(mut f: F)  {
    f("hello")
}
```
 
#### Which trait does the compiler prefer to use?
- Fn: the closure uses the captured value by reference (&T)
- FnMut: the closure uses the captured value by mutable reference (&mut T)
- FnOnce: the closure uses the captured value by value (T)

On a variable-by-variable basis, the compiler will capture variables in the least restrictive manner possible.

For instance, consider a parameter annotated as FnOnce. This specifies that the closure may capture by `&T`, `&mut T`, or `T`, but the compiler will ultimately choose based on how the captured variables are used in the closure.
Which trait to use is determined by what the closure does with captured value. 

This is because if a move is possible, then any type of borrow should also be possible. Note that the reverse is not true. If the parameter is annotated as `Fn`, then capturing variables by `&mut T` or `T` are not allowed.

7.🌟🌟
```rust,editable
/* Fill in the blank */

// A function which takes a closure as an argument and calls it.
// <F> denotes that F is a "Generic type parameter"
fn apply<F>(f: F) where
    // The closure takes no input and returns nothing.
    F: __ {

    f();
}

// A function which takes a closure and returns an `i32`.
fn apply_to_3<F>(f: F) -> i32 where
    // The closure takes an `i32` and returns an `i32`.
    F: Fn(i32) -> i32 {

    f(3)
}

fn main() {
    use std::mem;

    let greeting = "hello";
    // A non-copy type.
    // `to_owned` creates owned data from borrowed one
    let mut farewell = "goodbye".to_owned();

    // Capture 2 variables: `greeting` by reference and
    // `farewell` by value.
    let diary = || {
        // `greeting` is by reference: requires `Fn`.
        println!("I said {}.", greeting);

        // Mutation forces `farewell` to be captured by
        // mutable reference. Now requires `FnMut`.
        farewell.push_str("!!!");
        println!("Then I screamed {}.", farewell);
        println!("Now I can sleep. zzzzz");

        // Manually calling drop forces `farewell` to
        // be captured by value. Now requires `FnOnce`.
        mem::drop(farewell);
    };

    // Call the function which applies the closure.
    apply(diary);

    // `double` satisfies `apply_to_3`'s trait bound
    let double = |x| 2 * x;

    println!("3 doubled: {}", apply_to_3(double));
}
```

Move closures may still implement `Fn` or `FnMut`, even though they capture variables by move. This is because the traits implemented by a closure type are determined by what the closure does with captured values, not how it captures them. The `move` keyword only specifies the latter.

```rust
fn main() {
    let s = String::new();

    let update_string = move || println!("{}",s);

    exec(update_string);
}

fn exec<F: FnOnce()>(f: F)  {
    f()
}
```

The following code also has no error:
```rust
fn main() {
    let s = String::new();

    let update_string = move || println!("{}",s);

    exec(update_string);
}

fn exec<F: Fn()>(f: F)  {
    f()
}
```

8.🌟🌟
```rust,editable
/* Fill in the blank */
fn main() {
    let mut s = String::new();

    let update_string = |str| -> String {s.push_str(str); s };

    exec(update_string);
}

fn exec<'a, F: __>(mut f: F) {
    f("hello");
}
```


## Input functions
Since closure can be used as arguments, you might wonder can we use functions as arguments too? And indeed we can.

9.🌟🌟
```rust,editable

/* Implement `call_me` to make it work */
fn call_me {
    f();
}

fn function() {
    println!("I'm a function!");
}

fn main() {
    let closure = || println!("I'm a closure!");

    call_me(closure);
    call_me(function);
}
```

## Closure as return types
Returning a closure is much harder than you may have thought of.

10.🌟🌟
```rust,editable
/* Fill in the blank using two aproaches,
 and fix the errror */
fn create_fn() -> __ {
    let num = 5;

    // How does the following closure capture the environment variable `num`
    // &T, &mut T, T ?
    |x| x + num
}


fn main() {
    let fn_plain = create_fn();
    fn_plain(1);
}
```

11.🌟🌟
```rust,editable
/* Fill in the blank and fix the error*/
fn factory(x:i32) -> __ {

    let num = 5;

    if x > 1{
        move |x| x + num
    } else {
        move |x| x + num
    }
}
```


## Closure in structs

**Example**
```rust
struct Cacher<T,E>
where
    T: Fn(E) -> E,
    E: Copy
{
    query: T,
    value: Option<E>,
}

impl<T,E> Cacher<T,E>
where
    T: Fn(E) -> E,
    E: Copy
{
    fn new(query: T) -> Cacher<T,E> {
        Cacher {
            query,
            value: None,
        }
    }

    fn value(&mut self, arg: E) -> E {
        match self.value {
            Some(v) => v,
            None => {
                let v = (self.query)(arg);
                self.value = Some(v);
                v
            }
        }
    }
}
fn main() {
  
}

#[test]
fn call_with_different_values() {
    let mut c = Cacher::new(|a| a);

    let v1 = c.value(1);
    let v2 = c.value(2);

    assert_eq!(v2, 1);
}
```