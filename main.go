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
	WG sync.WaitGroup
}

func main() {
	// var wg sync.WaitGroup
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
		s.WG.Add(1)
		go s.worker(tar)
	}

	log.Println("Main: Waiting for workers to finish")
	s.WG.Wait()
	log.Println("Main: Completed")
}

func (s *set) worker(toTarget Source) {
	defer s.WG.Done()

	log.Println("Worker: Started")
	time.Sleep(time.Second * 2)
	target := &Target{
		UUID: toTarget.UUID,
		Name: toTarget.Name,
	}

	if err := s.DB.Create(&target).Error; err != nil {
		log.Printf("error %+v", err.Error())
	}
	log.Println("Worker: Finished")
}

func (s *set) getDBConnect() (err error) {
	s.DB, err = gorm.Open(sqlite.Open("dat.sqlite"), &gorm.Config{})
	return err
}
