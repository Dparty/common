package data

type Pair[L, R any] struct {
	L L `json:"left"`
	R R `json:"right"`
}
