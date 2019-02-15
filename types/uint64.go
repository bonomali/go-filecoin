package types

import (
	"encoding/base64"
	"strings"

	"gx/ipfs/QmNScbpMAm3r2D25kmfQ43JCbQ8QCtai4V4DNz5ebuXUuZ/refmt/obj/atlas"
	cbor "gx/ipfs/QmRZxJ7oybgnnwriuRub9JXp5YdFM9wiGSyRq38QC7swpS/go-ipld-cbor"
	"gx/ipfs/QmSKyB5faguXT4NqbrXpnRXqaVj5DhSm7x9BtzFydBY1UK/go-leb128"
)

func init() {
	cbor.RegisterCborType(uint64AtlasEntry)
}

var uint64AtlasEntry = atlas.BuildEntry(Uint64(0)).Transform().
	TransformMarshal(atlas.MakeMarshalTransformFunc(
		func(u Uint64) ([]byte, error) {
			return leb128.FromUInt64(uint64(u)), nil
		})).
	TransformUnmarshal(atlas.MakeUnmarshalTransformFunc(
		func(x []byte) (Uint64, error) {
			return Uint64(leb128.ToUInt64(x)), nil
		})).
	Complete()

// Uint64 is an unsigned 64-bit variable-length-encoded integer.
type Uint64 uint64

// MarshalJSON converts a Uint64 to a json string and returns it.
func (u Uint64) MarshalJSON() ([]byte, error) {
	encoded := base64.StdEncoding.EncodeToString(leb128.FromUInt64(uint64(u)))
	return []byte(`"` + encoded + `"`), nil
}

// UnmarshalJSON converts a json string to a Uint64.
func (u *Uint64) UnmarshalJSON(b []byte) error {
	jd, err := base64.StdEncoding.DecodeString(strings.Trim(string(b), `"`))
	if err != nil {
		return err
	}

	*u = Uint64(leb128.ToUInt64(jd))
	return nil
}
