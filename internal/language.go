// Package internal/interpreter - Language specification Interface
package internal

// Languager is the interface for creating a new language. Any object which conforms to this
// can be used as a language.
type Languager interface {
	ParseFile(filename string) (Ast, *HalError)
	CompileToFile(ast Ast, outputFileName string) *HalError
	Display() Displayer
}
