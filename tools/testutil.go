package tools

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/types"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
)

const (
	dataPath          = "../../data/heights"
	snapshotPath      = "../../data/snapshots"
	fileDataExtension = "json.gz"
	tracesPrefix      = "traces"
	tipsetPrefix      = "tipset"
	ethLogPrefix      = "ethlog"
	nativeLogPrefix   = "nativelog"
	NodeUrl           = "https://node-fil-mainnet-next.zondax.ch/rpc/v1"
	calibNextNodeUrl  = "https://hel1-node-fil-calibration-stable.zondax.ch/rpc/v1"
	feeType           = "fee"
	snapshotFileName  = "actor_snapshot.json"
)

func getLotusClient(url string) (api.FullNode, error) {
	lotusClient, _, err := client.NewFullNodeRPCV1(context.Background(), url, http.Header{})
	if err != nil {
		return nil, err
	}
	return lotusClient, nil
}

func GetLib(nodeURL string) (*rosettaFilecoinLib.RosettaConstructionFilecoin, error) {
	lotusClient, err := getLotusClient(nodeURL)
	if err != nil {
		return nil, err
	}

	lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(lotusClient)
	return lib, nil
}

func getFilename(prefix string, height int64) string {
	return fmt.Sprintf(`%s/%s_%d.%s`, dataPath, prefix, height, fileDataExtension)
}

func tracesFilename(height int64) string {
	return getFilename(tracesPrefix, height)
}

func ehtlogFilename(height int64) string {
	return getFilename(ethLogPrefix, height)
}

func nativeLogFilename(height int64) string {
	return getFilename(nativeLogPrefix, height)
}

func tipsetFilename(height int64) string {
	return getFilename(tipsetPrefix, height)
}

func read[T any](fileNameFn func(height int64) string, height int64) (*T, error) {
	raw, err := readGzFile(fileNameFn(height))
	if err != nil {
		return nil, err
	}
	var r T
	err = sonic.Unmarshal(raw, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func readGzFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("error creating gzip reader: %w", err)
	}
	defer gzipReader.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(gzipReader)
	if err != nil {
		return nil, fmt.Errorf("error reading from gzip reader: %w", err)
	}
	return buf.Bytes(), nil
}

func ReadTipset(height int64) (*types.ExtendedTipSet, error) {
	return read[types.ExtendedTipSet](tipsetFilename, height)
}

func ReadEthLogs(height int64) ([]types.EthLog, error) {
	logs, err := read[[]types.EthLog](ehtlogFilename, height)
	if err != nil {
		return nil, err
	}
	return *logs, nil
}

func ReadNativeLogs(height int64) ([]*filTypes.ActorEvent, error) {
	events, err := read[[]*filTypes.ActorEvent](nativeLogFilename, height)
	if err != nil {
		return nil, err
	}
	return *events, nil
}

func ReadTraces(height int64) ([]byte, error) {
	return readGzFile(tracesFilename(height))
}

func ReadActorSnapshot() ([]byte, error) {
	raw, err := os.ReadFile(fmt.Sprintf("%s/%s", snapshotPath, snapshotFileName))
	if err != nil {
		return nil, err
	}
	return raw, nil
}

func ComputeState[T any](height int64, version string) (*T, error) {
	traces, err := ReadTraces(height)
	if err != nil {
		return nil, err
	}

	var computeState T
	err = sonic.UnmarshalString(string(traces), &computeState)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling traces: %w", err)
	}
	return &computeState, nil

}

func CompareResult(result1, result2 map[string]any) bool {
	return reflect.DeepEqual(result1, result2)
}

type TestCase[T any] struct {
	Name      string
	Version   string
	Url       string
	Height    int64
	Network   string
	TipsetKey filTypes.TipSetKey
	Expected  T
	Address   *types.AddressInfo
}

func LoadTestData[T any](network string, fnName string, expected map[string]any) ([]TestCase[T], error) {
	var tests []TestCase[T]
	versions := GetSupportedVersions(network)
	for _, version := range versions {

		versionData := expected[version.String()]
		if versionData == nil {
			return nil, fmt.Errorf("version %s not found in expected data", version.String())
		}
		tmp, ok := versionData.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("version %s not found in expected data", version.String())
		}
		fnData := tmp[fnName]

		if fnData == nil {
			return nil, fmt.Errorf("function %s not found in version %s", fnName, version.String())
		}

		height := DeterministicTestHeight(version)
		if !version.IsSupported(network, height) {
			fmt.Println("skipping", version.String(), height)
			continue
		}

		tests = append(tests, TestCase[T]{
			Name:     fnName + "_" + version.String() + "_" + strconv.FormatInt(height, 10),
			Version:  version.String(),
			Url:      NodeUrl,
			Height:   height,
			Expected: fnData.(T),
		})
	}
	return tests, nil
}

func DeterministicTestHeight(version version) int64 {
	height := version.Height()
	if version != version.next() {
		height = (version.Height() + version.next().Height()) / 2
	}

	return height
}
