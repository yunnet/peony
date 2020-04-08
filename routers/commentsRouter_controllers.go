package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmDtuSettingController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmDtuSettingController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmDtuSettingController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmDtuSettingController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmDtuSettingController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmDtuSettingController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmDtuSettingController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmDtuSettingController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmDtuSettingController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmDtuSettingController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmTypesController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmTypesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmTypesController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmTypesController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmTypesController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmTypesController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmTypesController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmTypesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmTypesController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:AlarmTypesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentCustomerController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentCustomerController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentCustomerController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentCustomerController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentCustomerController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentCustomerController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentCustomerController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentCustomerController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentCustomerController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentCustomerController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentDtuConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentDtuConfigController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentDtuConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentDtuConfigController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentDtuConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentDtuConfigController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentDtuConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentDtuConfigController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentDtuConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentDtuConfigController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentGatewayController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentGatewayController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentGatewayController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentGatewayController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentGatewayController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentGatewayController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentGatewayController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentGatewayController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentGatewayController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentGatewayController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterAddrConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterAddrConfigController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterAddrConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterAddrConfigController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterAddrConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterAddrConfigController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterAddrConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterAddrConfigController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterAddrConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterAddrConfigController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigFieldController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigFieldController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigFieldController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigFieldController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigFieldController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigFieldController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigFieldController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigFieldController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigFieldController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterConfigFieldController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterRomConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterRomConfigController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterRomConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterRomConfigController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterRomConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterRomConfigController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterRomConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterRomConfigController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterRomConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterRomConfigController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterTypeController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterTypeController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterTypeController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterTypeController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterTypeController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterTypeController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterTypeController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterTypeController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterTypeController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentMeterTypeController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentRoomController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentRoomController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentRoomController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentRoomController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentRoomController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentRoomController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentRoomController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentRoomController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentRoomController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentRoomController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentSimcardsController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentSimcardsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentSimcardsController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentSimcardsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentSimcardsController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentSimcardsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentSimcardsController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentSimcardsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentSimcardsController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentSimcardsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentTableConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentTableConfigController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentTableConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentTableConfigController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentTableConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentTableConfigController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentTableConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentTableConfigController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentTableConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentTableConfigController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentVendorController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentVendorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentVendorController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentVendorController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentVendorController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentVendorController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentVendorController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentVendorController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentVendorController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:EquipmentVendorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:HomeController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:HomeController"],
		beego.ControllerComments{
			Method:           "GetConfig",
			Router:           `/config`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:HomeController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:HomeController"],
		beego.ControllerComments{
			Method:           "DoLogin",
			Router:           `/dologin`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendConfController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendConfController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendConfController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendConfController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendConfController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendConfController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendConfController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendConfController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendConfController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendConfController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserRmsRolesController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserRmsRolesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserRmsRolesController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserRmsRolesController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserRmsRolesController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserRmsRolesController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserRmsRolesController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserRmsRolesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserRmsRolesController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysBackendUserRmsRolesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysConnConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysConnConfigController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysConnConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysConnConfigController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysConnConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysConnConfigController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysConnConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysConnConfigController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysConnConfigController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysConnConfigController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysLogintraceController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysLogintraceController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysLogintraceController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysLogintraceController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysLogintraceController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysLogintraceController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysLogintraceController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysLogintraceController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysLogintraceController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysLogintraceController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysResourceController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysResourceController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysResourceController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysResourceController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysResourceController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysResourceController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysResourceController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysResourceController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysResourceController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysResourceController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleBackenduserRelController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleBackenduserRelController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleBackenduserRelController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleBackenduserRelController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleBackenduserRelController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleBackenduserRelController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleBackenduserRelController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleBackenduserRelController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleBackenduserRelController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleBackenduserRelController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleResourceRelController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleResourceRelController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleResourceRelController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleResourceRelController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleResourceRelController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleResourceRelController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleResourceRelController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleResourceRelController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleResourceRelController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysRoleResourceRelController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysValController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysValController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysValController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysValController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysValController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysValController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysValController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysValController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysValController"] = append(beego.GlobalControllerRouter["github.com/yunnet/peony/controllers:SysValController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

}
