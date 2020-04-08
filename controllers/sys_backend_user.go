package controllers

import (
	"encoding/json"
	"errors"
	"github.com/yunnet/peony/models"
	"strconv"
	"strings"

	"github.com/yunnet/peony/enums"
)

// SysBackendUserController operations for SysBackendUserGetAllSysBackendUser
type SysBackendUserController struct {
	BaseController
}

func (this *SysBackendUserController) Prepare() {
	this.BaseController.Prepare()
	this.checkAuthor()
}

// URLMapping ...
func (this *SysBackendUserController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("GetOne", this.GetOne)
	this.Mapping("GetAll", this.GetAll)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

// Post ...
// @Title Post
// @Description create SysBackendUser
// @Param	body		body 	models.SysBackendUser	true		"body for SysBackendUser content"
// @Success 201 {int} models.SysBackendUser
// @Failure 403 body is empty
// @router / [post]
func (c *SysBackendUserController) Post() {
	var v models.SysBackendUser
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSysBackendUser(&v); err == nil {
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
// @Description get SysBackendUser by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SysBackendUser
// @Failure 403 :id is empty
// @router /:id [get]
func (this *SysBackendUserController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSysBackendUserById(id)
	if err != nil {
		this.jsonResult(enums.JRCodeFailed, "", err.Error())
	} else {
		this.jsonResult(enums.JRCodeSucc, "", v)
	}
}

// GetAll ...
// @Title Get All
// @Description get SysBackendUser
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SysBackendUser
// @Failure 403
// @router / [get]
func (c *SysBackendUserController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

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
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllSysBackendUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the SysBackendUser
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SysBackendUser	true		"body for SysBackendUser content"
// @Success 200 {object} models.SysBackendUser
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SysBackendUserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SysBackendUser{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSysBackendUserById(&v); err == nil {
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
// @Description delete the SysBackendUser
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SysBackendUserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSysBackendUser(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
