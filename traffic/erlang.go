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

// ErlangB calculates erlang B formula
// See: https://en.wikipedia.org/wiki/Erlang_(unit)
func ErlangB(a float64, n int) float64 {
	fs := make([]uint64, n+1, n+1)

	var factorial func(i int) uint64
	factorial = func(i int) uint64 {
		if i <= 1 {
			return 1
		}
		if fs[i] != 0 {
			return fs[i]
		}
		fs[i] = uint64(i) * factorial(i-1)
		return fs[i]
	}

	es := make([]float64, n+1, n+1)
	var exponent func(i int) float64
	exponent = func(i int) float64 {
		if i == 0 {
			return 1
		}
		if es[i] != 0 {
			return es[i]
		}
		es[i] = a * exponent(i-1)
		return es[i]
	}

	denominator := 0.0
	for i := 0; i <= n; i++ {
		denominator += exponent(i) / float64(factorial(i))
	}
	return (exponent(n) / float64(factorial(n))) / denominator
}
