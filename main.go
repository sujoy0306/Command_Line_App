package main
import(
	"os"
	"time"

	"github.com/spf13/cobra"

)

func print1() *cobra.Command{
	return &cobra.Command{
		Use: "curtime",
		RunE: func(cmd *cobra.Command,args []string) error{
			now:=time.Now()
			prettytime:=now.Format(time.RubyDate)
			cmd.Println("Hey groppher, the curenttime is",prettytime)
			return nil
		},
	}
}

func project() *cobra.Command{
	return &cobra.Command{
		Use: "Project",
		RunE: func(cmd *cobra.Command,args []string) error{
			cmd.Println("I have done a project on blockchain")
			return nil
		},
	}
}

func main(){
	cmd:=&cobra.Command{
		Use: "portfolio",
		Short: "Simple CLI for portfolio",
		//Author: "Sujoy",
		Version: "1.0.0",
		SilenceUsage: true,
	}

	cmd.AddCommand(print1())
	cmd.AddCommand(project())
	if err:=cmd.Execute(); err!=nil{
		os.Exit(1)
	}
}