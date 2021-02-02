package models

package models

import (
	"fmt"
	"time"
)

type Status_Menu struct {
	Active_Dashboard string
	Active_Customers string
	Active_Products  string
	Active_Orders    string
	Active_Users     string
	Active_Roles     string
	Active_Device    string
	Active_Chat      string
	Active_Banner    string
}

type Date struct {
	Date_V string
}

type Form_Date struct {
	Form_Date_V string
}

func GetMenu(url string) (Status_Menu, error) {
	var status Status_Menu
	if url == "/" {
		status.Active_Dashboard = "active"
		status.Active_Customers = ""
		status.Active_Products = ""
		status.Active_Orders = ""
		status.Active_Users = ""
		status.Active_Roles = ""
		status.Active_Device = ""
		status.Active_Chat = ""
		status.Active_Banner = ""
	} else if url == "/customers" {
		status.Active_Dashboard = ""
		status.Active_Customers = "active"
		status.Active_Products = ""
		status.Active_Orders = ""
		status.Active_Users = ""
		status.Active_Roles = ""
		status.Active_Device = ""
		status.Active_Chat = ""
		status.Active_Banner = ""
	} else if url == "/customersaddnew" {
		status.Active_Dashboard = ""
		status.Active_Customers = "active"
		status.Active_Products = ""
		status.Active_Orders = ""
		status.Active_Users = ""
		status.Active_Roles = ""
		status.Active_Device = ""
		status.Active_Chat = ""
		status.Active_Banner = ""
	} else if url == "/customersedit" {
		status.Active_Dashboard = ""
		status.Active_Customers = "active"
		status.Active_Products = ""
		status.Active_Orders = ""
		status.Active_Users = ""
		status.Active_Roles = ""
		status.Active_Device = ""
		status.Active_Chat = ""
		status.Active_Banner = ""
	} else if url == "/products" {
		status.Active_Dashboard = ""
		status.Active_Customers = ""
		status.Active_Products = "active"
		status.Active_Orders = ""
		status.Active_Users = ""
		status.Active_Roles = ""
		status.Active_Device = ""
		status.Active_Chat = ""
		status.Active_Banner = ""
	} else if url == "/productsaddnew" {
		status.Active_Dashboard = ""
		status.Active_Customers = ""
		status.Active_Products = "active"
		status.Active_Orders = ""
		status.Active_Users = ""
		status.Active_Roles = ""
		status.Active_Device = ""
		status.Active_Chat = ""
		status.Active_Banner = ""
	} else if url == "/productsedit" {
		status.Active_Dashboard = ""
		status.Active_Customers = ""
		status.Active_Products = "active"
		status.Active_Orders = ""
		status.Active_Users = ""
		status.Active_Roles = ""
		status.Active_Device = ""
		status.Active_Chat = ""
		status.Active_Banner = ""
	} else if url == "/orders" {
		status.Active_Dashboard = ""
		status.Active_Customers = ""
		status.Active_Products = ""
		status.Active_Orders = "active"
		status.Active_Users = ""
		status.Active_Roles = ""
		status.Active_Device = ""
		status.Active_Chat = ""
		status.Active_Banner = ""
	} else if url == "/ordersdetail" {
		status.Active_Dashboard = ""
		status.Active_Customers = ""
		status.Active_Products = ""
		status.Active_Orders = "active"
		status.Active_Users = ""
		status.Active_Roles = ""
		status.Active_Device = ""
		status.Active_Chat = ""
		status.Active_Banner = ""
	} else if url == "/users" {
		status.Active_Dashboard = ""
		status.Active_Customers = ""
		status.Active_Products = ""
		status.Active_Orders = ""
		status.Active_Users = "active"
		status.Active_Roles = ""
		status.Active_Device = ""
		status.Active_Chat = ""
		status.Active_Banner = ""
	} else if url == "/roles" {
		status.Active_Dashboard = ""
		status.Active_Customers = ""
		status.Active_Products = ""
		status.Active_Orders = ""
		status.Active_Users = ""
		status.Active_Roles = "active"
		status.Active_Device = ""
		status.Active_Chat = ""
		status.Active_Banner = ""
	} else if url == "/device" {
		status.Active_Dashboard = ""
		status.Active_Customers = ""
		status.Active_Products = ""
		status.Active_Orders = ""
		status.Active_Users = ""
		status.Active_Roles = ""
		status.Active_Device = "active"
		status.Active_Chat = ""
		status.Active_Banner = ""
	} else if url == "/chat" {
		status.Active_Dashboard = ""
		status.Active_Customers = ""
		status.Active_Products = ""
		status.Active_Orders = ""
		status.Active_Users = ""
		status.Active_Roles = ""
		status.Active_Device = ""
		status.Active_Chat = "active"
		status.Active_Banner = ""
	} else if url == "/banner" {
		status.Active_Dashboard = ""
		status.Active_Customers = ""
		status.Active_Products = ""
		status.Active_Orders = ""
		status.Active_Users = ""
		status.Active_Roles = ""
		status.Active_Device = ""
		status.Active_Chat = ""
		status.Active_Banner = "active"
	} else {
		status.Active_Dashboard = ""
		status.Active_Customers = ""
		status.Active_Products = ""
		status.Active_Orders = ""
		status.Active_Users = ""
		status.Active_Roles = ""
		status.Active_Device = ""
		status.Active_Chat = ""
		status.Active_Banner = ""
	}

	return status, nil
}

func GetDate() (Date, error) {
	var date Date
	currentTime := time.Now()
	at := currentTime.Format("2006-01-02 15:04:05")
	date.Date_V = at
	fmt.Println(date)
	return date, nil
}

func GetDateParam(d string) (Form_Date, error) {
	var dt Form_Date
	str := d
	t, err := time.Parse(time.RFC3339, str)
	date := t.Format("2006-01-02")
	dt.Form_Date_V = date

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
	fmt.Println(t)
	fmt.Println(date)
	return dt, nil
}
