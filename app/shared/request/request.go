package request

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
)

var (
	decoder     = schema.NewDecoder()
	jsonDecoder = schema.NewDecoder()
)

func init() {
	jsonDecoder.SetAliasTag("json")
}

func Bind(c *gin.Context, req interface{}) error {
	if c.Request.Method == http.MethodGet {
		query := c.Request.URL.Query()
		return decoder.Decode(req, query)
	}
	if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json") {
		buf, err := io.ReadAll(c.Request.Body)
		if err != nil {
			return err
		}
		if msg, ok := req.(proto.Message); ok {
			return protojson.Unmarshal(buf, msg)
		}
		return json.Unmarshal(buf, req)
	}
	return c.ShouldBind(req)
}
