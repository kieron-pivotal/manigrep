package traverser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTraverser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Traverser Suite")
}
