package ettt_aws

import (
	"github.com/easy-to-test-tool/ettt"
	"github.com/google/uuid"
)

/*
S3Get
コマンドパラメータ
*/
type S3Get struct {
	Id     uuid.UUID
	Bucket string
	Prefix string
}

func (p S3Get) GetId() uuid.UUID {
	return p.Id
}

/*
Execute
コマンド実行
*/
func (p S3Get) Execute(
	gc ettt.GlobalContext, sc *ettt.ScenarioContext) {
	path, err := ettt.Replace(gc, *sc, p.Bucket+"/"+p.Prefix)
	if err != nil {
		sc.RegistrationCommandResult(ettt.CommandResult{
			Result: ettt.CommandFailure,
			Error:  err,
		})
	}
	println(path)
	sc.RegistrationCommandResult(ettt.CommandResult{})
}
