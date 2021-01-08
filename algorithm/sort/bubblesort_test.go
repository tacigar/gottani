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

package sort

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	slice := sort.IntSlice{8, 2, 4, 12, 3, 6, 10, 2, 5}
	BubbleSort(slice)
	assert.Equal(t, sort.IntSlice{2, 2, 3, 4, 5, 6, 8, 10, 12}, slice)
}