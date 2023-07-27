package game

// type API struct {
// 	logger     *logger.Logger
// 	validator  *validator.Validate
// 	repository *Repository
// }

// func New(logger *logger.Logger, validator *validator.Validate) *API {
// 	return &API{
// 		logger:     logger,
// 		validator:  validator,
// 		repository: NewRepository(db),
// 	}
// }

type API struct{}

func New() *API {
	return &API{}
}
