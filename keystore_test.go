// Copyright 2019 Thales e-Security, Inc
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package gose

import (
	"testing"

	"github.com/ThalesIgnite/gose/jose"
	"github.com/stretchr/testify/assert"
)

func BenchmarkNewTrustKeyStore(b *testing.B) {
	// Setup
	b.Helper()
	keys := map[string]jose.Jwk{
		"issuer": &jose.PublicRsaKey{},
	}
	for _, key := range keys {
		key.SetKid("123456")
	}
	for i := 0; i < b.N; i++ {
		_, _ = NewTrustKeyStore(keys)

	}
}

func TestNewTrustKeyStore(t *testing.T) {
	// Setup
	keys := map[string]jose.Jwk{
		"issuer": &jose.PublicRsaKey{},
	}
	for _, key := range keys {
		key.SetKid("123456")
	}
	// Act
	store, err := NewTrustKeyStore(keys)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, store)
}