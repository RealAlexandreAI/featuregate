package featuregate

const (
	ExampleFeature1 Feature = "ExampleFeature1"
	ExampleFeature2 Feature = "ExampleFeature2"
	ExampleFeature3 Feature = "ExampleFeature3"
)

// UserDefinedFeatureGate
//
//	Description:
//	return *FeatureGate
func UserDefinedFeatureGate() *FeatureGate {

	gate := NewFeatureGate()

	gate.AddFeatureSpec(ExampleFeature1, &BaseFeatureSpec{
		DefaultValue: true,
	})

	gate.AddFeatureSpec(ExampleFeature2, &BaseFeatureSpec{
		DefaultValue: false,
	})

	gate.AddFeatureSpec(ExampleFeature3, &BaseFeatureSpec{
		DefaultValue: false,
	})

	return gate
}
