package market

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	v10Market "github.com/filecoin-project/go-state-types/builtin/v10/market"
	v11Market "github.com/filecoin-project/go-state-types/builtin/v11/market"
	v12Market "github.com/filecoin-project/go-state-types/builtin/v12/market"
	v13Market "github.com/filecoin-project/go-state-types/builtin/v13/market"
	v14Market "github.com/filecoin-project/go-state-types/builtin/v14/market"
	v15Market "github.com/filecoin-project/go-state-types/builtin/v15/market"
	v8Market "github.com/filecoin-project/go-state-types/builtin/v8/market"
	v9Market "github.com/filecoin-project/go-state-types/builtin/v9/market"
)

func ParseAddBalance(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*address.Address, *address.Address](rawParams, nil, false)
	case 14:
		return parseGeneric[*address.Address, *address.Address](rawParams, nil, false)
	case 13:
		return parseGeneric[*address.Address, *address.Address](rawParams, nil, false)
	case 12:
		return parseGeneric[*address.Address, *address.Address](rawParams, nil, false)
	case 11:
		return parseGeneric[*address.Address, *address.Address](rawParams, nil, false)
	case 10:
		return parseGeneric[*address.Address, *address.Address](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}
func ParseWithdrawBalance(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*v15Market.WithdrawBalanceParams, *v15Market.WithdrawBalanceParams](rawParams, nil, false)
	case 14:
		return parseGeneric[*v14Market.WithdrawBalanceParams, *v14Market.WithdrawBalanceParams](rawParams, nil, false)
	case 13:
		return parseGeneric[*v13Market.WithdrawBalanceParams, *v13Market.WithdrawBalanceParams](rawParams, nil, false)
	case 12:
		return parseGeneric[*v12Market.WithdrawBalanceParams, *v12Market.WithdrawBalanceParams](rawParams, nil, false)
	case 11:
		return parseGeneric[*v11Market.WithdrawBalanceParams, *v11Market.WithdrawBalanceParams](rawParams, nil, false)
	case 10:
		return parseGeneric[*v10Market.WithdrawBalanceParams, *v10Market.WithdrawBalanceParams](rawParams, nil, false)
	case 9:
		return parseGeneric[*v9Market.WithdrawBalanceParams, *v9Market.WithdrawBalanceParams](rawParams, nil, false)
	case 8:
		return parseGeneric[*v8Market.WithdrawBalanceParams, *v8Market.WithdrawBalanceParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func PublishStorageDealsParams(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*v15Market.PublishStorageDealsParams, *v15Market.PublishStorageDealsReturn](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*v14Market.PublishStorageDealsParams, *v14Market.PublishStorageDealsReturn](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*v13Market.PublishStorageDealsParams, *v13Market.PublishStorageDealsReturn](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*v12Market.PublishStorageDealsParams, *v12Market.PublishStorageDealsReturn](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*v11Market.PublishStorageDealsParams, *v11Market.PublishStorageDealsReturn](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*v10Market.PublishStorageDealsParams, *v10Market.PublishStorageDealsReturn](rawParams, rawReturn, true)
	case 9:
		return parseGeneric[*v9Market.PublishStorageDealsParams, *v9Market.PublishStorageDealsReturn](rawParams, rawReturn, true)
	case 8:
		return parseGeneric[*v8Market.PublishStorageDealsParams, *v8Market.PublishStorageDealsReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}
func VerifyDealsForActivationParams(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*v15Market.VerifyDealsForActivationParams, *v15Market.VerifyDealsForActivationReturn](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*v14Market.VerifyDealsForActivationParams, *v14Market.VerifyDealsForActivationReturn](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*v13Market.VerifyDealsForActivationParams, *v13Market.VerifyDealsForActivationReturn](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*v12Market.VerifyDealsForActivationParams, *v12Market.VerifyDealsForActivationReturn](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*v11Market.VerifyDealsForActivationParams, *v11Market.VerifyDealsForActivationReturn](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*v10Market.VerifyDealsForActivationParams, *v10Market.VerifyDealsForActivationReturn](rawParams, rawReturn, true)
	case 9:
		return parseGeneric[*v9Market.VerifyDealsForActivationParams, *v9Market.VerifyDealsForActivationReturn](rawParams, rawReturn, true)
	case 8:
		return parseGeneric[*v8Market.VerifyDealsForActivationParams, *v8Market.VerifyDealsForActivationReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ActivateDealsParams(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*v15Market.ActivateDealsParams, *v15Market.ActivateDealsResult](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*v14Market.ActivateDealsParams, *v14Market.ActivateDealsResult](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*v13Market.ActivateDealsParams, *v13Market.ActivateDealsResult](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*v12Market.ActivateDealsParams, *v12Market.ActivateDealsResult](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*v11Market.ActivateDealsParams, *v11Market.ActivateDealsResult](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*v10Market.ActivateDealsParams, *v10Market.ActivateDealsResult](rawParams, rawReturn, true)
	case 9:
		return parseGeneric[*v9Market.ActivateDealsParams, *v9Market.ActivateDealsResult](rawParams, rawReturn, true)
	case 8:
		// return parseGeneric[*v8Market.ActivateDealsParams, *v8Market.ActivateDealsResult](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func OnMinerSectorsTerminateParams(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*v15Market.OnMinerSectorsTerminateParams, *v15Market.OnMinerSectorsTerminateParams](rawParams, nil, false)
	case 14:
		return parseGeneric[*v14Market.OnMinerSectorsTerminateParams, *v14Market.OnMinerSectorsTerminateParams](rawParams, nil, false)
	case 13:
		return parseGeneric[*v13Market.OnMinerSectorsTerminateParams, *v13Market.OnMinerSectorsTerminateParams](rawParams, nil, false)
	case 12:
		return parseGeneric[*v12Market.OnMinerSectorsTerminateParams, *v12Market.OnMinerSectorsTerminateParams](rawParams, nil, false)
	case 11:
		return parseGeneric[*v11Market.OnMinerSectorsTerminateParams, *v11Market.OnMinerSectorsTerminateParams](rawParams, nil, false)
	case 10:
		return parseGeneric[*v10Market.OnMinerSectorsTerminateParams, *v10Market.OnMinerSectorsTerminateParams](rawParams, nil, false)
	case 9:
		return parseGeneric[*v9Market.OnMinerSectorsTerminateParams, *v9Market.OnMinerSectorsTerminateParams](rawParams, nil, false)
	case 8:
		return parseGeneric[*v8Market.OnMinerSectorsTerminateParams, *v8Market.OnMinerSectorsTerminateParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ComputeDataCommitmentParams(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		// return parseGeneric[*v15Market.ComputeDataCommitmentParams, *v15Market.ComputeDataCommitmentReturn](rawParams, rawReturn, true)
	case 14:
		// return parseGeneric[*v14Market.ComputeDataCommitmentParams, *v14Market.ComputeDataCommitmentReturn](rawParams, rawReturn, true)
	case 13:
		// return parseGeneric[*v13Market.ComputeDataCommitmentParams, *v13Market.ComputeDataCommitmentReturn](rawParams, rawReturn, true)
	case 12:
		// return parseGeneric[*v12Market.ComputeDataCommitmentParams, *v12Market.ComputeDataCommitmentReturn](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*v11Market.ComputeDataCommitmentParams, *v11Market.ComputeDataCommitmentReturn](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*v10Market.ComputeDataCommitmentParams, *v10Market.ComputeDataCommitmentReturn](rawParams, rawReturn, true)
	case 9:
		return parseGeneric[*v9Market.ComputeDataCommitmentParams, *v9Market.ComputeDataCommitmentReturn](rawParams, rawReturn, true)
	case 8:
		return parseGeneric[*v8Market.ComputeDataCommitmentParams, *v8Market.ComputeDataCommitmentReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetBalanceParams(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*address.Address, *v15Market.GetBalanceReturn](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*address.Address, *v14Market.GetBalanceReturn](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*address.Address, *v13Market.GetBalanceReturn](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*address.Address, *v12Market.GetBalanceReturn](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*address.Address, *v11Market.GetBalanceReturn](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*address.Address, *v10Market.GetBalanceReturn](rawParams, rawReturn, true)
	case 9:
		// return parseGeneric[*address.Address, *v9Market.GetBalanceReturn](rawParams, rawReturn, true)
	case 8:
		// return parseGeneric[*address.Address, *v8Market.GetBalanceReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetDealDataCommitmentParams(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*v15Market.GetDealDataCommitmentParams, *v15Market.GetDealDataCommitmentReturn](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*v14Market.GetDealDataCommitmentParams, *v14Market.GetDealDataCommitmentReturn](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*v13Market.GetDealDataCommitmentParams, *v13Market.GetDealDataCommitmentReturn](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*v12Market.GetDealDataCommitmentParams, *v12Market.GetDealDataCommitmentReturn](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*v11Market.GetDealDataCommitmentParams, *v11Market.GetDealDataCommitmentReturn](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*v10Market.GetDealDataCommitmentParams, *v10Market.GetDealDataCommitmentReturn](rawParams, rawReturn, true)
	case 9:
		// return parseGeneric[*v9Market.GetDealDataCommitmentParams, *v9Market.GetDealDataCommitmentReturn](rawParams, rawReturn, true)
	case 8:
		// return parseGeneric[*v8Market.GetDealDataCommitmentParams, *v8Market.GetDealDataCommitmentReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetDealClientParams(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*v15Market.GetDealClientParams, *v15Market.GetDealClientReturn](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*v14Market.GetDealClientParams, *v14Market.GetDealClientReturn](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*v13Market.GetDealClientParams, *v13Market.GetDealClientReturn](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*v12Market.GetDealClientParams, *v12Market.GetDealClientReturn](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*v11Market.GetDealClientParams, *v11Market.GetDealClientReturn](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*v10Market.GetDealClientParams, *v10Market.GetDealClientReturn](rawParams, rawReturn, true)
	case 9:
		// return parseGeneric[*v9Market.GetDealClientParams, *v9Market.GetDealClientReturn](rawParams, rawReturn, true)
	case 8:
		// return parseGeneric[*v8Market.GetDealClientParams, *v8Market.GetDealClientReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetDealProviderParams(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*v15Market.GetDealProviderParams, *v15Market.GetDealProviderReturn](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*v14Market.GetDealProviderParams, *v14Market.GetDealProviderReturn](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*v13Market.GetDealProviderParams, *v13Market.GetDealProviderReturn](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*v12Market.GetDealProviderParams, *v12Market.GetDealProviderReturn](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*v11Market.GetDealProviderParams, *v11Market.GetDealProviderReturn](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*v10Market.GetDealProviderParams, *v10Market.GetDealProviderReturn](rawParams, rawReturn, true)
	case 9:
		// return parseGeneric[*v9Market.GetDealProviderParams, *v9Market.GetDealProviderReturn](rawParams, rawReturn, true)
	case 8:
		// return parseGeneric[*v8Market.GetDealProviderParams, *v8Market.GetDealProviderReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetDealLabelParams(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*v15Market.GetDealLabelParams, *v15Market.GetDealLabelReturn](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*v14Market.GetDealLabelParams, *v14Market.GetDealLabelReturn](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*v13Market.GetDealLabelParams, *v13Market.GetDealLabelReturn](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*v12Market.GetDealLabelParams, *v12Market.GetDealLabelReturn](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*v11Market.GetDealLabelParams, *v11Market.GetDealLabelReturn](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*v10Market.GetDealLabelParams, *v10Market.GetDealLabelReturn](rawParams, rawReturn, true)
	case 9:
		// return parseGeneric[*v9Market.GetDealLabelParams, *v9Market.GetDealLabelReturn](rawParams, rawReturn, true)
	case 8:
		// return parseGeneric[*v8Market.GetDealLabelParams, *v8Market.GetDealLabelReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetDealTermParams(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*v15Market.GetDealTermParams, *v15Market.GetDealTermReturn](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*v14Market.GetDealTermParams, *v14Market.GetDealTermReturn](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*v13Market.GetDealTermParams, *v13Market.GetDealTermReturn](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*v12Market.GetDealTermParams, *v12Market.GetDealTermReturn](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*v11Market.GetDealTermParams, *v11Market.GetDealTermReturn](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*v10Market.GetDealTermParams, *v10Market.GetDealTermReturn](rawParams, rawReturn, true)
	case 9:
		// return parseGeneric[*v9Market.GetDealTermParams, *v9Market.GetDealTermReturn](rawParams, rawReturn, true)
	case 8:
		// return parseGeneric[*v8Market.GetDealTermParams, *v8Market.GetDealTermReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetDealTotalPriceParams(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*v15Market.GetDealTotalPriceParams, *v15Market.GetDealTotalPriceReturn](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*v14Market.GetDealTotalPriceParams, *v14Market.GetDealTotalPriceReturn](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*v13Market.GetDealTotalPriceParams, *v13Market.GetDealTotalPriceReturn](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*v12Market.GetDealTotalPriceParams, *v12Market.GetDealTotalPriceReturn](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*v11Market.GetDealTotalPriceParams, *v11Market.GetDealTotalPriceReturn](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*v10Market.GetDealTotalPriceParams, *v10Market.GetDealTotalPriceReturn](rawParams, rawReturn, true)
	case 9:
		// return parseGeneric[*v9Market.GetDealTotalPriceParams, *v9Market.GetDealTotalPriceReturn](rawParams, rawReturn, true)
	case 8:
		// return parseGeneric[*v8Market.GetDealTotalPriceParams, *v8Market.GetDealTotalPriceReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetDealClientCollateralParams(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*v15Market.GetDealClientCollateralParams, *v15Market.GetDealClientCollateralReturn](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*v14Market.GetDealClientCollateralParams, *v14Market.GetDealClientCollateralReturn](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*v13Market.GetDealClientCollateralParams, *v13Market.GetDealClientCollateralReturn](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*v12Market.GetDealClientCollateralParams, *v12Market.GetDealClientCollateralReturn](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*v11Market.GetDealClientCollateralParams, *v11Market.GetDealClientCollateralReturn](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*v10Market.GetDealClientCollateralParams, *v10Market.GetDealClientCollateralReturn](rawParams, rawReturn, true)
	case 9:
		// return parseGeneric[*v9Market.GetDealClientCollateralParams, *v9Market.GetDealClientCollateralReturn](rawParams, rawReturn, true)
	case 8:
		// return parseGeneric[*v8Market.GetDealClientCollateralParams, *v8Market.GetDealClientCollateralReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetDealProviderCollateralParams(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*v15Market.GetDealProviderCollateralParams, *v15Market.GetDealProviderCollateralReturn](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*v14Market.GetDealProviderCollateralParams, *v14Market.GetDealProviderCollateralReturn](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*v13Market.GetDealProviderCollateralParams, *v13Market.GetDealProviderCollateralReturn](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*v12Market.GetDealProviderCollateralParams, *v12Market.GetDealProviderCollateralReturn](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*v11Market.GetDealProviderCollateralParams, *v11Market.GetDealProviderCollateralReturn](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*v10Market.GetDealProviderCollateralParams, *v10Market.GetDealProviderCollateralReturn](rawParams, rawReturn, true)
	case 9:
		// return parseGeneric[*v9Market.GetDealProviderCollateralParams, *v9Market.GetDealProviderCollateralReturn](rawParams, rawReturn, true)
	case 8:
		// return parseGeneric[*v8Market.GetDealProviderCollateralParams, *v8Market.GetDealProviderCollateralReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetDealVerifiedParams(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*v15Market.GetDealVerifiedParams, *v15Market.GetDealVerifiedReturn](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*v14Market.GetDealVerifiedParams, *v14Market.GetDealVerifiedReturn](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*v13Market.GetDealVerifiedParams, *v13Market.GetDealVerifiedReturn](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*v12Market.GetDealVerifiedParams, *v12Market.GetDealVerifiedReturn](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*v11Market.GetDealVerifiedParams, *v11Market.GetDealVerifiedReturn](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*v10Market.GetDealVerifiedParams, *v10Market.GetDealVerifiedReturn](rawParams, rawReturn, true)
	case 9:
		// return parseGeneric[*v9Market.GetDealVerifiedParams, *v9Market.GetDealVerifiedReturn](rawParams, rawReturn, true)
	case 8:
		// return parseGeneric[*v8Market.GetDealVerifiedParams, *v8Market.GetDealVerifiedReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetDealActivationParams(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*v15Market.GetDealActivationParams, *v15Market.GetDealActivationReturn](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*v14Market.GetDealActivationParams, *v14Market.GetDealActivationReturn](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*v13Market.GetDealActivationParams, *v13Market.GetDealActivationReturn](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*v12Market.GetDealActivationParams, *v12Market.GetDealActivationReturn](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*v11Market.GetDealActivationParams, *v11Market.GetDealActivationReturn](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*v10Market.GetDealActivationParams, *v10Market.GetDealActivationReturn](rawParams, rawReturn, true)
	case 9:
		// return parseGeneric[*v9Market.GetDealActivationParams, *v9Market.GetDealActivationReturn](rawParams, rawReturn, true)
	case 8:
		// return parseGeneric[*v8Market.GetDealActivationParams, *v8Market.GetDealActivationReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}
