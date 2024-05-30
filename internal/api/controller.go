package api

import (
	"context"
	"encoding/json"
	"main/api/fuelfinder"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ApiGateway struct {
	Handler http.Handler
	client  *fuelfinder.FuelFinderClient
}

func NewGateway(c *fuelfinder.FuelFinderClient, addr string) *http.Server {
	// create a new http server
	mux := http.NewServeMux()
	srv := &http.Server{Addr: addr, Handler: mux}
	gw := &ApiGateway{client: c}

	// attach funcs, then return
	mux.HandleFunc("/", gw.getStations)
	mux.HandleFunc("/ping", gw.getPing)
	mux.HandleFunc("/brands", gw.getBrands)

	//attach middleware to mux
	gw.Handler = log(mux)
	return srv
}

func (*ApiGateway) getPing(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct {
		Message   string `json:"message"`
		Timestamp string `json:"timestamp"`
	}{
		Message:   "pong",
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func (g *ApiGateway) getStations(w http.ResponseWriter, r *http.Request) {
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("latitude"), 32)
	long, _ := strconv.ParseFloat(r.URL.Query().Get("longitude"), 32)
	radius, _ := strconv.ParseFloat(r.URL.Query().Get("radius"), 32)
	brandQuery := r.URL.Query().Get("brands")

	if radius < 1 || radius > 20 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Radius must be between 1 and 20",
		})
		return
	}

	// get brands query string and calculate val
	brands := make([]string, 0)
	if brandQuery != "" {
		brands = strings.Split(brandQuery, ",")
	}

	service := *g.client
	val, err := service.QueryArea(context.TODO(), &fuelfinder.Geofence{
		Latitude:  float32(lat),
		Longitude: float32(long),
		Radius:    float32(radius),
		Brands:    brands,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// write succcessful response
	json.NewEncoder(w).Encode(val.Items)
}

func (g *ApiGateway) getBrands(w http.ResponseWriter, r *http.Request) {
	service := *g.client
	val, err := service.DistinctBrands(r.Context(), &fuelfinder.Empty{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// write succcessful response
	json.NewEncoder(w).Encode(val.Brands)
}
