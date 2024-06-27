# Vector
Vectors are resizable arrays. Like slices, their size is not known at compile time, but they can grow or shrink at any time. 

### Basic Operations
1. 🌟🌟🌟
```rust,editable

fn main() {
    let arr: [u8; 3] = [1, 2, 3];
    
    let v = Vec::from(arr);
    is_vec(&v);

    let v = vec![1, 2, 3];
    is_vec(&v);

    // vec!(..) and vec![..] are same macros, so
    let v = vec!(1, 2, 3);
    is_vec(&v);
    
    // In code below, v is Vec<[u8; 3]> , not Vec<u8>
    // USE Vec::new and `for` to rewrite the below code 
    let v1 = vec!(arr);
    is_vec(&v1);
 
    assert_eq!(v, v1);

    println!("Success!");
}

fn is_vec(v: &Vec<u8>) {}
```



2. 🌟🌟 A Vec can be extended with `extend` method
```rust,editable

// FILL in the blank
fn main() {
    let mut v1 = Vec::from([1, 2, 4]);
    v1.pop();
    v1.push(3);
    
    let mut v2 = Vec::new();
    v2.__;

    assert_eq!(v1, v2);

    println!("Success!");
}
```

### Turn X Into Vec
3. 🌟🌟🌟
```rust,editable

// FILL in the blanks
fn main() {
    // Array -> Vec
    // impl From<[T; N]> for Vec
    let arr = [1, 2, 3];
    let v1 = __(arr);
    let v2: Vec<i32> = arr.__();
 
    assert_eq!(v1, v2);
 
    
    // String -> Vec
    // impl From<String> for Vec
    let s = "hello".to_string();
    let v1: Vec<u8> = s.__();

    let s = "hello".to_string();
    let v2 = s.into_bytes();
    assert_eq!(v1, v2);

    // impl<'_> From<&'_ str> for Vec
    let s = "hello";
    let v3 = Vec::__(s);
    assert_eq!(v2, v3);

    // Iterators can be collected into vectors
    let v4: Vec<i32> = [0; 10].into_iter().collect();
    assert_eq!(v4, vec![0; 10]);

    println!("Success!");
 }
```

### Indexing
4. 🌟🌟🌟
```rust,editable

// FIX the error and IMPLEMENT the code
fn main() {
    let mut v = Vec::from([1, 2, 3]);
    for i in 0..5 {
        println!("{:?}", v[i])
    }

    for i in 0..5 {
       // IMPLEMENT the code here...
    }
    
    assert_eq!(v, vec![2, 3, 4, 5, 6]);

    println!("Success!");
}
```


### Slicing
Immutable or mutable slices of Vecs can be taken, using `&` or `&mut`, respectively.

In Rust, it’s more common to pass immutable slices as arguments rather than vectors when you just want to provide read access, as this is more flexible (no move) and efficient (no copy). The same goes for `String` and `&str`.

5. 🌟🌟
```rust,editable

// FIX the errors
fn main() {
    let mut v = vec![1, 2, 3];

    let slice1 = &v[..];
    // Out of bounds will cause a panic
    // You must use `v.len` here
    let slice2 = &v[0..4];
    
    assert_eq!(slice1, slice2);
    
    // A slice can also be mutable, in which
    // case mutating it will mutate its underlying Vec.
    // Note: slice and &Vec are different
    let vec_ref: &mut Vec<i32> = &mut v;
    (*vec_ref).push(4);
    let slice3 = &mut v[0..3];
    slice3[3] = 42;

    assert_eq!(slice3, &[1, 2, 3, 42]);
    assert_eq!(v, &[1, 2, 3, 42]);

    println!("Success!");
}
```
### Capacity
The capacity of a vector is the amount of space allocated for any future elements that will be added onto the vector. This is not to be confused with the length of a vector, which specifies the number of actual elements within the vector. If a vector’s length exceeds its capacity, its capacity will automatically be increased, but its elements will have to be reallocated.

For example, a vector with capacity 10 and length 0 would be an empty vector with space for 10 more elements. Pushing 10 or fewer elements onto the vector will not change its capacity or cause reallocation to occur. However, if the vector’s length is increased to 11, it will have to reallocate, which can be slow. For this reason, it is recommended to use `Vec::with_capacity `whenever possible to specify how big the vector is expected to get.

6. 🌟🌟
```rust,editable
// FIX the errors
fn main() {
    let mut vec = Vec::with_capacity(10);

    // The vector contains no items, even though it has capacity for more
    assert_eq!(vec.len(), __);
    assert_eq!(vec.capacity(), 10);

    // These are all done without reallocating...
    for i in 0..10 {
        vec.push(i);
    }
    assert_eq!(vec.len(), __);
    assert_eq!(vec.capacity(), __);

    // ...but this may make the vector reallocate
    vec.push(11);
    assert_eq!(vec.len(), 11);
    assert!(vec.capacity() >= 11);


    // Fill in an appropriate value to make the `for` done without reallocating 
    let mut vec = Vec::with_capacity(__);
    for i in 0..100 {
        vec.push(i);
    }

    assert_eq!(vec.len(), __);
    assert_eq!(vec.capacity(), __);
    
    println!("Success!");
}
```

### Store distinct types in Vector
The elements in a vector must be the same type, for example , the code below will cause an error:
```rust
fn main() {
   let v = vec![1, 2.0, 3];
}
```

But we can use enums or trait objects to store distinct types.

7. 🌟🌟
```rust,editable
#[derive(Debug)]
enum IpAddr {
    V4(String),
    V6(String),
}
fn main() {
    // FILL in the blank
    let v : Vec<IpAddr>= __;
    
    // Comparing two enums need to derive the PartialEq trait
    assert_eq!(v[0], IpAddr::V4("127.0.0.1".to_string()));
    assert_eq!(v[1], IpAddr::V6("::1".to_string()));

    println!("Success!");
}
```

8. 🌟🌟
```rust,editable
trait IpAddr {
    fn display(&self);
}

struct V4(String);
impl IpAddr for V4 {
    fn display(&self) {
        println!("ipv4: {:?}",self.0)
    }
}
struct V6(String);
impl IpAddr for V6 {
    fn display(&self) {
        println!("ipv6: {:?}",self.0)
    }
}

fn main() {
    // FILL in the blank
    let v: __= vec![
        Box::new(V4("127.0.0.1".to_string())),
        Box::new(V6("::1".to_string())),
    ];

    for ip in v {
        ip.display();
    }
}
```
