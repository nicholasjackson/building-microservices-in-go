package exception

import (
	"errors"
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestCallerInfo(t *testing.T) {
	a := assert.New(t)

	stackTrace := callerInfo()
	a.NotEmpty(stackTrace)
}

func TestNew(t *testing.T) {
	a := assert.New(t)
	ex := AsException(New("this is a test"))
	a.Equal("this is a test", ex.Message())
	a.NotEmpty(ex.StackTrace())
	a.Nil(ex.InnerException())
}

func TestError(t *testing.T) {
	a := assert.New(t)

	ex := AsException(New("this is a test"))
	message := ex.Error()
	a.NotEmpty(message)
}

func TestNewf(t *testing.T) {
	a := assert.New(t)
	ex := AsException(Newf("%s", "this is a test"))
	a.Equal("this is a test", ex.Message())
	a.NotEmpty(ex.StackTrace())
	a.Nil(ex.InnerException())
}

func TestWrapError(t *testing.T) {
	a := assert.New(t)
	ex := AsException(WrapError(errors.New("this is a test")))
	a.Equal("this is a test", ex.Message())
	a.NotEmpty(ex.StackTrace())
	a.Nil(ex.InnerException())
}

func TestWrapWithError(t *testing.T) {
	a := assert.New(t)

	err := errors.New("This is an error")

	wrappedErr := Wrap(err)
	a.NotNil(wrappedErr)
	typedWrapped := AsException(wrappedErr)
	a.NotNil(typedWrapped)
	a.Equal("This is an error", typedWrapped.Message())
}

func TestWrapWithException(t *testing.T) {
	a := assert.New(t)
	ex := New("This is an exception")
	wrappedEx := Wrap(ex)
	a.NotNil(wrappedEx)
	typedWrappedEx := AsException(wrappedEx)
	a.Equal("This is an exception", typedWrappedEx.Message())
	a.Equal(ex, typedWrappedEx)
}

func TestWrapWithNil(t *testing.T) {
	a := assert.New(t)

	shouldBeNil := Wrap(nil)
	a.Nil(shouldBeNil)
	a.Equal(nil, shouldBeNil)
}

func TestWrapWithTypedNil(t *testing.T) {
	a := assert.New(t)

	var nilError error
	a.Nil(nilError)
	a.Equal(nil, nilError)

	shouldBeNil := Wrap(nilError)
	a.Nil(shouldBeNil)
	a.True(shouldBeNil == nil)
}

func TestWrapWithReturnedNil(t *testing.T) {
	a := assert.New(t)

	returnsNil := func() error {
		return nil
	}

	shouldBeNil := Wrap(returnsNil())
	a.Nil(shouldBeNil)
	a.True(shouldBeNil == nil)

	returnsTypedNil := func() error {
		return Wrap(nil)
	}

	shouldAlsoBeNil := returnsTypedNil()
	a.Nil(shouldAlsoBeNil)
	a.True(shouldAlsoBeNil == nil)
}

func TestWrapPrefixWithError(t *testing.T) {
	a := assert.New(t)

	err := errors.New("This is an error")

	wrappedErr := WrapPrefix(err, "This is a prefix")
	a.NotNil(wrappedErr)
	typedWrapped := AsException(wrappedErr)
	a.NotNil(typedWrapped)
	a.Equal("This is a prefix: This is an error", typedWrapped.Message())
}

func TestWrapPrefixWithException(t *testing.T) {
	a := assert.New(t)
	ex := New("This is an exception")
	wrappedEx := WrapPrefix(ex, "This is a prefix")
	a.NotNil(wrappedEx)
	typedWrappedEx := AsException(wrappedEx)
	a.Equal("This is a prefix: This is an exception", typedWrappedEx.Message())
}

func TestWrapPrefixWithNil(t *testing.T) {
	a := assert.New(t)

	shouldBeNil := WrapPrefix(nil, "This is a prefix")
	a.Nil(shouldBeNil)
	a.Equal(nil, shouldBeNil)
}

func TestWrapPrefixWithTypedNil(t *testing.T) {
	a := assert.New(t)

	var nilError error
	a.Nil(nilError)
	a.Equal(nil, nilError)

	shouldBeNil := WrapPrefix(nilError, "This is a prefix")
	a.Nil(shouldBeNil)
	a.True(shouldBeNil == nil)
}

func TestWrapPrefixWithReturnedNil(t *testing.T) {
	a := assert.New(t)

	returnsNil := func() error {
		return nil
	}

	shouldBeNil := WrapPrefix(returnsNil(), "This is a prefix")
	a.Nil(shouldBeNil)
	a.True(shouldBeNil == nil)

	returnsTypedNil := func() error {
		return Wrap(nil)
	}

	shouldAlsoBeNil := returnsTypedNil()
	a.Nil(shouldAlsoBeNil)
	a.True(shouldAlsoBeNil == nil)
}

func TestWrapMany(t *testing.T) {
	a := assert.New(t)

	err := errors.New("This is an error")
	ex1 := New("Exception1")
	ex2 := New("Exception2")

	combined := AsException(WrapMany(ex1, ex2, err))

	a.NotNil(combined)
	a.NotNil(combined.InnerException())
	a.NotNil(combined.InnerException().InnerException())
	a.NotEmpty(combined.Error())
}

func TestWrapManyWithCycle(t *testing.T) {
	a := assert.New(t)

	ex1 := New("This is an error")
	err := WrapMany(ex1, ex1)

	a.NotNil(err)
	a.NotEmpty(err.Error())

	typedException := AsException(err)
	a.Equal(ex1, typedException)
}

func TestWrapManyNil(t *testing.T) {
	a := assert.New(t)

	var ex1 error
	var ex2 error
	var ex3 error

	err := WrapMany(ex1, ex2, ex3)
	a.Nil(err)
	a.Equal(nil, err)
}
