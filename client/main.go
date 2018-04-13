package main

import (
	"bufio"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	nbElement := flag.Int("n", 4096, "Number of elements added to the cache")
	keyLength := flag.Int("k", 16, "Length of the keys")
	port := flag.Int("port", 8080, "Port of the webserver")
	server := flag.String("server", "localhost", "Sever address")
	//nbQuery := flag.Int("q", 4096, "Number of queries")
	flag.Parse()

	f, err := os.Create("urls")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	urlBase := "http://" + *server + ":" + strconv.Itoa(*port) + "/"
	for i := 0; i < *nbElement; i++ {
		req, err := http.NewRequest("GET", urlBase+"add", nil)
		if err != nil {
			log.Fatalln(err)
		}

		q := req.URL.Query()
		key := randString(*keyLength)
		q.Add("key", key)
		q.Add("value", randString(350))
		req.URL.RawQuery = q.Encode()

		_, err = http.Get(req.URL.String())
		if err != nil {
			log.Fatalln(err)
		}
		_, err = w.WriteString(urlBase + "query?key=" + key + "\n")
		if err != nil {
			log.Fatalln(err)
		}
	}
	w.Flush()
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var randTime = time.Now().UnixNano()
var randSrc = rand.NewSource(randTime)

func resetRand() {
	randSrc = rand.NewSource(randTime)
}

func randString(n int) string {
	b := make([]byte, n)
	// A randSrc.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, randSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
