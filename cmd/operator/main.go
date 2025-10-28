package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gigvault/shared/pkg/logger"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
	// TODO: Add GigVault CRD schemes here
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool

	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false, "Enable leader election for controller manager.")
	flag.Parse()

	logger, err := logger.NewProduction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			// Ignore sync errors on stderr/stdout
			_ = err
		}
	}()

	logger.Info("Starting GigVault Operator",
		zap.String("metrics-addr", metricsAddr),
		zap.Bool("leader-election", enableLeaderElection),
	)

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:           scheme,
		LeaderElection:   enableLeaderElection,
		LeaderElectionID: "gigvault-operator",
	})
	if err != nil {
		logger.Fatal("Unable to start manager", zap.Error(err))
	}

	// TODO: Setup controllers here
	// if err = (&controller.CertificateAuthorityReconciler{
	// 	Client: mgr.GetClient(),
	// 	Scheme: mgr.GetScheme(),
	// }).SetupWithManager(mgr); err != nil {
	// 	logger.Fatal("Unable to create controller", zap.Error(err))
	// }

	logger.Info("Starting operator")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		logger.Fatal("Problem running manager", zap.Error(err))
	}
}
