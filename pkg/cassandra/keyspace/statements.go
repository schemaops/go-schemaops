package keyspace

import (
	"bytes"
	"strings"
	"text/template"
)

func CreateStatement(keyspace *Keyspace) (string, error) {

	t, err := template.New("cassandra_keyspace_create").Parse("CREATE KEYSPACE {{ .Name }} WITH REPLICATION = {{ .Replication }} AND DURABLE_WRITES = {{ .DurableWrites }};")

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, keyspace)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func UpdateStatement(oldKeyspace *Keyspace, newKeyspace *Keyspace) (string, error) {

	templateStr := "ALTER KEYSPACE {{ .Name }} WITH "
	var templateClauses []string

	if oldKeyspace.Replication != newKeyspace.Replication {
		templateClauses = append(templateClauses, "REPLICATION = {{ .Replication }}")
	}

	if oldKeyspace.DurableWrites != newKeyspace.DurableWrites {
		templateClauses = append(templateClauses, "DURABLE_WRITES = {{ .DurableWrites }}")
	}

	templateStr += strings.Join(templateClauses, " AND ") + ";"

	t, err := template.New("cassandra_keyspace_update").Parse(templateStr)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, newKeyspace)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}