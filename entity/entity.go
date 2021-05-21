package entity

type TestBO struct {
	Uid    int    `json:"uid"`
	Input  string `json:"input"`
	Output string `json:"output"`
}

type TestDO struct {
	IdBigint    int64 `db:"id_bigint" json:"id_bigint"`
	CodeVarchar string `db:"code_varchar" json:"code_varchar"`
	StockText   string `db:"stock_text" json:"stock_text"`
}

type ResultVO struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}
