package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUDPClient(t *testing.T) {
	config := UDPConfig{
		URL: "udp://localhost:8089",
	}
	client, err := NewUDP(config)
	assert.NoError(t, err)

	err = client.Query("ANY QUERY RETURNS NIL")
	assert.NoError(t, err)

	assert.NoError(t, client.Close())
}

func TestNewUDPClient_Errors(t *testing.T) {
	// DialUDP Error
	config := UDPConfig{
		URL: "localhost:8089",
	}
	_, err := NewUDP(config)
	assert.Error(t, err)

	// url.Parse Error
	config = UDPConfig{
		URL: "udp://localhost%35:8089",
	}
	_, err = NewUDP(config)
	assert.Error(t, err)

	// ResolveUDPAddr Error
	config = UDPConfig{
		URL: "udp://localhost:999999",
	}
	_, err = NewUDP(config)
	assert.Error(t, err)
}

func TestUDPClient_Write(t *testing.T) {
	config := UDPConfig{
		URL: "udp://localhost:8089",
	}
	client, err := NewUDP(config)
	assert.NoError(t, err)

	// TODO test write path here

	assert.NoError(t, client.Close())
}
