package main

import (
	"fmt"
	"sort"

	"github.com/hashicorp/go-version"
)

func main() {
	versions := make([]*version.Version, 0, 2)
	v, _ := version.NewVersion("2022.03.03-0303")
	versions = append(versions, v)
	v, _ = version.NewVersion("2022.04.04-0404")
	versions = append(versions, v)

	fmt.Printf("%s.Compare(%s): %d\n", versions[0], versions[1], versions[0].Compare(versions[1]))
	fmt.Printf("%s.Compare(%s): %d\n", versions[1], versions[0], versions[1].Compare(versions[0]))

	fmt.Println("before")
	fmt.Println(versions)
	sort.Sort(sort.Reverse(version.Collection(versions)))
	fmt.Println("after")
	fmt.Println(versions)
}
