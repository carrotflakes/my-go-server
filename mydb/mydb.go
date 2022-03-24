package mydb

type MyDB struct {
	tables map[string]*[]interface{}
}

func New() *MyDB {
	return &MyDB{
		tables: map[string]*[]interface{}{},
	}
}

func (d *MyDB) GetTable(tableName string) (*[]interface{}, error) {
	table, ok := d.tables[tableName]
	if !ok {
		table = &[]interface{}{}
		d.tables[tableName] = table
	}
	return table, nil
}
