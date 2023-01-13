package session

type store interface {

	//获得 session_id 的值
	ID() string

	//获得 session_id 在 cookie 中的名称
	Name() string

	//获得 session 某一项的值
	Get(name string) string

	//读取 session 某一项的值 如果不存在则返回默认值def
	DefaultGet(name string, def string) string

	//获取所有 session
	All() map[string]string

	//存储一项数据到 session
	Set(name string, value string)

	//删除 session 中一个条目
	Del(name string) bool

	//删除 session 中所有数据
	Clear() bool

	// //判断 Session 里是否存在一个条目，你可以使用 has 方法。如果条目存在且不为 null，has 方法返回 true
	// Has(name string) bool

	// //判断 Session 里是否存在一个即使结果值为 null 的条目
	// Exists(name string) bool

	// //获取并且删除一个条目
	// GetAndRomove(name string) string

	// //删除所有数据
	// Flush() bool
}
