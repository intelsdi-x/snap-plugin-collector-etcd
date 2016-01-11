/*
http://www.apache.org/licenses/LICENSE-2.0.txt
Copyright 2015 Intel Corporation
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package etcd

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core/ctypes"
)

const (
	// Name of plugin
	Name = "etcd"
	// Version of plugin
	Version = 1
	// Type of plugin
	Type = plugin.CollectorPluginType
)

var (
	errNoHost    = errors.New("getting metric types requires an etcd host")
	errBadHost   = errors.New("failed to parse given etcd_host")
	errReqFailed = errors.New("request to etcd api failed")
)

type Etcd struct{}

func (e *Etcd) CollectMetrics(mts []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {
	config := mts[0].Config().Table()
	hostcfg, ok := config["etcd_host"]
	if !ok {
		return nil, errNoHost
	}
	host, ok := hostcfg.(ctypes.ConfigValueStr)
	if !ok {
		return nil, errBadHost
	}

	filter := make([]string, len(mts))
	for i, m := range mts {
		filter[i] = m.Namespace_[len(m.Namespace_)-1]
	}

	return gathermts(host.Value, filter)
}

//GetMetricTypes returns metric types for testing
func (e *Etcd) GetMetricTypes(cfg plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {
	hostcfg, ok := cfg.Table()["etcd_host"]
	if !ok {
		return nil, errNoHost
	}
	host, ok := hostcfg.(ctypes.ConfigValueStr)
	if !ok {
		return nil, errBadHost
	}

	return gathermts(host.Value, []string{})
}

func (e *Etcd) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	c := cpolicy.New()
	rule, _ := cpolicy.NewStringRule("etcd_host", true)
	p := cpolicy.NewPolicyNode()
	p.Add(rule)
	c.Add([]string{"intel", "etcd"}, p)
	return c, nil
}

//Meta returns meta data for testing
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(
		Name,
		Version,
		Type,
		[]string{plugin.SnapGOBContentType},
		[]string{plugin.SnapGOBContentType},
		plugin.Unsecure(true),
		plugin.RoutingStrategy(plugin.DefaultRouting),
	)
}

func gathermts(host string, filter []string) ([]plugin.PluginMetricType, error) {
	resp, err := http.Get(fmt.Sprintf("%s/metrics", host))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errReqFailed
	}

	var mts []plugin.PluginMetricType
	scanner := bufio.NewScanner(resp.Body)

	hn, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	for scanner.Scan() {
		txt := scanner.Text()
		if !strings.Contains(txt, "{") && !strings.Contains(txt, "#") {
			nsslice := strings.Split(txt, " ")
			if len(filter) != 0 {
				for _, f := range filter {
					if strings.Contains(nsslice[0], f) {
						mts = append(mts, plugin.PluginMetricType{
							Namespace_: []string{"intel", "etcd", nsslice[0]},
							Data_:      nsslice[1],
							Source_:    hn,
							Timestamp_: time.Now(),
						})
					}
				}
			} else {
				mts = append(mts, plugin.PluginMetricType{
					Namespace_: []string{"intel", "etcd", nsslice[0]},
					Data_:      nsslice[1],
					Source_:    hn,
					Timestamp_: time.Now(),
				})
			}
		}
	}
	return mts, nil
}
