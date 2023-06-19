package dbaccess

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// func ConnectToDb() (context.Context, *mongo.Database, *mongo.Collection) {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)

// 	database := client.Database("UserService_BeMyPet")
// 	usersCollection := database.Collection("user")
// 	return ctx, database, usersCollection
// }

func ConnectToDb() *sql.DB {


	db, err := sql.Open("mysql", "ivi01xiryhw2nbk11rgb:pscale_pw_hC6omT7ldcKl9uiCdkdthnGd8Pi9zhTobmsPV1oQpKt@tcp(aws.connect.psdb.cloud)/bemypet?tls=true")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v", err)
	}
	log.Println("Successfully connected to PlanetScale!")
	return db
}
