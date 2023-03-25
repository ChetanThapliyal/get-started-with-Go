# Introduction

Go, or Golang, is an open source programming language developed by Google in 2009. It has a simple syntax and is designed for easy readability, making it a great language for beginners. 

Go is also known for its scalability and fast compile times, making it a great language for large scale projects. It is used in a variety of industries, such as web development, cloud computing, and data science. Go has a thriving open source community, with many libraries and frameworks available to make development faster and easier.

## Key GO Features

- A **compiled** language (Fast compilation)
- Less cumbersome code
- **Statically typed**
- With a built-in **concurrency** a system that is easy for developers to work on
- With a robust dependency management
- With a **garbage collector** (Unused memory freed automatically)

Go's syntax is similar to C, with some elements borrowed from other languages like Python and C++.

Go's standard library provides a wide range of functionality, including HTTP support, encryption, compression, and other commonly used features. Additionally, there is a large ecosystem of third-party libraries and frameworks available for Go.

Some of the benefits of Go include its simplicity, speed, and scalability. The language is easy to learn, even for developers new to programming. Go's built-in concurrency support makes it easy to write concurrent and parallel code, which can improve the performance of your applications.

> ### 📌 Concepts
> - **Concurrency** : A program is concurrent when tasks **can** be executed out-of-order or in a partial order. Concurrency in programming language is the ability of a program to run multiple tasks simultaneously. This is achieved by splitting tasks into sub-tasks that can be executed independently, allowing them to be processed in parallel rather than sequentially.
> - **Garbage collector** (often called GC): When we build programs, we need to store data and fetch data from memory. Memory is not an infinite resource. GC is a process in programming where unused elements stored in memory are destroyed from time to time. Putting some data into memory is called allocation; the inverse action, which consists of removing data from memory, is called deallocation. The garbage collector’s role is to deallocate memory when it is not used anymore.
> - **Statically Typed** : A statically typed language is a programming language where variables are explicitly declared before the program is executed. This means that the type of data that a variable can contain is determined at compile-time, and can not change during the program's execution. This provides a level of clarity and structure to the code, but also restricts the flexibility of the language in certain ways.

## Go Compared to Python and C++
---

| Go | Python | C++ |
| --- | --- | --- |
| Statically typed | Dynamically typed | Statically typed |
| Fast run time | Slow run time | Fast run time |
| Compiled | Interpreted | Compiled |
| Fast compile time | Interpreted | Slow compile time |
| Supports concurrency through goroutines and channel | No built-in concurrency mechanism | Supports concurrency through threads |
| Has automatic garbage collection | Has automatic garbage collection | Does not have automatic garbage collection |
| Does not support classes and objects | Has classes and objects | Has classes and objects |
| Does not support inheritance | Supports inheritance | Supports inheritance |
---

## Static vs Dynamic Typed Languages

| Static Type Languages | Dynamic Type Languages |
| --- | --- |
| 1. Static typed languages require that the data type of a variable be declared explicitly at the time of its declaration. The type of a variable is then checked at compile time, and any type mismatch will result in a compile-time error. This means that any errors related to type mismatches are caught before the program is run, making it easier to find and fix bugs. | 1. Dynamic typed languages do not require explicit type declaration. Instead, the type of a variable is determined at runtime, and the same variable can hold values of different types at different times. This means that type mismatches can only be detected at runtime, and the program might produce unexpected results or errors. |
| 2. Code written in a statically typed language is usually more concise. | 2. Code written in a dynamically typed language is usually more verbose. |
| 3. Statically typed languages tend to be faster and more efficient than dynamically typed languages. | 3. Dynamically typed languages can be more flexible and expressive than statically typed languages. |
| 4. Examples: Java, C++, and Go. | 4. Examples: Python, JavaScript, and Ruby. |
---
