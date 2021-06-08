package YOUR_FUNCTION_NAME

import "os"

// LyFnInputParams user fills up these parameters
// The struct name need to be static, but the internal composition of the struct can be changed to fit your usage
type LyFnInputParams struct {
    InputSample string
}

// LyFnOutputParams a struct that will be returned
// The struct name need to be static, but the internal composition of the struct can be changed to fit your usage
type LyFnOutputParams struct {
    OutputSample string
}

// PreRun (optional) will be executed prior to Run()
func PreRun() {
	
}

// Run will be the 
// The function name and definition need to be static, but the in
func Run(input LyFnInputParams) *LyFnOutputParams {
	response := LyFnOutputParams{OutputSample: input.InputSample + " " + os.Getenv("ENV1")}
	return &response
}

// PostRun (optional) will be executed after Run() is executed
func PostRun() {
	
}
