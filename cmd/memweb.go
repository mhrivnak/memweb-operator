package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mhrivnak/memweb-reconciler/pkg/handler"
	"github.com/mhrivnak/memweb-reconciler/pkg/reconciler"

	"github.com/operator-framework/operator-sdk-samples/memcached-operator/pkg/apis"
	"github.com/operator-framework/operator-sdk/pkg/log/zap"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("cmd")

func main() {
	logf.SetLogger(zap.Logger())

	config, err := config.GetConfig()
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	scheme := runtime.NewScheme()
	err = apis.AddToScheme(scheme)
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	err = clientgoscheme.AddToScheme(scheme)
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	// this is a non-caching client
	client, err := crclient.New(config, crclient.Options{Scheme: scheme})
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	h := handler.New(reconciler.New(client, scheme))

	http.HandleFunc("/", h.Handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Info("reconciler listening", "port", port)

	log.Error(http.ListenAndServe(fmt.Sprintf(":%s", port), nil), "")
}
