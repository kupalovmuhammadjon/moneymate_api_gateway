package client

import (
	"api_gateway/configs"

	pb "api_gateway/genproto/budgeting_service"
	pbu "api_gateway/genproto/users"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IServiceManager interface {
	UsersService() pbu.UsersServiceClient
	AccountService() pb.AccountServiceClient
	BudgetService() pb.BudgetServiceClient
	CategoryService() pb.CategoryServiceClient
	GoalService() pb.GoalServiceClient
	TransactionService() pb.TransactionServiceClient
}

type grpcClients struct {
	usersService       pbu.UsersServiceClient
	accountService     pb.AccountServiceClient
	budgetService      pb.BudgetServiceClient
	categoryService    pb.CategoryServiceClient
	goalService        pb.GoalServiceClient
	transactionService pb.TransactionServiceClient
}

func NewGrpcClients(cfg *configs.Config) (IServiceManager, error) {

	connUsersService, err := grpc.NewClient(cfg.UserServiceGrpcHost+cfg.UserServiceGrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connBudgetingService, err := grpc.NewClient(cfg.BudgetingServiceGrpcHost+cfg.BudgetingServiceGrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		usersService:       pbu.NewUsersServiceClient(connUsersService),
		accountService:     pb.NewAccountServiceClient(connBudgetingService),
		budgetService:      pb.NewBudgetServiceClient(connBudgetingService),
		categoryService:    pb.NewCategoryServiceClient(connBudgetingService),
		goalService:        pb.NewGoalServiceClient(connBudgetingService),
		transactionService: pb.NewTransactionServiceClient(connBudgetingService),
	}, nil
}

func (g *grpcClients) UsersService() pbu.UsersServiceClient {
	return g.usersService
}

func (g *grpcClients) AccountService() pb.AccountServiceClient {
	return g.accountService
}

func (g *grpcClients) BudgetService() pb.BudgetServiceClient {
	return g.budgetService
}

func (g *grpcClients) CategoryService() pb.CategoryServiceClient {
	return g.categoryService
}

func (g *grpcClients) GoalService() pb.GoalServiceClient {
	return g.goalService
}

func (g *grpcClients) TransactionService() pb.TransactionServiceClient {
	return g.transactionService
}
