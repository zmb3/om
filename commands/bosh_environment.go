package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/renderers"
)

const protocolPrefix = "://"

type BoshEnvironment struct {
	service         boshEnvironmentService
	logger          logger
	rendererFactory rendererFactory
	opsmanHost      string
	Options         struct {
		ShellType     string `long:"shell-type" description:"Prints for the given shell (posix|powershell)"`
		BoshVars      bool   `long:"bosh" short:"b" description:"Prints the BOSH director environment variables"`
		CredhubVars   bool   `long:"credhub" short:"c" description:"Prints the Credhub environment variables"`
		Unset         bool   `long:"unset" short:"u" description:"Prints unset commands for the environment variables"`
		SSHPrivateKey string `long:"ssh-private-key" short:"i" description:"Location of ssh private key to use to tunnel through the Ops Manager VM. Only necessary if bosh director is not reachable without a tunnel."`
	}
}

//counterfeiter:generate -o ./fakes/bosh_environment_service.go --fake-name BoshEnvironmentService . boshEnvironmentService
type boshEnvironmentService interface {
	GetBoshEnvironment() (api.GetBoshEnvironmentOutput, error)
	ListCertificateAuthorities() (api.CertificateAuthoritiesOutput, error)
}

//counterfeiter:generate -o ./fakes/renderer_factory.go --fake-name RendererFactory . rendererFactory
type rendererFactory interface {
	Create(shellType string) (renderers.Renderer, error)
}

func NewBoshEnvironment(service boshEnvironmentService, logger logger, opsmanHost string, rendererFactory rendererFactory) *BoshEnvironment {
	return &BoshEnvironment{
		service:         service,
		logger:          logger,
		rendererFactory: rendererFactory,
		opsmanHost:      opsmanHost,
	}
}
func (be BoshEnvironment) Target() string {
	if strings.Contains(be.opsmanHost, protocolPrefix) {
		parts := strings.SplitAfter(be.opsmanHost, protocolPrefix)
		return be.stripTrailingSlash(parts[1])
	}

	return be.stripTrailingSlash(be.opsmanHost)
}

func (be BoshEnvironment) stripTrailingSlash(host string) string {
	slashRegex := regexp.MustCompile(`/$`)
	return slashRegex.ReplaceAllString(host, "")
}

func (be BoshEnvironment) Execute(args []string) error {
	renderer, err := be.rendererFactory.Create(be.Options.ShellType)
	if err != nil {
		return err
	}
	boshEnvironment, err := be.service.GetBoshEnvironment()
	if err != nil {
		return err
	}

	certificateAuthorities, err := be.service.ListCertificateAuthorities()
	if err != nil {
		return err
	}

	var boshCACerts string

	for _, ca := range certificateAuthorities.CAs {
		if ca.Active {
			if boshCACerts != "" {
				boshCACerts = boshCACerts + "\n"
			}
			boshCACerts = boshCACerts + ca.CertPEM
		}
	}

	variables := make(map[string]string)
	var unsetVariables []string

	allEnvVars := !be.Options.BoshVars && !be.Options.CredhubVars

	if be.Options.BoshVars || allEnvVars {
		variables["BOSH_CLIENT"] = boshEnvironment.Client
		variables["BOSH_CLIENT_SECRET"] = boshEnvironment.ClientSecret
		variables["BOSH_ENVIRONMENT"] = boshEnvironment.Environment
		variables["BOSH_CA_CERT"] = boshCACerts

		unsetVariables = append(unsetVariables, []string{
			"BOSH_CLIENT",
			"BOSH_CLIENT_SECRET",
			"BOSH_ENVIRONMENT",
			"BOSH_CA_CERT",
			"BOSH_ALL_PROXY",
		}...)
	}

	if be.Options.CredhubVars || allEnvVars {
		variables["CREDHUB_CLIENT"] = boshEnvironment.Client
		variables["CREDHUB_SECRET"] = boshEnvironment.ClientSecret
		variables["CREDHUB_SERVER"] = fmt.Sprintf("https://%s:8844", boshEnvironment.Environment)
		variables["CREDHUB_CA_CERT"] = boshCACerts

		unsetVariables = append(unsetVariables, []string{
			"CREDHUB_CLIENT",
			"CREDHUB_SECRET",
			"CREDHUB_SERVER",
			"CREDHUB_CA_CERT",
			"CREDHUB_PROXY",
		}...)
	}

	if be.Options.SSHPrivateKey != "" {
		file, err := getKeyFilePath(be.Options.SSHPrivateKey)
		if err != nil {
			return err
		}

		proxy := fmt.Sprintf("ssh+socks5://ubuntu@%s:22?private-key=%s", be.Target(), file)

		if be.Options.BoshVars || allEnvVars {
			variables["BOSH_ALL_PROXY"] = proxy
		}

		if be.Options.CredhubVars || allEnvVars {
			variables["CREDHUB_PROXY"] = proxy
		}
	}

	if be.Options.Unset {
		be.renderUnsetVariables(renderer, unsetVariables)
	} else {
		be.renderVariables(renderer, variables)
	}

	return nil
}

func (be BoshEnvironment) renderVariables(renderer renderers.Renderer, variables map[string]string) {
	for k, v := range variables {
		be.logger.Println(renderer.RenderEnvironmentVariable(k, v))
	}
}

func (be BoshEnvironment) renderUnsetVariables(renderer renderers.Renderer, variables []string) {
	for _, variable := range variables {
		be.logger.Println(renderer.RenderUnsetVariable(variable))
	}
}

// Get the absolute path of a key file. Because we currently do no checking to see if the private key exists, if it doesn't, emulate existing behavior.
func getKeyFilePath(s string) (string, error) {
	f, err := os.Open(s)
	if err != nil {
		return "", fmt.Errorf("ssh key file '%s' does not exist", s)
	}
	defer f.Close()

	p, err := filepath.Abs(f.Name())
	if err != nil {
		return s, fmt.Errorf("ssh key file '%s' does not exist", s)
	}

	return p, nil
}
