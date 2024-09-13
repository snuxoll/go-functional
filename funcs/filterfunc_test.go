package funcs_test

import (
	"go.snuxoll.com/functional/funcs"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//func TestFuncs(t *testing.T) {
//	RegisterFailHandler(Fail)
//	RunSpecs(t, "Funcs Suite")
//}

type filterTestFunc[T any] func(filter funcs.FilterFunc[T], called *bool)

func withTrueFilter(f filterTestFunc[int]) func() {
	called := false
	filter := funcs.AsFilter(func(i int) bool {
		called = true
		return true
	})
	return func() {
		f(filter, &called)
	}
}

func withFalseFilter(f filterTestFunc[int]) func() {
	called := false
	filter := funcs.AsFilter(func(i int) bool {
		called = true
		return false
	})
	return func() {
		f(filter, &called)
	}
}

func TestFilterFunc(t *testing.T) {
	Convey("Given a FilterFunc that always returns true", t, withTrueFilter(
		func(f1 funcs.FilterFunc[int], f1Called *bool) {
			Convey("combined with another that always returns true", withTrueFilter(
				func(f2 funcs.FilterFunc[int], f2Called *bool) {
					combined := f1.Combine(f2)

					Convey("should only call the second function when evaluated", func() {
						combined(1)
						So(*f1Called, ShouldBeFalse)
						So(*f2Called, ShouldBeTrue)
					})

					Reset(func() {
						*f2Called = false
					})
				}))
			Reset(func() {
				*f1Called = false
			})

			Convey("combined with another that always returns false", withFalseFilter(
				func(f2 funcs.FilterFunc[int], f2Called *bool) {
					combined := f1.Combine(f2)

					Convey("should call both functions", func() {
						combined(1)
						So(*f1Called, ShouldBeTrue)
						So(*f2Called, ShouldBeTrue)
					})

					Convey("should return true", func() {
						So(combined(1), ShouldBeTrue)
					})

					Reset(func() {
						*f2Called = false
					})
				}))
		}))

	Convey("Given a FilterFunc that always returns false", t, withFalseFilter(
		func(f1 funcs.FilterFunc[int], f1Called *bool) {
			Convey("combined with another that always returns false", withFalseFilter(
				func(f2 funcs.FilterFunc[int], f2Called *bool) {
					combined := f1.Combine(f2)

					Convey("should call both functions when evaluated", func() {
						combined(1)
						So(*f1Called, ShouldBeTrue)
						So(*f2Called, ShouldBeTrue)
					})

					Convey("should return false", func() {
						So(combined(1), ShouldBeFalse)
					})

					Reset(func() {
						*f2Called = false
					})
				}))

			Reset(func() {
				*f1Called = false
			})
		}))
}
