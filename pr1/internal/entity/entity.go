package entity

type GenericDBResponse struct {
	tableName struct{} `pg:"some_table"`

	Id      int    `pg:"id,pk"`
	Column1 string `pg:"column1"`
	Column2 int    `pg:"column2"`
}
