package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	// temp
	// textObj := TextObject{"text object!", 9}
	app := cli.NewApp()
	app.Name = "goat"
	app.Usage = "time all your daily tasks!"
	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"st", "s"},
			Usage:   "create an entry for the given task name",
			Action: func(c *cli.Context) error {
				fmt.Printf("started task : %s\n", c.Args().First())
				Init()
				CreateTask(c.Args().First())
				WriteTasks()
				return nil

			},
		},
		{
			Name:    "current",
			Usage:   "print current tasks",
			Aliases: []string{"cur"},
			Action: func(c *cli.Context) error {
				Init()
				fmt.Printf("%s\n", CurrentTasks())
				return nil
			},
		},
		{
			Name:    "end",
			Usage:   "end a given task",
			Aliases: []string{"stop"},
			Action: func(c *cli.Context) error {
				Init()
				retString := EndTask(c.Args().First())
				WriteTasks()
				fmt.Printf("Ended Task\n%s\n", retString)
				return nil
			},
		},
		{
			Name:    "list",
			Usage:   "list tasks",
			Aliases: []string{"ls"},
			Action: func(c *cli.Context) error {
				Init()
				fmt.Printf("All Tasks:\n")
				tasks := ListTasks()
				for _, v := range tasks {
					fmt.Println(v)
				}
				// fmt.Printf("%v\n", ListTasks())
				return nil
			},
		},
		{
			Name:  "note",
			Usage: "create a note for a specific task\n\tgoat note [task name] \"a note\"", Aliases: []string{"nt"},
			Action: func(c *cli.Context) error {
				Init()
				fmt.Printf("adding note...\n")
				AddNote(c.Args().First(), c.Args()[1])
				WriteTasks()
				return nil
			},
		},
		{
			Name:    "info",
			Usage:   "detailed info on a task",
			Aliases: []string{"i"},
			Action: func(c *cli.Context) error {
				Init()
				fmt.Printf("%s\n", Info(c.Args()[0]))
				return nil
			},
		},
	}
	// app.Flags = []cli.Flag{
	// 	cli.StringFlag{
	// 		// name task name
	// 		Name: "name"
	// 		// usage
	// 		Usage: "name of the task to be executed"

	// 		// value
	// 		// destination
	// 		Destination: &the_name;

	// 	}
	// }
	// app.Action = func(c *cli.Context) error {
	// 	fmt.Printf("Hello %q\n", c.Args())
	// 	return nil
	// }
	app.Run(os.Args)

}
