package delivery

import (
	"bytes"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	//postgre implementation
	_ "github.com/lib/pq"

	"github.com/rtawormy14/cakman-go/util/database"
)

// Historier interface
type Historier interface {
	GetHistory(int64) ([]History, error)
	Insert(History, *sqlx.Tx) error
}

// History struct
type History struct {
	ID         int64     `db:"id" json:"id"`
	DeliveryID int64     `db:"delivery_id" json:"-"`
	Status     string    `db:"status" json:"status"`
	Note       string    `db:"note" json:"note"`
	CreateTime time.Time `db:"create_time" json:"timestamp"`
}

// NewHistory is ...
func NewHistory() Historier {
	return &History{}
}

// GetHistory will return History object if exist
func (d *History) GetHistory(deliveryID int64) (histories []History, err error) {
	db := database.DB

	histories = make([]History, 0)

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT * FROM delivery_history ")
	queryBuffer.WriteString("WHERE delivery_id = $1 ")
	queryBuffer.WriteString("ORDER BY create_time DESC ")

	query := db.Rebind(queryBuffer.String())
	err = db.Select(&histories, query, deliveryID)

	if err != nil && err != sql.ErrNoRows {
		log.Printf("[History][GetHistoryList] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}

	return
}

// Insert will insert History data
func (d *History) Insert(history History, tx *sqlx.Tx) (err error) {
	commitNow := false
	db := database.DB
	if tx == nil {
		tx = db.MustBegin()
		commitNow = true
	}

	query := "INSERT INTO delivery_history (delivery_id,status,note,create_time) VALUES ($1,$2,$3,$4)"
	query = tx.Rebind(query)
	tx.MustExec(query, history.DeliveryID, history.Status, history.Note, history.CreateTime)

	if commitNow {
		err = tx.Commit()
		if err != nil {
			log.Printf("[History][Insert] Error when commit to db : %s \n data : %v ", err, history)
			err = errors.New("error when insert data to History")
		}
	}
	return err
}
