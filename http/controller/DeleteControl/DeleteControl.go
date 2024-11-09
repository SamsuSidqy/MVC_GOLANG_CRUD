package DeleteControl

import(
	"fmt"
	"net/http"
	"crudMvc/db/models/Person"
	"encoding/json"
)


type Message struct{
	Pesan string
	Status int
}

func ControllerDelete(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		var resp Message
		params := r.URL.Path[8:]
		if len(params) == 0 {
			http.Redirect(w,r,"/",http.StatusSeeOther)
		}else{
			result := Person.CekData(params)
			if result{
				resultDelete := Person.DeleteData(params)
				fmt.Println(resultDelete)
				resp.Status = 200
				resp.Pesan = "Sukses Hapus"
				w.Header().Set("Content-Type", "application/json")
			    w.WriteHeader(http.StatusOK)
			    respJson,_ := json.Marshal(resp)
			    w.Write(respJson)
			}else{
				resp.Status = 400
				resp.Pesan = "Bad Request"
				w.Header().Set("Content-Type", "application/json")
			    w.WriteHeader(http.StatusBadRequest)
			    respJson,_ := json.Marshal(resp)
			    w.Write(respJson)
			}
		}

	}else{

	}
}