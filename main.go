package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jianhan/gocms/configs"
	cfgreader "github.com/jianhan/pkg/configs"
	micro "github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
)

func main() {
	// serviceConfigs, err := cfgreader.NewReader(os.Getenv("ENV")).Read()
	serviceConfigs, err := cfgreader.NewReader("development").Read()
	if err != nil {
		panic(fmt.Sprintf("error while reading configurations: %s", err.Error()))
	}
	spew.Dump(serviceConfigs)
	//db, err := db.GetDB()
	//if err != nil {
	//	panic(fmt.Sprintf("error while connect to database: %s", err.Error()))
	//}
	//defer db.Close()
	srv := micro.NewService(
		micro.Name(serviceConfigs.Name),
		micro.RegisterTTL(time.Duration(serviceConfigs.RegisterTTL)*time.Second),
		micro.RegisterInterval(time.Duration(serviceConfigs.RegisterInterval)*10),
		micro.Version(serviceConfigs.Version),
		micro.Metadata(serviceConfigs.Metadata),
	)

	srv.Init()
	// pcourses.RegisterCoursesHandler(
	// 	srv.Server(),
	// 	&handlers.Courses{
	// 		CourseRepository: repositories.NewCourseRepository(db),
	// 	},
	// )
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
