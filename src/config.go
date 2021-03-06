package mboximporter

// App configuration.
type Config struct {
    MongoURI string // the mongo connection string
    DBName string // db to use in MongoDB
    Filename string // filename of the file to import
    Count int // Number maximum of mails to import
    Workers int // Number maximum of workers
    Concurrency int // Max amount of messages in the pool to process
}
