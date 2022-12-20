package logging

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type DBLogger struct {
}

func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	b, err := q.FormattedQuery()
	if err != nil {
		return err
	}
	_, err = fmt.Println(string(b))
	return err
}
