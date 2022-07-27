package model

import (
	"errors"
	"todo/lib"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
)

type Task struct {
	Id          string `db:"Id" json:"id"`
	Title       string `db:"Title" json:"title"`
	Description string `db:"Description" json:"description"`
	Status      string `db:"Status" json:status`
}

func AddTask(c *gin.Context, task Task, dbmap *gorp.DbMap) (int64, error) {

	if task.Title == "" || task.Description == "" {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	} else {
		if insert, _ := dbmap.Exec(`INSERT INTO Tasks (Title, Description, Status) VALUES (?, ?, ?)`, task.Title, task.Description, task.Status); insert != nil {
			id, err := insert.LastInsertId()
			if err == nil {
				return id, nil
			} else {
				lib.CheckErr(err, "Insert failed")
				return -1, err
			}
		} else {
			c.JSON(400, gin.H{"error": "Fields are empty"})
		}
	}
	return -1, errors.New("of god !")
}

func ListTasks(c *gin.Context, dbmap *gorp.DbMap) ([]Task, error) {
	var tasks []Task
	_, err := dbmap.Select(&tasks, `select * from Tasks`)

	if err == nil {
		return tasks, nil
	} else {
		lib.CheckErr(err, "select failed")
		return []Task{}, err
	}
}

func UpdateTask(c *gin.Context, dbmap *gorp.DbMap) (Task, error) {
	var task Task
	id := c.Params.ByName("id")
	err := dbmap.SelectOne(&task, "SELECT * FROM user WHERE id=?", id)
	if err == nil {
		c.BindJSON(&task)
		_, err := dbmap.Update(&task)
		if err != nil {
			return Task{}, err
		} else {
			return task, nil
		}
	} else {
		return Task{}, nil
	}

}
