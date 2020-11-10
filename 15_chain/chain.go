package chain

import (
	"strings"
)

// SensitiveWordFilter is filter behave
type SensitiveWordFilter interface {
	Filter(content string) bool
}

// SensitiveWordFilterChain is filter chain
type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

// AddFilter for add new Filter
func (swfc *SensitiveWordFilterChain) AddFilter(f SensitiveWordFilter) {
	swfc.filters = append(swfc.filters, f)
}

// Filter is method of SensitiveWordFilterChain
func (swfc *SensitiveWordFilterChain) Filter(content string) bool {
	for _, swf := range swfc.filters {
		if swf.Filter(content) {
			return true
		}
	}
	return false
}

// AdSensitiveWordFilter is SensitiveWordFilter
type AdSensitiveWordFilter struct{}

// Filter is method of AdSensitiveWordFilter
func (aswf *AdSensitiveWordFilter) Filter(content string) bool {
	return strings.Contains(content, "taobao")
}

// PoliticalWordFilter is SensitiveWordFilter
type PoliticalWordFilter struct{}

// Filter is method of PoliticalWordFilter
func (pwf *PoliticalWordFilter) Filter(content string) bool {
	return strings.Contains(content, "china")
}
