package optimizer

import (
	"testing"

	. "github.com/pingcap/check"
	"github.com/pingcap/tidb/ast"
	"github.com/pingcap/tidb/infoschema"
	"github.com/pingcap/tidb/model"
	"github.com/pingcap/tidb/mysql"
	"github.com/pingcap/tidb/optimizer/plan"
	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/util/testleak"
)

var _ = Suite(&testPlanSuite{})

func TestT(t *testing.T) {
	TestingT(t)
}

type testPlanSuite struct{}

func newMockResolve(node ast.Node) error {
	indices := []*model.IndexInfo{
		{
			Name: model.NewCIStr("b"),
			Columns: []*model.IndexColumn{
				{
					Name: model.NewCIStr("b"),
				},
			},
		},
		{
			Name: model.NewCIStr("c_d_e"),
			Columns: []*model.IndexColumn{
				{
					Name: model.NewCIStr("c"),
				},
				{
					Name: model.NewCIStr("d"),
				},
				{
					Name: model.NewCIStr("e"),
				},
			},
		},
	}
	pkColumn := &model.ColumnInfo{
		State: model.StatePublic,
		Name:  model.NewCIStr("a"),
	}
	col1 := &model.ColumnInfo{
		State: model.StatePublic,
		Name:  model.NewCIStr("d"),
	}
	pkColumn.Flag = mysql.PriKeyFlag
	table := &model.TableInfo{
		Columns:    []*model.ColumnInfo{pkColumn, col1},
		Indices:    indices,
		Name:       model.NewCIStr("t"),
		PKIsHandle: true,
	}
	is := infoschema.MockInfoSchema(table)
	return MockResolveName(node, is, "test")
}

func (s *testPlanSuite) TestPredicatePushDown(c *C) {
	plan.UseNewPlanner = true
	defer testleak.AfterTest(c)()
	cases := []struct {
		sql  string
		best string
	}{
		{
			sql:  "select a from (select a from t where d = 0) k where k.a = 5",
			best: "Range(t)->Fields->Fields",
		},
		{
			sql:  "select a from (select 1+2 as a from t where d = 0) k where k.a = 5",
			best: "Table(t)->Fields->Filter->Fields",
		},
		{
			sql:  "select a from (select d as a from t where d = 0) k where k.a = 5",
			best: "Table(t)->Fields->Fields",
		},
		{
			sql:  "select * from t ta join t tb on ta.d = tb.d and ta.d > 1 where tb.a = 0",
			best: "Join{Table(t)->Range(t)}->Fields",
		},
		{
			sql:  "select * from t ta left outer join t tb on ta.d = tb.d and ta.d > 1 where tb.a = 0",
			best: "Join{Table(t)->Table(t)}->Filter->Fields",
		},
	}
	for _, ca := range cases {
		comment := Commentf("for %s", ca.sql)
		stmt, err := parser.ParseOneStmt(ca.sql, "", "")
		c.Assert(err, IsNil, comment)
		ast.SetFlag(stmt)

		err = newMockResolve(stmt)
		c.Assert(err, IsNil)

		p, err := plan.BuildPlan(stmt, nil)
		c.Assert(err, IsNil)

		_, err = plan.PredicatePushDown(p, []ast.ExprNode{})
		c.Assert(err, IsNil)
		err = plan.Refine(p)
		c.Assert(err, IsNil)
		c.Assert(plan.ToString(p), Equals, ca.best, Commentf("for %s", ca.sql))
	}
	plan.UseNewPlanner = false
}
