package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

func TestCacheStore(t *testing.T) {

	tests := []struct {
		testName     string
		checkAddFunc func(fn func(key, vlaue int))
		checkGetFunc func(fn func(key int) (value int, exist bool))
	}{
		{
			testName: "EmptyCacheStoreNotExistAnyValues",
			checkAddFunc: func(fn func(key, vlaue int)) {
			},
			checkGetFunc: func(fn func(key int) (value int, exist bool)) {
				value, ok := fn(12)
				assert.Zero(t, value)
				assert.False(t, ok)
			},
		},
		{
			testName: "AddedKeyValueShouldGetValueByKey",
			checkAddFunc: func(fn func(key, vlaue int)) {
				fn(12, 100)
			},
			checkGetFunc: func(fn func(key int) (value int, exist bool)) {

				value, ok := fn(12)
				assert.Equal(t, 100, value)
				assert.True(t, ok)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {

			addValue, getValue := CacheStore()

			test.checkAddFunc(addValue)
			test.checkGetFunc(getValue)
		})
	}
}
