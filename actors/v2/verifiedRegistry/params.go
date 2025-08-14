package verifiedRegistry

import (
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	verifregv10 "github.com/filecoin-project/go-state-types/builtin/v10/verifreg"
	verifregv11 "github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	verifregv12 "github.com/filecoin-project/go-state-types/builtin/v12/verifreg"
	verifregv13 "github.com/filecoin-project/go-state-types/builtin/v13/verifreg"
	verifregv14 "github.com/filecoin-project/go-state-types/builtin/v14/verifreg"
	verifregv15 "github.com/filecoin-project/go-state-types/builtin/v15/verifreg"
	verifregv16 "github.com/filecoin-project/go-state-types/builtin/v16/verifreg"
	verifregv8 "github.com/filecoin-project/go-state-types/builtin/v8/verifreg"
	verifregv9 "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"
	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/verifreg"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/verifreg"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/verifreg"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/verifreg"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/verifreg"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/verifreg"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/actors/v2/verifiedRegistry/types"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

// All methods can be found in the Actor.Exports method in
// the correct version package for "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"

func v1Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	v := &VerifiedRegistry{}
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		1: {
			Name:   parser.MethodConstructor,
			Method: actors.ParseConstructor,
		},
		2: {
			Name:   parser.MethodAddVerifier,
			Method: v.AddVerifier,
		},
		3: {
			Name:   parser.MethodRemoveVerifier,
			Method: v.RemoveVerifier,
		},
		4: {
			Name:   parser.MethodAddVerifiedClient,
			Method: v.AddVerifiedClientExported,
		},
		5: {
			Name:   parser.MethodUseBytes,
			Method: v.UseBytes,
		},
		6: {
			Name:   parser.MethodRestoreBytes,
			Method: v.RestoreBytes,
		},
	}
}

func v2Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v3Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v4Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v5Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v6Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v7Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	v := &VerifiedRegistry{}
	methods := v6Methods()
	methods[7] = nonLegacyBuiltin.MethodMeta{
		Name:   parser.MethodRemoveVerifiedClientDataCap,
		Method: v.RemoveVerifiedClientDataCap,
	}
	return methods
}

var addVerifierParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.AddVerifierParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.AddVerifierParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.AddVerifierParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.AddVerifierParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AddVerifierParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AddVerifierParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AddVerifierParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AddVerifierParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AddVerifierParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AddVerifierParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.AddVerifierParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.AddVerifierParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.AddVerifierParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.AddVerifierParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.AddVerifierParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.AddVerifierParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(verifregv8.AddVerifierParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.AddVerifierParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.AddVerifierParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.AddVerifierParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.AddVerifierParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.AddVerifierParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.AddVerifierParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.AddVerifierParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.AddVerifierParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.AddVerifierParams) },
}

var addVerifiedClientParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.AddVerifiedClientParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.AddVerifiedClientParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.AddVerifiedClientParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.AddVerifiedClientParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AddVerifiedClientParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AddVerifiedClientParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AddVerifiedClientParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AddVerifiedClientParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AddVerifiedClientParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AddVerifiedClientParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.AddVerifiedClientParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.AddVerifiedClientParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.AddVerifiedClientParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.AddVerifiedClientParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.AddVerifiedClientParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.AddVerifiedClientParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(verifregv8.AddVerifiedClientParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.AddVerifiedClientParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.AddVerifiedClientParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.AddVerifiedClientParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.AddVerifiedClientParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.AddVerifiedClientParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.AddVerifiedClientParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.AddVerifiedClientParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.AddVerifiedClientParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.AddVerifiedClientParams) },
}

var useBytesParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.UseBytesParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.UseBytesParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.UseBytesParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.UseBytesParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.UseBytesParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.UseBytesParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.UseBytesParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.UseBytesParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.UseBytesParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.UseBytesParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.UseBytesParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.UseBytesParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.UseBytesParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.UseBytesParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.UseBytesParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.UseBytesParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(verifregv8.UseBytesParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.UseBytesParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.UseBytesParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.UseBytesParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.UseBytesParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.UseBytesParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.UseBytesParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.UseBytesParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.UseBytesParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.UseBytesParams) },
}

var restoreBytesParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.RestoreBytesParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.RestoreBytesParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.RestoreBytesParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.RestoreBytesParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.RestoreBytesParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.RestoreBytesParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.RestoreBytesParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.RestoreBytesParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.RestoreBytesParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.RestoreBytesParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.RestoreBytesParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.RestoreBytesParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.RestoreBytesParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.RestoreBytesParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.RestoreBytesParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.RestoreBytesParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(verifregv8.RestoreBytesParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.RestoreBytesParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.RestoreBytesParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.RestoreBytesParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.RestoreBytesParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.RestoreBytesParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.RestoreBytesParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.RestoreBytesParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.RestoreBytesParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.RestoreBytesParams) },
}

var removedVerifiedClientDataCapParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.RemoveDataCapParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(verifregv8.RemoveDataCapParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.RemoveDataCapParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.RemoveDataCapParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.RemoveDataCapParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.RemoveDataCapParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.RemoveDataCapParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.RemoveDataCapParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.RemoveDataCapParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.RemoveDataCapParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.RemoveDataCapParams) },
}

var removedVerifiedClientDataCapReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.RemoveDataCapReturn) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.RemoveDataCapReturn) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.RemoveDataCapReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.RemoveDataCapReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.RemoveDataCapReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.RemoveDataCapReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.RemoveDataCapReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.RemoveDataCapReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.RemoveDataCapReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.RemoveDataCapReturn) },
}

var removeExpiredAllocationsParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.RemoveExpiredAllocationsParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.RemoveExpiredAllocationsParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.RemoveExpiredAllocationsParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.RemoveExpiredAllocationsParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.RemoveExpiredAllocationsParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.RemoveExpiredAllocationsParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.RemoveExpiredAllocationsParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.RemoveExpiredAllocationsParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.RemoveExpiredAllocationsParams) },
}

var removeExpiredAllocationsReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.RemoveExpiredAllocationsReturn) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.RemoveExpiredAllocationsReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.RemoveExpiredAllocationsReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.RemoveExpiredAllocationsReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.RemoveExpiredAllocationsReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.RemoveExpiredAllocationsReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.RemoveExpiredAllocationsReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.RemoveExpiredAllocationsReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.RemoveExpiredAllocationsReturn) },
}

var claimAllocationsParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.ClaimAllocationsParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.ClaimAllocationsParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.ClaimAllocationsParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.ClaimAllocationsParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.ClaimAllocationsParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.ClaimAllocationsParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.ClaimAllocationsParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.ClaimAllocationsParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.ClaimAllocationsParams) },
}

var claimAllocationsReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(types.ClaimAllocationsReturn) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(types.ClaimAllocationsReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(types.ClaimAllocationsReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(types.ClaimAllocationsReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(types.ClaimAllocationsReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(types.ClaimAllocationsReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(types.ClaimAllocationsReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(types.ClaimAllocationsReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(types.ClaimAllocationsReturn) },
}

var getClaimsParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.GetClaimsParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.GetClaimsParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.GetClaimsParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.GetClaimsParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.GetClaimsParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.GetClaimsParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.GetClaimsParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.GetClaimsParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.GetClaimsParams) },
}

var getClaimsReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.GetClaimsReturn) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.GetClaimsReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.GetClaimsReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.GetClaimsReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.GetClaimsReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.GetClaimsReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.GetClaimsReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.GetClaimsReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.GetClaimsReturn) },
}

var extendClaimTermsParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.ExtendClaimTermsParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.ExtendClaimTermsParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.ExtendClaimTermsParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.ExtendClaimTermsParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.ExtendClaimTermsParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.ExtendClaimTermsParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.ExtendClaimTermsParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.ExtendClaimTermsParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.ExtendClaimTermsParams) },
}

var extendClaimTermsReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.ExtendClaimTermsReturn) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.ExtendClaimTermsReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.ExtendClaimTermsReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.ExtendClaimTermsReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.ExtendClaimTermsReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.ExtendClaimTermsReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.ExtendClaimTermsReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.ExtendClaimTermsReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.ExtendClaimTermsReturn) },
}

var removeExpiredClaimsParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.RemoveExpiredClaimsParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.RemoveExpiredClaimsParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.RemoveExpiredClaimsParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.RemoveExpiredClaimsParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.RemoveExpiredClaimsParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.RemoveExpiredClaimsParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.RemoveExpiredClaimsParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.RemoveExpiredClaimsParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.RemoveExpiredClaimsParams) },
}

var removeExpiredClaimsReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.RemoveExpiredClaimsReturn) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.RemoveExpiredClaimsReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.RemoveExpiredClaimsReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.RemoveExpiredClaimsReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.RemoveExpiredClaimsReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.RemoveExpiredClaimsReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.RemoveExpiredClaimsReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.RemoveExpiredClaimsReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.RemoveExpiredClaimsReturn) },
}

var universalReceiverParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.UniversalReceiverParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.UniversalReceiverParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.UniversalReceiverParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.UniversalReceiverParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.UniversalReceiverParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.UniversalReceiverParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.UniversalReceiverParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.UniversalReceiverParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.UniversalReceiverParams) },
}

var allocationRequests = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(types.AllocationRequests) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.AllocationRequests) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.AllocationRequests) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.AllocationRequests) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.AllocationRequests) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.AllocationRequests) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.AllocationRequests) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.AllocationRequests) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.AllocationRequests) },
}

var allocationsResponse = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(verifregv9.AllocationsResponse) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(verifregv10.AllocationsResponse) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.AllocationsResponse) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(verifregv11.AllocationsResponse) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(verifregv12.AllocationsResponse) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(verifregv13.AllocationsResponse) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(verifregv14.AllocationsResponse) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(verifregv15.AllocationsResponse) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(verifregv16.AllocationsResponse) },
}

var VerifregTypes = map[string]map[string]func() cbg.CBORUnmarshaler{
	parser.MethodAddVerifiedClient:         addVerifiedClientParams,
	parser.MethodAddVerifiedClientExported: addVerifiedClientParams,
	parser.MethodAddVerifier:               addVerifierParams,
	//parser.MethodRemoveVerifier:                   &address.Address{},
	parser.MethodUseBytes:                         useBytesParams,
	parser.MethodRestoreBytes:                     restoreBytesParams,
	parser.MethodRemoveExpiredAllocations:         removeExpiredAllocationsParams,
	parser.MethodRemoveExpiredAllocationsExported: removeExpiredAllocationsParams,
	parser.MethodRemoveVerifiedClientDataCap:      removedVerifiedClientDataCapParams,
	parser.MethodVerifiedDeprecated1:              removedVerifiedClientDataCapParams,
	parser.MethodVerifiedDeprecated2:              removedVerifiedClientDataCapParams,
	parser.MethodGetClaims:                        getClaimsParams,
	parser.MethodGetClaimsExported:                getClaimsParams,
	parser.MethodExtendClaimTerms:                 extendClaimTermsParams,
	parser.MethodExtendClaimTermsExported:         extendClaimTermsParams,
	parser.MethodRemoveExpiredClaims:              removeExpiredClaimsParams,
	parser.MethodRemoveExpiredClaimsExported:      removeExpiredClaimsParams,
	parser.MethodUniversalReceiverHook:            universalReceiverParams,
	parser.MethodClaimAllocations:                 claimAllocationsParams,
}
