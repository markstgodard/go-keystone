package keystone_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/markstgodard/go-keystone/keystone"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const tokensResp = `{
    "token": {
        "methods": [
            "password"
        ],
        "expires_at": "2015-11-06T15:32:17.893769Z",
        "extras": {},
        "user": {
            "domain": {
                "id": "default",
                "name": "Default"
            },
            "id": "423f19a4ac1e4f48bbb4180756e6eb6c",
            "name": "admin",
            "password_expires_at": null
        },
        "audit_ids": [
            "ZzZwkUflQfygX7pdYDBCQQ"
        ],
        "issued_at": "2015-11-06T14:32:17.893797Z"
    }
}
}
`

var _ = Describe("Keystone API", func() {

	var (
		client *keystone.Client
		server *httptest.Server
	)

	Describe("NewClient", func() {
		var err error

		It("requires a URL", func() {
			client, err = keystone.NewClient("http://192.168.56.101:5000")
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when URL is missing", func() {
			It("returns an error", func() {
				client, err = keystone.NewClient("")
				Expect(err).To(MatchError("missing URL"))
			})
		})
	})

	Describe("Tokens", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set(keystone.X_SUBJECT_TOKEN_HEADER, "fake-but-valid-token")
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(tokensResp))
			}))
			var err error
			client, err = keystone.NewClient(server.URL)
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			server.Close()
		})

		Context("when user / password is valid", func() {
			It("returns a valid token", func() {
				auth := keystone.Auth{
					Identity: keystone.Identity{
						Methods: []string{"password"},
						Password: keystone.IdentityUser{
							User: keystone.User{
								Name:     "admin",
								Password: "password1",
								Domain: keystone.Domain{
									Name: "Default",
								},
							},
						},
					},
				}
				token, err := client.Tokens(auth)
				Expect(err).ToNot(HaveOccurred())
				Expect(token).To(Equal("fake-but-valid-token"))
			})
		})
	})

})
