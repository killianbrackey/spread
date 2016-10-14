package main

import (
	"fmt"

	// "sezzle/gorm-playground/"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {

	dbScript := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", DBUser, DBPassword, DBHost, DBPort, DBName)
	dbConn, err := gorm.Open(DBDriver, dbScript)

	if err != nil {
		fmt.Println(err)
	}

	// init
	dbConn.DB()
	dbConn.DB().Ping()
	dbConn.DB().SetMaxIdleConns(10)
	dbConn.DB().SetMaxOpenConns(100)

	//BelongsTo Example
	type Profile struct {
		gorm.Model
		Name string
	}
	type User struct {
		gorm.Model
		Profile   Profile
		ProfileID int
	}
	dbConn.AutoMigrate(&Profile{})
	dbConn.AutoMigrate(&User{})

	profile := new(Profile)
	profile.Name = "new profile"
	dbConn.Create(profile)

	user := new(User)
	user.Profile.Name = "new user profile"
	dbConn.Create(user)

	//HasOne
	type CreditCard struct {
		gorm.Model
		UserWithCCID uint
		Number       string
	}
	type UserWithCC struct {
		gorm.Model
		CreditCard CreditCard
	}
	dbConn.AutoMigrate(&CreditCard{})
	dbConn.AutoMigrate(&UserWithCC{})

	creditCard := new(CreditCard)
	creditCard.Number = "3232322323223232"
	dbConn.Create(creditCard)

	userWithCC := new(UserWithCC)
	userWithCC.CreditCard = CreditCard{
		Number: "1111111111",
	}
	dbConn.Create(userWithCC)

	//HasMany
	type Email struct {
		gorm.Model
		Email           string
		UserWithEmailID uint
	}
	type UserWithEmail struct {
		gorm.Model
		Emails []Email
	}
	dbConn.AutoMigrate(&Email{})
	dbConn.AutoMigrate(&UserWithEmail{})
	email := new(Email)
	email.Email = "nouser@nouser.com"
	dbConn.Create(email)
	userWithEmail := new(UserWithEmail)
	userWithEmail.Emails = []Email{
		Email{Email: "user@email.com"},
	}
	dbConn.Create(userWithEmail)

	//Many to Many
	type Language struct {
		gorm.Model
		Name string
	}
	type Person struct {
		gorm.Model
		Languages []Language `gorm:"many2many:user_languages;"`
	}
	dbConn.AutoMigrate(&Language{})
	dbConn.AutoMigrate(&Person{})

	language := new(Language)
	language.Name = "English"
	dbConn.Create(language)

	person := new(Person)
	person.Languages = []Language{
		Language{
			Name: "Spanish",
		},
		Language{
			Name: "Arabic",
		},
	}
	dbConn.Create(person)

	// Polymorphic Test
	type Toy struct {
		ID        int
		Name      string
		OwnerID   int
		OwnerType string
	}
	type Cat struct {
		ID   int
		Name string
		Toy  Toy `gorm:"polymorphic:Owner;"`
	}
	type Dog struct {
		ID   int
		Name string
		Toy  Toy `gorm:"polymorphic:Owner;"`
	}

	dbConn.AutoMigrate(&Toy{})
	dbConn.AutoMigrate(&Cat{})
	dbConn.AutoMigrate(&Dog{})

	toy := new(Toy)
	toy.Name = "test"
	dbConn.Create(toy)

	cat := new(Cat)
	cat.Toy = Toy{
		Name: "LeoToy",
	}
	cat.Name = "Leo"
	dbConn.Create(cat)

	dog := new(Dog)
	dog.Name = "Pooch"
	dbConn.Create(dog)
}
