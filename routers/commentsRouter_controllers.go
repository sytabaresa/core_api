package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AmparosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AmparosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AmparosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AmparosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AmparosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AmparosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AmparosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AmparosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AmparosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AmparosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ArlController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ArlController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ArlController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ArlController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ArlController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ArlController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ArlController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ArlController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ArlController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ArlController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CajaCompensacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CajaCompensacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CajaCompensacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CajaCompensacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CajaCompensacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ContactoSucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ContactoSucursalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ContactoSucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ContactoSucursalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ContactoSucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ContactoSucursalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ContactoSucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ContactoSucursalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ContactoSucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ContactoSucursalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EntidadAseguradoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EntidadAseguradoraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EntidadAseguradoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EntidadAseguradoraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EntidadAseguradoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EntidadAseguradoraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EntidadAseguradoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EntidadAseguradoraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EntidadAseguradoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EntidadAseguradoraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EpsController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EpsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EpsController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EpsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EpsController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EpsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EpsController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EpsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EpsController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EpsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EstadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EstadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EstadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EstadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:EstadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:FondoPensionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:FondoPensionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:FondoPensionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:FondoPensionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:FondoPensionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:InfoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:InfoEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:InfoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:InfoEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:InfoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:InfoEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:InfoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:InfoEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:InfoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:InfoEntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PaisController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PaisController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ParametroEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ParametroEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ParametroEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ParametroEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ParametroEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ParametroEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ParametroEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ParametroEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ParametroEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ParametroEntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupEspecialidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupEspecialidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupEspecialidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupEspecialidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupEspecialidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupEspecialidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupEspecialidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupEspecialidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupEspecialidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupEspecialidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupGrupoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupGrupoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupGrupoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupGrupoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupGrupoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupTipoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupTipoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupTipoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupTipoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RupTipoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoEntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
