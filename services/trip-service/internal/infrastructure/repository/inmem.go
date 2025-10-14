package repository

type inmemRepository struct {
	trips map[string]*TripModel
	rideFares map[string]*RideFareModel
}

func NewInmemRepository() *inmemRepository {
	return &inmemRepository{
		trups: make(map[string]*domain.TripModel),
		rideFares: make(map[string]*domain.RideFareModel),
	}
}

func (r *inmemRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	r.trips[trip.ID.Hex()] = trip
	return trip, nil
}