package database

// import (
// 	"CourseWork/internal/entities"
// 	"CourseWork/internal/process"
// 	"context"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"testing"

// 	"github.com/google/uuid"
// )

// func TestNewFullDataFile(t *testing.T) {
// 	fd, err := NewFullDataFile("test.data")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	log.Println(fd)
// }

// func TestWriteURL(t *testing.T) {
// 	fd, err := NewFullDataFile("test.data")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

// 	var url entities.UrlData

// 	url.Id = uuid.New()
// 	url.FullURL = "//test4.com/"
// 	url.ShortURL = "bitme.com/" + process.GenerateRandomString()

// 	err = fd.WriteURL(ctx, url)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestReadURL(t *testing.T) {
// 	fd, err := NewFullDataFile("test.data")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

// 	var url entities.UrlData

// 	url.ShortURL = "bitme.com/BpLnfgDsc4"

// 	urlData, err := fd.ReadURL(ctx, url)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	log.Println(urlData)
// }

// func TestWriteData(t *testing.T) {
// 	fd, err := NewFullDataFile("test.data")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

// 	id, err := uuid.Parse("b2a06f1f-1ecb-49a8-8c8c-85d3a2e1eca0")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	m := make(map[string]string)
// 	m["NumOfUses"] = "1"

// 	err = fd.WriteData(ctx, id, m)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }
