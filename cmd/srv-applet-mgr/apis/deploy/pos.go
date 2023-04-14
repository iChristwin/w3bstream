package deploy

import (
	"context"

	"github.com/machinefi/w3bstream/cmd/srv-applet-mgr/apis/middleware"
	"github.com/machinefi/w3bstream/pkg/depends/base/types"
	"github.com/machinefi/w3bstream/pkg/depends/kit/httptransport/httpx"
	"github.com/machinefi/w3bstream/pkg/modules/deploy"
)

type CreateInstance struct {
	httpx.MethodPost
	AppletID                           types.SFID `in:"path" name:"appletID"`
	deploy.CreateOrReDeployInstanceReq `in:"body"`
}

func (r *CreateInstance) Path() string {
	return "/applet/:appletID"
}

func (r *CreateInstance) Output(ctx context.Context) (interface{}, error) {
	ca := middleware.CurrentAccountFromContext(ctx)
	ctx, err := ca.WithAppletContext(ctx, r.AppletID)
	if err != nil {
		return nil, err
	}
	return deploy.CreateInstance(ctx, &r.CreateOrReDeployInstanceReq)
}

type ReDeployInstance struct {
	httpx.MethodPost
	AppletID                           types.SFID `in:"path" name:"appletID"`
	InstanceID                         types.SFID `in:"path" name:"instanceID"`
	deploy.CreateOrReDeployInstanceReq `in:"body"`
}

func (r *ReDeployInstance) Path() string {
	return "/applet/:appletID/:instanceID"
}

func (r *ReDeployInstance) Output(ctx context.Context) (interface{}, error) {
	ca := middleware.CurrentAccountFromContext(ctx)
	ctx, err := ca.WithAppletContext(ctx, r.AppletID)
	ctx, err = ca.WithInstanceContext(ctx, r.InstanceID)
	if err != nil {
		return nil, err
	}
	return deploy.ReDeployInstance(ctx, &r.CreateOrReDeployInstanceReq)
}
