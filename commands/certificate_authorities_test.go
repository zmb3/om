package commands_test

import (
    "fmt"

	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/commands/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("Certificate Authorities", func() {
	var (
		certificateAuthorities commands.CertificateAuthorities
		fakeCertificateAuthoritiesService *fakes.CertificateAuthoritiesService
		fakeTableWriter *fakes.TableWriter
	)

	BeforeEach(func() {
		fakeCertificateAuthoritiesService = &fakes.CertificateAuthoritiesService{}
		fakeTableWriter = &fakes.TableWriter{}
		certificateAuthorities = commands.NewCertificateAuthorities(fakeCertificateAuthoritiesService, fakeTableWriter)
	})

	Describe("Execute", func() {
		It("requests certificate authorities from the server", func() {
			err := certificateAuthorities.Execute([]string{})
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeCertificateAuthoritiesService.CertificateAuthoritiesCallCount()).To(Equal(1))
		})

		Context("when request for certificate authorities fails", func() {
			It("returns an error", func() {
				fakeCertificateAuthoritiesService.CertificateAuthoritiesReturns(
					[]commands.CA{},
					fmt.Errorf("could not get certificate authorities"),
				)

				err := certificateAuthorities.Execute([]string{})
				Expect(err).To(MatchError("could not get certificate authorities"))

			})

		})

		It("prints the certificate authorities to a table", func () {
			fakeCertificateAuthoritiesService.CertificateAuthoritiesReturns(
				[]commands.CA{
					{
						GUID: "some-guid",
						Issuer: "Pivotal",
						CreatedOn: "2017-01-09",
						ExpiresOn: "2021-01-09",
						Active: true,
						CertPEM: "-----BEGIN CERTIFICATE-----\nMIIC+zCCAeOgAwIBAgI....",
					},
					{
						GUID: "other-guid",
						Issuer: "Customer",
						CreatedOn: "2017-01-10",
						ExpiresOn: "2021-01-10",
						Active: false,
						CertPEM: "-----BEGIN CERTIFICATE-----\nMIIC+zCCAeOgAwIBBhI....",					
					},
				},
				nil,
			)

			err := certificateAuthorities.Execute([]string{})
			Expect(err).ToNot(HaveOccurred())

			Expect(fakeTableWriter.SetHeaderCallCount()).To(Equal(1))
			Expect(fakeTableWriter.SetHeaderArgsForCall(0)).To(Equal([]string{"id", "issuer", "active", "created on", "expired on", "certicate pem"}))

			Expect(fakeTableWriter.AppendCallCount()).To(Equal(2))
			Expect(fakeTableWriter.AppendArgsForCall(0)).To(Equal([]string{"some-guid", "Pivotal", "true", "2017-01-09", "2021-01-09", "-----BEGIN CERTIFICATE-----\nMIIC+zCCAeOgAwIBAgI...."}))
			Expect(fakeTableWriter.AppendArgsForCall(1)).To(Equal([]string{"other-guid", "Customer", "false", "2017-01-10", "2021-01-10", "-----BEGIN CERTIFICATE-----\nMIIC+zCCAeOgAwIBBhI...."}))

			Expect(fakeTableWriter.RenderCallCount()).To(Equal(1))
		})
	})

})