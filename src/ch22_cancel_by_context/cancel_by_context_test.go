package cancel

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// grpc的stream判断是否关闭，可以用该方式
func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	// context.WithCancel 创建一个传入Context的副本，并提供一个调用后能够关闭其中chan的函数
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	// 调用上面返回的取消函数，即可关闭该Context
	cancel()
	time.Sleep(time.Second * 1)
}
