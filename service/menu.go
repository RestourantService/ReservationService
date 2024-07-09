package service

// import (
// 	"context"
// 	"database/sql"
// 	pb "reservation_service/genproto/menu"
// 	"reservation_service/storage/postgres"
// 	"github.com/pkg/errors"
// )

// type MenuService struct {
// 	pb.UnimplementedMenuServer
// 	Repo *postgres.MenuRepo
// }

// func NewMenuService(db *sql.DB) *MenuService {
// 	return &MenuService{Repo: postgres.NewMenuRepo(db)}
// }

// func (m *MenuService) CreateMenu(ctx context.Context, req *pb.MenuDetails) (*pb.ID, error) {
// 	resp, err := m.Repo.CreateMenu(ctx, req)
//     if err!= nil {
//         return nil, errors.Wrap(err, "failed to create menu")
//     }
//     return resp, nil
// }

