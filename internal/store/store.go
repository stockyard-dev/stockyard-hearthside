package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Retro struct {
	ID string `json:"id"`
	SprintName string `json:"sprint_name"`
	WentWell string `json:"went_well"`
	ToImprove string `json:"to_improve"`
	ActionItems string `json:"action_items"`
	Participants string `json:"participants"`
	Date string `json:"date"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"hearthside.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS retros(id TEXT PRIMARY KEY,sprint_name TEXT NOT NULL,went_well TEXT DEFAULT '[]',to_improve TEXT DEFAULT '[]',action_items TEXT DEFAULT '[]',participants TEXT DEFAULT '',date TEXT DEFAULT '',status TEXT DEFAULT 'open',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Retro)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO retros(id,sprint_name,went_well,to_improve,action_items,participants,date,status,created_at)VALUES(?,?,?,?,?,?,?,?,?)`,e.ID,e.SprintName,e.WentWell,e.ToImprove,e.ActionItems,e.Participants,e.Date,e.Status,e.CreatedAt);return err}
func(d *DB)Get(id string)*Retro{var e Retro;if d.db.QueryRow(`SELECT id,sprint_name,went_well,to_improve,action_items,participants,date,status,created_at FROM retros WHERE id=?`,id).Scan(&e.ID,&e.SprintName,&e.WentWell,&e.ToImprove,&e.ActionItems,&e.Participants,&e.Date,&e.Status,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Retro{rows,_:=d.db.Query(`SELECT id,sprint_name,went_well,to_improve,action_items,participants,date,status,created_at FROM retros ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Retro;for rows.Next(){var e Retro;rows.Scan(&e.ID,&e.SprintName,&e.WentWell,&e.ToImprove,&e.ActionItems,&e.Participants,&e.Date,&e.Status,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Update(e *Retro)error{_,err:=d.db.Exec(`UPDATE retros SET sprint_name=?,went_well=?,to_improve=?,action_items=?,participants=?,date=?,status=? WHERE id=?`,e.SprintName,e.WentWell,e.ToImprove,e.ActionItems,e.Participants,e.Date,e.Status,e.ID);return err}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM retros WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM retros`).Scan(&n);return n}

func(d *DB)Search(q string, filters map[string]string)[]Retro{
    where:="1=1"
    args:=[]any{}
    if q!=""{
        where+=" AND (1=0)"
        
    }
    if v,ok:=filters["status"];ok&&v!=""{where+=" AND status=?";args=append(args,v)}
    rows,_:=d.db.Query(`SELECT id,sprint_name,went_well,to_improve,action_items,participants,date,status,created_at FROM retros WHERE `+where+` ORDER BY created_at DESC`,args...)
    if rows==nil{return nil};defer rows.Close()
    var o []Retro;for rows.Next(){var e Retro;rows.Scan(&e.ID,&e.SprintName,&e.WentWell,&e.ToImprove,&e.ActionItems,&e.Participants,&e.Date,&e.Status,&e.CreatedAt);o=append(o,e)};return o
}

func(d *DB)Stats()map[string]any{
    m:=map[string]any{"total":d.Count()}
    rows,_:=d.db.Query(`SELECT status,COUNT(*) FROM retros GROUP BY status`)
    if rows!=nil{defer rows.Close();by:=map[string]int{};for rows.Next(){var s string;var c int;rows.Scan(&s,&c);by[s]=c};m["by_status"]=by}
    return m
}
