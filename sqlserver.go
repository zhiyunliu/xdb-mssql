package sqlserver

import (
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
	"github.com/zhiyunliu/glue/config"
	contribxdb "github.com/zhiyunliu/glue/contrib/xdb"
	"github.com/zhiyunliu/glue/contrib/xdb/expression"
	"github.com/zhiyunliu/glue/xdb"
)

const Proto = "sqlserver"
const ArgumentPrefix = "p_"

type sqlserverResolver struct {
	name string
}

func (s *sqlserverResolver) Name() string {
	return s.name
}

func (s *sqlserverResolver) Resolve(connName string, setting config.Config, opts ...xdb.Option) (interface{}, error) {
	cfg := contribxdb.NewConfig(connName)
	err := setting.ScanTo(cfg.Cfg)
	if err != nil {
		return nil, fmt.Errorf("读取DB配置(%s):%w", connName, err)
	}
	return contribxdb.NewDB(Proto, cfg, opts...)
}

func init() {
	symbols := expression.DefaultSymbols

	tplMatcher := xdb.NewTemplateMatcher(
		expression.NewNormalExpressionMatcher(symbols),
		expression.NewCompareExpressionMatcher(symbols),
		expression.NewLikeExpressionMatcher(symbols),
		expression.NewInExpressionMatcher(symbols),
	)

	tplstmpProcessor := xdb.NewStmtDbTypeProcessor(DefaultDbTypeHandler...)

	xdb.Register(&sqlserverResolver{name: Proto})
	xdb.RegistTemplate(New(Proto, ArgumentPrefix, tplMatcher, tplstmpProcessor))

	xdb.Register(&sqlserverResolver{name: "mssql"})
	xdb.RegistTemplate(New("mssql", ArgumentPrefix, tplMatcher, tplstmpProcessor))
}
