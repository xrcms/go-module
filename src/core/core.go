package core

import "github.com/xrcms/globals"

/**
  Если какие либо из интерфейсов, вам точно не потребуются можете их закомментировать.
  Сэкономите, пару микросекунд, спасёте дерево! ;)
*/

type Core struct {
	Files globals.Fileser
	Comments globals.Commentser
	Infoblocks globals.Infoblocker
	Directories globals.Directoryer
	Groups globals.Groupser
	Options globals.Optionser
	Users globals.Userser
	Session globals.Sessioner
	Cookies globals.Cookieser
	Output globals.Outputer
	Permissions globals.Permissionser
	Email globals.Emailer
	EmailTemplates globals.EmailTemplateser
	Request globals.Requester
	DB globals.DataBaser
	Pages globals.Pageer
	Cache globals.Cacher
	Templates globals.Templateser
	Fields globals.Fieldser
	Modules globals.Modules
}

func Init(ins []interface{}) *Core {
	var c = new(Core)

	c.Files          = ins[0].(globals.Fileser)
	c.Comments       = ins[1].(globals.Commentser)
	c.Infoblocks     = ins[2].(globals.Infoblocker)
	c.Directories    = ins[3].(globals.Directoryer)
	c.Groups         = ins[4].(globals.Groupser)
	c.Options        = ins[5].(globals.Optionser)
	c.Users          = ins[6].(globals.Userser)
	c.Session        = ins[7].(globals.Sessioner)
	c.Cookies        = ins[8].(globals.Cookieser)
	c.Output         = ins[9].(globals.Outputer)
	c.Permissions    = ins[10].(globals.Permissionser)
	c.Email          = ins[11].(globals.Emailer)
	c.EmailTemplates = ins[12].(globals.EmailTemplateser)
	c.Request        = ins[13].(globals.Requester)
	c.DB             = ins[14].(globals.DataBaser)
	c.Pages          = ins[15].(globals.Pageer)
	c.Cache          = ins[16].(globals.Cacher)
	c.Templates      = ins[17].(globals.Templateser)
	c.Fields         = ins[18].(globals.Fieldser)
	c.Modules        = ins[19].(globals.Modules)

	return c
}