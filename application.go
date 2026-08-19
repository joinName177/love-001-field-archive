package fieldarchive

func NewApplication(s Store, g RenderGateway) *API {
	clips := NewClipService(s)
	return NewAPI(NewRenderService(clips, g))
}
