package dbr

type builder interface {
	Build() (string, error)
}

type functionBuilder interface {
	Build(db *db) (string, error)
}
