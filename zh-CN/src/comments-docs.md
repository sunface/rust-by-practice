# 注释和文档
本章的学习资料在[这里](https://course.rs/basic/comment.html)，大家可以先行学习后再来做题。

## 注释

1. 🌟🌟
```rust,editable

/* 只使用注释让下面代码工作! */
fn main() {
    todo!();
    unimplemented!();

    assert_eq!(6, 5 + 3 + 2 + 1 )
}
```


## 文档注释
文档注释会被解析为 HTML 文件，并支持 `Markdown` 语法。

在开始之前，我们需要创建一个新的项目用于后面的练习: `cargo new --lib doc-comments`.


### 行文档注释 `///` 
为 `add_one` 函数添加文档
```rust
// in lib.rs

/// Add one to the given value and return the value
///
/// # Examples
///
/// ```
/// let arg = 5;
/// let answer = my_crate::add_one(arg);
///
/// assert_eq!(6, answer);
/// ```
pub fn add_one(x: i32) -> i32 {
    x + 1
}
```

### Cargo doc
我们可以使用 `cargo doc --open` 来生成 HTML 文件，并自动在浏览器中打开网页。

### 块文档注释 `/** ... */`
为函数 `add_two` 添加文档:
```rust
/** Add two to the given value and return a new value

# Examples

let arg = 5;
let answer = my_crate::add_two(arg);

assert_eq!(7, answer);

*/
pub fn add_two(x: i32) -> i32 {
    x + 2
}
```

### 为包和模块创建文档注释
我们还可以创建包和模块的注释，用于描述它们的功能。

首先，来为我们的库包添加一些文档注释:

> 注意: 必须要将包、模块注释放置在包根或模块文件的最顶部

```rust
//! # 文档注释
//! 
//! 该库用于文档注释的教学

// in lib.rs
pub mod compute;
```

同样的，我们还可以使用块注释来达成目的:
```rust
/*! # 文档注释

 该库用于文档注释的教学 */
```

下一步，创建一个新的模块文件 `src/compute.rs`, 然后在其中添加以下注释:
```rust
//! 本模块用于处理一些复杂计算

// in compute.rs
```

然后运行 `cargo doc --open` 查看下结果。

### 文档测试
细心的同学可能会发现之前的 `add_one` 和 `add_tow` 的文档注释中，包含了两个示例代码块.

以上示例不仅仅是作为文档用于演示你的函数该如何使用，它的另一个作用就是用于文档测试 `cargo test`。

2. 🌟🌟 但是在这两个函数的示例中，存在错误，请修复它们并使用 `cargo test` 获取以下输出结果: 
```shell
running 0 tests

test result: ok. 0 passed; 0 failed; 0 ignored; 0 measured; 0 filtered out; finished in 0.00s

   Doc-tests doc-comments

running 2 tests
test src/lib.rs - add_one (line 11) ... ok
test src/lib.rs - add_two (line 26) ... ok

test result: ok. 2 passed; 0 failed; 0 ignored; 0 measured; 0 filtered out; finished in 0.55s
```

3. 🌟🌟 有时我们会期望示例的结果是一个 panic。将以下代码添加到 `src/compute.rs` ，并且让  `cargo test` 成功运行.

> 你只能修改注释，不要修改 `fn div`

```rust
// in src/compute.rs

/// # Panics
///
/// The function panics if the second argument is zero.
///
/// ```rust,should_panic
/// // panics on division by zero
/// doc_comments::compute::div(10, 0);
/// ```
pub fn div(a: i32, b: i32) -> i32 {
    if b == 0 {
        panic!("Divide-by-zero error");
    }

    a / b
}
```

4. 🌟🌟 有时我们会想要隐藏文档，但是保留文档测试

将以下代码添加到 `src/compute.rs` ,

```rust
// in src/compute.rs

/// ```
/// # fn try_main() -> Result<(), String> {
/// let res = doc_comments::compute::try_div(10, 0)?;
/// # Ok(()) // returning from try_main
/// # }
/// # fn main() { 
/// #    try_main().unwrap();
/// #
/// # }
/// ```
pub fn try_div(a: i32, b: i32) -> Result<i32, String> {
    if b == 0 {
        Err(String::from("Divide-by-zero"))
    } else {
        Ok(a / b)
    }
}
```

然后修改以上代码已实现两个目标:

- 文档注释不能出现在 `cargo doc --open` 生成的网页中
- 运行测试，并成功看到以下结果:

```shell
running 0 tests

test result: ok. 0 passed; 0 failed; 0 ignored; 0 measured; 0 filtered out; finished in 0.00s

   Doc-tests doc-comments

running 4 tests
test src/compute.rs - compute::div (line 7) ... ok
test src/lib.rs - add_two (line 27) ... ok
test src/lib.rs - add_one (line 11) ... ok
test src/compute.rs - compute::try_div (line 20) ... ok

test result: ok. 4 passed; 0 failed; 0 ignored; 0 measured; 0 filtered out; finished in 0.51s
```

### 代码跳转
Rust 为我们提供一个非常强大的特性：可以在文档注释中实现代码跳转。

将以下代码添加到 `src/lib.rs`:
```rust
// in lib.rs

/// Add one to the given value and return a [`Option`] type
pub fn add_three(x: i32) -> Option<i32> {
    Some(x + 3)
}
```

除了跳转到标准库中，我们还能跳转到项目中的其它模块。

```rust
// in lib.rs

mod a {
    /// Add four to the given value and return a [`Option`] type
    /// [`crate::MySpecialFormatter`]
    pub fn add_four(x: i32) -> Option<i32> {
        Some(x + 4)
    }
}

struct MySpecialFormatter;
```

### 文档属性
下面是很常用的 `#[doc]` 属性，该属性可以被 `rustdoc` 所使用。


### `inline`

可以用于内联文档, 而不是链接到一个单独的页面。

```rust,ignore
#[doc(inline)]
pub use bar::Bar;

/// bar docs
mod bar {
    /// the docs for Bar
    pub struct Bar;
}
```

### `no_inline`
用于防止链接到单独的页面或其它地方。

```rust,ignore
// Example from libcore/prelude
#[doc(no_inline)]
pub use crate::mem::drop;
```

### `hidden`

通过这个属性让 `rustdoc` 不要将下面的项包含在文档中:

```rust,editable,ignore
// Example from the futures-rs library
#[doc(hidden)]
pub use self::async_await::*;
```

对文档来说，`rustdoc` 被社区广泛采用，大家所看到的[标准库文档](https://doc.rust-lang.org/std/)也是基于此生成的。


### 完整的代码
`doc-comments` 的完整代码可以在[这里找到](https://github.com/sunface/rust-by-practice/tree/master/practices/doc-comments).
