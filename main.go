package main


import(
	"fmt"
	"github.com/toorop/go-bittrex"
)


const (
	API_KEY    = ""
	API_SECRET = ""
	MIN_ETH_THRESH = 0.1
)


func main() {

	B := bittrex.New( API_KEY, API_SECRET )

	EthBalance, Err := B.GetBalance( "ETH" )
	
	if Err == nil {

		fmt.Printf( "ETH balance: %f\n", EthBalance.Available )

		if EthBalance.Available >= MIN_ETH_THRESH {

			B.Withdraw( "", "ETH", EthBalance.Available )
		}
	}
}