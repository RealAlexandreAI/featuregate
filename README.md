# FeatureGate

Split pkg/featuregate from [mosn](https://github.com/mosn/mosn/). 

Remove some `init()`, and follow the original LICENSE too.

## usage

```go
fg := NewFeatureGate()

fg.AddFeatureSpec("ExampleFeature", &BaseFeatureSpec{
    DefaultValue: true,
})

// ---
assertEqual(fg.Enabled("ExampleFeature"), true)
``` 

More details in `example_test.go`.
