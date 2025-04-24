package dto

type UserProfile struct {
	UserId    uint    `json:"-"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
	Address   Address `json:"address"`
}

type Address struct {
	AddressLine1 *string `json:"address_line1"`
	AddressLine2 *string `json:"address_line2"`
	City         *string `json:"city"`
	PostCode     *string `json:"post_code"`
	Country      *string `json:"country"`
}
