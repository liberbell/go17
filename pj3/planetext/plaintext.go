package planetext

type CarHandler struct {
	car.demo.Car
}

func (h *CarHandler) ServeHTTP(w http.ResponseWriter, r http.Request) {
	fmt.Fprintf
}