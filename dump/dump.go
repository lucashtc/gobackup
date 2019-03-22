package dump

type gobackup interface {
	GetDatabase() ([]string, error)
}

// Dump ...
func Dump(g gobackup) {
	// m := mysql.New()

	// m.
}
