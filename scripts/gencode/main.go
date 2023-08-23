package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"lark/scripts/gencode/config"
	"lark/scripts/gencode/template"
	"lark/scripts/gencode/utils"
	"strings"
	"unicode"
)

func main() {
	var (
		serviceName      = "Payment"
		upperServiceName = toCamel(serviceName)
		lowerServiceName = firstLower(upperServiceName)
		packageName      = camelToSnake(upperServiceName)
	)
	conf := config.GenConfig{
		Path:        "",
		Prefix:      "",
		PackageName: packageName,
		Dict: map[string]interface{}{
			"UpperServiceName": upperServiceName,
			"LowerServiceName": lowerServiceName,
			"PackageName":      packageName,
		},
	}
	// apps
	generateAppsClientCode(conf)
	generateAppsCmdCode(conf)
	generateAppsConfigCode(conf)
	generateAppsDigCode(conf)
	generateAppsGrpcServerCode(conf)
	generateAppsGrpcServiceCode(conf)
	generateAppsServerCode(conf)
	generateAppsServiceCode(conf)
	generateAppsServiceConstCode(conf)

	// configs
	generateConfigsYamlCode(conf)

	// domain
	generateDomainCacheCode(conf)
	generateDomainRepoCode(conf)
	generateDomainPoCode(conf)

	// interfaces
	generateInterfacesDigCode(conf)
	generateInterfacesCtrlCode(conf)
	generateInterfacesDtoCode(conf)
	generateInterfacesRouterCode(conf)
	generateInterfacesServiceCode(conf)

	// pkg
	generatePkgProtoCode(conf)
	generatePkgProtoGoCode(conf)
	generatePkgProtoRespCode(conf)
}

func toCamel(s string) string {
	c := cases.Title(language.English, cases.NoLower)
	return c.String(s)
}

func firstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func firstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func camelToSnake(input string) string {
	var buffer bytes.Buffer
	for i, r := range input {
		if unicode.IsUpper(r) {
			if i > 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// Apps
func generateAppsClientCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/%s/client", conf.PackageName)
	conf.Filename = "client"
	utils.GenCode(template.AppsClientTemplate, &conf)
}

func generateAppsCmdCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/%s/cmd", conf.PackageName)
	conf.Prefix = "main_"
	utils.GenCode(template.AppsCmdTemplate, &conf)
}

func generateAppsConfigCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/%s/internal/config", conf.PackageName)
	code := strings.ReplaceAll(template.AppsConfigCode, "\"yaml:", "`yaml:")
	code = strings.ReplaceAll(code, "\"\"", "\"`")
	conf.Filename = "config"
	utils.GenCode(template.ParseTemplate(code), &conf)
}

func generateAppsDigCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/%s/dig", conf.PackageName)
	conf.Filename = "dig"
	utils.GenCode(template.AppsDigTemplate, &conf)
}

func generateAppsGrpcServerCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/%s/internal/server/%s", conf.PackageName, conf.PackageName)
	conf.Prefix = "svr_"
	utils.GenCode(template.AppsGrpcServerTemplate, &conf)
}

func generateAppsGrpcServiceCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/%s/internal/server/%s", conf.PackageName, conf.PackageName)
	conf.Prefix = "svr_"
	conf.Suffix = "_service"
	utils.GenCode(template.AppsGrpcServiceTemplate, &conf)
}

func generateAppsServerCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/%s/internal/server", conf.PackageName)
	utils.GenCode(template.AppsServerTemplate, &conf)
}

func generateAppsServiceCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/%s/internal/service", conf.PackageName)
	conf.Prefix = "svc_"
	utils.GenCode(template.AppsServiceTemplate, &conf)
}

func generateAppsServiceConstCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/%s/internal/service", conf.PackageName)
	conf.Prefix = "svc_"
	conf.Suffix = "_const"
	conf.Dict["AllUpperServiceName"] = strings.ToUpper(conf.PackageName)
	utils.GenCode(template.AppsServiceConstTemplate, &conf)
}

// Configs
func generateConfigsYamlCode(conf config.GenConfig) {
	conf.Path = "./configs"
	conf.FileType = config.FILE_TYPE_YAML
	utils.GenCode(template.ConfigsYamlTemplate, &conf)
}

func generateDomainCacheCode(conf config.GenConfig) {
	conf.Path = "./domain/cache"
	conf.Prefix = "cache_"
	utils.GenCode(template.DomainCacheTemplate, &conf)
}

func generateDomainPoCode(conf config.GenConfig) {
	conf.Path = "./domain/po"
	conf.Prefix = "po_"
	utils.GenCode(template.DomainPoTemplate, &conf)
}

func generateDomainRepoCode(conf config.GenConfig) {
	conf.Path = "./domain/repo"
	conf.Prefix = "repo_"
	utils.GenCode(template.DomainRepoTemplate, &conf)
}

func generateInterfacesDigCode(conf config.GenConfig) {
	conf.Path = "./apps/interfaces/dig"
	conf.Prefix = "dig_"
	utils.GenCode(template.InterfacesDigTemplate, &conf)
}

func generateInterfacesCtrlCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/interfaces/internal/ctrl/ctrl_%s", conf.PackageName)
	conf.Prefix = "ctrl_"
	utils.GenCode(template.InterfacesCtrlTemplate, &conf)
}

func generateInterfacesDtoCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/interfaces/internal/dto/dto_%s", conf.PackageName)
	conf.Prefix = "dto_"
	utils.GenCode(template.InterfacesDtoTemplate, &conf)
}

func generateInterfacesRouterCode(conf config.GenConfig) {
	conf.Path = "./apps/interfaces/internal/router"
	conf.Prefix = "router_"
	utils.GenCode(template.InterfacesRouterTemplate, &conf)
}

func generateInterfacesServiceCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./apps/interfaces/internal/service/svc_%s", conf.PackageName)
	conf.Prefix = "svc_"
	utils.GenCode(template.InterfacesServiceTemplate, &conf)
}

func generatePkgProtoCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./pkg/proto/pb_%s", conf.PackageName)
	conf.FileType = config.FILE_TYPE_PROTO
	utils.GenCode(template.PkgProtoTemplate, &conf)
}

func generatePkgProtoGoCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./pkg/proto/pb_%s", conf.PackageName)
	conf.Suffix = ".pb"
	utils.GenCode(template.PkgProtoGoTemplate, &conf)
}

func generatePkgProtoRespCode(conf config.GenConfig) {
	conf.Path = fmt.Sprintf("./pkg/proto/pb_%s", conf.PackageName)
	conf.PackageName = "resp"
	utils.GenCode(template.PkgProtoRespTemplate, &conf)
}
