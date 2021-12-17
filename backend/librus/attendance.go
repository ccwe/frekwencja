package librus

type Attendance struct {
	Id       int
	Type     struct{ Id int }
	Semester int
	Lesson   Lesson
}

func (L *Librus) GetAttendance(nodes map[int]*CalculatedNode) {
	usedNodes := map[int]bool{}
	for _, attendance := range L.Attendances {
		node, _ := nodes[attendance.Lesson.Id]
		node.Attendance[attendance.Semester-1][attendance.Type.Id%100] += 1
		usedNodes[attendance.Lesson.Id] = true
	}

	for key := range nodes {
		if !usedNodes[key] {
			delete(nodes, key)
		}
	}
}
