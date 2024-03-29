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

package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	{
		lhs := Point{28, 42}
		rhs := Point{12, 23}

		add := lhs.Add(rhs)
		assert.Equal(t, 40, add.X)
		assert.Equal(t, 65, add.Y)

		sub := lhs.Sub(rhs)
		assert.Equal(t, 16, sub.X)
		assert.Equal(t, 19, sub.Y)
	}
	{
		p := Point{28, 42}
		p.AddAssign(Point{3, 5})
		p.SubAssign(Point{8, 19})
		assert.Equal(t, 23, p.X)
		assert.Equal(t, 28, p.Y)
	}
}
