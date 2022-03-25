package main

//go:generate sqlboiler --wipe psql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ryskit/sqlboiler-sample/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {
	db, err := sql.Open("postgres", `dbname=gettingstarted host=localhost user=postgres password=password sslmode=disable`)
	dieIf(err)

	err = db.Ping()
	dieIf(err)

	fmt.Println("connected")

	u := &models.User{Name: "john"}
	err = u.Insert(context.Background(), db, boil.Infer())
	dieIf(err)

	fmt.Println("user id", u.ID)

	got, err := models.Users().One(context.Background(), db)
	dieIf(err)

	fmt.Println("got user:", got.ID)

	found, err := models.FindUser(context.Background(), db, u.ID)
	dieIf(err)

	fmt.Println("found user", found.ID)

	found.Name = "jane"
	rows, err := found.Update(context.Background(), db, boil.Whitelist(models.UserColumns.Name))
	dieIf(err)

	fmt.Println("updated", rows, "users")

	foundAgain, err := models.Users(qm.Where("name = ?", found.Name)).One(context.Background(), db)
	dieIf(err)

	fmt.Println("found again", foundAgain.ID, foundAgain.Name)

	exists, err := models.UserExists(context.Background(), db, foundAgain.ID)
	dieIf(err)

	fmt.Println("user", foundAgain.ID, "exists?", exists)

	count, err := models.Users().Count(context.Background(), db)
	dieIf(err)

	fmt.Println("there are", count, "users")

	videos, err := got.Videos().All(context.Background(), db)
	dieIf(err)

	fmt.Println("got:", len(videos), "videos")
	for _, v := range videos {
		fmt.Println("  got video:", v.ID, v.Name)
	}
}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}
