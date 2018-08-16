# Go NaCl Experiment

Native Client (NaCl) is a sandboxing technology by Google

## Background Reading

* https://blog.golang.org/playground
* https://github.com/golang/go/wiki/NativeClient
* https://golang.org/misc/nacl/README
* https://golang.org/s/go13nacl
* http://research.google.com/pubs/archive/34913.pdf (research paper)

## Prerequisites

1. Download the NaCl SDK [zip file](https://developer.chrome.com/native-client/sdk/download)

2. Unzip into `/usr/local/bin/nacl_sdk` and run `./naclsdk update`

    - If you run into SSL certificate errors, download this [zip file](https://github.com/Kagami/nacl_sdk) instead
    
3. Run the following (verify beforehand which version of `pepper` was installed in the update):

    - `ln -nfs /usr/local/bin/nacl_sdk/pepper_49/tools/sel_ldr_x86_32 /usr/local/bin/sel_ldr_x86_32`
    - `ln -nfs /usr/local/bin/nacl_sdk/pepper_49/tools/sel_ldr_x86_64 /usr/local/bin/sel_ldr_x86_64`
    - ``ln -nfs `go env GOROOT`/misc/nacl/go_nacl_amd64p32_exec /usr/local/bin/go_nacl_amd64p32_exec``
    - ``ln -nfs `go env GOROOT`/misc/nacl/go_nacl_386_exec /usr/local/bin/go_nacl_386_exec``
    
## Example 1 - Hello World

1. Take a look at the example hello world program `example1.go`

    - Run it regularly (not NaCl) with `go run example1.go`

2. Build an NaCl binary with `env GOOS=nacl GOARCH=amd64p32 go build example1.go`

3. Run with `sel_ldr_x86_64 -S -e ./example1`

4. You should see:
    ```
    [67861,2372682624:00:54:11.097894] Native Client module will be loaded at base address 0x00004fd100000000
    hello world
    ```
    
## Example 2 - File Access Protection

1. Take a look at the example hello world program `example2.go`

    - Run it regularly (not NaCl) with `go run example2.go` (the file `output.txt` will be created)

2. Build an NaCl binary with `env GOOS=nacl GOARCH=amd64p32 go build example2.go`

3. Run with `sel_ldr_x86_64 -S -e ./example2`

4. You should see that `output.txt` was not created
