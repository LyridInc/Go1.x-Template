package main

import (
	"fmt"

	"github.com/joho/godotenv"
	defaultFn "YOUR_APP_NAME.YOUR_MODULE_NAME/YOUR_FUNCTION_NAME"
)


// this function is not a part of the serverless and will not be processed by Lyrid,
// this is for user to be able to locally test and build their functions
func main() {

	// Loads env variables
    godotenv.Load()

	defaultFn.PreRun()

    defaultFunctionInput := defaultFn.LyFnInputParams{InputSample: "Hello"}
	defaultFunctionOutput := defaultFn.Run(defaultFunctionInput)

	fmt.Println("Function Input: " + defaultFunctionInput.InputSample)
	fmt.Println("Function Output: " + defaultFunctionOutput.OutputSample)

	defaultFn.PostRun()
}