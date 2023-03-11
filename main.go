package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type set struct {
	DB *gorm.DB
}

func main() {
	var wg sync.WaitGroup
	s := set{}
	if err := s.getDBConnect(); err != nil {
		panic(err)
	}

	var src *[]Source
	if err := s.DB.Find(&src).Error; err != nil {
		log.Printf("ERRORS: %+v\n", err)
	}

	fmt.Printf(">> %+v\n", len(*src))

	for i, tar := range *src {
		log.Println("Main: Starting worker", i)
		wg.Add(1)
		go s.worker(tar, &wg)
	}

	log.Println("Main: Waiting for workers to finish")
	wg.Wait()
	log.Println("Main: Completed")
}

func (s *set) worker(toTarget Source, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker: Started")
	time.Sleep(time.Second * 2)
	target := &Target{
		UUID: toTarget.UUID,
		Name: toTarget.Name,
	}

	if err := s.DB.Create(&target).Error; err != nil {
		log.Printf("error %+v", err.Error())
	} else {
		log.Printf("ok %+v", *target.Name)
	}
	log.Println("Worker: Finished")
}

func (s *set) getDBConnect() (err error) {
	s.DB, err = gorm.Open(sqlite.Open("dat.sqlite"), &gorm.Config{})
	return err
}
