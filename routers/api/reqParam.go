package api

type ReqParam struct {
	Apikey string `form:"apikey" json:"apikey"`
	Param  string `form:"param" json:"param"`
}
