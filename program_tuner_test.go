package koza

import (
	"testing"

	"github.com/MaxHalford/koza/tree"
	"github.com/MaxHalford/koza/tree/op"
)

func TestSetProgConstants(t *testing.T) {
	var (
		prog = Program{
			Tree: &tree.Tree{
				Operator: op.Sum{},
				Branches: []*tree.Tree{
					&tree.Tree{Operator: op.Constant{1}},
					&tree.Tree{Operator: op.Constant{2}},
				},
			},
		}
		progTuner = newProgramTuner(&prog)
	)
	// Set new Constants
	for i, c := range progTuner.ConstValues {
		progTuner.ConstValues[i] = c + 1
	}
	progTuner.setProgConstants()
	// Check with the Program's Constants
	for i, branch := range progTuner.Program.Tree.Branches {
		if branch.Operator.(op.Constant).Value != prog.Tree.Branches[i].Operator.(op.Constant).Value+1 {
			t.Errorf("Expected %v, got %v", prog.Tree.Branches[i], branch.Operator)
		}
	}
}

func TestJitterConstants(t *testing.T) {
	var (
		prog = Program{
			Tree: &tree.Tree{
				Operator: op.Sum{},
				Branches: []*tree.Tree{
					&tree.Tree{Operator: op.Constant{1}},
					&tree.Tree{Operator: op.Constant{2}},
				},
			},
		}
		progTuner = newProgramTuner(&prog)
	)
	// Jitter Constants
	progTuner.jitterConstants(newRand())
	// Compare with the Program's Constants
	for i, c := range progTuner.ConstValues {
		if c == prog.Tree.Branches[i].Operator.(op.Constant).Value {
			t.Errorf("Expected %v and %v to be different", prog.Tree.Branches[i], c)
		}
	}
}
