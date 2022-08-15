package main

import (
	"bytes"
	"strings"
	"text/template"
)

var asynqTemplate = `
{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
type {{.ServiceType}}JobServer interface {
{{- range .MethodSets}}
	{{.Name}}(context.Context, *{{.Request}}) (error)
{{- end}}
}

func Register{{.ServiceType}}JobServer(mux *asynq.ServeMux, srv {{.ServiceType}}JobServer) {
	{{- range .Methods}}
	mux.HandleFunc("{{.Typename}}", _{{$svrType}}_{{.Name}}_Job_Handler(srv))
	{{- end}}
}

{{range .Methods}}
func _{{$svrType}}_{{.Name}}_Job_Handler(srv {{$svrType}}JobServer) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, task *asynq.Task) error {
		var in {{.Request}}
		ctx, beforeCtx := trace.Before(ctx, "asynq", task.Type())
		scp := rkgrpcmid.GetServerContextPayload(ctx)		
		if err := json.Unmarshal(task.Payload(), &in); err != nil {
			trace.After(ctx, beforeCtx, err)
			return log.Errorln(task.Type(),log.Any("request",in),log.NewWhy(err))
		}
		scp["req"] = in		
		err := srv.{{.Name}}(ctx, &in)
		if err != nil{
			err=log.Errorln(task.Type(),log.Any("request",in),log.NewWhy(err))
		}else{
			log.Println(task.Type(),log.Any("request",in))
		}		
		trace.After(ctx, beforeCtx, err)
		return err
	}
}
{{end}}

type {{.ServiceType}}SvcJob struct {}
var {{.ServiceType}}Job {{.ServiceType}}SvcJob

{{range .MethodSets}}
func (j *{{$svrType}}SvcJob) {{.Name}}(in *{{.Request}}, opts ...asynq.Option) (*asynq.Task, error) {
	payload, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask("{{.Typename}}", payload, opts...)
	return task, nil
}
{{end}}

type {{.ServiceType}}JobClient interface {
{{- range .MethodSets}}
	{{.Name}}(ctx context.Context, req *{{.Request}}, opts ...asynq.Option) (info *asynq.TaskInfo, err error) 
{{- end}}
}

type {{.ServiceType}}JobClientImpl struct{
	cc *asynq.Client
}
	
func New{{.ServiceType}}JobClient (client *asynq.Client) {{.ServiceType}}JobClient {
	return &{{.ServiceType}}JobClientImpl{client}
}

{{range .MethodSets}}
func (c *{{$svrType}}JobClientImpl) {{.Name}}(ctx context.Context, in *{{.Request}}, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	task, err := {{$svrType}}Job.{{.Name}}(in, opts...)
	if err != nil {
		return nil, log.Errorln("{{$svrType}}Job.{{.Name}}", log.Any("request", in),log.NewWhy(err))	
	}
	log.Println("{{$svrType}}Job.{{.Name}}", log.Any("request", in))	
	info, err := c.cc.Enqueue(task)
	if err != nil {
		return nil, log.Errorln("{{$svrType}}Job.{{.Name}} Enqueue", log.Any("request", in),log.NewWhy(err))
	}
	return info, nil
}
{{end}}
`

type serviceDesc struct {
	ServiceType string // Greeter
	ServiceName string // helloworld.Greeter
	Metadata    string // api/helloworld/helloworld.proto
	Methods     []*methodDesc
	MethodSets  map[string]*methodDesc
}

type methodDesc struct {
	// method
	Name    string
	Num     int
	Request string
	Reply   string
	// asynq rule
	Typename string
}

func (s *serviceDesc) execute() string {
	s.MethodSets = make(map[string]*methodDesc)
	for _, m := range s.Methods {
		s.MethodSets[m.Name] = m
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("asynq").Parse(strings.TrimSpace(asynqTemplate))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
