# 2022 Dcard Web Frontend Intern Homework

Using Go with Fiber and SQLite. This is also my first project regarding Go and SQL databases.

## How to start

Before starting this program, you can configure the hostname and port where it will run at by editing `config.json`. There is also an option that decides whether SQLite uses an in memory database or a `urls.db` file on your disk drive.

To start without building it, you can type

```bash
go run .
```

Alternatively, you can build and run this program by

```bash
go build
./dcard-2022-backend-intern
```

## Choice Explanation

### Why Go?

Although I'm good at JavaScript and Express, I want to learn something new and prove that I can learn thing fast. What's more, Go is faster than JS, and has a better way handling concurrency than C/C++. As a result, server built with Go is definitely faster than JavaScript and Express.

### Why Fiber?

Fiber is a blazing fast web framework inspired by Express, the famous JS framework I'm familiar with.
In [TechEmpower Web Framework Benchmarks](https://www.techempower.com/benchmarks/#section=data-r20&hw=cl&test=composite&l=zijo5b-sf), Fiber won 14th place with the overall score of 830 (cloud environment). Moreover, it has a huge community support and got the most stars on GitHub among the competition.

### Why SQLite?

SQLite is a convenient RDBMS that follows ACID. It doesn't need a client server model, has a tiny size, and is simple.
