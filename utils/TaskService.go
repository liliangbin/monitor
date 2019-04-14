package utils

import (
	"log"
	"monitor/agent"
	"monitor/config"
)

func handle_tasks(tasks []config.Task) {

	for _, t := range tasks {

		var success string = "0"
		cmd := agent.Command{
			Cmd:  t.Script,
			Args: t.ScriptParam,
		}

		if _, err := cmd.Execute(); err == nil {
			success = "1"
			log.Println("task_id=" + t.Taskid + ",t.Script=" + t.Script + ",success=" + success)

		} else
		{

			log.Println("task_id=" + t.Taskid + ",t.Script=" + t.Script + ",success=" + success)
		}
	}

}
