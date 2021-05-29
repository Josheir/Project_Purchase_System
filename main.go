package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

type product struct {
	ProductQuantity int
	ProductName     string
	ProductCatTitle string
	ProductCost     int
}

type Product1 struct {
	
	ProductID       int
	ProductQuantity int
	ProductName     string
	DivID string
	ProductCatTitle string
	ProductCost     int
}

//spit back to last html page
type Product2 struct {
	ID                int
	QuantityAvailable int
}

var ProductList = []Product1{}
var ProductList2 = []Product2{}

//cited
//https://www.bing.com/videos/search?q=youtbe+golang+template&refig=e742578f4d004a2b8a5bd1f28849eb0f&ru=%2fsearch%3fq%3dyoutbe%2bgolang%2btemplate%26form%3dANNTH1%26refig%3de742578f4d004a2b8a5bd1f28849eb0f&view=detail&mmscn=vwrc&mid=BD040005A2743ACB801ABD040005A2743ACB801A&FORM=WRVORC
//http://localhost:8080/golangproj/

var globt *template.Template
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

var tpl *template.Template

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

//////////

//executes back to :  finalpage.html

//var list[]product

//var List1  []product
//var prod []product

func addElement(var1 int, var2 string, var3 string, var4 int) {

	var element product
	element.ProductQuantity = var1
	element.ProductName = var2
	element.ProductCatTitle = var3
	element.ProductCost = var4

	//prod = append(prod, element)

}

//arr.push({ID:key, Quant:item});
//type product1 struct {
//	ID    int `json:"ID"`
//	Quant int `json:"Quant"`
//}
type Product3 struct {
	ID    int
	Quant int
}

// i1 is product id, i2 is quant to add
func updateListForLastpage(index int , quantity int , amtInDatabase int ){

	ProductList2[index].QuantityAvailable = amtInDatabase + quantity
}


//arr.push({ID:key, Quant:item});

//pass an array to here from index.html
//called from finalpage. query string,
//there, checkout button pressed
//this function creates them template2!

//this last page is where the data is spat back to html to note any database changes that cause purchase impossible
func makeListForLastpage(id int, quant int) {

	//to spit back to html
	prod := Product2{

		ID:                id,
		QuantityAvailable: quant,
	}
	//list to spit back to html for rewriting all the quant
	ProductList2 = append(ProductList2, prod)
}

//called from template2, final purchase selected, so send this back to display
func spitBackAmounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")




	fmt.Println(ProductList2)

	json.NewEncoder(w).Encode(ProductList2)
	
	fmt.Println(ProductList2)
	

	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("GET params were:", r.URL.Query())

	query := r.URL.Query()

	
	allIds, present := query["id"]

	if !present || len(allIds) == 0 {
		fmt.Println("filters not present")
	}

	//fmt.Println(allIds[0])

	allQuants, present := query["quant"]

	if !present || len(allQuants) == 0 {
		fmt.Println("filters not present")
	}

	
	db := dbConn()

	var j = 0
	var i = 0
		for i = 0; i < len(allIds); i++ {
		
		stmt, err := db.Prepare("SELECT products.ProductQuantity  " +
			"FROM products WHERE " +
			"products.ProductID = ?")

		if err != nil {
			panic(err.Error())
		}
		rows, err := stmt.Query(allIds[j])
		

		if err != nil {
			panic(err.Error())
		}

		
		var ProductQuantity int

		var flag = "noOther"
		for rows.Next() {
			
			//This is the new exsact quantity available in the database for index i
			//the two arrays, below, are from the attempted values to purchase from setdata (template1)
			//this is a check of the most current database values
			err = rows.Scan(&ProductQuantity)
			if err != nil {
				panic(err.Error())
			}
			//assumption one unique index for whole slice 
			id, err := strconv.Atoi(allIds[i])
			if err == nil {
				fmt.Println(id)
			}
			quant, err1 := strconv.Atoi(allQuants[i])
			if err1 == nil {
				fmt.Println(quant)
			}
			for j = 0; j < len(ProductList2) ; j++{
			//	i3 := id 
				
				//id is already in product list keep id and add quant to it
				if (ProductList2[j].ID == id){
					flag = "thereIsAnother"
					var amtInDatabase = ProductQuantity
					updateListForLastpage(j , quant, amtInDatabase )
				}
			}
			//no current record
			if (flag != "ThereisAnother"){

			//snot in list yet
			makeListForLastpage(j , quant)
			}

		}

		//////////
	}
	if(len(ProductList2) != 0){
	//sends array of structs to template2.html
	json.NewEncoder(w).Encode(ProductList2)
	
	} else{
		fmt.Println("array length zero")
	}
}

func createTemplate2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	
	//////////

	/////////

	fmt.Println("GET params were:", r.URL.Query())

	query := r.URL.Query()

	//filters=["color", "price", "brand"]
	allIds, present := query["id"]

	if !present || len(allIds) == 0 {
		fmt.Println("filters not present")
	}

	//fmt.Println(allIds[0])

	allQuants, present := query["quant"]

	if !present || len(allQuants) == 0 {
		fmt.Println("filters not present")
	}

	/////////////

	//////////
	/////////

	string1 = ""

	db := dbConn()

	var var1 = "D"
	//yes this is right product starts at one
	var j = 0
	//counter := 0
	
	var i = 0
	//divID := "A"
	for i = 0; i < 1; i++ {

		
		DivID  := var1 + (strconv.Itoa(i))
		//pid := allIds[i]
		stmt, err := db.Prepare("SELECT products.ProductQuantity,products.ProductName,products.ProductCatTitle, products.ProductCost  " +
			"FROM products WHERE " +
			"products.ProductID = ?")

		if err != nil {
			panic(err.Error())
		}

		rows, err := stmt.Query(j)

		if err != nil {
			panic(err.Error())
		}

		//var counter = 0

		//var templ1 product

		var ProductQuantity, ProductCost int
		var ProductName, ProductCatTitle string

		for rows.Next() {

			//copies from database row to these variables
			err = rows.Scan(&ProductQuantity, &ProductName, &ProductCatTitle, &ProductCost)
			if err != nil {
				panic(err.Error())
			}

			//
			value1,err1 := strconv.Atoi(allIds[i]) 
			if err1 != nil {
				panic(err1.Error())
			}
			//value1 is product ID
			addProduct(value1, ProductQuantity, ProductName, DivID, ProductCatTitle, ProductCost)
		}
		

		globt = template.Must(template.ParseFiles("C:/wamp64/www/golangproj/template2.html"))

		err1 := globt.Execute(w, ProductList)

		if err1 != nil {
			fmt.Println("CC---------------")
			fmt.Println(err1.Error())

			panic(err1.Error())

		}

	}

	///////////
}
func addProduct(prodid int, quant int, name string, div string, cat string, cost int) {

	prod := Product1{
	
		ProductID:  prodid,
		ProductQuantity: quant,
		ProductName:     name,
		DivID: div,
		ProductCatTitle: cat,
		ProductCost:     cost,
	}
	flag := "nonefound"
	//create one record with total quantities, duplicate record is set to -1 as productID and 0 as quantity
	//could be done better, if time allows
	for i := 0; i<len(ProductList) ; i++{
		if((ProductList[i].ProductID) == prodid) {
		prod.ProductQuantity =  prod.ProductQuantity + ProductList[i].ProductQuantity
		ProductList[i].ProductQuantity = 0
		ProductList[i].ProductID = -1
		ProductList = append(ProductList, prod)
		flag = "found"
		} else{
			
		}
		if (flag != "found"){
		ProductList = append(ProductList, prod)
		}
	}



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

//var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

//this is for testin, not used anympre
func processSearch(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Rrrrrrraarg ")
	fmt.Fprintf(w, "got here1!")

}

type forTemplate struct {
	ProductID       int
	ProductCatTitle string
	//MainDiv         string
	TitleID string
	//ProductFilename    string
	ProductName        string
	DescID             string
	ProductDescription string
	CostID             string
	ProductCost        int
	QuantityID         string
	ProductQuantity    int
	Key1ID             string
	GKeyword1          string
	Key2ID             string
	GKeyword2          string
	Key3ID             string
	GKeyword3          string
	ProductFilename    string
	AmountToPurchaseID string
	AmountPurchasedID  string
	MainDivID          string
}

type Name struct {
	FName string
	LName string
}

type VAR1 struct {
	Var1 string
}

//var templ1 = forTemplate{str3, var18, var2, var3, var4, var5, var6, var7, str4, var9, str2, var11, var12, var13, var14, var15, var16}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var Var1 = "abc"
	var var2 = VAR1{Var1}
	//w.Header().Add("Content-Type", "application/html")

	//w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//data := VAR{"testname"}
	//var var1 = VAR{var2}

	globt := template.Must(template.ParseFiles("C:/wamp64/www/golangproj/twemplate1.html"))

	err1 := globt.Execute(w, var2)

	if err1 != nil {
		fmt.Println("B---------------")
		fmt.Println(err1.Error())

		panic(err1.Error())

	}

	//t, _ := template.ParseFiles("index1.html")
	//var t = template.Must(template.New("").Parse("index1.html"))
	//globt.Execute(w, "a")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	name := Name{"mindorks2", "Subject2"}
	template, _ := template.ParseFiles("index2.html")
	template.Execute(w, name)
}

func updateForm(w http.ResponseWriter, r *http.Request) {
	/*
		var var1 = r.FormValue("productID")
		var var18 = r.FormValue("ProductCatTitle")
		var var2 = r.FormValue("mainDiv")
		var var3 = r.FormValue("titleID")
		var var4 = r.FormValue("ProductName")
		var var5 = r.FormValue("descID")
		var var6 = r.FormValue("ProductDescription")
		var var7 = r.FormValue("costID")
		var var8 = r.FormValue("ProductCost")
		var var9 = r.FormValue("quantityID")
		var var10 = r.FormValue("ProductQuantity")
		var var11 = r.FormValue("key1ID")
		var var12 = r.FormValue("globKeyword")
		var var13 = r.FormValue("key2ID")
		var var14 = r.FormValue("globKeyword")
		var var15 = r.FormValue("key3ID")
		var var16 = r.FormValue("globKeyword")
		var var17 = r.FormValue("amountPurchased")
		var var19 = r.FormValue("ProductFilename")

		//enough product

		str, _ := strconv.Atoi(var17)
		str2, _ := strconv.Atoi(var10)
		str3, _ := strconv.Atoi(var1)
		str4, _ := strconv.Atoi(var8)

		if str <= str2 {

			var templ1 = forTemplate{str3, var18, var2, var3, var4, var5, var6, var7, str4, var9, str2, var11, var12, var13, var14, var15, var16}

			fmt.Println(templ1)

			_ = globt.Execute(w, templ1)

		} else {

			return
		}
	*/
}

/////////
func display1(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	string1 = ""

	fmt.Println("in display 1")

	db := dbConn()

	globKeyword = "apple1"

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

	var counter = 0

	var templ1 forTemplate

	for rows.Next() {

		var ProductCost, ProductQuantity int
		var gKeyword1, gKeyword2, gKeyword3, ProductName, ProductDescription, ProductCatTitle, ProductFilename, AmountToPurchaseID, AmountPurchasedID string

		err = rows.Scan(&gKeyword1, &gKeyword2, &gKeyword3, &ProductName, &ProductID, &ProductDescription, &ProductCost, &ProductQuantity, &ProductCatTitle, &ProductFilename)

		if err != nil {
			panic(err.Error())
		}

		counter = counter + 1
		str := strconv.Itoa(counter)

		//var inputID = "inputID" + str
		var mainDivID = "mainDivID" + str
		var titleID = "titleID" + str
		var descID = "descID" + str
		var costID = "costID" + str
		var quantityID = "quantityID" + str
		var key1ID = "key1ID" + str
		var key2ID = "key2ID" + str
		var key3ID = "key3ID" + str
		AmountToPurchaseID = "amountID" + str
		AmountPurchasedID = "amountPID" + str

		templ1 = forTemplate{ProductID, ProductCatTitle, titleID, ProductName, descID, ProductDescription, costID, ProductCost, quantityID, ProductQuantity,
			key1ID, globKeyword, key2ID, globKeyword, key3ID, globKeyword, ProductFilename, AmountToPurchaseID, AmountPurchasedID, mainDivID}

		fmt.Println(templ1)

		globt := template.Must(template.ParseFiles("C:/wamp64/www/golangproj/template1.html"))

		err1 := globt.Execute(w, templ1)

		if err1 != nil {
			fmt.Println("B---------------")
			fmt.Println(err1.Error())

			panic(err1.Error())

		}

	}

}

func submitfunc(w http.ResponseWriter, r *http.Request) {

	fmt.Println("aarg ")
	fmt.Println("aarg ")
}

//send from client to server and
//send form server to client
func getMessages(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	type User struct {
		Name string 
		Age  int    
		City string 
	}

	
	//	a := User{Name:"a" , Age: 10 , City:"s" };

	var user = []User{{

		Name: "John Doe",
		Age:  10,
		City: "richmond"}}

	var msg = new(User)
	msg.Name = "Test namee"
	msg.Age = 30
	msg.City = "here"
	user = append(user, *msg)

	msg = new(User)
	msg.Name = "namee"
	msg.Age = 20
	msg.City = "here2"
	user = append(user, *msg)

	json.NewEncoder(w).Encode(user)

	//w.Header().Set("Content-Type", "application/json")
	//w.Write(j)
	fmt.Println("--wwww---")
	fmt.Println(user)
}

//////

//type one struct {
//}

//type two struct {
//}

//func (m *one) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Listening on 8080: foo "))
//}
//func (m *two) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Listening on 8081: foo "))
//}
func main() {

	//finish := make(chan bool)

	//wg := new(sync.WaitGroup)
	//wg.Add(2)

	one := http.NewServeMux()

	//mux := http.NewServeMux()

	//has an id value passed in url
	one.HandleFunc("/updateForm/", updateForm)
	one.HandleFunc("/processSearch", processSearch)

	//button3 - just read session for right now
	one.HandleFunc("/getMessages", getMessages)

	one.HandleFunc("/display", display1)

	one.HandleFunc("/HelloWorld", HelloWorld)

	two := http.NewServeMux()

	//
	two.HandleFunc("/template2", createTemplate2)
	one.HandleFunc("/spitBackAmounts", spitBackAmounts)

	go func() {

		http.ListenAndServe(":8080", one)
	}()

	http.ListenAndServe(":8081", two)

	//wg.Wait()
}
