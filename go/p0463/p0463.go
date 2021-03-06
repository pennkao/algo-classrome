// Copyright (c) 2018 soren yang
//
// Licensed under the MIT License
// you may not use this file except in complicance with the License.
// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0463

func islandPerimeter(grid [][]int) int {
	perimeters := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				perimeters += 4
				if i > 0 {
					perimeters -= grid[i-1][j] * 2
				}
				if j > 0 {
					perimeters -= grid[i][j-1] * 2
				}
			}
		}
	}
	return perimeters
}
