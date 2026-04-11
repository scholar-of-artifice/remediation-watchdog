# Initializing Basic Projects

In this article, you will learn how I initialized the basic sub-projects that are part of this larger system.

## Creating the Directories

```bash
mkdir absurd-iguana bashful-yak calm-lynx dazzling-remora eager-marmot fearless-eagle giant-wasp
```

`mkdir` creates a separate directory where the code for each specific service will live.

## Initialize a Go Module in Each Directory

```bash
cd absurd-iguana
go mod init absurd-iguana
cd ..

# and so on and so forth...
```

The `go mod init <module-name-here>` command creates a `go.mod` file in the current directory. This file tracks the dependencies specifically for that service, keeping them isolated from the others.

## Initialize a Go Workspace

In the root directory of the project:

```bash
go work init
```

This generates a go.work file at the root of the project. A workspace allows yout to work across multiple modules simultaneously. It tells the IDE and toolchain to resolve local dependencies and intellisense acorss the entire workspace without needing to publish the modules or replace directives in the `go.mod` files.

## Add the Module sto the Workspace

In the root directory of the project:

```bash
go work use ./absurd-iguana ./bashful-yak ./calm-lynx ./dazzling-remora ./eager-marmot ./fearless-eagle ./giant-wasp 
```

Once this command is executes... the parts of the architecture will be initialized. Most IDEs will correctly recognize and support all independent Go services.


