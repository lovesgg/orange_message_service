package services

import (
	"fmt"
	"orange_message_service/app/components/mysql"
	models "orange_message_service/app/models/request"
)

/**
备注 需要您自己新建一张数据表如message
存储的字段由您根据实际参数决定 建议保留的字段:user_id,c_t,u_t,return_data,send_status
保留信息是为了后续排查发送状态等
您也可以忽略这过程 除非对于发送的 结果不care
 */


/**
放入mysql
*/
func InsertDataMysql(params models.ServerReq) bool {
	return true //如果需要保存信息 请修改这方法并注释这行代码

	fmt.Println("insert mysql")
	db := mysql.GetDb()
	defer db.Close()

	ret, err := db.Exec("INSERT INTO message (user_id, send_status, created) VALUES (?, ?, ?)", "18810832200", 1, "2016-06-21")
	if err != nil {
		fmt.Println(ret)
		return false
	}
	return true

}

/**
放入mongodb
*/
func InsertDataMongodb() {

}

/**
放入oracle
*/
func InsertDataOracle() {

}
