package quota

import (
	"encoding/xml"
	"fmt"
	"io"
)

type QuotaInfo struct {
	XMLNS      string   `xml:"xmlns,attr,omitempty"`
	XMLName    xml.Name `xml:"QuotaInfo"`
	MaxSpace   int64    `xml:"maxSpace"`
	MaxInodes  int64    `xml:"maxInodes"`
	UsedInodes int64    `xml:"usedInodes"`
	UsedSpace  int64    `xml:"usedSpace"`
	Dpath      string   `xml:"dpath"`
}

func ParseQuotaInfoConfig(r io.Reader) (*QuotaInfo, error) {
	var config QuotaInfo
	err := xml.NewDecoder(r).Decode(&config)
	if err != nil {
		return nil, err
	}
	if config.MaxSpace == 0 || config.MaxInodes == 0 {
		return &config, fmt.Errorf("max space space must > 0 and max inodes space must >0")
	}
	return &config, nil
}
