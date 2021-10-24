package model

type ErrorChangedColumn string

func (e ErrorChangedColumn) Error() string {
	return string(e)
}
