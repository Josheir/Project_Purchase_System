package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

var globKeyword = ""
var Test = 1

var ProductID = 0

type Rectangle struct {
	Length  int
	breadth int
	color   string
}

const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB

var string1 = ""

type App struct {
	Name string
}

type employee struct {
	gKeyword1           string
	gKeyword2           string
	gKeyword3           string
	ProductName         string
	ProductID           int
	ProductdDescription string
	ProductCost         int
	ProductQuantity     int
	ProductCatTitle     string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "ecommerce"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Rrrrrrraarg ")
}

////////

func receiveAjax(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		//data := r.FormValue("post_data")
		r.FormValue("post_data")
		fmt.Println("Receive ajax post data string ")

		w.Header().Add("Content-Type", "application/html")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Write([]byte(string1))

	}

}

////////https://stackoverflow.com/questions/21520244/how-to-simply-send-a-request-parameter-with-jquery-form-submit

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

//this is for testin, not used anympre
func processSearch(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "got here1!")

	err := r.ParseForm()
	if err != nil {
		fmt.Println("error")

	}
	fmt.Fprintf(w, "got here2!")
	fmt.Fprintln(w, "search :", r.Form.Get("search"))

	globKeyword = r.Form.Get("search")
	fmt.Println("here")
	//httpServletRequest.getParameter("myparam")

}

/////////
func display1(w http.ResponseWriter, r *http.Request) {

	/*
		err := r.ParseForm()
		if err != nil {
			fmt.Println("error")

		}

		fmt.Fprintf(w, "got here1!")
		fmt.Fprintln(w, "search :", r.Form.Get("search"))
	*/
	fmt.Println("in display 1")
	//fmt.Fprintf(w, "got here1!")
	//Var1 := "apple1"
	db := dbConn()

	///////////////////////////////////

	//https://dwahyudi.github.io/2020/05/15/mysql-where-in-query-with-golang.html
	//https://stackoverflow.com/questions/59005026/how-to-make-an-sql-query-in-golang-with-multiple-values-in-the-where-clause

	/*
	   	for results.Next() {
	   		var vehicle Vehicle
	   		err = results.Scan(&vehicle.ID, &vehicle.ProductionYear, &vehicle.Brand)
	   		panicError(err)
	   		fmt.Println(vehicle.ProductionYear)
	   		fmt.Println(vehicle.Brand)
	   		fmt.Println("=================")
	   	}
	   stmt, err := db.Prepare("select * from someTable where age = ? and hairColor = ?")
	   rows, err := stmt.Query(age,hairColor)
	*/
	////////////////////////////////////////

	//add products.ProductCatTitle = \"$titleOfSelectedDropDown\"
	//var q = "SELECT products.ProductKeyword1, products.ProductKeyword2, products.ProductKeyword3, products.ProductName, products.ProductID, " +
	//	"products.ProductDescription, products.ProductCost, products.ProductQuantity, products.ProductCatTitle , products.ProductFilename " +
	//	"FROM products WHERE " +
	//	"((products.ProductKeyWord1 = " + Var1 + ") OR " +
	//	"(products.ProductKeyWord2 = " + Var1 + ") OR (products.ProductKeyWord3 = " + Var1 + " ))"
	//add products.ProductCatTitle = \"$titleOfSelectedDropDown\"

	//var q = "SELECT products.ProductKeyword1, products.ProductKeyword2, products.ProductKeyword3, products.ProductName, products.ProductID, " +
	//	"products.ProductDescription, products.ProductCost, products.ProductQuantity, products.ProductCatTitle, products.ProductFilename " +
	//	//	"FROM products WHERE " +
	//	"FROM products WHERE " +
	//	"((products.ProductKeyWord1 = \"apple1\") OR " +
	//	"(products.ProductKeyWord2 = \"apple1\") OR (products.ProductKeyWord3 = \"apple1\" ))"

	stmt, err := db.Prepare("SELECT products.ProductKeyword1, products.ProductKeyword2, products.ProductKeyword3, products.ProductName, products.ProductID, " +
		"products.ProductDescription, products.ProductCost, products.ProductQuantity, products.ProductCatTitle , products.ProductFilename " +
		"FROM products WHERE " +
		"((products.ProductKeyWord1 = ?) OR " +
		"(products.ProductKeyWord2 = ?) OR (products.ProductKeyWord3 = ? ))")
	if err != nil {
		panic(err.Error())
	}

	rows, err := stmt.Query(globKeyword, globKeyword, globKeyword)

	if err != nil {
		panic(err.Error())
	}

	counter := 0

	for rows.Next() {

		var ProductCost, ProductQuantity int
		var gKeyword1, gKeyword2, gKeyword3, ProductName, ProductDescription, ProductCatTitle, ProductFilename string

		err = rows.Scan(&gKeyword1, &gKeyword2, &gKeyword3, &ProductName, &ProductID, &ProductDescription, &ProductCost, &ProductQuantity, &ProductCatTitle, &ProductFilename)

		if err != nil {
			panic(err.Error())
		}

		///////////////////
		/*string1 = string1 +

		"<iframe id=\"upload_target\" name=\"upload_target\"  style=\"width:0;height:0;border:0px solid #fff;\"></iframe>" +

		"<form " +
		"target = \"upload_target\" " +
		"id=\"form\"" +
		"enctype=\"multipart/form-data\"" +
		"action=\"http://localhost:8080/upload\"" +
		"method=\"POST\">" +
		"<input type=\"hidden\" id=\"custId\" name=\"custId\" value=\"3487\">" +
		"<input class=\"file\" id = \"file\" type=\"file\" name=\"file\" multiple />" +
		"<button class=\"button\" type=\"submit\">Submit for upload</button>" +
		"</form>" +

		"<p style=\"color:Tomato;\" ><b>Images can not exceed 50 megabytes</p>" +
		"<input onclick = \"refresh(   )\" id='button2' type='button' value='Confirm Image'>" +
		"<br><br><br>"
		*/

		//////////////////

		///////////////////
		counter = counter + 1
		str := strconv.Itoa(counter)

		var titleID = "titleID" + str
		var descID = "descID" + str
		var costID = "costID" + str
		var quantityID = "quantityID" + str
		var key1ID = "key1ID" + str
		var key2ID = "key2ID" + str
		var key3ID = "key3ID" + str

		string1 = string1 + "<p id = \"link1\">product id   : " + strconv.Itoa(ProductID) + " </p>" +
			"<p>category id  : " + ProductCatTitle + "</p>" +

			"<div class=\"row\" >" +
			"<div class=\"col\">" +
			"<h4><center><p id = \"\">Image</p></center></h4>" +
			//"<center><p id = \"\"> image </p></center>" +

			"<img  height=\"100\" width=\"100\"  src= /golangproj/uploads/" + ProductFilename + " alt=\"no product image\">" +

			"</div>" +

			"<div class=\"col\">" +
			"<h4><center><p id =\"\"  >Title</p></center></h4>" +

			"<center>      <p  >      <input id = " + titleID + " value = " + ProductName + " type=\"text\" name=\"title\" placeholder=\"\"></p></center>" +
			"</div>" +

			"<div class=\"col\">" +

			"<h4><center><p id = \"\">Desc</p></center></h4>" +

			"<center><textarea wrap id = " + descID + "   value = " + ProductDescription + "  type=\"text\" rows=\"5\" cols=\"34\"></textarea></center>" +
			"</div>" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\" >Cost</p></center></h4>" +
			"<center><p>	<input id = " + costID + " value =  " + strconv.Itoa(ProductCost) + "   type=\"number\" name=\"title\" placeholder=\"\">		</p></center>" +

			"</div>" +

			"</div>" +
			"</div>" +

			"<div class=\"row\" >" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\" >Quantity</p></center></h4>" +
			"<center><p> <input id = " + quantityID + " value = " + strconv.Itoa(ProductQuantity) + "  type=\"number\" name=\"title\" placeholder=\"\">	</p></center>" +
			"</div>" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\" >Keyword 1</p></center></h4>" +
			"<center><p>	<input id = " + key1ID + " value = " + gKeyword1 + " type=\"text\" name=\"title\" placeholder=\"\">		</p></center>" +

			"</div>" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\" >Keyword 2</p></center></h4>" +
			"<center><p>	<input id = " + key2ID + " value = " + gKeyword2 + " type=\"text\" name=\"title\" placeholder=\"\">		</p></center>" +

			"</div>" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\" >Keyword 3</p></center></h4>" +
			"<center><p>	<input id =   " + key3ID + "  value = " + gKeyword3 + " type=\"text\" name=\"title\" placeholder=\"\">		</p></center>" +

			"</div>" +

			"<br><br>" +

			"<div class=\"row\" >" +

			"<div class=\"col\">" +

			//<!--THIS LINE IS THE MAIN PROBLEM, THE FIRST PARAMETER WORKS BUT THE IDS HAVE A VALUE OF HTMLOBJECT
			//deleteFlag + "); , mainDiv , strconv.Itoa(ProductID) , titleID , descID,costID,quantityID ,key1ID,key2ID,key3ID
			"<center><button id = \"\" onclick = \"SaveProductItems(" + strconv.Itoa(ProductID) + ", " + titleID + ", " + descID + " )\">Submit</button></center>" +

			//<!--flag for determining if record delete will effect" sessioncount, 0 is no.-->
			//"<center><button id = \"\" onclick = \"deleteRecord( 1,  mainDiv, strconv.Itoa(ProductID) )\">Delete</button></center>" +
			//"<p><a href=\"#add\">To Add</a></p>" +
			//-->
			"</div>" +
			"<br><br>" +

			"</div>" +

			" <br><br><br><br>"

	} //for selDB.Next()

	receiveAjax(w, r)
}

///////
func uploadHandler(w http.ResponseWriter, r *http.Request) {

	//r.ParseMultipartForm()
	//fmt.Fprint(w, "aaaaaaa")
	//fmt.Println("upload handler1 ")

	//fmt.Fprint(w, ProductID)
	//***db := dbConn()

	////////////
	//	ProdID, ok := r.URL.Query()["ProductID"]
	//
	//	if !ok || len(ProdID[0]) < 1 {
	//		log.Println("Url Param 'key' is missing")
	//		return
	//	}
	//
	//	var productID = ProdID[0]

	//***var productID = 100

	//***var productID2 string
	//***productID2 = strconv.Itoa(productID)

	////////////

	/////////////////////////////

	////////////

	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.

	session, _ := store.Get(r, "session-name")
	// Set some session values.
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	// Save it before we write to the response/return from the handler.
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var a = session.Values["foo"]

	fmt.Println(a)

	//////////////
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
		return
	}
	////////////
	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	//////////
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {
		http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
		return
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/////////

	////////
	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//image is saved, now save to database
	fmt.Println("Received data string ")

	//////////////////////

	/////////

	////////////

	db := dbConn()

	//////////////////////////////////////////////////////////
	//DATABASE BELOW HERE, ALREADY SAVED FILE
	//////////////////////////////////////////////////////////

	var productFilename string

	var productFilename1 string

	var q0 = "SELECT ProductFilename FROM products WHERE ProductID = " + strconv.Itoa(ProductID)
	selDB, err := db.Query(q0)
	if err != nil {
		panic(err.Error())
	}
	//gets last number
	//for selDB.Next() {

	for selDB.Next() {
		_ = selDB.Scan(&productFilename)
	}

	productFilename1 = strings.ReplaceAll(productFilename, "image/", "")

	var path = "/uploads/" + productFilename1
	//err := os.Remove(path)
	os.Remove(path)

	//get latest number
	var Number int
	var q3 = "SELECT Number FROM  numbers"
	selDB1, err := db.Query(q3)
	if err != nil {
		panic(err.Error())
	}

	for selDB1.Next() {
		//gets last number
		_ = selDB1.Scan(&Number)
	}
	//var value1 = strconv.Itoa(Number)

	var value2 = strconv.Itoa(Number + 1)

	var Filename = "A" + (value2) + "." + filetype
	productFilename1 = strings.ReplaceAll(Filename, "image/", "")

	//put in database
	//stmt, e := db.Prepare("UPDATE products SET  ProductFilename = '" + Filename + "' WHERE ProductID = " +  strconv.Itoa(ProductID)
	stmt, e := db.Prepare("UPDATE products SET  ProductFilename = ? WHERE ProductID = ?")
	//selDB2, err := db.Query(q1)
	if e != nil {
		panic(err.Error())
	}
	stmt.Exec(productFilename1, ProductID)

	stmt, e = db.Prepare("UPDATE numbers SET Number =  ?")

	if e != nil {
		panic(err.Error())
	}

	stmt.Exec(value2)

}

func submitfunc(w http.ResponseWriter, r *http.Request) {

	fmt.Println("aarg ")
	fmt.Println("aarg ")
}

/////////

func getMessages(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//fmt.Println("method:", r.Method)

	////

	fmt.Println("in get messages ")
	fmt.Fprint(w, "or this")

	/*
		a := make([]string, 2)
		a[0] = "John"
		a[1] = "Sam"
		j, err := json.Marshal(a)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}

		w.Write(j)
	*/
}

//////

func main() {

	mux := http.NewServeMux()

	//button1

	//button2 - just make session right now

	mux.HandleFunc("/processSearch", processSearch)

	mux.HandleFunc("/upload", uploadHandler)

	//button3 - just read session for right now
	mux.HandleFunc("/getMessages", getMessages)
	//not used
	mux.HandleFunc("/display", display1)

	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
