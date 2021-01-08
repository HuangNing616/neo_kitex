package model

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //
	"strings"
	"time"
)

/**
	Author: huang ning
	Date: 2020/01/06
	Func: MySQL数据库连接并且实现CRUD操作
	Note:
	1. Go将数据库操作分为两类：Query与Exec, 区别在于前者会返回结果，而后者不会
	2. QueryRow表示只返回一行的查询，作为Query的一个常见特例。
	3. Prepare表示后续需要多次使用的语句，并且可以传入参数
	4. 打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	5. 方法名大写，就是public
*/

type FeatureGroupConfig struct {
	Id            int64     `json:"id" gorm:"column:id"`                                             //
	Name          string    `json:"name" gorm:"column:name"`                                         //
	ConfigDesc    string    `json:"config_desc" gorm:"column:config_desc"`                           //
	KeyType       string    `json:"key_type" gorm:"column:key_type"`                                 // abase key type, uid、did、uid_app_id, groupKey
	IsDeleted     string    `json:"is_deleted" gorm:"column:is_deleted;default:'0'"`                 // 未删除为0；已删除的记录删除时间
	Expire        string    `json:"expire" gorm:"column:expire"`                                     // 过期时间
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time;default:CURRENT_TIMESTAMP"` // 创建时间
	UpdateTime    time.Time `json:"update_time" gorm:"column:update_time;default:CURRENT_TIMESTAMP"` // 更新时间
	CreateUser    string    `json:"create_user" gorm:"column:create_user"`                           // 创建用户
	UpdateUser    string    `json:"update_user" gorm:"column:update_user"`                           // 更新用户
	VersionConfig string    `json:"version_config" gorm:"column:version_config"`                     // 版本信息
	SaveMergeLog  bool       `json:"save_merge_log" gorm:"column:save_merge_log"`					 // 是否保存merge日志
	AbaseConfig   string    `json:"abase_config" gorm:"abase_config"`                                // abase相关配置
}

//Db数据库连接池
var DB *sql.DB

//数据库配置
const (
	userName = "root"
	password = "rootroot"
	ip       = "10.227.8.243"
	port     = "3306"
	dbName   = "uap_db"
)

// 初始化数据库
func InitDB() {

	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)

	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("open database success!!!")
}

// 查询操作
func Query(){

	var featureGroupConfig FeatureGroupConfig
	sql := "select * from feature_group_config"
	rows, err := DB.Query(sql)
	if err == nil {
		_ = errors.New("query incur error")
	}
	for rows.Next() {
		e := rows.Scan(&featureGroupConfig.Id, &featureGroupConfig.Name, &featureGroupConfig.ConfigDesc, &featureGroupConfig.KeyType,
			&featureGroupConfig.IsDeleted, &featureGroupConfig.Expire, &featureGroupConfig.CreateTime, &featureGroupConfig.UpdateTime,
			&featureGroupConfig.CreateUser, &featureGroupConfig.UpdateUser, &featureGroupConfig.VersionConfig,
			&featureGroupConfig.SaveMergeLog, &featureGroupConfig.AbaseConfig)
		//if e != nil {
		//	fmt.Printf("%+v \n", featureGroupConfig)
		//}
		if e == nil{
			fmt.Println("当前查询为空")
		}
	}
	_ = rows.Close()

	stmt, _ := DB.Prepare("select * from feature_group_config where id=?")
	query, _ := stmt.Query(2)
	fmt.Println("===================")
	fmt.Println(query.Scan(&featureGroupConfig.Id))
}

// 插入操作
func Insert(id int, name string) bool {

	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}

	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO feature_group_config (`id`, `name`) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}

	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(id, name)
	if err != nil {
		fmt.Println("Exec insert fail")
		return false
	}

	//事务提交
	_ = tx.Commit()

	fmt.Println("上次插入id")
	fmt.Println(res.LastInsertId())
	return true
}

// 删除操作
func Delete(id int) bool{

	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}

	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM feature_group_config WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}

	//设置参数以及执行sql语句
	_, e := stmt.Exec(id)
	if e != nil {
		fmt.Println("Exec delete fail")
		return false
	}
	_ = tx.Commit()
	return true

}

