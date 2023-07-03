/*
Copyright 2023 Bill Nixon

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

package main

import "fmt"

func min(v int, va ...int) int {
	minValue := v

	fmt.Printf("min(%v", v)
	for _, n := range va {
		fmt.Printf(", %d", n)
		if n < minValue {
			minValue = n
		}
	}
	fmt.Printf(") = %v\n", minValue)

	return minValue
}

func main() {
	min(1)
	min(1, 0, 2)
	min(1, 2, 3)
	min(1, 2, 3, 4)
	min(4, 3, 2, 1)
	min(3, 1, 0, 1)
	min(1, 2, -1, 3)
}
