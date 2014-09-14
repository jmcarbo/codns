package codns

// The configuration file is a JSON file with a simple structure; the following
// configuration specify 3 forwarders: *.example.com and 10.1.2.* will be
// resolved by OpenDNS and a catch-all for resolving with Google DNS.
//
// {
//		"filters": [
// 				{
// 						"pattern": "example.com.",
// 						"addresses": [ "208.67.222.222", "208.67.220.220" ]
// 				},
// 				{
// 						"pattern": "2.1.10.in-addr.arpa.",
// 						"addresses": [ "208.67.222.222", "208.67.220.220" ]
// 				},
// 				{
// 						"pattern": ".",
// 						"addresses": [ "8.8.8.8" ]
// 				}
// 		]
// }
//

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type Configuration struct {
	Filters []ConfigFilter
}

type ConfigFilter struct {
	Pattern   string
	Addresses []string
}

const default_config = `
{
  "filters": [
    {
      "pattern": "consul.",
      "addresses": [ "127.0.0.1:8600" ]
    },
    {
      "pattern": ".",
      "addresses": [ "8.8.8.8", "8.8.4.4" ]
    }
  ]
}
`
// ReadConfig reads a JSON file and returns a Configuration object
// containing the raw elements.
func ReadConfig(filename string) Configuration {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
    file = []byte(default_config)
	}

	var jsonConfig Configuration
	json.Unmarshal(file, &jsonConfig)

	// Safety checks
	if len(jsonConfig.Filters) == 0 {
		log.Fatalf("Configuration contains no 'filters' section")
	}

	for _, filter := range jsonConfig.Filters {
		if filter.Pattern == "" || len(filter.Addresses) == 0 {
			log.Fatalf("Filter error: missing pattern or empty server list")
		}

		for i, address := range filter.Addresses {
			if !strings.Contains(address, ":") {
				filter.Addresses[i] = strings.Join([]string{address, "53"}, ":")
			}
		}
	}

	return jsonConfig
}
