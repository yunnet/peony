// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/yunnet/peony/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/alarm_dtu_setting",
			beego.NSInclude(
				&controllers.AlarmDtuSettingController{},
			),
		),

		beego.NSNamespace("/alarm_types",
			beego.NSInclude(
				&controllers.AlarmTypesController{},
			),
		),

		beego.NSNamespace("/equipment_customer",
			beego.NSInclude(
				&controllers.EquipmentCustomerController{},
			),
		),

		beego.NSNamespace("/equipment_dtu_config",
			beego.NSInclude(
				&controllers.EquipmentDtuConfigController{},
			),
		),

		beego.NSNamespace("/equipment_gateway",
			beego.NSInclude(
				&controllers.EquipmentGatewayController{},
			),
		),

		beego.NSNamespace("/equipment_meter_addr_config",
			beego.NSInclude(
				&controllers.EquipmentMeterAddrConfigController{},
			),
		),

		beego.NSNamespace("/equipment_meter_config",
			beego.NSInclude(
				&controllers.EquipmentMeterConfigController{},
			),
		),

		beego.NSNamespace("/equipment_meter_config_field",
			beego.NSInclude(
				&controllers.EquipmentMeterConfigFieldController{},
			),
		),

		beego.NSNamespace("/equipment_meter_rom_config",
			beego.NSInclude(
				&controllers.EquipmentMeterRomConfigController{},
			),
		),

		beego.NSNamespace("/equipment_meter_type",
			beego.NSInclude(
				&controllers.EquipmentMeterTypeController{},
			),
		),

		beego.NSNamespace("/equipment_room",
			beego.NSInclude(
				&controllers.EquipmentRoomController{},
			),
		),

		beego.NSNamespace("/equipment_simcards",
			beego.NSInclude(
				&controllers.EquipmentSimcardsController{},
			),
		),

		beego.NSNamespace("/equipment_table_config",
			beego.NSInclude(
				&controllers.EquipmentTableConfigController{},
			),
		),

		beego.NSNamespace("/equipment_vendor",
			beego.NSInclude(
				&controllers.EquipmentVendorController{},
			),
		),

		beego.NSNamespace("/sys_backend_conf",
			beego.NSInclude(
				&controllers.SysBackendConfController{},
			),
		),

		beego.NSNamespace("/sys_backend_user",
			beego.NSInclude(
				&controllers.SysBackendUserController{},
			),
		),

		beego.NSNamespace("/sys_backend_user_rms_roles",
			beego.NSInclude(
				&controllers.SysBackendUserRmsRolesController{},
			),
		),

		beego.NSNamespace("/sys_conn_config",
			beego.NSInclude(
				&controllers.SysConnConfigController{},
			),
		),

		beego.NSNamespace("/sys_logintrace",
			beego.NSInclude(
				&controllers.SysLogintraceController{},
			),
		),

		beego.NSNamespace("/sys_resource",
			beego.NSInclude(
				&controllers.SysResourceController{},
			),
		),

		beego.NSNamespace("/sys_role",
			beego.NSInclude(
				&controllers.SysRoleController{},
			),
		),

		beego.NSNamespace("/sys_role_backenduser_rel",
			beego.NSInclude(
				&controllers.SysRoleBackenduserRelController{},
			),
		),

		beego.NSNamespace("/sys_role_resource_rel",
			beego.NSInclude(
				&controllers.SysRoleResourceRelController{},
			),
		),

		beego.NSNamespace("/sys_val",
			beego.NSInclude(
				&controllers.SysValController{},
			),
		),

		beego.NSNamespace("/home",
			beego.NSInclude(
				&controllers.HomeController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
