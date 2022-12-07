package model

type Data struct {
	ID       int    `db:"id" json:"id"`
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"password"`
}

type Config struct {
	DataSourceName string `json:"dataSourceName"`
	Port           int    `json:"port"`
}
