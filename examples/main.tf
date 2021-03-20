resource "schemaops_cassandra_keyspace" "user" {
  name = "user"
  durable_writes = true
  replication = "{class = 'SimpleStrategy'}"
}