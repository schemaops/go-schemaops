package keyspace

import "testing"

func TestCreateStatement(t *testing.T) {

	keyspace := Keyspace{
		Name:          "user",
		DurableWrites: true,
		Replication:   "{ class = 'SimpleStrategy' }",
	}

	expected := "CREATE KEYSPACE user WITH REPLICATION = { class = 'SimpleStrategy' } AND DURABLE_WRITES = true;"
	actual, err := CreateStatement(&keyspace)

	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}

func TestUpdateStatement(t *testing.T) {

	keyspace1 := Keyspace{
		Name:          "user",
		DurableWrites: true,
		Replication:   "{ class = 'SimpleStrategy' }",
	}

	keyspace2 := Keyspace{
		Name:          "user",
		DurableWrites: false,
		Replication:   "{ class = 'SimpleStrategy' }",
	}

	keyspace3 := Keyspace{
		Name:          "user",
		DurableWrites: false,
		Replication:   "{ class = 'NetworkTopologyStrategy' }",
	}

	var testMatrix = []struct {
		OldKeyspace       *Keyspace
		NewKeyspace       *Keyspace
		ExpectedStatement string
	}{
		{
			OldKeyspace:       &keyspace1,
			NewKeyspace:       &keyspace2,
			ExpectedStatement: "ALTER KEYSPACE user WITH DURABLE_WRITES = false;",
		},
		{
			OldKeyspace:       &keyspace1,
			NewKeyspace:       &keyspace3,
			ExpectedStatement: "ALTER KEYSPACE user WITH REPLICATION = { class = 'NetworkTopologyStrategy' } AND DURABLE_WRITES = false;",
		},
		{
			OldKeyspace:       &keyspace3,
			NewKeyspace:       &keyspace2,
			ExpectedStatement: "ALTER KEYSPACE user WITH REPLICATION = { class = 'SimpleStrategy' };",
		},
	}

	for _, scenario := range testMatrix {
		actual, err := UpdateStatement(scenario.OldKeyspace, scenario.NewKeyspace)

		if err != nil {
			t.Error(err)
		}

		if actual != scenario.ExpectedStatement {
			t.Errorf("expected %s but got %s", scenario.ExpectedStatement, actual)
		}
	}
}
