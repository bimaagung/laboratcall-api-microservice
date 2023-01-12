package domain

import (
	"context"

	checkupentities "github.com/bimaagung/laboratcall-microservices/checkup/src/Domains/entities"
)

type CheckupRepository interface {
	Add(ctx context.Context, checkup *checkupentities.Checkup)(string)
}