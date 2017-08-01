# module
Шаблон модуля для XR.CMS    
Начните с функций public(*core.Core) и admin(*core.Core)    
Функция public должна возвращать HTML код публичной части модуля,   
функция admin, соответственно, HTML код административной части модуля   
По завершении разработки соберите ваш модуль командой:  
`go build -buildmode=plugin -o my_module_name.so src/module.go`

