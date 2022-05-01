# 数值类型

### 整数

1. 🌟 

> Tips: 如果我们没有显式的给予变量一个类型，那编译器会自动帮我们推导一个类型

```rust,editable

// 移除某个部分让代码工作
fn main() {
    let x: i32 = 5;
    let mut y: u32 = 5;

    y = x;
    
    let z = 10; // 这里 z 的类型是? 
}
```

2. 🌟 
```rust,editable

// 填空
fn main() {
    let v: u16 = 38_u8 as __;
}
```

3. 🌟🌟🌟 

> Tips: 如果我们没有显式的给予变量一个类型，那编译器会自动帮我们推导一个类型

```rust,editable

//  修改 `assert_eq!` 让代码工作
fn main() {
    let x = 5;
    assert_eq!("u32".to_string(), type_of(&x));
}

// 以下函数可以获取传入参数的类型，并返回类型的字符串形式，例如  "i8", "u8", "i32", "u32"
fn type_of<T>(_: &T) -> String {
    format!("{}", std::any::type_name::<T>())
}
```

4. 🌟🌟 
```rust,editable

// 填空，让代码工作
fn main() {
    assert_eq!(i8::MAX, __); 
    assert_eq!(u8::MAX, __); 
}
```

5. 🌟🌟 
```rust,editable

// 解决代码中的错误和 `panic`
fn main() {
   let v1 = 251_u8 + 8;
   let v2 = i8::checked_add(251, 8).unwrap();
   println!("{},{}",v1,v2);
}
```

6. 🌟🌟
```rust,editable

// 修改 `assert!` 让代码工作
fn main() {
    let v = 1_024 + 0xff + 0o77 + 0b1111_1111;
    assert!(v == 1579);
}
```


### 浮点数
7. 🌟 

```rust,editable

// 将 ? 替换成你的答案
fn main() {
    let x = 1_000.000_1; // ?
    let y: f32 = 0.12; // f32
    let z = 0.01_f64; // f64
}
```
8. 🌟🌟 使用两种方法来让下面代码工作


```rust,editable

fn main() {
    assert!(0.1+0.2==0.3);
}
```

### 序列Range
9. 🌟🌟 两个目标: 1. 修改 `assert!` 让它工作 2. 让 `println!` 输出: 97 - 122

```rust,editable
fn main() {
    let mut sum = 0;
    for i in -3..2 {
        sum += i
    }

    assert!(sum == -3);

    for c in 'a'..='z' {
        println!("{}",c);
    }
}
```

10. 🌟🌟 
```rust,editable

// 填空
use std::ops::{Range, RangeInclusive};
fn main() {
    assert_eq!((1..__), Range{ start: 1, end: 5 });
    assert_eq!((1..__), RangeInclusive::new(1, 5));
}
```

### 计算

11. 🌟 
```rust,editable

// 填空，并解决错误
fn main() {
    // 整数加法
    assert!(1u32 + 2 == __);

    // 整数减法
    assert!(1i32 - 2 == __);
    assert!(1u8 - 2 == -1);
    
    assert!(3 * 50 == __);

    assert!(9.6 / 3.2 == 3.0); // error ! 修改它让代码工作

    assert!(24 % 5 == __);
    
    // 逻辑与或非操作
    assert!(true && false == __);
    assert!(true || false == __);
    assert!(!true == __);

    // 位操作
    println!("0011 AND 0101 is {:04b}", 0b0011u32 & 0b0101);
    println!("0011 OR 0101 is {:04b}", 0b0011u32 | 0b0101);
    println!("0011 XOR 0101 is {:04b}", 0b0011u32 ^ 0b0101);
    println!("1 << 5 is {}", 1u32 << 5);
    println!("0x80 >> 2 is 0x{:x}", 0x80u32 >> 2);
}
```

> 你可以在[这里](https://github.com/sunface/rust-by-practice/blob/master/solutions/basic-types/numbers.md)找到答案(在 solutions 路径下) 
