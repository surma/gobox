package appletflag

var registered_packages = map[string]*FlagSet{}

// Equivalent to os.Argv, i.e. including the call name
var Parameters []string = []string{}

func getPackageFlagSet(name string) *FlagSet {
	flagset, ok := registered_packages[name]
	if !ok {
		flagset = NewFlagSet(name, ExitOnError)
		registered_packages[name] = flagset
	}
	return flagset
}
// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func Float64Var(p *float64, name string, value float64, usage string) {
	fs := getPackageFlagSet(getCallerPackage())
	fs.Float64Var(p, name, value, usage)
}

// Float64 defines an int flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func Float64(name string, value float64, usage string) *float64 {
	fs := getPackageFlagSet(getCallerPackage())
	return fs.Float64(name, value, usage)
}

// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
func Var(value Value, name string, usage string) {
	fs := getPackageFlagSet(getCallerPackage())
	fs.Var(value, name, usage)
}

// Parse parses the command-line flags from os.Args[1:].  Must be called
// after all flags are defined and before flags are accessed by the program.
func Parse() {
	fs := getPackageFlagSet(getCallerPackage())
	if len(Parameters) > 1 {
		fs.Parse(Parameters[1:])
	}
}

// Parsed returns true if the command-line flags have been parsed.
func Parsed() bool {
	fs := getPackageFlagSet(getCallerPackage())
	return fs.Parsed()
}

// VisitAll visits the command-line flags in lexicographical order, calling
// fn for each.  It visits all flags, even those not set.
func VisitAll(fn func(*Flag)) {
	fs := getPackageFlagSet(getCallerPackage())
	fs.VisitAll(fn)
}

// Visit visits the command-line flags in lexicographical order, calling fn
// for each.  It visits only those flags that have been set.
func Visit(fn func(*Flag)) {
	fs := getPackageFlagSet(getCallerPackage())
	fs.Visit(fn)
}

// Lookup returns the Flag structure of the named command-line flag,
// returning nil if none exists.
func Lookup(name string) *Flag {
	fs := getPackageFlagSet(getCallerPackage())
	return fs.Lookup(name)
}

// Set sets the value of the named command-line flag. It returns true if the
// set succeeded; false if there is no such flag defined.
func Set(name, value string) bool {
	fs := getPackageFlagSet(getCallerPackage())
	return fs.Set(name, value)
}

// PrintDefaults prints to standard error the default values of all defined command-line flags.
func PrintDefaults() {
	fs := getPackageFlagSet(getCallerPackage())
	fs.PrintDefaults()
}

// NFlag returns the number of command-line flags that have been set.
func NFlag() int {
	fs := getPackageFlagSet(getCallerPackage())
	return fs.NFlag()
}

// Arg returns the i'th command-line argument.  Arg(0) is the first remaining argument
// after flags have been processed.
func Arg(i int) string {
	fs := getPackageFlagSet(getCallerPackage())
	return fs.Arg(i)
}

// NArg is the number of arguments remaining after flags have been processed.
func NArg() int {
	fs := getPackageFlagSet(getCallerPackage())
	return fs.NArg()
}

// Args returns the non-flag command-line arguments.
func Args() []string {
	fs := getPackageFlagSet(getCallerPackage())
	return fs.Args()
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func BoolVar(p *bool, name string, value bool, usage string) {
	fs := getPackageFlagSet(getCallerPackage())
	fs.BoolVar(p, name, value, usage)
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func Bool(name string, value bool, usage string) *bool {
	fs := getPackageFlagSet(getCallerPackage())
	return fs.Bool(name, value, usage)
}

// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func IntVar(p *int, name string, value int, usage string) {
	fs := getPackageFlagSet(getCallerPackage())
	fs.IntVar(p, name, value, usage)
}

// Int defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func Int(name string, value int, usage string) *int {
	fs := getPackageFlagSet(getCallerPackage())
	return fs.Int(name, value, usage)
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func Int64Var(p *int64, name string, value int64, usage string) {
	fs := getPackageFlagSet(getCallerPackage())
	fs.Int64Var(p, name, value, usage)
}

// Int64 defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func Int64(name string, value int64, usage string) *int64 {
	fs := getPackageFlagSet(getCallerPackage())
	return fs.Int64(name, value, usage)
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint  variable in which to store the value of the flag.
func UintVar(p *uint, name string, value uint, usage string) {
	fs := getPackageFlagSet(getCallerPackage())
	fs.UintVar(p, name, value, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func Uint(name string, value uint, usage string) *uint {
	fs := getPackageFlagSet(getCallerPackage())
	return fs.Uint(name, value, usage)
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func Uint64Var(p *uint64, name string, value uint64, usage string) {
	fs := getPackageFlagSet(getCallerPackage())
	fs.Uint64Var(p, name, value, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func Uint64(name string, value uint64, usage string) *uint64 {
	fs := getPackageFlagSet(getCallerPackage())
	return fs.Uint64(name, value, usage)
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringVar(p *string, name string, value string, usage string) {
	fs := getPackageFlagSet(getCallerPackage())
	fs.StringVar(p, name, value, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func String(name string, value string, usage string) *string {
	fs := getPackageFlagSet(getCallerPackage())
	return fs.String(name, value, usage)
}
