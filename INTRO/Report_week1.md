# **INTRODUCTION TO GOLANG**

- [**INTRODUCTION TO GOLANG**](#introduction-to-golang)
  - [**Environment variables in go**](#environment-variables-in-go)
  - [**Build \& Run go files**](#build--run-go-files)
  - [**Module, package in go**](#module-package-in-go)
    - [**Package**](#package)
    - [**Module**](#module)

## **Environment variables in go**

Environment variables are used in building execution files from code. These are:

- `GOOS`: The target operation system of the target device: `linux`, `windows`,...

- `GOARCH`: The architecture of the target device: `arm64`, `amd64`

- `CGO_ENABLE`: `=1`: The built file can call the libraries in C right in the device when running. This makes the program dependent on the C libraries in the device. So, set `=0` to build all the dependencies, so that the final program can run without the fully equipped libraries in the target device.

- `GOPROXY`: Go downloads the libraries directly from proxy.golang.org by default. However, in companies, all the access to the Internet must be manage by proxy. Therefore, `GOPROXY` can help with connecting to the proxy, then, proxy connects to proxy.golang.org to download libraries.

## **Build & Run go files**

Set the environment variables first:

```bash
set GOOS=linux
set GOARCH=amd64
```

Build:

```bash
go build -o <final_execution_app_name> <go_code_file_name>.go
```

--> Built file is a file without extension, this file can run in the target computer.

Click into the built file to run, or build and run immediately by:

```bash
go run <go_code_file_name>.go
```

## **Module, package in go**

### **Package**

*Package* is a physical folder containing 1 or many files `.go`. This is used for grouping all the files with the same logical purpose. Otherwise, it also helps with managing functions'name.

All files must belong to a particular package: At the beginning of every go file, there must be a line declaring the package: `package <folder_name>`

The functions with the uppercase name are public with (can be called by) the functions in other packages. Otherwise, they can only be called and used inside a package.

Import:

```go
import (
    "<packageA>"
    "<path/to/packageB>"
)
```

--> The `"path/to/packageB"` is imported as an alias `packageB` automatically. When use functions in this package, directly use: `packageB.function`

If there are 2 packages with the same alias: `"module1/packageA"` and `"module2/packageA"`: import by using explicit alias

```go
import (
    AliasName1 "<module1/packageA>"
    AliasName2 "<module2/packageA>"
)
```

When use a function in the package: `AliasName1.function()`

### **Module**

*Module* is a repository including 1 or many packages. This plays a similar role as virtual environment in Python: managing the version of libraries.

Libraries and dependencies with different versions are stored in the same pool. Each project creates a module to manage the dependencies by command `go mod init <name_of_module>`. Then, a file `go.mod` will be created in the root folder of the project, this includes references to the necessary dependencies in the pool, helping with the storage by avoiding download another copy of library like venv in Python.

To use module:

- Create module: `go mod init <name_of_module>`
- Download module from Internet: `go get <link_to_module>`
  
    --> This will download the available module from the Internet into the pool, then, add a reference to the library into file `go.mod` of project
- Import packages in module:
  
    ```go
    import (
        "<module1/package1>"
        "<module2/package2>"
    )
    ```
