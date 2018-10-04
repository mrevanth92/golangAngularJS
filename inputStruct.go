package main
type Inputs struct  {
	Data []*Data `json:"inputs"`
}

type  Data struct{
	 Image *Image `json:"data"`
}

type  Image struct{
	Url *Url `json:"image"`
}

type Url struct {
	Value string `json:"url"`
}