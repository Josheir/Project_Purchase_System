//FOR TEMPLATE:  https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html

type product struct {
	Quantity string
	Title    string
	Category string
	Cost     int
}


w.Header().Set("Access-Control-Allow-Origin", "*")
	name := Name{"mindorks2", "Subject2"}
	template, _ := template.ParseFiles("index2.html")
	template.Execute(w, name)
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

//executes back to :  finalpage.html
func purgeHTML(w http.ResponseWriter, r *http.Request) {

	///////////////

	////////////////
	//incoming : productid and quantity
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.prepare("SELECT products.ProductQuantity, products.ProductName, products.ProductCatTitle, products.ProductCost FROM products WHERE id=1")
	if err != nil {
		panic(err.Error())
	}

	rows, err := stmt.Query()
	if err != nil {
		panic(err.Error())
	}



	prod := product{}
	for selDB.Next() {
		var quant, cost int
		var title, category string
		err = selDB.Scan(&quant, &title, &category, &Cost)
		if err != nil {
			panic(err.Error())
		}
		prod.Quantity = quant
		prod.Title = title
		prod.Category = category
		prod.Cost = cost
	}
	//tmpl.ExecuteTemplate(w, "Show", emp)
	
	//templ1 = product{ProductID, ProductCatTitle, titleID, ProductName, descID, ProductDescription, costID, ProductCost, quantityID, ProductQuantity,
	//	key1ID, globKeyword, key2ID, globKeyword, key3ID, globKeyword, ProductFilename, AmountToPurchaseID, AmountPurchasedID, mainDivID}

	//fmt.Println(templ1)

	globt := template.Must(template.ParseFiles("C:/wamp64/www/golangproj/template2.html"))

	err1 := globt.Execute(w, prod)
	
	
	
	
	defer db.Close()
}

func main() {

	mux := http.NewServeMux()

	//has an id value passed in url
	mux.HandleFunc("/purgeHTML", template2)

	http.ListenAndServe(":8080", mux)
}
