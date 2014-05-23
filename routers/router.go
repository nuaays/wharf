package routers

import (
	"github.com/astaxie/beego"
	"github.com/dockboard/docker-registry/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//Ping Router
	beego.Router("/_ping", &controllers.PingController{}, "get:GetPing")
	beego.Router("/_ping/", &controllers.PingController{}, "get:GetPing")
	beego.Router("/v1/_ping", &controllers.PingController{}, "get:GetPing")
	beego.Router("/v1/_ping/", &controllers.PingController{}, "get:GetPing")

	//User Router
	beego.Router("/v1/users/", &controllers.UsersController{}, "get:GetUsers")
	beego.Router("/v1/users/", &controllers.UsersController{}, "post:PostUsers")
	beego.Router("/v1/users", &controllers.UsersController{}, "get:GetUsers")
	beego.Router("/v1/users", &controllers.UsersController{}, "post:PostUsers")

	//Push Router Begin
	//1. 写入要上传的 Repository 的 JSON 信息，此 JSON 信息是一个包含所有 Image ID 的 JSON 字符串。
	beego.Router("/v1/repositories/:namespace/:repo_name/", &controllers.RepositoryController{}, "put:PutRepository")
	beego.Router("/v1/repositories/:namespace/:repo_name", &controllers.RepositoryController{}, "put:PutRepository")
	//2. 根据 Repository 的所有 Image 开始循环处理
	//2.1 根据 Image 的 ID 获取 Image 的 JSON 数据，如果返回 404 就要上传此 Image 的 JSON 数据和 Layer 文件
	beego.Router("/v1/images/:image_id/json/", &controllers.ImageController{}, "get:GetImageJson")
	beego.Router("/v1/images/:image_id/json", &controllers.ImageController{}, "get:GetImageJson")
	//2.2 上传 Image 的 JSON 文件
	beego.Router("/v1/images/:image_id/json/", &controllers.ImageController{}, "put:PutImageJson")
	beego.Router("/v1/images/:image_id/json", &controllers.ImageController{}, "put:PutImageJson")
	//2.3 上传 Image 的 Layer 文件
	beego.Router("/v1/images/:image_id/layer/", &controllers.ImageController{}, "put:PutImageIdLayer")
	beego.Router("/v1/images/:image_id/layer", &controllers.ImageController{}, "put:PutImageIdLayer")
	//2.4 上传 Image 的 Checksum 文件
	beego.Router("/v1/images/:image_id/checksum/", &controllers.ImageController{}, "put:PutChecksum")
	beego.Router("/v1/images/:image_id/checksum", &controllers.ImageController{}, "put:PutChecksum")
	//End Image 循环
	//3. 根据 Repository 的所有 Tag 信息循环写入所有的 Tag
	beego.Router("/v1/repositories/:namespace/:repository/tags/:tag/", &controllers.RepositoryController{}, "put:PutTag")
	beego.Router("/v1/repositories/:namespace/:repository/tags/:tag", &controllers.RepositoryController{}, "put:PutTag")
	//4. 最后执行，并没有上传任何有效数据
	beego.Router("/v1/repositories/:namespace/:repository/images/", &controllers.RepositoryController{}, "put:PutRepositoryImages")
	beego.Router("/v1/repositories/:namespace/:repository/images", &controllers.RepositoryController{}, "put:PutRepositoryImages")
	//Push Router End

	beego.Router("/_status", &controllers.StatusController{})
	beego.Router("/v1/_status", &controllers.StatusController{})
}
