package main

/* Expected output:
david@rat:~/go/src/github.com/x0ray/pmap/pmaptest$ ./pmaptest
2021/02/15 17:42:01 PMap test, pmaptest.go ver 0.0.1 of 13Feb2021
Start Pmap name: david.pmap path: /home/david/go/src/github.com/x0ray/pmap/pmaptest ...
  abc: 1239
  def: A dog
  fish: Very big fish
  pi: 1.2731477279816295e+06
  struct: {gg 4005 123.45 [106 107 108 109 110]}
  time: 2021-02-14T18:00:31.243383813-05:00
Len: 6 Size: 188
After add Pmap name: david.pmap path: /home/david/go/src/github.com/x0ray/pmap/pmaptest ...
  abc: 1239
  def: A dog
  fish: Very big fish
  pi: 3.4374988655504e+06
  struct: {gg 4005 123.45 [106 107 108 109 110]}
  time: 2021-02-14T18:00:31.243383813-05:00
Len: 6 Size: 188
After close and reload Pmap name: david.pmap path: /home/david/go/src/github.com/x0ray/pmap/pmaptest ...
  abc: 1240
  def: A dog
  fish: Very big fish
  pi: 3.4374988655504e+06
  struct: {gg 4005 123.45 [106 107 108 109 110]}
  time: 2021-02-15T17:42:01.258134043-05:00
Len: 6 Size: 188
2021/02/15 17:42:01 PMap test, pmaptest.go ended
*/
import (
	"encoding/gob"
	"log"
	"time"

	"github.com/x0ray/pmap"
)

const (
	pgm = "pmaptest.go"
	ver = "0.0.1"
	dat = "13Feb2021"
)

type T struct {
	A string
	B int
	C float64
	D []byte
}

func main() {
	gob.Register(T{}) // all structured types stored in pmap must be registered

	log.Printf("PMap test, %s ver %s of %s", pgm, ver, dat)
	tm, err := pmap.New("david", "")
	if err != nil {
		log.Printf("Problems: %v", err)
		return
	}
	tm.Print("Start ")
	if !tm.Exist("abc") {
		tm.Add("abc", 1234)
		tm.Add("pi", 3.1415926)
	}
	abc := tm.GetVal("abc").(int)
	if abc > 1240 {
		tn, err := tm.Copy("billy")
		if err != nil {
			log.Printf("Problems: %v", err)
			return
		}
		tm.Delete("abc")
		tn.Print("New copy ")
		tn.Close()
	}
	tm.Replace("def", "A dog")
	tm.Replace("struct", T{A: "gg", B: 4005, C: 123.45, D: []byte("jklmn")})
	iv := tm.Pm["pi"].(float64)
	tm.Pm["pi"] = iv * 2.7
	tm.Print("After add ")
	err = tm.Close()
	if err != nil {
		log.Printf("Close problem: %v", err)
	}

	mm, err := pmap.New("david", "")
	if err != nil {
		log.Printf("Problems: %v", err)
		return
	}
	mm.Add("fish", "Very big fish")
	mm.Replace("time", time.Now().Format(time.RFC3339Nano))
	mm.Increment("abc")
	mm.Print("After close and reload ")
	err = mm.Close()

	log.Printf("PMap test, %s ended", pgm)
}
