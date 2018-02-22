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
	c.Groups         = ins[3].(globals.Groupser)
	c.Options        = ins[4].(globals.Optionser)
	c.Users          = ins[5].(globals.Userser)
	c.Session        = ins[6].(globals.Sessioner)
	c.Cookies        = ins[7].(globals.Cookieser)
	c.Output         = ins[8].(globals.Outputer)
	c.Permissions    = ins[9].(globals.Permissionser)
	c.Email          = ins[10].(globals.Emailer)
	c.EmailTemplates = ins[11].(globals.EmailTemplateser)
	c.Request        = ins[12].(globals.Requester)
	c.DB             = ins[13].(globals.DataBaser)
	c.Pages          = ins[14].(globals.Pageer)
	c.Cache          = ins[15].(globals.Cacher)
	c.Templates      = ins[16].(globals.Templateser)
	c.Fields         = ins[17].(globals.Fieldser)
	c.Modules        = ins[18].(globals.Modules)

	return c
}