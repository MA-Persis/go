// 在前面的例子中，我们看到了如何使用原子操作来管理简单的计数器。对于更加复杂的情况，我们可以使用一个互斥锁来在 Go 协程间安全的访问数据。

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int) // 在我们的例子中，state 是一个 map。

	var mutex = &sync.Mutex{} // 这里的 mutex 将同步对 state 的访问。

	var ops int64 = 0 // we'll see later, ops will count how manyoperations we perform against the state.
	// 为了比较基于互斥锁的处理方式和我们后面将要看到的其他方式，ops 将记录我们对 state 的操作次数。

	// 这里我们运行 100 个 Go 协程来重复读取 state。
	for r := 0; r < 100; r++ { // 为了确保这个 Go 协程不会在调度中饿死，我们在每次操作后明确的使用 runtime.Gosched()进行释放。
		// 这个释放一般是自动处理的，像例如每个通道操作后或者 time.Sleep 的阻塞调用后相似，
		// 但是在这个例子中我们需要手动的处理。
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	// 同样的，我们运行 10 个 Go 协程来模拟写入操作，使用和读取相同的模式。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
