# OCM-Meta
Go server for to provide a list of OnChainMonkey metas from [Metagood OCM Repo](https://github.com/metagood/OnChainMonkeyData)


## Installation

```bash
$ go get -u github.com/themobilecoder/ocm-meta
```


## Usage

```go
package main

import (
	"fmt"

	"github.com/themobilecoder/ocm-meta/meta"
)

func main() {
	monkeys := meta.GetOnChainMonkeys()

	fmt.Printf("%+v\n", monkeys[4642-1])
	fmt.Printf("%+v\n", monkeys[1179-1])
	fmt.Printf("%+v\n", monkeys[5961-1])
	fmt.Printf("%+v\n", monkeys[8059-1])
	fmt.Printf("%+v\n", monkeys[7753-1])
	fmt.Printf("%+v\n", monkeys[2301-1])
	fmt.Printf("%+v\n", monkeys[965-1])
}

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)