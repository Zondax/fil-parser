package types

import (
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
)

type AddressInfoMapSuite struct {
	suite.Suite
	aim *AddressInfoMap
}

func TestAddressInfoMapSuite(t *testing.T) {
	suite.Run(t, new(AddressInfoMapSuite))
}

func (suite *AddressInfoMapSuite) SetupTest() {
	suite.aim = NewAddressInfoMap()
}

func (suite *AddressInfoMapSuite) TestNewAddressInfoMap() {
	suite.NotNil(suite.aim)
}

func (suite *AddressInfoMapSuite) TestSetAndGet() {
	address := &AddressInfo{
		Short:         "short",
		Robust:        "robust",
		EthAddress:    "ethAddress",
		ActorCid:      "actorCid",
		ActorType:     "actorType",
		CreationTxCid: "creationTxCid",
	}
	suite.aim.Set("key", address)

	retrieved, ok := suite.aim.Get("key")

	suite.True(ok)
	suite.Equal(address, retrieved)
}

func (suite *AddressInfoMapSuite) TestLen() {
	suite.Equal(0, suite.aim.Len())
	suite.aim.Set("key", &AddressInfo{})
	suite.Equal(1, suite.aim.Len())
}

func (suite *AddressInfoMapSuite) TestRange() {
	address := &AddressInfo{
		Short:  "short",
		Robust: "robust",
	}
	suite.aim.Set("key", address)

	var once sync.Once
	suite.aim.Range(func(key string, value *AddressInfo) bool {
		once.Do(func() {
			suite.Equal("key", key)
			suite.Equal(address, value)
		})
		return true
	})
}
