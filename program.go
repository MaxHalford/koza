package xgp

import (
	"github.com/MaxHalford/xgp/tree"
)

// A Program is simply an abstraction of top of a Tree.
type Program struct {
	Tree      *tree.Tree             `json:"tree"`
	Task      Task                   `json:"task"`
	DRS       *DynamicRangeSelection `json:"drs"`
	Estimator *Estimator             `json:"-"`
}

// String representation of a Program.
func (prog Program) String() string {
	return prog.Tree.String()
}

// Clone a Program.
func (prog Program) clone() *Program {
	var clone = &Program{
		Tree:      prog.Tree.Clone(),
		Task:      prog.Task,
		Estimator: prog.Estimator,
	}
	if prog.DRS != nil {
		clone.DRS = prog.DRS.clone()
	}
	return clone
}

// Predict predicts the output of a slice of features.
func (prog *Program) Predict(X [][]float64, predictProba bool) (yPred []float64, err error) {
	// Check if a cache exists
	var cache *tree.Cache
	if prog.Estimator != nil {
		cache = prog.Estimator.cache
	}
	yPred, err = prog.Tree.EvaluateCols(X, cache)
	if err != nil {
		return nil, err
	}
	// Binary classification
	if prog.Task.binaryClassification() {
		if predictProba {
			return sigmoid(yPred), nil
		}
		return binary(yPred), nil
	}
	// Multi-class classification
	if prog.Task.multiClassClassification() {
		return prog.DRS.Predict(yPred), nil
	}
	// Regression
	return yPred, nil
}
