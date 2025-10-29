package models_test

import (
	"testing"

	"github.com/formancehq/connector-framework/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestAccountTypeFromString(t *testing.T) {
	assert.Equal(t, models.ACCOUNT_TYPE_INTERNAL, models.AccountTypeFromString("INTERNAL"))
	assert.Equal(t, models.ACCOUNT_TYPE_EXTERNAL, models.AccountTypeFromString("EXTERNAL"))
	assert.Equal(t, models.ACCOUNT_TYPE_UNKNOWN, models.AccountTypeFromString(""))
	assert.Equal(t, models.ACCOUNT_TYPE_UNKNOWN, models.AccountTypeFromString("someval"))
}
