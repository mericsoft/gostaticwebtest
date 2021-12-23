package main
import (
"fmt"
"net/http"
_"strings"
_"log"
"github.com/gorilla/mux"
"html/template"
)

type HtmlData struct{
	Ad string
	Yas string
}

var Welcome = http.HandlerFunc(
	func (w http.ResponseWriter,r *http.Request){
		w.Write([]byte("welcome mux test"))
})

var Sinifac = http.HandlerFunc(
		func (w http.ResponseWriter,r *http.Request){
			vars:=mux.Vars(r)
			fmt.Fprintf(w,"%s sinifi %s hoca ile açıldı",vars["sinifad"],vars["hocaad"])
})

var Mesajat = http.HandlerFunc(
	func (w http.ResponseWriter,r *http.Request){
		frm:="<html> <body><form name='mesaj' method='post' action='/mesajoku'><input name='isim'></input> <input name='yas' </input> <input type='submit' value='yolla'></input> </form> </body> </html>"
		fmt.Fprintf(w,frm)
		
		
	})

var Mesajal = http.HandlerFunc(
	func (w http.ResponseWriter,r *http.Request){
		
		dt:=HtmlData{Ad:r.FormValue("isim"),Yas:r.FormValue("yas")}
		tmp:=template.Must(template.ParseFiles("htmls/mesajoku.html"))
		//fmt.Fprintf(w,"sayin %s %s yaşında",r.FormValue("isim"),r.FormValue("yas"))
		tmp.Execute(w,dt)
		
		
	})





func main(){
	
	mx:=mux.NewRouter()
	mx.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	mx.Handle("/",Welcome).Methods("GET")
	mx.Handle("/sinif/{sinifad}/hoca/{hocaad}",Sinifac)
	mx.Handle("/mesaj",Mesajat)
	mx.Handle("/mesajoku",Mesajal)
	
	
	
	
	
	
	
 http.ListenAndServe("127.0.0.10:9000", mx) // set listen port
	
}
