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

func BubbleSort(slice Interface) {
	for i := 0; i < slice.Len()-1; i++ {
		done := true

		for j := 1; j < slice.Len()-i; j++ {
			if slice.Less(j, j-1) {
				slice.Swap(j, j-1)
				done = false
			}
		}
		if done {
			return
		}
	}
}
