# 语句与表达式

### 示例
```rust,editable
fn main() {
    let x = 5u32;

    let y = {
        let x_squared = x * x;
        let x_cube = x_squared * x;

        // 下面表达式的值将被赋给 `y`
        x_cube + x_squared + x
    };

    let z = {
        // 分号让表达式变成了语句，因此返回的不再是表达式 `2 * x` 的值，而是语句的值 `()`
        2 * x;
    };

    println!("x is {:?}", x);
    println!("y is {:?}", y);
    println!("z is {:?}", z);
}
```

### 练习
🌟🌟
```rust,editable
// 使用两种方法修改内部的 {} 中的内容
fn main() {
   let v = {
       let mut x = 1;
       x += 2
   };

   assert_eq!(v, 3);
}
```

🌟
```rust,editable

fn main() {
   let v = (let x = 3);

   assert!(v == 3);
}
```

🌟
```rust,editable

fn main() {}

fn sum(x: i32, y: i32) -> i32 {
    x + y;
}
```