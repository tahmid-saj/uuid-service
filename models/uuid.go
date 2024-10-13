package models

import "uuid-service/uuid"

type Response struct {
	Ok       bool
	Response interface{}
}

func GenerateUUID() (*Response, error) {
	uuidGenerated := uuid.GenerateUUID()

	return &Response{
		Ok: true,
		Response: uuidGenerated,
	}, nil
}