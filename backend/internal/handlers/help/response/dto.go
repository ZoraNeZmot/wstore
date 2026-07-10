package response

/*
  DTO is mock data transfer object for the response documentation.
  DTO is used in swagger.
  They are just a children of Response struct.

  If such structure CAN'T be implemented using Swagger comments, then create a new DTO.
*/

type FlagDTO struct {
	Success bool `json:"success"`
}
