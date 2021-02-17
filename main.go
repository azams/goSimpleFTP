package main
import (
        "os"
        "log"
        filedriver "github.com/goftp/file-driver"
        "github.com/goftp/server"
)
func main() {
        root, err := os.Getwd()
        if err != nil {
                panic(err)
        }
        user := "horangi"
        pass := "horangutan"
        port := 7090
        host := "0.0.0.0"
        factory := &filedriver.FileDriverFactory{
                RootPath: root,
                Perm:     server.NewSimplePerm("user", "group"),
        }
        opts := &server.ServerOpts{
                Factory:  factory,
                Port:     port,
                Hostname: host,
                Auth:     &server.SimpleAuth{Name: user, Password: pass},
        }
        log.Printf("Starting ftp server on %v:%v", opts.Hostname, opts.Port)
        log.Printf("Username %v, Password %v", user, pass)
        server := server.NewServer(opts)
        err = server.ListenAndServe()
        if err != nil {
                log.Fatal("Error starting server:", err)
        }
}
