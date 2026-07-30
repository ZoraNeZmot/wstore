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

type MessageDTO struct {
	Success bool `json:"success"`
	Data    struct {
		Message string `json:"message"`
	} `json:"data"`
}

type ErrorDTO struct {
	Success bool           `json:"success"`
	Error   *ResponseError `json:"error"`
}

// USERS DTOs
type UsersListDTO struct {
	Success bool `json:"success"`
	Data    struct {
		Users []*UserResponse `json:"users"`
	} `json:"data"`
}

type UserDTO struct {
	Success bool          `json:"success"`
	Data    *UserResponse `json:"data"`
}

type TokenDTO struct {
	Success bool `json:"success"`
	Data    struct {
		AccessToken string `json:"accessToken"`
	} `json:"data"`
}

// SALES DTOs
type ItemsListDTO struct {
	Success bool `json:"success"`
	Data    struct {
		Items []*ItemResponse `json:"items"`
	} `json:"data"`
}

type ItemDTO struct {
	Success bool          `json:"success"`
	Data    *ItemResponse `json:"data"`
}
