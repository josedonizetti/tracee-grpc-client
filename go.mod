module github.com/josedonizetti/tracee-grpc-client

go 1.19

require (
	github.com/aquasecurity/tracee/types v0.0.0-20230803171448-12101ff4a06d
	google.golang.org/grpc v1.56.2
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace github.com/aquasecurity/tracee/types => ../tracee/types
