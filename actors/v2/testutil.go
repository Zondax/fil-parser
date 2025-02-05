package v2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/rt"
)

var utilityMethods = []string{"Exports", "Code", "IsSingleton", "State"}

func runtimeActorMethods(actor rt.VMActor) []string {
	t := reflect.TypeOf(actor)
	methods := []string{}
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		if slices.Contains(utilityMethods, method.Name) {
			continue
		}
		methods = append(methods, method.Name)
	}
	return methods
}

func builtinActorMethods(methodMap map[abi.MethodNum]builtin.MethodMeta) []string {
	methods := []string{}
	for _, method := range methodMap {
		methods = append(methods, method.Name)
	}
	return methods
}

func MissingMethods(actor Actor, actorVersions []any) []string {
	missingMethods := map[string]bool{}
	txTypes := actor.TransactionTypes()
	for _, actorVersion := range actorVersions {
		switch actorVersion := actorVersion.(type) {
		case rt.VMActor:
			methods := runtimeActorMethods(actorVersion)
			for _, method := range methods {
				_, ok := txTypes[method]
				if !ok {
					missingMethods[method] = true
				}
			}
		case map[abi.MethodNum]builtin.MethodMeta:
			methods := builtinActorMethods(actorVersion)
			for _, method := range methods {
				_, ok := txTypes[method]
				if !ok {
					missingMethods[method] = true
				}
			}
		}
	}
	missingMethodsList := []string{}
	for method := range missingMethods {
		missingMethodsList = append(missingMethodsList, method)
	}

	return missingMethodsList
}

// LatestBuiltinActorVersion retrieves the latest release version from the GitHub releases page here: https://github.com/filecoin-project/builtin-actors/releases
// It calls GitHub's API to get the latest release and then parses the 'tag_name'.
// It expects the tag to begin with a "v" (e.g. "v10.5.0") and returns the major version (in this case, 10).
func LatestBuiltinActorVersion() (uint64, error) {
	var latestVersion uint64
	const url = "https://api.github.com/repos/filecoin-project/builtin-actors/releases/latest"

	resp, err := http.Get(url)
	if err != nil {
		return latestVersion, fmt.Errorf("failed to fetch latest builtin actors release: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return latestVersion, fmt.Errorf("unexpected status code %d when fetching latest release", resp.StatusCode)
	}

	// Decode JSON response.
	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return latestVersion, fmt.Errorf("failed to decode release info: %v", err)
	}

	tag := release.TagName
	if len(tag) == 0 {
		return latestVersion, fmt.Errorf("empty tag name received")
	}

	// Remove potential leading 'v' or 'V'
	if tag[0] == 'v' || tag[0] == 'V' {
		tag = tag[1:]
	}

	// Split the version by '.' and parse the major version number.
	parts := strings.Split(tag, ".")
	if len(parts) == 0 {
		return latestVersion, fmt.Errorf("invalid tag format")
	}
	major, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return latestVersion, fmt.Errorf("failed to parse major version from tag %q: %v", release.TagName, err)
	}
	latestVersion = major

	return latestVersion, nil
}
