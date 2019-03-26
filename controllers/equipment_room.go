package controllers

import (
	"encoding/json"
	"github.com/yunnet/peony/models"
	"strconv"
	"strings"

	"github.com/yunnet/peony/enums"
)

// EquipmentRoomController operations for EquipmentRoom
type EquipmentRoomController struct {
	BaseController
}

func (this *EquipmentRoomController) Prepare() {
	this.BaseController.Prepare()
	this.checkAuthor()
}

// URLMapping ...
func (c *EquipmentRoomController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create EquipmentRoom
// @Param	body		body 	models.EquipmentRoom	true		"body for EquipmentRoom content"
// @Success 201 {int} models.EquipmentRoom
// @Failure 403 body is empty
// @router / [post]
func (c *EquipmentRoomController) Post() {
	var v models.EquipmentRoom
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddEquipmentRoom(&v); err == nil {
			c.jsonResult(enums.JRCodeSucc, "增加成功.", v)
		} else {
			c.jsonResult(enums.JRCodeFailed, "增加错误", err.Error())
		}
	} else {
		c.jsonResult(enums.JRCodeFailed, "增加错误", err.Error())
	}
}

// GetOne ...
// @Title Get One
// @Description get EquipmentRoom by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.EquipmentRoom
// @Failure 403 :id is empty
// @router /:id [get]
func (c *EquipmentRoomController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetEquipmentRoomById(id)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", err.Error())
	} else {
		c.jsonResult(enums.JRCodeSucc, "", v)
	}
}

// GetAll ...
// @Title Get All
// @Description get EquipmentRoom
// @Param   token   header  string  true    "token"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Param   paging  query   bool    false   "Paging return"
// @Success 200 {object} models.EquipmentRoom
// @Failure 403
// @router / [get]
func (c *EquipmentRoomController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0
	var paging = false

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	//paging
	if v, err := c.GetBool("paging", false); err == nil {
		paging = v
	}

	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.jsonResult(enums.JRCodeFailed, "Error: invalid query key/value pair", nil)
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, total, err := models.GetAllEquipmentRoom(query, fields, sortby, order, offset, limit, paging)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", err.Error())
	} else {
		if paging {
			result := &models.JsonPaging{*(&models.Paging{limit, offset, total}), l}
			c.jsonResult(enums.JRCodeSucc, "", result)
		} else {
			c.jsonResult(enums.JRCodeSucc, "", l)
		}
	}
}

// Put ...
// @Title Put
// @Description update the EquipmentRoom
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.EquipmentRoom	true		"body for EquipmentRoom content"
// @Success 200 {object} models.EquipmentRoom
// @Failure 403 :id is not int
// @router /:id [put]
func (c *EquipmentRoomController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.EquipmentRoom{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateEquipmentRoomById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the EquipmentRoom
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *EquipmentRoomController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteEquipmentRoom(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
