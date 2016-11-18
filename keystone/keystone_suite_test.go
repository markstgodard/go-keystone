package keystone_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKeystone(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keystone Suite")
}
