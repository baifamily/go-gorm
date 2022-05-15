package root

import (
	"fmt"
	//"github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	//"github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Name string
func init(){
	modelToDataCmd.Flags().StringVarP(&Name,"DbName","t","","表所在的数据库")
}

var modelToDataCmd = &cobra.Command{
	Use: "move",
	Short: "转移表",
	Long: "Constructs a table to the specified database",
	Run: func(cmd *cobra.Command, args []string) {
		err := moveStruct(Name)
		if err != nil {
			log.Fatal(err)
		}
		//cmd.SetHelpCommand(cmd *Command)
		//cmd.SetHelpFunc(f func(*Command, []string))
		//cmd.SetHelpTemplate(s string)
	},
}

func moveStruct(DbName string) error{
	v := viper.New()
	v.AddConfigPath("./conf")//路径
	v.SetConfigName("config")
	v.SetConfigType("ini")

	err := v.ReadInConfig()
	if err !=nil {
		panic(err)
	}
	//[section]default
	s := v.GetString("db.driver")
	fmt.Printf("s: %v\n",s)
	arg := v.GetString("db.username")+":"+v.GetString("db.password")+"@tcp(127.0.0.1:"+v.GetString("db.port")+")/"
	ak := fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local",DbName)
	arg = arg+ak
	fmt.Printf("arg: %v\n",arg)
	db0,err := gorm.Open(mysql.Open(arg), &gorm.Config{})
	fmt.Println(db0)
	db0.AutoMigrate(User{})
	if err != nil {
		log.Fatal("数据库打开出现了问题？？？",err)
		return err
	}
	//db0.Ping()
	if err != nil {
		log.Fatal("数据库连接出现了问题：", err)
		return err
	}
	fmt.Printf("arg: %v\n",arg)	//dsn := fmt.Sprintf("root:bqz113550@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",DbName)
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil{
	//	return err
	//}
	//err = db.AutoMigrate(&User{})
	//if err != nil {
	//	return err
	//}
	return nil
}
//
//// SetHelpFunc sets help function. Can be defined by Application.
//func (c *Command) SetHelpFunc(f func(*Command, []string)){
//	c.helpFunc = f
//}
//
//// SetHlpCommand sets help command.
//func (c *Command) SetHelpCommand(cmd *Command) {
//	c.helpCommand = cmd
//}
//
//// SetHelpTemplate sets help template to be used. Application can use it to set custom template.
//func(c *Command) SetHelpTemplate(s string) {
//	c.helpTemplate = s
//}
//
//
//
//func (c *Command) InitDefaultHelpCmd(){
//	if !c.HasSubCommands(){
//		return
//	}
//	if c.helpCommand == nil{
//		c.helpCommand = &Command{
//			Use: "help [command]",
//			Short: "Help about any command",
//			Long: `Help provides help for any command in the application.
//           Simply type ` + c.Name() + `help [path to command] for full details.`,
//			Run: func(c *Command, args []string){
//				cmd, _, e := c.Root().Find(args)
//				if cmd == nil || e !=nil {
//					c.Printf("Unknown help topic %#q\n", args)
//					c.Root().Usage()
//				} else {
//					cmd.InitDefaultHelpFlag() // make possible 'help' flag to be shown
//					cmd.Help()
//				}
//			},
//		}
//	}
//	c.RemoveCommand(c.helpCommand)
//	c.AddCommand(c.helpCommand)
//}



