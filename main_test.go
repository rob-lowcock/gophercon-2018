package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rob-lowcock/gophercon-2018"
)

var _ = Describe("House calculator", func() {
	It("includes the original square", func() {
		Expect(main.CalculateHouses("")).To(Equal(1))
	})

	It("should respond to basic instructions", func() {
		Expect(main.CalculateHouses(">")).To(Equal(2))
	})

	It("can handle going in a square", func() {
		Expect(main.CalculateHouses("^>v<")).To(Equal(4))
	})

	It("can handle a deliverer who switches between two houses", func() {
		Expect(main.CalculateHouses("^v^v^v^v^v^v^v^v^v^v")).To(Equal(2))
	})

	It("can handle some vaguely complicated instructions", func() {
		Expect(main.CalculateHouses("<<^>vv>^<<v")).To(Equal(8))
	})

	It("can handle some extensive instructions", func() {
		Expect(main.CalculateHouses(">>>^^<v>v>>vvvvv<<<<<<^^^^>vv>><<^^<<v>^<v^^<<>>>^^^^>>vvvv<<>vv^^^^^<<^^<<vvvv<<vv^^^>>>>>vv")).
			To(Equal(65))
	})
})
