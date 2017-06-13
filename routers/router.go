// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/core_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/amparos",
			beego.NSInclude(
				&controllers.AmparosController{},
			),
		),

		beego.NSNamespace("/arl",
			beego.NSInclude(
				&controllers.ArlController{},
			),
		),

		beego.NSNamespace("/caja_compensacion",
			beego.NSInclude(
				&controllers.CajaCompensacionController{},
			),
		),

		beego.NSNamespace("/info_entidad",
			beego.NSInclude(
				&controllers.InfoEntidadController{},
			),
		),

		beego.NSNamespace("/departamento",
			beego.NSInclude(
				&controllers.DepartamentoController{},
			),
		),

		beego.NSNamespace("/contacto_sucursal",
			beego.NSInclude(
				&controllers.ContactoSucursalController{},
			),
		),

		beego.NSNamespace("/ciiu_tipo",
			beego.NSInclude(
				&controllers.CiiuTipoController{},
			),
		),

		beego.NSNamespace("/eps",
			beego.NSInclude(
				&controllers.EpsController{},
			),
		),

		beego.NSNamespace("/fondo_pension",
			beego.NSInclude(
				&controllers.FondoPensionController{},
			),
		),

		beego.NSNamespace("/jefe_dependencia",
			beego.NSInclude(
				&controllers.JefeDependenciaController{},
			),
		),

		beego.NSNamespace("/ordenador_gasto",
			beego.NSInclude(
				&controllers.OrdenadorGastoController{},
			),
		),

		beego.NSNamespace("/ciiu_clase",
			beego.NSInclude(
				&controllers.CiiuClaseController{},
			),
		),

		beego.NSNamespace("/ciudad",
			beego.NSInclude(
				&controllers.CiudadController{},
			),
		),

		beego.NSNamespace("/ciiu_division",
			beego.NSInclude(
				&controllers.CiiuDivisionController{},
			),
		),

		beego.NSNamespace("/parametro_entidad",
			beego.NSInclude(
				&controllers.ParametroEntidadController{},
			),
		),

		beego.NSNamespace("/pais",
			beego.NSInclude(
				&controllers.PaisController{},
			),
		),

		beego.NSNamespace("/rubros_dependencia",
			beego.NSInclude(
				&controllers.RubrosDependenciaController{},
			),
		),

		beego.NSNamespace("/rubros_ordenador",
			beego.NSInclude(
				&controllers.RubrosOrdenadorController{},
			),
		),

		beego.NSNamespace("/salario_minimo",
			beego.NSInclude(
				&controllers.SalarioMinimoController{},
			),
		),

		beego.NSNamespace("/rup_tipo",
			beego.NSInclude(
				&controllers.RupTipoController{},
			),
		),

		beego.NSNamespace("/tipo_documento",
			beego.NSInclude(
				&controllers.TipoDocumentoController{},
			),
		),

		beego.NSNamespace("/tipo_contacto",
			beego.NSInclude(
				&controllers.TipoContactoController{},
			),
		),

		beego.NSNamespace("/snies_area",
			beego.NSInclude(
				&controllers.SniesAreaController{},
			),
		),

		beego.NSNamespace("/banco",
			beego.NSInclude(
				&controllers.BancoController{},
			),
		),

		beego.NSNamespace("/tipo_entidad",
			beego.NSInclude(
				&controllers.TipoEntidadController{},
			),
		),

		beego.NSNamespace("/estado",
			beego.NSInclude(
				&controllers.EstadoController{},
			),
		),

		beego.NSNamespace("/entidad_aseguradora",
			beego.NSInclude(
				&controllers.EntidadAseguradoraController{},
			),
		),

		beego.NSNamespace("/ciiu_subclase",
			beego.NSInclude(
				&controllers.CiiuSubclaseController{},
			),
		),

		beego.NSNamespace("/rup_grupo",
			beego.NSInclude(
				&controllers.RupGrupoController{},
			),
		),

		beego.NSNamespace("/snies_nucleo_basico",
			beego.NSInclude(
				&controllers.SniesNucleoBasicoController{},
			),
		),

		beego.NSNamespace("/rup_especialidad",
			beego.NSInclude(
				&controllers.RupEspecialidadController{},
			),
		),

		beego.NSNamespace("/sucursal",
			beego.NSInclude(
				&controllers.SucursalController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
