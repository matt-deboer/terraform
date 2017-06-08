package discovery

import (
	version "github.com/hashicorp/go-version"
)

// A VersionStr is a string containing a possibly-invalid representation
// of a semver version number. Call Parse on it to obtain a real Version
// object, or discover that it is invalid.
type VersionStr string

// Parse transforms a VersionStr into a Version if it is
// syntactically valid. If it isn't then an error is returned instead.
func (s VersionStr) Parse() (Version, error) {
	raw, err := version.NewVersion(string(s))
	if err != nil {
		return Version{}, err
	}
	return Version{raw}, nil
}

// MustParse transforms a VersionStr into a Version if it is
// syntactically valid. If it isn't then it panics.
func (s VersionStr) MustParse() Version {
	ret, err := s.Parse()
	if err != nil {
		panic(err)
	}
	return ret
}

// Version represents a version number that has been parsed from
// a semver string and known to be valid.
type Version struct {
	// We wrap this here just because it avoids a proliferation of
	// direct go-version imports all over the place, and keeps the
	// version-processing details within this package.
	raw *version.Version
}

func (v Version) String() string {
	return v.raw.String()
}

func (v Version) NewerThan(other Version) bool {
	return v.raw.GreaterThan(other.raw)
}
