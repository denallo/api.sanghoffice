package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["api.sanghoffice/controllers:ItemController"] = append(beego.GlobalControllerRouter["api.sanghoffice/controllers:ItemController"],
        beego.ControllerComments{
            Method: "GetBrief",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api.sanghoffice/controllers:KutiController"] = append(beego.GlobalControllerRouter["api.sanghoffice/controllers:KutiController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api.sanghoffice/controllers:KutiController"] = append(beego.GlobalControllerRouter["api.sanghoffice/controllers:KutiController"],
        beego.ControllerComments{
            Method: "UpdateBrokenStatus",
            Router: `/status`,
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api.sanghoffice/controllers:ResiStatusCtrl"] = append(beego.GlobalControllerRouter["api.sanghoffice/controllers:ResiStatusCtrl"],
        beego.ControllerComments{
            Method: "AddResiStatus",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api.sanghoffice/controllers:ResidentCtrl"] = append(beego.GlobalControllerRouter["api.sanghoffice/controllers:ResidentCtrl"],
        beego.ControllerComments{
            Method: "GetResidents",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api.sanghoffice/controllers:ResidentCtrl"] = append(beego.GlobalControllerRouter["api.sanghoffice/controllers:ResidentCtrl"],
        beego.ControllerComments{
            Method: "GetResidentInfo",
            Router: `/info`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api.sanghoffice/controllers:ResidentCtrl"] = append(beego.GlobalControllerRouter["api.sanghoffice/controllers:ResidentCtrl"],
        beego.ControllerComments{
            Method: "UpdateResidentInfo",
            Router: `/info`,
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
