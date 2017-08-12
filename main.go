package main

import (
	"fmt"
	"reflect"
	"time"
	//"unsafe"

	"test_orm/orm"

	_ "github.com/mattn/go-oci8"
)

type Order struct {
	Id   int    `orm:"column(ID)"`
	Name string `orm:"column(Name)"`
}

func (u *Order) TableName() string {
	return "ORDER"
}

type User struct {
	Id    int
	Name  string
	Image string
}

type Post struct {
	Id    int    `orm:"auto"`
	Title string `orm:"size(100)"`
	User  *User  `orm:"rel(fk)"`
}

func (u *Post) TableName() string {
	return "post"
}

var aa = 1

func (u *User) TableName() string {
	return "USER"
}

func init() {
	orm.RegisterDriver("oci8", orm.DROracle)
	orm.RegisterDataBase("default", "oci8", "LIULIANG/123456@127.0.0.1:1521/ORCL?loc=America%2FLos_Angeles", 50, 100)
	orm.RegisterModel(new(User), new(Order), new(Post))
	orm.Debug = true

}

func main() {

	o := orm.NewOrm()

	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	if 1 == 1 {

		user := Order{Id: 1026, Name: "AAAAAAAAAA"}
		val := reflect.ValueOf(&user)
		ind := reflect.Indirect(val)
		typ := ind.Type()
		fmt.Println(val.Type())
		fmt.Println(val, ind, typ)
		fmt.Println(ind.Type().Name(), ind.Type().PkgPath())
		o.Test(&user)
		//num, err := o.InsertOrUpdate(&user)
		//fmt.Printf("Affected Num: %v, %v", num, err)
		//fmt.Printf(*(*string)(unsafe.Pointer(uintptr(num))))
		//num, err := o.QueryTable("Post").Filter("title", "313241").Update(orm.Params{
		//	"user_id": "1556889",
		//})
		//fmt.Printf("Affected Num: %v, %v", num, err)
		// SET name = "astaixe" WHERE name = "slene"

		//num, err := o.QueryTable("Post").Filter("title", "313241").Delete()
		//fmt.Printf("Affected Num: %v, %v", num, err)
		// DELETE FROM user WHERE name = "slene"

		/*
			exist := o.QueryTable("Post").Filter("title", "Name").Exist()
			fmt.Printf("Is Exist: %v", exist)

			cnt, err := o.QueryTable("Post").Count() // SELECT COUNT(*) FROM USER
			fmt.Printf("Count Num: %v, %v", cnt, err)
		*/

		///res, err := o.Raw(`select "ID" from "ORDER"  FOR UPDATE`).Exec()
		//if err == nil {
		//	num, _ := res.RowsAffected()
		////	fmt.Println("mysql row affected nums: ", num)
		//}

		res, err := o.Raw(`update  "ORDER" set "Name"='bbbbbbbbbbbAAAAAAA' WHERE ID=100`).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			fmt.Println("mysql row affected nums: ", num)
		}

		res, err = o.Raw(`commit`).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			fmt.Println("mysql row affected nums: ", num)
		}

	} else {
		order := new(Order)
		order.Id = 1007
		order.Name = "是的吗"

		//fmt.Println(o.Insert(order))

		user := Order{Id: 1006}

		err := o.Read(&user)

		fmt.Println(err)
		fmt.Println(user)

		// update
		user.Name = "qqqqqqqqqqqqqqqqqqaBBBBBBBBBBBBBBBBBBBBBBstaxie222222222222222222222222222222222222"
		num, err := o.Update(&user)
		fmt.Printf("NUM: %d, ERR: %v\n", num, err)

		//num, err = o.Delete(&user)
		//fmt.Printf("NUM: %d, ERR: %v\n", num, err)

		{
			for i := 0; i < 2; i++ {
				go func() {
					var posts []*Post
					qs := o.QueryTable("Post")
					num, err := qs.Filter("User__Name", "slene").Filter("User__Image", "kkbbcc.jp").All(&posts)
					fmt.Printf("NUM: %d, ERR: %v\n", num, err)
					//fmt.Printf("%v", posts[0])
					time.Sleep(time.Millisecond)
				}()
			}
		}

		var r orm.RawSeter
		r = o.Raw(`UPDATE "USER" SET  "name"=:name,"image"=:image WHERE "id"=:OLD_id`, "aavvcc", "aaaaaaccddd12312", 100)
		r.Exec()

		{

			users := []User{
				{Id: 318, Name: "slene"},
				{Id: 319, Name: "astaxie"},
				{Id: 310, Name: "unknown"},
			}
			num, err := o.InsertMulti(3, users)
			fmt.Printf("NUM: %d, ERR: %v\n", num, err)
		}
	}
}
