package memdb

type MemDb struct {
	tables map[string]*[]interface{}
}

func New() *MemDb {
	return &MemDb{
		tables: map[string]*[]interface{}{},
	}
}

func (d *MemDb) GetTable(tableName string) (*[]interface{}, error) {
	table, ok := d.tables[tableName]
	if !ok {
		table = &[]interface{}{}
		d.tables[tableName] = table
	}
	return table, nil
}
