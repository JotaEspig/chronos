package config

import "github.com/microcosm-cc/bluemonday"

var StrictPolicy = bluemonday.StrictPolicy()
var UGCPolicy = bluemonday.UGCPolicy()
