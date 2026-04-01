package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-hearthside/internal/store")
func(s *Server)handleListRetros(w http.ResponseWriter,r *http.Request){list,_:=s.db.ListRetros();if list==nil{list=[]store.Retro{}};writeJSON(w,200,list)}
func(s *Server)handleCreateRetro(w http.ResponseWriter,r *http.Request){var ret store.Retro;json.NewDecoder(r.Body).Decode(&ret);if ret.Sprint==""{writeError(w,400,"sprint required");return};s.db.CreateRetro(&ret);writeJSON(w,201,ret)}
func(s *Server)handleListItems(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);list,_:=s.db.ListItems(id);if list==nil{list=[]store.Item{}};writeJSON(w,200,list)}
func(s *Server)handleAddItem(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);var item store.Item;json.NewDecoder(r.Body).Decode(&item);item.RetroID=id;if item.Body==""{writeError(w,400,"body required");return};if item.Category==""{item.Category="good"};s.db.AddItem(&item);writeJSON(w,201,item)}
func(s *Server)handleVoteItem(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.VoteItem(id);writeJSON(w,200,map[string]string{"status":"voted"})}
func(s *Server)handleAddAction(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);var a store.Action;json.NewDecoder(r.Body).Decode(&a);a.RetroID=id;if a.Description==""{writeError(w,400,"description required");return};s.db.AddAction(&a);writeJSON(w,201,a)}
func(s *Server)handleOverview(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}
