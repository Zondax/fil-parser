package types_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zondax/fil-parser/types"
)

type AddressInfoMapSuite struct {
	suite.Suite
	aim *types.AddressInfoMap
}

func TestAddressInfoMapSuite(t *testing.T) {
	suite.Run(t, new(AddressInfoMapSuite))
}

func (suite *AddressInfoMapSuite) SetupTest() {
	suite.aim = types.NewAddressInfoMap()
}

func (suite *AddressInfoMapSuite) TestNewAddressInfoMap() {
	suite.NotNil(suite.aim)
}

func (suite *AddressInfoMapSuite) TestSetAndGet() {
	address := &types.AddressInfo{
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
	suite.aim.Set("key", &types.AddressInfo{})
	suite.Equal(1, suite.aim.Len())
}

func (suite *AddressInfoMapSuite) TestRange() {
	address := &types.AddressInfo{
		Short:  "short",
		Robust: "robust",
	}
	suite.aim.Set("key", address)

	var once sync.Once
	suite.aim.Range(func(key string, value *types.AddressInfo) bool {
		once.Do(func() {
			suite.Equal("key", key)
			suite.Equal(address, value)
		})
		return true
	})
}

func (suite *AddressInfoMapSuite) TestCopy() {
	address1 := &types.AddressInfo{
		Short:  "short1",
		Robust: "robust1",
	}
	address2 := &types.AddressInfo{
		Short:  "short2",
		Robust: "robust2",
	}

	suite.aim.Set("key1", address1)
	suite.aim.Set("key2", address2)

	copiedMap := suite.aim.Copy()

	suite.Equal(len(copiedMap), suite.aim.Len())

	suite.Equal(address1, copiedMap["key1"])
	suite.Equal(address2, copiedMap["key2"])
}
