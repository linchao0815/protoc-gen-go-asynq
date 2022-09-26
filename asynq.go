package main

import (
	"fmt"

	asynq "github.com/linchao0815/protoc-gen-go-asynq/proto"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	contextPackage = protogen.GoImportPath("context")
	asynqPackage   = protogen.GoImportPath("github.com/hibiken/asynq")
	emptyPackage   = protogen.GoImportPath("google.golang.org/protobuf/types/known/emptypb")
	protoPackage   = protogen.GoImportPath("google.golang.org/protobuf/proto")
	jsonPackage    = protogen.GoImportPath("encoding/json")
)

var methodSets = make(map[string]int)

// generateFile generates a _asynq.pb.go file.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 || (!hasAsynqRule(file.Services)) {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + "_asynq.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-asynq. DO NOT EDIT.")
	g.P("// versions:")
	g.P(fmt.Sprintf("// protoc-gen-go-asynq %s", release))
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	generateFileContent(gen, file, g)
	return g
}

// generateFileContent generates the kratos errors definitions, excluding the package statement.
func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}
	g.P("import	\"igspkg/logger\"")
	g.P("import	\"igspkg/igstrace\"")
	g.P("import	\"net/http\"")
	g.P("import	\"strings\"")
	g.P("import	\"go.opentelemetry.io/otel/propagation\"")
	g.P("import	\"go.opentelemetry.io/otel/attribute\"")
	//g.P("import rkgrpcmid \"github.com/rookie-ninja/rk-grpc/v2/middleware\"")
	g.P("import rkgrpcctx \"github.com/rookie-ninja/rk-grpc/v2/middleware/context\"")
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the asynq package it is being compiled against.")
	g.P("var _ = new(", contextPackage.Ident("Context"), ")")
	g.P("var _ = new(", asynqPackage.Ident("Task"), ")")
	g.P("var _ = new(", emptyPackage.Ident("Empty"), ")")
	g.P("var _ = new(", protoPackage.Ident("Message"), ")")
	g.P("var _ = new(", jsonPackage.Ident("InvalidUTF8Error"), ")")

	for _, service := range file.Services {
		genService(gen, file, g, service)
	}
}

func genService(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
		g.P(deprecationComment)
	}
	// Job Server.
	sd := &serviceDesc{
		ServiceType: service.GoName,
		ServiceName: string(service.Desc.FullName()),
		Metadata:    file.Desc.Path(),
	}
	for _, method := range service.Methods {
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			continue
		}
		rule, ok := proto.GetExtension(method.Desc.Options(), asynq.E_Task).(*asynq.Task)
		if rule != nil && ok {
			sd.Methods = append(sd.Methods, buildMethodDesc(g, method, rule.Typename))
		}
	}
	if len(sd.Methods) != 0 {
		g.P(sd.execute())
	}
}

func hasAsynqRule(services []*protogen.Service) bool {
	for _, service := range services {
		for _, method := range service.Methods {
			if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
				continue
			}
			rule, ok := proto.GetExtension(method.Desc.Options(), asynq.E_Task).(*asynq.Task)
			if rule != nil && ok {
				return true
			}
		}
	}
	return false
}

func buildMethodDesc(g *protogen.GeneratedFile, m *protogen.Method, t string) *methodDesc {
	return &methodDesc{
		Name:     m.GoName,
		Num:      methodSets[m.GoName],
		Request:  g.QualifiedGoIdent(m.Input.GoIdent),
		Reply:    g.QualifiedGoIdent(m.Output.GoIdent),
		Typename: t,
	}
}

const deprecationComment = "// Deprecated: Do not use."
