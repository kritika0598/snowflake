## Snowflake

To create 64 bit ids which are time sortable and can be created at large scale, twitter proposed [snowflake](https://blog.twitter.com/engineering/en_us/a/2010/announcing-snowflake.html).

To generate the roughly-sorted 64 bit ids in an uncoordinated manner, we settled on a composition of: timestamp, worker number and sequence number.

Sequence numbers are per-thread and worker numbers are chosen at startup via zookeeper (though that’s overridable via a config file).

- Time — ***41*** bits (millisecond precision w/a custom epoch of Twitter gives 69 years)
- Configured machine id — ***10*** bits — gives us up to 1024 machines
- Sequence number — ***12*** bits — rolls over every 4096 per machine (with protection to avoid rollover in the same ms)
- 1 bit is reserved for future use and value is set as 0

## Run
`go run main.go`