// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0874

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   [][]int
	target int
}

var values = []result{
	{
		arg1:   []int{4, -1, 3},
		arg2:   [][]int{},
		target: 25,
	},
	{
		arg1: []int{4, -1, 4, -2, 4},
		arg2: [][]int{
			[]int{2, 4},
		},
		target: 65,
	},
	{
		arg1: []int{-2, 8, 3, 7, -1},
		arg2: [][]int{
			[]int{-4, -1},
			[]int{1, -1},
			[]int{1, 4},
			[]int{5, 0},
			[]int{4, 5},
			[]int{-2, -1},
			[]int{2, -5},
			[]int{5, 1},
			[]int{-3, -1},
			[]int{5, -3},
		},
		target: 324,
	},
	{
		arg1: []int{-2, -1, -2, 3, 7},
		arg2: [][]int{
			[]int{1, -3},
			[]int{2, -3},
			[]int{4, 0},
			[]int{-2, 5},
			[]int{-5, 2},
			[]int{0, 0},
			[]int{4, -4},
			[]int{-2, -5},
			[]int{-1, -2},
			[]int{0, 2},
		},
		target: 100,
	},
	{
		arg1: []int{1, 2, -2, 5, -1, -2, -1, 8, 3, -1, 9, 4, -2, 3, 2, 4, 3, 9, 2, -1, -1, -2, 1, 3, -2, 4, 1, 4, -1, 1, 9, -1, -2, 5, -1, 5, 5, -2, 6, 6, 7, 7, 2, 8, 9, -1, 7, 4, 6, 9, 9, 9, -1, 5, 1, 3, 3, -1, 5, 9, 7, 4, 8, -1, -2, 1, 3, 2, 9, 3, -1, -2, 8, 8, 7, 5, -2, 6, 8, 4, 6, 2, 7, 2, -1, 7, -2, 3, 3, 2, -2, 6, 9, 8, 1, -2, -1, 1, 4, 7},
		arg2: [][]int{[]int{-57, -58},
			[]int{-72, 91},
			[]int{-55, 35},
			[]int{-20, 29},
			[]int{51, 70},
			[]int{-61, 88},
			[]int{-62, 99},
			[]int{52, 17},
			[]int{-75, -32},
			[]int{91, -22},
			[]int{54, 33},
			[]int{-45, -59},
			[]int{47, -48},
			[]int{53, -98},
			[]int{-91, 83},
			[]int{81, 12},
			[]int{-34, -90},
			[]int{-79, -82},
			[]int{-15, -86},
			[]int{-24, 66},
			[]int{-35, 35},
			[]int{3, 31},
			[]int{87, 93},
			[]int{2, -19},
			[]int{87, -93},
			[]int{24, -10},
			[]int{84, -53},
			[]int{86, 87},
			[]int{-88, -18},
			[]int{-51, 89},
			[]int{96, 66},
			[]int{-77, -94},
			[]int{-39, -1},
			[]int{89, 51},
			[]int{-23, -72},
			[]int{27, 24},
			[]int{53, -80},
			[]int{52, -33},
			[]int{32, 4},
			[]int{78, -55},
			[]int{-25, 18},
			[]int{-23, 47},
			[]int{79, -5},
			[]int{-23, -22},
			[]int{14, -25},
			[]int{-11, 69},
			[]int{63, 36},
			[]int{35, -99},
			[]int{-24, 82},
			[]int{-29, -98},
			[]int{-50, -70},
			[]int{72, 95},
			[]int{80, 80},
			[]int{-68, -40},
			[]int{65, 70},
			[]int{-92, 78},
			[]int{-45, -63},
			[]int{1, 34},
			[]int{81, 50},
			[]int{14, 91},
			[]int{-77, -54},
			[]int{13, -88},
			[]int{24, 37},
			[]int{-12, 59},
			[]int{-48, -62},
			[]int{57, -22},
			[]int{-8, 85},
			[]int{48, 71},
			[]int{12, 1},
			[]int{-20, 36},
			[]int{-32, -14},
			[]int{39, 46},
			[]int{-41, 75},
			[]int{13, -23},
			[]int{98, 10},
			[]int{-88, 64},
			[]int{50, 37},
			[]int{-95, -32},
			[]int{46, -91},
			[]int{10, 79},
			[]int{-11, 43},
			[]int{-94, 98},
			[]int{79, 42},
			[]int{51, 71},
			[]int{4, -30},
			[]int{2, 74},
			[]int{4, 10},
			[]int{61, 98},
			[]int{57, 98},
			[]int{46, 43},
			[]int{-16, 72},
			[]int{53, -69},
			[]int{54, -96},
			[]int{22, 0},
			[]int{-7, 92},
			[]int{-69, 80},
			[]int{68, -73},
			[]int{-24, -92},
			[]int{-21, 82},
			[]int{32, -1},
			[]int{-6, 16},
			[]int{15, -29},
			[]int{70, -66},
			[]int{-85, 80},
			[]int{50, -3},
			[]int{6, 13},
			[]int{-30, -98},
			[]int{-30, 59},
			[]int{-67, 40},
			[]int{17, 72},
			[]int{79, 82},
			[]int{89, -100},
			[]int{2, 79},
			[]int{-95, -46},
			[]int{17, 68},
			[]int{-46, 81},
			[]int{-5, -57},
			[]int{7, 58},
			[]int{-42, 68},
			[]int{19, -95},
			[]int{-17, -76},
			[]int{81, -86},
			[]int{79, 78},
			[]int{-82, -67},
			[]int{6, 0},
			[]int{35, -16},
			[]int{98, 83},
			[]int{-81, 100},
			[]int{-11, 46},
			[]int{-21, -38},
			[]int{-30, -41},
			[]int{86, 18},
			[]int{-68, 6},
			[]int{80, 75},
			[]int{-96, -44},
			[]int{-19, 66},
			[]int{21, 84},
			[]int{-56, -64},
			[]int{39, -15},
			[]int{0, 45},
			[]int{-81, -54},
			[]int{-66, -93},
			[]int{-4, 2},
			[]int{-42, -67},
			[]int{-15, -33},
			[]int{1, -32},
			[]int{-74, -24},
			[]int{7, 18},
			[]int{-62, 84},
			[]int{19, 61},
			[]int{39, 79},
			[]int{60, -98},
			[]int{-76, 45},
			[]int{58, -98},
			[]int{33, 26},
			[]int{-74, -95},
			[]int{22, 30},
			[]int{-68, -62},
			[]int{-59, 4},
			[]int{-62, 35},
			[]int{-78, 80},
			[]int{-82, 54},
			[]int{-42, 81},
			[]int{56, -15},
			[]int{32, -19},
			[]int{34, 93},
			[]int{57, -100},
			[]int{-1, -87},
			[]int{68, -26},
			[]int{18, 86},
			[]int{-55, -19},
			[]int{-68, -99},
			[]int{-9, 47},
			[]int{24, 94},
			[]int{92, 97},
			[]int{5, 67},
			[]int{97, -71},
			[]int{63, -57},
			[]int{-52, -14},
			[]int{-86, -78},
			[]int{-17, 92},
			[]int{-61, -83},
			[]int{-84, -10},
			[]int{20, 13},
			[]int{-68, -47},
			[]int{7, 28},
			[]int{66, 89},
			[]int{-41, -17},
			[]int{-14, -46},
			[]int{-72, -91},
			[]int{4, 52},
			[]int{-17, -59},
			[]int{-85, -46},
			[]int{-94, -23},
			[]int{-48, -3},
			[]int{-64, -37},
			[]int{2, 26},
			[]int{76, 88},
			[]int{-8, -46},
			[]int{-19, -68},
		},
		target: 5140,
	},
}

type p0874TestSuite struct {
	suite.Suite
}

func (s *p0874TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, robotSim(v.arg1, v.arg2))
	}
}

func TestP0874TestSuite(t *testing.T) {
	s := &p0874TestSuite{}
	suite.Run(t, s)
}