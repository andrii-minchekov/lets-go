package impl

import (
	"github.com/andrii-minchekov/lets-go/app/impl/db"
	"github.com/andrii-minchekov/lets-go/app/usecases"
	cfg "github.com/andrii-minchekov/lets-go/domain"
)

type composedUseCases struct {
	uc.SnippetUseCase
	uc.UserUseCase
}

func NewComposedUseCases(config cfg.Config) uc.UseCases {
	return composedUseCases{
		uc.NewSnippetUseCase(db.NewOrmSnippetRepository()),
		uc.NewUserUseCase(db.NewDbUserRepository()),
	}
}
