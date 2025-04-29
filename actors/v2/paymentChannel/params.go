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

	cbg "github.com/whyrusleeping/cbor-gen"
)

var constructorParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ConstructorParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ConstructorParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ConstructorParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ConstructorParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ConstructorParams) },

	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ConstructorParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(paychv8.ConstructorParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(paychv9.ConstructorParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(paychv10.ConstructorParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(paychv11.ConstructorParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(paychv11.ConstructorParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(paychv12.ConstructorParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(paychv13.ConstructorParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(paychv14.ConstructorParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(paychv15.ConstructorParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(paychv16.ConstructorParams) },
}

var updateChannelStateParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.UpdateChannelStateParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.UpdateChannelStateParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.UpdateChannelStateParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.UpdateChannelStateParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.UpdateChannelStateParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.UpdateChannelStateParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.UpdateChannelStateParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.UpdateChannelStateParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.UpdateChannelStateParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.UpdateChannelStateParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.UpdateChannelStateParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.UpdateChannelStateParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.UpdateChannelStateParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.UpdateChannelStateParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.UpdateChannelStateParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(paychv8.UpdateChannelStateParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(paychv9.UpdateChannelStateParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(paychv10.UpdateChannelStateParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(paychv11.UpdateChannelStateParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(paychv11.UpdateChannelStateParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(paychv12.UpdateChannelStateParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(paychv13.UpdateChannelStateParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(paychv14.UpdateChannelStateParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(paychv15.UpdateChannelStateParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(paychv16.UpdateChannelStateParams) },
}
