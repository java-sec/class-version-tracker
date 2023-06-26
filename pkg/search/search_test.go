package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindClass(t *testing.T) {
	classBytes, err := FindClass("./test_data/spring-webmvc-5.3.19.jar", "org.springframework.web.servlet.mvc.method.RequestMappingInfo")
	assert.Nil(t, err)
	assert.NotNil(t, classBytes)
}
