//
// @copyright: 2022 by Pauli Rikula <pauli.rikula@gmail.com>
// @license: MIT <http://www.opensource.org/licenses/mit-license.php>
//

package gopareto

import (
	"fmt"
	"testing"
)

func TestComparisonChainExample(t *testing.T) {
	// [(2,2,2), (0,1,1), (0,0,1), (0,1,0), (1,0,0), (0,0,0)]
	values := [][]Item{
		{Int(2), Int(2), Int(2)},
		{Int(0), Int(1), Int(1)},
		{Int(0), Int(0), Int(1)},
		{Int(0), Int(1), Int(0)},
		{Int(1), Int(0), Int(0)},
		{Int(0), Int(0), Int(0)}}

	//fmt.Printf("%v\n", values)
	//  [[2 2 2] [0 1 1] [0 0 1] [0 1 0] [1 0 0] [0 0 0]]
	s := fmt.Sprintf("%v", values)
	if s != "[[2 2 2] [0 1 1] [0 0 1] [0 1 0] [1 0 0] [0 0 0]]" {
		t.Fatalf("value representation changed")
	}

	// TODO ...
}
