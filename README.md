# FeatureGate

Split pkg/featuregate from [mosn](https://github.com/mosn/mosn/). 

Remove `init()`, and follow the original LICENSE too.

## Usage

```go
fg := NewFeatureGate()

fg.AddFeatureSpec("ExampleFeature", &BaseFeatureSpec{
    DefaultValue: true,
})

// ---
assertEqual(fg.Enabled("ExampleFeature"), true)
``` 

More details in `example.go` and `example_test.go`.
