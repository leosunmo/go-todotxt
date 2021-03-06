/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package todotxt

import (
	"testing"
	"time"
)

var (
	testInputTask = "testdata/task_todo.txt"
)

func TestNewTask(t *testing.T) {
	task := NewTask()

	testExpected = 0
	testGot = task.ID
	if testGot != testExpected {
		t.Errorf("Expected new Task to have default ID [%d], but got [%d]", testExpected, testGot)
	}

	testExpected = ""
	testGot = task.Original
	if testGot != testExpected {
		t.Errorf("Expected new Task to be empty, but got [%s]", testGot)
	}

	testExpected = ""
	testGot = task.Todo
	if testGot != testExpected {
		t.Errorf("Expected new Task to be empty, but got [%s]", testGot)
	}

	testExpected = false
	testGot = task.HasPriority()
	if testGot != testExpected {
		t.Errorf("Expected new Task to have no priority, but got [%v]", testGot)
	}

	testExpected = 0
	testGot = len(task.Projects)
	if testGot != testExpected {
		t.Errorf("Expected new Task to have %d projects, but got [%d]", testExpected, testGot)
	}

	testExpected = 0
	testGot = len(task.Contexts)
	if testGot != testExpected {
		t.Errorf("Expected new Task to have %d contexts, but got [%d]", testExpected, testGot)
	}

	testExpected = 0
	testGot = len(task.AdditionalTags)
	if testGot != testExpected {
		t.Errorf("Expected new Task to have %d additional tags, but got [%d]", testExpected, testGot)
	}

	testExpected = true
	testGot = task.HasCreatedDate()
	if testGot != testExpected {
		t.Errorf("Expected new Task to have a created date, but got [%v]", testGot)
	}

	testExpected = false
	testGot = task.HasCompletedDate()
	if testGot != testExpected {
		t.Errorf("Expected new Task to not have a completed date, but got [%v]", testGot)
	}

	testExpected = false
	testGot = task.HasDueDate()
	if testGot != testExpected {
		t.Errorf("Expected new Task to not have a due date, but got [%v]", testGot)
	}

	testExpected = false
	testGot = task.Completed
	if testGot != testExpected {
		t.Errorf("Expected new Task to not be completed, but got [%v]", testGot)
	}
}

func TestParseTask(t *testing.T) {
	task, err := ParseTask("x (C) 2014-01-01 @Go due:2014-01-12 Create golang library documentation +go-todotxt   ")
	if err != nil {
		t.Error(err)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = task.Task()
	if testGot != testExpected {
		t.Errorf("Expected Task to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = 0
	testGot = task.ID
	if testGot != testExpected {
		t.Errorf("Expected Task to have default ID [%d], but got [%d]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 @Go due:2014-01-12 Create golang library documentation +go-todotxt"
	testGot = task.Original
	if testGot != testExpected {
		t.Errorf("Expected Task to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "Create golang library documentation"
	testGot = task.Todo
	if testGot != testExpected {
		t.Errorf("Expected Task to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = true
	testGot = task.HasPriority()
	if testGot != testExpected {
		t.Errorf("Expected Task to have no priority, but got [%v]", testGot)
	}

	testExpected = "C"
	testGot = task.Priority
	if testGot != testExpected {
		t.Errorf("Expected Task to have priority [%v], but got [%v]", testExpected, testGot)
	}

	testExpected = 1
	testGot = len(task.Projects)
	if testGot != testExpected {
		t.Errorf("Expected Task to have %d projects, but got [%d]", testExpected, testGot)
	}

	testExpected = 1
	testGot = len(task.Contexts)
	if testGot != testExpected {
		t.Errorf("Expected Task to have %d contexts, but got [%d]", testExpected, testGot)
	}

	testExpected = 0
	testGot = len(task.AdditionalTags)
	if testGot != testExpected {
		t.Errorf("Expected Task to have %d additional tags, but got [%d]", testExpected, testGot)
	}

	testExpected = true
	testGot = task.HasCreatedDate()
	if testGot != testExpected {
		t.Errorf("Expected Task to have a created date, but got [%v]", testGot)
	}

	testExpected = false
	testGot = task.HasCompletedDate()
	if testGot != testExpected {
		t.Errorf("Expected Task to not have a completed date, but got [%v]", testGot)
	}

	testExpected = true
	testGot = task.HasDueDate()
	if testGot != testExpected {
		t.Errorf("Expected Task to have a due date, but got [%v]", testGot)
	}

	testExpected = true
	testGot = task.Completed
	if testGot != testExpected {
		t.Errorf("Expected Task to be completed, but got [%v]", testGot)
	}
}

func TestTaskID(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)

	taskID := 1
	testGot = testTasklist[taskID-1].ID
	if testGot != taskID {
		t.Errorf("Expected Task[%d] to have ID [%d], but got [%d]", taskID, taskID, testGot)
	}

	taskID = 5
	testGot = testTasklist[taskID-1].ID
	if testGot != taskID {
		t.Errorf("Expected Task[%d] to have ID [%d], but got [%d]", taskID, taskID, testGot)
	}

	taskID = 27
	testGot = testTasklist[taskID-1].ID
	if testGot != taskID {
		t.Errorf("Expected Task[%d] to have ID [%d], but got [%d]", taskID, taskID, testGot)
	}
	taskID++
}

func TestTaskString(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskID := 1

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[taskID-1].String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskID, testExpected, testGot)
	}
	taskID++

	testExpected = "x Download Todo.txt mobile app @Phone"
	testGot = testTasklist[taskID-1].String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskID, testExpected, testGot)
	}
	taskID++

	testExpected = "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = testTasklist[taskID-1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskID, testExpected, testGot)
	}
	taskID++

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[taskID-1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskID, testExpected, testGot)
	}
	taskID++

	testExpected = "x 2014-01-03 2014-01-01 Create some more golang library test cases @Go +go-todotxt"
	testGot = testTasklist[taskID-1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskID, testExpected, testGot)
	}
	taskID++
}

func TestTaskPriority(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskID := 6

	testExpected = "B"
	testGot = testTasklist[taskID-1].Priority
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have priority '%s', but got '%s'", taskID, testExpected, testGot)
	}
	taskID++

	testExpected = "C"
	testGot = testTasklist[taskID-1].Priority
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have priority '%s', but got '%s'", taskID, testExpected, testGot)
	}
	taskID++

	testExpected = "B"
	testGot = testTasklist[taskID-1].Priority
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have priority '%s', but got '%s'", taskID, testExpected, testGot)
	}
	taskID++

	if testTasklist[taskID-1].HasPriority() {
		t.Errorf("Expected Task[%d] to have no priority, but got '%s'", taskID, testTasklist[4].Priority)
	}
	taskID++
}

func TestTaskCreatedDate(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskID := 10

	testExpected, err := time.Parse(DateLayout, "2012-01-30")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskID-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	testExpected, err = time.Parse(DateLayout, "2013-02-22")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskID-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	testExpected, err = time.Parse(DateLayout, "2014-01-01")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskID-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	testExpected, err = time.Parse(DateLayout, "2013-12-30")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskID-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	testExpected, err = time.Parse(DateLayout, "2014-01-01")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskID-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	if testTasklist[taskID-1].HasCreatedDate() {
		t.Errorf("Expected Task[%d] to have no created date, but got '%v'", taskID, testTasklist[4].CreatedDate)
	}
	taskID++
}

func TestTaskContexts(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskID := 16

	testExpected = []string{"Call", "Phone"}
	testGot = testTasklist[taskID-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have contexts '%v', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	testExpected = []string{"Office"}
	testGot = testTasklist[taskID-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have contexts '%v', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	testExpected = []string{"Electricity", "Home", "Of_Super-Importance", "Television"}
	testGot = testTasklist[taskID-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have contexts '%v', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	testExpected = []string{}
	testGot = testTasklist[taskID-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have no contexts, but got '%v'", taskID, testGot)
	}
	taskID++
}

func TestTasksProjects(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskID := 20

	testExpected = []string{"Gardening", "Improving", "Planning", "Relaxing-Work"}
	testGot = testTasklist[taskID-1].Projects
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have projects '%v', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	testExpected = []string{"Novel"}
	testGot = testTasklist[taskID-1].Projects
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have projects '%v', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	testExpected = []string{}
	testGot = testTasklist[taskID-1].Projects
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have no projects, but got '%v'", taskID, testGot)
	}
	taskID++
}

func TestTaskDueDate(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskID := 23

	testExpected, err := time.Parse(DateLayout, "2014-02-17")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskID-1].DueDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have due date '%s', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	if testTasklist[taskID-1].HasDueDate() {
		t.Errorf("Expected Task[%d] to have no due date, but got '%v'", taskID, testTasklist[taskID-1].DueDate)
	}
	taskID++
}

func TestTaskAddonTags(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskID := 25

	testExpected = map[string]string{"Level": "5", "private": "false"}
	testGot = testTasklist[taskID-1].AdditionalTags
	if len(testGot.(map[string]string)) != 2 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have addon tags '%v', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	testExpected = map[string]string{"Importance": "Very!"}
	testGot = testTasklist[taskID-1].AdditionalTags
	if len(testGot.(map[string]string)) != 1 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have projects '%v', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	testExpected = map[string]string{}
	testGot = testTasklist[taskID-1].AdditionalTags
	if len(testGot.(map[string]string)) != 0 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have no additional tags, but got '%v'", taskID, testGot)
	}
	taskID++

	testExpected = map[string]string{}
	testGot = testTasklist[taskID-1].AdditionalTags
	if len(testGot.(map[string]string)) != 0 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have no additional tags, but got '%v'", taskID, testGot)
	}
	taskID++
}

func TestTaskCompleted(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskID := 29

	testExpected = true
	testGot = testTasklist[taskID-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskID, testGot)
	}
	taskID++

	testExpected = true
	testGot = testTasklist[taskID-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskID, testGot)
	}
	taskID++

	testExpected = true
	testGot = testTasklist[taskID-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskID, testGot)
	}
	taskID++

	testExpected = false
	testGot = testTasklist[taskID-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] not to be completed, but got '%v'", taskID, testGot)
	}
	taskID++

	testExpected = false
	testGot = testTasklist[taskID-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] not to be completed, but got '%v'", taskID, testGot)
	}
	taskID++
}

func TestTaskCompletedDate(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskID := 34

	if testTasklist[taskID-1].HasCompletedDate() {
		t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskID, testTasklist[taskID-1].CompletedDate)
	}
	taskID++

	testExpected, err := time.Parse(DateLayout, "2014-01-03")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskID-1].CompletedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have completed date '%s', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	if testTasklist[taskID-1].HasCompletedDate() {
		t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskID, testTasklist[taskID-1].CompletedDate)
	}
	taskID++

	testExpected, err = time.Parse(DateLayout, "2014-01-02")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskID-1].CompletedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have completed date '%s', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	testExpected, err = time.Parse(DateLayout, "2014-01-03")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskID-1].CompletedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have completed date '%s', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++

	if testTasklist[taskID-1].HasCompletedDate() {
		t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskID, testTasklist[taskID-1].CompletedDate)
	}
	taskID++
}

func TestTaskIsOverdue(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskID := 40

	testGot = testTasklist[taskID-1].IsOverdue()
	if !testGot.(bool) {
		t.Errorf("Expected Task[%d] to be overdue, but got '%v'", taskID, testGot)
	}
	taskID++

	testGot = testTasklist[taskID-1].IsOverdue()
	if testGot.(bool) {
		t.Errorf("Expected Task[%d] not to be overdue, but got '%v'", taskID, testGot)
	}
	testTasklist[taskID-1].DueDate = time.Now().AddDate(0, 0, 1)
	testGot = testTasklist[taskID-1].Due()
	if testGot.(time.Duration).Hours() < 23 ||
		testGot.(time.Duration).Hours() > 25 {
		t.Errorf("Expected Task[%d] to be due in 24 hours, but got '%v'", taskID, testGot)
	}
	taskID++

	testGot = testTasklist[taskID-1].IsOverdue()
	if !testGot.(bool) {
		t.Errorf("Expected Task[%d] to be overdue, but got '%v'", taskID, testGot)
	}
	testTasklist[taskID-1].DueDate = time.Now().AddDate(0, 0, -3)
	testGot = testTasklist[taskID-1].Due()
	if testGot.(time.Duration).Hours() < 71 ||
		testGot.(time.Duration).Hours() > 73 {
		t.Errorf("Expected Task[%d] to be due since 72 hours, but got '%v'", taskID, testGot)
	}
	taskID++

	testGot = testTasklist[taskID-1].IsOverdue()
	if testGot.(bool) {
		t.Errorf("Expected Task[%d] not to be overdue, but got '%v'", taskID, testGot)
	}
	taskID++
}

func TestTaskComplete(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskID := 44

	// first 4 tasks should all match the same tests
	for i := 0; i < 4; i++ {
		testExpected = false
		testGot = testTasklist[taskID-1].Completed
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] not to be completed, but got '%v'", taskID, testGot)
		}
		testGot = testTasklist[taskID-1].HasCompletedDate()
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] not to have a completed date, but got '%v'", taskID, testGot)
		}
		testTasklist[taskID-1].Complete()
		testExpected = true
		testGot = testTasklist[taskID-1].Completed
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskID, testGot)
		}
		testGot = testTasklist[taskID-1].HasCompletedDate()
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to have a completed date, but got '%v'", taskID, testGot)
		}
		testExpected = time.Now().Format(DateLayout)
		testGot = testTasklist[taskID-1].CompletedDate.Format(DateLayout)
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to have a completed date of '%v', but got '%v'", taskID, testExpected, testGot)
		}
		taskID++
	}

	testExpected = true
	testGot = testTasklist[taskID-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskID, testGot)
	}
	testGot = testTasklist[taskID-1].HasCompletedDate()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have a completed date, but got '%v'", taskID, testGot)
	}
	testTasklist[taskID-1].Complete()
	testGot = testTasklist[taskID-1].Completed // should be unchanged
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskID, testGot)
	}
	testGot = testTasklist[taskID-1].HasCompletedDate() // should be unchanged
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have a completed date, but got '%v'", taskID, testGot)
	}
	testExpected = "2012-01-01" // should be unchanged
	testGot = testTasklist[taskID-1].CompletedDate.Format(DateLayout)
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have a completed date of '%v', but got '%v'", taskID, testExpected, testGot)
	}
	taskID++
}

func TestTaskReopen(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskID := 49

	// the first 2 tasks should match the same tests
	for i := 0; i < 2; i++ {
		testExpected = true
		testGot = testTasklist[taskID-1].Completed
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskID, testGot)
		}
		testExpected = false
		testGot = testTasklist[taskID-1].HasCompletedDate()
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to have a completed date, but got '%v'", taskID, testGot)
		}
		testTasklist[taskID-1].Reopen()
		testExpected = false
		testGot = testTasklist[taskID-1].Completed
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to not be completed, but got '%v'", taskID, testGot)
		}
		testGot = testTasklist[taskID-1].HasCompletedDate()
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskID, testGot)
		}
		taskID++
	}

	// the next 3 tasks should all match the same tests
	for i := 0; i < 3; i++ {
		testExpected = true
		testGot = testTasklist[taskID-1].Completed
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskID, testGot)
		}
		testGot = testTasklist[taskID-1].HasCompletedDate()
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to have a completed date, but got '%v'", taskID, testGot)
		}
		testTasklist[taskID-1].Reopen()
		testExpected = false
		testGot = testTasklist[taskID-1].Completed
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to not be completed, but got '%v'", taskID, testGot)
		}
		testGot = testTasklist[taskID-1].HasCompletedDate()
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskID, testGot)
		}
		taskID++
	}

	testExpected = false
	testGot = testTasklist[taskID-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskID, testGot)
	}
	testGot = testTasklist[taskID-1].HasCompletedDate()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have a completed date, but got '%v'", taskID, testGot)
	}
	testTasklist[taskID-1].Reopen()
	testGot = testTasklist[taskID-1].Completed // should be unchanged
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskID, testGot)
	}
	testGot = testTasklist[taskID-1].HasCompletedDate() // should be unchanged
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have a completed date, but got '%v'", taskID, testGot)
	}
	taskID++
}

func compareSlices(list1 []string, list2 []string) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i := range list1 {
		if list1[i] != list2[i] {
			return false
		}
	}

	return true
}

func compareMaps(map1 map[string]string, map2 map[string]string) bool {
	if len(map1) != len(map2) {
		return false
	}

	compare := func(map1 map[string]string, map2 map[string]string) bool {
		for key, value := range map1 {
			if value2, found := map2[key]; !found {
				return false
			} else if value != value2 {
				return false
			}
		}
		return true
	}

	return compare(map1, map2) && compare(map2, map1)
}
