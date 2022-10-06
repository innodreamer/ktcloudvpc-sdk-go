/*
Package noauth creates a "noauth" *ktvpcsdk.ServiceClient for use in Cinder
environments configured with the noauth authentication middleware.

Example of Creating a noauth Service Client

	provider, err := noauth.NewClient(ktvpcsdk.AuthOptions{
		Username:   os.Getenv("OS_USERNAME"),
		TenantName: os.Getenv("OS_TENANT_NAME"),
	})
	client, err := noauth.NewBlockStorageNoAuthV2(provider, noauth.EndpointOpts{
		CinderEndpoint: os.Getenv("CINDER_ENDPOINT"),
	})

	An example of a CinderEndpoint would be: http://example.com:8776/v2,
*/
package noauth
