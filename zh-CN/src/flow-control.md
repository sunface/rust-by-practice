# 流程控制

### if/else
🌟 
```rust,editable

// 填空
fn main() {
    let n = 5;

    if n < 0 {
        println!("{} is negative", n);
    } __ n > 0 {
        println!("{} is positive", n);
    } __ {
        println!("{} is zero", n);
    }
} 
```

🌟🌟 if/else 可以用作表达式来进行赋值
```rust,editable

// 修复错误
fn main() {
    let n = 5;

    let big_n =
        if n < 10 && n > -10 {
            println!(", and is a small number, increase ten-fold");

            10 * n
        } else {
            println!(", and is a big number, halve the number");

            n / 2.0 ;
        }

    println!("{} -> {}", n, big_n);
} 
```

### for
🌟 The `for in` construct can be used to iterate through an Iterator, e.g a range `a..b`.

```rust,editable

fn main() {
    for n in 1..=100 { // modify this line to make the code work
        if n == 100 {
            panic!("NEVER LET THIS RUN")
        }
    }
} 
```


🌟🌟 
```rust,editable

// 修复错误，不要新增或删除代码行
fn main() {
    let names = [String::from("liming"),String::from("hanmeimei")];
    for name in names {
        // do something with name...
    }

    println!("{:?}", names);

    let numbers = [1, 2, 3];
    // numbers中的元素实现了 Copy，因此无需转移所有权
    for n in numbers {
        // do something with name...
    }
    
    println!("{:?}", numbers);
} 
```

🌟
```rust,editable
fn main() {
    let a = [4,3,2,1];

    // iterate the indexing and value in 'a'
    for (i,v) in a.__ {
        println!("第{}个元素是{}",i+1,v);
    }
}
```

### while
The `while` keyword can be used to run a loop when a condition is true.

```rust,editable

// fill in the blanks to make the last println! work !
fn main() {
    // A counter variable
    let mut n = 1;

    // Loop while the condition is true
    while n __ 10 {
        if n % 15 == 0 {
            println!("fizzbuzz");
        } else if n % 3 == 0 {
            println!("fizz");
        } else if n % 5 == 0 {
            println!("buzz");
        } else {
            println!("{}", n);
        }


        __;
    }

    println!("n reached {}, soloop is over",n);
}
```

### continue and break
🌟 使用 `break` 可以跳出循环
```rust,editable

// 填空，不要修改其它代码
fn main() {
    let mut n = 0;
    for i in 0..=100 {
       if n == 66 {
           __
       }
       n += 1;
    }

    assert_eq!(n, 66);
}
```

🌟🌟 `continue` 会结束当次循环并立即开始下一次循环
```rust,editable

// 填空，不要修改其它代码
fn main() {
    let mut n = 0;
    for i in 0..=100 {
       if n != 66 {
           n+=1;
           __;
       }
       
       __
    }

    assert_eq!(n, 66);
}
```

### loop 

🌟🌟 loop 一般都需要配合 `break` 或 `continue` 一起使用。

```rust,editable

// 填空，不要修改其它代码
fn main() {
    let mut count = 0u32;

    println!("Let's count until infinity!");

    // Infinite loop
    loop {
        count += 1;

        if count == 3 {
            println!("three");

            // Skip the rest of this iteration
            __;
        }

        println!("{}", count);

        if count == 5 {
            println!("OK, that's enough");

            __;
        }
    }

    assert_eq!(count, 5);
}
```

🌟🌟 loop is an expression, so we can use it with `break` to return a value
```rust,editable

// fill in the blank
fn main() {
    let mut counter = 0;

    let result = loop {
        counter += 1;

        if counter == 10 {
            __;
        }
    };

    assert_eq!(result, 20);
}
```

🌟🌟🌟 It's possible to break or continue outer loops when dealing with nested loops. In these cases, the loops must be annotated with some 'label, and the label must be passed to the break/continue statement.

```rust,editable

// 填空
fn main() {
    let mut count = 0;
    'outer: loop {
        'inner1: loop {
            if count >= 20 {
                // This would break only the inner1 loop
                break 'inner1; // `break` is also ok 
            }
            count += 2;
        }

        count += 5;

        'inner2: loop {
            if count >= 30 {
                // This breaks the outer loop
                break 'outer;
            }

            // This will continue the outer loop
            continue 'outer;
        }
    }

    assert!(count == __)
}
```