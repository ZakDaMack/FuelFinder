package fuelfinder

import (
	"context"
	"log/slog"
	"main/api/fuelfinder"
	"main/internal/database"

	"google.golang.org/grpc"
)

type Server struct {
	database *database.MongoStore
	fuelfinder.UnimplementedFuelFinderServer
}

func NewGrpcServer(db, uri string) (*grpc.Server, error) {
	server := grpc.NewServer()
	fds, err := NewServer(db, uri)
	if err != nil {
		return nil, err
	}

	// on server start, ensure that the indexes are working, or the queries wont work
	fds.EnsureIndexes()

	// attach server to grpc registrar
	fuelfinder.RegisterFuelFinderServer(server, fds)

	return server, nil
}

func NewServer(db, uri string) (*Server, error) {
	conn, err := database.NewMongoConnection(uri, db)
	if err != nil {
		return nil, err
	}

	fds := &Server{database: conn}
	return fds, nil
}

func (s *Server) QueryArea(ctx context.Context, fence *fuelfinder.Geofence) (*fuelfinder.StationItems, error) {
	queryRes, err := s.database.QueryArea(
		float64(fence.Latitude),
		float64(fence.Longitude),
		int(fence.Radius),
		fence.Brands,
	)

	if err != nil {
		return nil, err
	}

	// REVIEW - is there a better way for this?
	// map to pointers
	items := make([]*fuelfinder.StationItem, len(queryRes))
	for i := 0; i < len(queryRes); i++ {
		items[i] = &queryRes[i]
	}

	results := &fuelfinder.StationItems{Items: items}
	return results, nil
}

func (s *Server) Upload(ctx context.Context, items *fuelfinder.StationItems) (*fuelfinder.UploadedItems, error) {
	res := &fuelfinder.UploadedItems{}

	if len(items.Items) == 0 {
		res.Count = 0
		return res, nil
	}

	// check if files already exists, if so, return
	exists, err := s.database.Exists(items.Items[0].CreatedAt, items.Items[0].SiteId)
	if err != nil {
		return nil, err
	}

	if exists {
		res.Count = 0
		return res, nil
	}

	writeRes, err := s.database.Write(items.Items)
	if err != nil {
		return nil, err
	}

	slog.Debug("uploaded station data", "stations", writeRes)
	res.Count = int32(writeRes)

	return res, nil
}

func (s *Server) DistinctBrands(ctx context.Context, _ *fuelfinder.Empty) (*fuelfinder.Brands, error) {
	brands, err := s.database.GetDistinctBrands()
	if err != nil {
		return nil, err
	}

	res := &fuelfinder.Brands{Brands: brands}
	return res, nil
}

func (s *Server) EnsureIndexes() error {
	ixs, err := s.database.GetIndexes()
	if err != nil {
		return err
	}

	// do these indexes exist?
	s.createIfDoesntExist(ixs, "location", "2dsphere")
	s.createIfDoesntExist(ixs, "created_at", 1)
	s.createIfDoesntExist(ixs, "site_id", 1)
	return nil
}

func (s *Server) createIfDoesntExist(indexes map[string]interface{}, key string, val interface{}) {
	_, ok := indexes[key]
	if !ok {
		s.database.CreateIndex(key, val)
	}
}
