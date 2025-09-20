# Pocket-Sized (Go) Projects 🤭

This is my attempt at tackling, (re)creating, and building the projects found in the excellent [Learn Go with Pocket-Sized Projects](https://www.manning.com/books/learn-go-with-pocket-sized-projects) book.

## Projects 🗂️

- [x] [Hello, earth!](./hello)
- [x] [A bookworm's digest: Playing with loops and maps](./bookworms)
- [x] A log story: Creating a library
- [ ] Gordle: Play a word game in your terminal
- [ ] Money Converter: CLI around an HTTP call

## Usage 🔧

Navigate to the project folder and run or test the code:

```bash
$ cd hello        # Go into the project directory
$ go run main.go  # Run the project
$ go test         # Run the tests (if any)
```

> **Tip**: Running `go run .` will compile and execute the current project. Handy for quick experiments! 🚀

## Using go.work for Multi-Project Management 🛠️

[Go workspaces](https://go.dev/doc/tutorial/workspaces) (go.work) let you manage multiple Go modules together,
which is perfect for juggling all your `pocket-sized` projects.

### Initialize a workspace

From the root folder containing all your projects:

```bash
$ go work init ./hello ./bookworms
```

This creates a `go.work` file linking your projects.

### Add a new project to the workspace

If you create a new project, for example `calculator`:

```bash
$ go work use ./calculator
```

Your `go.work` file now includes it automatically.

### Running projects in the workspace

You can run any project as usual from its folder:

```bash
$ cd hello
$ go run .
```

Or run from the workspace root using the relative path:

```bash
$ go run ./hello
```

### Running tests in the workspace

You can test a single project:

```bash
$ cd hello
$ go test
```

## License 📄

These projects are open-source and licensed under the [MIT License](./LICENSE) 😉

Happy hacking! 🤩
