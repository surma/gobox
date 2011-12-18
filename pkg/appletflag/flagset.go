package appletflag

import (
	"flag"
)

var ErrHelp = flag.ErrHelp

// Usage prints to standard error a usage message documenting all defined command-line flags.
// The function is a variable that may be changed to point to a custom function.
var Usage = flag.Usage

// A FlagSet represents a set of defined flags.
type FlagSet flag.FlagSet

// Value is the interface to the dynamic value stored in a flag.
// (The default value is represented as a string.)
type Value flag.Value

// ErrorHandling defines how to handle flag parsing errors.
type ErrorHandling flag.ErrorHandling

const (
	ContinueOnError ErrorHandling = ErrorHandling(flag.ContinueOnError)
	ExitOnError     ErrorHandling = ErrorHandling(flag.ExitOnError)
	PanicOnError    ErrorHandling = ErrorHandling(flag.PanicOnError)
)

// A Flag represents the state of a flag.
type Flag flag.Flag

// NewFlagSet returns a new, empty flag set with the specified name and
// error handling property.
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet {
	return (*FlagSet)(flag.NewFlagSet(name, flag.ErrorHandling(errorHandling)))
}

func FnWrapper(fn func(*Flag)) func(*flag.Flag) {
	return func(f *flag.Flag) {
		fn((*Flag)(f))
	}
}

// VisitAll visits the flags in lexicographical order, calling fn for each.
// It visits all flags, even those not set.
func (f *FlagSet) VisitAll(fn func(*Flag)) {
	(*flag.FlagSet)(f).VisitAll(FnWrapper(fn))
}

// Visit visits the flags in lexicographical order, calling fn for each.
// It visits only those flags that have been set.
func (f *FlagSet) Visit(fn func(*Flag)) {
	(*flag.FlagSet)(f).Visit(FnWrapper(fn))
}

// Lookup returns the Flag structure of the named flag, returning nil if none exists.
func (f *FlagSet) Lookup(name string) *Flag {
	return (*Flag)((*flag.FlagSet)(f).Lookup(name))
}

// Set sets the value of the named flag.  It returns true if the set succeeded; false if
// there is no such flag defined.
func (f *FlagSet) Set(name, value string) bool {
	return (*flag.FlagSet)(f).Set(name, value)
}

// PrintDefaults prints to standard error the default values of all defined flags in the set.
func (f *FlagSet) PrintDefaults() {
	(*flag.FlagSet)(f).PrintDefaults()
}

// NFlag returns the number of flags that have been set.
func (f *FlagSet) NFlag() int { return (*flag.FlagSet)(f).NFlag() }

// Arg returns the i'th argument.  Arg(0) is the first remaining argument
// after flags have been processed.
func (f *FlagSet) Arg(i int) string {
	return (*flag.FlagSet)(f).Arg(i)
}

// NArg is the number of arguments remaining after flags have been processed.
func (f *FlagSet) NArg() int { return (*flag.FlagSet)(f).NArg() }

// Args returns the non-flag arguments.
func (f *FlagSet) Args() []string { return (*flag.FlagSet)(f).Args() }

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string) {
	(*flag.FlagSet)(f).BoolVar(p, name, value, usage)
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func (f *FlagSet) Bool(name string, value bool, usage string) *bool {
	return (*flag.FlagSet)(f).Bool(name, value, usage)
}

// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func (f *FlagSet) IntVar(p *int, name string, value int, usage string) {
	(*flag.FlagSet)(f).IntVar(p, name, value, usage)
}

// Int defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func (f *FlagSet) Int(name string, value int, usage string) *int {
	return (*flag.FlagSet)(f).Int(name, value, usage)
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string) {
	(*flag.FlagSet)(f).Int64Var(p, name, value, usage)
}

// Int64 defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func (f *FlagSet) Int64(name string, value int64, usage string) *int64 {
	return (*flag.FlagSet)(f).Int64(name, value, usage)
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string) {
	(*flag.FlagSet)(f).UintVar(p, name, value, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (f *FlagSet) Uint(name string, value uint, usage string) *uint {
	return (*flag.FlagSet)(f).Uint(name, value, usage)
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string) {
	(*flag.FlagSet)(f).Uint64Var(p, name, value, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64 {
	return (*flag.FlagSet)(f).Uint64(name, value, usage)
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (f *FlagSet) StringVar(p *string, name string, value string, usage string) {
	(*flag.FlagSet)(f).StringVar(p, name, value, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (f *FlagSet) String(name string, value string, usage string) *string {
	return (*flag.FlagSet)(f).String(name, value, usage)
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string) {
	(*flag.FlagSet)(f).Float64Var(p, name, value, usage)
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func (f *FlagSet) Float64(name string, value float64, usage string) *float64 {
	return (*flag.FlagSet)(f).Float64(name, value, usage)
}

// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
func (f *FlagSet) Var(value Value, name string, usage string) {
	(*flag.FlagSet)(f).Var(flag.Value(value), name, usage)
}

// Parse parses flag definitions from the argument list, which should not
// include the command name.  Must be called after all flags in the FlagSet
// are defined and before flags are accessed by the program.
// The return value will be ErrHelp if -help was set but not defined.
func (f *FlagSet) Parse(arguments []string) error {
	return (*flag.FlagSet)(f).Parse(arguments)
}

// Parsed reports whether f.Parse has been called.
func (f *FlagSet) Parsed() bool {
	return (*flag.FlagSet)(f).Parsed()
}

// Init sets the name and error handling property for a flag set.
// By default, the zero FlagSet uses an empty name and the
// ContinueOnError error handling policy.
func (f *FlagSet) Init(name string, errorHandling ErrorHandling) {
	(*flag.FlagSet)(f).Init(name, flag.ErrorHandling(errorHandling))
}
