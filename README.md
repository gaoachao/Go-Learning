# Go Learning

## 聊天室



## gin

### 匿名结构体HTTP处理函数中JSON序列化和反序列化

我们在处理http请求时，通常会和JSON数据打交道。

比如post请求的content-type使用application/json时，服务器接收过来的json数据是key:value格式，不同key的value的类型可以不一样，可能是数字、字符串、数组等，因此会遇到使用`json.Unmarshal`和`map[string]interface{}`来接收JSON反序列化后的数据。

但是使用map[string]interface{}有几个问题：

- 没有类型检查：比如json的某个value本来预期是string类型，但是请求传过来的是bool类型，使用json.Unmarshal解析到map[string]interface{}是不会报错的，因为空interface可以接受任何类型数据。
- map是模糊的：Unmarshal后得到了map，我们还得判断这个key在map里是否存在。否则拿不存在的key的value，得到的可能是给nil值，如果不做检查，直接对nil指针做*操作，会引发panic。
- 代码比较冗长：得先判断key是否存在，如果存在，要显示转换成对应的数据类型，并且还得判断转换是否成功。代码会比较冗长。

这个时候我们就可以使用匿名结构体来接收反序列化后的数据，代码会更简洁。

```go
func main() {
	r := gin.Default()
	r.POST("/ping", func(c *gin.Context) {
		body:= struct {
			Message string `json:"message"`
		}{}
		err := c.BindJSON(&body)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(string(body.Message))
		c.JSON(http.StatusOK, gin.H{
			"message": body.Message,
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
```

