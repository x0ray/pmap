// pmap package creates a persistant map of key value pairs
package pmap

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Pmap - persistant map
type Pmap struct {
	Name string
	Path string
	Pm   map[string]interface{}
}

// New creates a new persistent map
func New(name string, path string) (*Pmap, error) {
	var err error
	if name == "" {
		err = errors.New("empty pmap name not valid")
		return nil, err
	}
	if !strings.HasSuffix(name, ".pmap") {
		name = name + ".pmap"
	}
	p := new(Pmap)
	p.Pm = make(map[string]interface{})
	p.Name = name
	p.Path = path
	if p.Path == "" {
		p.Path, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	}
	err = p.load()
	return p, err
}

// Copy a pmap to a new  pmap
func (p *Pmap) Copy(newname string) (*Pmap, error) {
	var err error
	if newname == "" {
		err = errors.New("empty pmap name not valid")
		return nil, err
	}
	if !strings.HasSuffix(newname, ".pmap") {
		newname = newname + ".pmap"
	}
	if newname == p.Name {
		err = errors.New("new pmap name equals old pmap name, not valid")
		return nil, err
	}
	q := new(Pmap)
	q.Pm = make(map[string]interface{})
	// copy the old to new pmap data
	q.Name = newname
	q.Path = p.Path
	for k, v := range p.Pm {
		q.Pm[k] = v
	}
	return q, err
}

// Prints the contents of the pmap with an optional title
func (p *Pmap) Print(title string) {
	var ss []string
	// make slice of keys
	for k, _ := range p.Pm {
		ss = append(ss, k)
	}
	// sort keys
	sort.Strings(ss)
	// print sorted keys and their values
	fmt.Printf("%sPmap name: %s path: %s ...\n", title, p.Name, p.Path)
	for _, k := range ss {
		v := p.Pm[k]
		fmt.Printf("  %s: %v\n", k, v)
	}
	fmt.Printf("Len: %d Size: %d \n", p.Len(), p.Size())
}

// loads a pmap from persistent storage
func (p *Pmap) load() error {
	filename := filepath.Join(p.Path, p.Name)
	f, err := os.Open(filename)
	if err == nil {
		decoder := gob.NewDecoder(f)
		err = decoder.Decode(p)
		if err != nil {
			return err
		}
	}
	f.Close()
	return nil
}

// Closes and saves the pmap
func (p *Pmap) Close() error {
	filename := filepath.Join(p.Path, p.Name)
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	encoder := gob.NewEncoder(f)
	err = encoder.Encode(&p)
	if err != nil {
		return err
	}
	err = f.Close()
	return err
}

// Add a new key with its value to the map
func (p *Pmap) Add(key string, val interface{}) {
	if _, ok := p.Pm[key]; !ok {
		p.Pm[key] = val
	}
}

// Replace an existing key with a new value
func (p *Pmap) Replace(key string, val interface{}) {
	p.Pm[key] = val
}

// Delete a key from the map
func (p *Pmap) Delete(key string) {
	delete(p.Pm, key) // remove element
}

// Exist tests if the specified key is in the map
func (p *Pmap) Exist(key string) bool {
	_, ok := p.Pm[key] // check if key exists
	return ok
}

// GetVal gets the value for the specified key
func (p *Pmap) GetVal(key string) interface{} {
	return p.Pm[key]
}

// Size returns the size in bytes of the storage used by the entire map
func (p *Pmap) Size() int {
	var sz int
	sz += len(p.Name) + len(p.Path)
	for k, v := range p.Pm {
		sz += len(k)      // add size of the key
		switch v.(type) { // add size of the value
		case uint:
			sz += 8
		case int:
			sz += 8
		case uint64:
			sz += 8
		case int64:
			sz += 8
		case uint32:
			sz += 4
		case int32:
			sz += 4
		case uint16:
			sz += 2
		case int16:
			sz += 2
		case uint8:
			sz += 1
		case int8:
			sz += 1
		case float64:
			sz += 8
		case float32:
			sz += 4
		case string:
			sz += len(v.(string))
		case []byte:
			sz += len(v.([]byte))
		default:
			func() {
				// defer a recover func to trap panic on Sprintf()
				defer func() {
					if x := recover(); x != nil {
						log.Printf("run time panic: %v", x)
					}
				}()
				sz += len(fmt.Sprint(v)) // approximation
			}()
		}
	}
	return sz
}

// Len returns the number of entries in the map
func (p *Pmap) Len() int {
	return len(p.Pm)
}

// Increment adds one to a key with a numeric value
func (p *Pmap) Increment(key string) {
	if val, ok := p.Pm[key].(int); ok {
		p.Pm[key] = val + 1
	} else if val, ok := p.Pm[key].(float64); ok {
		p.Pm[key] = val + 1.0
	}
}

// Decrement subtracts one from a key with a numeric value
func (p *Pmap) Decrement(key string) {
	if val, ok := p.Pm[key].(int); ok {
		p.Pm[key] = val - 1
	} else if val, ok := p.Pm[key].(float64); ok {
		p.Pm[key] = val - 1.0
	}
}

// Update saves the current map to persistent storage
func (p *Pmap) Update() error {
	err := p.Close()
	if err != nil {
		return err
	}
	err = p.load()
	return err
}
