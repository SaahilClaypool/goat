package main

import (
	"fmt"
	"os"
	"os/user"
	"encoding/json"
	"log"
	"time"
	"io"
	"bytes"
)
var taskMap map[string]Task 
var fileName string
// read the init file for all of the tasks
func Init(){
	usr , err := user.Current()
	fileName = usr.HomeDir + "/.config/goat"
	taskMap = make(map[string]Task)// create task map
	// open and read file
	input , err := os.Open(fileName)
	if err != nil{
		fmt.Printf("%v\n", err)
		log.Fatal(err)
	}

	dec := json.NewDecoder(input)

	for{
		var a_task Task; 
		err = dec.Decode(&a_task)
		if(err != nil){
			if(err == io.EOF ){
				break; 
			}
			log.Fatal(err)
		}
		if a_task.Active {
			a_task.TimeIntervals[0].End  = time.Now()
		}
		// add task to map
		taskMap[a_task.Name] = a_task
	}
}



func CreateTask (name string) string{
	task , exist := taskMap[name]
	if(exist){
		if task.Active{
			return fmt.Sprintf("task %s is already active" , name)
		} else{
			// append a new time value
			newInt := TimeInterval{
				time.Now(), 
				time.Now(),
			}
			task.Active = true
			task.TimeIntervals = append([]TimeInterval{newInt}, task.TimeIntervals ...)
			taskMap[name] = task 
			
			return fmt.Sprintf("restarted task: %s", name)

		}
	}
	new_task := Task{
		true, 
		name,
		make([]TimeInterval, 0),
		make([]string, 0), 
	}
	now := time.Now()
	new_task.TimeIntervals = append(new_task.TimeIntervals, TimeInterval{now, now})

	taskMap[name] = new_task
	return fmt.Sprintf("created new task: %s", name)

}
func EndTask(name string){
	now := time.Now()

	task , exist := taskMap[name]
	if(exist){
		task.Active = false
		task.TimeIntervals[0].End = now 
		taskMap[name] = task
		fmt.Printf("removed : %s\n" , name)
	}
}
func ListTasks() []string{
	keys := make([]string, 0)
	for k,_ := range(taskMap){
		 keys = append(keys, k)
	}
	return keys
}
func AddNote(taskname string, note string){
	task := taskMap[taskname]
	task.Notes = append(task.Notes, note)
	taskMap[taskname] = task
}

func Info(taskname string) string{
	task := taskMap[taskname]
	return fmt.Sprintf("%s\n", task.String())
}
func WriteTasks (){
	output, err := os.Create(fileName)
	if err != nil{
		log.Fatal(err); 
	}

	for _, v := range (taskMap){
		fmt.Fprintf(output, "%s\n", v.ToJson()); 
	}
}

func CurrentTasks() []Task {
	currentTasks := make([]Task, 0)
	for _,v := range(taskMap){
		if(v.Active){
			currentTasks = append(currentTasks, v)
		}
	}
	return currentTasks
}

type TimeInterval struct {
	Start time.Time ;
	End time.Time ; 
}
func (timeInterval *TimeInterval) duration() time.Duration {
	return timeInterval.End.Sub( timeInterval.Start)
}

type Task struct {
	Active bool 
	Name string
	TimeIntervals []TimeInterval
	// Start_time time / int
	Notes []string; 
}

func (task *Task) totalTime() time.Duration {
	var sum time.Duration
	for _ , interval := range ( task.TimeIntervals){
		sum += interval.duration()
	}
	return sum

}
func (task *Task) String() string{
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Name: %s\nActive: %t\nLast Run Time: %s \nNotes:\n", task.Name, task.Active, task.TimeIntervals[0].duration()))
	for _, nt := range(task.Notes){
		buffer.WriteString(fmt.Sprintf("\t%s\n", nt))
	}
	return buffer.String(); 
	
}
func (task *Task) Write(file *os.File) error{
	fmt.Fprintf(file, "%s\n", task.ToJson())
	return nil
}
func (task *Task) ToJson () []byte {
	jsonStr , err := json.Marshal(task)
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}
	
	return jsonStr
}

