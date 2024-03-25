package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lucky-xin/xyz-common-go/r"
	"io"
	"net/http"
	"strconv"
)

func SendMsg(endpoint, chatId, content string) (re r.Resp[any]) {
	params := fmt.Sprintf(`{
					"chatId": %s,
					"msg": {
						"msgtype": "text",
						"text": {
							"content": %s,
							"atAll": true
						}
					}
				}`, strconv.Quote(chatId), strconv.Quote(content))
	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer([]byte(params)))
	defer func(b io.ReadCloser) {
		if er := b.Close(); er != nil {
			panic(er)
		}
	}(response.Body)

	if err != nil {
		re = r.Failed(err.Error())
		return
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		re = r.Failed(err.Error())
		return
	}
	err = json.Unmarshal(b, &re)
	if err != nil {
		re = r.Failed(err.Error())
	}
	return
}
