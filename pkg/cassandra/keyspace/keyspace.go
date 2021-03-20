package keyspace

type Keyspace struct {
	Name          string `hcl:"name"`
	DurableWrites bool   `hcl:"durable_writes"`
	Replication   string `hcl:"replication"`
}