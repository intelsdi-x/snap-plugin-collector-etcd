# Snap collector plugin - etcd
This plugin collects metrics from etcd's `/metrics` endpoint.

It's used in the [Snap framework](http://github.com/intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements
* [etcd](https://github.com/coreos/etcd)
* [golang 1.5+](https://golang.org/dl/)

### Operating systems
All OSs currently supported by Snap:
* Linux/amd64
* Darwin/amd64

### Installation

#### Download the plugin binary:

You can get the pre-built binaries for your OS and architecture from the plugin's [GitHub Releases](https://github.com/intelsdi-x/snap-plugin-collector-etcd/releasess) page. Download the plugin from the latest release and load it into `snapd` (`/opt/snap/plugins` is the default location for Snap packages).

#### To build the plugin binary:

Fork https://github.com/intelsdi-x/snap-plugin-collector-etcd
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-etcd.git
```

Build the Snap etcd plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `./build/`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Load the plugin and create a task, see example in [Examples](https://github.com/intelsdi-x/snap-plugin-collector-etcd/blob/master/README.md#examples).

## Documentation

### Collected Metrics
This plugin collects the raw counters and a few other metrics from etcd.  It is focused on the ones which are easily parsable and excludes ones which are formatted specifically for a Prometheus metric.  There are also a handful of derivatives which take a sum over count average.

Here are the derivatives:

```
/intel/etcd/derivative/etcd_server_proposal_durations_avg
/intel/etcd/derivative/etcd_snapshot_save_marshalling_durations_seconds_avg
/intel/etcd/derivative/etcd_storage_db_compaction_pause_duration_milliseconds_avg
/intel/etcd/derivative/etcd_storage_db_compaction_total_duration_milliseconds_avg
/intel/etcd/derivative/etcd_storage_index_compaction_pause_duration_milliseconds_avg
/intel/etcd/derivative/etcd_wal_fsync_durations_seconds_avg
```

And the raw metrics:

```
/intel/etcd/etcd_server_file_descriptors_used_total
/intel/etcd/etcd_server_pending_proposal_total
/intel/etcd/etcd_server_proposal_durations_seconds_count
/intel/etcd/etcd_server_proposal_durations_seconds_sum
/intel/etcd/etcd_server_proposal_failed_total
/intel/etcd/etcd_snapshot_save_marshalling_durations_seconds_count
/intel/etcd/etcd_snapshot_save_marshalling_durations_seconds_sum
/intel/etcd/etcd_snapshot_save_total_durations_seconds_count
/intel/etcd/etcd_snapshot_save_total_durations_seconds_sum
/intel/etcd/etcd_storage_db_compaction_pause_duration_milliseconds_count
/intel/etcd/etcd_storage_db_compaction_pause_duration_milliseconds_sum
/intel/etcd/etcd_storage_db_compaction_total_duration_milliseconds_count
/intel/etcd/etcd_storage_db_compaction_total_duration_milliseconds_sum
/intel/etcd/etcd_storage_db_total_size_in_bytes
/intel/etcd/etcd_storage_delete_total
/intel/etcd/etcd_storage_events_total
/intel/etcd/etcd_storage_index_compaction_pause_duration_milliseconds_count
/intel/etcd/etcd_storage_index_compaction_pause_duration_milliseconds_sum
/intel/etcd/etcd_storage_keys_total
/intel/etcd/etcd_storage_pending_events_total
/intel/etcd/etcd_storage_put_total
/intel/etcd/etcd_storage_range_total
/intel/etcd/etcd_storage_slow_watcher_total
/intel/etcd/etcd_storage_txn_total
/intel/etcd/etcd_storage_watch_stream_total
/intel/etcd/etcd_storage_watcher_total
/intel/etcd/etcd_store_expires_total
/intel/etcd/etcd_store_watch_requests_total
/intel/etcd/etcd_store_watchers
/intel/etcd/etcd_wal_fsync_durations_seconds_count
/intel/etcd/etcd_wal_fsync_durations_seconds_sum
/intel/etcd/etcd_wal_last_index_saved
/intel/etcd/go_gc_duration_seconds_count
/intel/etcd/go_gc_duration_seconds_sum
/intel/etcd/go_goroutines
/intel/etcd/go_memstats_alloc_bytes
/intel/etcd/go_memstats_alloc_bytes_total
/intel/etcd/go_memstats_buck_hash_sys_bytes
/intel/etcd/go_memstats_frees_total
/intel/etcd/go_memstats_gc_sys_bytes
/intel/etcd/go_memstats_heap_alloc_bytes
/intel/etcd/go_memstats_heap_idle_bytes
/intel/etcd/go_memstats_heap_inuse_bytes
/intel/etcd/go_memstats_heap_objects
/intel/etcd/go_memstats_heap_released_bytes_total
/intel/etcd/go_memstats_heap_sys_bytes
/intel/etcd/go_memstats_last_gc_time_seconds
/intel/etcd/go_memstats_lookups_total
/intel/etcd/go_memstats_mallocs_total
/intel/etcd/go_memstats_mcache_inuse_bytes
/intel/etcd/go_memstats_mcache_sys_bytes
/intel/etcd/go_memstats_mspan_inuse_bytes
/intel/etcd/go_memstats_mspan_sys_bytes
/intel/etcd/go_memstats_next_gc_bytes
/intel/etcd/go_memstats_other_sys_bytes
/intel/etcd/go_memstats_stack_inuse_bytes
/intel/etcd/go_memstats_stack_sys_bytes
/intel/etcd/go_memstats_sys_bytes
/intel/etcd/process_cpu_seconds_total
/intel/etcd/process_max_fds
/intel/etcd/process_open_fds
/intel/etcd/process_resident_memory_bytes
/intel/etcd/process_start_time_seconds
/intel/etcd/process_virtual_memory_bytes
```

### Examples
Example of running Snap etcd collector and writing data to file.

Ensure [Snap daemon is running](https://github.com/intelsdi-x/snap#running-snap):
* initd: `service snap-telemetry start`
* systemd: `systemctl start snap-telemetry`
* command line: `snapd -l 1 -t 0 &`

Download and load Snap plugins:
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-etcd/latest/linux/x86_64/snap-plugin-collector-etcd
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest/linux/x86_64/snap-plugin-publisher-file
$ chmod 755 snap-plugin-*
$ snapctl plugin load snap-plugin-collector-etcd
$ snapctl plugin load snap-plugin-publisher-file
```

See all available metrics:

```
$ snapctl metric list
```

Download an [example task file](https://github.com/intelsdi-x/snap-plugin-collector-etcd/blob/master/examples/tasks/) and load it:
```
$ curl -sfLO https://raw.githubusercontent.com/intelsdi-x/snap-plugin-collector-etcd/master/examples/tasks/task-etcd.json
$ snapctl task create -t task-etcd.json
Using task manifest to create task
Task created
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
Name: Task-02dd7ff4-8106-47e9-8b86-70067cd0a850
State: Running
```

See realtime output from `snapctl task watch <task_id>` (CTRL+C to exit)
```
snapctl --url http://localhost:8182 task watch 02dd7ff4-8106-47e9-8b86-70067cd0a850
Watching Task (02dd7ff4-8106-47e9-8b86-70067cd0a850):
NAMESPACE                                                        DATA                    TIMESTAMP                                       SOURCE
/intel/etcd/derivative/etcd_wal_fsync_durations_seconds_avg      0.0004962833257018023   2016-01-11 16:33:22.728866474 -0800 PST         snap1
/intel/etcd/etcd_wal_fsync_durations_seconds_count               21053                   2016-01-11 16:33:22.728717791 -0800 PST         snap1
/intel/etcd/etcd_wal_fsync_durations_seconds_sum                 10.448252856000042      2016-01-11 16:33:22.728715613 -0800 PST         snap1
```

This data is published to a file `/tmp/published_etcd.log` per task specification

Stop task:
```
$ snapctl task stop 02dd7ff4-8106-47e9-8b86-70067cd0a850
Task stopped:
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-etcd/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-etcd/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap.

To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support).

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[Snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [@danielscottt](https://github.com/danielscottt/)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
