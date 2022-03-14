1. `cargo new hello-package`

2. `cargo new --lib hello-package1`

3. `hello-package` has a binary crate named `hello-package`, `src/main.rs` is the crate root. 
   
`hello-pacakge1` has a library crate named `hello-package1`, `src/lib.rs` is the crate root.

4. `hello-package1`

5.
```shell
# FILL in the blanks
.
├── Cargo.lock
├── Cargo.toml
├── src
│   ├── main.rs
│   └── lib.rs
```

6.
```shell
# Create a package which contains 
# 1. three binary crates: `hello-package`, `main1` and `main2`
# 2. one library crate
# describe the directory tree below
.
├── Cargo.toml
├── Cargo.lock
├── src
│   ├── main.rs
│   ├── lib.rs
│   └── bin
│       └── main1.rs
│       └── main2.rs
├── tests # directory for integrated tests files
│   └── some_integration_tests.rs
├── benches # dir for benchmark files
│   └── simple_bench.rs
└── examples # dir for example files
    └── simple_example.rs
```