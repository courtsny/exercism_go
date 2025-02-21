package main

import "fmt"

type ElectionResult struct {
	Name  string
	Votes int
}

func NewVoteCounter(initial_votes int) *int {
	var vote_counter = &initial_votes
	return vote_counter
}

func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

func IncrementVoteCount(counter *int, increment_amount int) bool {
	if counter == nil {
		return false
	}
	*counter += increment_amount
	return true
}

func NewElectionResult(name string, vote_count int) *ElectionResult {
	var result = ElectionResult{Name: name, Votes: vote_count}
	var result_ptr = &result
	return result_ptr
}

func DisplayResult(results *ElectionResult) {
	fmt.Printf("%s won with %d vote(s)\n", results.Name, results.Votes)
}

func DecrementVotesOfCandidate(result_map map[string]int, name string) {
	result_map[name] -= 1
}

func PointersTest() {
	init_votes := 2
	var init_votes_ptr, nil_vote_ptr *int
	init_votes_ptr = NewVoteCounter(init_votes)
	fmt.Printf("init votes = %d, ptr = %d\n", init_votes, *init_votes_ptr)
	more_votes := 3
	more_votes_ptr := &more_votes
	fmt.Printf("vote count = %d, nil vote count = %d\n", VoteCount(more_votes_ptr), VoteCount(nil_vote_ptr))
	fmt.Printf("vote count before increment = %d, vote count success after adding %d = %t, after increment = %d\n",
		VoteCount(init_votes_ptr), *&more_votes, IncrementVoteCount(init_votes_ptr, *&more_votes), VoteCount(init_votes_ptr))
	var result_ptr = NewElectionResult("Peter", 3)
	DisplayResult(result_ptr)
	var result_map = map[string]int{
		"Mary": 10,
		"John": 51,
	}
	fmt.Printf("mary had %d votes before the recount\n", result_map["Mary"])
	DecrementVotesOfCandidate(result_map, "Mary")
	fmt.Printf("mary had %d votes after the recount\n", result_map["Mary"])
}
