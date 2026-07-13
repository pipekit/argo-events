/*
Copyright 2020 The Argoproj Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package customtrigger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBearerToken_GetRequestMetadata(t *testing.T) {
	t.Run("defaults to the authorization header", func(t *testing.T) {
		b := bearerToken{token: "my-secret-token", requireTransportSecurity: true}

		md, err := b.GetRequestMetadata(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, map[string]string{"authorization": "Bearer my-secret-token"}, md)
	})

	t.Run("uses a custom header, lowercased", func(t *testing.T) {
		b := bearerToken{header: "X-Api-Key", token: "my-secret-token"}

		md, err := b.GetRequestMetadata(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, map[string]string{"x-api-key": "Bearer my-secret-token"}, md)
	})
}

func TestBearerToken_RequireTransportSecurity(t *testing.T) {
	assert.True(t, bearerToken{requireTransportSecurity: true}.RequireTransportSecurity())
	assert.False(t, bearerToken{requireTransportSecurity: false}.RequireTransportSecurity())
}
