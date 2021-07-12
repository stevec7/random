package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/allegro/bigcache"
	"github.com/gomodule/redigo/redis"
)

type record struct {
	k string
	v string
}

func plainClient(iq <-chan string, up chan<- record) error {
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)
	for i:= 0; i < 8; i++ {
	//for i:= 0; i < 5; i++ {
		go func(iiq <-chan string, uup chan<- record, rng *rand.Rand, thrId int) {
			for {
				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond * 5000))
				defer cancel()
				//ctx := context.Background()
				select {
				case query := <-iq:
					v := rng.Intn(1000)
					vv := strconv.Itoa(v)
					fmt.Printf("PLAINHIT_%d: k: %s, v: %s\n", thrId, query, vv)
					fmt.Printf("PLAIN_%d: Sending, k: %v, v: %v, UP\n", thrId, query, vv)
					up <- record{k: query, v: vv}
					fmt.Printf("PLAIN_%d: SUCCESS Sending, k: %v, v: %v, UP\n", thrId, query, vv)
				case <-ctx.Done():
						fmt.Printf("PLAIN_%d: Timed out\n", thrId)
				}
			}
		}(iq, up, r, i)
	}
	return nil
}

func bigcacheClient(c *bigcache.BigCache, iq <-chan string, oq chan<- string, up <-chan record) error {
	for {
		select {
		case query := <-iq:
			entry, err := c.Get(query)
			if err != nil {
				fmt.Printf("BCMISS: Sending %s to redis lookup\n", query)
				oq <- query
				fmt.Printf("BC: SUCCESS Sending %s to redis lookup\n", query)
			} else {
				fmt.Printf("BCHIT: k: %v, v: %s\n", query, string(entry))
			}
			case rec := <-up:
				fmt.Printf("BC: Adding k: %v, v: %v to bigcache\n", rec.k, rec.v)
				err := c.Set(rec.k, []byte(rec.v))
				if err != nil {
					fmt.Printf("BC: Error adding, %s\n", err)
				}
		}
	}
	return nil
}

func RLookup(conn redis.Conn, key string) (string, error) {
	v, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return v, nil
}

func RSet(conn redis.Conn, key, value string) error {
	_, err := redis.String(conn.Do("SET", key, value))
	if err != nil {
		return err
	}
	return nil
}

func redisClient(r redis.Pool, iq <-chan string, oq chan<- string, up <-chan record) error {

	for i := 0; i < 4; i++ {
		go func(iiq <-chan string, uup <-chan record) {
			conn := r.Get()
			defer conn.Close()
			for {
				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond * 5000))
				defer cancel()
				//ctx := context.Background()
				select {
				case query := <-iiq:
					//entry, err := redis.String(r.Do("GET", query))
					entry, err := RLookup(conn, query)
					if err != nil {
						fmt.Printf("RMISS: Sending %s to lctl lookup\n", query)
						oq <- query
						fmt.Printf("R: SUCCESS: Sending %s to lctl lookup\n", query)
					} else {
						fmt.Printf("RHIT: k: %v, v: %s\n", query, string(entry))
					}
					case rec := <-uup:
						fmt.Printf("R: Adding k: %v, v: %v to redis\n", rec.k, rec.v)
						//_, err := redis.String(r.Do("SET", rec.k, rec.v))
						err := RSet(conn, rec.k, rec.v)
						if err != nil {
							fmt.Printf("R: Error adding, %s\n", err)
						}
					case <-ctx.Done():
						fmt.Printf("R: Timed out")
				}
			}
		}(iq, up)
	}
	return nil
}

func lookup(key string, bcq chan<- string) error {
	//fmt.Printf("LOOKUP: Sending k: %s to big cache\n", key)
	fmt.Printf("LOOKUP: Sending k: %s to redis\n", key)
	bcq <- key
	return nil
}


func main() {
	fmt.Printf("Hello\n")

	//cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	//if err != nil {
	//	panic(fmt.Sprintf("initializing bigcache, %s\n", err))
	//}

	//rconn, err := redis.Dial("tcp", ":6379")
	//if err != nil {
	//	panic(fmt.Sprintf("connecting to redis, %s\n", err))
	//}
	pool := redis.Pool{
		MaxActive: 8,
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(fmt.Sprintf("creating redis pool, %s", err))
			}
			return c, nil
		},
		TestOnBorrow: func (c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	//cache.Set("0", []byte("val0"))

	////n, err := rconn.Do("GET", "1")
	//n, err := redis.String(rconn.Do("GET", "1"))
	//if err != nil {
	//	panic(fmt.Sprintf("getting val, %s", err))
	//}
	//fmt.Printf("n: %+v\n", n)
	//bigCacheCh := make(chan string)
	lookupCh := make(chan string)
	//redisCh := make(chan string)
	plainCh := make(chan string)
	popCh := make(chan record)

	// start workers
	//go bigcacheClient(cache, bigCacheCh, redisCh, popCh)
	go redisClient(pool, lookupCh, plainCh, popCh)
	go plainClient(plainCh, popCh)

	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)

	x := 100000
	for i := 0; i < x; i++ {
		j := strconv.Itoa(r.Intn(10000000))
		_ = lookup(string(j), lookupCh)
	}

	time.Sleep(time.Duration(time.Second * 1))
}
