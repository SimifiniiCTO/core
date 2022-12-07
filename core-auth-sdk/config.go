// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package core_auth_sdk

const (
	DefaultKeychainTTL = 60
)

// Config is a configuration struct for Client
type Config struct {
	Issuer         string //the base url of the service handling authentication
	PrivateBaseURL string //overrides the base url for private endpoints
	Audience       string //the domain (host) of the main application
	Username       string //the http basic auth username for accessing private endpoints of the lib issuer
	Password       string //the http basic auth password for accessing private endpoints of the lib issuer
	KeychainTTL    int    //TTL for a key in keychain in minutes
}

func (c *Config) setDefaults() {
	if c.KeychainTTL == 0 {
		c.KeychainTTL = DefaultKeychainTTL
	}
	if c.PrivateBaseURL == "" {
		c.PrivateBaseURL = c.Issuer
	}
}
