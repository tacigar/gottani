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

func TestPointF(t *testing.T) {
	{
		lhs := PointF{28.4, 42.3}
		rhs := PointF{12.9, 23.1}

		add := lhs.Add(rhs)
		assert.InDelta(t, 41.3, add.X, 0.00001)
		assert.InDelta(t, 65.4, add.Y, 0.00001)

		sub := lhs.Sub(rhs)
		assert.InDelta(t, 15.5, sub.X, 0.00001)
		assert.InDelta(t, 19.2, sub.Y, 0.00001)
	}
	{
		p := PointF{28.4, 42.3}
		p.AddAssign(PointF{3.2, 5.1})
		p.SubAssign(PointF{8.4, 19.2})
		assert.InDelta(t, 23.2, p.X, 0.00001)
		assert.InDelta(t, 28.2, p.Y, 0.00001)
	}
}
