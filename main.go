package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "encoding/json"
    "io/ioutil"
    "strconv"

    "github.com/gorilla/mux"

    "currency-converter/database"
    "currency-converter/tradermade"
)

type ConvertData struct {
    Date time.Time `json:"date"`
    From string `json:"from"`
    To string `json:"to"`
    Amount json.Number `json:"amount"`
}

type Transaction struct {
    Id int `json:"id"`
    Date time.Time `json:"date"`
    FromCurrencyCode string `json:"from_currency_code"`
    ToCurrencyCode string `json:"to_currency_code"`
    Amount float64 `json:"amount"`
    ConversionRate float64 `json:"conversion_rate"`
    ConvertedAmount float64 `json:"converted_amount"`
}


func convertCurrency(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: convertCurrency")

    var data ConvertData
    reqBody, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(reqBody, &data)

    var amount, convertedAmount float64
    amount, _ = strconv.ParseFloat(fmt.Sprint(data.Amount), 64)

    // get the current live conversion rate from TraderMade API
    rate := TraderMade.GetCurrencyRate(data.From, data.To)
    convertedAmount = amount * rate

    db, err := DatabaseManager.OpenDatabase()

    sqlStatement := `
        INSERT INTO hasurapg.transactions
            (date, from_currency_code, to_currency_code, amount, conversion_rate, converted_amount)
        VALUES
            ($1, $2, $3, $4, $5, $6)`
    _, err = db.Exec(sqlStatement, data.Date, data.From, data.To, amount, rate, convertedAmount)
    if err != nil {
        panic(err)
    }

    DatabaseManager.CloseDatabase(db)

    json.NewEncoder(w).Encode(data)
}

func returnAllTransactions(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllTransactions")

    db, err := DatabaseManager.OpenDatabase()

    rows, err := db.Query(
        `SELECT
            id, date, from_currency_code, to_currency_code, amount, conversion_rate, converted_amount
        FROM hasurapg.transactions`)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    var transactions []Transaction

    for rows.Next() {
        var transaction Transaction
        err := rows.Scan(
            &transaction.Id, &transaction.Date, &transaction.FromCurrencyCode,
            &transaction.ToCurrencyCode, &transaction.Amount,
            &transaction.ConversionRate, &transaction.ConvertedAmount)
        transactions = append(transactions, transaction)

        if err != nil {
            panic(err)
        }
    }

    DatabaseManager.CloseDatabase(db)

    json.NewEncoder(w).Encode(transactions)
}

func handleRequests() {
    // create a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/transactions", returnAllTransactions)
    myRouter.HandleFunc("/convert", convertCurrency).Methods("POST")

    http.Handle("/", myRouter)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    fmt.Println("GoLang API - Mux Router - Hasura - PostgreSQL " +
        "- Converting Currency - Live Rates - Saving Each Transaction")
    handleRequests()
}
