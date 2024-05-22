package featuregate

import (
	"fmt"
	"testing"
)

func TestFeatureGate_example(t *testing.T) {

	gate1 := UserDefinedFeatureGate()
	gate1.SetFeatureState(ExampleFeature1, true)
	gate1.SetFeatureState(ExampleFeature2, true)

	features := gate1.KnownFeatures()

	assertEqual(features["ExampleFeature1"], true)
	assertEqual(features["ExampleFeature2"], true)

	//---

	gate2 := UserDefinedFeatureGate()
	gate2.SetFromMap(features)

	assertEqual(gate2.Enabled(ExampleFeature2), false)
	assertEqual(gate2.Enabled(ExampleFeature3), true)
}

func assertEqual(left, right bool) {
	if left != right {
		fmt.Printf("not match. left: %t, right: %t", left, right)
		fmt.Println()
	}
}
