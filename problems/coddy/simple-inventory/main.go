package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Product struct{
	Price float64
	Quantity int
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	scanner.Scan()
	storeInfo := scanner.Text()

	storeArr:=strings.Split(storeInfo,",")
	storeName:=storeArr[0]
	location:=storeArr[1]
	
	scanner.Scan()
	productData := scanner.Text()
	
	productArr:=strings.Split(productData,",")

	inventory := make(map[string]Product)
	var allProductsName []string
	totalStocks:=0
	totalValue:=0.0
	
	for _,product:=range productArr{
		productInfo := strings.Split(product,":")
		productPrice,_ := strconv.ParseFloat(productInfo[1],64)
		productQuantity,_:=strconv.Atoi(productInfo[2])
		inventory[productInfo[0]]=Product{productPrice,productQuantity}
		allProductsName= append(allProductsName, productInfo[0])
		totalStocks+=productQuantity
		totalValue+= float64(productQuantity)*productPrice
	}

	fmt.Printf("=== %s Inventory System ===\n",storeName)
	fmt.Printf("Location: %s\n",location)
	fmt.Printf("Inventory initialized with %d products\n",len(productArr))

	sort.Strings(allProductsName)
	fmt.Println("Current Inventory:")
	for _,product:=range allProductsName{
		fmt.Printf("- %s: $%.2f (Stock: %d)\n",product,inventory[product].Price,inventory[product].Quantity)
	}

	fmt.Println("Inventory Statistics:")
	fmt.Printf("Total Products: %d\n",len(allProductsName))
	fmt.Printf("Total Items in Stock: %d\n",totalStocks)
	fmt.Printf("Total Inventory Value: $%.2f\n",totalValue)

	fmt.Println("System Status: Ready")
	fmt.Println("Inventory management system initialized successfully")
	
}