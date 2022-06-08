package main

import "github.com/gin-gonic/gin"

func main() {
	/*
		user := "123,144,155,2314,2355,9,"
		//pos := strings.LastIndex(user, "9")
		//fmt.Println(user[:pos] + user[pos+len("9")+1:])
		a := strings.Split(user, ",")
		for i := 0; i < len(a)-1; i++ {
			fmt.Println(a[i])
		}
		fmt.Println("this")
	*/

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
