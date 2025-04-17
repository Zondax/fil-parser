package verifiedRegistry

import (
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
	"github.com/zondax/fil-parser/tools"
)

func addVerifierParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.AddVerifierParams{},

		tools.V8.String(): &legacyv2.AddVerifierParams{},
		tools.V9.String(): &legacyv2.AddVerifierParams{},

		tools.V10.String(): &legacyv3.AddVerifierParams{},
		tools.V11.String(): &legacyv3.AddVerifierParams{},

		tools.V12.String(): &legacyv4.AddVerifierParams{},
		tools.V13.String(): &legacyv5.AddVerifierParams{},
		tools.V14.String(): &legacyv6.AddVerifierParams{},
		tools.V15.String(): &legacyv7.AddVerifierParams{},
		tools.V16.String(): &verifregv8.AddVerifierParams{},
		tools.V17.String(): &verifregv9.AddVerifierParams{},
		tools.V18.String(): &verifregv10.AddVerifierParams{},

		tools.V19.String(): &verifregv11.AddVerifierParams{},
		tools.V20.String(): &verifregv11.AddVerifierParams{},

		tools.V21.String(): &verifregv12.AddVerifierParams{},
		tools.V22.String(): &verifregv13.AddVerifierParams{},
		tools.V23.String(): &verifregv14.AddVerifierParams{},
		tools.V24.String(): &verifregv15.AddVerifierParams{},
		tools.V25.String(): &verifregv16.AddVerifierParams{},
	}
}
func addVerifiedClientParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.AddVerifiedClientParams{},

		tools.V8.String(): &legacyv2.AddVerifiedClientParams{},
		tools.V9.String(): &legacyv2.AddVerifiedClientParams{},

		tools.V10.String(): &legacyv3.AddVerifiedClientParams{},
		tools.V11.String(): &legacyv3.AddVerifiedClientParams{},

		tools.V12.String(): &legacyv4.AddVerifiedClientParams{},
		tools.V13.String(): &legacyv5.AddVerifiedClientParams{},
		tools.V14.String(): &legacyv6.AddVerifiedClientParams{},
		tools.V15.String(): &legacyv7.AddVerifiedClientParams{},
		tools.V16.String(): &verifregv8.AddVerifiedClientParams{},
		tools.V17.String(): &verifregv9.AddVerifiedClientParams{},
		tools.V18.String(): &verifregv10.AddVerifiedClientParams{},

		tools.V19.String(): &verifregv11.AddVerifiedClientParams{},
		tools.V20.String(): &verifregv11.AddVerifiedClientParams{},

		tools.V21.String(): &verifregv12.AddVerifiedClientParams{},
		tools.V22.String(): &verifregv13.AddVerifiedClientParams{},
		tools.V23.String(): &verifregv14.AddVerifiedClientParams{},
		tools.V24.String(): &verifregv15.AddVerifiedClientParams{},
		tools.V25.String(): &verifregv16.AddVerifiedClientParams{},
	}
}

func useBytesParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.UseBytesParams{},

		tools.V8.String(): &legacyv2.UseBytesParams{},
		tools.V9.String(): &legacyv2.UseBytesParams{},

		tools.V10.String(): &legacyv3.UseBytesParams{},
		tools.V11.String(): &legacyv3.UseBytesParams{},

		tools.V12.String(): &legacyv4.UseBytesParams{},
		tools.V13.String(): &legacyv5.UseBytesParams{},
		tools.V14.String(): &legacyv6.UseBytesParams{},
		tools.V15.String(): &legacyv7.UseBytesParams{},
		tools.V16.String(): &verifregv8.UseBytesParams{},
		tools.V17.String(): &verifregv9.UseBytesParams{},
		tools.V18.String(): &verifregv10.UseBytesParams{},

		tools.V19.String(): &verifregv11.UseBytesParams{},
		tools.V20.String(): &verifregv11.UseBytesParams{},

		tools.V21.String(): &verifregv12.UseBytesParams{},
		tools.V22.String(): &verifregv13.UseBytesParams{},
		tools.V23.String(): &verifregv14.UseBytesParams{},
		tools.V24.String(): &verifregv15.UseBytesParams{},
		tools.V25.String(): &verifregv16.UseBytesParams{},
	}
}

func restoreBytesParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.RestoreBytesParams{},

		tools.V8.String(): &legacyv2.RestoreBytesParams{},
		tools.V9.String(): &legacyv2.RestoreBytesParams{},

		tools.V10.String(): &legacyv3.RestoreBytesParams{},
		tools.V11.String(): &legacyv3.RestoreBytesParams{},

		tools.V12.String(): &legacyv4.RestoreBytesParams{},
		tools.V13.String(): &legacyv5.RestoreBytesParams{},
		tools.V14.String(): &legacyv6.RestoreBytesParams{},
		tools.V15.String(): &legacyv7.RestoreBytesParams{},
		tools.V16.String(): &verifregv8.RestoreBytesParams{},
		tools.V17.String(): &verifregv9.RestoreBytesParams{},
		tools.V18.String(): &verifregv10.RestoreBytesParams{},

		tools.V19.String(): &verifregv11.RestoreBytesParams{},
		tools.V20.String(): &verifregv11.RestoreBytesParams{},

		tools.V21.String(): &verifregv12.RestoreBytesParams{},
		tools.V22.String(): &verifregv13.RestoreBytesParams{},
		tools.V23.String(): &verifregv14.RestoreBytesParams{},
		tools.V24.String(): &verifregv15.RestoreBytesParams{},
		tools.V25.String(): &verifregv16.RestoreBytesParams{},
	}
}

func dataCap() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.DataCap{},

		tools.V8.String(): &legacyv2.DataCap{},
		tools.V9.String(): &legacyv2.DataCap{},

		tools.V10.String(): &legacyv3.DataCap{},
		tools.V11.String(): &legacyv3.DataCap{},

		tools.V12.String(): &legacyv4.DataCap{},
		tools.V13.String(): &legacyv5.DataCap{},
		tools.V14.String(): &legacyv6.DataCap{},
		tools.V15.String(): &legacyv7.DataCap{},
		tools.V16.String(): &verifregv8.DataCap{},
		tools.V17.String(): &verifregv9.DataCap{},
		tools.V18.String(): &verifregv10.DataCap{},

		tools.V19.String(): &verifregv11.DataCap{},
		tools.V20.String(): &verifregv11.DataCap{},

		tools.V21.String(): &verifregv12.DataCap{},
		tools.V22.String(): &verifregv13.DataCap{},
		tools.V23.String(): &verifregv14.DataCap{},
		tools.V24.String(): &verifregv15.DataCap{},
		tools.V25.String(): &verifregv16.DataCap{},
	}
}

func removeExpiredAllocationsParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &verifregv9.RemoveExpiredAllocationsParams{},
		tools.V18.String(): &verifregv10.RemoveExpiredAllocationsParams{},

		tools.V19.String(): &verifregv11.RemoveExpiredAllocationsParams{},
		tools.V20.String(): &verifregv11.RemoveExpiredAllocationsParams{},

		tools.V21.String(): &verifregv12.RemoveExpiredAllocationsParams{},
		tools.V22.String(): &verifregv13.RemoveExpiredAllocationsParams{},
		tools.V23.String(): &verifregv14.RemoveExpiredAllocationsParams{},
		tools.V24.String(): &verifregv15.RemoveExpiredAllocationsParams{},
		tools.V25.String(): &verifregv16.RemoveExpiredAllocationsParams{},
	}
}

func removeExpiredAllocationsReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &verifregv9.RemoveExpiredAllocationsReturn{},
		tools.V18.String(): &verifregv10.RemoveExpiredAllocationsReturn{},

		tools.V19.String(): &verifregv11.RemoveExpiredAllocationsReturn{},
		tools.V20.String(): &verifregv11.RemoveExpiredAllocationsReturn{},

		tools.V21.String(): &verifregv12.RemoveExpiredAllocationsReturn{},
		tools.V22.String(): &verifregv13.RemoveExpiredAllocationsReturn{},
		tools.V23.String(): &verifregv14.RemoveExpiredAllocationsReturn{},
		tools.V24.String(): &verifregv15.RemoveExpiredAllocationsReturn{},
		tools.V25.String(): &verifregv16.RemoveExpiredAllocationsReturn{},
	}
}

func claimAllocationsParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &verifregv9.ClaimAllocationsParams{},
		tools.V18.String(): &verifregv10.ClaimAllocationsParams{},

		tools.V19.String(): &verifregv11.ClaimAllocationsParams{},
		tools.V20.String(): &verifregv11.ClaimAllocationsParams{},

		tools.V21.String(): &verifregv12.ClaimAllocationsParams{},
		tools.V22.String(): &verifregv13.ClaimAllocationsParams{},
		tools.V23.String(): &verifregv14.ClaimAllocationsParams{},
		tools.V24.String(): &verifregv15.ClaimAllocationsParams{},
		tools.V25.String(): &verifregv16.ClaimAllocationsParams{},
	}
}

func claimAllocationsReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &verifregv9.ClaimAllocationsReturn{},
		tools.V18.String(): &verifregv10.ClaimAllocationsReturn{},

		tools.V19.String(): &verifregv11.ClaimAllocationsReturn{},
		tools.V20.String(): &verifregv11.ClaimAllocationsReturn{},

		tools.V21.String(): &verifregv12.ClaimAllocationsReturn{},
		tools.V22.String(): &verifregv13.ClaimAllocationsReturn{},
		tools.V23.String(): &verifregv14.ClaimAllocationsReturn{},
		tools.V24.String(): &verifregv15.ClaimAllocationsReturn{},
		tools.V25.String(): &verifregv16.ClaimAllocationsReturn{},
	}
}

func getClaimsParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &verifregv9.GetClaimsParams{},
		tools.V18.String(): &verifregv10.GetClaimsParams{},

		tools.V19.String(): &verifregv11.GetClaimsParams{},
		tools.V20.String(): &verifregv11.GetClaimsParams{},

		tools.V21.String(): &verifregv12.GetClaimsParams{},
		tools.V22.String(): &verifregv13.GetClaimsParams{},
		tools.V23.String(): &verifregv14.GetClaimsParams{},
		tools.V24.String(): &verifregv15.GetClaimsParams{},
		tools.V25.String(): &verifregv16.GetClaimsParams{},
	}
}

func getClaimsReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &verifregv9.GetClaimsReturn{},
		tools.V18.String(): &verifregv10.GetClaimsReturn{},

		tools.V19.String(): &verifregv11.GetClaimsReturn{},
		tools.V20.String(): &verifregv11.GetClaimsReturn{},

		tools.V21.String(): &verifregv12.GetClaimsReturn{},
		tools.V22.String(): &verifregv13.GetClaimsReturn{},
		tools.V23.String(): &verifregv14.GetClaimsReturn{},
		tools.V24.String(): &verifregv15.GetClaimsReturn{},
		tools.V25.String(): &verifregv16.GetClaimsReturn{},
	}
}

func extendClaimTermsParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &verifregv9.ExtendClaimTermsParams{},
		tools.V18.String(): &verifregv10.ExtendClaimTermsParams{},

		tools.V19.String(): &verifregv11.ExtendClaimTermsParams{},
		tools.V20.String(): &verifregv11.ExtendClaimTermsParams{},

		tools.V21.String(): &verifregv12.ExtendClaimTermsParams{},
		tools.V22.String(): &verifregv13.ExtendClaimTermsParams{},
		tools.V23.String(): &verifregv14.ExtendClaimTermsParams{},
		tools.V24.String(): &verifregv15.ExtendClaimTermsParams{},
		tools.V25.String(): &verifregv16.ExtendClaimTermsParams{},
	}
}

func extendClaimTermsReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &verifregv9.ExtendClaimTermsReturn{},
		tools.V18.String(): &verifregv10.ExtendClaimTermsReturn{},

		tools.V19.String(): &verifregv11.ExtendClaimTermsReturn{},
		tools.V20.String(): &verifregv11.ExtendClaimTermsReturn{},

		tools.V21.String(): &verifregv12.ExtendClaimTermsReturn{},
		tools.V22.String(): &verifregv13.ExtendClaimTermsReturn{},
		tools.V23.String(): &verifregv14.ExtendClaimTermsReturn{},
		tools.V24.String(): &verifregv15.ExtendClaimTermsReturn{},
		tools.V25.String(): &verifregv16.ExtendClaimTermsReturn{},
	}
}

func removeExpiredClaimsParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &verifregv9.RemoveExpiredClaimsParams{},
		tools.V18.String(): &verifregv10.RemoveExpiredClaimsParams{},

		tools.V19.String(): &verifregv11.RemoveExpiredClaimsParams{},
		tools.V20.String(): &verifregv11.RemoveExpiredClaimsParams{},

		tools.V21.String(): &verifregv12.RemoveExpiredClaimsParams{},
		tools.V22.String(): &verifregv13.RemoveExpiredClaimsParams{},
		tools.V23.String(): &verifregv14.RemoveExpiredClaimsParams{},
		tools.V24.String(): &verifregv15.RemoveExpiredClaimsParams{},
		tools.V25.String(): &verifregv16.RemoveExpiredClaimsParams{},
	}
}

func removeExpiredClaimsReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &verifregv9.RemoveExpiredClaimsReturn{},
		tools.V18.String(): &verifregv10.RemoveExpiredClaimsReturn{},

		tools.V19.String(): &verifregv11.RemoveExpiredClaimsReturn{},
		tools.V20.String(): &verifregv11.RemoveExpiredClaimsReturn{},

		tools.V21.String(): &verifregv12.RemoveExpiredClaimsReturn{},
		tools.V22.String(): &verifregv13.RemoveExpiredClaimsReturn{},
		tools.V23.String(): &verifregv14.RemoveExpiredClaimsReturn{},
		tools.V24.String(): &verifregv15.RemoveExpiredClaimsReturn{},
		tools.V25.String(): &verifregv16.RemoveExpiredClaimsReturn{},
	}
}

func universalReceiverParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &verifregv9.UniversalReceiverParams{},
		tools.V18.String(): &verifregv10.UniversalReceiverParams{},

		tools.V19.String(): &verifregv11.UniversalReceiverParams{},
		tools.V20.String(): &verifregv11.UniversalReceiverParams{},

		tools.V21.String(): &verifregv12.UniversalReceiverParams{},
		tools.V22.String(): &verifregv13.UniversalReceiverParams{},
		tools.V23.String(): &verifregv14.UniversalReceiverParams{},
		tools.V24.String(): &verifregv15.UniversalReceiverParams{},
		tools.V25.String(): &verifregv16.UniversalReceiverParams{},
	}
}

func allocationsResponse() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &verifregv9.AllocationsResponse{},
		tools.V18.String(): &verifregv10.AllocationsResponse{},

		tools.V19.String(): &verifregv11.AllocationsResponse{},
		tools.V20.String(): &verifregv11.AllocationsResponse{},

		tools.V21.String(): &verifregv12.AllocationsResponse{},
		tools.V22.String(): &verifregv13.AllocationsResponse{},
		tools.V23.String(): &verifregv14.AllocationsResponse{},
		tools.V24.String(): &verifregv15.AllocationsResponse{},
		tools.V25.String(): &verifregv16.AllocationsResponse{},
	}
}
