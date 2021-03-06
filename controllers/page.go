package controllers

import (
	"github.com/astaxie/beego"
	"showdoc/consts"
	"showdoc/models"
	"strings"
	"showdoc/helper"
)

// 分类模块
type PageController struct {
	beego.Controller
}

// @Title  分类列表
// @Description 分类列表
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.page
// @router /uploadImg [post]
func (this *PageController) UploadImg() {

	var ext string
	var filename string //文件名称
	var fileUrl string //文件 url
	savePath := beego.AppConfig.String("upload_path")
	f, h, _ := this.GetFile("editormd-image-file")                  //获取上传的文件

	ext = h.Filename[strings.LastIndex( h.Filename, "."):]
	filename = helper.UniqueId() + ext

	path := "."+savePath + "/" + filename  //文件目录
	fileUrl = this.Ctx.Request.URL.Host + savePath + "/" + filename

	f.Close()                                          //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	this.SaveToFile("editormd-image-file", path)                    //存文件
	println(path)
	this.Data["json"] = map[string]interface{}{"url":fileUrl,"success":1}
	this.ServeJSON()

}

// @Title  分类列表
// @Description 分类列表
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.catalogs
// @router /save [post]
func (this *PageController) Save() {
	/**
	page_id=0&
	item_id=2&
	s_number=&
	page_title=tes&
	page_content=%E6%AC%A2%E8%BF%8E%E4%BD&
	cat_id=0

	 */
	uid := this.GetSession(consts.SESSION_UID)
	if uid == nil {
		this.Abort("403")
	} else {

		id,_ :=this.GetInt("page_id")
		item_id,_ :=this.GetInt("item_id")
		s_number,_ := this.GetInt("s_number")
		page_title := this.GetString("page_title")
		page_content := this.GetString("page_content")
		cat_id,_ := this.GetInt("cat_id")


		var page models.Page
		page.Id = id
		page.ItemId = item_id
		page.SNumber = s_number
		page.PageTitle = page_title
		page.PageContent = page_content
		page.CatId = cat_id

		page.SavePage()

		json := consts.Json{}
		json.SetData(page)
		this.Data["json"] = json.VendorOk()
		this.ServeJSON()
	}


}

// @Title  分类列表
// @Description 分类列表
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.catalogs
// @router /info [post]
func (this *PageController) Info() {
	json := consts.Json{}
	id,_ := this.GetInt("page_id")
	//uid := this.GetSession(consts.SESSION_UID)

	_,item := models.GetOnePage(id)

	json.SetData(item)
	this.Data["json"] = json.VendorOk()
	this.ServeJSON()

}

// @Title  删除item
// @Description 分类列表
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.catalogs
// @router /delete [post]
func (this *PageController) Delete() {
	json := consts.Json{}
	id,_ := this.GetInt("page_id")
	uid := this.GetSession(consts.SESSION_UID)
	if uid == nil {
		this.Abort("403")
	} else {
		ret,page := models.GetOnePage(id)
		if ret {
			err := page.Delete()
			if err == nil {
				json.Set(0,"删除成功")
			} else {
				json.Set(404,"删除失败")
			}
		} else {
			json.Set(404,"数据不存在")
		}
		this.Data["json"] = json.VendorOk()

		this.ServeJSON()
	}
}
