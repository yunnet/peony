package models

type VSysBackendConf struct {
	Id_RENAME  int    `orm:"column(id)" description:"序号"`
	Appname    string `orm:"column(appname);size(50)"`
	Appversion string `orm:"column(appversion);size(50)"`
	Deploy     string `orm:"column(deploy);size(50)"`
	Configtext string `orm:"column(configtext);size(5000);null"`
}
