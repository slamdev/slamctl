package internal

import (
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const rootCmdName = "slamctl"

var rootCmd = &cobra.Command{
	Use:   rootCmdName,
	Short: rootCmdName + " controls execution of various utilities",
	Long: rootCmdName + ` controls execution of various utilities

 Find more information at: https://github.com/slamdev/` + rootCmdName,
	PersistentPreRunE: func(*cobra.Command, []string) error { return setup() },
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		logrus.WithError(err).Fatal("failed to get home directory")
	}

	rootCmd.PersistentFlags().String("config", home+"/."+rootCmdName+".yaml", "config file")
	rootCmd.PersistentFlags().String("log.format", "text", "log format (json or text)")
	rootCmd.PersistentFlags().String("log.level", "error", "log level (debug, info, warn or error)")

	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		logrus.WithError(err).Fatal("failed to bind config flags")
	}
}

func setup() error {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(rootCmdName)

	config, err := parseRootConfig()
	if err != nil {
		return errors.WithStack(err)
	}

	config, err = initConfigFile(config)
	if err != nil {
		return errors.WithStack(err)
	}

	return initLogger(config)
}

func parseRootConfig() (rootConfig, error) {
	var config rootConfig
	err := viper.Unmarshal(&config)
	return config, errors.WithStack(err)
}

func initConfigFile(config rootConfig) (rootConfig, error) {
	if config.ConfigFile != "" {
		// Use usersConfig file from the flag.
		viper.SetConfigFile(config.ConfigFile)
	}

	// If a usersConfig file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.WithError(err).Debug("config file is not found")
		} else {
			return rootConfig{}, errors.WithStack(err)
		}
	} else {
		logrus.Debugf("Using config file: %v", viper.ConfigFileUsed())
	}

	return parseRootConfig()
}

func initLogger(config rootConfig) error {
	level, err := logrus.ParseLevel(config.Log.Level)
	if err != nil {
		return errors.WithStack(err)
	}
	logrus.SetLevel(level)

	if config.Log.Format == "text" {
		logrus.SetFormatter(&logrus.TextFormatter{})
	} else if config.Log.Format == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		return errors.New("unsupported log format")
	}

	return nil
}

func ExecuteCmd() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

type rootConfig struct {
	ConfigFile string `mapstructure:"config"`
	Log        rootConfigLog
}

type rootConfigLog struct {
	Format string
	Level  string
}
