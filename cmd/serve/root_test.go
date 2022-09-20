package serve_test

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ory/kratos/internal/testhelpers"
)

func TestServe(t *testing.T) {
	_, _ = testhelpers.StartE2EServer(t, "./stub/kratos.yml", nil)
}

func TestServeTLSBase64(t *testing.T) {
	certPath, keyPath := testhelpers.GenerateTLSCertificateFilesForTests(t)

	certPEM, err := os.ReadFile(certPath)
	require.NoError(t, err)

	keyPEM, err := os.ReadFile(keyPath)
	require.NoError(t, err)

	certBase64 := base64.StdEncoding.EncodeToString(certPEM)
	keyBase64 := base64.StdEncoding.EncodeToString(keyPEM)
	publicPort, adminPort := testhelpers.StartE2EServerOnly(t,
		"./stub/kratos.yml",
		true,
		testhelpers.ConfigOptions{
			"serve.public.tls.key.base64":  keyBase64,
			"serve.public.tls.cert.base64": certBase64,
			"serve.admin.tls.key.base64":   keyBase64,
			"serve.admin.tls.cert.base64":  certBase64,
		},
	)
	testhelpers.CheckE2EServerOnHTTPS(t, publicPort, adminPort)
}

func TestServeTLSPaths(t *testing.T) {
	certPath, keyPath := testhelpers.GenerateTLSCertificateFilesForTests(t)

	publicPort, adminPort := testhelpers.StartE2EServerOnly(t,
		"./stub/kratos.yml",
		true,
		testhelpers.ConfigOptions{
			"serve.public.tls.key.path":  keyPath,
			"serve.public.tls.cert.path": certPath,
			"serve.admin.tls.key.path":   keyPath,
			"serve.admin.tls.cert.path":  certPath,
		},
	)
	testhelpers.CheckE2EServerOnHTTPS(t, publicPort, adminPort)
}
