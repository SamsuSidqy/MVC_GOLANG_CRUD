package UpdateControl

import(
	"fmt"
	"net/http"
	"crudMvc/db/models/Person"
)

func UpdateController(w http.ResponseWriter, r *http.Request){
	fmt.Println("update")
	fmt.Println(r.Method)
	if r.Method == "POST"{		
		result := Person.UpdateData(r.FormValue("id"),r.FormValue("nama"),r.FormValue("alamat"))
		if result {
			http.Redirect(w,r,"/",http.StatusSeeOther)
		}else{
			http.Redirect(w,r,"/",http.StatusSeeOther)			
		}
	}else{
		http.Redirect(w,r,"/",http.StatusSeeOther)
	}

}