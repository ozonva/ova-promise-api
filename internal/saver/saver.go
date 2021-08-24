package saver

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/ozonva/ova-promise-api/internal/domain"
	"github.com/ozonva/ova-promise-api/internal/usecase"
)

type Saver interface {
	Save(promise domain.Promise) error
	Init()
	Close()
}

type saver struct {
	sync.Mutex
	ticker         *time.Ticker
	tickerInterval uint64
	ch             chan interface{}
	buffer         []domain.Promise
	bufferCapacity int
	ucHandler      usecase.Handler
	logger         *zap.Logger
	ctx            context.Context
}

//nolint:gocritic //saver need clone of struct
func (s *saver) Save(promise domain.Promise) error {
	s.Lock()
	defer s.Unlock()

	if len(s.buffer) >= s.bufferCapacity {
		s.logger.Info("trying to flush buffer")
		s.flush()

		if len(s.buffer) >= s.bufferCapacity {
			return ErrFullBuffer
		}
	}

	s.buffer = append(s.buffer, promise)

	return nil
}

func (s *saver) Close() {
	s.flush()
	close(s.ch)
}

func NewSaver(ctx context.Context, tickerInterval uint64, capacity int, ucHandler usecase.Handler) Saver {
	return &saver{
		tickerInterval: tickerInterval,
		ucHandler:      ucHandler,
		buffer:         make([]domain.Promise, 0, capacity),
		bufferCapacity: capacity,
		ctx:            ctx,
	}
}

func (s *saver) Init() {
	s.ticker = time.NewTicker(time.Duration(s.tickerInterval) * time.Second)
	s.ch = make(chan interface{})

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func(ch <-chan interface{}) {
		wg.Done()

		for {
			select {
			case _, ok := <-ch:
				if !ok {
					s.ticker.Stop()

					break
				}
			case <-s.ticker.C:
				s.flush()
			}
		}
	}(s.ch)

	wg.Wait()
}

func (s *saver) flush() {
	s.Lock()
	defer s.Unlock()

	if len(s.buffer) > 0 {
		s.buffer = s.ucHandler.Flush(s.ctx, s.buffer)

		if len(s.buffer) > 0 {
			s.logger.Warn("unsaved data", zap.Any("promises", s.buffer))
		}
	}
}
