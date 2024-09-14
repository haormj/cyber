package record

import (
	"errors"
	"fmt"

	"github.com/haormj/cyber/pb"
	"github.com/haormj/util"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type ProtoDesc struct {
	RegistryFiles *protoregistry.Files
}

func NewProtoDesc(message proto.Message) (*ProtoDesc, error) {
	return NewProtoDescFromFileDescriptor(message.ProtoReflect().Descriptor().ParentFile())
}

func NewProtoDescFromFiles(files *protoregistry.Files) *ProtoDesc {
	return &ProtoDesc{
		RegistryFiles: files,
	}
}

func NewProtoDescFromFileDescriptor(fd protoreflect.FileDescriptor) (*ProtoDesc, error) {
	p := &ProtoDesc{
		RegistryFiles: protoregistry.GlobalFiles,
	}
	node, err := p.createDependencyTree(fd)
	if err != nil {
		return nil, fmt.Errorf("createDependencyTree %w", err)
	}

	var fds []protoreflect.FileDescriptor
	node.DFSPostOrder(func(n *util.Node) {
		fds = append(fds, n.Data["fd"].(protoreflect.FileDescriptor))
	})

	files := new(protoregistry.Files)
	for _, fd := range fds {
		_ = files.RegisterFile(fd)
	}

	p.RegistryFiles = files
	return p, nil
}

func NewProtoDescFromBytes(data []byte) (*ProtoDesc, error) {
	p := &ProtoDesc{}

	var recordProtoDesc pb.ProtoDesc
	if err := proto.Unmarshal(data, &recordProtoDesc); err != nil {
		return nil, err
	}

	fdps, err := p.getFileDescriptorProtos(&recordProtoDesc)
	if err != nil {
		return nil, err
	}
	if len(fdps) == 0 {
		return nil, errors.New("not find proto desc")
	}

	files := new(protoregistry.Files)
	for i := len(fdps) - 1; i >= 0; i-- {
		fileDescrptor, err := protodesc.NewFile(fdps[i], files)
		if err != nil {
			return nil, err
		}
		_ = files.RegisterFile(fileDescrptor)
	}
	p.RegistryFiles = files
	return p, nil
}

func (p *ProtoDesc) createDependencyTree(fd protoreflect.FileDescriptor) (*util.Node, error) {
	node := util.NewNode()
	node.Put("path", fd.Path())
	node.Put("fd", fd)

	for i := 0; i < fd.Imports().Len(); i++ {
		fileImport := fd.Imports().Get(i)
		fd, err := p.RegistryFiles.FindFileByPath(fileImport.Path())
		if err != nil {
			return nil, fmt.Errorf("FindFileByPath: %w", err)
		}

		n, err := p.createDependencyTree(fd)
		if err != nil {
			return nil, err
		}

		node.AddChild(n)
	}

	return node, nil
}

func (p *ProtoDesc) getFileDescriptorProtos(recordProtoDesc *pb.ProtoDesc) ([]*descriptorpb.FileDescriptorProto, error) {
	if len(recordProtoDesc.Desc) == 0 {
		return nil, nil
	}

	var fdps []*descriptorpb.FileDescriptorProto

	var fdp descriptorpb.FileDescriptorProto
	if err := proto.Unmarshal(recordProtoDesc.Desc, &fdp); err != nil {
		return nil, err
	}
	fdps = append(fdps, &fdp)

	if len(recordProtoDesc.Dependencies) == 0 {
		return fdps, nil
	}

	for _, rpd := range recordProtoDesc.Dependencies {
		t, err := p.getFileDescriptorProtos(rpd)
		if err != nil {
			return nil, err
		}

		if t == nil {
			continue
		}

		fdps = append(fdps, t...)
	}

	return fdps, nil
}

func (p *ProtoDesc) bytes(fds []protoreflect.FileDescriptor) ([]byte, error) {
	var recordProtoDesc *pb.ProtoDesc
	for i, fd := range fds {
		fdp := protodesc.ToFileDescriptorProto(fd)
		data, err := proto.Marshal(fdp)
		if err != nil {
			return nil, err
		}

		desc := &pb.ProtoDesc{
			Desc: data,
		}

		if i == 0 {
			recordProtoDesc = desc
			continue
		}

		recordProtoDesc.Dependencies = append(recordProtoDesc.Dependencies, desc)
	}

	b, err := proto.Marshal(recordProtoDesc)
	if err != nil {
		return nil, fmt.Errorf("proto.Marshal: %w", err)
	}

	return b, nil
}

func (p *ProtoDesc) Bytes(topicType string) ([]byte, error) {
	d, err := p.RegistryFiles.FindDescriptorByName(protoreflect.FullName(topicType))
	if err != nil {
		return nil, fmt.Errorf("FindDescriptorByName: %w", err)
	}

	n, err := p.createDependencyTree(d.ParentFile())
	if err != nil {
		return nil, fmt.Errorf("createDependencyTree: %w", err)
	}

	var fds []protoreflect.FileDescriptor
	n.DFSPreOrder(func(n *util.Node) {
		fds = append(fds, n.Data["fd"].(protoreflect.FileDescriptor))
	})

	return p.bytes(fds)
}
