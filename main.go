package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	//"log"
	"math"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	//"math/big"
	//"github.com/shopspring/decimal"
	//"github.com/leekchan/accounting"
)

type product struct {
	ProductQuantity int
	ProductName     string
	ProductCatTitle string
	ProductCost     int
}

type Product1 struct {
	GrandTotalStringID string

	GrandTotalString string
	BoughtID         string
	Bought           int
	TotalCost        string
	TotalCostID      string
	CostID           string
	AmountToBuyID    string
	Condition        int
	Condition2       int
	ProductID        int
	ProductQuantity  int
	ProductName      string
	DivID            string
	ProductCatTitle  string
	ProductCost      string
}

//spit back to last html page
type Product2 struct {
	ID                  int
	QuantityAvailable   int
	IsNotEnoughQuantity string
}

var ProductList = []Product1{}
var ProductList2 = []Product2{}
var ProductList2A = []Product2{}

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

// USD represents US dollar amount in terms of cents
type USD int64

// Float64 converts a USD to float64
func (m USD) Float64() float64 {
	x := float64(m)
	x = x / 100
	return x
}

// Multiply safely multiplies a USD value by a float64, rounding
// to the nearest cent.
func (m USD) Multiply(f float64) USD {
	x := (float64(m) * f) + 0.5
	return USD(x)
}

// String returns a formatted USD value
func (m USD) String() string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("$%.2f", x)
}

////////////
// ToUSD converts a float64 to USD
// e.g. 1.23 to $1.23, 1.345 to $1.35
func ToUSD(f float64) USD {
	return USD((f * 100) + 0.5)
}

//////////

//executes back to :  finalpage.html

//var list[]product

//var List1  []product
//var prod []product

//https://www.bing.com/search?q=receiver%20int%20golang&qs=n&form=QBRE&sp=-1&pq=receiver%20int%20golang&sc=0-19&sk=&cvid=14C3226BD73C46F09A57AA46291441EA
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
//func updateListForLastpage(index int, quantity int, amtInDatabase int) {
//
//	ProductList2[index].QuantityAvailable = amtInDatabase + quantity
//}

func makeListForLastpageA(enough string, id int, quant int) {

	//to spit back to html
	prod := Product2{

		IsNotEnoughQuantity: enough,
		QuantityAvailable:   quant,
		ID:                  id,
	}
	//list to spit back to html for rewriting all the quant
	ProductList2A = append(ProductList2A, prod)
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

		QuantityAvailable: quant,
		ID:                id,
	}
	//list to spit back to html for rewriting all the quant
	ProductList2 = append(ProductList2, prod)
}

var orderid1 = 100

/*
//for testing
func trycommit(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	db := dbConn()
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}

	var cust = 1
	var cost = 107
	var date = "2021-01-1"

	//var ins *sql.Stmt
	//stmt, err := tx.Prepare("INSERT INTO orders set OrderID=?, OrderDate=?,OrderCost=?, CustomerID=?");
	stmt, err := tx.Prepare("INSERT INTO orders (OrderID, OrderDate,OrderCost, CustomerID) values(?,?,?,?)")

	if err != nil {
		fmt.Println(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(orderid1, date, cost, cust)
	orderid1++

	if err != nil {
		fmt.Println(err)
	}

	err5 := tx.Commit()
	if err5 != nil {
		fmt.Println(err5)
	}

}
*/
//called from template2, final purchase selected, so send this back to display
func spitBackAmounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	//fmt.Println(ProductList2)

	//json.NewEncoder(w).Encode(ProductList2)

	//fmt.Println(ProductList2)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	//fmt.Println("GET params were:", r.URL.Query())

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

	var isEnoughInDatabase = "yes"

	i := 0
	j := 0
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}

	var counter = 0
	var quant = 0
	//any fail means isnotenoughindatabase
	for j = 0; j < len(allIds); j++ {

		DatabaseQuantity := 0

		//gets quantity for each product id
		quant, _ = (strconv.Atoi(allQuants[j]))

		if quant == 0 {
			continue
		}

		/////////IMPORTANT!!!!!!/////
		if counter == 0 {
			//quant = 2000
			counter++
		}

		row := tx.QueryRow("SELECT products.ProductQuantity FROM products WHERE products.ProductID = ?", allIds[j])
		err = row.Scan(&DatabaseQuantity)
		if err != nil {
			fmt.Println(err)
		}

		val1, err1 := strconv.Atoi(allIds[j])
		if err1 != nil {
			fmt.Println(err)
		}

		//listing : ProductList2A - new values of quantity for id
		//allIds2[i] = val1
		//product amt, whats in database
		//allQuants2[i] = DatabaseQuantity

		enough := ""
		//not enough to buy this product
		if DatabaseQuantity <= quant {

			enough = "yes"
		} else {

			enough = "no"
		}

		//makes productline2A for new values to pass to template2.html
		//appends
		makeListForLastpageA(enough, (val1), DatabaseQuantity)

		//amount of product changed, there is no longer enough product for purchase
		//just send all values without delete
		//if there is no fail than checkout completes with
		//quant is amount purchasing
		//if counter == 0 {
		//	quant = 2000
		//	counter++
		//}
		//any one fail means just write the new amounts and do  not change the database
		if (quant) > DatabaseQuantity || isEnoughInDatabase == "no" {

			isEnoughInDatabase = "no"
			//continue

		}

		//err := tx.Commit();
		//if err != nil {
		//fmt.Println(err)
		//}

	}

	//all the products may be removed from the database so delete  them and create an order and create a new product for each product
	//proceed with checkout
	if isEnoughInDatabase == "yes" {

		//set new record for order, with the quantity subtracted

		var insertOrderFlag = "yes"
		var orderid = 0
		var nextProductID = 0
		//var ins *sql.Stmt
		for i = 0; i < len(allIds); i++ {
			insertOrderFlag = "yes"

			intQuant, err := strconv.Atoi(allQuants[i])
			if err != nil {
				fmt.Println(err)
			}

			if intQuant <= 0 {
				continue
			}

			//quant, err2 := (strconv.Atoi(allQuants[i]))
			//if err2 != nil {
			//	fmt.Println(err)
			//}

			orderid = 0
			//get last order id
			row := tx.QueryRow("SELECT OrderID FROM orders ORDER BY OrderID DESC LIMIT 1")
			err3 := row.Scan(&orderid)
			if err3 != nil {
				fmt.Println(err)
			}
			orderid = orderid + 1
			//get last productid
			row = tx.QueryRow("SELECT ProductID FROM products ORDER BY ProductID DESC LIMIT 1")
			err3 = row.Scan(&nextProductID)
			if err3 != nil {
				fmt.Println(err)
			}

			nextProductID = nextProductID + 1

			//////////////

			//this gets the record before it is updated for quantity and orderid to supply the insert

			rows := tx.QueryRow("SELECT * FROM products   WHERE products.ProductID = ?", allIds[i])

			var ProductCost float64
			var ProductQuantity, ProductID, CustomerID, OrderID, AdminID int
			var gKeyword1, gKeyword2, gKeyword3, ProductName, ProductDescription, ProductCatTitle, ProductFilename, ProductStatus string
			//https://devtidbits.com/2020/08/03/go-sql-error-converting-null-to-string-is-unsupported/
			err4 := rows.Scan(&ProductID, &ProductFilename, &ProductName, &ProductDescription, &ProductCost, &ProductQuantity, &ProductCatTitle,
				&gKeyword1, &gKeyword2, &gKeyword3, &CustomerID, &OrderID, &ProductStatus, &AdminID)

			if err4 != nil {
				fmt.Println(err4)
			}

			//var ProductQuantity2 = ProductQuantity - quant

			/////////////

			//creates same product with quantitiy of this product minus how much purchased (quant)
			//database amount - amount purchased
			//100 in database , 10 bought -> 90 left   so there is ten purchased in order (bought)
			//intQuant, err := strconv.Atoi(allQuants[i])
			_ = tx.QueryRow("Update products SET ProductQuantity = ? WHERE products.ProductID = ?", ProductQuantity-intQuant, allIds[i])

			//////////////
			//https://idineshkrishnan.com/crud-operations-with-mysql-in-go-language/
			//new record with next orderid - amount of products purchased
			//	_ = tx.QueryRow("INSERT INTO products VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)", orderid, ProductFilename,ProductName, ProductDescription,
			//	ProductCost, quant, ProductCatTitle, &gKeyword1, &gKeyword2, &gKeyword3, CustomerID, 0, ProductStatus,AdminID )

			////////////////

			//get this with logon : custid

			if insertOrderFlag == "yes" {
				var cust = 1
				var cost = 100
				var date = "2021-01-1"

				//stmt, err := tx.Prepare("INSERT INTO orders set OrderID=?, OrderDate=?,OrderCost=?, CustomerID=?");
				stmt, err := tx.Prepare("INSERT INTO orders (OrderID, OrderDate,OrderCost, CustomerID) values(?,?,?,?)")

				if err != nil {
					fmt.Println(err)
				}

				defer stmt.Close()

				_, err = stmt.Exec(orderid, date, cost, cust)

				if err != nil {
					fmt.Println(err)
				}

				insertOrderFlag = "no"

			}
			////////////////

			stmt, err := tx.Prepare(`INSERT INTO products set  ProductID=?,ProductFilename=?,ProductName =?, ProductDescription=?, ProductCost =?, ProductQuantity=?, ProductCatTitle = ?,ProductKeyword1=?, ProductKeyword2=?, ProductKeyword3=?, CustomerID = ?,OrderID=?, ProductStatus=?, AdminID=?`)

			if err != nil {
				fmt.Println(err)
			}

			ProductCatTitle = "purchased"

			_, err = stmt.Exec(nextProductID, ProductFilename, ProductName, ProductDescription, ProductCost, intQuant, ProductCatTitle,
				gKeyword1, gKeyword2, gKeyword3, CustomerID, orderid, ProductStatus, AdminID)

			if err != nil {
				fmt.Println(err)
			}

			/////////////////////

		} //for

	} //is enough

	err5 := tx.Commit()
	if err5 != nil {
		fmt.Println(err5)
	}

	if len(ProductList2A) != 0 && isEnoughInDatabase == "no" {
		//sends array of structs to template2.html
		json.NewEncoder(w).Encode(ProductList2A)

	} else {

		//fmt.Println("array length zero")
		http.Redirect(w, r, "http://www.localhost/donePurchase", 301)
	}

}

//https://github.com/strongo/decimal
//https://programming.guide/go/convert-int64-to-string.html

//stackoverflow.com/questions/54362751/how-can-i-truncate-float64-number-to-a-particular-precision
//stackoverflow.com/questions/4187146/truncate-number-to-two-decimal-places-without-rounding#:~:text=General%20solution%20to%20truncate%20%28no%20rounding%29%20a%20number,with%20exactly%20n%20decimal%20digits%2C%20for%20any%20n%E2%89%A50.

func createTemplate2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("GET params were:", r.URL.Query())

	query := r.URL.Query()

	//filters=["color", "price", "brand"]
	allIds, present := query["id"]
	fmt.Println("allIds")
	fmt.Println(allIds)

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
	var var2 = "A"
	var var3 = "C"
	var var4 = "TC"
	var var5 = "B"
	var var6 = "GT"
	//yes this is right product starts at one

	//var j = 1
	//var ProductID = 2
	var i = 0

	var Condition = 0
	var Condition2 = 0

	//error when making two product and pressing checkout displays all table data below main table with another checkout button
	//6/26/21

	ID := 0
	bought := 0
	numTotal := 0.0

	for i = 0; i < len(allIds); i++ {

		fmt.Println("length")
		fmt.Println(len(allIds))

		//Condition++
		Condition2++
		DivID := var1 + (strconv.Itoa(i))
		AmountToBuyID := var2 + (strconv.Itoa(i))
		CostID := var3 + (strconv.Itoa(i))
		TotalCostID := var4 + (strconv.Itoa(i))
		BoughtID := var5 + (strconv.Itoa(i))
		GrandTotalStringID := var6 + (strconv.Itoa(i))

		var prodid, err = (strconv.Atoi(allIds[i]))
		if err != nil {
			fmt.Println(err)
		}

		//Total := ""
		stmt, err := db.Prepare("SELECT products.ProductQuantity,products.ProductName,products.ProductCatTitle, products.ProductCost  " +
			"FROM products WHERE " +
			"products.ProductID = ? AND products.ProductStatus = 'ready'")

		if err != nil {
			fmt.Println(err)
		}

		rows, err := stmt.Query(prodid)

		if err != nil {
			fmt.Println(err)
		}

		//var counter = 0

		//var templ1 product

		var ProductQuantity int

		var ProductName, ProductCatTitle, ProductCost, TotalCost string

		fmt.Println("ProductList")
		fmt.Println(ProductList)

		for rows.Next() {

			fmt.Println("ProductList1")
			fmt.Println(ProductList)

			//copies from database row to these variables
			err = rows.Scan(&ProductQuantity, &ProductName, &ProductCatTitle, &ProductCost)
			if err != nil {
				fmt.Println(err)
			}

			var i = 0
			for i = 0; i < len(allIds); i++ {

				bought, err = (strconv.Atoi(allQuants[i]))
				if err != nil {
					fmt.Println(var2)
				}

				ID, err = strconv.Atoi(allIds[i])
				if err != nil {
					fmt.Println(err)
				}

				//there is boutght total that goes with this id
				if prodid == ID {

					//subtract bought from quantity
					ProductQuantity = ProductQuantity - bought
					break
				}

			} //for

			//value1 is product ID
			if i == 0 {
				Condition = 1
			} else {
				Condition = 0

			}
			if i == (len(allIds) - 1) {
				Condition2 = -1
			}

			aQuant, err3 := strconv.Atoi(allQuants[i])
			if err3 == nil {
				fmt.Println(var2)
			}

			//https://yourbasic.org/golang/round-float-2-decimal-places/
			//https://stackoverflow.com/questions/20596428/how-to-represent-currency-in-go
			//https://www.bing.com/search?q=put%20commas%20in%20string%20golang&qs=n&form=QBRE&sp=-1&pq=put%20commas%20in%20string%20golang&sc=0-27&sk=&cvid=D3A2A7E4E0E141BCAA5BA7E7EE279532
			//quantity
			//However, whole numner
			var QuantityFloat float64 = float64(aQuant)
			//in cents, no decimal
			ProductCostString := ProductCost

			//cents
			ProductCostFloat, err := strconv.ParseFloat(ProductCostString, 64)
			if err != nil {
				fmt.Println(err)
			}
			//cents
			ProductCostFloat2 := ProductCostFloat

			//move decimal
			ProductCostFloat = ProductCostFloat / (math.Pow(10, 2))
			//take care of decimals - ready for display
			ProductCostString = fmt.Sprintf("%.2f", ProductCostFloat)

			//in cents and wholenumber
			TotalCostFloat := QuantityFloat * ProductCostFloat2

			//cents
			TotalCostFloat2 := TotalCostFloat
			//move decimal
			TotalCostFloat = TotalCostFloat / (math.Pow(10, 2))
			//take care of decimals
			TotalCost = fmt.Sprintf("%.2f", TotalCostFloat)

			//cents
			numTotal = numTotal + TotalCostFloat2

			tax := 0.0

			//numTotal = numTotal  * .05
			//this is the tax  amount
			var GrandTotalString = "this text doesnt display"
			if i == (len(allIds) - 1) {
				tax = numTotal * 5
				numTotal = numTotal * 100
				numTotal = numTotal + tax

				numTotal = numTotal / math.Pow(10, 4)
				GrandTotalString = fmt.Sprintf("%.2f", numTotal)
			}

			addProduct(GrandTotalStringID, GrandTotalString, BoughtID, bought, TotalCost, TotalCostID, ProductQuantity, CostID, AmountToBuyID, Condition, Condition2, ID, ProductQuantity, ProductName, DivID, ProductCatTitle, ProductCostString)

		}
		fmt.Println("ProductList")
		fmt.Println(ProductList)

	} //for next loop

	///////////

	//https://stackoverflow.com/questions/24755509/using-conditions-inside-templates
	globt = template.Must(template.ParseFiles("C:/wamp64/www/golangproj/template2.html"))
	fmt.Println("ProductList2")
	fmt.Println(ProductList)

	err1 := globt.Execute(w, ProductList)

	if err1 != nil {
		fmt.Println("CC---------------")
		fmt.Println(err1.Error())

		panic(err1.Error())

	}

	///////////
}
func addProduct(totalID string, total string, boughtid string, bought int, totalcost string, totalcostid string, ProductQuantity int, costid string, amountid string, condition int, condition2, prodid int, quant int, name string, div string, cat string, cost string) {

	prod := Product1{
		GrandTotalStringID: totalID,
		GrandTotalString:   total,
		BoughtID:           boughtid,
		Bought:             bought,
		TotalCost:          totalcost,
		TotalCostID:        totalcostid,
		CostID:             costid,
		AmountToBuyID:      amountid,
		Condition:          condition,
		Condition2:         condition2,
		ProductID:          prodid,
		ProductQuantity:    quant,
		ProductName:        name,
		DivID:              div,
		ProductCatTitle:    cat,
		ProductCost:        cost,
	}
	flag := "nonefound"

	//ProductList = append(ProductList, prod)

	//could be done better, if time allows
	for i := 0; i < len(ProductList); i++ {
		if (ProductList[i].ProductID) == prodid {
			//ProductQuantityFloat = float64(ProductQuantity)
			//prod.Bought = prod.Bought + bought
			ProductList[i].ProductQuantity = ProductQuantity
			ProductList[i].Bought = bought
			ProductList[i].TotalCost = totalcost
			//break out

			flag = "found"
			i = 100
		}
	}

	if flag != "found" {
		//prod.ProductQuantity = prod.ProductQuantity - prod.Bought
		ProductList = append(ProductList, prod)
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
	AmountPurchased int
	ProductID       int
	ProductCatTitle string
	//MainDiv         string
	TitleID string
	//ProductFilename    string
	ProductName        string
	DescID             string
	ProductDescription string
	CostID             string
	ProductCost        float64
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

	query := r.URL.Query()

	key, present := query["var"]

	if !present || len(key) == 0 {
		fmt.Println("filters not present")
	}

	keyTotalAmountBought, present2 := query["quant"]
	if !present2 || len(keyTotalAmountBought) == 0 {
		fmt.Println("filters not present")
	}
	ProdID, present3 := query["id"]
	if !present3 || len(ProdID) == 0 {
		fmt.Println("filters not present")
	}

	globKeyword = key[0]

	w.Header().Set("Access-Control-Allow-Origin", "*")

	string1 = ""

	fmt.Println("in display 1")

	db := dbConn()

	stmt, err := db.Prepare("SELECT products.ProductKeyword1, products.ProductKeyword2, products.ProductKeyword3, products.ProductName, products.ProductID, " +
		"products.ProductDescription, products.ProductCost, products.ProductQuantity, products.ProductCatTitle , products.ProductFilename " +
		"FROM products WHERE " +
		"((products.ProductKeyWord1 = ?) OR " +
		"(products.ProductKeyWord2 = ?) OR (products.ProductKeyWord3 = ? )) AND products.ProductStatus = 'ready'")
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

		var ProductCost float64
		var ProductQuantity int
		var gKeyword1, gKeyword2, gKeyword3, ProductName, ProductDescription, ProductCatTitle, ProductFilename, AmountToPurchaseID, AmountPurchasedID string

		err = rows.Scan(&gKeyword1, &gKeyword2, &gKeyword3, &ProductName, &ProductID, &ProductDescription, &ProductCost, &ProductQuantity, &ProductCatTitle, &ProductFilename)

		if err != nil {
			panic(err.Error())
		}

		i := 0
		prodBoughtInt := 0
		isAmountPurchased := "no"

		for i = 0; i < len(ProdID); i++ {
			prodIDStr := ProdID[i]

			prodIDInt, err := strconv.Atoi(prodIDStr)
			if err != nil {
			}

			prodBoughtStr := keyTotalAmountBought[i]
			prodBoughtInt, err = strconv.Atoi(prodBoughtStr)
			if err != nil {
			}

			if prodIDInt == ProductID {
				ProductQuantity = ProductQuantity - prodBoughtInt
				isAmountPurchased = "yes"
				break
			}

		}

		counter = counter + 1
		str := strconv.Itoa(counter)
		AmountPurchased := 0

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

		if isAmountPurchased == "yes" {
			AmountPurchased = prodBoughtInt
		} else {
			AmountPurchased = 0
		}

		templ1 = forTemplate{AmountPurchased, ProductID, ProductCatTitle, titleID, ProductName, descID, ProductDescription, costID, ProductCost, quantityID, ProductQuantity,
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
