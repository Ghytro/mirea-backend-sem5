package entity

type Form struct {
	tableName struct{} `pg:"forms"`

	Name    string `pg:"name" json:"name"`
	Email   string `pg:"email" json:"email"`
	Message string `pg:"message" json:"message"`
}

type Review struct {
	tableName struct{} `pg:"reviews"`

	Name    string  `pg:"name" json:"name"`
	Rating  int     `pg:"rating" json:"rating"`
	Message *string `pg:"message" json:"message,omitempty"`
}
