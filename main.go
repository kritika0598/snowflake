package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	// Unix timestamp for 2023-01-01 00:00:00 UTC in milliseconds
	epoch    = 1672531200000
	nodeBits = 10
	seqBits  = 12
	maxNode  = 1 << nodeBits
	maxSeq   = 1 << seqBits
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var (
	mutex    sync.Mutex
	sequence int
)

type Snowflake int64

func GenerateSnowflake() Snowflake {
	node := seededRand.Intn(maxNode)
	mutex.Lock()
	defer mutex.Unlock()

	currentTime := time.Now().UnixMilli()
	sequence = (sequence + 1) % maxSeq

	id := int64(0)
	id |= (int64(currentTime-epoch) << (nodeBits + seqBits))
	id |= (int64(node) << seqBits)
	id |= (int64(sequence))
	return Snowflake(id)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println(GenerateSnowflake())
	}()

	go func() {
		defer wg.Done()
		fmt.Println(GenerateSnowflake())
	}()

	wg.Wait()
}
