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
	router.GET("/communitypool", getCommunityPool)
	router.GET("/proposal", getProposals)
	router.GET("/proposal/filtered", getProposalsFiltered)
	router.GET("/totalWithdrawnAmount", getTotalWithdrawnAmount)

	router.Run("localhost:8080")

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
	var result ResponseProposalFiltered
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "ERROR"})
	}

	sum := big.NewInt(0)
	z := &big.Int{}

	for _, rec := range result.PropsalFiltered {
		for _, loop2 := range rec.TotalDeposits {
			z, success := z.SetString(loop2.Amount, 10)
			if !success {
				panic("Error converting ... ")
			}
			fmt.Println("z =", z)
			sum = sum.Add(sum, z)
		}

	}
	total := sum.String()
	c.IndentedJSON(http.StatusOK, total)
}


func check() gin.HandlerFunc {
	return func(c *gin.Context) {
		validPath := map[string]bool{
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
