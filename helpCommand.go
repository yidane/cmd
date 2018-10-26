package main

type HelpCommand struct {
}

func (HelpCommand) Name() string {
	return "help"
}

func (HelpCommand) Exec(ctx *Context) error {
	return nil
}

func (HelpCommand) Usage() string {
	return "help or help command"
}

func init() {
	Register(HelpCommand{})
}

//func Help(commandName ...string) {
//	buf := bytes.Buffer{}
//	errCmd := []string{}
//	if len(commandName) > 0 {
//		for _, c := range commandName {
//			if cmd, ok := commands[c]; ok {
//				buf.WriteString(fmt.Sprintf(`	|%-20s|%s%s`, c, (*cmd).Usage(), "\n"))
//			} else {
//				errCmd = append(errCmd, c)
//			}
//		}
//	} else {
//		for _, cmd := range commands {
//			buf.WriteString(fmt.Sprintf(`	|%-20s|%s%s`, (*cmd).Name(), (*cmd).Usage(), "\n"))
//		}
//	}
//
//	if buf.Len() > 0 {
//		log.Succeed("command usage as follows：")
//		fmt.Printf(buf.String())
//	}
//	if len(errCmd) > 0 {
//		log.Error(fmt.Sprintf("no such commands:%s", strings.Join(errCmd, ",")))
//	}
//}