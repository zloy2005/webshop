package server

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/zloy2005/webshop/internal/service/payment"
	"github.com/zloy2005/webshop/internal/storage"
	"github.com/zloy2005/webshop/web"
)

var config Config

func init() {
	Cmd.Flags().AddFlagSet(config.Flags())
}

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "Start webshop http server",
	RunE: func(cmd *cobra.Command, args []string) error {
		initConfigs(cmd)

		logrus.Info("initialize database")
		db, err := gorm.Open("postgres", config.Storage.String())
		if err != nil {
			return errors.Wrap(err, "failed to open db connection")

		}
		defer db.Close()

		storage := storage.New(db)
		storage.Migrate()

		paySystem := payment.New(storage, config.Payment)
		worker := payment.NewWorker(storage, paySystem)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		worker.Run(ctx)

		srv := web.New(storage, paySystem, config.Server)
		if err := srv.ListenAndServe(); err != nil {
			return err
		}

		return nil
	},
}

func initConfigs(cmd *cobra.Command) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if config.ConfigFile != "" {
		viper.SetConfigFile(config.ConfigFile)
	}
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		logrus.Error(err)
	}
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		logrus.Error(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		logrus.Error(err)
	}

}
