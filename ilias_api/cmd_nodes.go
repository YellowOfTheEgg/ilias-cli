package ilias_api

import (	
	"io/ioutil"
	"gopkg.in/yaml.v3"
	"fmt"
	"os"
)


type Cmd struct {
	Name string `yaml:"cmdName"`
	Node string `yaml:"cmdNode"`
}

func getCmdNodeFromFile(cmd_name string) (node string, err error) {
	buffer, err := ioutil.ReadFile("cmd_nodes.yml")
	if err != nil {		
		return "", err
	}

	var cmds []Cmd
	if err := yaml.Unmarshal(buffer, &cmds);err !=nil {		
		return "", err
	}
	
	for _, cmd:=range cmds{
		if cmd.Name == cmd_name {				
			return cmd.Node, nil
		}
	}
	fmt.Fprintf(os.Stderr,"Error: could not find a node hash for cmdName: %v. Check cmd_nodes.yml\n", cmd_name)
	os.Exit(1)
	return "", nil
}