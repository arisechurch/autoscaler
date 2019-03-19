package flag

import (
	"flag"

	kube_flag "k8s.io/apiserver/pkg/util/flag"
	"k8s.io/klog"
)

// InitFlags initializes the command line flags.
func InitFlags() {
	kube_flag.InitFlags()
	SyncKlogFlags()
}

// SyncKlogFlags makes sure glog and klog can co-exists.
//
// See: https://github.com/kubernetes/klog/blob/master/examples/coexist_glog/coexist_glog.go
func SyncKlogFlags() {
	klogFlags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlags)

	flag.CommandLine.VisitAll(func(f1 *flag.Flag) {
		f2 := klogFlags.Lookup(f1.Name)
		if f2 != nil {
			value := f1.Value.String()
			f2.Value.Set(value)
		}
	})
}
