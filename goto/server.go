package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

import (
	"gopkg.in/redis.v5"
)

var (
	redisOption redis.Options
	redisClient *redis.Client
)

func main() {
	flag.Parse()
	initRedisClient()

	http.HandleFunc("/", redirect)
	http.HandleFunc("/add", addUrl)
	http.ListenAndServe(DOMAIN_PORT, nil)
}

func init() {
	flag.StringVar(&redisOption.Addr, "redis", "", "ip:port of the redis service.")
}

func initRedisClient() {
	redisClient = redis.NewClient(&redisOption)
	if redisClient.Ping().Err() != nil {
		log.Fatalln("Failed to connect redis server at", redisOption.Addr)
	} else {
		log.Println("Connected to redis server at", redisOption.Addr)
	}
}

func redirect(response http.ResponseWriter, request *http.Request) {
	key := request.URL.EscapedPath()[1:]
	url, err := redisClient.Get(key).Result()
	if err != nil && err != redis.Nil {
		log.Printf("Error happened while serve %s: %v", request.URL.EscapedPath(), err)
		http.Error(response, "Unable to serve this request!", http.StatusInternalServerError)
		return
	}
	if err == redis.Nil || len(url) == 0 {
		http.Error(response, key+" is not on record!", http.StatusNotFound)
	} else {
		http.Redirect(response, request, url, http.StatusFound)
	}
}

const (
	DOMAIN_PORT = "localhost:8080"
	FORM        = `<html><body><form method="POST" action="/add">
		URL: <input type="text" name="url"><input type="submit" value="Add">
		</form></body></html>`
)

func addUrl(response http.ResponseWriter, request *http.Request) {
	url := request.FormValue("url")
	if len(url) == 0 {
		fmt.Fprint(response, FORM)
		return
	}
	for {
		digest := md5.Sum([]byte(url + strconv.Itoa(rand.Int())))
		key := fmt.Sprintf("%x", digest)[:4+rand.Intn(4)]
		success, err := redisClient.SetNX(key, url, 0).Result()
		if err != nil && err != redis.Nil {
			log.Printf("Error happened while adding %s: %v", url, err)
			http.Error(response, "Unable to serve this request!", http.StatusInternalServerError)
			return
		}
		if success {
			fmt.Fprintf(response, "<html><body><pre>%s</pre> is added as http://%s/%s</body></html>", url, DOMAIN_PORT, key)
			break
		} else {
			log.Printf("Conflict key: ", key)
		}
	}
}
