# DISCONTINUATION OF PROJECT 

**This project will no longer be maintained by Intel.  Intel will not provide or guarantee development of or support for this project, including but not limited to, maintenance, bug fixes, new releases or updates.  Patches to this project are no longer accepted by Intel. If you have an ongoing need to use this project, are interested in independently developing it, or would like to maintain patches for the community, please create your own fork of the project.**


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

You can get the pre-built binaries for your OS and architecture from the plugin's [GitHub Releases](https://github.com/intelsdi-x/snap-plugin-collector-etcd/releases) page. Download the plugin from the latest release and load it into `snapteld` (`/opt/snap/plugins` is the default location for Snap packages).

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
* Load the plugin and create a task, see example in [Examples](#examples).

## Documentation

### Collected Metrics
This plugin collects the raw counters and a few other metrics from etcd.  It is focused on the ones which are easily parsable and excludes ones which are formatted specifically for a Prometheus metric.  There are also a handful of derivatives which take a sum over count average.
List of metrics collected by this plugin can be found in [METRICS.md file](METRICS.md).

### Examples
Example of running Snap etcd collector and writing data to file.

Ensure [Snap daemon is running](https://github.com/intelsdi-x/snap#running-snap):
* initd: `service snap-telemetry start`
* systemd: `systemctl start snap-telemetry`
* command line: `snapteld -l 1 -t 0 &`

Download and load Snap plugins:
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-etcd/latest/linux/x86_64/snap-plugin-collector-etcd
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest/linux/x86_64/snap-plugin-publisher-file
$ chmod 755 snap-plugin-*
$ snaptel plugin load snap-plugin-collector-etcd
$ snaptel plugin load snap-plugin-publisher-file
```

See all available metrics:

```
$ snaptel metric list
```

Download an [example task file](https://github.com/intelsdi-x/snap-plugin-collector-etcd/blob/master/examples/tasks/task-etcd.json) and load it:
```
$ curl -sfLO https://raw.githubusercontent.com/intelsdi-x/snap-plugin-collector-etcd/master/examples/tasks/task-etcd.json
$ snaptel task create -t task-etcd.json
Using task manifest to create task
Task created
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
Name: Task-02dd7ff4-8106-47e9-8b86-70067cd0a850
State: Running
```

See realtime output from `snaptel task watch <task_id>` (CTRL+C to exit)
```
snaptel --url http://localhost:8182 task watch 02dd7ff4-8106-47e9-8b86-70067cd0a850
Watching Task (02dd7ff4-8106-47e9-8b86-70067cd0a850):
NAMESPACE                                                        DATA                    TIMESTAMP                                       SOURCE
/intel/etcd/derivative/etcd_wal_fsync_durations_seconds_avg      0.0004962833257018023   2016-01-11 16:33:22.728866474 -0800 PST         snap1
/intel/etcd/etcd_wal_fsync_durations_seconds_count               21053                   2016-01-11 16:33:22.728717791 -0800 PST         snap1
/intel/etcd/etcd_wal_fsync_durations_seconds_sum                 10.448252856000042      2016-01-11 16:33:22.728715613 -0800 PST         snap1
```

This data is published to a file `/tmp/published_etcd.log` per task specification

Stop task:
```
$ snaptel task stop 02dd7ff4-8106-47e9-8b86-70067cd0a850
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
