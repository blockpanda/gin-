package main

import (
	"fmt"
	"gin/routers"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

func intMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("5555555555zjjj")
	c.Next()
	fmt.Println("666666666666zjjj")
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}
func main() {
	r := gin.Default()
	//加载模板
	r.LoadHTMLGlob("templates/**/*")

	//配置静态WEB目录  第一个参数标识路由，第二个参数标识映射的目录
	r.Static("/static", "./static")

	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)

	r.GET("/zjj", intMiddleware,
		func(c *gin.Context) {
			c.String(200, "中间件测试")
			fmt.Println("7777777777777ddd")
		})

	// 前台     GET 请求传值
	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "default/index.html", gin.H{
	// 		"title": "首页",
	// 		"msg":   "我是msg",
	// 		"hobby": []string{"吃饭", "睡觉", "写代码"},
	// 		"newlist": []interface{}{
	// 			&Article{
	// 				Title:   "新闻标题1111",
	// 				Content: "新闻内容1111",
	// 			},
	// 			&Article{
	// 				Title:   "新闻标题2222",
	// 				Content: "新闻内容22222",
	// 			},
	// 		},
	// 	})
	// })
	// r.GET("/news", func(c *gin.Context) {
	// 	news := &Article{
	// 		Title:   "新闻标题",
	// 		Content: "新闻内容",
	// 	}
	// 	c.HTML(http.StatusOK, "default/news.html", gin.H{
	// 		"title": "新闻页面",
	// 		"news":  news,
	// 	})
	// })

	// //后台配置
	// r.GET("/admin", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "admin/index.html", gin.H{
	// 		"title": "后台首页",
	// 	})
	// })
	// r.GET("/admin/news", func(c *gin.Context) {
	// 	news := &Article{
	// 		Title:   "后台新闻标题",
	// 		Content: "后台新闻内容",
	// 	}
	// 	c.HTML(http.StatusOK, "admin/news.html", gin.H{
	// 		"title": "后台新闻页面",
	// 		"news":  news,
	// 	})
	// })

	r.Run()
}
