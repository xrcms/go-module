/**
	Шаблон модуля для XR.CMS
	Начните с функций public(*core.Core) и admin(*core.Core)
	Функция public должна возвращать HTML код публичной части модуля,
	функция admin, соответственно, HTML код административной части модуля
	По завершении разработки соберите ваш модуль командой:
	go build -buildmode=plugin -o my_module_name.so src/module.go
 */
package main

import (
	"./core"
	"encoding/json"
	"fmt"
	"github.com/xrcms/globals"
)

// Описание модуля
var manifest, _ = json.Marshal(globals.Manifest{
	Name: "Шаблон модуля",
	Icon: "fa-font",
	Version: "1.00.00",
	ReleaseDate: "2017-07-23",
	Description: "Заготовка модуля для XR.CMS",
	Author: globals.Author {
		Name: "Вершинин Павел Николаевич",
		Email: "feedback@xr-cms.ru",
		Site: "https://xr-cms.ru",
	},
	HomePage: "https://xr-cms.ru",
	MinimumCmsVersion: "1.01.000",
})

func Init(arguments ...interface{}) []byte {
	var entryPoint = arguments[0].(string)

	// Программа запросила информацию о модуле
	// Сразу, ничего не делая, вернём наш манифест
	if entryPoint == "manifest" {
		return manifest
	}

	// Подготовим ядро API
	var c = core.Init(arguments[1:])

	switch entryPoint {
		case "public": // Запрошена публичная часть модуля, за неё отвечает функция public(*core.Core)
			return public(c)
		case "admin": // Запрошена административная часть модуля, за неё отвечает функция admin(*core.Core)
			return admin(c)
		// ...
		// Любые другие, если необходимо, точки входа, дописываем тут, по аналогии с public и admin
	}

	return []byte(fmt.Sprintf(`Неверная точка входа "%s"`, entryPoint))
}

func public(core *core.Core) []byte {
	// @TODO Код необходимый для пользовательской части модуля
	return nil
}

func admin(core *core.Core) []byte {
	if allow, err := core.Permissions.Check("admin/modules", core.Session.User().Groups); !allow {
		return []byte(err.Error())
	}
	// @TODO Код, необходимый для административной части модуля, если модулю админка не нужна, можно оставить так
	return nil
}