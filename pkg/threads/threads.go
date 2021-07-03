package threads

import (
	"context"
	"errors"
	"log"
	"sync"
)

type LockPool struct {
	mu      sync.Mutex
	locks   []lock
	threads []thread
}

type lock struct {
	mu      sync.Mutex
	threads map[int]bool
}

type thread struct {
	locks map[int]bool
}

func NewLockPool(size int) *LockPool {
	locks := make([]lock, size)
	for i := range locks {
		locks[i].threads = make(map[int]bool)
	}

	return &LockPool{locks: locks}
}

func (l *LockPool) RegisterThread(ctx context.Context) context.Context {
	l.mu.Lock()
	defer l.mu.Unlock()

	newThreadID := len(l.threads)
	l.threads = append(l.threads, thread{
		locks: make(map[int]bool),
	})
	return context.WithValue(ctx, "threadID", newThreadID)
}

func (l *LockPool) Lock(ctx context.Context, lockID int) error {
	if lockID < 0 || lockID >= len(l.locks) {
		return errors.New("invalid lock ID")
	}

	threadID, ok := ctx.Value("threadID").(int)
	if !ok {
		return errors.New("need to register thread")
	}

	if err := l.lockUpdateState(threadID, lockID); err != nil {
		return err
	}

	log.Print("no lock detected - locking ", lockID)

	l.locks[lockID].mu.Lock()
	return nil
}

func (l *LockPool) lockUpdateState(threadID, lockID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if err := l.checkDeadLock(threadID, threadID, lockID, nil); err != nil {
		return err
	}

	l.locks[lockID].threads[threadID] = true
	l.threads[threadID].locks[lockID] = true

	return nil
}

func (l *LockPool) Unlock(ctx context.Context, lockID int) error {
	if lockID < 0 || lockID >= len(l.locks) {
		return errors.New("invalid lock ID")
	}

	threadID, ok := ctx.Value("threadID").(int)
	if !ok {
		return errors.New("need to register thread")
	}

	if err := l.unlockUpdateState(threadID, lockID); err != nil {
		return err
	}

	l.locks[lockID].mu.Unlock()
	return nil
}

func (l *LockPool) unlockUpdateState(threadID, lockID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.locks[lockID].threads[threadID] = false
	l.threads[threadID].locks[lockID] = false

	return nil
}

func (l *LockPool) checkDeadLock(
	acquiringThreadID,
	checkThreadID,
	lockID int,
	threadsChecked map[int]bool,
) error {
	log.Printf("checking deadlock: acquiringThreadID=%d, checkThreadID=%d, lockID=%d", acquiringThreadID, checkThreadID, lockID)

	// check that this thread does not already hold the lock.
	if l.locks[lockID].threads[acquiringThreadID] {
		return errors.New("deadlock detected")
	}

	newThreadsChecked := make(map[int]bool, len(threadsChecked)+1)
	for t := range threadsChecked {
		newThreadsChecked[t] = true
	}
	newThreadsChecked[checkThreadID] = true

	// check that other threads waiting on this lock aren't blocked by this
	// thread.
	for thread, waiting := range l.locks[lockID].threads {
		if !waiting {
			continue
		}

		if threadsChecked[thread] {
			continue
		}

		for lock, ok := range l.threads[thread].locks {
			if !ok {
				continue
			}

			if lock == lockID {
				continue
			}

			err := l.checkDeadLock(acquiringThreadID, thread, lock, newThreadsChecked)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
