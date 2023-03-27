package govalidator

import (
	"errors"
	"fmt"
	"go-es/global"
	"strings"

	"github.com/thedevsaddam/govalidator"
)

// 自定义规则 not_exists，验证请求数据必须不存在于数据库中。
func init() {
	// 常用于保证数据库某个字段的值唯一，如用户名、邮箱、手机号、或者分类的名称。
	// not_exists:users,phone 检查数据库表里是否存在同一条信息
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		tableName := rng[0]            // 第一个参数，表名称，如 users
		dbFiled := rng[1]              // 第二个参数，字段名称，如 phone
		requestValue := value.(string) // 用户请求过来的数据

		// 查询数据库
		var count int64
		global.GDB.Table(tableName).Where(fmt.Sprintf("%v = ?", dbFiled), requestValue).Count(&count)
		if count != 0 {
			if message != "" {
				return errors.New(message)
			}

			return fmt.Errorf("%v 已存在", requestValue)
		}

		return nil
	})
}
