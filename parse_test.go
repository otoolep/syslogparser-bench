package bench

import (
	"fmt"
	"testing"

	"github.com/jeromer/syslogparser/rfc5424"
)

func Test_Parsing(t *testing.T) {
	//	msg := `<134>0 2015-05-05T21:20:00.493320+00:00 fisher apache-access - - [exampleSDID@32473 iut="3" eventSource="Application" eventID="1011"] 173.247.206.174 - - [05/May/2015:21:19:52 +0000] "GET /2013/11/ HTTP/1.   1" 200    22056 "http://www.philipotoole.com/" "Wget/1.15 (linux-gnu)"`
	msg := `<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut="3" eventSource="Application" eventID="1011"] An        application event log entry...`
	buf := []byte(msg)
	p := rfc5424.NewParser(buf)

	err := p.Parse()
	if err != nil {
		t.Fatalf("failed to parse RFC5424 message: %s", err.Error())
	}

	for k, v := range p.Dump() {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func Benchmark_Parsing(b *testing.B) {
	msg := `<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut="3" eventSource="Application" eventID="1011"] An        application event log entry...`
	buf := []byte(msg)

	for n := 0; n < b.N; n++ {
		p := rfc5424.NewParser(buf)
		err := p.Parse()
		if err != nil {
			fmt.Println(err.Error())
			panic("unable to parse message during benchmarking")
		}
	}
}
