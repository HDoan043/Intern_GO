# **BASIC GO**

**TABLE OF CONTENT**

- [**1. VARIABLES**](#1-variables)
- [**2. CONTROLLING STRUCTURE**](#2-controlling-structure)
  - [***Branch***](#branch)
  - [***Loop***](#loop)
- [**3. DATA STRUCTURE**](#3-data-structure)
  - [**3.1| ARRAY**](#31-array)
    - [***Define and initialize value***](#define-and-initialize-value)
    - [***Copy array***](#copy-array)
    - [***Access to an element***](#access-to-an-element)
  - [**3.2| STRUCT**](#32-struct)
    - [***Define new data type***](#define-new-data-type)
    - [***Define new object with the custom type***](#define-new-object-with-the-custom-type)
    - [***Initialize value***](#initialize-value)
  - [**3.3| SLICE**](#33-slice)
    - [***Define and initialize value***](#define-and-initialize-value-1)
    - [***Access and append elements***](#access-and-append-elements)
    - [***Copy slice***](#copy-slice)
  - [**3.4| MAP**](#34-map)
    - [***Define and initialize map***](#define-and-initialize-map)
    - [***Add/change element in map***](#addchange-element-in-map)
    - [***Get an element***](#get-an-element)
    - [***Delete an element***](#delete-an-element)
- [**4. FUNCTION**](#4-function)
  - [***Global function***](#global-function)
  - [***Local function***](#local-function)
  - [***Anonymous function***](#anonymous-function)
  - [***Function of struct***](#function-of-struct)
- [**5. ERROR PROCESSING**](#5-error-processing)
  - [**Key word `defer`**](#key-word-defer)
  - [**Catch errors**](#catch-errors)
- [**6. CONCURRENCE PROCESSING**](#6-concurrence-processing)
  - [**6.1 Concurrence processing with goroutines**](#61-concurrence-processing-with-goroutines)
  - [**6.2 Communication among routines**](#62-communication-among-routines)
    - [***Transporting many messages***](#transporting-many-messages)
    - [***Many goroutines at the same time***](#many-goroutines-at-the-same-time)

## **1. VARIABLES**

Define variables: global and local (inside function)

- Global: using the key word `var` and must care about its type: `var <var_name> <type>`.

    If combine define and initialize the value for variables: `var <var_name> <type> = <value>`

    Can define global variables separately or in group:

    ```go
    package main

    // Define separately
    var globalCount int = 0
    var globalName string = "MyName"

    // Define in group
    var (
        gCount int = 0
        gName string = "MyName2"
    )
    ```

- Local: can explicitly define by using key word `var`, and initialize value later, or use operator `:=` to define and initialize the value at the same time, so that go can automatically interpret the type of variables.

    ```go
    function main() {
        //Explicitly define and initialize value
        var count int = 0
        var name string
        str = "LocalName"

        //Shortly define and initialize value 
        localCount := 0
        localName := "LocalName"
    }
    ```

    **NOTE**: The local variables defined ***MUST*** be used in code, if not, compiler cannot build successfully.

## **2. CONTROLLING STRUCTURE**

### ***Branch***

```go
if <condition> {
}
else if <condition> {
}
else {
}
```

### ***Loop***

In go, there is no `while`, just `for`.

- Looping index:  

    ```go
    // <=> (Python) for i in range(N):
    for i := 0; i < N; i++ {

    }
    ```

- Looping item of list:  

    ```go
    // <=> (Python) for i, item in enumerate(lst):
    for i, item := range lst {

    }
    ```

## **3. DATA STRUCTURE**

### **3.1| ARRAY**

*Array* is a data structure like array in C: continuous memory area in RAM, array contains elements with the same type; unchangeable size, we cannot append or remove any elements from array.

```txt
RAM

_________________________________
   | arr[0] | arr[1] | arr[2] |     
___|________|________|________|__
                             
   <------------arr-----------> 
```

#### ***Define and initialize value***

Similarly to variables, to define an array and initialize value:

- Explicitly define: `var <array_name> <type_of_array>`. `<type_of_array>` is not simply `int`, `float` which is the type of the elements in array like C, but must be `[<size>]<type_of_elements_in_array>` (`[3]int`, `[4]float`). This means the size of array is also a part of the type of array.

    The value of an array is `<type_of_array>{value1, value2,....}`. To initialize value:

    ```go
    var array [3]int
    array = [3]int{1,2,3}
    ```

- Shortly define and initialize value:

    ```go
    array := [3]int{1,2,3}
    array2:= [...]int{1,2,3,4} // go can automatically count elements and define array without explicitly determine size
    ```

#### ***Copy array***

When pass an array into a function, or use an available array to initialize value for a new array, the passed value and the initialized value of new array are just the copies of the original array.

```go
array := [...]int{1,2,3}
new_array := array // array in this expression is just the copy of the original array
```

#### ***Access to an element***

Get an element through its index similarly as in C:

```go
array := [...]int{1,2,3,4,5}
first_element := array[0] // first_element = 1
array[2] = 100 // array: [5]int{1,2,100,4,5}
```

### **3.2| STRUCT**

*Struct* is similar to class in other language. Struct allows to group the variables or functions into an object. The type of object is a special type defined by key word `struct`

#### ***Define new data type***

```go
type <Type_name> struct {
    <property1> <type1>
    <property2> <type2>
}
```

#### ***Define new object with the custom type***

```go
type NewDataType struct {
    ID int
    Name string
}

var new_object NewDataType
```

#### ***Initialize value***

The value of an object is:

```go
value = <Custom_type>{
    property1: value1,
    property2: value2,
}
```

Example:

```go
// Explicitly define and initialize value later
var new_object NewDataType
new_object = NewDataType{
    ID: 1,           
    Name: "Hung",
}

//Define and initialize
new_object := NewDataType{
    ID: 1,           
    Name: "Hung",
}
```

### **3.3| SLICE**

*Slice* in go is similar to the vector in C and list in Python: an array of elements with the same type, changeable size, we can append elements. In fact, slice is implemented by struct with 3 properties:

- Data: an array
- Length: the length of the current slice, accessing to the elements with index lower than this length is allowed and there is no error index out of bound
- Capacity: the capacity shows the total size of the used area and area in reserve which will be used in the future when the slice is broadened

```txt
Slice
____________________________________________
| slice[0] | slice[1] | not use  | not use |
|__________|__________|__________|_________|
    
<---Length(in use)----><-----(reserve)----->

<------------------Capacity---------------->

```

#### ***Define and initialize value***

Define and initialize value by:

- Function `make`. This function receives 3 parameters `[]<type_of_element_in_slice>`, `length`, `capacity`:

  ```go
  slice := make([]int, 0, 5)
  ```

- Operator `:=`. The value is `[]<type_of_element_in_slice>` (similarly to array but no size). (This is also created by `make` with `capacity` = `length`)

    ```go
    slice := []int{1,2,3} 
    ```

#### ***Access and append elements***

An element is accessed by index. Only access an index lower than the slice's length. A subslice is accessed by index too: `subslice = slice[i:j]`.

Additional elements can be appended into available slice by function `append`:

```go
slice := []int{1,2,3,4}
slice = append(slice, 5) //slice: []int{1,2,3,4,5}
```

The reserved areas will be allocated for the appended elements. This allocation is very fast.

However, if there are more elements than the number of reserved areas, the slice will be broadened. In particular:

- A new array with the twice larger capacity will be created, having more reserved areas.
- Then, the elements in old slice will be copied into the new array, and the new reserved areas will be allocated for the appended elements.
- Finally, the old array is destructed and the old slice will point to the new array. 

This process is slow, therefore, when use slice to contain a **LARGE** number of elements, define the **LARGE CAPACITY** too.

```txt
Slice

__________________________
| slice[0]  |  not use   |   <- new element
|___________|____________|

<--length---><-reserve--->

            |
            | [Append a new element]
            V
__________________________
| slice[0]  |  slice[1]  |   <- new element
|___________|____________|

<---------length--------->

            |
            | [Broaden]
            V
____________________________________________
| slice[0]  |  slice[1]  |not use |not use |   
|___________|____________|________|________|

<---------length---------><--new reserve--->

            |
            | [Append a new element]
            V
____________________________________________
| slice[0]  |  slice[1]  |slice[2]|not use |   
|___________|____________|________|________|

<--------------length------------->

```

#### ***Copy slice***

When passing a slice to a function as parameter, or use an available slice to initialize value for a new slice, the copy of the slice will be used instead of the original one. However, the slice, basically, is a struct with its property data being a pointer to the real array in RAM. So that the copy of slice is also the copy of pointer, and making changes in the copy, in fact, is making changes in the original data.

```go
original_slice := []int{1,2,3,4}
new_slice := original_slice
new_slice[0] = 100 // original_slice: [100,2,3,4]
```

### **3.4| MAP**

*Map* is a data structure saving the *key-value* elements. Map has no order, `for range` loops over the elements in map with random order.

#### ***Define and initialize map***

Map must be defined and initialized by operator `:=` so that it can be allocated with a memory area.

```go
// empty map with function make
my_map := make(map[string]<type_of_value>)

// map with value
my_map := map[string]<type_of_value>{
    "key1": value1,
    "key2": value2,
}
```

#### ***Add/change element in map***

```go
new_map["key"] = <value>
```

#### ***Get an element***

Getting an element is usually used with checking the existence of that element.

```go
value, ok = new_map["Hung"]

//ok: True <=> the key exists, False: the key does not exist
```

#### ***Delete an element***

Use function `delete`

```go
delete(new_map, "Hung")
```

## **4. FUNCTION**

### ***Global function***

Global functions are functions defined outside the main function. Use key word `func` to define function.

```go
package main

// Public function has its first letter in name being uppercase
func PublicGlobalFunction(param1 int, param2 string) int {
    return ...
}

// Private function 
func privateGlobalFunction(param1 int, param2 string) int {
    return ...
}
```

### ***Local function***

Local functions are functions defined inside the main function. This kind of functions do not have official names. Instead, these functions are usually attached to variables, and use variables as the name of functions in use.

```go
func main(){
    // Define anonymous function
    variable := func(param1 int, param2 string) int {
        return param1*2
    }

    // Call function
    result := variable(20, "Hung")
}
```

### ***Anonymous function***

Anonymous functions are function having no name and usually used once.

```go
func main() {

    // Define and call immediately (by passing parameters right after the definition of function)
    result := func(param1 int, param2 int) {
        return param1 + param2
    }(1,2)
}
```

### ***Function of struct***

Struct is a special data type in Go, it allows to group the relative variables into an object so that the management becomes easier. Beside variables, struct also owns its functions, but functions of struct are not defined directly inside struct, struct only contains data which are variables. Instead, functions are defined outside the struct, then they are linked to the struct by **Receiver**.

If the struct is defined globally, its functions must be defined globally, too. The scope of the functions is independent on the scope of struct or variables inside struct.

```go
package main

type MyStruct struct {
    ID int
    Name string
}

// (m *MyStruct) is the definition of a receiver, passing pointer of the MyStruct object into receiver allows functions to change the properties inside object.
func (m *MyStruct) PublicFunction(param1 int, param2 string) int {
    m.ID = m.ID + param1
    m.Name = param2
    return m.ID
}

func (m *MyStruct) privateFunction(param1, param2 string) {

}
```

Similarly to the OOP language, to call functions of the struct:

```go
func main() {
    // Define an object
    object := MyStruct{
        ID: 1,
        Name: "A",
    }

    // Call functions
    object.PublicFunction(3, "B")
}
```

## **5. ERROR PROCESSING**

### **Key word `defer`**

When working with resources like database, file, .... sometimes we open resources very soon, and close after working with them in 1000 lines of code. This becomes danger because we may forget to close. Go uses the key word `defer` to deal with this situation.

```go
func ProcessResource() int {
    Openfile()
    defer Closefile()
    Processfile()

    return 0
}
```

The above code is similar to:

```go
func ProcessResource() int {
    Openfile()
    Processfile()
    Closefile()

    return 0
}
```

In fact, the key word `defer` can pause the execution of the function `Closefile()` until the function `ProcessResource()` prepares to return. Thanks to this, we can write `Closefile()` right after `Openfile()`, avoiding forgeting to close the resource.

In case of many `defer`, the function call after `defer` are stacked. Right before the function `ProcessResource()` returns, the function calls are poped and executed.

```go
func ProcessResource() int {
    Openfile1()
    defer Closefile1()
    Processfile1()

    Openfile2()
    defer Closefile2()
    Openfile3()
    defer Closefile3()
    Processfile2()
    Processfile3()

    return 0
}
```

The above code is similar to:

```go
func ProcessResource() int {
    Openfile1()
    Processfile1()

    Openfile2()
    Openfile3()
    Processfile2()
    Processfile3()

    Closefile3()
    Closefile2()
    Closefile1()

    return 0
}
```

### **Catch errors**

In other language, errors or exceptions are caught by `try ... except ...` (Python) or `try... catch ...` (Java). Howerver, go encourages developers on trapping and handling the errors explicitly, requires developers to predict which errors will occur, so that the code will be clear, maintainance will be easier. The error is considered as a variable with the type of `error`.

```go
var new_error error
```

There are 3 levels of throwing errors, corresponding to 3 levels of handling errors.

- No error:

    ```go
    var new_error error
    new_error = nil
    ```

- Print notification of static errors. This define a new error once. In future, if any functions need to throw this error, they can directly attach: `throw_error = defined_error`, and in handling error, we can also check if an error is the one defined directly by `==`:

    ```go
    import (
        "errors"
        "fmt"
    )

    // Define new error
    var new_error error
    New_error = errors.New("Description of error")

    // Throw error: function must return error
    func Func(param1 int, param2 string) (int, error){
        if !checkCondition(param1, param2) {
            throw_error := New_error
        }
        else {
            throw_error := nil
        }

        return param1, throw_error
    }

    // Check error:
    func main() {
        // Call function -> throw error
        value, e  = Func(1, "Hello")
        // Checking error
        if e == New_error {
            fmt.PrintLn("This is the new error defined")
        }
    }
    ```

- Print notification with context of errors. The notification shows the information of errors according to the context ("User not found: A", "User not found: B", ...). If create the static errors: `errorA := errors.New("User not found: A")` and `errorB := errors.New("User not found: B")`, `errorA` and `errorB` cannot be compared: `errorA == errorB` is always False while they are the same type of error "User not found". To deal with this situation: create a general error "User not found", then, other particular errors like "User not found: A" and "User not found: B" are wrapped by the general error.

    ```go
    import (
        "fmt"
        "errors"
    )

    //Define general error
    var General_error error = errors.New("General Error")

    //Throw error with context, wrap with the general error with '%w'
    func Func(param1 int, param2 string) (int, error) {
        if !checkCondition(param1, param2){
            context = makeContext(param1, param2)
            throw_error = fmt.Errorf("Context: %s, %w", context, General_error)
        }
        else{
            throw_error = nil
        }
        return param1, throw_error
    }

    // Checking error:
    func main() {
        // Call function -> throw error
        value, e  = Func(1, "Hello")
        // Checking error
        if errors.Is(e, General_error) {
            fmt.PrintLn("This is the general error defined")
        }
    }
    ```

## **6. CONCURRENCE PROCESSING**

Go can automatically manage the logical threads for concurrence processing. The developer do not need to create, manage threads as well as shared resources manually.

### **6.1 Concurrence processing with goroutines**

Use key word `go` in front of the calling of functions to register with go that the functions must be executed in an independent threads (called routine). The main function is also executed in an independent routine. However, this is the root routine, if this closes, the other sub routine created will close.

```go
import "time"

func main() {
    // Func1 is executed in an independent routine
    go Func1(1)
    // Func2 is executed in another routine
    go Func2("Hung")

    // main is in an independent root routine
    // waiting for the Func1 and Func2 routine of immediately close to avoid close sub routines
    time.Sleep(30)
}
```

Go will ask the OS for the physical threads (8 CPU cores in CPU - 8 physical threads), and manage to run all of the logical routines registerd in the physical threads. Mapping `m` logical routines to `n` physical threads.

```txt
            Developers
                |
                | [Register logical threads]
    ____________|_____________
    |           |            |   
    |           |            |   
Logical     Logical         Logical
routine 1    routine 2      routine 3
    |           |            |
    |           |            |
 ___V___________V____________V___
 |                              |
 |             Go               |
 |______________________________|
    |           /            |
    |          /             |
    |         /              |
    |        /               |
    |       /                |   
    |      /                 |   
    V     V                  V
    Physical            Physical
    thread 1            thread 2

```

### **6.2 Communication among routines**

Many routines are basically executed at the same time. The shared resources (like variables, RAM,...) are distributed to one routine at a moment, the other routines using those resources are blocked. In go, the shared resources are called **CHANNELS**. At a moment, there is totally one routine having permission to work with the *channel*. So that, the shared resources are processed sequentially by routines. This phenomenon creates an illusion that the calculated result are passed from a routine to another in a pipe, the *channel* plays the role of a *pipe*.

```txt
____________________
    Routine 1       --> processed result    
____________________                |
                                    |
                                    V
                                |       |
                                |       |
                                |       |
                                |channel|
                                |       |
                                |       |
                                |       |
                                    |
                                    |
____________________                V
    Routine 2       <-- processed result
____________________                                                        
```

```go
package main
import (
    "fmt"
    "time"
)

// create channel by make(chan <type_of_data_transported>)
var ch = make(chan int)

func Func1(param1 int, param2 int){
    processed_result := calculate1(param1, param2)

    // Push calculated result into channels, using the operator "<-"
    ch <- result
}

func main() {
    // routine calculates in an independent thread
    go Func1(1,2)

    // in main thread, waiting for result from channel, so that the main thread is blocked until the channel has data
    waiting_result := <- ch
    fmt.PrintLn(waiting_result)

    // if the main thread does not wait for the channel, it should sleep to wait for the completion of sub routines
    //time.Sleep(30)
}
```

Because the logical goroutines base on the physical threads of computer, we have to consider before create goroutines:

- If the functions are the Input-output tasks: create goroutines because the IO processes usually lasts long. When a routine is waiting for IO, it is blocked, and the physical thread can execute the other goroutines, avoiding wasting the threads.

- If the functions are calculations: avoid creating too many goroutines( <= the number of physical threads in computer) because many goroutines in this case can not help faster calculations, but waste time to transfer thread's processor among virtual goroutines.

#### ***Transporting many messages***

The above code is an example of transporting a calculated result from a goroutine to another. In fact, the producer may want to send many results, but if the customer is slow, the producer has to wait for the channel. Therefore, the channel is implemented with waited queue, this queue plays a role of buffer (called **buffer channel**) storing what the producer send, so that the producer can send continuously without waiting for the customer. We can loop over the channel by `for range`, this loop can wait for data in channel until the producer `close` channel.

```go
func main() {
    // Create channel with the size of 2
    ch := make(chan string, 2)

    go func() {
        a := "A"    // producer produces 1 result
        ch <- a     // ch has 1 element
        
        b := "B"    // producer produces 2 results
        ch <- b     // ch has 2 elements (full -> ch are blocked)

        c := "C"    // producer produces 3 results
        ch <- c     // cannot push element into ch because ch is full and blocked

        close(ch)   // close channel, this is for announcing the producer to stop waiting
        
    }()

    // Pop the element out from channel respectively by for... range...
    for item := range ch {
        process(item) // once the first item is poped from the channel, the producer routine can continue to push "c" into channel
    }

}
```

#### ***Many goroutines at the same time***

The main function is executed in an independent goroutine, and it should wait for the sub goroutines. But, not all of the sub goroutines can end, in this case, go uses `select` for waiting one sub goroutines to complete and determine the order of processing data in channels. `select` can help execute the earliest data from all channels.

```go
import "time"

func main() {
    ch1 := make(chan int, 2)
    ch2 := make(chan string, 2)

    // goroutine1: producer 1 produces message and push it into channel 1
    go func() {
        one := 1 // producer 1 produces 1 element
        ch1 <- one

        two := 2
        ch1 <- two

        close(ch1)
    }

    // goroutine1: producer 2 produces message and push it into channel 2
    go func() {
        a := "a" // producer 2 produces 1 element
        ch2 <- a

        b := "b"
        ch2 <- b

        c := "c"
        ch2 <- c

        close(ch2)
    }

    // main goroutine: receive and process
    // loop infinitely
    for {
        select {
        case msg1, ok1 := <- ch1:
            if ok1 == false {
                ch1 = nil // attach ch1 with nil so that select can automatically skip waiting for this channel
                continue
            }
        case msg2, ok2 := <- ch2:
            if ok2 == false {
                ch2 = nil
                continue
            }
        }
        // Waiting for the limited time instead of waiting for all channels infinitely
        // time.After returns a channel, so, use operator "<-" to get the value in the channel 
        case <-time.After(10*time.Second):
            break
    }
}
```
