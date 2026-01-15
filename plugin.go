package plugin

import (
	"path"
	gplugin "plugin"
	"strings"

	"github.com/puzpuzpuz/xsync/v4"
)

var pluginmap = xsync.NewMap[string, gplugin.Symbol]()

func Load[T any](plugpath, name string) T {
	if len(plugpath) == 0 {
		plugpath = "./plugins"
	}
	name = strings.ToLower(name)
	obj, ok := pluginmap.Load(name)
	if ok {
		return obj.(T)
	}
	so := path.Join(plugpath, name+".so")
	plug, err := gplugin.Open(so)
	if err != nil {
		panic("插件 " + name + " 加载失败！" + err.Error())
	}
	obj, err = plug.Lookup("Plugin")
	if err != nil {
		panic("插件 " + name + " 加载，无Plugin符号！" + err.Error())
	}
	t, ok := obj.(T)
	if !ok {
		panic("插件 " + name + " 转换T失败！")
	}
	pluginmap.Store(name, t)
	return t
}

func Reload[T any](plugpath, name string) T {
	pluginmap.Delete(name)
	return Load[T](plugpath, name)
}
