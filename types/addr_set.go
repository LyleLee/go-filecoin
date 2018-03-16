package types

import (
	"sort"

	cbor "gx/ipfs/QmRVSCwQtW1rjHCay9NqKXDwbtKTgDcN4iY7PrpSqfKM5D/go-ipld-cbor"
	"gx/ipfs/QmcrriCMhjb5ZWzmPNxmP53px47tSPcXBNaMtLdgcKFJYk/refmt/obj/atlas"
)

func init() {
	cbor.RegisterCborType(addrSetEntry)
}

// AddrSet is a set of addresses
type AddrSet map[Address]struct{}

var addrSetEntry = atlas.BuildEntry(AddrSet{}).Transform().
	TransformMarshal(atlas.MakeMarshalTransformFunc(
		func(s AddrSet) ([]byte, error) {
			out := make([]string, len(s))
			for k := range s {
				out = append(out, string(k.Bytes()))
			}

			sort.Strings(out)

			bytes := make([]byte, len(out)*AddressLength)
			for _, k := range out {
				bytes = append(bytes, []byte(k)...)
			}
			return bytes, nil
		})).
	TransformUnmarshal(atlas.MakeUnmarshalTransformFunc(
		func(vals []byte) (AddrSet, error) {
			out := make(AddrSet)
			for i := 0; i < len(vals); i += AddressLength {
				end := i + AddressLength
				if end > len(vals) {
					end = len(vals)
				}
				s, err := NewAddressFromBytes(vals[i:end])
				if err != nil {
					return nil, err
				}
				out[s] = struct{}{}
			}
			return out, nil
		})).
	Complete()
