package Test

import (
	"database/sql"
	"github.com/pkg/errors"
	"testing"
)

//1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
//以上作业，要求提交到自己的 GitHub 上面，然后把自己的 GitHub 地址填写到班班提供的表单中： https://jinshuju.net/f/o1NSZD
//作业提交截止时间为 3 月 27 日（周日）晚 24:00 分前。
//===========================================
//需要使用；
//原因：此层为业务层，需要通过wrap包装原始堆栈错误信息
//以下代码部分：
func Test02(t *testing.T) {
	var dbObj sql.DB
	rows, err := dbObj.Query("SELECT * FROM TableName")
	if err != nil {
		var msg string
		if err == sql.ErrNoRows {
			msg = "0 records"
		} else {
			msg = "common error"
		}
		errors.Wrap(err, msg)
	}

	for rows.Next(){
		//....
	}
	if err := rows.Close(); err != nil {
		//...
	}
}


