package student

import (
	"errors"
	"github.com/ntuladhar/student-webservice/db"
	"log"
)

type Student struct {
	Id int `db:"id"`
	Name string `db:"name"`
	City string `db:"city"`
}

type Students struct {
	students []Student
}

func (this *Student) AddStudents(student Student) (Student, error) {
	//newEntry := Student{name, city}
	db, err := db.GetDBConnection()
	if err == nil {
		defer db.Close()
		_, err := db.Exec("insert into student(name,city) values ($1,$2)", student.Name, student.City)
		if err == nil {
			return student, nil
		} else {
			return Student{}, errors.New("Unable to save student to database")
		}

	} else {
		return Student{}, errors.New("Unable to save session to database")
	}

}

func (this *Student) GetStudent(name string) ([]Student, error) {
	if len(name) > 0 {
		db, err := db.GetDBConnection()
		if err == nil {
			defer db.Close()

			//rows:=db.QueryRow(`Select uname,pword from login`)
			results := []Student{}
			db.Select(&results, "Select * from student where name=$1", name)
			//err=rows.Scan(&results.uname,&results.pword)
			if err == nil {

				return results, nil
			} else {
				return results, errors.New("Unable to find Student with student name: " + name)
			}

		} else {
			return []Student{}, errors.New("Unable to find Student with student name: " + name)
		}
	} else {
		return []Student{}, errors.New("Please provide the name ")

	}
}

func (this *Student) GetAllStudents() []Student {
	//students:=new(Students)
	db, err := db.GetDBConnection()
	if err == nil {
		defer db.Close()
		results := []Student{}
		db.Select(&results, "Select * from student")
		log.Println(results)
		return results

	} else {
		return []Student{}
	}

}

