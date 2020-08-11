package mailer_test

import (
	"context"
	"testing"

	"allaboutapps.dev/aw/go-starter/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMailerSendPasswordReset(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fixtures := test.Fixtures()

	m := test.NewTestMailer(t)

	passwordResetLink := "http://localhost/password/reset/12345"
	err := m.SendPasswordReset(ctx, fixtures.User1.Username.String, passwordResetLink)
	require.NoError(t, err)

	mt := test.GetTestMailerMockTransport(t, m)

	mail := mt.GetLastSentMail()
	mails := mt.GetSentMails()
	require.NotNil(t, mail)
	require.Len(t, mails, 1)
	assert.Equal(t, m.Config.DefaultSender, mail.From)
	assert.Len(t, mail.To, 1)
	assert.Equal(t, fixtures.User1.Username.String, mail.To[0])
	assert.Equal(t, test.TestMailerDefaultSender, mail.From)
	assert.Equal(t, "Password reset", mail.Subject)
	assert.Contains(t, string(mail.HTML), passwordResetLink)
}
