package core

import (
	`fmt`
	libUrl `net/url`
	`strings`

	`github.com/storezhang/gox`
)

type (
	// OrgConfig 机构配置
	OrgConfig struct {
		// Url 机构地址
		Url string `json:"url" mapstructure:"url" validate:"required"`
		// Name 机构名称
		Name string `json:"name" mapstructure:"name" validate:"required"`
	}
)

func (oc *OrgConfig) StructToMap() (model map[string]interface{}, err error) {
	return gox.StructToMap(oc)
}

func (oc *OrgConfig) MapToStruct(model map[string]interface{}) (err error) {
	return gox.MapToStruct(model, oc)
}

func (oc *OrgConfig) Exist() bool {
	return "" != oc.Url && "" != oc.Name
}

func (oc *OrgConfig) Domain() (domain string, err error) {
	var url *libUrl.URL

	if url, err = libUrl.Parse(oc.Url); nil != err {
		return
	}

	domain = url.Hostname()

	return
}

func (oc *OrgConfig) PackageName() (name string, err error) {
	var domain string
	if domain, err = oc.Domain(); nil != err {
		return
	}

	// 解析包名
	name = gox.ParsePackageName(domain)

	return
}

func (oc *OrgConfig) ConfigUpdateUrl(name ConfigName) (path string, err error) {
	return oc.ApiUrl(OrgApiConfigUpdateUrl, map[string]string{"name": string(name)}, ApiVersionDefault)
}

func (oc *OrgConfig) PackageNotifyUrl(client BaseClient) (path string, err error) {
	return oc.ApiUrl(OrgApiClientPackageNotifyUrl, map[string]string{"id": client.IdString()}, ApiVersionDefault)
}

func (oc *OrgConfig) ApiUrl(
	api ApiPath,
	pathParams map[string]string,
	version ApiVersion,
) (path string, err error) {
	return oc.GetUrl(api, pathParams, UrlApiPrefix, version)
}

func (oc *OrgConfig) GetUrl(
	path ApiPath,
	pathParams map[string]string,
	prefix string,
	version ApiVersion,
) (url string, err error) {
	var sb strings.Builder

	if _, err = sb.WriteString(oc.Url); nil != err {
		return
	}

	if "" != prefix {
		if _, err = sb.WriteString(fmt.Sprintf("/%s", prefix)); nil != err {
			return
		}
	}

	if ApiVersionDefault != version {
		if _, err = sb.WriteString(fmt.Sprintf("/%s", version)); nil != err {
			return
		}
	}
	sb.WriteString(fmt.Sprintf("/%s", path))

	url = sb.String()
	// 处理路径参数
	if 0 < len(pathParams) {
		for param, value := range pathParams {
			url = strings.Replace(url, fmt.Sprintf("{%s}", param), libUrl.PathEscape(value), -1)
		}
	}

	return
}
