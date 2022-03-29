package fileopt

type SaveFileOption struct {
	OnConflict onFileConflictAction
}

type onFileConflictAction int

const (
	ReturnErr onFileConflictAction = iota + 1
	Overwrite
	Append
	NoAction
)
