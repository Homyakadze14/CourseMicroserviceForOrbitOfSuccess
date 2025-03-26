package entities

type Course struct {
	ID              int
	Title           string
	Description     string
	FullDescription string
	Work            string
	Difficulty      string
	Duration        int32
	Image           string
}

type Theme struct {
	ID       int
	CourseID int
	Title    string
}

type Lesson struct {
	ID           int
	CourseID     int
	ThemeID      int
	Title        string
	Type         string
	Duration     int32
	Content      string
	PracticeType string
	Task         string
}
