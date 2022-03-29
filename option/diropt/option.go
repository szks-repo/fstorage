package diropt

type onDirConflictAction int

const (
	NoAction onDirConflictAction = iota + 1
	Remove
	RemoveAll
	ReturnErr
)

type MkdirOption struct {
	OnConflict onDirConflictAction
}

func DefaultMkdirOption() *MkdirOption {
	return &MkdirOption{OnConflict: NoAction}
}
