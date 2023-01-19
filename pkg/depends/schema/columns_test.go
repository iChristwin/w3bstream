package schema_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/machinefi/w3bstream/pkg/depends/schema"
)

func TestColumns_Ex(t *testing.T) {
	ctx := context.Background()

	cs, err := schema.Cols("f_id", "f_name")
	NewWithT(t).Expect(err).To(BeNil())

	ex := cs.Ex(ctx)
	NewWithT(t).Expect(ex.Query()).To(Equal("f_id,f_name"))
}

func TestColumns_AutoIncrement(t *testing.T) {
	cs, err := schema.Cols()
	NewWithT(t).Expect(err).To(BeNil())

	col := schema.Col("f_id")
	col.Constrains.AutoIncrement = true

	err = cs.Add(col)
	NewWithT(t).Expect(err).To(BeNil())

	NewWithT(t).Expect(cs.AutoIncrement()).To(Equal(col))
}
