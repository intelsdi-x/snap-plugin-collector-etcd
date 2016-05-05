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
	"strconv"
	"strings"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
	"github.com/intelsdi-x/snap/core/ctypes"
)

const (
	// Name of plugin
	Name = "etcd"
	// Version of plugin
	Version = 2
	// Type of plugin
	Type = plugin.CollectorPluginType
)

var (
	errNoHost    = errors.New("getting metric types requires an etcd host")
	errBadHost   = errors.New("failed to parse given etcd_host")
	errReqFailed = errors.New("request to etcd api failed")

	// the derivatives are a sum / count avg
	derivatives = map[string][2]string{
		"etcd_server_proposal_durations_avg": [2]string{
			"etcd_server_proposal_durations_seconds_sum",
			"etcd_server_proposal_durations_seconds_count",
		},
		"etcd_snapshot_save_marshalling_durations_seconds_avg": [2]string{
			"etcd_snapshot_save_marshalling_durations_seconds_sum",
			"etcd_snapshot_save_marshalling_durations_seconds_count",
		},
		"etcd_storage_db_compaction_pause_duration_milliseconds_avg": [2]string{
			"etcd_storage_db_compaction_pause_duration_milliseconds_sum",
			"etcd_storage_db_compaction_pause_duration_milliseconds_count",
		},
		"etcd_storage_db_compaction_total_duration_milliseconds_avg": [2]string{
			"etcd_storage_db_compaction_total_duration_milliseconds_sum",
			"etcd_storage_db_compaction_total_duration_milliseconds_count",
		},
		"etcd_storage_index_compaction_pause_duration_milliseconds_avg": [2]string{
			"etcd_storage_index_compaction_pause_duration_milliseconds_sum",
			"etcd_storage_index_compaction_pause_duration_milliseconds_count",
		},
		"etcd_wal_fsync_durations_seconds_avg": [2]string{
			"etcd_wal_fsync_durations_seconds_sum",
			"etcd_wal_fsync_durations_seconds_count",
		},
	}
)

type Etcd struct{}

func (e *Etcd) CollectMetrics(mts []plugin.MetricType) ([]plugin.MetricType, error) {
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
		filter[i] = m.Namespace().Strings()[len(m.Namespace().Strings())-1]
	}

	return gathermts(host.Value, filter)
}

//GetMetricTypes returns metric types for testing
func (e *Etcd) GetMetricTypes(cfg plugin.ConfigType) ([]plugin.MetricType, error) {
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

func gathermts(host string, filter []string) ([]plugin.MetricType, error) {
	resp, err := http.Get(fmt.Sprintf("%s/metrics", host))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errReqFailed
	}

	mtsmap := make(map[string]plugin.MetricType)
	scanner := bufio.NewScanner(resp.Body)

	if err != nil {
		return nil, err
	}

	for scanner.Scan() {
		txt := scanner.Text()
		if !strings.Contains(txt, "{") && !strings.Contains(txt, "#") {
			nsslice := strings.Split(txt, " ")
			mtsmap[nsslice[0]] = plugin.MetricType{
				Namespace_: core.NewNamespace("intel", "etcd", nsslice[0]),
				Data_:      nsslice[1],
				Timestamp_: time.Now(),
			}
		}
	}

	// No filter given; this was a GetMetricTypes call.
	if len(filter) == 0 {
		mts := make([]plugin.MetricType, 0, len(mtsmap)+len(derivatives))
		for _, v := range mtsmap {
			mts = append(mts, v)
		}
		for k := range derivatives {
			mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace("intel", "etcd", "derivative", k)})
		}
		return mts, nil
	}

	// Walk through filter and pluck out metrics.
	// if we find the requested metric in derivatives,
	// then derive the value from `from`.
	mts := make([]plugin.MetricType, 0, len(filter))
	for _, f := range filter {
		from, ok := derivatives[f]
		if ok {
			mt := plugin.MetricType{
				Namespace_: core.NewNamespace("intel", "etcd", "derivative", f),
				Timestamp_: time.Now(),
			}
			sum, err := strconv.ParseFloat(mtsmap[from[0]].Data_.(string), 64)
			if err != nil {
				return nil, err
			}
			count, err := strconv.ParseFloat(mtsmap[from[1]].Data_.(string), 64)
			if err != nil {
				return nil, err
			}
			mt.Data_ = sum / count
			mts = append(mts, mt)
			continue
		}

		mt, ok := mtsmap[f]
		if ok {
			mts = append(mts, mt)
		}
	}
	return mts, nil
}
