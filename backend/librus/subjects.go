package librus

type Subject struct {
	Id                int
	Name              string
	IsExtracurricular bool
}

type Lesson struct {
	Id      int
	Subject struct{ Id int }
}

func (L *Librus) MakeLessonSubjectMap() map[int]*CalculatedNode {
	result := map[int]*CalculatedNode{}
	subjects := map[int]string{}

	for _, subject := range L.Subjects {
		if !subject.IsExtracurricular {
			subjects[subject.Id] = subject.Name
		} else {
			subjects[subject.Id] = "*" + subject.Name
		}
	}

	for _, lesson := range L.Lessons {
		node := CalculatedNode{Name: subjects[lesson.Subject.Id]}
		result[lesson.Id] = &node
	}

	return result
}
