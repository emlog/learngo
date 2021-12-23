package note

import (
	"log"

	"github.com/emlog/goexample/model/request"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type NoteDao struct {
	*xorm.Engine // 数据库连接
}

func NewNoteDao() *NoteDao {
	// DSN (Data Source Name)
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	engine, err := xorm.NewEngine("mysql", "root:12345678@(127.0.0.1)/test_note?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}

	engine.ShowSQL(true)

	return &NoteDao{engine}
}

// Create 创建
func (d NoteDao) Create(note *request.ReqNote) (id int64, err error) {

	// 执行插入操作
	affected, err := d.Table("note").Insert(note)
	if err != nil {
		return 0, err
	}

	// 返回影响行数
	return affected, nil
}
