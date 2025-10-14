package domain

type RideFareModel struct {
	ID primitive.ObjectID
	UserID string
	PackageSlug string //ex: van, luxury, sedan
	TotalPriceInCents float64
}
