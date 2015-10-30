package server

import (
	"encoding/xml"
	"errors"
	"fmt"
	"sync"
)

var (
	ErrAlreadyExists = errors.New("album already exists")
)

// The DB interface defines methods to manipulate the albums.
type DB interface {
	Get(id int) *Album
	GetAll() []*Album
	Find(band, title string, year int) []*Album
	Add(a *Album) (int, error)
	Update(a *Album) error
	Delete(id int)
}

// Thread-safe in-memory map of albums.
type albumsDB struct {
	sync.RWMutex
	m   map[int]*Album
	seq int
}

// The one and only database instance.
var DBInstance DB

func init() {
	DBInstance = &albumsDB{
		m: make(map[int]*Album),
	}
	// Fill the database
	DBInstance.Add(&Album{Id: 1, Band: "Slayer", Title: "Reign In Blood", Year: 1986})
	DBInstance.Add(&Album{Id: 2, Band: "Slayer", Title: "Seasons In The Abyss", Year: 1990})
	DBInstance.Add(&Album{Id: 3, Band: "Bruce Springsteen", Title: "Born To Run", Year: 1975})
}

// GetAll returns all albums from the database.
func (DBInstance *albumsDB) GetAll() []*Album {
	DBInstance.RLock()
	defer DBInstance.RUnlock()
	if len(DBInstance.m) == 0 {
		return nil
	}
	ar := make([]*Album, len(DBInstance.m))
	i := 0
	for _, v := range DBInstance.m {
		ar[i] = v
		i++
	}
	return ar
}

// Find returns albums that match the search criteria.
func (DBInstance *albumsDB) Find(band, title string, year int) []*Album {
	DBInstance.RLock()
	defer DBInstance.RUnlock()
	var res []*Album
	for _, v := range DBInstance.m {
		if v.Band == band || band == "" {
			if v.Title == title || title == "" {
				if v.Year == year || year == 0 {
					res = append(res, v)
				}
			}
		}
	}
	return res
}

// Get returns the album identified by the id, or nil.
func (DBInstance *albumsDB) Get(id int) *Album {
	DBInstance.RLock()
	defer DBInstance.RUnlock()
	return DBInstance.m[id]
}

// Add creates a new album and returns its id, or an error.
func (DBInstance *albumsDB) Add(a *Album) (int, error) {
	DBInstance.Lock()
	defer DBInstance.Unlock()
	// Return an error if band-title already exists
	if !DBInstance.isUnique(a) {
		return 0, ErrAlreadyExists
	}
	// Get the unique ID
	DBInstance.seq++
	a.Id = DBInstance.seq
	// Store
	DBInstance.m[a.Id] = a
	return a.Id, nil
}

// Update changes the album identified by the id. It returns an error if the
// updated album is a duplicate.
func (DBInstance *albumsDB) Update(a *Album) error {
	DBInstance.Lock()
	defer DBInstance.Unlock()
	if !DBInstance.isUnique(a) {
		return ErrAlreadyExists
	}
	DBInstance.m[a.Id] = a
	return nil
}

// Delete removes the album identified by the id from the database. It is a no-op
// if the id does not exist.
func (DBInstance *albumsDB) Delete(id int) {
	DBInstance.Lock()
	defer DBInstance.Unlock()
	delete(DBInstance.m, id)
}

// Checks if the album already exists in the database, based on the Band and Title
// fields.
func (DBInstance *albumsDB) isUnique(a *Album) bool {
	for _, v := range DBInstance.m {
		if v.Band == a.Band && v.Title == a.Title && v.Id != a.Id {
			return false
		}
	}
	return true
}

// The Album data structure, serializable in JSON, XML and text using the Stringer interface.
type Album struct {
	XMLName xml.Name `json:"-" xml:"album"`
	Id      int      `json:"id" xml:"id,attr"`
	Band    string   `json:"band" xml:"band"`
	Title   string   `json:"title" xml:"title"`
	Year    int      `json:"year" xml:"year"`
}

func (a *Album) String() string {
	return fmt.Sprintf("%s - %s (%d)", a.Band, a.Title, a.Year)
}
