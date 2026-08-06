# **ADVANCED GO**

## **1. FORMATING**

### **1.1| Json**

JSON is a popular communicating format. In go, Json message is usually slice of bytes instead of string because it can help with convenient communication and parsing into struct. To work with json: we have to import `"encoding/json"`

Create a Json message:

* From a string:

    ```go
    // Message in string type
    stringMessage := `{"key1": value1, "key2": value2}`
    // Parse string type to slice of bytes
    JsonMessage := []byte(stringMessage)

    ```

* From a struct object (marshal): Each property of object is a key. When define the struct, we must add tag `json:"key_name"` after each property.

    ```go
    package main

    import "encoding/json"

    // DEFINE STRUCT TYPE (Must have tag `json:"key_name"` after each property)
    // To define a struct corresponding to a multi-level json:
    // 1. Sub JSON struct
    type Mark struct {
        Math float `json:"math"`
        Physic float `json:"physic"`
        Chemistry float `json:"chemistry"`
    }
    // 2. Embed into parent struct
    type Student struct {
        Name string `json:"name"`
        Id int `json:"id"`
        mark Mark `json:"mark"`
    }

    // PARSE INTO SLICE OF BYTES (MARSHAL)
    func main(){
        // struct object
        student := Student{
            Name: "A",
            ID: 1,
            mark: Mark{
                Math: 10.0,
                Physic: 9.8,
                Chemistry: 9.0,
            },
        }
        // JsonMessage is the slice of byte
        // ok is nil if no error, else, ok is error
        JsonMessage, ok = json.Marshal(student)
    }
    ```

When receive a Json message (slice of bytes), we have to parse it into a go struct object (unmarshal) to process:

```go
package main

import "encoding/json"

// DEFINE STRUCT TYPE (Must have tag `json:"key_name"` after each property)
// To define a struct corresponding to a multi-level json:
// 1. Sub JSON struct
type Mark struct {
    Math float `json:"math"`
    Physic float `json:"physic"`
    Chemistry float `json:"chemistry"`
}
// 2. Embed into parent struct
type Student struct {
    Name string `json:"name"`
    Id int `json:"id"`
    mark Mark `json:"mark"`
}


// PARSE INTO SLICE OF BYTES (MARSHAL)
func main(){
    // JsonMessage is the slice of byte
    JsonMessage := []byte(`{"Name": "A", "ID": 1, "mark": {"Math": 10.0,"Physic": 9.8, "Chemistry": 9.0}}`)
    // struct object to receive messgae
    var student Student
    
    // Parse message into struct object (pass pointer into the function)
    json.Unmarshal(JsonMessage, &student)
}
```

**NOTE**: Properties in struct must be ***Public***, so that the function `json.Marshal()` and `json.Unmarshal()` in the package `encoding` can access and parse.

### **1.2| YAML**

**YAML** is a popular format for configuration. To work with yaml:

* Install: `go get gopkg.in/yaml.v3`

* Import: `import "gopkg.in/yaml.v3"`

* Define struct to receive configuration from a yaml format: similar as json, but the tags are `yaml:` instead of `json:`

* `yaml.Unmarshal()` and `yaml.Marshal()` are similar as `json.Unmarshal()` and `json.Marshal()`

**NOTE**: Properties in struct must be ***Public***, so that the function `json.Marshal()` and `json.Unmarshal()` in the package `encoding` can access and parse.

## **2. NETWORK COMMUNICATION**

### **2.1 TCP/IP Model**

TCP/IP model is a five-layer conceptual framework that defines how data is sent and received over the internet. Data is broken into small packets and these packets are routed across the network and reassembled at the destination. Five layers are:

* **Application Layer**:
  * Interacts directly with app to packaging data
  * Add header information including: sender, receiver, protocol,...
  * Protocols in application layer: HTTP, HTTPS, FPT,...

* **Transportation Layer**:

  As we have known, data is broken into smaller parts and then reassembled at the destination. These progresses are done by transportation layer:
  * Splitting the packet into segments when sending
  * Combining all segments when receiving, ensuring there order and then forwarding to the suitable application through port
  * Add the header into segments with some information including special flag to control the order of segment,...
  * Protocols in application layer: TCP, UDP, ...

* **Network Layer**:
  * Add header into each segment -> each packet
  * Routing the packets across different LANs: Determine which hop to send the packages next.
  * Protocols in network layer: IP

* **Data-link Layer**:
  * Add header into each packet (including: MAC address,...) -> frame
  * Routing the frame inside LAN

* **Physical Layer**:
  * Convert digital data to physical data
  * Transport the data through physical channels

```txt

                    Message
                       |
    _________ ___________________________
    | header |       Full data          |
    |________|__________________________|      
                       |
                       V
    ____ ___ ___ ___ ___ ___ ___ ___ ____
    |   |   |   |   |   |   |   |   |   |   
    |___|___|___|___|___|___|___|___|___|
                              |  
                     _______ _V_   
                     |header|   |
                     |______|___|
                     <--segment->
                            |
              _______ ______V____
              |header|           |
              |______|___________|
              <------packet------>
                        |
       _______ _________V________
       |header|                  |
       |______|__________________|
       <-----------frame--------->
                    |
                    V
                Physical

```

### **2.1| Http**

HTTP is short term of Hypertext Transfer Protocol, this is a protocol in application layer. The main tasks of the protocol is:

* when sending: wrap the message and add information of protocol.
* when receiving: unbox the wrapper to get the main message

When clients want servers to serve something, they will send **HTTP Request**, and the servers will return **HTTP Response**.

#### ***Http request***

A Htttp request includes:

* **METHOD**: What clients want to do with servers:

  `GET`(Ask server for information)

  `POST`(Send data to server to save in database)

  `PUT`/`PATCH`(Update available data in database)
  
  `DELETE`(Ask server to delete information in database)

* **URL**: Where the server is and What clients want to access. Example: `gemini.google.com/app/2ef0c9827c6bb690?hl=vi`. This includes 3 parts:

  * ***Domain***: `gemini.google.com`: IP or domain name of server

  * ***Path***: `/app/2ef0c9827c6bb690`: path to the resource that clients want to access

  * ***Query***: `?hl=vi`: parameters after `?`: to filter some parts of the resource

* **HEADER**: Metadata describes the request in `key-value` format. Header includes:
  
  * `Host`
  * `Content-Type`: What type of the information in body is (`application/json`, `text/html`,...)
  * `Authorization`: Security token
  * `User-Agent`: Who is sending? (Chrome or go applications,...)

* **BODY**: The message that clients want to send to server. If method is `GET`, this must be empty.

Example of a Http request:

```http
POST /api/v1/users?lang=vi HTTP/1.1
Host: mycompany.com
Authorization: Bearer xyz123
Content-Type: application/json

{"name": "Hung", "role": "admin"}
```

#### ***Http response***

When receives requests from clients, the server will handle the request by handler and return **Http response**. This includes:

* **STATUS LINE**: This shows the status of the request:

  * `2xx`: Success
  * `3xx`: Current server does not serve the requests, send to other server
  * `4xx`: Fail because of inavailable resources, wrong format,...
  * `5xx`: Error inside server

* **RESPONSE HEADER**: Metadata in `key-value` format describing the response. This includes:
  * `Content-Type`
  * `Content-Length`
  * `Set-cookie`
  * `Cache-Control`
  * `CORS Header`

* **BODY**

Example of a Http response:

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Content-Length: 42
Server: my-viettel-gateway

{"status": "success", "user_id": 8888}
```

### **2.2| Implementation**

Microservices communicate with each other by HTTP and the sender plays role of client, the receiver plays role of serverl.

In go, library `network/http` is responsible for http network communication.

#### **Data structure**

* `http.Request`: This is a struct in library `http`, implements a Http request.

  When clients ask the server, `http.Request` objects have to be created manualy by clients.

  In server's side, TCP gathers all the segments of the request in the right order, then forwards the full data to the application layer. Server opens a port to listen. Whenever the port receives data from TCP connection, a `http.Request` object is created automatically to parse the data.
  
  `http.Request` has some fields:

  |**Field**|**Type**|**Data**|
  |:-------:|:------:|:------|
  |`Method`|`string`|`GET`, `POST`, `PUT`, `DELETE`|
  |`URL`|`*url.URL`|Url|
  |`Header`|`http.Header`|A map contains metadata|
  |`Body`|`io.ReadCloser`|Stream byte|

* `http.ResponseWriter`: This is an interface in library `http`, helps servers to write http response to clients. There are 3 functions in this interface, must call respectively:
  * `Header()`: attach header. An object `http.response` is created and set header.
  * `WriteHeader(statusCode int)`: Write status code into the object `http.response`.
  * `Write([]byte)`: The calling of function `Write` will transfer the `http.response` object to a text in format `HTTP`. And then, this text is written to a buffer. This buffer is then forwarded to the transportation layer.

#### **Operating server**

The following functions are used in server's side:

* `http.HandlerFunc()`: This function is used to handle requests which match a particular pattern (API) and write response to clients.
  
  ```go
  http.HandlerFunc(
    // pattern
    "/url-pattern",
    // a handler function receiving a http.ResponseWriter to write response and a http.Request pointer (which is automatically created by go whenever the listening port receives a message from TCP connection)
    func(w http.ResponseWriter, r *http.Request){
        // Processing request r
        // Create http.response and set header
        w.Header()...
        // Set status code in http.response
        w.WriteHeader()...
        // Transfer http.response into HTTP text and save in a buffer, then forward buffer to TCP connection
        w.Write()
    }
  )
  ```

  --> When the server receives a request and its url matchs the pattern, the object of `http.Request` is created to parse request and the pointer of object is passed into the handler function. The handler function is immediately activated to process the `http.Request`, then write response by `http.ResponseWriter` to client.

* `http.NewServeMux()`: This separates handlers into different muxes processing different APIs. `http.HandlerFunc()` is the default mux processing API. However, implementing the same handler by `http.HandlerFunc()` to process all the APIs can allow clients to access the private APIs. Therefore, `http.NewServeMux()` is created to wrap the handler for scope management, and this serveMux will be called instead of default `http.HandlerFunc()`.

  ```go
  publicMux := http.NewServeMux()
  publicMux.HandlerFunc("/api/buy", handleBuy)
  publicMux.HandlerFunc("/api/pay", handlePay)

  privateMux := http.NewServeMux()
  privateMux.HandleFunc("/admin/delete-db", handleDelete)

  // public /api/buy is listened at port 8080
  go http.ListenAndServe(":8080", publicMux) 
  // private /admin/delete-db is listened at port :9090
  go http.ListenAndServe(":9090", privateMux)
  ```

* `http.ListenAndServe()`: Open a port to listen requests from a client. A port is responsible for listening a group of APIs and they are handled by the same mux. This port is run independently in a goroutine, the number of goroutines corresponds to the number of clients.

  ```go
  // serveMux
  http.ListenAndServe(":port", serveMux)
  ```

#### **Client-side**

The following functions are used in client-side to simulate a virtual browser

* `http.Client`: this is a struct in library `http`, with some fields: `Timeout`(the time of waiting for response from the server), ... In client-side, `http.Request` has to be created manually, and it will be sent to server by `http.Client`

  ```go
  client := &http.Client{
    // After 5 seconds, if there is no response, the client will interrupt
    Timeout: 5*time.Second,
  }

  // Step 1: Create http.Request
  jsonData := []byte(`{"name": "A"}`)
  req, err := http.NewRequest(
    http.MethodPost, // METHOD
    "http://10.0.0.5:8080//api/data", // URL
    bytes.NewBuffer(jsonData) // BODY 
  )
  req.Header.Set("Content-Type", "application/json") // Set header by req.Header.Set("key-name", value)
  req.Header.Set("Authorization", "Bearer token123")

  // Step 2: Send request and receive response
  // Response is type *http.Response
  resp, err := client.Do(req)
  if err == nil {
    // Sending request sucessfully
  }
  defer resp.Body.Close()

  // Step 3: Process response
  if resp.StatusCode != http.StatusOk{ // Response fail
  }
  else{
    // Create a middle variable to store all the content in body in response
    // But this does not optimize performance
    bodyBytes, _ := io.ReadAll(resp.Body)
    json.Unmarshal(bodyBytes, &structObject)
    
    // Read directly on buffer
    json.NewDecoder(resp.Body).Decode(&structObject)
  }
  ```

#### **Context**

##### <medium>**[ROLE]**</medium>

In go, context is used to solve the problem of waiting:

* Remember that when create a client microservice, we have to initialize the time out. But, responses of different APIs need to be waited with different durations (uploading files through APIs `/api/upload/` may last longer than getting information `/api/get_users/`). So that, we can not use the same initialized time out for all APIs.

  --> use **different contexts** for waiting different APIs'responses.

* When time out, the client will unilaterally stop connection to server. In this case, the server does not need to continue processing the clients'requests, and it can interrupt to save the resources. However, how the server can know when the clients stop connection.

  --> use **context** to announce to server when TCP connection is broken.

##### <medium>**[IMPLEMENTATION]**</medium>

Context is responsible for following the lifecycle of a request.

When processing a request from client, the server may call to APIs of other microservices. Therefore, some contexts are created to follow these APIs.  

```txt
    _
    |  --> Request from client    
    |         |                              
    |         |                              
    |      _  |                     
    | ctx1|   --> APIs calculating  
ctx0|     |_  <-- calculated result 
    |         |
    |         |                              
    |      _  |                              
    | ctx2|   --> APIs writting to database
    |     |_  <-- finish
    |         |
    |         |                              
    |         |                              
    |_  <-- Write response to client

ctx0: context follows the processing of request
ctx1: context follows the calculation
ctx2: context follows writting in database
```

In the above example, `ctx0` is used to follow the request processing, while `ctx1` and `ctx2` are used to follow the sub processing.

* `ctx0` cannot wait for both `ctx1` and `ctx2` until they finish because in case of both contexts are lag and return no responses, the client has to wait infinitely. So, each context has a limited waiting time, but, they may be different (`ctx0: 10s`, `ctx1: 8s`, `ctx2: 4s`) because of the different workload of the distinct tasks.

* However, waiting for both `ctx1` and `ctx2` time out (`8s+4s=12s`) lasts too long. When `ctx1` time out, it mostly returns nothing, so there is nothing to write in database and continue to wait is time-wasting. Therefore, `ctx0` should break and cancel when time out (`10s`), ignore the completion of the subcontexts.

  In that case, it is unneccessary for `ctx1` and `ctx2` to continue to process. So that: `ctx0` cancelation should cancel `ctx1` and `ctx2`.

  But, if the `ctx1` can return result and then it is canceled, the result must be continued to write into database. So that the cancelation of the subcontexts like `ctx1` must not cancel the parent context `ctx0`.

  ➡️ Therefore, go uses tree-structrure to organizes the contexts which follow a request: `ctx0` is root, `ctx1` is the child of `ctx0`, `ctx2` is the child of `ctx1`; or `ctx0` is root, `ctx1` and `ctx2` are children of root `ctx0`. If a context is cancelled, all of its descendants will be cancelled but its parent and ancestor won't.

#### <medium>**[CODE]**</medium>

* Import: `import "context"`

* `context.Context`: an interface in library `context`. To follow a request, after being initialized, the context must be attached to the API of the request. The most important function in interface is `Done()`, this return a value in an empty channel (Explain later after code)

* `context.Background()`: a function to initialize an empty context as root of tree. This context does not follow any request.

* `context.WithTimeout()`: a function in library `context`.

  * Input:
    * A parent context (type `context.Context`)

    * The duration time (type `time.Duration`)

  * Output:
    * A child context (type `context.Context`): this context has waiting time, if timeout, a cancellation is automatically called to interrupt waiting.

    * The cancellation function: This function must be always called manually to avoid waiting for time out in case of that the server can return response earlier than the deadline.

* `context.WithCancel()`: This is a function creating a context with manual cancellation.
  
  * Input:
    * A parent context (type `context.Context`)

  * Output:
    * A child context (type `context.Context`): this context has no waiting time, this context will wait infinitely if it is not cancelled manually.

    * The cancellation function: This function must be always called manually.

* `http.NewRequestWithContext()`: This helps attach a context to a request:
  
  * Input:

    * Context (type `context.Context`)
    * Method (type `string`): `"GET"`, `"POST"`,...
    * Url (type `string`)
    * Body (tpe `[]byte`)
  
  * Output:

    * A `http.Request`
    * Error of initializing request

Flow of usage:

```go
package main
import (
  "net/http"
  "context"
  "time"
  "fmt"
  "encoding/json"
)

type User struct{
  Name string
  Id int
}

func main() {
  client := &http.Client{
    Timeout: 10*time.Second,
  }

  // Step 1: Initialize context with duration
  ctx, cancel := context.WithTimeout(
    context.Background(),
    3*time.Second
  )
  defer cancel() // cancellation function must be always called manually

  // Step 2: Initialize request with context
  request, _ := http.NewRequestWithContext(
    ctx,
    "GET",
    "http://192.168.1.1:8080/api/user-info",
    nil
  )

  // Step 3: Send request and get response (or error)
  res, err := client.Do(request)
  defer res.Body.Close()

  // Step 4: Wait for result, or timeout
  if err != nil {
    if ctx.Err() == context.DeadlineExceeded {
      fmt.Println("Time out")
    } else {
      fmt.Println("Other error", err)
    }
    return
  }

  var user User
  err = json.NewDecoder(res.Body).Decode(&user)
  if err != nil {
    fmt.Println("Error in parsing JSON:", err)
    return
  }
}
```

In the above code, how the program can interrupt when time out?

* Step 1: When initializing a new context with time out: a new context will be created and an empty channel is created, too. This channel is used only for block the thread while waiting for time out, not used to pass data.

* Step 3: `client.Do(request)` use `select` to wait for the ealier between the time out and the response from server. The second case in `select` block(`case <- ctx.Done()`) waits for something in the channel returned by function `ctx.Done()` (this is also the channel initialized in step 1), but that is an empty channel, and nothing can be get, so the thread is blocked until there are something in the channel.

* When time out, Go runtime closes the channel and awakes the thread, the value returned is `nil`. Then, the thread wakes up and interrupt TCP connection to stop waiting.

* So we can see, the channel here is playing the role of an alarm. We do not care what is in the channel, but we care about when the channel is closed. Thanks to the channel, the thread is blocked while waiting, saving the CPU resources.
