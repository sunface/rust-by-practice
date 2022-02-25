# 所有权

🌟🌟
```rust,editable

fn main() {
    // 只能修改下面这行代码！ 使用尽可能多的方法来通过编译
    let x = String::from("hello, world");
    let y = x;
    println!("{},{}",x,y);
}
```

🌟🌟
```rust,editable
// 不要修改 main 中的代码
fn main() {
    let s1 = String::from("hello, world");
    let s2 = take_ownership(s1);

    println!("{}", s2);
}

// 只能修改下面的代码!
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

// 只能修改下面的代码!
fn give_ownership() -> String {
    let s = String::from("hello, world");
    // convert String to Vec
    // 将 String 转换成 Vec 类型
    let _s = s.into_bytes();
    s
}
```

🌟🌟
```rust,editable
// 使用 clone 来让代码通过编译
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
// 不要使用 clone，使用 copy 的方式替代
fn main() {
    let x = (1, 2, (), "hello");
    let y = x.clone();
    println!("{:?}, {:?}", x, y);
}
```

