package track

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrack(t *testing.T) {
	//track, err := Track(context.Background(), "org.springframework", "spring-webmvc", "org.springframework.web.servlet.mvc.method.RequestMappingInfo")
	track, err := Track(context.Background(), "org.apache.dubbo", "dubbo", "org.apache.dubbo.rpc.protocol.DelegateExporterMap")
	assert.Nil(t, err)
	t.Log(track)
}
