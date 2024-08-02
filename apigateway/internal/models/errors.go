package models
// @Failure			403  {object} models.ForbiddenError
// @Failure			404  {object} models.NotFound
// @Failure			500  {object} models.StandartError

type ForbiddenError struct{
	Error string `json:"error"`
}

type NotFound struct{
	Error string `json:"error"`
}

type StandartError struct{
	Error string `json:"error"`
}