## Learning GO

---

### Value type on GO

int, float, string, bool, structs
Note: Use pointers to changes these in a function

### Reference types on GO

slices, map, channels, pointers, function
Note: No need to pass as pointer

### Notes

# Use of go.sum and go.mod

- go.mod is critical for defining your project's dependencies.
- go.sum ensures the dependencies are downloaded correctly and consistently.
- Do not remove them unless absolutely necessary and you know what you're doing. If you’re troubleshooting, regenerate them using go mod tidy.
