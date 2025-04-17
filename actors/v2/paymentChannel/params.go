package paymentChannel

import (
	paychv10 "github.com/filecoin-project/go-state-types/builtin/v10/paych"
	paychv11 "github.com/filecoin-project/go-state-types/builtin/v11/paych"
	paychv12 "github.com/filecoin-project/go-state-types/builtin/v12/paych"
	paychv13 "github.com/filecoin-project/go-state-types/builtin/v13/paych"
	paychv14 "github.com/filecoin-project/go-state-types/builtin/v14/paych"
	paychv15 "github.com/filecoin-project/go-state-types/builtin/v15/paych"
	paychv16 "github.com/filecoin-project/go-state-types/builtin/v16/paych"
	paychv8 "github.com/filecoin-project/go-state-types/builtin/v8/paych"
	paychv9 "github.com/filecoin-project/go-state-types/builtin/v9/paych"
	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/paych"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/paych"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/paych"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/paych"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/paych"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/paych"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/paych"
	"github.com/zondax/fil-parser/tools"
)

func constructorParams() map[string]paymentChannelParams {
	return map[string]paymentChannelParams{
		tools.V7.String(): &legacyv1.ConstructorParams{},

		tools.V8.String(): &legacyv2.ConstructorParams{},
		tools.V9.String(): &legacyv2.ConstructorParams{},

		tools.V10.String(): &legacyv3.ConstructorParams{},
		tools.V11.String(): &legacyv3.ConstructorParams{},

		tools.V12.String(): &legacyv4.ConstructorParams{},
		tools.V13.String(): &legacyv5.ConstructorParams{},
		tools.V14.String(): &legacyv6.ConstructorParams{},

		tools.V15.String(): &legacyv7.ConstructorParams{},
		tools.V16.String(): &paychv8.ConstructorParams{},
		tools.V17.String(): &paychv9.ConstructorParams{},
		tools.V18.String(): &paychv10.ConstructorParams{},

		tools.V19.String(): &paychv11.ConstructorParams{},
		tools.V20.String(): &paychv11.ConstructorParams{},

		tools.V21.String(): &paychv12.ConstructorParams{},
		tools.V22.String(): &paychv13.ConstructorParams{},
		tools.V23.String(): &paychv14.ConstructorParams{},
		tools.V24.String(): &paychv15.ConstructorParams{},
		tools.V25.String(): &paychv16.ConstructorParams{},
	}
}

func updateChannelStateParams() map[string]paymentChannelParams {
	return map[string]paymentChannelParams{
		tools.V7.String(): &legacyv1.UpdateChannelStateParams{},

		tools.V8.String(): &legacyv2.UpdateChannelStateParams{},
		tools.V9.String(): &legacyv2.UpdateChannelStateParams{},

		tools.V10.String(): &legacyv3.UpdateChannelStateParams{},
		tools.V11.String(): &legacyv3.UpdateChannelStateParams{},

		tools.V12.String(): &legacyv4.UpdateChannelStateParams{},
		tools.V13.String(): &legacyv5.UpdateChannelStateParams{},
		tools.V14.String(): &legacyv6.UpdateChannelStateParams{},
		tools.V15.String(): &legacyv7.UpdateChannelStateParams{},
		tools.V16.String(): &paychv8.UpdateChannelStateParams{},
		tools.V17.String(): &paychv9.UpdateChannelStateParams{},
		tools.V18.String(): &paychv10.UpdateChannelStateParams{},

		tools.V19.String(): &paychv11.UpdateChannelStateParams{},
		tools.V20.String(): &paychv11.UpdateChannelStateParams{},

		tools.V21.String(): &paychv12.UpdateChannelStateParams{},
		tools.V22.String(): &paychv13.UpdateChannelStateParams{},
		tools.V23.String(): &paychv14.UpdateChannelStateParams{},
		tools.V24.String(): &paychv15.UpdateChannelStateParams{},
		tools.V25.String(): &paychv16.UpdateChannelStateParams{},
	}
}
