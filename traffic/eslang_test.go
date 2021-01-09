// Copyright 2021 Tacigar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package traffic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErlangB(t *testing.T) {
	d := 0.001
	assert.InDelta(t, 0.001, ErlangB(0.001, 1), d)
	assert.InDelta(t, 0.002, ErlangB(0.0653, 2), d)
	assert.InDelta(t, 0.03, ErlangB(1.26, 4), d)
	assert.InDelta(t, 0.1, ErlangB(3.76, 6), d)
}
