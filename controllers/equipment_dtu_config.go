package controllers

import (
	"encoding/json"
	"github.com/yunnet/peony/models"
	"strconv"
	"strings"

	"github.com/yunnet/peony/enums"
)

// EquipmentDtuConfigController operations for EquipmentDtuConfig
type EquipmentDtuConfigController struct {
	BaseController
}

func (this *EquipmentDtuConfigController) Prepare() {
	this.BaseController.Prepare()
	this.checkAuthor()
}

// URLMapping ...
func (c *EquipmentDtuConfigController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create EquipmentDtuConfig
// @Param   token   header  string  true    "token"
// @Param	body		body 	models.EquipmentDtuConfig	true		"body for EquipmentDtuConfig content"
// @Success 201 {int} models.EquipmentDtuConfig
// @Failure 403 body is empty
// @router / [post]
func (c *EquipmentDtuConfigController) Post() {
	var v models.EquipmentDtuConfig
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddEquipmentDtuConfig(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get EquipmentDtuConfig by id
// @Param   token   header  string  true    "token"
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.EquipmentDtuConfig
// @Failure 403 :id is empty
// @router /:id [get]
func (c *EquipmentDtuConfigController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetEquipmentDtuConfigById(id)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", err.Error())
	} else {
		c.jsonResult(enums.JRCodeSucc, "", v)
	}
}

// GetAll ...
// @Title Get All
// @Description get EquipmentDtuConfig
// @Param   token   header  string  true    "token"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Param   paging  query   bool    false   "Paging return"
// @Success 200 {object} models.EquipmentDtuConfig
// @Failure 403
// @router / [get]
func (c *EquipmentDtuConfigController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0
	var paging bool = false

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

	l, total, err := models.GetAllEquipmentDtuConfig(query, fields, sortby, order, offset, limit, paging)
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
// @Description update the EquipmentDtuConfig
// @Param   token   header  string  true    "token"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.EquipmentDtuConfig	true		"body for EquipmentDtuConfig content"
// @Success 200 {object} models.EquipmentDtuConfig
// @Failure 403 :id is not int
// @router /:id [put]
func (c *EquipmentDtuConfigController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.EquipmentDtuConfig{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateEquipmentDtuConfigById(&v); err == nil {
			c.jsonResult(enums.JRCodeSucc, "更新数据成功", nil)
		} else {
			c.jsonResult(enums.JRCodeFailed, "更新数据失败", err.Error())
		}
	} else {
		c.jsonResult(enums.JRCodeFailed, "更新数据失败", err.Error())
	}
}

// Delete ...
// @Title Delete
// @Description delete the EquipmentDtuConfig
// @Param   token   header  string  true    "token"
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *EquipmentDtuConfigController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteEquipmentDtuConfig(id); err == nil {
		c.jsonResult(enums.JRCodeSucc, "删除数据成功", nil)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除数据失败", err.Error())
	}
}
