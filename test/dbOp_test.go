package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/boltdb/bolt"
	service "github.com/liuyh73/GraphQL-Service"
)

func TestGetFilmByID(t *testing.T) {
	film2 := &service.Film{}
	db, _ := bolt.Open("../data/data.db", 0600, nil)
	defer db.Close()
	err, film1 := service.GetFilmByID("2", db)
	if err != nil {
		t.Error("测试失败")
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Film"))
		json.Unmarshal(b.Get([]byte("2")), film2)
		return nil
	})
	if film2.ID == film1.ID && film2.Title == film1.Title && *(film2.EpisodeID) == *(film1.EpisodeID) {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}

func TestGetPeopleByID(t *testing.T) {
	people2 := &service.People{}
	db, _ := bolt.Open("../data/data.db", 0600, nil)
	defer db.Close()
	err, people1 := service.GetPeopleByID("1", db)
	if err != nil {
		t.Error("测试失败")
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("People"))
		json.Unmarshal(b.Get([]byte("1")), people2)
		return nil
	})
	fmt.Println(people2.ID, people1.ID, people2.Name, people1.Name)
	if people2.ID == people1.ID && people2.Name == people1.Name {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}

func TestGetPlanetByID(t *testing.T) {
	Planet2 := &service.Planet{}
	db, _ := bolt.Open("../data/data.db", 0600, nil)
	defer db.Close()
	err, Planet1 := service.GetPlanetByID("1", db)
	if err != nil {
		t.Error("测试失败")
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Planet"))
		json.Unmarshal(b.Get([]byte("1")), Planet2)
		return nil
	})
	fmt.Println(Planet2.ID, Planet1.ID, Planet2.Name, Planet1.Name)
	if Planet2.ID == Planet1.ID && Planet2.Name == Planet1.Name {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}

func TestGetSpecieByID(t *testing.T) {
	Specie2 := &service.Specie{}
	db, _ := bolt.Open("../data/data.db", 0600, nil)
	defer db.Close()
	err, Specie1 := service.GetSpeciesByID("1", db)
	if err != nil {
		t.Error("测试失败")
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Specie"))
		json.Unmarshal(b.Get([]byte("1")), Specie2)
		return nil
	})
	fmt.Println(Specie2.ID, Specie1.ID, Specie2.Name, Specie1.Name)
	if Specie2.ID == Specie1.ID && Specie2.Name == Specie1.Name {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}

func TestGetStarshipByID(t *testing.T) {
	Starship2 := &service.Starship{}
	db, _ := bolt.Open("../data/data.db", 0600, nil)
	defer db.Close()
	err, Starship1 := service.GetStarshipByID("1", db)
	if err != nil {
		t.Error("测试失败")
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Starship"))
		json.Unmarshal(b.Get([]byte("1")), Starship2)
		return nil
	})
	fmt.Println(Starship2.ID, Starship1.ID, Starship2.Name, Starship1.Name)
	if Starship2.ID == Starship1.ID && Starship2.Name == Starship1.Name {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}
func TestGetVehicleByID(t *testing.T) {
	Vehicle2 := &service.Vehicle{}
	db, _ := bolt.Open("../data/data.db", 0600, nil)
	defer db.Close()
	err, Vehicle1 := service.GetVehicleByID("1", db)
	if err != nil {
		t.Error("测试失败")
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Vehicle"))
		json.Unmarshal(b.Get([]byte("1")), Vehicle2)
		return nil
	})
	fmt.Println(Vehicle2.ID, Vehicle1.ID, Vehicle2.Name, Vehicle1.Name)
	if Vehicle2.ID == Vehicle1.ID && Vehicle2.Name == Vehicle1.Name {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}
