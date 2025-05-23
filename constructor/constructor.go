package constructor

import "fmt"

//Constructor Function = ฟังก์ชันธรรมดาที่ return struct (หรือ pointer ของ struct)
//Constructor คือ "ประตูแรก" ที่ควบคุมและเตรียม object ให้พร้อมใช้งาน

type person struct {
	Name string
	Age  int
}

// ฟังก์ชัน constructor สำหรับ User
func newUser(name string, age int) *person {
	return &person{
		Name: name,
		Age:  age,
	}
}

func MainConstructor() {
	// สร้าง instance ของ User ด้วย constructor
	user := newUser("สมชาย", 30)
	fmt.Println(*user)
}
