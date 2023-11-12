package main

type Student struct {
	ID    string
	Name  string
	Grade int32
}

var students = []*Student{}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, eachElm := range students {
		if eachElm.ID == id {
			return eachElm
		}
	}

	return nil
}
