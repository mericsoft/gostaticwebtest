package main
import (
"fmt"
"net/http"
_"strings"
"log"
"time"
"os"

)



func main() {
	//static files
	fs:=http.FileServer(http.Dir("static/"))
	http.Handle("/static/",http.StripPrefix("/static/",fs))
	http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request){ // set router
		fmt.Fprintf(w,"ilk sayfa, isteğiniz %s \n",r.URL.Path)
		yas:=r.FormValue("yas")
		if yas!=""{
			fmt.Fprintf(w,"yaşınız : %s ",yas)
		}
	})
	
	http.HandleFunc("/okul",func(w http.ResponseWriter,r *http.Request){
		html:="<html> <head> <link rel='stylesheet' href='static/stil.css'/> <script src='static/main.js'></script> </head> <body>selamlar <div>bugün güzel</div> <div id='a'></div> </body> </html>"
		w.Header().Set("121112","token")
		fmt.Fprintf(w,html)
		
	})
	
	http.HandleFunc("/cerezata",func(w http.ResponseWriter, r *http.Request){
		cdeg:=r.FormValue("cerez")
		if cdeg!=""{
			//exp:=time.Now().Add(24*time.Hour) //zaman aşımı
			cook:=http.Cookie{Name:"cereztest",Value:cdeg,MaxAge:60*60*24} // çerezi oluştur
			http.SetCookie(w,&cook) //cerezi ata
		}
	})

	http.HandleFunc("/cerezoku",func(w http.ResponseWriter, r *http.Request){
		getcook,_:=r.Cookie("cereztest")
		fmt.Fprintf(w,"çerez değeri =  %s",getcook)

	})
	
	http.HandleFunc("/cerezsil",func(w http.ResponseWriter, r *http.Request){
		delcook := http.Cookie{Name:"cereztest",Value:"",MaxAge:-1}
		http.SetCookie(w,&delcook)
		http.Redirect(w,r,"/",301)
		//w.Header().Add("Location" ,"http://127.0.0.0:9000")
		//w.WriteHeader(http.StatusOK)


	})
	Port:=os.Getenv("PORT")
	s:=&http.Server{
	  Addr: ":"+port,
    	ReadTimeout: 10 * time.Minute,
    	WriteTimeout: 10 * time.Minute,
    	MaxHeaderBytes: 0,
	}
	err:=s.ListenAndServe()
		
	//err := http.ListenAndServe(":"+os.Getenv("PORT"), nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
