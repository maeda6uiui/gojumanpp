package gojumanpp

type Morpheme struct{
	Midasi string
	Yomi string
	Genkei string
	Hinsi string
	Hinsi_id int
	Bunrui string
	Bunrui_id int
	Katuyou1 string
	Katuyou1_id int
	Katuyou2 string
	Katuyou2_id int
	Imis string
	Fstring string
	Repname string
}

func NewMorpheme() *Morpheme{
	m:=new(Morpheme)
	return m
}

