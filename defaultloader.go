package envcfg

import "fmt"

// this file ensures that a default loader is created and available on the package, so users with
// simple cases can just do envcfg.Load.

var defaultLoader *Loader

func init() {
	// we can only fail here if one of the hardcoded default parsers has the wrong function
	// signature.  If that does fail, fail hard.
	var err error
	defaultLoader, err = New()
	if err != nil {
		panic(fmt.Sprintf("could not init default loader: %v", err))
	}
}

// Load loads config from the environment into the provided struct.
func Load(c interface{}) error {
	return defaultLoader.LoadWithDefaults(c)
}

// Load loads config from the environment into the provided struct.
func LoadWithDefaults(c interface{}) error {
	return defaultLoader.LoadWithDefaults(c)
}

// Load loads config from the environment into the provided struct.
func LoadWithoutDefaults(c interface{}) error {
	return defaultLoader.LoadWithoutDefaults(c)
}


// LoadFromMap loads config from the provided map into the provided struct.
func LoadFromMap(vals map[string]string, c interface{}) error {
	return defaultLoader.LoadFromMap(vals, c, true)
}

// RegisterParser takes a func (string) (<anytype>, error) and registers it on the default loader
// as the parser for <anytype>.
func RegisterParser(f interface{}) error {
	return defaultLoader.RegisterParser(f)
}
