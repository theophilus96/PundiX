// Question: Using Golang, create a simple REST API that will query the following information from our mainnet blockchain.

// Query the outstanding amount in the community-spend pool

// 		You may find the answer by running this command in the CLI:

// 		fxcored query distribution community-pool

// Query all the withdrawals from the community-spend pool

// Information (filtered) should include:

// Amount withdrawn

// Timestamp

// Receiver address

// Sum of all withdrawn amount

// 	You may find the answer by running this command in the CLI:

// fxcored q gov proposals

// By inputting the HTTPS request in a browser eg. http://localhost:3000/query/community-pool/outstanding will return a json in the browser with the output, just display the information, no css required.

// Ensure proper error handling if the wrong ‘input’ field is called

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(check())
	router.GET("/albums", getAlbums)
	router.GET("/communitypool", getCommunityPool)
	router.GET("/proposal", getProposals)
	router.GET("/proposal/filtered", getProposalsFiltered)
	router.GET("/totalWithdrawnAmount", getTotalWithdrawnAmount)

	router.Run("localhost:8080")

}

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type CommunityPool struct {
	Pool []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"pool"`
}

type ResponseProposal struct {
	Proposals  Proposals
	Pagination Pagination `json:"pagination"`
}

type ResponseProposalFiltered struct {
	PropsalFiltered PropsalFiltered `json:"proposals,omitempty"`
	Pagination      Pagination      `json:"-"`
}

type Proposals []struct {
	ProposalID       string  `json:"proposal_id"`
	Content          Content `json:"content,omitempty"`
	Status           string  `json:"status"`
	FinalTallyResult struct {
		Yes        string `json:"yes"`
		Abstain    string `json:"abstain"`
		No         string `json:"no"`
		NoWithVeto string `json:"no_with_veto"`
	} `json:"final_tally_result"`
	SubmitTime     time.Time `json:"submit_time"`
	DepositEndTime time.Time `json:"deposit_end_time"`
	TotalDeposit   []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"total_deposit"`
	VotingStartTime time.Time `json:"voting_start_time"`
	VotingEndTime   time.Time `json:"voting_end_time"`
	// Content2        Content2  `json:"proposals"`
}

type PropsalFiltered []struct {
	ProposalID       string          `json:"proposal_id"`
	ContentFiltered  ContentFiltered `json:"content,omitempty"`
	Status           string          `json:"-"`
	FinalTallyResult struct {
		Yes        string `json:"yes"`
		Abstain    string `json:"abstain"`
		No         string `json:"no"`
		NoWithVeto string `json:"no_with_veto"`
	} `json:"-"`
	SubmitTime      time.Time     `json:"submit_time"`
	DepositEndTime  time.Time     `json:"-"`
	TotalDeposits   TotalDeposits `json:"total_deposit"`
	VotingStartTime time.Time     `json:"-"`
	VotingEndTime   time.Time     `json:"-"`
}

type TotalDeposits []TotalDeposit

type TotalDeposit struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type Pagination struct {
	NextKey interface{} `json:"next_key"`
	Total   string      `json:"total"`
}

// type Content2 struct {
// 	Type        string `json:"@type"`
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	Recipient   string `json:"recipient"`
// 	Amount      []struct {
// 		Denom  string `json:"denom"`
// 		Amount string `json:"amount"`
// 	} `json:"amount"`
// }

type Content struct {
	Type        string `json:"@type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Params      struct {
		GravityID                         string   `json:"gravity_id"`
		AverageBlockTime                  string   `json:"average_block_time"`
		ExternalBatchTimeout              string   `json:"external_batch_timeout"`
		AverageExternalBlockTime          string   `json:"average_external_block_time"`
		SignedWindow                      string   `json:"signed_window"`
		SlashFraction                     string   `json:"slash_fraction"`
		OracleSetUpdatePowerChangePercent string   `json:"oracle_set_update_power_change_percent"`
		IbcTransferTimeoutHeight          string   `json:"ibc_transfer_timeout_height"`
		Oracles                           []string `json:"oracles"`
		DepositThreshold                  struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"deposit_threshold"`
	} `json:"params"`
	ChainName string `json:"chain_name"`
}

type ContentFiltered struct {
	Type        string `json:"@type"`
	Title       string `json:"title"`
	Description string `json:"-"`
	Recipient   string `json:"recipient,omitempty"`
	Params      struct {
		GravityID                         string   `json:"gravity_id"`
		AverageBlockTime                  string   `json:"average_block_time"`
		ExternalBatchTimeout              string   `json:"external_batch_timeout"`
		AverageExternalBlockTime          string   `json:"average_external_block_time"`
		SignedWindow                      string   `json:"signed_window"`
		SlashFraction                     string   `json:"slash_fraction"`
		OracleSetUpdatePowerChangePercent string   `json:"oracle_set_update_power_change_percent"`
		IbcTransferTimeoutHeight          string   `json:"ibc_transfer_timeout_height"`
		Oracles                           []string `json:"oracles"`
		DepositThreshold                  struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"deposit_threshold"`
	} `json:"-"`
	ChainName string `json:"-"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getCommunityPool(c *gin.Context) {

	url := "http://35.225.177.4:1317/cosmos/distribution/v1beta1/community_pool"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// snippet only
	var result CommunityPool
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(result)
	c.IndentedJSON(http.StatusOK, result)
}

func getProposals(c *gin.Context) {

	url := "http://35.225.177.4:1317/cosmos/gov/v1beta1/proposals"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))
	var result ResponseProposal
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(result)
	c.IndentedJSON(http.StatusOK, result)
}

func getProposalsFiltered(c *gin.Context) {

	url := "http://35.225.177.4:1317/cosmos/gov/v1beta1/proposals"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))
	var result ResponseProposalFiltered
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(result)
	c.IndentedJSON(http.StatusOK, result)
}

func getTotalWithdrawnAmount(c *gin.Context) {

	url := "http://35.225.177.4:1317/cosmos/gov/v1beta1/proposals"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))
	var result ResponseProposalFiltered
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "ERROR"})
	}
	// fmt.Println(result)

	var list []string

	for _, rec := range result.PropsalFiltered {
		// fmt.Println(rec.TotalDeposits.WithdrawnTotal())
		list = append(list, rec.TotalDeposits.WithdrawnTotal())
	}

	fmt.Println("list", list)

	sum := big.NewInt(0)
	for i := 0; i < len(list); i++ {
		z := &big.Int{}
		z, success := z.SetString(list[i], 10)
		if !success {
			panic("Error converting ... ")
		}
		sum = sum.Add(sum, z)
	}

	total := sum.String()

	c.IndentedJSON(http.StatusOK, total)

}

func (u TotalDeposits) WithdrawnTotal() string {
	var list []string

	for _, TotalDeposit := range u {
		list = append(list, TotalDeposit.Amount)
	}

	fmt.Println("list", list)

	sum := big.NewInt(0)
	for i := 0; i < len(list); i++ {
		z := &big.Int{}
		z, success := z.SetString(list[i], 10)
		if !success {
			panic("Error converting ... ")
		}
		sum = sum.Add(sum, z)
	}

	total := sum.String()

	// for _, TotalDeposit := range u {
	// 	intVar, err := strconv.ParseInt(TotalDeposit.Amount, 10, 64)
	// 	if err != nil {
	// 		// handle error
	// 		fmt.Println(err)
	// 	}
	// 	total += intVar
	// }
	return total
}

func check() gin.HandlerFunc {
	return func(c *gin.Context) {
		validPath := map[string]bool{
			"/albums":               true,
			"/communitypool":        true,
			"/proposal":             true,
			"/proposal/filtered":    true,
			"/totalWithdrawnAmount": true,
		}

		if !validPath[c.Request.URL.Path] {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "page not found"})
		}
	}
}
