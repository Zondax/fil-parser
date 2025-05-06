package miner

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/filecoin-project/go-address"
	miner10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	miner16 "github.com/filecoin-project/go-state-types/builtin/v16/miner"
	miner8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"

	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/miner"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func parseGeneric[T minerParam, R minerReturn](rawParams, rawReturn []byte, customReturn bool, params T, r R, key string) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, fmt.Errorf("error unmarshalling params: %w", err)
	}
	metadata[key] = params
	if !customReturn {
		return metadata, nil
	}
	if len(rawReturn) > 0 {
		reader = bytes.NewReader(rawReturn)
		err = r.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, err
		}
		metadata[parser.ReturnKey] = r
	}
	return metadata, nil
}

func parseControlReturn[R minerReturn](rawParams, rawReturn []byte, controlReturn R) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if rawParams != nil {
		metadata[parser.ParamsKey] = base64.StdEncoding.EncodeToString(rawParams)
	}
	reader := bytes.NewReader(rawReturn)
	err := controlReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	controlAddress, err := getControlAddress(controlReturn)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = controlAddress
	return metadata, nil
}

func getControlAddress(controlReturn any) (parser.ControlAddress, error) {
	controlAddress := parser.ControlAddress{}
	setControlReturn := func(owner, worker string, controlAddrs []string) {
		controlAddress.Owner = owner
		controlAddress.Worker = worker
		controlAddress.ControlAddrs = controlAddrs
	}

	switch v := controlReturn.(type) {
	case *legacyv1.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *legacyv7.GetControlAddressesReturn: // all previous legacy versions upto v1 are the same exact type, adding to the switch case will cause a compile time error
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner8.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner9.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner10.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner11.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner12.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner13.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner14.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner15.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner16.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	default:
		return controlAddress, fmt.Errorf("unsupported control return type: %T", v)
	}
	return controlAddress, nil

}

func getBeneficiaryReturn(network string, height int64, rawReturn []byte) (parser.GetBeneficiaryReturn, error) {
	reader := bytes.NewReader(rawReturn)

	var (
		beneficiary string
		quota       string
		usedQuota   string
		expiration  int64

		newBeneficiary        string
		newQuota              string
		newExpiration         int64
		approvedByBeneficiary bool
		approvedByNominee     bool
	)
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return parser.GetBeneficiaryReturn{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V17.IsSupported(network, height):
		tmp := &miner9.GetBeneficiaryReturn{}
		err := tmp.UnmarshalCBOR(reader)
		if err != nil {
			return parser.GetBeneficiaryReturn{}, err
		}
		beneficiary = tmp.Active.Beneficiary.String()
		quota = tmp.Active.Term.Quota.String()
		usedQuota = tmp.Active.Term.UsedQuota.String()
		expiration = int64(tmp.Active.Term.Expiration)

		newBeneficiary = tmp.Proposed.NewBeneficiary.String()
		newQuota = tmp.Proposed.NewQuota.String()
		newExpiration = int64(tmp.Proposed.NewExpiration)
		approvedByBeneficiary = tmp.Proposed.ApprovedByBeneficiary
		approvedByNominee = tmp.Proposed.ApprovedByNominee

	case tools.V18.IsSupported(network, height):
		tmp := &miner10.GetBeneficiaryReturn{}
		err := tmp.UnmarshalCBOR(reader)
		if err != nil {
			return parser.GetBeneficiaryReturn{}, err
		}
		beneficiary = tmp.Active.Beneficiary.String()
		quota = tmp.Active.Term.Quota.String()
		usedQuota = tmp.Active.Term.UsedQuota.String()
		expiration = int64(tmp.Active.Term.Expiration)

		newBeneficiary = tmp.Proposed.NewBeneficiary.String()
		newQuota = tmp.Proposed.NewQuota.String()
		newExpiration = int64(tmp.Proposed.NewExpiration)
		approvedByBeneficiary = tmp.Proposed.ApprovedByBeneficiary
		approvedByNominee = tmp.Proposed.ApprovedByNominee
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		tmp := &miner11.GetBeneficiaryReturn{}
		err := tmp.UnmarshalCBOR(reader)
		if err != nil {
			return parser.GetBeneficiaryReturn{}, err
		}
		beneficiary = tmp.Active.Beneficiary.String()
		quota = tmp.Active.Term.Quota.String()
		usedQuota = tmp.Active.Term.UsedQuota.String()
		expiration = int64(tmp.Active.Term.Expiration)

		newBeneficiary = tmp.Proposed.NewBeneficiary.String()
		newQuota = tmp.Proposed.NewQuota.String()
		newExpiration = int64(tmp.Proposed.NewExpiration)
		approvedByBeneficiary = tmp.Proposed.ApprovedByBeneficiary
		approvedByNominee = tmp.Proposed.ApprovedByNominee
	case tools.V21.IsSupported(network, height):
		tmp := &miner12.GetBeneficiaryReturn{}
		err := tmp.UnmarshalCBOR(reader)
		if err != nil {
			return parser.GetBeneficiaryReturn{}, err
		}
		beneficiary = tmp.Active.Beneficiary.String()
		quota = tmp.Active.Term.Quota.String()
		usedQuota = tmp.Active.Term.UsedQuota.String()
		expiration = int64(tmp.Active.Term.Expiration)

		newBeneficiary = tmp.Proposed.NewBeneficiary.String()
		newQuota = tmp.Proposed.NewQuota.String()
		newExpiration = int64(tmp.Proposed.NewExpiration)
		approvedByBeneficiary = tmp.Proposed.ApprovedByBeneficiary
		approvedByNominee = tmp.Proposed.ApprovedByNominee
	case tools.V22.IsSupported(network, height):
		tmp := &miner13.GetBeneficiaryReturn{}
		err := tmp.UnmarshalCBOR(reader)
		if err != nil {
			return parser.GetBeneficiaryReturn{}, err
		}
		beneficiary = tmp.Active.Beneficiary.String()
		quota = tmp.Active.Term.Quota.String()
		usedQuota = tmp.Active.Term.UsedQuota.String()
		expiration = int64(tmp.Active.Term.Expiration)

		newBeneficiary = tmp.Proposed.NewBeneficiary.String()
		newQuota = tmp.Proposed.NewQuota.String()
		newExpiration = int64(tmp.Proposed.NewExpiration)
		approvedByBeneficiary = tmp.Proposed.ApprovedByBeneficiary
		approvedByNominee = tmp.Proposed.ApprovedByNominee
	case tools.V23.IsSupported(network, height):
		tmp := &miner14.GetBeneficiaryReturn{}
		err := tmp.UnmarshalCBOR(reader)
		if err != nil {
			return parser.GetBeneficiaryReturn{}, err
		}
		beneficiary = tmp.Active.Beneficiary.String()
		quota = tmp.Active.Term.Quota.String()
		usedQuota = tmp.Active.Term.UsedQuota.String()
		expiration = int64(tmp.Active.Term.Expiration)

		newBeneficiary = tmp.Proposed.NewBeneficiary.String()
		newQuota = tmp.Proposed.NewQuota.String()
		newExpiration = int64(tmp.Proposed.NewExpiration)
		approvedByBeneficiary = tmp.Proposed.ApprovedByBeneficiary
		approvedByNominee = tmp.Proposed.ApprovedByNominee
	case tools.V24.IsSupported(network, height):
		tmp := &miner15.GetBeneficiaryReturn{}
		err := tmp.UnmarshalCBOR(reader)
		if err != nil {
			return parser.GetBeneficiaryReturn{}, err
		}
		beneficiary = tmp.Active.Beneficiary.String()
		quota = tmp.Active.Term.Quota.String()
		usedQuota = tmp.Active.Term.UsedQuota.String()
		expiration = int64(tmp.Active.Term.Expiration)

		newBeneficiary = tmp.Proposed.NewBeneficiary.String()
		newQuota = tmp.Proposed.NewQuota.String()
		newExpiration = int64(tmp.Proposed.NewExpiration)
		approvedByBeneficiary = tmp.Proposed.ApprovedByBeneficiary
		approvedByNominee = tmp.Proposed.ApprovedByNominee
	case tools.V25.IsSupported(network, height):
		tmp := &miner16.GetBeneficiaryReturn{}
		err := tmp.UnmarshalCBOR(reader)
		if err != nil {
			return parser.GetBeneficiaryReturn{}, err
		}
		beneficiary = tmp.Active.Beneficiary.String()
		quota = tmp.Active.Term.Quota.String()
		usedQuota = tmp.Active.Term.UsedQuota.String()
		expiration = int64(tmp.Active.Term.Expiration)

		newBeneficiary = tmp.Proposed.NewBeneficiary.String()
		newQuota = tmp.Proposed.NewQuota.String()
		newExpiration = int64(tmp.Proposed.NewExpiration)
		approvedByBeneficiary = tmp.Proposed.ApprovedByBeneficiary
		approvedByNominee = tmp.Proposed.ApprovedByNominee
	default:
		return parser.GetBeneficiaryReturn{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parser.GetBeneficiaryReturn{
		Active: parser.ActiveBeneficiary{
			Beneficiary: beneficiary,
			Term: parser.BeneficiaryTerm{
				Quota:      quota,
				UsedQuota:  usedQuota,
				Expiration: expiration,
			},
		},
		Proposed: parser.Proposed{
			NewBeneficiary:        newBeneficiary,
			NewQuota:              newQuota,
			NewExpiration:         newExpiration,
			ApprovedByBeneficiary: approvedByBeneficiary,
			ApprovedByNominee:     approvedByNominee,
		},
	}, nil
}

func getControlAddrs(addrs []address.Address) []string {
	r := make([]string, len(addrs))
	for i, addr := range addrs {
		r[i] = addr.String()
	}
	return r
}
