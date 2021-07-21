package encrypt_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEncrypt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Encrypt Suite")
}
