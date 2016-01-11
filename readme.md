Snap Etcd Collector
===================

This collector retrieves metrics from Etcd's `/metrics` endpoint.  At this point it time, it is only returning the counters, as the bulk of the metrics available are either summaries or historgams which are specially formatted for Prometheus consumption.

```
NAMESPACE                                                                                VERSIONS
/intel/etcd/derivative/etcd_server_proposal_durations_avg                                1
/intel/etcd/derivative/etcd_snapshot_save_marshalling_durations_seconds_avg              1
/intel/etcd/derivative/etcd_storage_db_compaction_pause_duration_milliseconds_avg        1
/intel/etcd/derivative/etcd_storage_db_compaction_total_duration_milliseconds_avg        1
/intel/etcd/derivative/etcd_storage_index_compaction_pause_duration_milliseconds_avg     1
/intel/etcd/derivative/etcd_wal_fsync_durations_seconds_avg                              1
/intel/etcd/etcd_server_file_descriptors_used_total                                      1
/intel/etcd/etcd_server_pending_proposal_total                                           1
/intel/etcd/etcd_server_proposal_durations_seconds_count                                 1
/intel/etcd/etcd_server_proposal_durations_seconds_sum                                   1
/intel/etcd/etcd_server_proposal_failed_total                                            1
/intel/etcd/etcd_snapshot_save_marshalling_durations_seconds_count                       1
/intel/etcd/etcd_snapshot_save_marshalling_durations_seconds_sum                         1
/intel/etcd/etcd_snapshot_save_total_durations_seconds_count                             1
/intel/etcd/etcd_snapshot_save_total_durations_seconds_sum                               1
/intel/etcd/etcd_storage_db_compaction_pause_duration_milliseconds_count                 1
/intel/etcd/etcd_storage_db_compaction_pause_duration_milliseconds_sum                   1
/intel/etcd/etcd_storage_db_compaction_total_duration_milliseconds_count                 1
/intel/etcd/etcd_storage_db_compaction_total_duration_milliseconds_sum                   1
/intel/etcd/etcd_storage_db_total_size_in_bytes                                          1
/intel/etcd/etcd_storage_delete_total                                                    1
/intel/etcd/etcd_storage_events_total                                                    1
/intel/etcd/etcd_storage_index_compaction_pause_duration_milliseconds_count              1
/intel/etcd/etcd_storage_index_compaction_pause_duration_milliseconds_sum                1
/intel/etcd/etcd_storage_keys_total                                                      1
/intel/etcd/etcd_storage_pending_events_total                                            1
/intel/etcd/etcd_storage_put_total                                                       1
/intel/etcd/etcd_storage_range_total                                                     1
/intel/etcd/etcd_storage_slow_watcher_total                                              1
/intel/etcd/etcd_storage_txn_total                                                       1
/intel/etcd/etcd_storage_watch_stream_total                                              1
/intel/etcd/etcd_storage_watcher_total                                                   1
/intel/etcd/etcd_store_expires_total                                                     1
/intel/etcd/etcd_store_watch_requests_total                                              1
/intel/etcd/etcd_store_watchers                                                          1
/intel/etcd/etcd_wal_fsync_durations_seconds_count                                       1
/intel/etcd/etcd_wal_fsync_durations_seconds_sum                                         1
/intel/etcd/etcd_wal_last_index_saved                                                    1
/intel/etcd/go_gc_duration_seconds_count                                                 1
/intel/etcd/go_gc_duration_seconds_sum                                                   1
/intel/etcd/go_goroutines                                                                1
/intel/etcd/go_memstats_alloc_bytes                                                      1
/intel/etcd/go_memstats_alloc_bytes_total                                                1
/intel/etcd/go_memstats_buck_hash_sys_bytes                                              1
/intel/etcd/go_memstats_frees_total                                                      1
/intel/etcd/go_memstats_gc_sys_bytes                                                     1
/intel/etcd/go_memstats_heap_alloc_bytes                                                 1
/intel/etcd/go_memstats_heap_idle_bytes                                                  1
/intel/etcd/go_memstats_heap_inuse_bytes                                                 1
/intel/etcd/go_memstats_heap_objects                                                     1
/intel/etcd/go_memstats_heap_released_bytes_total                                        1
/intel/etcd/go_memstats_heap_sys_bytes                                                   1
/intel/etcd/go_memstats_last_gc_time_seconds                                             1
/intel/etcd/go_memstats_lookups_total                                                    1
/intel/etcd/go_memstats_mallocs_total                                                    1
/intel/etcd/go_memstats_mcache_inuse_bytes                                               1
/intel/etcd/go_memstats_mcache_sys_bytes                                                 1
/intel/etcd/go_memstats_mspan_inuse_bytes                                                1
/intel/etcd/go_memstats_mspan_sys_bytes                                                  1
/intel/etcd/go_memstats_next_gc_bytes                                                    1
/intel/etcd/go_memstats_other_sys_bytes                                                  1
/intel/etcd/go_memstats_stack_inuse_bytes                                                1
/intel/etcd/go_memstats_stack_sys_bytes                                                  1
/intel/etcd/go_memstats_sys_bytes                                                        1
/intel/etcd/process_cpu_seconds_total                                                    1
/intel/etcd/process_max_fds                                                              1
/intel/etcd/process_open_fds                                                             1
/intel/etcd/process_resident_memory_bytes                                                1
/intel/etcd/process_start_time_seconds                                                   1
/intel/etcd/process_virtual_memory_bytes                                                 1
```
