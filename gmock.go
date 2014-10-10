// Mocking Framework

package gmock

type GMock struct {
	GMockFuncs []*MockRoutine
}

type MockRoutine struct {
	Name    string
	Args    []interface{}
	RetVals []interface{}
}

type GMockResult struct {
	Result []interface{}
}

//Returns the parameter as a bool
// It by default returns false
func (gr *GMockResult) Bool(i int) bool {

	if gr.Contains(i) {
		return gr.Result[i].(bool)
	} else {
		return false
	}

}

// Returns the parameter as float32
// Default returns 0
func (gr *GMockResult) Float32(i int) float32 {

	if gr.Contains(i) {
		return gr.Result[i].(float32)
	} else {
		return 0
	}

}

// Returns the parameter as float64
// Default returns 0
func (gr *GMockResult) Float64(i int) float64 {

	if gr.Contains(i) {
		return gr.Result[i].(float64)
	} else {
		return 0
	}
}

func (gr *GMockResult) Int(i int) int {

	if gr.Contains(i) {
		return gr.Result[i].(int)
	} else {
		return 0
	}
}

func (gr *GMockResult) Int8(i int) int8 {

	if gr.Contains(i) {
		return gr.Result[i].(int8)
	} else {
		return 0
	}
}

func (gr *GMockResult) Int32(i int) int32 {

	if gr.Contains(i) {

		return gr.Result[i].(int32)
	} else {

		return 0
	}

}

func (gr *GMockResult) Int64(i int) int64 {

	if gr.Contains(i) {

		return gr.Result[i].(int64)
	} else {

		return 0
	}
}

// Returns the parameter as a string
func (gr *GMockResult) String(i int) string {

	if gr.Contains(i) {
		return gr.Result[i].(string)
	} else {
		return ""
	}
}
