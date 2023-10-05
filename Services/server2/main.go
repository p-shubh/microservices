package main

import (
	databaseconnectionServer2 "server2/DatabaseConnection"
	routesServer2 "server2/Routes"
)

func main() {
	databaseconnectionServer2.Databaseconnectionserver2()
	routesServer2.RoutesServer2()
}
