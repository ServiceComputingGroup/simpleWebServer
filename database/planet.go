package database

import (
	"encoding/json"
	"fmt"

	"github.com/ServiceComputingGroup/simpleWebServer/entity"
)

func GetPlanet(key string) string {
	key = "https://swapi.co/api/planets/" + key + "/"
	var str string

	datas := GetPlanets()
	for _, v := range datas {

		if v.Url == key {
			encoded, _ := json.MarshalIndent(v, "", "\t")
			str = string(encoded)

		}
	}
	return str
}
func GetPlanets() []entity.Planet {
	//执行查询语句
	rows, err := DB.Query("SELECT * from planet")
	if err != nil {
		fmt.Println("查询出错了")
	}
	var datas []entity.Planet
	//循环读取结果
	for rows.Next() {
		var data entity.Planet
		var str string
		var id int
		err := rows.Scan(&id, &str)
		json.Unmarshal([]byte(str), &data)
		if err != nil {
			fmt.Println("rows fail")
		}
		//将user追加到users的这个数组中
		datas = append(datas, data)
	}
	return datas
}
