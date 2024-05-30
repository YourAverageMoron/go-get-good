package tree

import (
	"testing"
)

func TestPreOrderTraverse(t *testing.T) {
	lln := Node{
		value: 3,
	}
	lrn := Node{
		value: 4,
	}
	ln := Node{
		value: 2,
		left:  &lln,
		right: &lrn,
	}
	rln := Node{
		value: 6,
	}
	rrn := Node{
		value: 7,
	}
	rn := Node{
		value: 5,
		left:  &rln,
		right: &rrn,
	}
	root := Node{
		value: 1,
		left:  &ln,
		right: &rn,
	}
	PreOrderTraverse(&root)
}
