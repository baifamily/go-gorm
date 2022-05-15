package root

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)
var cmd  = &cobra.Command{}//创建初始化对象Command

func Execute(){
	cmd.Execute()
}

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
)
func init(){
	cobra.OnInitialize(initConfig)
	//modelToDataCmd.Flags().Boolp("verbose","v",false,"是否显示测试详情")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	cmd.PersistentFlags().StringP("author", "a", "bqzfamily", "author name for copyright attribution")
	cmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	cmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", cmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", cmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	//modelToDataCmd.AddCommand(addCmd)
	cmd.AddCommand(versionCmd)
	cmd.AddCommand(modelToDataCmd)

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
