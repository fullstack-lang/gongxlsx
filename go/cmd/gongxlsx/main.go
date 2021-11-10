package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongxlsx"
	"github.com/fullstack-lang/gongxlsx/go/controllers"
	"github.com/fullstack-lang/gongxlsx/go/models"
	"github.com/fullstack-lang/gongxlsx/go/orm"
)

var (
	logDBFlag   = flag.Bool("logDB", false, "log mode for db")
	logGINFlag  = flag.Bool("logGIN", false, "log mode for gin")
	compareFlag = flag.String("compare", "sampleFile", "compare to the other file")
)

func main() {

	log.SetPrefix("gongxlsx: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	// setup GORM
	db := orm.SetupModels(*logDBFlag, ":memory:")
	dbDB, err := db.DB()

	// since gongsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	dbDB.SetMaxOpenConns(1)

	controllers.RegisterControllers(r)

	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(gongxlsx.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	nbArgs := flag.Args()
	if len(nbArgs) < 1 {
		log.Panicf("there should be at least one file argument")
	}

	file := new(models.XLFile).Stage()
	file.Open(flag.Arg(0))

	if *compareFlag == "sampleFile" {
		log.Println("no file to compare")
	} else {
		fileToCompare := new(models.XLFile).Stage()
		fileToCompare.Open(*compareFlag)

	}

	models.Stage.Commit()

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
