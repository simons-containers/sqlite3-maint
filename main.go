package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type multiFlag []string

func (m *multiFlag) String() string {
	return strings.Join(*m, ",")
}

func (m *multiFlag) Set(v string) error {
	*m = append(*m, v)
	return nil
}

func maintain(path string) error {
	db, err := sql.Open("sqlite3", "file:"+path+"?mode=rw&_busy_timeout=5000")
	if err != nil {
		return err
	}
	defer db.Close()
	fmt.Fprintf(os.Stdout, "running VACUUM/ANALYZE maintenance on %s\n", path)


	_, err = db.Exec("VACUUM; ANALYZE;")
	fmt.Fprintf(os.Stdout, "completed maintenance on %s\n", path)
	return err
}

func collectDBs(files []string, dirs []string) ([]string, error) {
	set := map[string]struct{}{}

	for _, f := range files {
		set[f] = struct{}{}
	}

	for _, d := range dirs {
		entries, err := os.ReadDir(d)
		if err != nil {
			return nil, err
		}

		for _, e := range entries {
			if e.IsDir() {
				continue
			}
			if strings.HasSuffix(e.Name(), ".db") {
				set[filepath.Join(d, e.Name())] = struct{}{}
			}
		}
	}

	var out []string
	for k := range set {
		out = append(out, k)
	}

	return out, nil
}

func main() {
	var dbFlags multiFlag
	var dbDirFlags multiFlag

	flag.Var(&dbFlags, "db", "database file (repeatable)")
	flag.Var(&dbDirFlags, "dbdir", "directory containing .db files (repeatable)")
	flag.Parse()

	dbs, err := collectDBs(dbFlags, dbDirFlags)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(dbs) == 0 {
		return
	}

	workers := runtime.NumCPU()
	sem := make(chan struct{}, workers)

	var wg sync.WaitGroup
	var failed bool
	var mu sync.Mutex

	for _, path := range dbs {
		wg.Add(1)

		go func(p string) {
			defer wg.Done()

			sem <- struct{}{}
			err := maintain(p)
			<-sem

			if err != nil {
				mu.Lock()
				failed = true
				mu.Unlock()
				fmt.Fprintf(os.Stderr, "error: %s: %v\n", p, err)
			}
		}(path)
	}

	wg.Wait()

	if failed {
		os.Exit(1)
	}
}