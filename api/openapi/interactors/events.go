package interactors

/*
func useEvents(service *web.Service, store api.Store) {
	useCRUDL(
		service,
		withPrefix("/events"),
		withCreate(usecase.NewInteractor(func(ctx context.Context, input api.CreateEventRequest, output api.CreateEventResponse) error {
			_, err := store.CreateEvent(input)
			if err != nil {
				return err
			}

			return nil
		})),
		withRead(usecase.NewInteractor(func(ctx context.Context, input api.ReadEventRequest, output api.ReadEventResponse) error {
			_, err := store.ReadEvent(input)
			if err != nil {
				return err
			}

			return nil
		})),
		withUpdate(usecase.NewInteractor(func(ctx context.Context, input api.UpdateEventRequest, output api.UpdateEventResponse) error {
			_, err := store.UpdateEvent(input)
			if err != nil {
				return err
			}

			return nil
		})),
		withDelete(usecase.NewInteractor(func(ctx context.Context, input api.DeleteEventRequest, output api.DeleteEventResponse) error {
			_, err := store.DeleteEvent(input)
			if err != nil {
				return err
			}

			return nil
		})),
		withList(usecase.NewInteractor(func(ctx context.Context, input api.ListEventsRequest, output api.ListEventsResponse) error {
			_, err := store.ListEvents(input)
			if err != nil {
				return err
			}

			return nil
		})),
	)
}
*/
