package shared

var (
	log logrus.FieldLogger = logrus.New().WithField("package", "shared")
)

type ErrorHandler func(err error)
type ErrorHandlers []ErrorHandler
