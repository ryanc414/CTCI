package threads

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	ctx := context.Background()
	pool := NewLockPool(5)

	t.Run("validation", func(t *testing.T) {
		// invalid lock ID
		err := pool.Lock(ctx, -1)
		require.Error(t, err)
		assert.Equal(t, "invalid lock ID", err.Error())

		// no register
		err = pool.Lock(ctx, 0)
		require.Error(t, err)
		assert.Equal(t, "need to register thread", err.Error())
	})

	t.Run("acquire/re-acquire", func(t *testing.T) {
		// acquire some locks
		ctx = pool.RegisterThread(ctx)

		for i := 0; i < 5; i++ {
			err := pool.Lock(ctx, i)
			require.NoError(t, err)
		}

		t.Log("acquired all locks")

		// cannot acquire same lock twice
		err := pool.Lock(ctx, 0)
		require.Error(t, err)
		assert.Equal(t, "deadlock detected", err.Error())

		t.Log("deadlock detected")

		// unlock them all
		for i := 0; i < 5; i++ {
			err = pool.Unlock(ctx, i)
			require.NoError(t, err)
		}

		t.Log("unlocked all")

		// can re-acquire now
		err = pool.Lock(ctx, 0)
		require.NoError(t, err)
		err = pool.Unlock(ctx, 0)
		require.NoError(t, err)

		t.Log("re-acquired")
	})

	t.Run("multiple threads", func(t *testing.T) {
		var wg sync.WaitGroup
		cond := sync.NewCond(&sync.Mutex{})
		ch := make(chan struct{})

		for i := 0; i < 4; i++ {
			j := i
			wg.Add(1)
			go func() {
				ctx := pool.RegisterThread(ctx)

				t.Logf("thread %d locking %d", j, j)
				err := pool.Lock(ctx, j)
				require.NoError(t, err)

				ch <- struct{}{}

				cond.L.Lock()
				cond.Wait()
				cond.L.Unlock()

				t.Logf("thread %d locking %d", j, j+1)
				err = pool.Lock(ctx, j+1)
				require.NoError(t, err)

				err = pool.Unlock(ctx, j+1)
				require.NoError(t, err)

				err = pool.Unlock(ctx, j)
				require.NoError(t, err)

				wg.Done()
			}()
		}

		t.Log("locking up")

		ctx := pool.RegisterThread(ctx)
		err := pool.Lock(ctx, 4)
		require.NoError(t, err)

		t.Log("all locked")

		for i := 0; i < 4; i++ {
			<-ch
		}

		cond.Broadcast()
		<-time.After(time.Second)

		t.Log("main thread locking 0")
		err = pool.Lock(ctx, 0)
		require.Error(t, err)
		assert.Equal(t, "deadlock detected", err.Error())

		t.Log("deadlock detected")

		err = pool.Unlock(ctx, 4)
		require.NoError(t, err)

		wg.Wait()
	})
}
