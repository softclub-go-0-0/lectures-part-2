package models

import "database/sql"

type Teacher struct {
	ID        int64
	Name      string
	Surname   string
	Phone     string
	Email     sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type Course struct {
	ID         int64
	Title      string
	MonthlyFee int64
	Duration   int64
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
	DeletedAt  sql.NullTime
}

type TimeTable struct {
	ID        int64
	Classroom string
	Start     int64
	Finish    int64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type Student struct {
	ID        int64
	Name      string
	Surname   string
	Phone     string
	Email     sql.NullString
	GroupID   int64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type Group struct {
	ID          int64
	Title       string
	TeacherID   sql.NullInt64
	GroupID     int64
	StartDate   sql.NullTime
	TimetableID sql.NullInt64
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	DeletedAt   sql.NullTime
}
