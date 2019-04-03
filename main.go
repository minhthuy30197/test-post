package main

import (
	"github.com/kataras/iris"
)

type Post struct {
	Title   string
	Author  string
	Content string
}

type Resource struct {
	Title   string
	Author  string
	Content string
	Links   []string
}

type Project struct {
	Title   string
	Author  string
	Link    string
	Course  string
}

type Experience struct {
	Title   string
	Author  string
	Course  string
}

func main() {
	app := iris.Default()

	app.RegisterView(iris.HTML("./templates", ".html").Layout("layout.html"))

	app.Get("/bai-viet", func(ctx iris.Context) {
		var postType, viewTemplate string
		postType = ctx.URLParam("type")

		// TODO: Lấy dữ liệu từ database
		switch postType {
		case "do-an":
			viewTemplate = "project.html"
			// TODO: parse json và nhét dữ liệu theo struct Project
			// Fake data 
			var post = []Project {
				{
					Title: "Website bất động sản",
					Author: "Đoàn Thanh Hương",
					Link: "https://google.com.vn",
					Course: "Lập trình web frontend cho người mới bắt đầu",
				},
				{
					Title: "Website bán đồ nội thất",
					Author: "Đỗ Trung Hiếu",
					Link: "https://google.com.vn",
					Course: "Lập trình web frontend cho người mới bắt đầu",
				},
				{
					Title: "Ứng dụng nhắc uống nước mỗi ngày",
					Author: "Phan Thị Hạnh",
					Link: "https://google.com.vn",
					Course: "Lập trình IOS Swift",
				},
			} 
			ctx.ViewData("Posts", post)
		case "tai-lieu":
			viewTemplate = "resource.html"
			// TODO: parse json và nhét dữ liệu theo struct Resource 
			// Fake data 
			var post = []Resource {
				{
					Title: "10 mẫu template web bán hàng",
					Author: "Robin Huy",
					Content: "Like share để nhận ngay 10 mẫu template mới nhất",
					Links: []string{"https://google.com.vn", "https://google.com.vn"},
				},
				{
					Title: "Source code phần mềm quản lý thử viện",
					Author: "Nguyễn Văn A",
					Content: "Source code phần mềm quản lý thư viện viết bằng C#",
					Links: []string{"https://google.com.vn"},
				},
			}
			ctx.ViewData("Posts", post)
		case "trai-nghiem":
			viewTemplate = "experience.html"
			// TODO: parse json và nhét dữ liệu theo struct Experience
			// Fake data
			var post = []Experience {
				{
					Title: "Câu chuyện về bước ngoặt từ bỏ 13 năm học và làm ngành Điện để theo đuổi đam mê Lập Trình",
					Author: "Nguyễn Minh Trung",
					Course: "Lập trình web frontend",
				},
				{
					Title: "Quyết định khó khăn ở tuổi 30",
					Author: "Phạm Anh Tuấn",
					Course: "Lập trình PHP",				
				},
				{
					Title: "Techmaster là nhà",
					Author: "Nguyễn Thành Đạt",
					Course: "Lập trình web frontend",				
				},
			}
			ctx.ViewData("Posts", post)
		default:
			viewTemplate = "post.html"
			// Fake data
			var post = []Post {
				{
					Title: "Đây là bài viết 1",
					Author: "Nguyễn Văn A",
					Content: "Đội là tóm tắt của bài viết 1...",					
				},
				{
					Title: "Đây là bài viết 2",
					Author: "Nguyễn Thị B",
					Content: "Đội là tóm tắt của bài viết 2...",					
				},
				{
					Title: "Đây là bài viết 3",
					Author: "Phạm Văn C",
					Content: "Đội là tóm tắt của bài viết 3...",					
				},
				{
					Title: "Đây là bài viết 4",
					Author: "Mai D",
					Content: "Đội là tóm tắt của bài viết 4...",					
				},
			}
			ctx.ViewData("Posts", post)
		}
		ctx.View(viewTemplate)
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":8080"))
}
