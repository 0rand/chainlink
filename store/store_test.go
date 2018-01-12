package store_test

import (
	"syscall"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/smartcontractkit/chainlink-go/internal/cltest"
)

func TestGracefulShutdown(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	store, cleanup := cltest.NewStore()
	defer cleanup()

	var completed bool
	store.Exiter = func(code int) {
		completed = true
	}

	store.Start()
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)

	Eventually(func() bool {
		return completed
	}).Should(BeTrue())
}