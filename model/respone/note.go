package respone

type RespNote struct {
	Message string `json:"message"` // 返回信息
	Id      int64  `json:"id"`      // 返回文章id
}
