# pmap
Persistent map package for Golang

This package generates a Golang map with a string key and an empty interface as the value part. This allows the pmap to be used as a key value store, which can store any number of keys and associated values of any type.

## Installing the pmap package

``` sh
cd ../src
go get github.com/x0ray/pmap
```

## Using the pmap package

Create and save a pmap with one entry in the KV store.
``` go
package main
import "github.com/x0ray/pmap"
func main() {
    m,_ := pmap.New("test","")     // create pmap called "test.pmap" in cwd
    m.Add("abc", 123)              // add a value for key "abc"
    m.Close()                      // close and save
}
```

Access the previously saved KV store
``` go
package main
import "github.com/x0ray/pmap"
func main() {
    m,_ := pmap.New("test","")     // access the pmap called "test.pmap" in cwd
    m.Replace("abc", 456)          // replace the key "abc" with new value
    m.Close()                      // close and save
}
``` 

## Available pmap methods

### New
New creates a new persistent map
``` go
func New(name string, path string) (*Pmap, error) 
```

### Print
Prints the contents of the pmap with an optional title
``` go
func (p *Pmap) Print(title string) 
```

### Close
Closes and saves the pmap
``` go
func (p *Pmap) Close() error 
```

### Add
Add a new key with its value to the map
``` go
func (p *Pmap) Add(key string, val interface{}) 
```

### Replace
Replace an existing key with a new value
``` go
func (p *Pmap) Replace(key string, val interface{}) 
```

### Delete
Delete a key from the map
``` go
func (p *Pmap) Delete(key string) 
```

### Exist
Exist tests if the specified key is in the map
``` go
func (p *Pmap) Exist(key string) bool 
```

### GetVal
GetVal gets the value for the specified key
``` go
func (p *Pmap) GetVal(key string) interface{} 
```

### Size
Size returns the size in bytes of the storage used by the entire map
``` go
func (p *Pmap) Size() int 
```

### Len
Len returns the number of entries in the map
``` go
func (p *Pmap) Len() int 
```

### Increment
Increment adds one to a key with a numeric value
``` go
func (p *Pmap) Increment(key string) 
```

### Decrement
Decrement subtracts one from a key with a numeric value
``` go
func (p *Pmap) Decrement(key string) 
```

### Update
Update saves the current map to persistent storage
``` go
func (p *Pmap) Update() error 
```