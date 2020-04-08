package controllers

import (
	"encoding/json"
	"github.com/yunnet/peony/models"
	"strconv"
	"strings"

	"github.com/yunnet/peony/enums"
)

// EquipmentSimcardsController operations for EquipmentSimcards
type EquipmentSimcardsController struct {
	BaseController
}

func (this *EquipmentSimcardsController) Prepare() {
	this.BaseController.Prepare()
	this.checkAuthor()
}

// URLMapping ...
func (c *EquipmentSimcardsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create EquipmentSimcards
// @Param	body		body 	models.EquipmentSimcards	true		"body for EquipmentSimcards content"
// @Success 201 {int} models.EquipmentSimcards
// @Failure 403 body is empty
// @router / [post]
func (c *EquipmentSimcardsController) Post() {
	var v models.EquipmentSimcards
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddEquipmentSimcards(&v); err == nil {
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
// @Description get EquipmentSimcards by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.EquipmentSimcards
// @Failure 403 :id is empty
// @router /:id [get]
func (c *EquipmentSimcardsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetEquipmentSimcardsById(id)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", err.Error())
	} else {
		c.jsonResult(enums.JRCodeSucc, "", v)
	}
}

// GetAll ...
// @Title Get All
// @Description get EquipmentSimcards
// @Param   token   header  string  true    "token"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Param   paging  query   bool    false   "Paging return"
// @Success 200 {object} models.EquipmentSimcards
// @Failure 403
// @router / [get]
func (c *EquipmentSimcardsController) GetAll() {
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

	l, total, err := models.GetAllEquipmentSimcards(query, fields, sortby, order, offset, limit, paging)
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
// @Description update the EquipmentSimcards
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.EquipmentSimcards	true		"body for EquipmentSimcards content"
// @Success 200 {object} models.EquipmentSimcards
// @Failure 403 :id is not int
// @router /:id [put]
func (c *EquipmentSimcardsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.EquipmentSimcards{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateEquipmentSimcardsById(&v); err == nil {
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
// @Description delete the EquipmentSimcards
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *EquipmentSimcardsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteEquipmentSimcards(id); err == nil {
		c.jsonResult(enums.JRCodeFailed, "更新数据失败", err.Error())
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除数据失败", err.Error())
	}
}
