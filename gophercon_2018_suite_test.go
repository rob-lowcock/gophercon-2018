package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGophercon2018(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gophercon2018 Suite")
}
