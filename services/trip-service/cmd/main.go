package main

func main() {
	inmemRepo := repository.NewInmemRepository()
	tripService := service.NewService(inmemRepo)

	fare := &domain.RideFareModel{
		UserId: "42",
	}
	t, err := tripService.CreateTrip(context.Background(), )
	if err != nil {
		log.Println(err)
	}

	log.Println(t)
	
	//temporary
	for {
		time.Sleep(1 * time.Second)
	}
}