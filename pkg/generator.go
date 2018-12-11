package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

// Statement is a struct and something to encapsulate the
// string and the type of word that gets filled in
// note that we are using Capital letters for the variables
// this is so we can access them directly
type Statement struct {
	Sentence string // the sentence to substitute a word in
	Dictionary []string    // words that can be used.
}

// nouns is an array assignment with no specific limits
// that means other words can be added here
var nouns = []string{
	"efficiency",
	"competitive advantage",
	"ecosystems",
	"synergy",
	"learning organization",
	"network",
	"social media",
	"revolution",
	"big data",
	"security",
	"internet of things",
	"digital business",
	"data leaders",
	"big data",
	"insight from data",
	"platform",
	"culture"}

var adjectives = []string{
	"digital first",
	"agile",
	"open",
	"innovative",
	"networked",
	"collaborative",
	"cloud based",
	"growth focused",
	"secure",
	"customer focused",
	"disruptive",
	"platform based",
	"sustainable",
	"value added"}

var adverbs = []string {
	"competitively",
	"creatively",
	"effectively",
	"independently",
	"methodically",
	"substantially",
}

// Each statement has a %v which is for string substitution
// in Go there is no tuple substitution like python so the
// whole statement needs to be divided in this way
var statements = []Statement{
	Statement{"Our strategy is %v. ", adjectives},
	Statement{"We will lead a %v ", adjectives},
	Statement{"effort of the market through our use of %v ", nouns},
	Statement{"and %v ", nouns},
	Statement{"to build a %v. ", nouns},
	Statement{"By being both %v ", adjectives},
	Statement{"and %v, ", adjectives},
	Statement{"our %v ", adjectives},
	Statement{"approach will drive %v throughout the organization. ", nouns},
	Statement{"Synergies between our %v ", nouns},
	Statement{"and %v ", nouns},
	Statement{"will enable us to %v ", adverbs},
	Statement{"capture the upside by becoming %v ", adjectives},
	Statement{"in a %v world. ", nouns},
	Statement{"These transformations combined with %v ", nouns},
	Statement{"due to our %v ", nouns},
	Statement{"will create a %v ", nouns},
	Statement{"through %v ", nouns},
	Statement{"and %v.", nouns}}

// Generate produces a random business statedgy statement.
// Note: any of the functions called from the packages imported
// has to start with a Capital letter this is because of the rule
// that only functions starting with a capital letter are 'exported'
// so you can call them here from another package.
// It's not really the same as public and private but more like direct
// and indirect reference or use
// https://www.ardanlabs.com/blog/2014/03/exportedunexported-identifiers-in-go.html
func Generate() string {

	// implicit variable assignment
	// not available outside the function
	statement := ""

	// rand uses a global package called Source
	// https://golang.org/pkg/math/rand/#Source
	// The source has to be initialized by some seed
	// the seed needs to change for each Generation
	// or else the same sequence will be produced.
	rand.Seed(time.Now().UnixNano())
	for i := range statements {
		s := statements[i]
		RandIndex := rand.Intn(len(s.Dictionary))
		word := s.Dictionary[RandIndex]
		statement += fmt.Sprintf(s.Sentence, word)
	}

	return statement
}
