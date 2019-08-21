// Copyright 2019 The OpenSDS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
This module implements a entry into the OpenSDS file share controller service.

*/
package fileshare

import (
	"reflect"
	"testing"
	"context"

	"github.com/opensds/opensds/pkg/dock/client"
	"github.com/opensds/opensds/pkg/model"
	pb "github.com/opensds/opensds/pkg/model/proto"
	. "github.com/opensds/opensds/testutils/collection"
	"google.golang.org/grpc"
)

func TestNewController(t *testing.T) {
	tests := []struct {
		name string
		want Controller
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controller_CreateFileShareAcl(t *testing.T) {
	type fields struct {
		Client   client.Client
		DockInfo *model.DockSpec
	}
	type args struct {
		opt *pb.CreateFileShareAclOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.FileShareAclSpec
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				Client:   tt.fields.Client,
				DockInfo: tt.fields.DockInfo,
			}
			got, err := c.CreateFileShareAcl(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("controller.CreateFileShareAcl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("controller.CreateFileShareAcl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controller_DeleteFileShareAcl(t *testing.T) {
	type fields struct {
		Client   client.Client
		DockInfo *model.DockSpec
	}
	type args struct {
		opt *pb.DeleteFileShareAclOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				Client:   tt.fields.Client,
				DockInfo: tt.fields.DockInfo,
			}
			if err := c.DeleteFileShareAcl(tt.args.opt); (err != nil) != tt.wantErr {
				t.Errorf("controller.DeleteFileShareAcl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type fakefileshareClient struct{}

func (f fakefileshareClient) CreateVolume(ctx context.Context, in *pb.CreateVolumeOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) DeleteVolume(ctx context.Context, in *pb.DeleteVolumeOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) ExtendVolume(ctx context.Context, in *pb.ExtendVolumeOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) CreateVolumeSnapshot(ctx context.Context, in *pb.CreateVolumeSnapshotOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) DeleteVolumeSnapshot(ctx context.Context, in *pb.DeleteVolumeSnapshotOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) CreateVolumeAttachment(ctx context.Context, in *pb.CreateVolumeAttachmentOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) DeleteVolumeAttachment(ctx context.Context, in *pb.DeleteVolumeAttachmentOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) CreateReplication(ctx context.Context, in *pb.CreateReplicationOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) DeleteReplication(ctx context.Context, in *pb.DeleteReplicationOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) EnableReplication(ctx context.Context, in *pb.EnableReplicationOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) DisableReplication(ctx context.Context, in *pb.DisableReplicationOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) FailoverReplication(ctx context.Context, in *pb.FailoverReplicationOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) CreateVolumeGroup(ctx context.Context, in *pb.CreateVolumeGroupOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) UpdateVolumeGroup(ctx context.Context, in *pb.UpdateVolumeGroupOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) DeleteVolumeGroup(ctx context.Context, in *pb.DeleteVolumeGroupOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) CollectMetrics(ctx context.Context, in *pb.CollectMetricsOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) GetMetrics(ctx context.Context, in *pb.GetMetricsOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) GetUrls(ctx context.Context, in *pb.NoParams, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) AttachVolume(ctx context.Context, in *pb.AttachVolumeOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) DetachVolume(ctx context.Context, in *pb.DetachVolumeOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) CreateFileShare(ctx context.Context, in *pb.CreateFileShareOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) DeleteFileShare(ctx context.Context, in *pb.DeleteFileShareOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) CreateFileShareSnapshot(ctx context.Context, in *pb.CreateFileShareSnapshotOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) DeleteFileShareSnapshot(ctx context.Context, in *pb.DeleteFileShareSnapshotOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) CreateFileShareAcl(ctx context.Context, in *pb.CreateFileShareAclOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) DeleteFileShareAcl(ctx context.Context, in *pb.DeleteFileShareAclOpts, opts ...grpc.CallOption) (*pb.GenericResponse, error) {
	return nil, nil
}

func (f fakefileshareClient) Connect(edp string) error {
	return nil
}

func (f fakefileshareClient) Close() {
	return
}

func NewFakeClient() client.Client {
	return &fakefileshareClient{}
}

func NewFakeController() Controller {
	return &controller{
		Client: client.NewClient(),
	}
}

type fakecontroller struct {
	client.Client
	DockInfo *model.DockSpec
}

func (c controller) SetDock(dockInfo *model.DockSpec) {}

func (c controller) CreateFileShare(opt *pb.CreateFileShareOpts) (*model.FileShareSpec, error) {
	return &pb.GenericResponse{
		Reply: &pb.GenericResponse_Result_{
			Result: &pb.GenericResponse_Result{
				Message: ByteFileShare,
			},
		},
	}, nil
}

func (c controller) CreateFileShareAcl(opt *pb.CreateFileShareAclOpts) (*model.FileShareAclSpec, error) {
	return &pb.GenericResponse{
		Reply: &pb.GenericResponse_Result_{
			Result: &pb.GenericResponse_Result{
				Message: ByteFileShareAcl,
			},
		},
	}, nil
}

func (c controller) DeleteFileShareAcl(opt *pb.DeleteFileShareAclOpts) error {
	panic("implement me")
}

func (c controller) DeleteFileShare(opt *pb.DeleteFileShareOpts) error {
	panic("implement me")
}

func (c controller) CreateFileShareSnapshot(opt *pb.CreateFileShareSnapshotOpts) (*model.FileShareSnapshotSpec, error) {
	panic("implement me")
}

func (c controller) DeleteFileShareSnapshot(opts *pb.DeleteFileShareSnapshotOpts) error {
	panic("implement me")
}

func Test_controller_CreateFileShare(t *testing.T) {
	type fields struct {
		Client   client.Client
		DockInfo *model.DockSpec
	}
	type args struct {
		opt *pb.CreateFileShareOpts
	}
	prf := &SampleFileShareProfiles[0]
	req := &pb.CreateFileShareOpts{
		Id:          "d2975ebe-d82c-430f-b28e-f373746a71ca",
		Name:        "sample-fileshare-01",
		Description: "This is a sample fileshare for testing",
		Size:        int64(1),
		Profile:     prf.ToJson(),
		//Context:     c.NewAdminContext().ToJson(),
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.FileShareSpec
		wantErr bool
	}{
		{name:"test1", fields:fields{
			Client:   ,
			DockInfo: nil,
		},args:args{req}, want:&SampleFileShares[0], wantErr:false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				Client:   tt.fields.Client,
				DockInfo: tt.fields.DockInfo,
			}
			_, err := c.CreateFileShare(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("controller.CreateFileShare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("controller.CreateFileShare() = %v, want %v", got, tt.want)
			//}
		})
	}
}

func Test_controller_DeleteFileShare(t *testing.T) {
	type fields struct {
		Client   client.Client
		DockInfo *model.DockSpec
	}
	type args struct {
		opt *pb.DeleteFileShareOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				Client:   tt.fields.Client,
				DockInfo: tt.fields.DockInfo,
			}
			if err := c.DeleteFileShare(tt.args.opt); (err != nil) != tt.wantErr {
				t.Errorf("controller.DeleteFileShare() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_controller_SetDock(t *testing.T) {
	type fields struct {
		Client   client.Client
		DockInfo *model.DockSpec
	}
	type args struct {
		dockInfo *model.DockSpec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				Client:   tt.fields.Client,
				DockInfo: tt.fields.DockInfo,
			}
			c.SetDock(tt.args.dockInfo)
		})
	}
}

func Test_controller_CreateFileShareSnapshot(t *testing.T) {
	type fields struct {
		Client   client.Client
		DockInfo *model.DockSpec
	}
	type args struct {
		opt *pb.CreateFileShareSnapshotOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.FileShareSnapshotSpec
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				Client:   tt.fields.Client,
				DockInfo: tt.fields.DockInfo,
			}
			got, err := c.CreateFileShareSnapshot(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("controller.CreateFileShareSnapshot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("controller.CreateFileShareSnapshot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controller_DeleteFileShareSnapshot(t *testing.T) {
	type fields struct {
		Client   client.Client
		DockInfo *model.DockSpec
	}
	type args struct {
		opt *pb.DeleteFileShareSnapshotOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				Client:   tt.fields.Client,
				DockInfo: tt.fields.DockInfo,
			}
			if err := c.DeleteFileShareSnapshot(tt.args.opt); (err != nil) != tt.wantErr {
				t.Errorf("controller.DeleteFileShareSnapshot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
