package main

import (
	databaseconnection "server1/DatabaseConnection"
	routes "server1/Routes"
)

func main() {
	databaseconnection.Databaseconnection()
	routes.RoutesServer1()
}
