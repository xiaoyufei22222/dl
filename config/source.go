package config

import "github.com/xiaoyufei22222/proxypool-luo/pkg/tool"

type Source struct {
	Type    string       `json:"type" yaml:"type"`
	Options tool.Options `json:"options" yaml:"options"`
}
