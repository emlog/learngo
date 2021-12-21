package note

import (
	"database/sql"

	"github.com/emlog/goexample/model/request"
	_ "github.com/go-sql-driver/mysql"
)

type NoteDao struct {
	*sql.DB // 数据库连接
}

func NewNoteDao() *NoteDao {
	// DSN (Data Source Name)
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	db, err := sql.Open("mysql", "root:12345678@(127.0.0.1)/test_note")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	return &NoteDao{db}
}

// Create 创建
func (d NoteDao) Create(req *request.ReqNote) (id int64, err error) {
	// 执行插入操作
	result, err := d.Exec("INSERT INTO note(content) VALUES(?)", req.Content)
	if err != nil {
		return
	}

	// 获取插入数据的主键
	id, err = result.LastInsertId()
	if err != nil {
		return
	}

	return id, nil
}
