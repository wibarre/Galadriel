// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package datastore

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createJoinTokenStmt, err = db.PrepareContext(ctx, createJoinToken); err != nil {
		return nil, fmt.Errorf("error preparing query CreateJoinToken: %w", err)
	}
	if q.createTrustDomainStmt, err = db.PrepareContext(ctx, createTrustDomain); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTrustDomain: %w", err)
	}
	if q.deleteJoinTokenStmt, err = db.PrepareContext(ctx, deleteJoinToken); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteJoinToken: %w", err)
	}
	if q.deleteTrustDomainStmt, err = db.PrepareContext(ctx, deleteTrustDomain); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTrustDomain: %w", err)
	}
	if q.findJoinTokenStmt, err = db.PrepareContext(ctx, findJoinToken); err != nil {
		return nil, fmt.Errorf("error preparing query FindJoinToken: %w", err)
	}
	if q.findJoinTokenByIDStmt, err = db.PrepareContext(ctx, findJoinTokenByID); err != nil {
		return nil, fmt.Errorf("error preparing query FindJoinTokenByID: %w", err)
	}
	if q.findJoinTokensByTrustDomainIDStmt, err = db.PrepareContext(ctx, findJoinTokensByTrustDomainID); err != nil {
		return nil, fmt.Errorf("error preparing query FindJoinTokensByTrustDomainID: %w", err)
	}
	if q.findJoinTokensByTrustDomainNameStmt, err = db.PrepareContext(ctx, findJoinTokensByTrustDomainName); err != nil {
		return nil, fmt.Errorf("error preparing query FindJoinTokensByTrustDomainName: %w", err)
	}
	if q.findTrustDomainByIDStmt, err = db.PrepareContext(ctx, findTrustDomainByID); err != nil {
		return nil, fmt.Errorf("error preparing query FindTrustDomainByID: %w", err)
	}
	if q.findTrustDomainByNameStmt, err = db.PrepareContext(ctx, findTrustDomainByName); err != nil {
		return nil, fmt.Errorf("error preparing query FindTrustDomainByName: %w", err)
	}
	if q.listJoinTokensStmt, err = db.PrepareContext(ctx, listJoinTokens); err != nil {
		return nil, fmt.Errorf("error preparing query ListJoinTokens: %w", err)
	}
	if q.listTrustDomainsStmt, err = db.PrepareContext(ctx, listTrustDomains); err != nil {
		return nil, fmt.Errorf("error preparing query ListTrustDomains: %w", err)
	}
	if q.updateJoinTokenStmt, err = db.PrepareContext(ctx, updateJoinToken); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateJoinToken: %w", err)
	}
	if q.updateTrustDomainStmt, err = db.PrepareContext(ctx, updateTrustDomain); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTrustDomain: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createJoinTokenStmt != nil {
		if cerr := q.createJoinTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createJoinTokenStmt: %w", cerr)
		}
	}
	if q.createTrustDomainStmt != nil {
		if cerr := q.createTrustDomainStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTrustDomainStmt: %w", cerr)
		}
	}
	if q.deleteJoinTokenStmt != nil {
		if cerr := q.deleteJoinTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteJoinTokenStmt: %w", cerr)
		}
	}
	if q.deleteTrustDomainStmt != nil {
		if cerr := q.deleteTrustDomainStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTrustDomainStmt: %w", cerr)
		}
	}
	if q.findJoinTokenStmt != nil {
		if cerr := q.findJoinTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findJoinTokenStmt: %w", cerr)
		}
	}
	if q.findJoinTokenByIDStmt != nil {
		if cerr := q.findJoinTokenByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findJoinTokenByIDStmt: %w", cerr)
		}
	}
	if q.findJoinTokensByTrustDomainIDStmt != nil {
		if cerr := q.findJoinTokensByTrustDomainIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findJoinTokensByTrustDomainIDStmt: %w", cerr)
		}
	}
	if q.findJoinTokensByTrustDomainNameStmt != nil {
		if cerr := q.findJoinTokensByTrustDomainNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findJoinTokensByTrustDomainNameStmt: %w", cerr)
		}
	}
	if q.findTrustDomainByIDStmt != nil {
		if cerr := q.findTrustDomainByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findTrustDomainByIDStmt: %w", cerr)
		}
	}
	if q.findTrustDomainByNameStmt != nil {
		if cerr := q.findTrustDomainByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findTrustDomainByNameStmt: %w", cerr)
		}
	}
	if q.listJoinTokensStmt != nil {
		if cerr := q.listJoinTokensStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listJoinTokensStmt: %w", cerr)
		}
	}
	if q.listTrustDomainsStmt != nil {
		if cerr := q.listTrustDomainsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTrustDomainsStmt: %w", cerr)
		}
	}
	if q.updateJoinTokenStmt != nil {
		if cerr := q.updateJoinTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateJoinTokenStmt: %w", cerr)
		}
	}
	if q.updateTrustDomainStmt != nil {
		if cerr := q.updateTrustDomainStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTrustDomainStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                  DBTX
	tx                                  *sql.Tx
	createJoinTokenStmt                 *sql.Stmt
	createTrustDomainStmt               *sql.Stmt
	deleteJoinTokenStmt                 *sql.Stmt
	deleteTrustDomainStmt               *sql.Stmt
	findJoinTokenStmt                   *sql.Stmt
	findJoinTokenByIDStmt               *sql.Stmt
	findJoinTokensByTrustDomainIDStmt   *sql.Stmt
	findJoinTokensByTrustDomainNameStmt *sql.Stmt
	findTrustDomainByIDStmt             *sql.Stmt
	findTrustDomainByNameStmt           *sql.Stmt
	listJoinTokensStmt                  *sql.Stmt
	listTrustDomainsStmt                *sql.Stmt
	updateJoinTokenStmt                 *sql.Stmt
	updateTrustDomainStmt               *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                  tx,
		tx:                                  tx,
		createJoinTokenStmt:                 q.createJoinTokenStmt,
		createTrustDomainStmt:               q.createTrustDomainStmt,
		deleteJoinTokenStmt:                 q.deleteJoinTokenStmt,
		deleteTrustDomainStmt:               q.deleteTrustDomainStmt,
		findJoinTokenStmt:                   q.findJoinTokenStmt,
		findJoinTokenByIDStmt:               q.findJoinTokenByIDStmt,
		findJoinTokensByTrustDomainIDStmt:   q.findJoinTokensByTrustDomainIDStmt,
		findJoinTokensByTrustDomainNameStmt: q.findJoinTokensByTrustDomainNameStmt,
		findTrustDomainByIDStmt:             q.findTrustDomainByIDStmt,
		findTrustDomainByNameStmt:           q.findTrustDomainByNameStmt,
		listJoinTokensStmt:                  q.listJoinTokensStmt,
		listTrustDomainsStmt:                q.listTrustDomainsStmt,
		updateJoinTokenStmt:                 q.updateJoinTokenStmt,
		updateTrustDomainStmt:               q.updateTrustDomainStmt,
	}
}