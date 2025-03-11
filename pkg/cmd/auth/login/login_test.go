package login

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/internal/config"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	mockfactory "github.com/berrybytes/zocli/pkg/utils/factory/mock"

	"github.com/stretchr/testify/assert"
)

func TestLoginCmd(t *testing.T) {
	tt := []struct {
		name string
		args []string
		out  string
		err  string
	}{
		{
			name: "help",
			out:  grammar.LoginHelp,
			args: []string{"--help"},
			err:  "",
		},
		{
			name: "random input",
			out:  "",
			args: []string{"--random input"},
			err:  "unknown flag: --random input",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			f := factory.New(context.Background(), config.New())
			f.LoggedIn = true
			loginCmd := NewLoginCommand(f)

			loginCmd.SetArgs(tc.args)

			stdin := new(bytes.Buffer)
			stderr := new(bytes.Buffer)

			loginCmd.SetIn(stdin)
			loginCmd.SetOut(stderr)
			loginCmd.SetErr(stderr)

			_, err := loginCmd.ExecuteC()
			if tc.err != "" {
				assert.EqualError(t, err, tc.err)
			} else {
				assert.NoError(t, err)
			}

			if tc.out != "" {
				assert.Contains(t, stderr.String(), tc.out)
			}
		})
	}
}

func TestSaveDetails(t *testing.T) {
	invalidConfigFolder := config.New()
	invalidConfigFolder.ConfigFolder = "/root/invalid"

	tt := []struct {
		name   string
		option Opts
		err    string
		panics bool
	}{
		{
			name: "invalid config folder",
			option: Opts{
				F: factory.New(context.Background(), invalidConfigFolder),
			},
			err:    "permission denied",
			panics: false,
		},
		{
			name: "invalid config folder and auth file",
			option: Opts{
				F: factory.New(context.Background(), &config.Config{
					AuthFile:     "/invalid",
					ConfigFolder: "/root/invalid"}),
			},
			err:    "permission denied",
			panics: false,
		},
		{
			name: "valid everything",
			option: Opts{
				F: mockfactory.NewFactory(),
				LoginResponse: &api.LoginResponse{
					Data: api.Data{
						AuthToken: "this is the test token",
						User: api.User{
							Id:    1,
							Email: "fakemail@mail.com",
						},
					},
				},
			},
			err:    "",
			panics: true,
		},
	}

	for _, tc := range tt {
		ti := tc
		t.Run(ti.name, func(t *testing.T) {
			opts := ti.option
			defer func() {
				opts.F.CleanUpTest()
			}()
			if ti.panics {
				assert.Panicsf(t, func() {
					err := opts.saveDetails()
					if err != nil {
						fmt.Println(err)
					}
				}, "panic expected")
			} else {
				err := opts.saveDetails()
				assert.Contains(t, err.Error(), ti.err)
			}
		})
	}
}
