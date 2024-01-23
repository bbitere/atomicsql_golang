package gomigration_dialect

type GenericDialectArg struct {
	Connection string
}

func NewGenericDialectArg(connection string) *GenericDialectArg {
	return &GenericDialectArg{Connection: connection}
}
