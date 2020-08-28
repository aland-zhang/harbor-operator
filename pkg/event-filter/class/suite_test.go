package class_test

import (
	"context"
	"testing"

	. "github.com/goharbor/harbor-operator/pkg/event-filter/class"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/goharbor/harbor-operator/pkg/factories/logger"
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t, "EventFilter", []Reporter{printer.NewlineReporter{}})
}

func setupTest(ctx context.Context) (*Filter, context.Context) {
	logger.Set(&ctx, zap.LoggerTo(GinkgoWriter, true))

	return &Filter{}, ctx
}
