package metadata

import (
	"encoding/json"
)

func (metadata MetadataOffChain) CreateURL() string {
	bytes, _ := json.Marshal(metadata)

	return string(bytes)
}
