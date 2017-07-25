package broker

import (
	"context"
	"errors"
	"testing"
	"time"

	fakeAPI "github.com/Azure/azure-service-broker/pkg/api/fake"
	fakeAsync "github.com/Azure/azure-service-broker/pkg/async/fake"
	"github.com/stretchr/testify/assert"
)

var errSome = errors.New("an error")

func TestBrokerStartBlocksUntilAsyncEngineErrors(t *testing.T) {
	var apiServerStopped bool
	svr := fakeAPI.NewServer()
	svr.RunBehavior = func(ctx context.Context) error {
		<-ctx.Done()
		apiServerStopped = true
		return ctx.Err()
	}
	e := fakeAsync.NewEngine()
	e.RunBehavior = func(context.Context) error {
		return errSome
	}
	b, err := NewBroker(nil, nil, nil)
	assert.Nil(t, err)
	brk := b.(*broker)
	brk.asyncEngine = e
	brk.apiServer = svr
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = b.Start(ctx)
	assert.Equal(t, &errAsyncEngineStopped{err: errSome}, err)
	time.Sleep(time.Second)
	assert.True(t, apiServerStopped)
}

func TestBrokerStartBlocksUntilAsyncEngineReturns(t *testing.T) {
	var apiServerStopped bool
	svr := fakeAPI.NewServer()
	svr.RunBehavior = func(ctx context.Context) error {
		<-ctx.Done()
		apiServerStopped = true
		return ctx.Err()
	}
	e := fakeAsync.NewEngine()
	e.RunBehavior = func(context.Context) error {
		return nil
	}
	b, err := NewBroker(nil, nil, nil)
	assert.Nil(t, err)
	brk := b.(*broker)
	brk.asyncEngine = e
	brk.apiServer = svr
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = b.Start(ctx)
	assert.Equal(t, &errAsyncEngineStopped{}, err)
	time.Sleep(time.Second)
	assert.True(t, apiServerStopped)
}

func TestBrokerStartBlocksUntilAPIServerErrors(t *testing.T) {
	svr := fakeAPI.NewServer()
	svr.RunBehavior = func(context.Context) error {
		return errSome
	}
	var asyncEngineStopped bool
	e := fakeAsync.NewEngine()
	e.RunBehavior = func(ctx context.Context) error {
		<-ctx.Done()
		asyncEngineStopped = true
		return ctx.Err()
	}
	b, err := NewBroker(nil, nil, nil)
	assert.Nil(t, err)
	brk := b.(*broker)
	brk.asyncEngine = e
	brk.apiServer = svr
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = b.Start(ctx)
	assert.Equal(t, &errAPIServerStopped{err: errSome}, err)
	time.Sleep(time.Second)
	assert.True(t, asyncEngineStopped)
}

func TestBrokerStartBlocksUntilAPIServerReturns(t *testing.T) {
	svr := fakeAPI.NewServer()
	svr.RunBehavior = func(context.Context) error {
		return nil
	}
	var asyncEngineStopped bool
	e := fakeAsync.NewEngine()
	e.RunBehavior = func(ctx context.Context) error {
		<-ctx.Done()
		asyncEngineStopped = true
		return ctx.Err()
	}
	b, err := NewBroker(nil, nil, nil)
	assert.Nil(t, err)
	brk := b.(*broker)
	brk.asyncEngine = e
	brk.apiServer = svr
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = b.Start(ctx)
	assert.Equal(t, &errAPIServerStopped{}, err)
	time.Sleep(time.Second)
	assert.True(t, asyncEngineStopped)
}

func TestBrokerStartBlocksUntilContextCanceled(t *testing.T) {
	var apiServerStopped bool
	svr := fakeAPI.NewServer()
	svr.RunBehavior = func(ctx context.Context) error {
		<-ctx.Done()
		apiServerStopped = true
		return ctx.Err()
	}
	var asyncEngineStopped bool
	e := fakeAsync.NewEngine()
	e.RunBehavior = func(ctx context.Context) error {
		<-ctx.Done()
		asyncEngineStopped = true
		return ctx.Err()
	}
	b, err := NewBroker(nil, nil, nil)
	assert.Nil(t, err)
	brk := b.(*broker)
	brk.asyncEngine = e
	brk.apiServer = svr
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = b.Start(ctx)
	assert.Equal(t, ctx.Err(), err)
	time.Sleep(time.Second)
	assert.True(t, apiServerStopped)
	assert.True(t, asyncEngineStopped)
}
