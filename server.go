package main

import (
	"github.com/codegangsta/martini"
	"github.com/ntuladhar/student-webservice/student"
	"github.com/ntuladhar/student-webservice/webservice"
)

func main(){
	martiniClassic:=martini.Classic()
	student:=new (student.Student)
	webservice.RegisterWebService(student,martiniClassic)
	martiniClassic.Run()
}
