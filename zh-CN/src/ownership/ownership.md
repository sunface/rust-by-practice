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

#### 可变性
当所有权转移时，可变性也可以随之改变。

🌟
```rust,editable

fn main() {
    let s = String::from("hello, ");
    
    // 只修改下面这行代码 !
    let s1 = s;

    s1.push_str("world")
}
```

🌟🌟🌟
```rust,editable

fn main() {
    let x = Box::new(5);
    
    let ...      // 完成该行代码，不要修改其它行！
    
    *y = 4;
    
    assert_eq!(*x, 5);
}
```

### 部分 move
当解构一个变量时，可以同时使用 `move` 和引用模式绑定的方式。当这么做时，部分 `move` 就会发生：变量中一部分的所有权被转移给其它变量，而另一部分我们获取了它的引用。

在这种情况下，原变量将无法再被使用，但是它没有转移所有权的那一部分依然可以使用，也就是之前被引用的那部分。

#### 示例
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

    // 通过这种解构式模式匹配，person.name 的所有权被转移给新的变量 `name`
    // 但是，这里 `age` 变量确是对 person.age 的引用, 这里 ref 的使用相当于: let age = &person.age 
    let Person { name, ref age } = person;

    println!("The person's age is {}", age);

    println!("The person's name is {}", name);

    // Error! 原因是 person 的一部分已经被转移了所有权，因此我们无法再使用它
    //println!("The person struct is {:?}", person);

    // 虽然 `person` 作为一个整体无法再被使用，但是 `person.age` 依然可以使用
    println!("The person's age from person struct is {}", person.age);
}
```

#### 练习

🌟
```rust,editable

fn main() {
   let t = (String::from("hello"), String::from("world"));

   let _s = t.0;

   // 仅修改下面这行代码，且不要使用 `_s`
   println!("{:?}", t);
}
```

🌟🌟
```rust,editable

fn main() {
   let t = (String::from("hello"), String::from("world"));

   // 填空，不要修改其它代码
   let (__, __) = t;

   println!("{:?}, {:?}, {:?}", s1, s2, t);
}
```