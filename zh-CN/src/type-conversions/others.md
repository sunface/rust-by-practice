# 其它转换

### 将任何类型转换成 String
只要为一个类型实现了 `ToString`，就可以将任何类型转换成 `String`。事实上，这种方式并不是最好的，大家还记得 `fmt::Display` 特征吗？它可以控制一个类型如何打印，在实现它的时候还会自动实现 `ToString`。


1. 🌟🌟
```rust,editable
use std::fmt;

struct Point {
    x: i32,
    y: i32,
}

impl fmt::Display for Point {
    // 实现 fmt 方法
}

fn main() {
    let origin = Point { x: 0, y: 0 };
    // 填空
    assert_eq!(origin.__, "The point is (0, 0)");
    assert_eq!(format!(__), "The point is (0, 0)");

    println!("Success!")
}
```

### 解析 String
2. 🌟🌟🌟 使用 `parse` 方法可以将一个 `String` 转换成 `i32` 数字，这是因为在标准库中为 `i32` 类型实现了 `FromStr`: : `impl FromStr for i32`
```rust,editable
// 为了使用 `from_str` 方法, 你需要引入该特征到当前作用域中
use std::str::FromStr;
fn main() {
    let parsed: i32 = "5".__.unwrap();
    let turbo_parsed = "10".__.unwrap();
    let from_str = __.unwrap();
    let sum = parsed + turbo_parsed + from_str;
    assert_eq!(sum, 35);

    println!("Success!")
}
```


3. 🌟🌟 还可以为自定义类型实现 `FromStr` 特征
```rust,editable
use std::str::FromStr;
use std::num::ParseIntError;

#[derive(Debug, PartialEq)]
struct Point {
    x: i32,
    y: i32
}

impl FromStr for Point {
    type Err = ParseIntError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let coords: Vec<&str> = s.trim_matches(|p| p == '(' || p == ')' )
                                 .split(',')
                                 .map(|x| x.trim())
                                 .collect();

        let x_fromstr = coords[0].parse::<i32>()?;
        let y_fromstr = coords[1].parse::<i32>()?;

        Ok(Point { x: x_fromstr, y: y_fromstr })
    }
}
fn main() {
    // 使用两种方式填空
    // 不要修改其它地方的代码
    let p = __;
    assert_eq!(p.unwrap(), Point{ x: 3, y: 4} );

    println!("Success!")
}
```

### Deref 特征
Deref 特征在[智能指针 - Deref](https://practice.rs/smart-pointers/deref.html)章节中有更加详细的介绍。

### transmute
`std::mem::transmute` 是一个 unsafe 函数，可以把一个类型按位解释为另一个类型，其中这两个类型必须有同样的位数( bits )。

`transmute` 相当于将一个类型按位移动到另一个类型，它会将源值的所有位拷贝到目标值中，然后遗忘源值。该函数跟 C 语言中的 `memcpy` 函数类似。

正因为此，**`transmute` 非常非常不安全!** 调用者必须要自己保证代码的安全性，当然这也是 unsafe 的目的。

#### 示例
1. `transmute` 可以将一个指针转换成一个函数指针，该转换并不具备可移植性，原因是在不同机器上，函数指针和数据指针可能有不同的位数( size )。

```rust,editable
fn foo() -> i32 {
    0
}

fn main() {
    let pointer = foo as *const ();
    let function = unsafe {
        std::mem::transmute::<*const (), fn() -> i32>(pointer)
    };
    assert_eq!(function(), 0);
}
```

2. `transmute` 还可以扩展或缩短一个不变量的生命周期，将 Unsafe Rust 的不安全性体现的淋漓尽致!
```rust,editable
struct R<'a>(&'a i32);
unsafe fn extend_lifetime<'b>(r: R<'b>) -> R<'static> {
    std::mem::transmute::<R<'b>, R<'static>>(r)
}

unsafe fn shorten_invariant_lifetime<'b, 'c>(r: &'b mut R<'static>)
                                             -> &'b mut R<'c> {
    std::mem::transmute::<&'b mut R<'static>, &'b mut R<'c>>(r)
}
```

3. 事实上我们还可以使用一些安全的方法来替代 `transmute`.
```rust,editable
fn main() {
    /*Turning raw bytes(&[u8]) to u32, f64, etc.: */
    let raw_bytes = [0x78, 0x56, 0x34, 0x12];

    let num = unsafe { std::mem::transmute::<[u8; 4], u32>(raw_bytes) };

    // use `u32::from_ne_bytes` instead
    let num = u32::from_ne_bytes(raw_bytes);
    // or use `u32::from_le_bytes` or `u32::from_be_bytes` to specify the endianness
    let num = u32::from_le_bytes(raw_bytes);
    assert_eq!(num, 0x12345678);
    let num = u32::from_be_bytes(raw_bytes);
    assert_eq!(num, 0x78563412);

    /*Turning a pointer into a usize: */
    let ptr = &0;
    let ptr_num_transmute = unsafe { std::mem::transmute::<&i32, usize>(ptr) };

    // Use an `as` cast instead
    let ptr_num_cast = ptr as *const i32 as usize;

    /*Turning an &mut T into an &mut U: */
    let ptr = &mut 0;
    let val_transmuted = unsafe { std::mem::transmute::<&mut i32, &mut u32>(ptr) };

    // Now, put together `as` and reborrowing - note the chaining of `as`
    // `as` is not transitive
    let val_casts = unsafe { &mut *(ptr as *mut i32 as *mut u32) };

    /*Turning an &str into a &[u8]: */
    // this is not a good way to do this.
    let slice = unsafe { std::mem::transmute::<&str, &[u8]>("Rust") };
    assert_eq!(slice, &[82, 117, 115, 116]);

    // You could use `str::as_bytes`
    let slice = "Rust".as_bytes();
    assert_eq!(slice, &[82, 117, 115, 116]);

    // Or, just use a byte string, if you have control over the string
    // literal
    assert_eq!(b"Rust", &[82, 117, 115, 116]);
}
```

> 你可以在[这里](https://github.com/sunface/rust-by-practice/blob/master/solutions/type-conversions/others.md)找到答案(在 solutions 路径下) 
