# Mesosbeat

Mesosbeat is an [Elastic Beat](https://www.elastic.co/products/beats) that reads stats from the Mesos and indexes them into Elasticsearch or Logstash.

## Description

> [Apache Mesos](http://mesos.apache.org/) is a cluster manager that provides efficient resource isolation and sharing across distributed applications or frameworks.

## Configuration

Adjust the `mesosbeat.yml` configuration file to your needs.

### `period`
Defines how often to read metrics. Default to `30` s.

### `port`
Defines the envoy proxy port serviced. Default to `5050`

### `host`
Host name of ElasticSearch. Default to `localhost`

## Document Example

<pre>

{
  "master": {
    "allocator": {
      "mesos": {
        "allocation_run_latency_ms": 0.166912,
        "allocation_run_latency_MS": {
          "p50": 0.153088,
          "p90": 0.165888,
          "p999": 0.34411084799999886,
          "p95": 0.175872,
          "min": 0.036096,
          "max": 0.390912,
          "count": 1000,
          "p99": 0.19791872,
          "p9999": 0.3862318848000031
        },
        "event_queue_dispatches": 6,
        "resources": {
          "disk": {
            "offered_or_allocated": 0,
            "total": 172196
          },
          "mem": {
            "total": 14886,
            "offered_or_allocated": 0
          },
          "cpus": {
            "offered_or_allocated": 0,
            "total": 12
          }
        },
        "allocation_run_ms": 0.231936,
        "allocation_run_MS": {
          "max": 0.323072,
          "p50": 0.228096,
          "p90": 0.2439936,
          "count": 1000,
          "p95": 0.26039039999999997,
          "p999": 0.3100290559999997,
          "min": 0.070144,
          "p9999": 0.3217677056000009,
          "p99": 0.28597248
        },
        "allocation_runs": 2110
      }
    },
    "master": {
      "valid_executor_to_framework_messages": 0,
      "tasks_unreachable": 0,
      "gpus_used": 0,
      "gpus_revocable_total": 0,
      "tasks_staging": 0,
      "messages_revive_offers": 0,
      "dropped_messages": 0,
      "messages_launch_tasks": 0,
      "messages_deactivate_framework": 0,
      "mem_percent": 0,
      "messages_reconcile_operations": 0,
      "frameworks_active": 0,
      "disk_revocable_percent": 0,
      "slave_registrations": 0,
      "valid_framework_to_executor_messages": 0,
      "messages_register_framework": 0,
      "messages_authenticate": 0,
      "frameworks_connected": 0,
      "invalid_status_update_acknowledgements": 0,
      "messages_resource_request": 0,
      "tasks_starting": 0,
      "messages_exited_executor": 0,
      "tasks_finished": 0,
      "outstanding_offers": 0,
      "event_queue_dispatches": 8,
      "messages_decline_offers": 0,
      "messages_reregister_slave": 1,
      "messages_reconcile_tasks": 0,
      "gpus_revocable_used": 0,
      "mem_total": 14886,
      "disk_revocable_used": 0,
      "messages_operation_status_update_acknowledgement": 0,
      "valid_status_update_acknowledgements": 0,
      "messages_framework_to_executor": 0,
      "elected": 1,
      "messages_status_update": 0,
      "cpus_revocable_total": 0,
      "invalid_status_updates": 0,
      "invalid_framework_to_executor_messages": 0,
      "cpus_revocable_used": 0,
      "invalid_executor_to_framework_messages": 0,
      "tasks_failed": 0,
      "operator_event_stream_subscribers": 0,
      "disk_percent": 0,
      "messages_kill_task": 0,
      "slave_unreachable_scheduled": 0,
      "messages_unregister_framework": 0,
      "cpus_used": 0,
      "gpus_revocable_percent": 0,
      "tasks_killing": 0,
      "cpus_total": 12,
      "messages_reregister_framework": 0,
      "disk_total": 172196,
      "recovery_slave_removals": 0,
      "uptime_secs": 2113.715932928,
      "tasks_running": 0,
      "tasks_killed": 0,
      "slave_Removals": {
        "reason_registered": 0,
        "reason_unhealthy": 0,
        "reason_unregistered": 0
      },
      "cpus_revocable_percent": 0,
      "slaves_inactive": 0,
      "messages_update_slave": 1,
      "messages_register_slave": 0,
      "event_queue_http_requests": 0,
      "cpus_percent": 0,
      "messages_executor_to_framework": 0,
      "frameworks_disconnected": 0,
      "mem_used": 0,
      "messages_status_update_acknowledgement": 0,
      "slave_unreachable_canceled": 0,
      "slave_reregistrations": 1,
      "tasks_lost": 0,
      "event_queue_messages": 0,
      "messages_unregister_slave": 0,
      "slave_unreachable_completed": 0,
      "disk_revocable_total": 0,
      "slaves_connected": 1,
      "gpus_total": 0,
      "slaves_disconnected": 0,
      "frameworks_inactive": 0,
      "mem_revocable_total": 0,
      "valid_status_updates": 0,
      "slaves_active": 1,
      "invalid_operation_status_update_acknowledgements": 0,
      "disk_used": 0,
      "gpus_percent": 0,
      "slaves_unreachable": 0,
      "tasks_error": 0,
      "mem_revocable_percent": 0,
      "slave_removals": 0,
      "mem_revocable_used": 0,
      "valid_operation_status_update_acknowledgements": 0
    },
    "system": {
      "cpus_total": 12,
      "load_15min": 0.37,
      "load_5min": 0.52,
      "load_1min": 0.63,
      "mem_free_bytes": 9097818112,
      "mem_total_bytes": 16683503616
    },
    "registrar": {
      "log": {
        "ensemble_size": 1,
        "recovered": 1
      },
      "state_fetch_ms": 28.32896,
      "state_store_ms": 18.193152,
      "state_store_MS": {}
    }
}

</pre>

## Getting Started with Mesosbeat

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/berfinsari/mesosbeat`

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Mesosbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Mesosbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/berfinsari/mesosbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Mesosbeat run the command below. This will generate a binary
in the same directory with the name mesosbeat.

```
make
```


### Run

To run Mesosbeat with debugging output enabled, run:

```
./mesosbeat -c mesosbeat.yml -e -d "*"
```


### Test

To test Mesosbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  Mesosbeat source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Mesosbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/berfinsari/mesosbeat
git clone https://github.com/berfinsari/mesosbeat ${GOPATH}/src/github.com/berfinsari/mesosbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
