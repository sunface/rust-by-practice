# Lifetime


**Lifetimes**

1. 🌟
```rust,editable
/* Make it work */

#[derive(Debug)]
struct NoCopyType {}

#[derive(Debug)]
struct Example<'a, 'b> {
    a: &'a u32,
    b: &'b NoCopyType
}

fn main()
{ 
  /* 'a tied to fn-main stackframe */
  let var_a = 35;
  let example: Example;
  
  // {
    /* lifetime 'b tied to new stackframe/scope */ 
    let var_b = NoCopyType {};
    
    /* fixme */
    example = Example { a: &var_a, b: &var_b };
  // }
  
  println!("(Success!) {:?}", example);
}
```


2. 🌟
```rust,editable

#[derive(Debug)]
struct NoCopyType {}

#[derive(Debug)]
#[allow(dead_code)]
struct Example<'a, 'b> {
    a: &'a u32,
    b: &'b NoCopyType
}

/* Fix function signature */
fn fix_me<'b>(foo: &Example<'_, 'b>) -> &'b NoCopyType
{ foo.b }

fn main()
{
    let no_copy = NoCopyType {};
    let example = Example { a: &1, b: &no_copy };
    fix_me(&example);
    print!("Success!")
}
```




