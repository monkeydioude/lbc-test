package entity

type Params struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Inter int    `json:"-"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}
