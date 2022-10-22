package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const tmpl = `[slack-sender]
token = {{.Token}}
`

type TemplateArgs struct {
	Token string
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate initialize configration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := template.New("config").Parse(tmpl)
		if err != nil {
			return err
		}

		tmplArgs := TemplateArgs{
			Token: viper.GetString("init.token"),
		}

		var fs = afero.NewOsFs()

		fmt.Printf("%v\n", viper.AllSettings())
		file, err := fs.OpenFile(viper.GetString("config"), os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}

		if err = t.Execute(file, tmplArgs); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	initCmd.Flags().String("token", "", "token")
	initCmd.Flags().Bool("force", false, "Force write")
	viper.BindPFlag("init.token", initCmd.Flags().Lookup("token"))
	viper.BindPFlag("init.force", initCmd.Flags().Lookup("force"))
}
