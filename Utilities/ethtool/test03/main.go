package main

import (
	"fmt"
	"github.com/safchain/ethtool"
	"strings"
)

func main() {
	ethHandle, err := ethtool.NewEthtool()
	defer ethHandle.Close()
	if err != nil {
		panic(err)
	}

	intf := "wlp0s20f3"
	var config = map[string]bool{
		"rx-gro": true,
	}
	//var config = map[string]bool{
	//	"generic-receive-offload": false,
	//}
	////
	err = ethHandle.Change(intf, config)
	if err != nil {
		panic(err)
	}

	result, err := ethHandle.Features(intf)
	//result, err := ethHandle.FeatureNames(intf)
	if err != nil {
		panic(err)
	}

	for k, v := range result {
		if strings.Contains(k, "gro") {
			fmt.Printf("(K,V)=('%s','%v')\n", k, v)
		}
	}

	//bob := result["Xgeneric-receive-offload"]
	//fmt.Printf("sgbFeatures:'%v'\n", bob)
}
