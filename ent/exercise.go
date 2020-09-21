// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/francismarcus/eg/ent/exercise"
)

// Exercise is the model entity for the Exercise schema.
type Exercise struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Exercise) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Exercise fields.
func (e *Exercise) assignValues(values ...interface{}) error {
	if m, n := len(values), len(exercise.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	e.ID = int(value.Int64)
	values = values[1:]
	return nil
}

// Update returns a builder for updating this Exercise.
// Note that, you need to call Exercise.Unwrap() before calling this method, if this Exercise
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Exercise) Update() *ExerciseUpdateOne {
	return (&ExerciseClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (e *Exercise) Unwrap() *Exercise {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Exercise is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Exercise) String() string {
	var builder strings.Builder
	builder.WriteString("Exercise(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Exercises is a parsable slice of Exercise.
type Exercises []*Exercise

func (e Exercises) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
