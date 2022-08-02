package client

import (
	"context"
	gen "github.com/hashicorp-dev-advocates/waypoint-client/pkg/waypoint"
)

type OiDCConfig struct {
	// Client ID of OIDC provider
	ClientId string
	// Client secret of OIDC provider
	ClientSecret string
	// List of scopes to request from teh provider
	Scopes []string
	// List of valid audience values to accept login.
	// This can be used to restrict only certain folks in a shared OIDC domain.
	Auds []string
	// List of allowed redirect URIs, since our redirect URIs are somewhat dynamic (UI for web UI, localhost server for CLI, etc.).
	// This protects against attack since this is not generally recommended
	AllowedRedirectUris []string
	// The signing algorithms supported by the OIDC connect server.
	// If this isn't specified, this will default to RS256 since that should be supported according to the RFC.
	// The string values here should be valid OIDC signing algorithms.
	SigningAlgs []string
	// Discovery URL endpoint to get other information.
	// Required by OIDC.
	DiscoveryUrl string
	// Optional CA certificate chain to validate the discovery URL.
	// Multiple CA certificates can be specified to support easier rotation.
	DiscoveryCaPem []string
	// Mapping claims to keys for usage in selectors such as the "access_selector" on the root auth method.
	// Claim mappings are available as "value.<name>" and list mappings are available as "list.<name>".
	ClaimMappings map[string]string
	// List claim mappings
	ListClaimMappings map[string]string
}

func DefaultOidcConfig() OiDCConfig {
	return OiDCConfig{
		ClientId:            "",
		ClientSecret:        "",
		Scopes:              nil,
		Auds:                nil,
		AllowedRedirectUris: nil,
		SigningAlgs:         nil,
		DiscoveryUrl:        "",
		DiscoveryCaPem:      nil,
		ClaimMappings:       nil,
		ListClaimMappings:   nil,
	}
}

type AuthMethodConfig struct {
	Name           string
	DisplayName    string
	Description    string
	AccessSelector string
	Method         gen.AuthMethod_Oidc
}

func DefaultAuthMethodConfig() AuthMethodConfig {
	return AuthMethodConfig{
		Name:           "",
		DisplayName:    "",
		Description:    "",
		AccessSelector: "",
		Method:         gen.AuthMethod_Oidc{},
	}
}

func (c *waypointImpl) UpsertOidc(ctx context.Context,
	config OiDCConfig,
	amc AuthMethodConfig) (*gen.AuthMethod, error) {

	oidcConfig := gen.AuthMethod_OIDC{
		ClientId:            config.ClientId,
		ClientSecret:        config.ClientSecret,
		Scopes:              config.Scopes,
		Auds:                config.Auds,
		AllowedRedirectUris: config.AllowedRedirectUris,
		SigningAlgs:         config.SigningAlgs,
		DiscoveryUrl:        config.DiscoveryUrl,
		DiscoveryCaPem:      config.DiscoveryCaPem,
		ClaimMappings:       config.ClaimMappings,
		ListClaimMappings:   config.ListClaimMappings,
	}

	uor := &gen.UpsertAuthMethodRequest{
		AuthMethod: &gen.AuthMethod{
			Name:           amc.Name,
			DisplayName:    amc.DisplayName,
			Description:    amc.Description,
			AccessSelector: amc.AccessSelector,
			Method:         &gen.AuthMethod_Oidc{Oidc: &oidcConfig},
		},
	}

	uam, err := c.client.UpsertAuthMethod(ctx, uor)
	if err != nil {
		return nil, err
	}
	return uam.AuthMethod, nil
}

func (c *waypointImpl) DeleteOidc(ctx context.Context, name string) error {
	damr := &gen.DeleteAuthMethodRequest{AuthMethod: &gen.Ref_AuthMethod{Name: name}}

	_, err := c.client.DeleteAuthMethod(ctx, damr)
	if err != nil {
		return err
	}

	return nil
}

func (c *waypointImpl) GetOidcAuthMethod(ctx context.Context, name string) (*gen.GetAuthMethodResponse, error) {

	gamr := &gen.GetAuthMethodRequest{
		AuthMethod: &gen.Ref_AuthMethod{
			Name: name,
		},
	}
	oidcAm, err := c.client.GetAuthMethod(ctx, gamr)
	if err != nil {
		return nil, err
	}

	return oidcAm, nil

}

