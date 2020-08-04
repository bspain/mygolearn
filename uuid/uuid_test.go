package uuid_test

import (
	"os"
	"bufio"
	"testing"

	"github.com/satori/go.uuid"
	"github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("satori go.uuid", func() {
		g.It("should parse Jenkins generated UUID's", func() {
			file, _ := os.Open("uuid-test-data")

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)

			for scanner.Scan() {
				_, err := uuid.FromString(scanner.Text())
				g.Assert(err).Eql(nil)	
			}
		})
	})
}

