package internal

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"slamctl/pkg/users"
)

var u users.Users
var c usersConfig

var usersCmd = &cobra.Command{
	Use:              "users",
	Short:            "Operations on users",
	TraverseChildren: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		u = users.NewUsers(cmd.OutOrStdout())
		var err error
		c, err = parseUsersConfig()
		return errors.WithStack(err)
	},
}

var createUsersCmd = &cobra.Command{
	Use:              "create",
	Short:            "Create a new user",
	TraverseChildren: true,
	RunE: func(*cobra.Command, []string) error {
		return u.Create(c.Username, c.Force)
	},
}

func init() {
	usersCmd.PersistentFlags().StringP("username", "u", "", "username")
	if err := usersCmd.MarkPersistentFlagRequired("username"); err != nil {
		logrus.WithError(err).Fatal("failed mark flag as required")
	}
	if err := viper.BindPFlags(usersCmd.PersistentFlags()); err != nil {
		logrus.WithError(err).Fatal("failed to bind config flags")
	}
	rootCmd.AddCommand(usersCmd)

	createUsersCmd.Flags().BoolP("force", "f", false, "deletes user if it is already exist")
	if err := viper.BindPFlags(createUsersCmd.Flags()); err != nil {
		logrus.WithError(err).Fatal("failed to bind config flags")
	}
	usersCmd.AddCommand(createUsersCmd)
}

func parseUsersConfig() (usersConfig, error) {
	var config usersConfig
	err := viper.Unmarshal(&config)
	return config, errors.WithStack(err)
}

type usersConfig struct {
	Username string
	Force    bool
}
