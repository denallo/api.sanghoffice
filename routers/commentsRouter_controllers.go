package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

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

    beego.GlobalControllerRouter["api.sanghoffice/controllers:ResidentController"] = append(beego.GlobalControllerRouter["api.sanghoffice/controllers:ResidentController"],
        beego.ControllerComments{
            Method: "AddResident",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
